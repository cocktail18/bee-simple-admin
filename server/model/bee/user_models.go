// 自动生成模板UserModels
package bee

import (
	"time"
)

// 用户信息 结构体  UserModels
type UserModels struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    Domain  string `json:"domain" form:"domain" gorm:"column:domain;comment:;size:100;"`  //domain字段 
    Username  string `json:"username" form:"username" gorm:"column:username;comment:;size:100;"`  //username字段 
    Password  string `json:"password" form:"password" gorm:"column:password;comment:;size:100;"`  //password字段 
    Salt  string `json:"salt" form:"salt" gorm:"column:salt;comment:;size:100;"`  //salt字段 
    Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;size:100;"`  //phone字段 
}


// TableName 用户信息 UserModels自定义表名 user_models
func (UserModels) TableName() string {
    return "user_models"
}

