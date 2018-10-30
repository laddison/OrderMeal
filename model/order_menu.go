package model

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Order struct {
}

type OrderData struct {
	Id       int
	Username string
	Menu     string
	Created  string
	MenuId   int
}

func (this *Order) GetList() []OrderData {
	db, err := sql.Open("sqlite3", "./OrderMeal.db")

	rows, err := db.Query("SELECT id, username, menu, created, menu_id FROM OrderMenu where status = 0 order by id desc")
	defer rows.Close()
	this.CheckErr(err)

	Datas := []OrderData{}

	for rows.Next() {
		info := OrderData{}
		err = rows.Scan(&info.Id, &info.Username, &info.Menu, &info.Created, &info.MenuId)
		this.CheckErr(err)

		Datas = append(Datas, info)
	}

	return Datas
}

func (this *Order) GetInfoByUid(userId int) OrderData {
	db, err := sql.Open("sqlite3", "./OrderMeal.db")
	this.CheckErr(err)

	info := OrderData{}
	db.QueryRow("SELECT id, username, menu, created, menu_id FROM OrderMenu where status = 0 and  user_id = ? order by id desc", userId).Scan(&info.Id, &info.Username, &info.Menu, &info.Created, &info.MenuId)
	defer db.Close()

	return info
}

func (this *Order) Insert(username string, menu string, create_time string, userId int, menuId int) int64 {
	db, err := sql.Open("sqlite3", "./OrderMeal.db")
	stmt, err := db.Prepare("INSERT INTO OrderMenu(username, menu, created, user_id, menu_id) values(?,?,?,?,?)")
	res, err := stmt.Exec(username, menu, create_time, userId, menuId)
	id, err := res.LastInsertId()
	this.CheckErr(err)

	return id
}

func (this *Order) Update(menu string, userId int, menuId int) int64 {
	db, err := sql.Open("sqlite3", "./OrderMeal.db")
	stmt, err := db.Prepare("UPDATE OrderMenu set menu = ?, menu_id = ? where and status = 0 user_id= ?")
	res, err := stmt.Exec(menu, menuId, userId)
	num, err := res.RowsAffected()
	this.CheckErr(err)

	return num
}

func (this *Order) UpdateStatus() int64 {
	db, err := sql.Open("sqlite3", "./OrderMeal.db")
	stmt, err := db.Prepare("UPDATE OrderMenu set status = 1 ")
	res, err := stmt.Exec()
	num, err := res.RowsAffected()
	this.CheckErr(err)

	return num
}

// 统计
func (this *Order) GetSumTotal(list []OrderData) [3]int {
	var (
		menu1 int
		menu2 int
		menu3 int
	)
	menu1 = 0
	menu2 = 0
	menu3 = 0

	for _, value := range list {
		if value.MenuId == 1 {
			menu1 += 1
		} else if value.MenuId == 2 {
			menu2 += 1
		} else if value.MenuId == 3 {
			menu3 += 1
		}
	}

	results := [3]int{menu1, menu2, menu3}

	return results
}

func (this *Order) CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
