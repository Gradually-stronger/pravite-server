package boot

import (
	"context"
	redisLib "github.com/go-redis/redis/v8"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go.uber.org/dig"
	"gorm.io/gorm"
	"gxt-api-frame/app/bll/impl"
	"gxt-api-frame/library/gplus"
	"gxt-api-frame/library/logger"
	"gxt-api-frame/library/middleware"
	"gxt-api-frame/library/redis"
	"gxt-api-frame/library/utils"
	"gxt-api-frame/router"
	"os"
)

// VERSION 定义应用版本号
const VERSION = "1.0.0"

func init() {
	// 初始化logger
	logger.Init(g.Cfg().GetString("common.run_mode"))
	logger.SetVersion(VERSION)
	logger.SetTraceIdFunc(utils.NewTraceID)
	ctx := logger.NewTraceIDContext(context.Background(), utils.NewTraceID())
	Init(ctx)
}

// 初始化App,
// TODO: 返回释放回调，暂时没调用
func Init(ctx context.Context) func() {
	logger.Printf(ctx, "服务启动，运行模式：%s，版本号：%s，进程号：%d", g.Cfg().Get("common.run_mode"), VERSION, os.Getpid())
	// 初始化依赖注入容器
	container, call := buildContainer(ctx)
	// 初始化路由注册
	s := g.Server()
	// 每个请求生成新的追踪Id,如果上下文件中没有trace-id
	s.Use(middleware.TraceIdMiddleware())
	// 统一处理内部错误
	s.Use(func(r *ghttp.Request) {
		r.Middleware.Next()
		if err := r.GetError(); err != nil {
			gplus.ResError(r, err)
		}
	})
	router.InitRouters(s, container)
	return func() {
		if call != nil {
			call()
		}
	}

}

// 初始化redis
func initRedis(ctx context.Context, container *dig.Container) func() {
	addr := g.Cfg().GetString("redis.addr")
	password := g.Cfg().GetString("redis.password")
	db := g.Cfg().GetInt("redis.db")
	redisCli := redis.Init(ctx, addr, password, db)
	logger.Printf(ctx, "REDIS初始化成功，当前服务器[%s]", addr)
	// 注入redis client
	_ = container.Provide(func() *redisLib.Client {
		return redisCli
	})
	return func() {
		_ = redisCli.Close()
	}
}

// 初始化存储，目前只初始化gorm
func initStore(ctx context.Context, container *dig.Container) (func(), error) {
	var storeCall func()
	db, err := initGorm()
	if err != nil {
		return storeCall, err
	}
	// 如果自动映射数据表
	if g.Cfg().GetBool("gorm.enable_auto_migrate") {
		err = autoMigrate(db)
		if err != nil {
			return storeCall, err
		}
	}
	// 注入DB
	_ = container.Provide(func() *gorm.DB { return db })
	// 注入model接口
	_ = InjectModel(container)
	storeCall = func() {
		sqlDb, _ := db.DB()
		_ = sqlDb.Close()
	}
	logger.Printf(ctx, "MYSQL初始化成功, 服务器[%s], 数据库[%s]",
		g.Cfg().GetString("mysql.host"),
		g.Cfg().GetString("mysql.db_name"))
	return storeCall, nil
}

// 构建依赖注入容器
func buildContainer(ctx context.Context) (*dig.Container, func()) {
	container := dig.New()
	// 初始化存储模块
	storeCall, err := initStore(ctx, container)
	if err != nil {
		panic(err)
	}
	// 初始化redis模块
	var redisCall func()
	if g.Cfg().GetBool("redis.enable") {
		redisCall = initRedis(ctx, container)
	}
	// 注入bll
	impl.Inject(container)
	return container, func() {
		if storeCall != nil {
			storeCall()
		}
		if redisCall != nil {
			redisCall()
		}
	}
}
