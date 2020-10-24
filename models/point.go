package models

import "github.com/jinzhu/gorm"

type Point struct {
	gorm.Model
	SotenA int
	SotenB int
	SotenC int
	SotenD int
}

// データインサート処理
func dbInsert(sotenA int, sotenB int, sotenC int, sotenD int) {
	db := gormConnect()

	defer db.Close()
	// Insert処理
	db.Create(&Point{
		SotenA: sotenA,
		SotenB: sotenB,
		SotenC: sotenC,
		SotenD: sotenD,
	})
}

//DB更新
func dbUpdate(id int, sotenAValue int, sotenBValue int, sotenCValue int, sotenDValue int) {
	db := gormConnect()
	var point Point
	db.First(&point, id)
	point.SotenA = sotenAValue
	point.SotenB = sotenBValue
	point.SotenC = sotenCValue
	point.SotenD = sotenDValue
	db.Save(&point)
	db.Close()
}

// 全件取得
func dbGetAll() []Point {
	db := gormConnect()

	defer db.Close()
	var pointList []Point
	// FindでDB名を指定して取得した後、orderで登録順に並び替え
	db.Order("created_at desc").Find(&pointList)
	return pointList
}

//DB一つ取得
func dbGetOne(id int) Point {
	db := gormConnect()
	var point Point
	db.First(&point, id)
	db.Close()
	return point
}

//DB削除
func dbDelete(id int) {
	db := gormConnect()
	var point Point
	db.First(&point, id)
	db.Delete(&point)
	db.Close()
}
