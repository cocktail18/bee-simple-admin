package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeShopFreightTplService struct{}

// CreateBeeShopFreightTpl 创建商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) CreateBeeShopFreightTpl(beeShopFreightTpl *bee.BeeShopFreightTpl) (err error) {
	beeShopFreightTpl.DateAdd = utils.NowPtr()
	beeShopFreightTpl.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopFreightTpl).Error
	return err
}

// DeleteBeeShopFreightTpl 删除商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) DeleteBeeShopFreightTpl(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTpl{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopFreightTplByIds 批量删除商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) DeleteBeeShopFreightTplByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTpl{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopFreightTpl 更新商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) UpdateBeeShopFreightTpl(beeShopFreightTpl bee.BeeShopFreightTpl, shopUserId int) (err error) {
	beeShopFreightTpl.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTpl{}).Where("id = ? and user_id = ?", beeShopFreightTpl.Id, shopUserId).Updates(&beeShopFreightTpl).Error
	return err
}

// GetBeeShopFreightTpl 根据id获取商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) GetBeeShopFreightTpl(id string, shopUserId int) (beeShopFreightTpl bee.BeeShopFreightTpl, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopFreightTpl).Error
	return
}

// GetBeeShopFreightTplInfoList 分页获取商城运费模版记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplService *BeeShopFreightTplService) GetBeeShopFreightTplInfoList(info beeReq.BeeShopFreightTplSearch, shopUserId int) (list []bee.BeeShopFreightTpl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTpl{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopFreightTpls []bee.BeeShopFreightTpl
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.FeeType != nil {
		db = db.Where("fee_type = ?", info.FeeType)
	}
	if info.IsFree != nil {
		db = db.Where("is_free = ?", info.IsFree)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeShopFreightTpls).Error
	return beeShopFreightTpls, total, err
}
