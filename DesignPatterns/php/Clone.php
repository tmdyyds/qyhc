<?php
/*
原型模式,如果对象的创建成本比较大（比如要经过复杂计算 排序 计算hash值 rpc或者数据库等比较慢的io中），而同一个类的不同对象之间差别不大（大部分字段都相同），在这种情况下，我们可以利用对已有对象（原型）进行复制（或者叫拷贝）的方式来创建新对象，以达到节省创建时间的目的
*/

//php深浅拷贝
class Test
{
    public $test1 = 1;
}

class Test2
{
    public $test2=111;

    public function __construct(){
        $this->obj = new Test();
    }
}

// $n = new Test;
/*$m = $n; //浅拷贝
$m->test1 = 2;
$n->test1 = 20;*/

/*$m = $n; //浅拷贝
$m->test2 = 10;
$m->obj->test1 = 111111;*/


//test2的属性为深拷贝,属性对象为浅拷贝
/*$n = new Test2;
$p = clone $n;//深拷贝
$p->test2 = 10;
$p->obj->test1 = 111111;*/

//若要对象也实现深拷贝 则修改如下
/*public function __clone(){
    $this->obj = clone $this->obj;
}*/
// 或者序列化
//方法二，序列化反序列化实现对象深拷贝
/*$n = serialize($m);
$n = unserialize($n);

var_dump($p, $n);*/



//原型模式摘自https://refactoringguru.cn/design-patterns/prototype/php/example#example-1
class Prototype
{
    public $primitive;
    public $component;
    public $circularReference;

    /**
     * PHP has built-in cloning support. You can `clone` an object without
     * defining any special methods as long as it has fields of primitive types.
     * Fields containing objects retain their references in a cloned object.
     * Therefore, in some cases, you might want to clone those referenced
     * objects as well. You can do this in a special `__clone()` method.
     */
    public function __clone()
    {
        $this->component = clone $this->component;

        // Cloning an object that has a nested object with backreference
        // requires special treatment. After the cloning is completed, the
        // nested object should point to the cloned object, instead of the
        // original object.
        $this->circularReference = clone $this->circularReference;
        $this->circularReference->prototype = $this;
    }
}

class ComponentWithBackReference
{
    public $prototype;

    /**
     * Note that the constructor won't be executed during cloning. If you have
     * complex logic inside the constructor, you may need to execute it in the
     * `__clone` method as well.
     */
    public function __construct(Prototype $prototype)
    {
        $this->prototype = $prototype;
    }
}

/**
 * The client code.
 */
function clientCode()
{
    $p1 = new Prototype();
    $p1->primitive = 245;
    $p1->component = new \DateTime();
    $p1->circularReference = new ComponentWithBackReference($p1);

    $p2 = clone $p1;
    if ($p1->primitive === $p2->primitive) {
        echo "Primitive field values have been carried over to a clone. Yay!\n";
    } else {
        echo "Primitive field values have not been copied. Booo!\n";
    }
    if ($p1->component === $p2->component) {
        echo "Simple component has not been cloned. Booo!\n";
    } else {
        echo "Simple component has been cloned. Yay!\n";
    }

    if ($p1->circularReference === $p2->circularReference) {
        echo "Component with back reference has not been cloned. Booo!\n";
    } else {
        echo "Component with back reference has been cloned. Yay!\n";
    }

    if ($p1->circularReference->prototype === $p2->circularReference->prototype) {
        echo "Component with back reference is linked to original object. Booo!\n";
    } else {
        echo "Component with back reference is linked to the clone. Yay!\n";
    }
}

clientCode();