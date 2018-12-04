
## 1.解决 protoc-gen-micro: program not found or is not executable 

~~~
# tom @ tom-pc in ~/goprojects/src/shop-micro/service/user-service on git:master x [23:11:50]
$ make build
protoc --proto_path=/Users/tom/goprojects/src:. --micro_out=. --gofast_out=. proto/*.proto
protoc-gen-micro: program not found or is not executable
--micro_out: protoc-gen-micro: Plugin failed with status code 1.
make: *** [proto] Error 1

~~~
两方面考虑:
