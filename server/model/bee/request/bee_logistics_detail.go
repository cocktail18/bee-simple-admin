package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeLogisticsDetailSearch struct{
    LogisticsId  *int `json:"logisticsId" form:"logisticsId" `
    Type  *int `json:"type" form:"type" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
