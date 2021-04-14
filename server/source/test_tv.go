package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"github.com/gookit/color"
	"gorm.io/gorm"
	"time"
)

var TestTv = new(testTv)

type testTv struct{}

var testTvs = []model.TestTv{
	{GVA_MODEL: global.GVA_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "testTv01", Channel: "001"},
	{GVA_MODEL: global.GVA_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "testTv02", Channel: "002"},
}


//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (tv *testTv) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.TestTv{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> test_tv 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&testTvs).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> test_tv 表初始数据成功!")
		return nil
	})
}


