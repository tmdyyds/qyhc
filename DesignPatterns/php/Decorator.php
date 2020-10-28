<?php
/*
装饰器模式主要解决继承关系过于复杂的问题，通过组合来替代继承。它主要的作用是给原始类添加增强功能。这也是判断是否该用装饰器模式的一个重要的依据。除此之外，装饰器模式还有一个特点，那就是可以对原始类嵌套使用多个装饰器

TIP1:
    对功能的增强
    装饰器和原始类继承相同的父类,这样可以对原始类嵌套多个装饰器类

TIP2(代理与装饰器区别):
    Decorator关注为对象动态的添加功能, Proxy关注对象的信息隐藏及访问控制.
    Decorator体现多态性, Proxy体现封装性.
*/

//原数据接口
interface DataSource
{
    //对数据加密和解压
    public function readData();
}

//接口具体实现逻辑
class CacheDataSource implements DataSource
{
    public function readData()
    {
        //从缓存取数据
        return 'cache data';
    }
}

//装饰器基类继承原数据接口 具体装饰器基于该基类
class DataSourceDecorator implements DataSource
{
    public $dataSource;

    public function __construct(DataSource $dataSource)
    {
        $this->dataSource = $dataSource;
    }

    public function readData()
    {
        //具体装饰时 增加额外其他功能
        return $this->dataSource->readData();
    }
}

//具体装饰 加密解密
class EncryptionDecorator extends DataSourceDecorator
{
    public function readData()
    {
        //其他逻辑处理
        $data = parent::readData();
        //其他逻辑

        return '加密数据 = ' . $data;
    }
}

//解压缩
class CompressionDecorator extends DataSourceDecorator
{
    public function readData()
    {
        //其他逻辑处理
        $data = parent::readData();
        //其他逻辑

        return '解压 = ' . $data;
    }
}

//比较类似一个栈,先New 后执行
$c = new CacheDataSource;
$decorate = new EncryptionDecorator($c);
$decorate = new CompressionDecorator($decorate);

var_dump($decorate->readData());










