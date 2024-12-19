# 微服务商城
## 框架使用 go-zero
## rpc服务
- 用户服务 user-rpc
- 商品服务 product-rpc
- 订单服务 order-rpc
- 支付服务 payment-rpc
- 库存服务 stock
- 购物车服务 cart
- 评论服务 comment
- 搜索服务 search
- 消息服务 message

## api服务
app有两层api
- app-api app相关的处理
- external-api 商城后台的操作

用户相关
- 用户注册
-- 存储到数据库的密码不是明文，需要是加密的
- 用户登录


12-19
整体框架搭好【mysql、redis、es、】
脚本、goctl自定义模板、common的一些插件