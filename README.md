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


