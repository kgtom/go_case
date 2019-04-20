
### grep 常用命令
#### 1.grep 实现or的操作
~~~
第一种：grep -E 'a|b' test.log
第二种： grep -e 'a' -e 'b' text.log
第三种： grep -E 'a*|b*' test.log
~~~


#### 2.grep 实现and 的操作
~~~
grep Query* test.log |grep  tom
tail -f test.log | grep getUserInfo
~~~



#### 3.grep 实现not 反查的操作
~~~
tail -f text.log | grep -v info | grep -v debug
~~~
