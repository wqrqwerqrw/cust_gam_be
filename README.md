# cust_gam_backend

科研训练大作业。
实现一个健身房管理系统。主要是B端服务。所以直接暴露了数据库的自增id。

项目实现方式比较简单。框架是gin和gorm。

设计思想遵守大厂的设计思想。

数据库配置：mysql

需要写一个db.json

````
{
  "user": 账号
  "pass": 密码
  "path": 路径,
  "database": 数据库名
}