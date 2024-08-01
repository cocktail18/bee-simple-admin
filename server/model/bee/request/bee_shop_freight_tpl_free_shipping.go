package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeShopFreightTplFreeShippingSearch struct{
    CityId  string `json:"cityId" form:"cityId" `
    ProvinceId  string `json:"provinceId" form:"provinceId" `
    TplId  *int `json:"tplId" form:"tplId" `
    Type  *int `json:"type" form:"type" `
    request.PageInfo
}
