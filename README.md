# Gin+Gorm+Mysql+Vue 货币汇率项目

## 环境准备

### MySQL+MySQL Workbench 下载

[MySQL下载链接](https://dev.mysql.com/downloads/mysql/)
<img width="1491" height="915" alt="image" src="https://github.com/user-attachments/assets/50297f98-d3fd-4d6e-8270-b949770b9172" />
这里可以选择**msi版本或者zip版本**<br>

我选择直接下载压缩包，msi就是有图形界面的下载版本（但是我的不知道为什么一直安装不了）<br>

[指路安装包教程](https://blog.csdn.net/a876106354/article/details/154486697?ops_request_misc=elastic_search_misc&request_id=f81b762bab9ca68f0eeeca9bb697d006&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~all~ElasticSearch~search_v2-2-154486697-null-null.142^v102^pc_search_result_base7&utm_term=mysql%2Bworkbench%E7%A4%BE%E5%8C%BA%E7%89%88%E5%AE%89%E8%A3%85%E6%95%99%E7%A8%8B&spm=1018.2226.3001.4187)<br>
按照这个教程成功<br>

[Workbench下载链接](https://dev.mysql.com/downloads/workbench/)<br>
按照指引下载即可<br>

### Postman下载
[下载链接](https://www.postman.com/downloads/)

## 前端准备
直接将`Exchangeapp_frontend`文件拉取<br>

## 后端
新建与`Exchangeapp_frontend`同一级的后端文件夹`Exchangeapp_backend`<br>
在该文件夹下创建`config`文件夹负责**配置文件**<br>
`config`下面接着新建一个`config.go`和`config.yml`文件<br>

### `config.yml`<br>
```
app:
  name: CurrencyExchangeApp
  port: :3000

database:
  host: localhost
  port: :3306
  user: your_username
  password: your_password
  name: currency_exchange_db
```
同时创建一个与`config`同级的文件`main.go`<br>

### 敲命令
`go mod init`<br>
`go get -u github.com/gin-gonic/gin`<br>
`go get github.com/spf13/viper`<br>

### [体验gin框架](https://gin-gonic.com/zh-cn/docs/quickstart/)
`main.go`
```
package main

import (
	"Exchangeapp/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```
接着输入命令`go run .`<br>
来到postman
<img width="2337" height="1452" alt="image" src="https://github.com/user-attachments/assets/ff6b35a6-ed6b-4471-9795-557e4315672e" />

### [路由管理](https://gin-gonic.com/zh-cn/docs/examples/http-method/)
新建`router`文件夹统一管理，下建`router.go`<br>

```
package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "login success",
			})
		})
		auth.POST("/register", func(ctx *gin.Context) {
			ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "register success",
			})
		})
	}
	return r
}
```

修改`main.go`<br>

```
package main

import (
	"Exchangeapp/config"
	"Exchangeapp/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	port := config.AppConfig.App.Port
	if port == "" {
		port = "8080"
	}
	r.Run(port) 
}
```
将router.go中的文件导入main.go,同时万一忘记了配置port,所以在main.go中加层防护<br>

### [Gorm+MySQL](https://gorm.io/zh_CN/docs/)
`go get -u gorm.io/gorm`
安装gorm<br>
在`config`下创建`db.go`<br>
```
package config

import (
	"Exchangeapp/global"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	dsn := AppConfig.Database.Dsn

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database,got err:%v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(AppConfig.Database.MaxIdleConns)
	sqlDB.SetMaxIdleConns(AppConfig.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		log.Fatalf("Failed to config database,got err:%v", err)
	}

	global.Db = db
}
```
连接数据库，一些重要配置应该写在`config.yml`中，详见文件<br>
创建全局文件`global`,再建`global.go`,将全局变量放在里面，比如数据库<br>

### 写注册功能
新建`controllers`文件下建`auth_controller.go`<br>
新建`models`文件下建`user.go`<br>
`user.go`
```
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model		//gorm框架下的模型
	Username string `gorm:"unique"`	//用户名，同时增加一个用户名唯一的验证
	Password string					//密码
}
```
