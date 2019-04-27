
### 一、修改线上db，t_new_orders_N 表 在status字段后面增加字段order_total

~~~
delimiter $$

drop procedure if exists AddOrdersColumn$$   

create procedure AddOrdersColumn()
begin
 declare i int;        
 set i = 1;     
 while i <= 10 do          
  SET @STMT :=CONCAT("alter table t_new_orders_",i," ADD column order_total int(10) unsigned DEFAULT NULL after status");

  PREPARE STMT FROM @STMT;   
  EXECUTE STMT; 
  set i = i +1;
end while;
end;
$$
call AddOrdersColumn();

~~~


### 二、 新增表内容

~~~
delimiter $$

drop procedure if exists pro10 $$   

create procedure pro10()
     begin
     declare i int;
     set i=1;
     while i<5 do
         insert into t1(val) values(i);
         set i=i+1;
    
end while;
end;
$$
call pro10();

~~~



