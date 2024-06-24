## gin
文档： [x](https://gin-gonic.com/zh-cn/docs/)

## gorm
文档： [x](https://gorm.io/zh_CN/docs)

### 查询数据操作
```go
// ....省略了数据库连接以及表结构的定义
var res []Todo
db.First(&res)
```