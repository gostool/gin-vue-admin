// 自动生成模板TestGame
package model

import (
	"gin-vue-admin/global"
    "time"
)

// 如果含有time.Time 请自行import time包
type TestGame struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名字;type:varchar(100);size:100;"`
      Area  string `json:"area" form:"area" gorm:"column:area;comment:游戏区域;type:varchar(20);size:20;"`
      UpTime  time.Time `json:"upTime" form:"upTime" gorm:"column:up_time;comment:上线时间;type:timestamp;size:13;"`
}


func (TestGame) TableName() string {
  return "test_game"
}

