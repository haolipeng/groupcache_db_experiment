# groupcache_db_experiment
groupcache实验代码

组件介绍
1、slowdb
模仿慢速数据库操作
写入操作：瞬时写入
读取操作：300ms延迟

2、frontend
实现了Get读取操作缓存中内容
开启了rpc，通过默认的(80001 + 1000)端口可以访问Get函数

3、dbserver
实现了Set将内容写入数据库
实现了Get读取数据库内容
开启了rpc，通过默认的9090端口可以访问Set和Get函数


