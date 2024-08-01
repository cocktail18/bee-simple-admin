package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeShopLogisticsExceptionService struct {}

// CreateBeeShopLogisticsException 创建运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService) CreateBeeShopLogisticsException(beeShopLogisticsException *bee.BeeShopLogisticsException) (err error) {
    beeShopLogisticsException.DateAdd = utils.NowPtr()
    beeShopLogisticsException.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopLogisticsException).Error
	return err
}

// DeleteBeeShopLogisticsException 删除运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService)DeleteBeeShopLogisticsException(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopLogisticsException{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
        Updates(map[string]interface{}{
            "is_deleted":  1,
            "date_delete": utils.NowPtr(),
        }).Error
	return err
}

// DeleteBeeShopLogisticsExceptionByIds 批量删除运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService)DeleteBeeShopLogisticsExceptionByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopLogisticsException{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
              Updates(map[string]interface{}{
                  "is_deleted":  1,
                  "date_delete": utils.NowPtr(),
              }).Error
	return err
}

// UpdateBeeShopLogisticsException 更新运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService)UpdateBeeShopLogisticsException(beeShopLogisticsException bee.BeeShopLogisticsException, shopUserId int) (err error) {
    beeShopLogisticsException.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopLogisticsException{}).Where("id = ? and user_id = ?",beeShopLogisticsException.Id, shopUserId).Updates(&beeShopLogisticsException).Error
	return err
}

// GetBeeShopLogisticsException 根据id获取运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService)GetBeeShopLogisticsException(id string, shopUserId int) (beeShopLogisticsException bee.BeeShopLogisticsException, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopLogisticsException).Error
	return
}

// GetBeeShopLogisticsExceptionInfoList 分页获取运费排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopLogisticsExceptionService *BeeShopLogisticsExceptionService)GetBeeShopLogisticsExceptionInfoList(info beeReq.BeeShopLogisticsExceptionSearch, shopUserId int) (list []bee.BeeShopLogisticsException, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopLogisticsException{})
	db = db.Where("user_id = ?", shopUserId)
    var beeShopLogisticsExceptions []bee.BeeShopLogisticsException
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LogisticsId != nil {
        db = db.Where("logistics_id = ?",info.LogisticsId)
    }
    if info.ProvinceId != "" {
        db = db.Where("province_id = ?",info.ProvinceId)
    }
    if info.CityId != "" {
        db = db.Where("city_id = ?",info.CityId)
    }
    if info.RegionId != "" {
        db = db.Where("region_id = ?",info.RegionId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["date_add"] = true
         	orderMap["logistics_id"] = true
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
	
	err = db.Find(&beeShopLogisticsExceptions).Error
	return  beeShopLogisticsExceptions, total, err
}