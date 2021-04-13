package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTestTv
//@description: 创建TestTv记录
//@param: testTv model.TestTv
//@return: err error

func CreateTestTv(testTv model.TestTv) (err error) {
	err = global.GVA_DB.Create(&testTv).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestTv
//@description: 删除TestTv记录
//@param: testTv model.TestTv
//@return: err error

func DeleteTestTv(testTv model.TestTv) (err error) {
	err = global.GVA_DB.Delete(&testTv).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestTvByIds
//@description: 批量删除TestTv记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTestTvByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TestTv{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTestTv
//@description: 更新TestTv记录
//@param: testTv *model.TestTv
//@return: err error

func UpdateTestTv(testTv model.TestTv) (err error) {
	err = global.GVA_DB.Save(&testTv).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestTv
//@description: 根据id获取TestTv记录
//@param: id uint
//@return: err error, testTv model.TestTv

func GetTestTv(id uint) (err error, testTv model.TestTv) {
	err = global.GVA_DB.Where("id = ?", id).First(&testTv).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestTvInfoList
//@description: 分页获取TestTv记录
//@param: info request.TestTvSearch
//@return: err error, list interface{}, total int64

func GetTestTvInfoList(info request.TestTvSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TestTv{})
    var testTvs []model.TestTv
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` = ?",info.Name)
    }
    if info.Channel != "" {
        db = db.Where("`channel` = ?",info.Channel)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&testTvs).Error
	return err, testTvs, total
}