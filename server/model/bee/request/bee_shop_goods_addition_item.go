package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeShopGoodsAdditionItemSearch struct {
	Id              *int     `json:"id" form:"id" `
	GoodsId         *int     `json:"goodsId" form:"goodsId" `
	GoodsAdditionId *int     `json:"goodsAdditionId" form:"goodsAdditionId" `
	Name            string   `json:"name" form:"name" `
	StartPrice      *float64 `json:"startPrice" form:"startPrice"`
	EndPrice        *float64 `json:"endPrice" form:"endPrice"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
