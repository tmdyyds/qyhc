<?php
/*
策略模式定义一族算法类，将每个算法分别封装起来，让它们可以互相替换。策略模式可以使算法的变化独立于使用它们的客户端（这里的客户端代指使用算法的代码）

TIP1:
    侧重于算法(行为)实现
    实现主要通过接口
    目的实现方便地替换不同的算法类
    策略模式侧重“策略”或“算法”这个特定的应用场景，用来解决根据运行时状态从一组策略中选择不同策略的问题，而工厂模式侧重封装对象的创建过程，这里的对象没有任何业务场景的限定，可以是策略，但也可以是其他东西。从设计意图上来，这两个模式完全是两回事儿。
    在策略模式中，不同的策略具有相同的目的、不同的实现、互相之间可以替换
*/

//策略接口
interface DiscountStrategy
{
    public function calDiscount();
}

//常规策略(策略具体实现)
class NormalDiscountStrategy implements DiscountStrategy
{
    public function calDiscount()
    {
        return '标准折扣';
    }
}

//组团策略
class GrouponDiscountStrategy implements DiscountStrategy
{
    public function calDiscount()
    {
        return '团购价';
    }
}

//利用工厂模式
class DiscountStrategyFactory
{
    //有if else时 可以通过反射去掉if else
/*    public function getDiscountStrategy($type)
    {
        if ($type == 1) {
            return new GrouponDiscountStrategy;
        } elseif ($type == 2) {
            return new NormalDiscountStrategy;
        } else {
            // ....
        }
    }*/

    public function getDiscountStrategy($class)
    {
        $reflector = new ReflectionClass($class);
        // $constructor = $reflector->getConstructor(); //没有construct故不做其他处理
        if ($reflector->isInstantiable()) {
            $class = $reflector->getName();
            return new $class;
        }

        throw new \Exception("Error Processing Request", 1);

    }
}

//if else
/*$c = new DiscountStrategyFactory;
$strategy = $c->getDiscountStrategy('1');
var_dump($strategy->calDiscount());*/


$c = new DiscountStrategyFactory;
$strategy = $c->getDiscountStrategy('GrouponDiscountStrategy');
var_dump($strategy->calDiscount());




