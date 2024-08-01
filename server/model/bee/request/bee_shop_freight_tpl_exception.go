package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeShopFreightTplExceptionSearch struct{
    DetailId  *int `json:"detailId" form:"detailId" `
    ProvinceId  string `json:"provinceId" form:"provinceId" `
    RegionId  string `json:"regionId" form:"regionId" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
