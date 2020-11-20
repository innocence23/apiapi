## 模块
- git管理
- curd操作
- 配置模块
- 日志模块
- mysql初始化
- web中间件
- 参数验证
- 分页和多条件搜索
- 接口文档
- 目录结构
    + router middleware
    + controller
    + service
    + repository
    + model
    + dto
    + Util
    + runtime, config ...
    
# 正文

###初始化项目 及目录机构
- git init
- go mod init zyf/go_prefer_api
- touch main.go
- touch Makefile
- mkdir -p router router/middleware controller service repository model dto pkg util doc runtime test config 


###配置文件 viper
- [viper](https://github.com/spf13/viper)
- pkg.config

###日志文件 logrus
- [logrus](https://github.com/sirupsen/logrus)
- pkg.xlog

###mysql数据库 gorm
- [gorm](http://gorm.book.jasperxu.com/)
- pkg.database

###gin web 入口文件
- [gin](https://github.com/gin-gonic/gin)

###增删改查
- model
- repository
- service
- controller
- 设置路由


### 验证参数


  
  

    
