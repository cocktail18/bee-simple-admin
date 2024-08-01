package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type UserModelsSearch struct{
    Id  *int `json:"id" form:"id" `
    Username  string `json:"username" form:"username" `
    Phone  string `json:"phone" form:"phone" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
