// 自动生成模板BeeShopFreightTplFreeShipping
package bee

import (
	"time"
)

// 运费模版免运费配置 结构体  BeeShopFreightTplFreeShipping
type BeeShopFreightTplFreeShipping struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    CityId  string `json:"cityId" form:"cityId" gorm:"column:city_id;comment:;size:100;"`  //cityId字段 
    ProvinceId  string `json:"provinceId" form:"provinceId" gorm:"column:province_id;comment:;size:100;"`  //provinceId字段 
    Number  *float64 `json:"number" form:"number" gorm:"column:number;comment:;size:20;"`  //number字段 
    TplId  *int `json:"tplId" form:"tplId" gorm:"column:tpl_id;comment:;size:19;"`  //tplId字段 
    Type  *int `json:"type" form:"type" gorm:"column:type;comment:0件，1重量，3金额;size:10;"`  //0件，1重量，3金额 
    CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:;"`  //createdAt字段 
    UpdatedAt  *time.Time `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:;"`  //updatedAt字段 
    DeletedAt  *time.Time `json:"deletedAt" form:"deletedAt" gorm:"column:deleted_at;comment:;"`  //deletedAt字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
}


// TableName 运费模版免运费配置 BeeShopFreightTplFreeShipping自定义表名 bee_shop_freight_tpl_free_shipping
func (BeeShopFreightTplFreeShipping) TableName() string {
    return "bee_shop_freight_tpl_free_shipping"
}

