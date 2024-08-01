package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BeeLogisticsDetailService struct {}

// CreateBeeLogisticsDetail 创建运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService) CreateBeeLogisticsDetail(beeLogisticsDetail *bee.BeeLogisticsDetail) (err error) {
    beeLogisticsDetail.DateAdd = utils.NowPtr()
    beeLogisticsDetail.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(beeLogisticsDetail).Error
	return err
}

// DeleteBeeLogisticsDetail 删除运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService)DeleteBeeLogisticsDetail(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogisticsDetail{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
        Updates(map[string]interface{}{
            "is_deleted":  1,
            "date_delete": utils.NowPtr(),
        }).Error
	return err
}

// DeleteBeeLogisticsDetailByIds 批量删除运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService)DeleteBeeLogisticsDetailByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogisticsDetail{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
              Updates(map[string]interface{}{
                  "is_deleted":  1,
                  "date_delete": utils.NowPtr(),
              }).Error
	return err
}

// UpdateBeeLogisticsDetail 更新运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService)UpdateBeeLogisticsDetail(beeLogisticsDetail bee.BeeLogisticsDetail, shopUserId int) (err error) {
    beeLogisticsDetail.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogisticsDetail{}).Where("id = ? and user_id = ?",beeLogisticsDetail.Id, shopUserId).Updates(&beeLogisticsDetail).Error
	return err
}

// GetBeeLogisticsDetail 根据id获取运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService)GetBeeLogisticsDetail(id string, shopUserId int) (beeLogisticsDetail bee.BeeLogisticsDetail, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&beeLogisticsDetail).Error
	return
}

// GetBeeLogisticsDetailInfoList 分页获取运费模版详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (beeLogisticsDetailService *BeeLogisticsDetailService)GetBeeLogisticsDetailInfoList(info beeReq.BeeLogisticsDetailSearch, shopUserId int) (list []bee.BeeLogisticsDetail, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.BeeLogisticsDetail{})
	db = db.Where("user_id = ?", shopUserId)
    var beeLogisticsDetails []bee.BeeLogisticsDetail
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.LogisticsId != nil {
        db = db.Where("logistics_id = ?",info.LogisticsId)
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["id"] = true
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
	
	err = db.Find(&beeLogisticsDetails).Error
	return  beeLogisticsDetails, total, err
}