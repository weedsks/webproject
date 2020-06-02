# beego_blog
#### 配置文件
./conf/app.conf
```
appname = webproject
httpport = 8080
runmode = prod
sessionon = true
Graceful = true

TemplateLeft = "<#"
TemplateRight = "#>"

dbprefix = ks_
dbhost = 127.0.0.1
dbport = 3306
dbuser = 
dbpwd = 
dbname = 

#模块前缀
adminfix=admin
homefix=blog
pwdsalt=weeds

#七牛云账号
QiNiu_AccessKey=
QiNiu_SecretKey=
QiNiu_bucket=
QiNiu_domain=git
```
#### 数据初始化
导入sql.sql文件，配置数据库
#### 执行命令
```shell 
bee run
```
#### 地址
-- 前台
http://127.0.0.1:8080
-- 后台
http://127.0.0.1:8080/admin
#### 账号默认密码
weeds 123456
