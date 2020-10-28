<?php
    //mysql基础

    sql在mysql中执行流程: sql语句->查询缓存(mysql8.0已删除)->解析器->优化器->执行器
    show profile;获取上一次请求的执行时间(解析模块、优化模块等各个时间)(select @@profiling; set profiling=1)
    show profile for query 2(查询id query_id) 指定查询id

    DDL: CREATE DROP ALTER //修改表结构
    DML: SELECT DELETE UPDATE //增删改查

    CREATE DATABASE test
    DROP DATABASE test

    DROP TABLE IF EXISTS test;
    CREATE TABLE test (
        //字段 数据类型
        id int(11) NOT NULL AUTO_INCREMENT,
        uid int(11) NOT NULL DEFAULT '0' COMMENT '',

        //索引
        PRIMARY KEY (id) USING BTREE,
        UNIQUE INDEX idx_uid(uid) USING BTREE //唯一索引
        KEY idx_uid(uid) //普通索引
    ) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='';
    //) ENGINE=InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

    //字段
    ALTER TABLE test ADD age int(11) unsigned; //添加字段
    ALTER TABLE test CHANGE old_uid new_id int(11) NOT NULL; //修改字段名
    ALTER TABLE test MODIFY COLUMN uid bigint(20) NOT NULL; //修改字段类型
    ALTER TABLE test DROP COLUMN uid; //删除字段

    //索引
    ALTER TABLE test ADD PRIMARY_KEY (column)
        | UNIQUE (column)
        | INDEX index_name (column)
        | FUllTEXT (column)
        | INDEX index_name (column, column ...)
    CREATE INDEX|UNIQUE... index_name ON test(column)
    //删除索引
    ALTER TABLE test DROP index index_name
    ALTER TABLE test DROP primary key
    DROP index index_name on test

    //增删改查
    SELECT ... FROM ... WHERE ... GROUP BY ... HAVING ... ORDER BY ... LIMIT ...
    //某课程讲解select执行顺序
    /*
        SELECT DISTINCT player_id, player_name, count(*) as num #顺序5
        FROM player JOIN team ON player.team_id = team.team_id #顺序1
        WHERE height > 1.80 #顺序2
        GROUP BY player.team_id #顺序3
        HAVING num > 2 #顺序4
        ORDER BY num DESC #顺序6
        LIMIT 2 #顺序7

        完整的SELECT语句内部执行顺序是：
        1、FROM子句组装数据（包括通过ON进行连接）
        2、WHERE子句进行条件筛选
        3、GROUP BY分组
        4、使用聚集函数进行计算；
        5、HAVING筛选分组；
        6、计算所有的表达式；
        7、SELECT 的字段；
        8、ORDER BY排序
        9、LIMIT筛选
    */

    //视图
    CREATE VIEW view_name AS SELECT column1, column2 FROM table WHERE condition
    ALTER VIEW view_name AS SELECT column1, column2 FROM table WHERE condition //修改视图
    DROP TABLE view_name

    //存储过程
    CREATE PROCEDURE 存储过程名称([参数列表])
    BEGIN
        需要执行的语句
    END
    /*
    //例子
    DELIMITER //
    CREATE PROCEDURE `add_num`(IN n INT)
    BEGIN
           DECLARE i INT;
           DECLARE sum INT;

           SET i = 1;
           SET sum = 0;
           WHILE i <= n DO
                  SET sum = sum + i;
                  SET i = i +1;
           END WHILE;
           SELECT sum;
    END //
    DELIMITER ;
    */
    DROP TABLE view_name //删除存储过程

    //隔离级别
    脏读: A事务读到B事务修正但是还没有commit的数据,读未提交的数据,称为脏读
    不可重复读: 对某数据进行读取,两次读取的结果不同,这因为其他事务对改数据进行了修改或删除
    幻读(专指插入数据): 在事务中,前后两次查询到符合条件数据总数不一致,其他事务插入数据



TIP: 临时表只在当前连接存在,关闭连接后,临时表就会自动释放

    //优化

    COUNT(*) = COUNT(1) > COUNT(字段)
    //ORDRE BY
    SQL中，可以在WHERE子句和ORDER BY子句中使用索引，目的是在WHERE子句中避免全表扫描，ORDER BY子句避免使用FileSort排序。
    尽量Using Index完成ORDER BY排序。如果WHERE和ORDER BY相同列就使用单索引列；如果不同使用联合索引。
    无法Using Index时，对FileSort方式进行调优。
    //IN 和 between
    在索引字段后，慎用IN和NOT IN，如果是连续的数值，可以考虑用BETWEEN进行替换


