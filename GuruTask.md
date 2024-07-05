# 任务

## 思路：
### 业务代码服务器
1. 提供一个服务
2. 里面需要一个定时任务
3. 定时任务帮助检查ssl是否过期
4. 过期了就帮助申请
5. 还需要帮助替换nginx的配置，以及重启nginx
6. 这里CA证书需要sudo权限
7. 怎么找到nginx得配置也是个问题
8. 部署验证

### cdn/oss 这里需要用我自己的服务器
1. 提供一个界面，让用户填写基本信息并注册。
2. 代申请ssl证书
3. 调用对应的ssl证书上传接口
4. 在自己服务器上跑定时任务。 过期了就又从 第二步开始
5. 证书的部署成功验证(这里api一般都支持)

### 关于证书的归滚
1. 通过 domain/yyyymmdd 保存证书
2. 数据库证书网站一对多
3. 最新的一条作为当前证书
4. 全查出来让用户自己选择
5. 回滚时验证旧证书是否可用 

## 拆解任务

### 用户登录注册

### 证书的申请
1. 不包含收费证书
2. 证书申请的流程
   1. CRS文件制作
   2. 向CA发起请求验证以及获取私钥
   3. 验证
3. 实现阿里云。七牛云的cdn或者oss自动证书替换
   1. 七牛云CND https://developer.qiniu.com/fusion/8593/interface-related-certificate
   2. 阿里云OSS https://help.aliyun.com/zh/ssl-certificate/developer-reference/api-cas-2020-04-07-dir-csr-management/?spm=a2c4g.11186623.0.0.64cc26b4psSSRL
4. 免费域名提供商
   1. https://letsencrypt.org/zh-cn/getting-started/
5. 收费证书 TODO (暂时没找到id对接收费证书的文档。或许可以去收费证书官网看看)
   1. 设计包括规格
   2. 对接支付
   3. 订单
   4. 退费
6. 关于下单
   1. 全部走下单流程
   2. 免费的让她0元购，直接返回支付成功就是，这里主要校验，不要看价格是0元就0元
   3. 好吧一定要注意验证价格
   4. 

### 数据库设计

大概就是需要一个注册页面来注册。以及分别保存阿里云和七牛云的注册信息并加以区分。
#### 用户表 users

|     字段名     |      类型      |   描述   | 是否为NULL |          备注          |
|:-----------:|:------------:|:------:|:-------:|:--------------------:|
|     id      |     int      |  自增id  |    x    | 作为主键关联,autoincrement |
|   account   | varchar(128) |   账号   |    x    |          -           |
|    name     | varchar(255) |   名字   |    x    |          -           |
|  password   | varchar(128) |   密码   |    x    |         SM3          |
|    salt     | varchar(128) |   密码   |    x    |         SM3          |
|   mobile    | varchar(128) |  手机号   |    y    |        作为通知手段        |
|    email    | varchar(128) |   邮箱   |    y    |        作为通知手段        |
|   status    |     int      |   状态   |    x    |          -           |
| create_time |  timestamp   |  创建时间  |    x    |          -           |
| update_time |  timestamp   |  更新时间  |    x    |          -           |
| delete_time |  timestamp   | 删除时间时间 |    y    |          -           |


#### 类别基础表 categories

