<?php
/*
建造者(生成器)模式使你能够分步骤创建复杂对象。 该模式允许你使用相同的创建代码生成不同类型和形式的对象

TIP1:
    工厂模式创建不同但是相关类型的对象(继承同一父类或接口的一组子类)，由给定的参数来决定创建哪种类型的对象。建造者模式是用来创建一种类型的复杂对象，通过设置不同的可选参数，“定制化”地创建不同的对象
TIP2:
    生成器重点关注如何分步生成复杂对象。 抽象工厂专门用于生产一系列相关对象。 抽象工厂会马上返回产品， 生成器则允许你在获取产品前执行一些额外构造步骤。
*/

//汽车(产品)
class Car
{
    public $seats  = 4;
    public $engine = '超级引擎';
    public $gps    = '牛逼牌gps';
}

//汽车手册
class Manual
{
    //属性省略
}

//构造者 生成器接口声明了创建产品对象不同部件的方法
interface Builder
{
    public function setSeats();
    public function setEngine();
    public function setGPS();
    //...其他省略
}

//具体生成器 实现产品具体生成过程
//提供构造过程的不同实现
class CarBuilder implements Builder
{
    private $car;

    public function __construct()
    {
        $this->reset();
    }

    //一个新的生成器实例必须包含一个在后续组装过程中使用的空产品对象
    public function reset()
    {
        $this->car = new Car;
    }

    public function setSeats()
    {
        //这里可以判断一些属性限制 大小、类型等
        $this->car->seats = 8;

    }

    public function setEngine()
    {
        $this->car->engine = '生成超级引擎';
    }

    public function setGPS()
    {
        $this->car->gps = '生成GPS';
    }

    public function getProduct()
    {
        $product = $this->car;
        $this->reset();

        //这里可以处理一些属性逻辑 例如属性大小、类型等各种限制

        return $product;
    }
}

class CarManual implements Builder
{
    private $manual;

    public function __construct()
    {
        $this->reset();
    }

    //一个新的生成器实例必须包含一个在后续组装过程中使用的空产品对象
    public function reset()
    {
        $this->manual = new Manual;
    }

    public function setSeats()
    {
        $this->manual->seats = '四个座位德国牛皮';
    }

    public function setEngine()
    {
        $this->manual->engine = '引擎德国纯手工制作';
    }

    public function setGPS()
    {
        $this->manual->gps = 'GPS林志玲姐姐语音提示';
    }

    public function getProduct()
    {
        $product = $this->manual;

        return $product;
    }
}

//主管 主管只负责按照特定顺序执行生成步骤(可有可无)
//定义调用构造步骤的顺序， 这样你就可以创建和复用特定的产品配置。
class Direcotr
{
    //主管可使用同样的生成步骤创建多个产品变体。
    public function constructSportsCar(Builder $builder)
    {
        $builder->reset();
        $builder->setSeats();
        $builder->setEngine();
        $builder->setGPS();

        return $builder;
    }

    public function constructSuv(Builder $builder)
    {
        //....省略
    }
}

$director = new Direcotr;
$builder = new CarBuilder;

//主管类组装生成过程(如何生成)
$director->constructSportsCar($builder);

//返回成品(组装后的产品)
var_dump($builder->getProduct());
