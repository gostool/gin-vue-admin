package request

import "gin-vue-admin/model"

type TestGameSearch struct{
    model.TestGame
    PageInfo
}