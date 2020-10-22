<?php
/*
单利模式一个类只有一个实例,并提供一个访问该实例的全局节点
场景: 1 处理资源访问冲突, 2 表示全局唯一类(例如配置文件、DB连接)

TIP: 单利模式分懒汉模式和恶汉模式,php中我只见过懒汉模式,恶汉模式在一些语言中会出现(比如JAVA)
*/

/**
 * 1 类中添加私有变量 保存单例实例
 * 2 声明一个公用静态方法用于获取单利实例
 * 3 构造函数私有化,其他对象不能调用,类的静态方法仍然可以调用
 * 4 禁止克隆和反序列化类
 */
class Singleton
{
    //保存单例实例的成员变量必须被声明为静态类型
    private static $instances = [];

    //私有化 防止new,但是单利类还会执行构造函数
    private function __construct()
    {

    }

    //禁止克隆
    protected function __clone() {}

    public function __wakeup()
    {
        throw new \Throwable('禁止反序列化');
    }

    //对外提供访问对象的实例的方法 且必是唯一方式
    public static function getInstance()
    {
        //延迟初始化
        $subclass = static::class;
        if (!self::$instances[$subclass]) {
            self::$instances[$subclass] = new static;
        }

        return self::$instances[$subclass];
    }
}

//日志写入文件
class Logger extends Singleton
{
    public static function Log(string $message)
    {
        $logger = static::getInstance();
        $logger->writeLog($message);
    }

    private function writeLog(string $message)
    {
        //可以写入本地文件，此处直接输出内容
        echo $message . '...';
    }
}

//日志写入mysql
class Mysql extends Singleton
{

}

Logger::Log('log写入文件');

