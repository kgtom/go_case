### 一、go-callvis
提供可视化go函数调用关系图，便于了解项目结构，尤其是大型项目。

* 安装
~~~
go get -u github.com/TrueFurby/go-callvis
cd $GOPATH/src/github.com/TrueFurby/go-callvis && make
~~~
* 使用
~~~
# tom @ tom-pc in ~/goprojects/src [15:07:09] C:130
$ go-callvis -file= otwm

2019/01/23 15:07:31 http serving at http://localhost:7878
2019/01/23 15:07:31 converting dot to svg..
2019/01/23 15:07:32 serving file: /var/folders/_k/h8fwv1gx54576mgnkdk4nvjh0000gq/T/go-callvis_export.svg
~~~

### 二、revie
go语言代码质量检查工具。
* 安装
~~~
 # tom @ tom-pc in ~/goprojects/src [16:26:22] C:1
$ export http_proxy=http://127.0.0.1:1087;export https_proxy=http://127.0.0.1:1087;

# tom @ tom-pc in ~/goprojects/src [16:33:49]
$ go get -u github.com/mgechev/revive
~~~
* 使用
~~~
# tom @ tom-pc in ~/goprojects/src [17:34:07] C:1
$ revive -config otwm/default.toml -formatter friendly otwm
  ⚠  https://revive.run/r#var-naming  don't use underscores in Go names; func page_not_found should be pageNotFound
  otwm/main.go:35:6

  ⚠  https://revive.run/r#var-naming  don't use underscores in Go names; func page_note_permission should be pageNotePermission
  otwm/main.go:42:6

  ⚠  https://revive.run/r#exported  exported var FilterUser should have comment or be unexported
  otwm/main.go:23:5

⚠ 3 problems (0 errors, 3 warnings)

Warnings:
  2  var-naming
  1  exported


~~~

### 三、
