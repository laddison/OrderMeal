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
	http.HandleFunc("/nelab", showNelab)
	http.HandleFunc("/save", saveOrder)
	http.HandleFunc("/upd", cronUpdate)
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

//首页展示
func showIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	type data struct {
		UserName       map[int]string
		MenuName       map[int]string
		ListOrder      []model.OrderData
		TotalMenus     [5]int
		ThreeOrderList []model.OrderData
		GroupName      string
	}

	orderModel := model.Order{}
	userModel := model.User{}
	list := orderModel.GetList("services")
	GroupName := "services"

	ShowData := data{userModel.GetUser(), userModel.GetMenu(), list, orderModel.GetSumTotal(list), orderModel.GetPreThre("services"), GroupName}
	t.Execute(w, ShowData)
}

// showNelab 游戏点餐
func showNelab(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")

	type data struct {
		UserName       map[int]string
		MenuName       map[int]string
		ListOrder      []model.OrderData
		TotalMenus     [5]int
		ThreeOrderList []model.OrderData
		GroupName      string
	}

	orderModel := model.Order{}
	userModel := model.User{}
	list := orderModel.GetList("nelab")
	GroupName := "nelab"

	ShowData := data{userModel.GetNelabUser(), userModel.GetMenu(), list, orderModel.GetSumTotal(list), orderModel.GetPreThre("nelab"), GroupName}
	t.Execute(w, ShowData)
}

//保存订单
func saveOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	groupName := r.Form.Get("groupName")

	orderModel := model.Order{}
	userModel := model.User{}

	var userList map[int]string
	if groupName == "services" {
		userList = userModel.GetUser()
	} else {
		userList = userModel.GetNelabUser()
	}

	menuList := userModel.GetMenu()

	userId, err := strconv.Atoi(r.Form.Get("name"))
	orderModel.CheckErr(err)
	if userId == 0 {
		fmt.Fprint(w, 3)
		return
	}

	menuId, err := strconv.Atoi(r.Form.Get("menu"))
	orderModel.CheckErr(err)

	create_time := time.Now().Format("2006-01-02 15:04:05")

	info := orderModel.GetInfoByUid(userId, groupName)
	if info.Id > 0 {
		orderModel.Update(info.Id, menuList[menuId], userId, menuId)

		fmt.Fprint(w, 1)
		return
	} else {
		orderModel.Insert(userList[userId], menuList[menuId], create_time, userId, menuId, groupName)
		fmt.Fprint(w, 2)
		return
	}
}

//更新状态
func cronUpdate(w http.ResponseWriter, r *http.Request) {
	orderModel := model.Order{}
	affectNum := orderModel.UpdateStatus()
	fmt.Fprint(w, affectNum)
}
