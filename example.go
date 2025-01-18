package main

import (
	_ "database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func prepareDb() error {

	dataSourceName := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("sql.Open error ", err)
	}

	return nil
}

func endDb() {
	if sqlDB, err := db.DB(); err == nil {
		sqlDB.Close()
	}

	fmt.Println("end gorm-taiso!")
}

type User struct {
	ID        int
	UserName  string
	Email     string
	Password  string
	Location  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	fmt.Println("start gorm-taiso!")
	if err := prepareDb(); err != nil {
		fmt.Println("prepareDb error", err)
		return
	}
	defer endDb()

	// 起動確認として 1 件データを取得して表示します
	var userTest User
	errTest := db.First(&userTest, "id = ?", 1).Error
	if errTest != nil {
		if errors.Is(errTest, gorm.ErrRecordNotFound) {
			fmt.Println("err:", errTest)
		} else {
			fmt.Println("err:", errTest)
		}
		return
	} else {
		fmt.Println("test id:", userTest.ID, "user_name:", userTest.UserName)
	}

	////////////////////////////////////////
	// 練習問題
	// ※以下に回答を記入して実行し手確認してください
	////////////////////////////////////////

	// 練習：users テーブルにデータを insert しましょう
	// ※変数 db が、GORM の DB 接続インスタンスです。db に対して操作を実行します。
	//   ＜例＞ db.Create(&User{UserName: "test", ～
	// ※User構造体を上の方に定義してあるので、よければ利用してください
	newUser := User{
		UserName: "taro",
		Email:    "taro@aaa.com",
		Password: "password",
		Location: "tokyo",
	}

	err := db.Create(&newUser).Error
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("insert success")

	// 練習：insert したデータの id を表示しましょう
	fmt.Println("id:", newUser.ID)

	// 練習：insert したデータを select で取得してみましょう
	user := User{}
	err = db.First(&user, "id = ?", newUser.ID).Error
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("User:", user)

	// 練習：データの user_name を更新してみましょう
	user.UserName = "hanako"
	err = db.Save(&user).Error
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	// 練習：データを delete しましょう
	err = db.Delete(&user).Error
	if err != nil {
		fmt.Println("err:", err)
		return
	}

}
