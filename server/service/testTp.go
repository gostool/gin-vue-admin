package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTestTp
//@description: 创建TestTp记录
//@param: testTp model.TestTp
//@return: err error

func CreateTestTp(testTp model.TestTp) (err error) {
	err = global.GVA_DB.Create(&testTp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestTp
//@description: 删除TestTp记录
//@param: testTp model.TestTp
//@return: err error

func DeleteTestTp(testTp model.TestTp) (err error) {
	err = global.GVA_DB.Delete(&testTp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestTpByIds
//@description: 批量删除TestTp记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTestTpByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TestTp{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTestTp
//@description: 更新TestTp记录
//@param: testTp *model.TestTp
//@return: err error

func UpdateTestTp(testTp model.TestTp) (err error) {
	err = global.GVA_DB.Save(&testTp).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestTp
//@description: 根据id获取TestTp记录
//@param: id uint
//@return: err error, testTp model.TestTp

func GetTestTp(id uint) (err error, testTp model.TestTp) {
	err = global.GVA_DB.Where("id = ?", id).First(&testTp).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestTpInfoList
//@description: 分页获取TestTp记录
//@param: info request.TestTpSearch
//@return: err error, list interface{}, total int64

func GetTestTpInfoList(info request.TestTpSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TestTp{})
    var testTps []model.TestTp
    // 如果有条件搜索 下方会自动创建搜索语句
	if info.Hot > 0 {
		db = db.Where("`hot` > ?",info.Hot)
	}
	if !info.Uptime.IsZero() {
		db = db.Where("`uptime` >= ?",info.Uptime)
	}
	if info.Name != "" {
		db = db.Where("`name` = ?",info.Name)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&testTps).Error
	return err, testTps, total
}