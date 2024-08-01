package bee

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
    beeReq "github.com/flipped-aurora/gin-vue-admin/server/model/bee/request"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type UserModelsService struct {}

// CreateUserModels 创建用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService) CreateUserModels(userModels *bee.UserModels) (err error) {
    userModels.DateAdd = utils.NowPtr()
    userModels.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Create(userModels).Error
	return err
}

// DeleteUserModels 删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService)DeleteUserModels(id string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.UserModels{}).Where("id = ?", id).Where("user_id = ?", shopUserId).
        Updates(map[string]interface{}{
            "is_deleted":  1,
            "date_delete": utils.NowPtr(),
        }).Error
	return err
}

// DeleteUserModelsByIds 批量删除用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService)DeleteUserModelsByIds(ids []string, shopUserId int) (err error) {
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.UserModels{}).Where("id = ?", ids).Where("user_id = ?", shopUserId).
              Updates(map[string]interface{}{
                  "is_deleted":  1,
                  "date_delete": utils.NowPtr(),
              }).Error
	return err
}

// UpdateUserModels 更新用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService)UpdateUserModels(userModels bee.UserModels, shopUserId int) (err error) {
    userModels.DateUpdate = utils.NowPtr()
	err = global.MustGetGlobalDBByDBName("bee").Model(&bee.UserModels{}).Where("id = ? and user_id = ?",userModels.Id, shopUserId).Updates(&userModels).Error
	return err
}

// GetUserModels 根据id获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService)GetUserModels(id string, shopUserId int) (userModels bee.UserModels, err error) {
	err = global.MustGetGlobalDBByDBName("bee").Where("id = ? and user_id = ?", id, shopUserId).First(&userModels).Error
	return
}

// GetUserModelsInfoList 分页获取用户信息记录
// Author [piexlmax](https://github.com/piexlmax)
func (userModelsService *UserModelsService)GetUserModelsInfoList(info beeReq.UserModelsSearch, shopUserId int) (list []bee.UserModels, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("bee").Model(&bee.UserModels{})
	db = db.Where("user_id = ?", shopUserId)
    var userModelss []bee.UserModels
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Id != nil {
        db = db.Where("id = ?",info.Id)
    }
    if info.Username != "" {
        db = db.Where("username = ?",info.Username)
    }
    if info.Phone != "" {
        db = db.Where("phone = ?",info.Phone)
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
	
	err = db.Find(&userModelss).Error
	return  userModelss, total, err
}