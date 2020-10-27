<?php
/*
适配器模式(类适配和对象适配)让原本不兼容而不能一起工作的类可以一起工作

TIP1:
    类适配器使用继承关系来实现，对象适配器使用组合关系来实现
*/

/*
如果 Adaptee 接口并不多，那两种实现方式都可以。

如果 Adaptee 接口很多，而且 Adaptee 和 ITarget接口定义大部分都相同，那我们推荐使用类适配器，因为 Adaptor 复用父类 Adaptee 的接口，比起对象适配器的实现方式，Adaptor 的代码量要少一些。

如果 Adaptee 接口很多，而且 Adaptee 和 ITarget接口定义大部分都不相同，那我们推荐使用对象适配器，因为组合结构相对于继承更加灵活。(摘自极客时间某课程)

TIP1:
    门面(外观)模式做接口整合，解决的是多接口调用带来的问题。
    适配器是做接口转换，解决的是原接口和目标接口不匹配的问题。
    适配器模式能为被封装对象提供不同的接口， 代理模式能为对象提供相同的接口， 装饰模式则能为对象提供加强的接口。
*/

//如下是摘自极客时间设计模式之美中一段代码
// 类适配器: 基于继承
/*public interface ITarget {
  void f1();
  void f2();
  void fc();
}

public class Adaptee {
  public void fa() { //... }
  public void fb() { //... }
  public void fc() { //... }
}

public class Adaptor extends Adaptee implements ITarget {
  public void f1() {
    super.fa();
  }

  public void f2() {
    //...重新实现f2()...
  }

  // 这里fc()不需要实现，直接继承自Adaptee，这是跟对象适配器最大的不同点
}

// 对象适配器：基于组合
public interface ITarget {
  void f1();
  void f2();
  void fc();
}

public class Adaptee {
  public void fa() { //... }
  public void fb() { //... }
  public void fc() { //... }
}

public class Adaptor implements ITarget {
  private Adaptee adaptee;

  public Adaptor(Adaptee adaptee) {
    this.adaptee = adaptee;
  }

  public void f1() {
    adaptee.fa(); //委托给Adaptee
  }

  public void f2() {
    //...重新实现f2()...
  }

  public void fc() {
    adaptee.fc();
  }
}*/

/*
封装有缺陷的接口设计
统一多个类的接口设计
替换依赖的外部系统
兼容老版本接口
适配不同格式的数据
*/

//外部过滤词系统A
class WordsFilterAdapteeA
{
    //text是原始文本，函数输出用***替换敏感词之后的文本
    public function filterSexyWords($text)
    {
        echo '***' . $text;
    }
}

//外部过滤词系统B
class WordsFilterAdapteeB
{
    //text是原始文本，函数输出用***替换敏感词之后的文本
    public function filterSexyWordsB($text)
    {
        echo '***' . $text;
    }
}

//外部过滤词系统C
/*class WordsFilterAdapteeC
{

}*/

//目标过滤词
interface WordsFilterItarget
{
    public function filter();
}

//适配对象
class AWordsFilterAdaptor implements WordsFilterItarget
{
    public function filter()
    {
        return (new WordsFilterAdapteeA)->filterSexyWords('A');
    }
}

class BWordsFilterAdaptor implements WordsFilterItarget
{
    public function filter()
    {
        return (new WordsFilterAdapteeB)->filterSexyWords('B');
    }
}

class Client
{
    public $itarget;

    //基于接口而非实现
    public function __construct(WordsFilterItarget $itarget)
    {
        $this->itarget = $itarget;
    }

    public function reqFilterWord()
    {
        //其他逻辑
        $this->itarget->filter();
    }
}


$c = new Client(new AWordsFilterAdaptor);
$c->reqFilterWord();








