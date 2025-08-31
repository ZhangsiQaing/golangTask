package main

import (
	"fmt"
	"task3/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//基本CRUD
	// dsn := "root:root123456@tcp(zsq-mysql:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&sql.Students{})
	//创建
	// student := sql.Students{
	// 	Name:  "zhangiang",
	// 	Age:   19,
	// 	Grade: "八年级",
	// }
	// db.Model(&sql.Students{}).Create(&student)
	//查询
	// var students []sql.Students
	// db.Model(&Students{}).Where("age >= ?", 18).Find(&students)
	// fmt.Println(students)
	//更新
	// db.Model(&sql.Students{}).Where("name = ?", "张三").Update("grade", "四年级")
	//删除
	// db.Model(&sql.Students{}).Where("age < ?", 15).Delete(&sql.Students{})

	//事物语句
	// dsn := "root:root123456@tcp(zsq-mysql:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&sql.Account{}, &sql.Transaction{})
	// if err := sql.Transfer(db, 1, 2, 100); err != nil {
	// 	fmt.Println("转账失败:", err)
	// } else {
	// 	fmt.Println("转账成功！")
	// }

	// sqlx1
	// dsn := "root:root123456@tcp(zsq-mysql:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&sql.Employees{})
	// var employees []sql.Employees
	// db.Raw("SELECT id,name,department,salary From employees WHERE department = ?", "技术部").Scan(&employees)
	// fmt.Println(employees)

	// var employees sql.Employees
	// db.Raw("SELECT id,name,department,salary From employees Order By salary desc limit 1").Scan(&employees)
	// fmt.Println(employees)

	// sqlx2
	// dsn := "root:root123456@tcp(zsq-mysql:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db.AutoMigrate(&sql.Books{})
	// var books []sql.Books
	// db.Raw("SELECT id,title,author,price FROM books WHERE price > ?", 50).Scan(&books)
	// fmt.Println(books)

	//使用sqlx
	dsn := "root:root123456@tcp(zsq-mysql:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&sql.User{}, &sql.Post{}, &sql.Comment{})

	//查询用户所有文章
	// user, err := sql.QueryAllInfomationByUid(db, 1)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(user)

	//查询评论数量最多的文章
	// post, commentCount, err := sql.QueryMostCommentedPost(db)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(post, commentCount)

	//hook
	// post := sql.Post{
	// 	Title:   "Golang 学习笔记",
	// 	Content: "这是一篇关于 GORM 的文章",
	// 	UserID:  1,
	// }
	// if err := db.Model(sql.Post{}).Create(&post).Error; err != nil {
	// 	fmt.Println("创建文章失败:", err)
	// } else {
	// 	fmt.Println("文章创建成功，并自动更新了用户的文章数！")
	// }

	var comment sql.Comment
	if err := db.Model(&sql.Comment{}).First(&comment, 4).Error; err != nil {
		// 处理错误
		fmt.Println("评论不存在：", err)
		return
	}

	if err := db.Model(&sql.Comment{}).Delete(&comment).Error; err != nil {
		fmt.Println("删除评论失败:", err)
		return
	}
	fmt.Println("评论删除成功！")

}
