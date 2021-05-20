## docker 

(base) xx-Pro :: ~/xxx » docker run -it --rm --name=centos_dev -p 8010:8010 -p 8210:8210 -p 8310:8310  -v $PWD/xxx/:/root/xxx centos:7.5.1804 /bin/bash

yum install -y git which wget gcc gcc-c++ automake autoconf libtool make zlib-devel net-tools psmisc
yum -y install telnet.x86_64

tar  xvf yasm-1.3.0.tar.gz
     ls
     cd yasm-1.3.0
     ls
    ./configure
    make
make install



 wget https://golang.org/dl/go1.15.12.linux-amd64.tar.gz
tar xvf go1.15.12.linux-amd64.tar.gz  -C /usr/local
vim /etc/profile

export GOROOT=/usr/local/go
export GOPATH=/root/goproject
export PATH=$PATH:$GOROOT/bin

source /etc/profile

删除未运行的容器： docker rm $(sudo docker ps -a -q)

删除未使用镜像： docker image prune -a
