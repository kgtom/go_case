### 本节目录


### 一、运行 mongo容器
~~~

# tom @ tom-pc in / [9:36:49]
$ docker run -p 27017:27017 -v /tempMongo/db:/data/db -d mongo
efb9032a2a9f3164dad23bd9be32af25c0d9b545629d5dee49535e2e725dd347

# tom @ tom-pc in / [9:36:54]
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                      NAMES
efb9032a2a9f        mongo               "docker-entrypoint.s…"   4 seconds ago       Up 3 seconds        0.0.0.0:27017->27017/tcp   optimistic_minsky

~~~
* -v：挂载宿主机目录/tempMongo/db 到容器/data/db目录，避免容器重启后，数据丢失。
* -d: 后台运行容器，并返回容器ID

### 二、命令行终端运行mongodb客户端

~~~
# tom @ tom-pc in / [9:37:27]
$ docker run -it mongo mongo --host 192.168.0.197 --port 27017
MongoDB shell version v4.0.4
connecting to: mongodb://192.168.0.197:27017/
Implicit session: session { "id" : UUID("199bfada-caf4-4155-9fc7-fe3663cf04d7") }
MongoDB server version: 4.0.4
Welcome to the MongoDB shell.
.....

> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
otw     0.000GB
> use otw
switched to db otw
> show collections
crm
> db.crm.find()
{ "_id" : ObjectId("5c3e8b0f73c09dc24bb323a3"), "id" : "crm001", "type" : 1, "name" : "客户甲", "xxx_nounkeyedliteral" : {  }, "xxx_unrecognized" : BinData(0,""), "xxx_sizecache" : 0 }
{ "_id" : ObjectId("5c3e8b0f73c09dc24bb323a5"), "id" : "crm002", "type" : 2, "name" : "客户乙", "xxx_nounkeyedliteral" : {  }, "xxx_unrecognized" : BinData(0,""), "xxx_sizecache" : 0 }
> ^C
bye
2019-01-16T01:39:22.580+0000 E -        [main] Error saving history file: FileOpenFailed: Unable to open() file /home/mongodb/.dbshell: Unknown error

~~~

### 三、使用mongo-express管理mondodb

 * 语法 docker run -it --rm -p 8081:8081 --link <mongo容器ID>:mongo mongo-express

~~~
# tom @ tom-pc in / [9:39:56] C:125
$  docker run -it --rm -p 8081:8081 --link efb9032a2a9f:mongo mongo-express

Unable to find image 'mongo-express:latest' locally
latest: Pulling from library/mongo-express
......
Waiting for mongo:27017...
Welcome to mongo-express
------------------------


Mongo Express server listening at http://0.0.0.0:8081
Server is open to allow connections from anyone (0.0.0.0)
basicAuth credentials are "admin:pass", it is recommended you change this in your config.js!
Database connected
Admin Database connected

~~~
* 浏览器查看：http://localhost:8081

### 四、使用 mongoclient 管理mongodb

* 格式：docker run --name mongoclient -d -p 3000:3000 -v /tmp/db:/data/db -e MONGO_URL=mongodb://宿主机ip:27017/  mongoclient/mongoclient
~~~
# tom @ tom-pc in / [11:33:59]
$ docker run --name mongoclient -d -p 3000:3000 -v /tmp/db:/data/db -e MONGO_URL=mongodb://192.168.0.197:27017/  mongoclient/mongoclient
51c5a623ab67512ebb47b81d38d0a2c53670965e8329e7cd89a8b762cd735c4d

# tom @ tom-pc in / [11:34:02]
$ docker ps
CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS              PORTS                      NAMES
51c5a623ab67        mongoclient/mongoclient   "./entrypoint.sh nod…"   5 seconds ago       Up 4 seconds        0.0.0.0:3000->3000/tcp     mongoclient
7f66469ed437        mongo-express             "tini -- /docker-ent…"   2 hours ago         Up 2 hours          0.0.0.0:8081->8081/tcp     mystifying_goldwasser
efb9032a2a9f        mongo                     "docker-entrypoint.s…"   2 hours ago         Up 2 hours          0.0.0.0:27017->27017/tcp   optimistic_minsky

~~~
* 访问 http://localhost:3000/

