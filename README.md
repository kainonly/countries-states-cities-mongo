# Countries States Cities Mongo

国家地区信息同步至 MongoDB，数据支持来自：

- https://github.com/dr5hn/countries-states-cities-database

镜像支持：

- ghcr.io/kainonly/cscm:latest
- hkccr.ccs.tencentyun.com/kainonly/cscm:latest（国内）

## 定时触发

推荐使用腾讯云云函数部署实现

## 环境变量

- **DATABASE_URI** MongoDB 连接 URI
- **DATABASE_DBNAME** 数据库名称