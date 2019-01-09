# Golang点餐小程序

## 数据库使用sqlite3 

## 创建数据库
> sqlite3 OrderMenu.db  

## 创建表 
> 再当前根目录下创建 sqlite3 OrderMeal.db  

```
CREATE TABLE "OrderMenu" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "username" VARCHAR(64),
  "menu" VARCHAR(64),
  "created" DATE,
  "user_id" integer DEFAULT 0,
  "status" integer DEFAULT 0,
  "menu_id" integer DEFAULT 1
);

CREATE INDEX "user_id"
ON "OrderMenu" (
  "user_id" ASC
);
```
unrecognized import path "golang.org/x/net/html"
解决办法：
1：本身就是挂着vpn的，可以打开google，github等各种网站，就是下载不下来；
2：参考文章做了下，果然搞定，文章链接：当go get遇到墙时
操作的办法就是：
$mkdir -p $GOPATH/src/golang.org/x/
$cd $GOPATH/src/golang.org/x/
$git clone https://github.com/golang/net.git net 
$go install net
安装了这个net包，然后再去安装gin，就安装成功了。
