<?php
/*
桥连模式将抽象和实现解耦，让它们可以独立变化(这个模式提倡组合由于继承...)(这太JB抽象了...)

TIP1:
    如果你想要拆分或重组一个具有多重功能的庞杂类 （例如能与多个数据库服务器进行交互的类）， 可以使用桥接模式
    如果你希望在几个独立维度上扩展一个类， 可使用该模式
*/

//根据消息发送方式维度和消息紧急级别2个维度拆分
//发送消息 "实现"基类
interface MsgSender
{
    public function send(string $msg);
}

//具体实现类
class TelephoneMsgSender implements MsgSender
{
    public function send(string $msg)
    {
        return '发送手机消息';
    }
}

class EmailMsgSender implements MsgSender
{
    public function send(string $msg)
    {
        return '发送邮件消息';
    }
}
//其他消息


//"抽象"
abstract class Notification
{
    public $msgSender;
    public function __construct(MsgSender $msgSender)
    {
        $this->msgSender = $msgSender;
    }

    abstract public function notify(string $msg);
}

//严重消息通知
class SevereNotification extends Notification
{
    public function notify(string $msg)
    {
        return $this->msgSender->send($msg);
    }
}

//普通消息
class TrivialNotification extends Notification
{
    public function notify(string $msg)
    {
        return $this->msgSender->send($msg);
    }
}

$telephoneMsgSender = new TelephoneMsgSender;


$notify = new SevereNotification($telephoneMsgSender);
var_dump($notify->notify("通知"));exit;





