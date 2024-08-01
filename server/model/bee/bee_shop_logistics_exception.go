// 自动生成模板BeeShopLogisticsException
package bee

import (
	"time"
)

// 运费排除项 结构体  BeeShopLogisticsException
type BeeShopLogisticsException struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    LogisticsId  *int `json:"logisticsId" form:"logisticsId" gorm:"column:logistics_id;comment:模板详情id;size:19;"`  //模板详情id 
    DetailId  *int `json:"detailId" form:"detailId" gorm:"column:detail_id;comment:模板详情id;size:19;"`  //模板详情id 
    ProvinceId  string `json:"provinceId" form:"provinceId" gorm:"column:province_id;comment:;size:100;"`  //provinceId字段 
    CityId  string `json:"cityId" form:"cityId" gorm:"column:city_id;comment:;size:100;"`  //cityId字段 
    RegionId  string `json:"regionId" form:"regionId" gorm:"column:region_id;comment:;size:100;"`  //regionId字段 
    RegionStr  string `json:"regionStr" form:"regionStr" gorm:"column:region_str;comment:;size:100;"`  //regionStr字段 
    AddAmount  *float64 `json:"addAmount" form:"addAmount" gorm:"column:add_amount;comment:;size:20;"`  //addAmount字段 
    AddNumber  *float64 `json:"addNumber" form:"addNumber" gorm:"column:add_number;comment:;size:20;"`  //addNumber字段 
    FirstAmount  *float64 `json:"firstAmount" form:"firstAmount" gorm:"column:first_amount;comment:;size:20;"`  //firstAmount字段 
    FirstNumber  *float64 `json:"firstNumber" form:"firstNumber" gorm:"column:first_number;comment:;size:20;"`  //firstNumber字段 
}


// TableName 运费排除项 BeeShopLogisticsException自定义表名 bee_shop_logistics_exception
func (BeeShopLogisticsException) TableName() string {
    return "bee_shop_logistics_exception"
}

