<?php
//栈实现队列
class Queue2Stack
{
    private $queue;
    private $top; //栈顶元素

    public function __construct()
    {
        $this->queue = new SplQueue;
    }

    //删除栈顶元素并返回
    public function pop()
    {
        //1 2 3 队列元素 栈元素 3 1 2
        //这里不模拟转换 直接使用top 等获取
        $top = $this->queue->top();
        $this->queue->pop();

        $this->top = $this->queue->offsetGet($this->queue->count() - 1);

        return $top;
    }

    //入栈 入栈到栈顶
    public function push($x)
    {
        $this->queue->push($x);
        $this->top = $x;
    }

    //返回栈顶元素
    public function top()
    {
        return $this->top;
    }

    //栈是否为空
    public function isEmpty()
    {
        return $this->queue->isEmpty();
    }
}

$queue2Stack = new Queue2Stack;
for ($i=1; $i <= 3; $i++) {
    $queue2Stack->push($i);
}

$queue2Stack->pop();
var_dump($queue2Stack);exit;

//栈实现队列
class Stack2Queue
{
    private $stack1;
    private $stack2;

    public function __construct()
    {
        $this->stack1 = new SplStack;
        $this->stack2 = new SplStack;
    }

    //入队(队尾)
    public function push(int $data)
    {
        $this->stack1->push($data);
    }

    //出队(删除对头)并返回元素
    public function pop()
    {
        $this->peek(); //保证stack2不能为空

        return $this->stack2->pop();
    }

    //返回队头元素(保证stack2不能为空)
    public function peek()
    {
        //若队列为空 则把stack1数据移到stack2中
        if ($this->stack2->isEmpty()) {
            while (!$this->stack1->isEmpty()) {
                $this->stack2->push($this->stack1->pop());
            }
        }

        return $this->stack2->top();
    }

    //队列是否为空
    private function isEmpty()
    {
        //两个栈全为空
        return $this->stack1 && $this->stack2;
    }
}

$stack2queue = new Stack2Queue;

for ($i=0; $i < 10; $i++) {
    $stack2queue->push($i);
}

$stack2queue->pop();
var_dump($stack2queue->pop());