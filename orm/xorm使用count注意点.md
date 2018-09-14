
### Count方法

统计数据使用Count方法，Count方法的参数为struct的指针并且成为查询条件。

~~~
user := new(User)
total, err := engine.Where("id >?", user.id).Count(user)
~~~

假如 此时 user 做为方法参数传递过来，&user{id:1,name:"tom"},使用上面语句，最终形成sql中where id>1 and name="tom"，
count(user)会自动将user中其它字段自动追加到sql中， 所以可以使用下面语句：

~~~
user := new(User)
total, err := db.Table(f).Where("id>?", user.id).Count()

~~~

