<?php

class A
{
    private $a = 5;

    public function __construct($param = 5) //注意这个默认值
    {
        $this->a = $param;
    }

    public function display()
    {
        echo $this->a;
    }
}

class B
{
    private $obj;

    public function __construct(A $a)
    {
        $this->obj = $a;
    }

    public function display()
    {
        $this->obj->display();
    }
}

class C
{
    private $b;

    public function __construct(B $b)
    {
        $this->b = $b;
    }

    public function display()
    {
        $this->b->display();
    }
}

function getDependencies(array $params)
{
    $dependencies = [];
    foreach ($params as $value) {

        //ReflectionParameter的getClass(); 获取类型提示类，返回ReflectionClass
        $dependency = $value->getClass();
        if (is_null($dependency)) {
            var_dump($dependency);exit;
            if ($value->isDefaultValueAvailable()) {
                $dependencies[] = $value->getDefaultValue();
            } else {
                throw new Exception('没有默认值');
            }
        } else {
            $dependencies[] = getObject($value->getClass()->name);
            var_dump($dependencies, '===');
        }
    }

    return $dependencies;
}
function getObject(string $className)
{
    $reflector = new ReflectionClass($className);
    if (!$reflector->isInstantiable()) {
        throw new Exception('不可实例化');
    }

    //获取类的构造函数 返回 ReflectionMethod
    $constructor = $reflector->getConstructor();
    if (is_null($constructor)) {
        return new $className;
    }

    //获取构造函数的参数,返回ReflectionParameter
    $dependencies = $constructor->getParameters();

    $instances = getDependencies($dependencies);
    return $reflector->newInstanceArgs($instances);
}


try {
    getObject('C')->display();
} catch (\Throwable $e) {
    var_dump($e->getTraceAsString());
}