|     字段名     |      类型      |  描述  | 是否为NULL |          备注           |
|:-----------:|:------------:|:----:|:-------:|:---------------------:|
|     id      |    serial    | 自增id |    x    |           -           |
|    name     | varchar(255) |  名字  |    x    | ali_oss/qi`niu_cdn... |
|    desc     | varchar(255) |  简述  |    y    |                       |
|   pre_id    |     int      | 上级id |    x    |           -           |
|    pres     | varchar(255) | 所有上级 |    x    |         快速查询          |
|   status    |     int      |  状态  |    x    |           -           |
| create_time |  timestamp   | 创建时间 |    y    |           -           |
| update_time |  timestamp   | 更新时间 |    y    |           -           |

#### 类别表 category_info

|     字段名     |      类型      |   描述   | 是否为NULL |         备注         |
|:-----------:|:------------:|:------:|:-------:|:------------------:|
|     id      |    serial    |  自增id  |    x    |         -          |
| category_id | varchar(255) |   类别   |    x    |      oss/cdn       |
|  access_id  | varchar(255) | 第三方账号  |    y    |    看有没有必要做填充加解密    | 
| access_key  | varchar(255) | 第三方密码  |    y    |    看有没有必要做填充加解密    | 
|    auth     |     int      | 是否需要验证 |    x    | 需要鉴权，access两个都需要填写 | 
|   status    |     int      |   状态   |    x    |         -          |
| create_time |  timestamp   |  创建时间  |    y    |         -          |
| update_time |  timestamp   |  更新时间  |    y    |         -          |
1. 这里包括但是不限于，oss、cdn

#### 用户网站表 user_domains

|       字段名        |      类型      |   描述   | 是否为NULL | 备注 |
|:----------------:|:------------:|:------:|:-------:|:--:|
|        id        |    serial    |  自增id  |    x    | -  | 
|     user_id      |     int      | 用户表的id |    x    | -  |
|      domain      | varchar(255) |   网站   |    x    | -  |
| category_info_id | varchar(255) |   类别   |    x    | -  |
|      status      |     int      |   状态   |    x    | -  |
|   create_time    |  timestamp   |  创建时间  |    y    | -  |
|   update_time    |  timestamp   |  更新时间  |    y    | -  |


#### 资源表 resources

|       字段名        |      类型      |   描述   | 是否为NULL |       备注       |
|:----------------:|:------------:|:------:|:-------:|:--------------:|
|        id        |    serial    |  自增id  |    x    |       -        |
|     cert_id      | varchar(255) |  证书编号  |    x    |       -        |
|    cert_name     | varchar(255) | 证书文件名  |    x    |                | 
| private_key_name | varchar(255) | 私钥文件名  |    x    |                | 
|     location     | varchar(128) |   位置   |    x    | basePath/网站名字/ | 
|      status      |     int      |   状态   |    x    |       -        |
|   create_time    |  timestamp   |  创建时间  |    x    |       -        |
|   update_time    |  timestamp   | 删除时间时间 |    y    |       -        |

#### 网站/资源表 domain_resource

|     字段名     |    类型     |  描述  | 是否为NULL | 备注 |
|:-----------:|:---------:|:----:|:-------:|:--:|
|     id      |  serial   | 自增id |    x    | -  | 
|  domain_id  |    int    | 网站id |    x    | -  |
| resource_id |    int    | 资源id |    x    | -  |
|   status    |    int    |  状态  |    x    | -  |
| create_time | timestamp | 创建时间 |    y    | -  |
| update_time | timestamp | 更新时间 |    y    | -  |

#### 规格表-主表 sku

|     字段名     |      类型      |  描述  | 是否为NULL |       备注        |
|:-----------:|:------------:|:----:|:-------:|:---------------:|
|     id      |    serial    | 自增id |    x    |        -        | 
|    name     |     int      | 规格名称 |    x    |        -        |
|    desc     | VARCHAR(255) | 规格名称 |    x    |        -        |
|    path     | VARCHAR(255) | 规格名称 |    x    | 记录上下级关系。不要则递归生成 |
|   status    |     int      |  状态  |    x    |        -        |
| create_time |  timestamp   | 创建时间 |    y    |        -        |
| update_time |  timestamp   | 更新时间 |    y    |        -        |

#### 规格表-价格 sku_value

|     字段名     |    类型     |  描述  | 是否为NULL |      备注       |
|:-----------:|:---------:|:----:|:-------:|:-------------:|
|     id      |  serial   | 自增id |    x    |       -       | 
|   sku_id    |    int    | 规格id |    x    |       -       |
|    value    |    int    |  价格  |    x    |  整数保存，缩放100倍  |
|    stock    |    int    |  库存  |    x    | 拆分商品的库存到各个规格上 |
|   status    |    int    |  状态  |    x    |       -       |
| create_time | timestamp | 创建时间 |    y    |       -       |
| update_time | timestamp | 更新时间 |    y    |       -       |

#### 商品规格表 商品 <-> 多对多 goods_sku

|     字段名     |    类型     |  描述  | 是否为NULL |     备注      |
|:-----------:|:---------:|:----:|:-------:|:-----------:|
|     id      |  serial   | 自增id |    x    |      -      | 
|   sku_id    |    int    | 规格id |    x    |      -      |
|    value    |    int    |  价格  |    x    | 整数保存，缩放100倍 |
|   status    |    int    |  状态  |    x    |      -      |
| create_time | timestamp | 创建时间 |    y    |      -      |
| update_time | timestamp | 更新时间 |    y    |      -      |