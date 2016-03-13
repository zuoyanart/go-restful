## restful 框架
这是一个根据neo的restful框架编写的一个restful流程部署应用框架，包含:

 > 1. api文档生成，使用基于nodejs的apidoc，配置文件apidoc.json
 > 2. restful 路由采用框架 [neo](https://github.com/ivpusic/neo "neo"),当然使用[gin](https://github.com/gin-gonic/gin "gin")也是可以的，并且很容易做切换
 > 3. 数据验证：[https://github.com/go-playground/validator](https://github.com/go-playground/validator "validator")
 > 4. orm使用gorm

## 文档目录结构
-- controller //控制器

-- logic //业务逻辑

-- model //数据模型

-- tools //工具类

--apidoc.json //api文档生成配置

--conf.toml //neo配置

--main.go

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build gopkg.in/alecthomas
