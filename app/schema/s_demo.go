package schema

import "time"

// Demo demo对象
type Demo struct {
	RecordID  string    `json:"record_id"`                       // 记录ID
	Code      string    `json:"code" v:"required#顶过"`            // 编号
	Name      string    `json:"name" v:"required"`               // 名称
	Memo      string    `json:"memo"`                            // 备注
	Status    int       `json:"status" v:"required|max:2|min:1"` // 状态(1:启用 2:停用)
	Creator   string    `json:"creator"`                         // 创建者
	CreatedAt time.Time `json:"created_at"`                      // 创建时间
}

// DemoQueryParam 查询条件
type DemoQueryParam struct {
	Code     string `form:"-"`        // 编号
	Status   int    `form:"status"`   // 状态(1:启用 2:停用)
	LikeCode string `form:"likeCode"` // 编号(模糊查询)
	LikeName string `form:"likeName"` // 名称(模糊查询)
}

// DemoQueryOptions demo对象查询可选参数项
type DemoQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// DemoQueryResult demo对象查询结果
type DemoQueryResult struct {
	Data       []*Demo
	PageResult *PaginationResult
}
