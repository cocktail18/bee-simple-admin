package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeOrderGoodsSearch struct{
    StartDateAdd  *time.Time  `json:"startDateAdd" form:"startDateAdd"`
    EndDateAdd  *time.Time  `json:"endDateAdd" form:"endDateAdd"`
    StartDateUpdate  *time.Time  `json:"startDateUpdate" form:"startDateUpdate"`
    EndDateUpdate  *time.Time  `json:"endDateUpdate" form:"endDateUpdate"`
    OrderId  *int `json:"orderId" form:"orderId" `
    Purchase  string `json:"purchase" form:"purchase" `
    RefundStatus  *int `json:"refundStatus" form:"refundStatus" `
    Uid  *int `json:"uid" form:"uid" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
