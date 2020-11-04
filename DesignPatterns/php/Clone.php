<?php


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

$n = new Test;
/*$m = $n; //浅拷贝
$m->test1 = 2;
$n->test1 = 20;*/

/*$m = $n; //浅拷贝
$m->test2 = 10;
$m->obj->test1 = 111111;*/


//test2的属性为深拷贝,属性对象为浅拷贝
$n = new Test2;
$p = clone $n;//深拷贝
$p->test2 = 10;
$p->obj->test1 = 111111;

//若要对象也实现深拷贝 则修改如下
/*public function __clone(){
    $this->obj = clone $this->obj;
}*/
// 或者序列化
//方法二，序列化反序列化实现对象深拷贝
$n = serialize($m);
$n = unserialize($n);

var_dump($p, $n);


