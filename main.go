package main

import (
	"OrderMeal/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

// 路由
func main() {
	http.HandleFunc("/", showIndex)
	http.HandleFunc("/save", saveOrder)
	http.HandleFunc("/upd", cronUpdate)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	err := http.ListenAndServe(":8293", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//首页展示
func showIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	type data struct {
		UserName   map[int]string
		MenuName   map[int]string
		ListOrder  []model.OrderData
		TotalMenus [3]int
	}

	orderModel := model.Order{}
	userModel := model.User{}
	list := orderModel.GetList()

	ShowData := data{userModel.GetUser(), userModel.GetMenu(), list, orderModel.GetSumTotal(list)}
	t.Execute(w, ShowData)
}

//保存订单
func saveOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	orderModel := model.Order{}
	userModel := model.User{}

	userList := userModel.GetUser()
	menuList := userModel.GetMenu()

	userId, err := strconv.Atoi(r.Form.Get("name"))
	orderModel.CheckErr(err)

	menuId, err := strconv.Atoi(r.Form.Get("menu"))
	orderModel.CheckErr(err)

	create_time := time.Now().Format("2006-01-02 15:04:05")

	info := orderModel.GetInfoByUid(userId)
	if info.Id > 0 {
		orderModel.Update(menuList[menuId], userId, menuId)

		fmt.Fprint(w, 1)
	} else {
		orderModel.Insert(userList[userId], menuList[menuId], create_time, userId, menuId)

		fmt.Fprint(w, 2)
	}
}

//更新状态
func cronUpdate(w http.ResponseWriter, r *http.Request) {
	orderModel := model.Order{}
	affectNum := orderModel.UpdateStatus()
	fmt.Fprint(w, affectNum)
}
