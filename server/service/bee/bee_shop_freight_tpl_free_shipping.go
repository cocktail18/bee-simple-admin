package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeShopFreightTplFreeShippingService struct {}

// CreateBeeShopFreightTplFreeShipping 创建运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService) CreateBeeShopFreightTplFreeShipping(beeShopFreightTplFreeShipping *bee.BeeShopFreightTplFreeShipping) (err error) {
    beeShopFreightTplFreeShipping.DateAdd = utils.NowPtr()
    beeShopFreightTplFreeShipping.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopFreightTplFreeShipping).Error
	return err
}

// DeleteBeeShopFreightTplFreeShipping 删除运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService)DeleteBeeShopFreightTplFreeShipping(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplFreeShipping{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
        Updates(map[string]interface{}{
            "is_deleted":  1,
            "date_delete": utils.NowPtr(),
        }).Error
	return err
}

// DeleteBeeShopFreightTplFreeShippingByIds 批量删除运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService)DeleteBeeShopFreightTplFreeShippingByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplFreeShipping{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
              Updates(map[string]interface{}{
                  "is_deleted":  1,
                  "date_delete": utils.NowPtr(),
              }).Error
	return err
}

// UpdateBeeShopFreightTplFreeShipping 更新运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService)UpdateBeeShopFreightTplFreeShipping(beeShopFreightTplFreeShipping bee.BeeShopFreightTplFreeShipping, shopUserId int) (err error) {
    beeShopFreightTplFreeShipping.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplFreeShipping{}).Where("id = ? and user_id = ?",beeShopFreightTplFreeShipping.Id, shopUserId).Updates(&beeShopFreightTplFreeShipping).Error
	return err
}

// GetBeeShopFreightTplFreeShipping 根据id获取运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService)GetBeeShopFreightTplFreeShipping(id string, shopUserId int) (beeShopFreightTplFreeShipping bee.BeeShopFreightTplFreeShipping, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopFreightTplFreeShipping).Error
	return
}

// GetBeeShopFreightTplFreeShippingInfoList 分页获取运费模版免运费配置记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplFreeShippingService *BeeShopFreightTplFreeShippingService)GetBeeShopFreightTplFreeShippingInfoList(info beeReq.BeeShopFreightTplFreeShippingSearch, shopUserId int) (list []bee.BeeShopFreightTplFreeShipping, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplFreeShipping{})
	db = db.Where("user_id = ?", shopUserId)
    var beeShopFreightTplFreeShippings []bee.BeeShopFreightTplFreeShipping
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.CityId != "" {
        db = db.Where("city_id = ?",info.CityId)
    }
    if info.ProvinceId != "" {
        db = db.Where("province_id = ?",info.ProvinceId)
    }
    if info.TplId != nil {
        db = db.Where("tpl_id = ?",info.TplId)
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&beeShopFreightTplFreeShippings).Error
	return  beeShopFreightTplFreeShippings, total, err
}