<?php
/*
简单工厂模式由工厂对象决定创建哪一种产品类,适合创建产品类种类比较少的情况
工厂方法模式定义了一个创建对象的接口，但由子类决定要实例化的类是哪一个。工厂方法让类把实例化推迟到子类
抽象工厂模式提供一个接口，用于创建相关或依赖对象的家族，而不需要明确指定具体类
TIP1:
    当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，我们推荐使用工厂方法模式，将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。而使用简单工厂模式，将所有的创建逻辑都放到一个工厂类中，会导致这个工厂类变得很复杂。
*/

class RuleConfigSource
{
    private $ruleConfigParserFactory;

    public function load()
    {
        //其他

        $this->init();

        //后缀名
        $ext = 'json';

        //根据文件名后缀 调用不同解析文件的类
        $parse = $this->ruleConfigParserFactory->createParser($ext);
        ;

        //其他逻辑

        //这里直接返回
        return $parse->parse();
    }

    private function init()
    {
        if (!$this->ruleConfigParserFactory) {
            $this->ruleConfigParserFactory = new RuleConfigParserFactory;
        }
    }
}

//简单工厂模式
class RuleConfigParserFactory
{
    //在new class不复杂的情况下 直接放在简单工厂
    /*有多处if分支判断逻辑，违背开闭原则，但权衡扩展性和可读性，这样的代码实现在大多数情况下（比如，不需要频繁地添加parser，也没有太多的 parser）是没有问题的。*/
    public function createParser(string $fileExtension)
    {
        if ($fileExtension == 'json') {
            $parse = new JsonRuleConfigParser;
        } elseif ($fileExtension == 'xml') {
            $parse = new XmlRuleConfigParser;
        } else {
            $parse = new YamlRuleConfigParser;
        }

        return $parse;
    }
}

//start
//工厂方法模式
interface IRuleConfigParserFactory
{
    public function createParser();
}

//工厂方法模式利用多态 去掉简单工厂中if else
//但是if else 会放在调用方(耦合进Load方法中)
//json解析
class JsonRuleConfigParserFactory implements IRuleConfigParserFactory
{
    public function createParser()
    {
        return new JsonRuleConfigParser;
    }
}

//xml解析
class XmlRuleConfigParserFactory implements IRuleConfigParserFactory
{
    public function createParser()
    {
        return new XmlRuleConfigParser;
    }
}

//其他

//end

//start
//抽象工厂方法 特殊场景 比如上面都是根据文件后缀来分类解析,现在有多种分类方式
//抽象工厂模式可以简单理解为帮你组装好了一个生产类的流水线
interface IConfigParserFactory
{
    public function createRuleParser(); //根据文件后缀
    public function createSystemParser(); //根据系统
}

class JsonConfigParserFactory implements IConfigParserFactory
{
    public function createRuleParser()
    {
        return new JsonRuleConfigParser;
    }

    public function createSystemParser()
    {
        return new JsonSystemConfigParser;
    }
}
//end

//解析类
class JsonRuleConfigParser
{
    //真正解析
    public function parse()
    {
        //其他

        //解析成功
        return true;
    }
}

$c = new RuleConfigSource;
var_dump($c->load());