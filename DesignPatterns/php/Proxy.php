<?php
/*
代理模式在不改变原始类（或叫被代理类）代码的情况下，通过引入代理类来给原始类附加功能

TIP1:
    原始类定义接口,代理类implements接口实现代理
    原始类没有定义接口(第三方接口), 代理类可以继承的方式实现
    动态代理可利用反射实现
TIP2:
    适配器模式能为被封装对象提供不同的接口， 代理模式能为对象提供相同的接口， 装饰模式则能为对象提供加强的接口。
*/

interface Downloader
{
    public function download(string $url): string;
}

//真正实现下载类
class SimpleDownloader implements Downloader
{
    public function download(string $url): string
    {
        $result = file_get_contents($url);
        echo "Downloaded bytes: " . strlen($result) . "\n";

        return $result;
    }
}

//代理类缓存
class CachingDownloader implements Downloader
{
    private $downloader;
    private $cache = [];

    public function __construct(SimpleDownloader $downloader)
    {
        $this->downloader = $downloader;
    }

    public function download(string $url): string
    {
        if (!isset($this->cache[$url])) {
            $this->cache[$url] = $this->downloader->download($url);
        }

        return $this->cache[$url];
    }
}

$realSubject = new SimpleDownloader;
$proxy = new CachingDownloader($realSubject);
$proxy->download('https://localhost/test.txt');

/*
interface IUserController
{
    public function login(String $telephone, String $password);
    public function register(String $telephone, String $password);
}

class UserController implements IUserController
{
    public function login(String $telephone, String $password)
    {
        echo 'is Login' . PHP_EOL;
    }

    public function register(String $telephone, String $password)
    {

    }
}

class MetricsCollector
{
    public function recordRequest($requestInfo)
    {
    }
}

class RequestInfo
{
    public function __construct($apiName, $responseTime, $startTimestamp)
    {
    }
}

class MetricsCollectorProxy
{
    private $proxiedObject;

    private $metricsCollector;

    public function __construct(MetricsCollector $metricsCollector)
    {
        $this->metricsCollector = $metricsCollector;
    }

    public function createProxy(object $object)
    {
        $this->proxiedObject = $object;
        return $this;
    }

    public function __call($method, $arguments)
    {
        $ref = new ReflectionClass($this->proxiedObject);
        if (!$ref->hasMethod($method))
            throw new Exception("method not existed");

        $method = $ref->getMethod($method);
        $startTimestamp = time();
        $userVo = $this->callMethod($method, $arguments);

        $endTimeStamp = time();
        $responseTime = $endTimeStamp - $startTimestamp;
        $requestInfo = new RequestInfo("login", $responseTime, $startTimestamp);
        $this->metricsCollector->recordRequest($requestInfo);

        return $userVo;
    }


    private function callMethod(\ReflectionMethod $method, $arguments)
    {
        //前置判断省略
        $method->invokeArgs($this->proxiedObject, $arguments);
    }

}

$proxy = new MetricsCollectorProxy(new MetricsCollector);
$userController = $proxy->createProxy(new UserController);
$userController->login(13800138000, 'pwd');*/

