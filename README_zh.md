# Golang RESTful Api

(转到 >> [英文说明](https://github.com/levblanc/golang-restful-api/blob/master/README.md))

本项目使用 Golang 和 Mongodb，搭建了一套简单的信息流系统的 RESTful API。

数据库使用了 [mLab](https://mlab.com) 的免费沙盒套餐，只有0.5G空间，所以项目仅作展示和测试用途。

(**所以如果项目运行起来后报任何数据库相关的问题，请在issue里告诉我，谢谢！**)

克隆项目到本地后，可跟随下面的指引运行起来。

[API 文档](https://github.com/levblanc/golang-restful-api/blob/master/README.md#api-docs) 和 [postman collection](https://github.com/levblanc/golang-restful-api/blob/master/README.md#make-request-now) 已经在项目中准备好，方便你的使用。

## Get Started

项目进行开发使用的 Go 版本如下：

```
$ go version
go version go1.11.5 darwin/amd64
```

对于其他 Go 版本的兼容尚未有机会进行测试（非常抱歉），但我想你起码需要把本机的 Go 版本升级到 `v1.11`，才能把项目跑起来。

项目中使用了下面这些 package：

- [xid](https://github.com/rs/xid): unique id generation
- [gorilla/mux](https://github.com/gorilla/mux): request routing
- [golang/x/crypto/bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt): password hashing
- [mgo](https://github.com/globalsign/mgo): mongodb driver
- [bson](https://godoc.org/github.com/globalsign/mgo/bson): BSON specification

由于墙的原因，`golang/x/crypto`在国内无法正常进行下载，所以我使用 `go mod vendor` 命令把 package 都下载了在 `src/vendor` 下。

## Run 
```bash
# cd 到项目目录下
$ cd /path/to/golang-restful-api
# 使用 mod flag 运行项目
$ go run -mod=vendor src/main.go
2019/02/05 11:16:09 Connected to mongodb database: mstream
2019/02/05 11:16:10 Server started at: http://127.0.0.1:8080
```

## API Docs

我使用 [ApiDoc](http://apidocjs.com) 生成了 API 文档。

在浏览器中打开 `doc/index.html` 即可查看。

![api-doc](images/api-doc.png)

## Routes, Cookie and Session

项目使用 cookie 进行用户验证。为了方便测试，用户 session 时长设置为**30分钟**。

到目前为止实现了下列的 API: 

- user: signup, login, logout, get one/all users
- post: CURD
- comment: create

## Make Request Now!

Postman collection `mstream.postman_collection.json` 放在项目根目录下。

Import 到你的 Postman app 即可开始测试了。

![postman](images/postman.png)
