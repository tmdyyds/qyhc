<?php
/*
观察者模式在对象之间定义一个一对多的依赖，当一个对象状态改变的时候，所有依赖的对象都会自动收到通知

TIP1:
    观察者和被观察者类比laravel框架中事件:监听&被监听
*/

//观察者基类 定义接口数据
interface RegObserver
{
    public function handleRegSuccess(int $uid);
}

//发送通知
class RegNotificationObserver implements RegObserver
{
    private $notificationService;

    public function __construct()
    {
        $this->notificationService = new NotificationService;
    }

    public function handleRegSuccess(int $uid)
    {
        return '消息已发送成功';
    }
}

//发优惠券
class RegCouponsObserver implements RegObserver
{
    private $couponsService;

    public function __construct()
    {
        $this->couponsService = new CouponsService;
    }

    public function handleRegSuccess(int $uid)
    {
        return '优惠券已发送成功';
    }
}

class NotificationService
{
}

class CouponsService
{
}

class UserController
{
    public function getObserverClass()
    {
        //可以放在配置文件中也可以单独加入到内存中
        return [
            new RegNotificationObserver,
            new RegCouponsObserver,
        ];
    }

    public function register($telephone, $passwd)
    {
        //注册类逻辑
        $uid = 1;

        //这里可以解耦 同步阻塞
        //异步阻塞 可以用mq队列 或者重启一个线程 golang可以重启一个goroutine
        foreach ($this->getObserverClass() as $observer) {
            $observer->handleRegSuccess($uid);
        }

        return true;
    }
}

$u = new UserController;
var_dump($u->register('test', '123456'));



//下面出处
//https://refactoringguru.cn/design-patterns/observer/php/example#example-1

// class UserRepository implements \SplSubject
// {
//     /**
//      * @var array The list of users.
//      */
//     private $users = [];

//     // Here goes the actual Observer management infrastructure. Note that it's
//     // not everything that our class is responsible for. Its primary business
//     // logic is listed below these methods.

//     /**
//      * @var array
//      */
//     private $observers = [];

//     public function __construct()
//     {
//         // A special event group for observers that want to listen to all
//         // events.
//         $this->observers["*"] = [];
//     }

//     private function initEventGroup(string $event = "*"): void
//     {
//         if (!isset($this->observers[$event])) {
//             $this->observers[$event] = [];
//         }
//     }

//     private function getEventObservers(string $event = "*"): array
//     {
//         $this->initEventGroup($event);
//         $group = $this->observers[$event];
//         $all = $this->observers["*"];

//         return array_merge($group, $all);
//     }

//     public function attach(\SplObserver $observer, string $event = "*"): void
//     {
//         $this->initEventGroup($event);

//         $this->observers[$event][] = $observer;
//     }

//     public function detach(\SplObserver $observer, string $event = "*"): void
//     {
//         foreach ($this->getEventObservers($event) as $key => $s) {
//             if ($s === $observer) {
//                 unset($this->observers[$event][$key]);
//             }
//         }
//     }

//     public function notify(string $event = "*", $data = null): void
//     {
//         $event = '*';
//         echo "UserRepository: Broadcasting the '$event' event.\n";
//         foreach ($this->getEventObservers($event) as $observer) {
//             $observer->update($this, $event, $data);
//         }
//     }

//     // Here are the methods representing the business logic of the class.

//     public function initialize($filename): void
//     {
//         echo "UserRepository: Loading user records from a file.\n";
//         // ...
//         $this->notify("users:init", $filename);
//     }

//     public function createUser(array $data): User
//     {
//         echo "UserRepository: Creating a user.\n";

//         $user = new User();
//         $user->update($data);

//         $id = bin2hex(openssl_random_pseudo_bytes(16));
//         $user->update(["id" => $id]);
//         $this->users[$id] = $user;

//         $this->notify("users:created", $user);

//         return $user;
//     }

//     public function updateUser(User $user, array $data): User
//     {
//         echo "UserRepository: Updating a user.\n";

//         $id = $user->attributes["id"];
//         if (!isset($this->users[$id])) {
//             return null;
//         }

//         $user = $this->users[$id];
//         $user->update($data);

//         $this->notify("users:updated", $user);

//         return $user;
//     }

//     public function deleteUser(User $user): void
//     {
//         echo "UserRepository: Deleting a user.\n";

//         $id = $user->attributes["id"];
//         if (!isset($this->users[$id])) {
//             return;
//         }

//         unset($this->users[$id]);

//         $this->notify("users:deleted", $user);
//     }
// }

// /**
//  * Let's keep the User class trivial since it's not the focus of our example.
//  */
// class User
// {
//     public $attributes = [];

//     public function update($data): void
//     {
//         $this->attributes = array_merge($this->attributes, $data);
//     }
// }

// /**
//  * This Concrete Component logs any events it's subscribed to.
//  */
// class Logger implements \SplObserver
// {
//     private $filename;

//     public function __construct($filename)
//     {
//         $this->filename = $filename;
//         if (file_exists($this->filename)) {
//             unlink($this->filename);
//         }
//     }

//     public function update(\SplSubject $repository, string $event = null, $data = null): void
//     {
//         $entry = date("Y-m-d H:i:s") . ": '$event' with data '" . json_encode($data) . "'\n";
//         file_put_contents($this->filename, $entry, FILE_APPEND);

//         echo "Logger: I've written '$event' entry to the log.\n";
//     }
// }

// /**
//  * This Concrete Component sends initial instructions to new users. The client
//  * is responsible for attaching this component to a proper user creation event.
//  */
// class OnboardingNotification implements \SplObserver
// {
//     private $adminEmail;

//     public function __construct($adminEmail)
//     {
//         $this->adminEmail = $adminEmail;
//     }

//     public function update(\SplSubject $repository, string $event = null, $data = null): void
//     {
//         // mail($this->adminEmail,
//         //     "Onboarding required",
//         //     "We have a new user. Here's his info: " .json_encode($data));

//         echo "OnboardingNotification: The notification has been emailed!\n";
//     }
// }

// /**
//  * The client code.
//  */

// $repository = new UserRepository();
// $repository->attach(new Logger(__DIR__ . "/log.txt"), "*");
// $repository->attach(new OnboardingNotification("1@example.com"), "users:created");

// $repository->initialize(__DIR__ . "/users.csv");

// // ...

// $user = $repository->createUser([
//     "name" => "John Smith",
//     "email" => "john99@example.com",
// ]);

// // ...

// $repository->deleteUser($user);









