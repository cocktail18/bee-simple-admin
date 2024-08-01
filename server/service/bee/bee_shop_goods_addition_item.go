package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeShopGoodsAdditionItemService struct{}

// CreateBeeShopGoodsAdditionItem 创建商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) CreateBeeShopGoodsAdditionItem(beeShopGoodsAdditionItem *bee.BeeShopGoodsAdditionItem) (err error) {
	beeShopGoodsAdditionItem.DateAdd = utils.NowPtr()
	beeShopGoodsAdditionItem.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopGoodsAdditionItem).Error
	return err
}

// DeleteBeeShopGoodsAdditionItem 删除商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) DeleteBeeShopGoodsAdditionItem(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsAdditionItem{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// DeleteBeeShopGoodsAdditionItemByIds 批量删除商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) DeleteBeeShopGoodsAdditionItemByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsAdditionItem{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
		Updates(map[string]interface{}{
			"is_deleted":  1,
			"date_delete": utils.NowPtr(),
		}).Error
	return err
}

// UpdateBeeShopGoodsAdditionItem 更新商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) UpdateBeeShopGoodsAdditionItem(beeShopGoodsAdditionItem bee.BeeShopGoodsAdditionItem, shopUserId int) (err error) {
	beeShopGoodsAdditionItem.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsAdditionItem{}).Where("id = ? and user_id = ?", beeShopGoodsAdditionItem.Id, shopUserId).Updates(&beeShopGoodsAdditionItem).Error
	return err
}

// GetBeeShopGoodsAdditionItem 根据id获取商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) GetBeeShopGoodsAdditionItem(id string, shopUserId int) (beeShopGoodsAdditionItem bee.BeeShopGoodsAdditionItem, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopGoodsAdditionItem).Error
	return
}

// GetBeeShopGoodsAdditionItemInfoList 分页获取商品额外信息绑定记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopGoodsAdditionItemService *BeeShopGoodsAdditionItemService) GetBeeShopGoodsAdditionItemInfoList(info beeReq.BeeShopGoodsAdditionItemSearch, shopUserId int) (list []bee.BeeShopGoodsAdditionItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopGoodsAdditionItem{})
	db = db.Where("user_id = ?", shopUserId)
	var beeShopGoodsAdditionItems []bee.BeeShopGoodsAdditionItem
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Id != nil {
		db = db.Where("id = ?", info.Id)
	}
	if info.GoodsId != nil {
		db = db.Where("goods_id = ?", info.GoodsId)
	}
	if info.GoodsAdditionId != nil {
		db = db.Where("goods_addition_id = ?", info.GoodsAdditionId)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.StartPrice != nil && info.EndPrice != nil {
		db = db.Where("price BETWEEN ? AND ? ", info.StartPrice, info.EndPrice)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["id"] = true
	orderMap["goods_id"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&beeShopGoodsAdditionItems).Error
	return beeShopGoodsAdditionItems, total, err
}
