package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeShopLogisticsExceptionSearch struct{
    LogisticsId  *int `json:"logisticsId" form:"logisticsId" `
    ProvinceId  string `json:"provinceId" form:"provinceId" `
    CityId  string `json:"cityId" form:"cityId" `
    RegionId  string `json:"regionId" form:"regionId" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
