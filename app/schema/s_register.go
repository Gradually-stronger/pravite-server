package schema

import "time"

// Register register对象
type Register struct {
	RecordId  string    `json:"record_id"`              // 记录ID
	UserName  string    `json:"user_name" v:"required"` // 用户名称
	PassWord  string    `json:"password" v:"required"`  // 用户密码
	CreatedAt time.Time `json:"created_at"`             //创建时间
}
