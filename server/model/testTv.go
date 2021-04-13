// 自动生成模板TestTv
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type TestTv struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名字;type:varchar(100);size:100;"`
      Channel  string `json:"channel" form:"channel" gorm:"column:channel;comment:频道;type:varchar(100);size:100;"`
}


func (TestTv) TableName() string {
  return "testTv"
}

