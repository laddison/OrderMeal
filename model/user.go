package model

type User struct {
}

type ShowAlert struct {
	Show    bool
	Content string
}

func (this *User) GetUser() map[int]string {
	user := make(map[int]string)
	user[0] = "请选择名字"
	user[1] = "余航"
	user[2] = "刘羽"
	user[3] = "郑剑峰"
	user[4] = "陈枞"
	user[5] = "范孝荣"
	//user[6] = "黄业炜"
	user[7] = "吴志钢"
	user[8] = "李秋水"
	//user[9] = "丛磊"
	user[10] = "郑陈菲"
	user[11] = "傅锎若"
	user[12] = "吴磊"
	user[13] = "兰增伟"
	user[14] = "刘雪奇"
	user[15] = "王志敏"
	user[16] = "林伟"
	user[17] = "蔡广运"
	user[18] = "赵世昌"
	user[19] = "冯德伟"
	user[20] = "郑少琼"
	user[21] = "刘宗斌"
	//user[22] = "郑小龙"
	user[23] = "陈静康"
	user[24] = "陈鸿"
	user[25] = "陈东海"
	user[26] = "孙德冰"
	user[27] = "金林平"
	return user
}

func (this *User) GetNelabUser() map[int]string {
	user := make(map[int]string)
	user[0] = "请选择名字"
	user[1] = "郑及罕"
	user[2] = "郑俊"
	user[3] = "赖俊林"
	user[4] = "张玉润"
	user[5] = "金楠"
	user[6] = "许彬彬"
	user[7] = "茅煦鹏"
	user[8] = "范郑唐"
	user[9] = "李春青"
	user[10] = "杨海鹏"
	user[11] = "陈林键"
	user[12] = "连金祥"
	user[13] = "孙德冰"
	user[14] = "陈永校"
	user[15] = "陈润发"
	user[16] = "洪景雯"
	user[17] = "陈弘晟"
	user[18] = "李国仁"
	return user
}

func (this *User) GetMenu() map[int]string {
	menu := make(map[int]string)
	menu[1] = "A餐"
	menu[2] = "B餐"
	menu[3] = "C餐"
	menu[4] = "D餐"
	menu[5] = "E餐"

	return menu
}
