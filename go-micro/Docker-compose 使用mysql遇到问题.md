
## 问题
~~~
MySQL Connection Error: (2054) The server requested authentication method unknown to the client  
~~~
原因：
在MySQL 8.0以后版本，caching_sha2_password 是默认的身份验证插件而不是 mysql_native_password。问题在于密码验证不同了，前者安全性比后者更高。所以客户端连接出现问题。如果想使用之前的连接方式，则需要修改：
[官方文档](pluggable)
* 1.进入容器，修改 /etc/my.cnf 或者/etc/mysql/my.cnf

~~~
# tom @ tom-pc in ~/goprojects/src/shop-micro on git:master x [9:33:11]
$ docker exec -it shop-micro_mysql_1 bash
[root@03957200e1de /]# echo $MYSQL_ROOT_PASSWORD
123456
[root@03957200e1de etc]# cd my
my.cnf    my.cnf.d/ mysql/
[root@03957200e1de etc]# vi my.cnf
~~~

添加这一行配置：
~~~
[mysqld]
default_authentication_plugin=mysql_native_password
~~~

* 2.授权登录账号的密码验证插件
~~~
root@3b5ab69bc54e:~# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 17
Server version: 8.0.13 MySQL Community Server - GPL

Copyright (c) 2000, 2018, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> use mysql;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
mysql> ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';
Query OK, 0 rows affected (0.10 sec)

mysql> exit;
~~~

* 3.停止并重启容器
~~~

$ docker-compose stop

$ docker-compose restart

~~~
