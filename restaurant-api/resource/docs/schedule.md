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

### 使用docker构建mongodb服务

### mongodb 创建索引
- 定义一个方法，通过集合和字段名创建索引
- 完成

### 日志输出
- 完成

### 验证入参
- 完成

### 密码加密
- 完成

### Token

### Middleware

### 使用docker构建 Redis 服务， 模拟发送短信

### 专注业务开发