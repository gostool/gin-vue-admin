package request

import "gin-vue-admin/model"

type TestTvSearch struct{
    model.TestTv
    PageInfo
}