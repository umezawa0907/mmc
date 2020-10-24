package models

import "github.com/jinzhu/gorm"

// Point は素点の型定義
type Point struct {
	gorm.Model
	SotenA int
	SotenB int
	SotenC int
	SotenD int
}

// DbInsert はデータインサート処理
func DbInsert(sotenA int, sotenB int, sotenC int, sotenD int) {
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

// DbUpdate はDB更新
func DbUpdate(id int, sotenAValue int, sotenBValue int, sotenCValue int, sotenDValue int) {
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

// DbGetAll は全件取得
func DbGetAll() []Point {
	db := gormConnect()

	defer db.Close()
	var pointList []Point
	// FindでDB名を指定して取得した後、orderで登録順に並び替え
	db.Order("created_at desc").Find(&pointList)
	return pointList
}

// DbGetOne はDB一つ取得
func DbGetOne(id int) Point {
	db := gormConnect()
	var point Point
	db.First(&point, id)
	db.Close()
	return point
}

// DbDelete はDB削除
func DbDelete(id int) {
	db := gormConnect()
	var point Point
	db.First(&point, id)
	db.Delete(&point)
	db.Close()
}
