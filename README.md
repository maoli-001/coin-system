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
  name: currency_exchange_db```
同时创建一个与`config`同级的文件`main.go`<br>
### 敲命令
`go mod init`<br>
`go get -u github.com/gin-gonic/gin`<br>
`go get github.com/spf13/viper`<br>
### 体验gin框架
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




