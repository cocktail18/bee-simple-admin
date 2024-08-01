package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeShopFreightTplSearch struct{
    UserId  *int `json:"userId" form:"userId" `
    Name  string `json:"name" form:"name" `
    FeeType  *int `json:"feeType" form:"feeType" `
    IsFree  *bool `json:"isFree" form:"isFree" `
    request.PageInfo
}
