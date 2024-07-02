# 任务

## 思路：
### 业务代码服务器
1. 提供一个服务
2. 里面需要一个定时任务
3. 定时任务帮助检查ssl是否过期
4. 过期了就帮助申请
5. 还需要帮助替换nginx的配置，以及重启nginx

### cdn/oss 这里需要用我自己的服务器
1. 提供一个界面，让用户填写基本信息并注册。
2. 代申请ssl证书
3. 调用对应的ssl证书上传接口
4. 在自己服务器上跑定时任务。 过期了就又从 第二步开始 

### 自己服务
1. 业务代码通过这个接口来上报，以及通知用户
2. 

## 拆解任务

1. 这个算是个脚本吧，还是需要在用户自己的服务器上跑
### 用户登录注册

### 定时任务查询证书是否过期

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

### 数据库设计

大概就是需要一个注册页面来注册。以及分别保存阿里云和七牛云的注册信息并加以区分。
#### 用户表
|     字段名     |      类型      |   描述   | 是否为NULL |          备注          |
|:-----------:|:------------:|:------:|:-------:|:--------------------:|
|     id      |     int      |  自增id  |    x    | 作为主键关联,autoincrement |
|   account   | varchar(128) |   账号   |    x    |          -           |
|    name     | varchar(255) |   名字   |    x    |          -           |
|  password   | varchar(128) |   密码   |    x    |         SM3          |
|   mobile    |     int      |  手机号   |    y    |        作为通知手段        |
|    email    |     int      |  手机号   |    y    |        作为通知手段        |
|   status    |     int      |   状态   |    x    |          -           |
| create_time |     date     |  创建时间  |    x    |          -           |
| update_time |     date     |  更新时间  |    x    |          -           |
| delete_time |     date     | 删除时间时间 |    x    |          -           |


#### 用户网站表

|     字段名     |      类型      |   描述   | 是否为NULL |   备注    |
|:-----------:|:------------:|:------:|:-------:|:-------:|
|     id      |     int      |  自增id  |    x    |    -    | 
|   user_id   |     int      | 用户表的id |    x    |    -    | 
|   origin    | varchar(255) |   来源   |    x    | oss/cdn |
|   domain    | varchar(255) |   网站   |    x    |    -    |
|     key     | varchar(255) |  配置名称  |    x    |    -    | 
|    value    | varchar(128) |  配置内容  |    x    |    -    | 
|   status    |     int      |   状态   |    x    |    -    |
| create_time |     date     |  创建时间  |    x    |    -    |
| update_time |     date     |  更新时间  |    x    |    -    |
| delete_time |     date     | 删除时间时间 |    x    |    -    |
