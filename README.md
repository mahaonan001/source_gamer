

# source_gamer web框架




 
## 目录

- [source\_gamer web框架](#source_gamer-web框架)
  - [目录](#目录)
          - [开发前的配置要求](#开发前的配置要求)
          - [**安装步骤**](#安装步骤)
    - [文件目录说明](#文件目录说明)
    - [路由说明](#路由说明)
    - [使用到的框架](#使用到的框架)
    - [贡献者](#贡献者)
      - [如何参与开源项目](#如何参与开源项目)
    - [版本控制](#版本控制)
    - [作者](#作者)
    - [鸣谢](#鸣谢)




###### 开发前的配置要求
1. 需要配备go1.12及以上版本
2. 一款可以编写go语言的编译器或者任意命令行工具
3. 需要安装mysql
###### **安装步骤**

```sh
git clone https://github.com/mahaonan001/source_gamer.git
```
1. cd source_gamer
2. go mod init source_gamer&go mod tidy
3. Mac or Linux
- go build -o server main.go
- ./server
1. Windows
- go build -o server.exe main.go  
- .\server.exe


### 文件目录说明


```
|   .gitignore
|   go.mod
|   go.sum
|   main.go
|   README.md
|
+---common
|       db.go
|       jwt.go
|
+---config
|       config.yml
|
+---controller
|       admin.go
|       cg_info.go
|       info.go
|       login.go
|       register.go
|       send_mail.go
|
+---mail
|       mail.go
|
+---middle
|       auth_token.go
|       a_token.go
|
+---model
|       admin.go
|       user.go
|
+---response
|       response.go
|
+---router
|       router.go
|
\---utils
        util.go

```
### 路由说明
```
http://localhost/api/get_code #获取验证码 需要参数：email
http://localhost/api/user/register  #用户注册  需要参数：email，name，password，code（验证码）
http://localhost/api/user/login #用户登录 需要参数：email，password 返回bear token
http://localhost/api/user/info #获取个人资料 需要参数：bear token


```


### 使用到的框架

- [Gin](https://github.com/gin-gonic/gin)
- [Gorm](https://github.com/go-gorm/gorm.io)
- [Viper](https://github.com/spf13/viper)

### 贡献者

暂无
#### 如何参与开源项目

贡献使开源社区成为一个学习、激励和创造的绝佳场所。你所作的任何贡献都是**非常感谢**的。


1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



### 版本控制

该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

### 作者

mahaonan001
1649801526@qq.com

 qq:1649801526    

 *您也可以在贡献者名单中参看所有参与该项目的开发者。*


### 鸣谢


- [shaojintian](https://github.com/shaojintian/Best_README_template)大佬的README.md模板




