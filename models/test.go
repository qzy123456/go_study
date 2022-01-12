package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserInfos struct {
	Uid                int       `gorm:"uid"`
	Silver_coin        int       `gorm:"silver_coin"`
	Gold_coin          int       `gorm:"gold_coin"`
	Task_turnover      int       `gorm:"task_turnover"`
	Shop_grade         int       `gorm:"shop_grade"`
	Lives              int       `gorm:"lives"`
	Last               int       `gorm:"last"`
	New_level          int       `gorm:"new_level"`
	Stars              int       `gorm:"stars"`
	Chest_last         int       `gorm:"chest_last"`
	Chest_number       int       `gorm:"chest_number"`
}
var db *gorm.DB
func main()  {

	//利用伪造单例模式 得到 db1还是db2
  var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ro_api_db2_dev?charset=utf8,utf8mb4;")
	if err != nil {
		panic(err)
	}
	//日志
	db.LogMode(true)
	cc := Getinfo(57)
	fmt.Println(cc)
}

func (UserInfos) TableName() string {

	return "user_info_1"
}

//传入db，以及其他参数
func Getinfo(uid int) (v *UserInfos) {
	user := new(UserInfos)
	defer db.Close()

	//拼接条件
	if err :=db.Where("uid=?", uid).First(&user).Error; err!=nil{
		fmt.Println("没找到",err)
	}
	return user
}

//传入db，以及其他参数
func QueryAll() (v *[]UserInfos) {
	user := new([]UserInfos)
	defer db.Close()

	//拼接条件
	db.Find(&user)
	return user
}
func(u *UserInfos) BeforeUpdate()(err error){
	if u.Stars == 0 {
		err = errors.New("read only user")
	}
	return
}
