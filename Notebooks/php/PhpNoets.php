<?php
/*CGI：是 Web Server 与 Web Application 之间数据交换的一种协议。
FastCGI: 是Web服务器(如：Nginx、Apache)和处理程序之间的一种通信协议，它是与Http类似的一种应用层通信协议，注意：它只是一种协议！同 CGI，是一种通信协议，但比 CGI 在效率上做了一些优化。同样，SCGI 协议与 FastCGI 类似。
PHP-CGI：是 PHP （Web Application）对 Web Server 提供的 CGI 协议的接口程序。
php-Fpm: 是 PHP（Web Application）对 Web Server 提供的 FastCGI 协议的接口程序，额外还提供了相对智能一些任务管理。一个master进程(fork)多个子进程(主进程通过发送信号的方式通知子进程),*/

//fastcgi&fpm
https://www.awaimai.com/371.html
https://www.miss77.net/2017/08/php-fpm-working-mechanism/
http://www.stelin.me/2017/06/09/php-fpm%E8%AF%A6%E8%A7%A3 //fastcgi包含 502 504情况
https://blog.wangzhipeng.top/2019/04/18/2019_04_18/
https://juejin.im/post/6844903784670298126