package models

import (
	"fmt"
	"testing"
)

// User 用户信息
// orm会默认根据结构体创建 table
// orm采用的是 linux 命名方式  即小写加下划线，且会在名字后面加 s
// 会创建users 表
type User struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Hobby string
}

// 自动迁移
// 若该表不存在则创建该表，若该表存在且结构体发生变化则更新表结构
func migrateUser() {
	db.AutoMigrate(&User{})
}

// 插入数据
func createRow(name string, hobby string) {
	db.Create(&User{Name: name, Hobby: hobby})
	fmt.Printf("插入成功 \n")
}

// 查询单条数据记录
func queryRow(name string) {
	var user User
	db.First(&user, "name = ?", name)
	fmt.Println("查看查询记录:", user)
}

// 更新数据
func updateRow(name string, newHobby string) {
	var user User
	db.First(&user, "name = ?", name)
	db.Model(&user).Update("hobby", newHobby)
}

// 删除数据
func deleteRow(name string) {
	var user User
	db.First(&user, "name = ?", name)

	// 删除记录
	db.Delete(&user)
	fmt.Printf("删除成功 \n")
}

// 查询多行
func queryMultiRow() {
	var users []User
	db.Find(&users)
	fmt.Println(users)
}

func TestMigrate(t *testing.T) {
	// 初始化表
	migrateUser()
}

func TestInsert(t *testing.T) {
	// 插入数据
	createRow("吴越", "音乐")
	createRow("李霜", "唱歌")
	createRow("夏以生", "爬山")
}

func TestQuery(t *testing.T) {
	// 查询多行数据
	queryMultiRow()

}

func TestUpdate(t *testing.T) {
	// 更新数据
	updateRow("李霜", "看电影")
	// 查询单行数据
	queryRow("李霜")
}

func TestDelete(t1 *testing.T) {
	// 删除数据
	deleteRow("吴越")
	// 查询多行数据
	queryMultiRow()

}
