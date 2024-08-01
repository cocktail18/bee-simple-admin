package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeShopFreightTplExceptionService struct {}

// CreateBeeShopFreightTplException 创建运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService) CreateBeeShopFreightTplException(beeShopFreightTplException *bee.BeeShopFreightTplException) (err error) {
    beeShopFreightTplException.DateAdd = utils.NowPtr()
    beeShopFreightTplException.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeShopFreightTplException).Error
	return err
}

// DeleteBeeShopFreightTplException 删除运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService)DeleteBeeShopFreightTplException(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplException{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
        Updates(map[string]interface{}{
            "is_deleted":  1,
            "date_delete": utils.NowPtr(),
        }).Error
	return err
}

// DeleteBeeShopFreightTplExceptionByIds 批量删除运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService)DeleteBeeShopFreightTplExceptionByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplException{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
              Updates(map[string]interface{}{
                  "is_deleted":  1,
                  "date_delete": utils.NowPtr(),
              }).Error
	return err
}

// UpdateBeeShopFreightTplException 更新运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService)UpdateBeeShopFreightTplException(beeShopFreightTplException bee.BeeShopFreightTplException, shopUserId int) (err error) {
    beeShopFreightTplException.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplException{}).Where("id = ? and user_id = ?",beeShopFreightTplException.Id, shopUserId).Updates(&beeShopFreightTplException).Error
	return err
}

// GetBeeShopFreightTplException 根据id获取运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService)GetBeeShopFreightTplException(id string, shopUserId int) (beeShopFreightTplException bee.BeeShopFreightTplException, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeShopFreightTplException).Error
	return
}

// GetBeeShopFreightTplExceptionInfoList 分页获取运费模版排除项记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeShopFreightTplExceptionService *BeeShopFreightTplExceptionService)GetBeeShopFreightTplExceptionInfoList(info beeReq.BeeShopFreightTplExceptionSearch, shopUserId int) (list []bee.BeeShopFreightTplException, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeShopFreightTplException{})
	db = db.Where("user_id = ?", shopUserId)
    var beeShopFreightTplExceptions []bee.BeeShopFreightTplException
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.DetailId != nil {
        db = db.Where("detail_id = ?",info.DetailId)
    }
    if info.ProvinceId != "" {
        db = db.Where("province_id = ?",info.ProvinceId)
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
         	orderMap["id"] = true
         	orderMap["date_add"] = true
         	orderMap["date_update"] = true
         	orderMap["date_delete"] = true
         	orderMap["detail_id"] = true
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
	
	err = db.Find(&beeShopFreightTplExceptions).Error
	return  beeShopFreightTplExceptions, total, err
}