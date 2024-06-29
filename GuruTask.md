# 任务

## 拆解任务

### 用户登录注册

### 定时任务查询证书是否过期

### 证书的申请
1. 是否包含收费证书
2. 包含的收费对接哪家第三方支付
   1. 微信支付似乎支持了个人开发者的微信支付开通
   2. 这里有个看起来的坑是，说是只能用于线下实体业务
3. 证书申请的流程
   1. CRS文件制作
   2. 向CA发起请求验证以及获取私钥
   3. 验证
4. 实现阿里云。七牛云的cdn或者oss自动证书替换
   1. 七牛云CND https://developer.qiniu.com/fusion/8593/interface-related-certificate
   2. 阿里云OSS https://help.aliyun.com/zh/ssl-certificate/developer-reference/api-cas-2020-04-07-dir-csr-management/?spm=a2c4g.11186623.0.0.64cc26b4psSSRL
5. 免费域名提供商
   1. https://letsencrypt.org/zh-cn/getting-started/

### 封装

#### 封装后端工具代码 Gin

#### 封装前端js代码 用React