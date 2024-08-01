// 自动生成模板BeeShopGoodsAdditionItem
package bee

import (
	"time"
)

// 商品额外信息绑定 结构体  BeeShopGoodsAdditionItem
type BeeShopGoodsAdditionItem struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`  //创建时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    GoodsId  *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:;size:19;"`  //goodsId字段 
    GoodsAdditionId  *int `json:"goodsAdditionId" form:"goodsAdditionId" gorm:"column:goods_addition_id;comment:;size:20;"`  //goodsAdditionId字段 
    Name  string `json:"name" form:"name" gorm:"column:name;comment:;size:100;"`  //name字段 
    Price  *float64 `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`  //价格 
}


// TableName 商品额外信息绑定 BeeShopGoodsAdditionItem自定义表名 bee_shop_goods_addition_item
func (BeeShopGoodsAdditionItem) TableName() string {
    return "bee_shop_goods_addition_item"
}

