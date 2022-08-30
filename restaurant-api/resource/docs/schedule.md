# 规划安排整个项目开发步骤

## 系统大概设置

### 数据结构设计
- Admin 管理表, 管理员登陆后可管理订餐系统
    - id
    - name
    - phone [unique]
    - password
- Shop 店铺表
    - id
    - name [unique]
    - logo
- Menu 菜单表
    - id
    - name 
- Food 菜品
    - id
    - name
    - price
    - image_url
    - stock
    - meun_id
- Order_Master 订单主表
    - id
    - amount
    - table_id
    - status
    - customer_id
- Order_detail 订单明细
    - id
    - quantity
    - unit_price
    - order_id
    - food_id
    - food_name
    - food_image
- Table 桌子
    - id
    - code [unique]
    - max_seat
- Customer 消费者
    - id
    - phone [unique]
    - username
    - password 
    - avatar

--- 为了简洁先这样设计吧

```对于基础的信息表，只创建一些基础的路由```


## 开发规划

### 1. 使用docker构建mongodb服务

### 2. mongodb 创建索引
- 定义一个方法，通过集合和字段名创建索引
- 完成

### 3. 日志输出
- 完成

### 4. 验证入参
- 完成

### 5. 密码加密
- 完成

### 6. Token
- 完成

### 7. Middleware
- auth - 完成
- 限流控制 - 完成
- 统一超时 - 完成
- 记录每一个请求入参出参 - 完成 - 部分数据获取不到

### 8. Refresh Token
- 创建一个session集合 - 完成
- 存/读session - 完成
- 必要时，多配置一个jwt密钥

### 9. 处理monggodb错误

### 10. 规划响应数据结构（异常处理，正确响应）

### 11. 使用docker构建 Redis 服务， 模拟发送短信
- 

### 12. 专注业务开发

### 13. 文件上传

### 14. 建立前端后台管理项目

### 15. 建立前端前台项目

### 16. 微信授权

### 17. 测试支付

### 18. 链路追踪

### 19. 申请并阿里/腾讯短信验证

### 20. 申请并使用阿里/腾讯/七牛云OSS

