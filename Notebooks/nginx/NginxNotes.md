<?php

//nginx日常记录

//nginx&fpm 502 504
http://www.stelin.me/2017/06/09/php-fpm%E8%AF%A6%E8%A7%A3 //fastcgi包含 502 504情况

### nginx服务

#### 启动停止nginx命令 kill SINGAL PID/nginx.pid路径
- kill HUP PID 平滑重启
- kill INT|TERM|QUIT PID 快速停止
- kill WINCH PID 平缓停止workprocess,nginx平滑升级
- kill USR1 PID 重新打开log,用户log切割
- kill USR2 PID 使用新版本Nginx文件启动服务,之后平缓停止原有Nginx进程,即平滑升级
- nginx -g INT|TERM(快速停止)|QUIT(平缓)|HUP(-c filepath)

#### nginx服务参数
- nginx -c