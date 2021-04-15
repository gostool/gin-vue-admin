// 自动生成模板TestTp
package model

import (
	"gin-vue-admin/global"
    "time"
)

// 如果含有time.Time 请自行import time包
type TestTp struct {
      global.GVA_MODEL
      Hot  int `json:"hot" form:"hot" gorm:"column:hot;comment:热度;type:smallint;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名字;type:varchar(100);size:100;"`
      Uptime  time.Time `json:"uptme" form:"uptime" gorm:"column:uptime;comment:上映时间;type:timestamp;"`
}


func (TestTp) TableName() string {
  return "test_tp"
}

