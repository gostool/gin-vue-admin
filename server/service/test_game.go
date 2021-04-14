package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTestGame
//@description: 创建TestGame记录
//@param: testGame model.TestGame
//@return: err error

func CreateTestGame(testGame model.TestGame) (err error) {
	err = global.GVA_DB.Create(&testGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestGame
//@description: 删除TestGame记录
//@param: testGame model.TestGame
//@return: err error

func DeleteTestGame(testGame model.TestGame) (err error) {
	err = global.GVA_DB.Delete(&testGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTestGameByIds
//@description: 批量删除TestGame记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTestGameByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.TestGame{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTestGame
//@description: 更新TestGame记录
//@param: testGame *model.TestGame
//@return: err error

func UpdateTestGame(testGame model.TestGame) (err error) {
	err = global.GVA_DB.Save(&testGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestGame
//@description: 根据id获取TestGame记录
//@param: id uint
//@return: err error, testGame model.TestGame

func GetTestGame(id uint) (err error, testGame model.TestGame) {
	err = global.GVA_DB.Where("id = ?", id).First(&testGame).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTestGameInfoList
//@description: 分页获取TestGame记录
//@param: info request.TestGameSearch
//@return: err error, list interface{}, total int64

func GetTestGameInfoList(info request.TestGameSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.TestGame{})
    var testGames []model.TestGame
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Area != "" {
        db = db.Where("`area` = ?",info.Area)
    }
    if !info.UpTime.IsZero() {
         db = db.Where("`up_time` <> ?",info.UpTime)
    }
	if info.Name != "" {
		db = db.Where("`name` = ?",info.Name)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&testGames).Error
	return err, testGames, total
}