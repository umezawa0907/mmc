package models

import "github.com/jinzhu/gorm"

func gormConnect() *gorm.DB {
	DBMS := "postgres"
	USER := "postgres"
	PASS := "hockey35"
	DBNAME := "mmcdb"
	CONNECT := "user=" + USER + " dbname=" + DBNAME + " password=" + PASS + " sslmode=disable"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// DbInit はDBの初期化
func DbInit() {
	db := gormConnect()

	// コネクション解放解放
	defer db.Close()
	db.AutoMigrate(&Point{}) //構造体に基づいてテーブルを作成
}
