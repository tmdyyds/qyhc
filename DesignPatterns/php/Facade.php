<?php
/*
门面模式为子系统提供一组统一的高层接口,让子系统更易用

场景:
    1 封装底层复杂实现,对外提供易用的接口
    2 解决性能,客户端调用服务端多个接口产生的网络通信耗时,Facade减少请求

TIP1:
    门面(外观)模式做接口整合，解决的是多接口调用带来的问题。
    适配器是做接口转换，解决的是原接口和目标接口不匹配的问题。
*/

//外观模式 提供了访问子系统功能的function
class Facade
{
    public function fetchFunc1()
    {
        return (new Subsystem1)->func1();
    }

    public function fetchFunc2()
    {
        return (new Subsystem2)->func2();
    }

    //组合方式
    public function fetchFunc3()
    {
        return (new Subsystem2)->func2() . (new Subsystem1)->func1();
    }
}

class Subsystem1
{
    public function func1()
    {
        echo '子系统1';
    }
}

class Subsystem2
{
    public function func2()
    {
        echo '子系统2';
    }
}

//...其他子系统SubsystemN

//请求方
$n = new Facade;
$n->fetchFunc3();

