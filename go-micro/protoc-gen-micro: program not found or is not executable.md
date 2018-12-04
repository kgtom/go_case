
## 在make bulid 时遇到： protoc-gen-micro: program not found or is not executable 

~~~
# tom @ tom-pc in ~/goprojects/src/shop-micro/service/user-service on git:master x [23:11:50]
$ make build
protoc --proto_path=/Users/tom/goprojects/src:. --micro_out=. --gofast_out=. proto/*.proto
protoc-gen-micro: program not found or is not executable
--micro_out: protoc-gen-micro: Plugin failed with status code 1.
make: *** [proto] Error 1

~~~
两方面考虑:
### gopath问题
* 1.安装protoc-gen-go 完成后，确保在protoc-gen-go 在(当前gopath的）src/bin中，
~~~
go get -u github.com/golang/protobuf/protoc-gen-go
~~~
* 2.确保当前gopath配置环境变量中
~~~
export GOPATH=/Users/tom/goprojects
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN
~~~
或者
~~~
export PATH=$PATH:$GOPATH/bin
~~~

### go-micro 及protoc-gen-micro 重新安装一下
~~~
go get github.com/micro/go-micro
go get -u -v github.com/micro/protoc-gen-micro
~~~

