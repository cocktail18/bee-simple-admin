// 自动生成模板BeeLogisticsDetail
package bee

import (
	"time"
)

// 运费模版详情 结构体  BeeLogisticsDetail
type BeeLogisticsDetail struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    LogisticsId  *int `json:"logisticsId" form:"logisticsId" gorm:"column:logistics_id;comment:;size:19;"`  //logisticsId字段 
    AddAmount  *float64 `json:"addAmount" form:"addAmount" gorm:"column:add_amount;comment:条件包邮,满xx包邮;size:20;"`  //条件包邮,满xx包邮 
    AddNumber  *float64 `json:"addNumber" form:"addNumber" gorm:"column:add_number;comment:条件包邮,满xx件包邮;size:20;"`  //条件包邮,满xx件包邮 
    FirstAmount  *float64 `json:"firstAmount" form:"firstAmount" gorm:"column:first_amount;comment:首件价格;size:20;"`  //首件价格 
    FirstNumber  *float64 `json:"firstNumber" form:"firstNumber" gorm:"column:first_number;comment:首件数量;size:20;"`  //首件数量 
    Type  *int `json:"type" form:"type" gorm:"column:type;comment:类型，1首件，2续件;size:10;"`  //类型，1首件，2续件 
}


// TableName 运费模版详情 BeeLogisticsDetail自定义表名 bee_logistics_detail
func (BeeLogisticsDetail) TableName() string {
    return "bee_logistics_detail"
}

