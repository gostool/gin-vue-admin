package request

import "gin-vue-admin/model"

type TestTpSearch struct{
    model.TestTp
    PageInfo
}