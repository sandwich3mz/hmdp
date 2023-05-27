# hmdp

#### 使用gin + ent 不完全重构黑马点评项目

#### tb_seckill_voucher表中添加了新主键id
#### 项目并未上传config.yaml配置文件
``` app:
  env: local # 环境名称
  port: 8081 # 服务监听该端口号
  app_name: gin-app # 应用名称
  app_url: http://localhost # 应用域名

database:
  driver: mysql # 数据库驱动
  host: 127.0.0.1 # 域名
  port: 3307 # 端口号
  database: hmdp # 数据库名称
  username: xxxx # 用户名
  password: xxxx # 密码
  charset: utf8mb4 # 编码格式
  max_idle_conns: 10 # 空闲连接池中连接的最大数量
  max_open_conns: 100 # 打开数据库连接的最大数量
  log_mode: info  # 日志级别
  enable_file_log_writer: true # 是否启用日志
  log_filename: sql.log # 日志文件名称

redis:
  host: 127.0.0.1
  port: 6379
  db: 0
  password: xxxxx
