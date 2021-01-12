## 重点知识
#### 命令
- go build：用于测试编译包，在项目目录下生成可执行文件（有main包）,不能生成包文件, go install 可以生成包文件
- go install：主要用来生成库和工具。一是编译包文件（无main包），将编译后的包文件放到 pkg 目录下（$GOPATH/pkg）。二是编译生成可执行文件（有main包），将可执行文件放到 bin 目录（$GOPATH/bin）,生成可执行文件在当前目录下， go install 生成可执行文件在bin目录下（$GOPATH/bin）

#### 环境变量
- $GOROOT 表示 Go 在你的电脑上的安装位置，它的值一般都是 $HOME/go，当然，你也可以安装在别的地方
- $GOARCH 表示目标机器的处理器架构，它的值可以是 386( 32 位系统)、amd64(64位) 或 arm
- $GOOS 表示目标机器的操作系统，它的值可以是 darwin、freebsd、linux 或 windows
- $GOBIN 表示编译器和链接器的安装位置，默认是 $GOROOT/bin，如果你使用的是 Go 1.0.3 及以后的版本，一般情况下你可以将它的值设置为空，Go 将会使用前面提到的默认值
- $GOPATH 默认采用和 $GOROOT 一样的值，但从 Go 1.1 版本开始，你必须修改为其它路径。它可以包含多个 Go 语言源码文件、包文件和可执行文件的路径，而这些路径下又必须分别包含三个规定的目录：src、pkg 和 bin，这三个目录分别用于存放源码文件、包文件和可执行文件
- $GOARM 专门针对基于 arm 架构的处理器，它的值可以是 5 或 6，默认为 6
- $GOMAXPROCS 用于设置应用程序可使用的处理器个数与核数
目标机器是指你打算运行你的 Go 应用程序的机器。

Go编译器支持交叉编译，也就是说你可以在一台机器上构建运行在具有不同操作系统和处理器架构上运行的应用程序，也就是说编写源代码的机器可以和目标机器有完全不同的特性（操作系统与处理器架构）。

为了区分本地机器和目标机器，你可以使用 $GOHOSTOS 和 $GOHOSTARCH 设置本地机器的操作系统名称和编译体系结构，这两个变量只有在进行交叉编译的时候才会用到，如果你不进行显示设置，他们的值会和本地机器（$GOOS 和 $GOARCH）一样。

**GOROOT目录文件**: 
- /bin：包含可执行文件，如：编译器，Go 工具
- /doc：包含示例程序，代码工具，本地文档等
- /lib：包含文档模版
- /misc：包含与支持 Go 编辑器有关的配置文件以及 cgo 的示例
- /os_arch：包含标准库的包的对象文件（.a）
- /src：包含源代码构建脚本和标准库的包的完整源代码（Go 是一门开源语言）
- /src/cmd：包含 Go 和 C 的编译器和命令行脚本

#### import
- 标准库包引入时,路径从GOROOT
- import自定义引包时,路径从GOPATH的src下开始,不用带src,编译器会自动从src中开始引入
#### go生成目标文件
- window env GOOS=windows GOARCH=amd64 go build -o wangbin.exe main.go

#### channel(关闭后,能读不能写)
- channe遍历使用for range
    + 在遍历时,若channel没关闭,则报deadlock错误
    + 如果channel已关闭,则会正常遍历,遍历完后就会退出遍历
    
#### Slice
- 检查slice是否为空(用len(s) == 0, s != nil时也会有为空)
```
var s []int //len(s) == 0, s == nil
s = nil //len(s) == 0, s == nil
s = []int(nil) //len(s) == 0, s == nil
s = []int{} //len(s) == 0, s != nil
```
- slice两个冒号c[:3:5]
```
c := []int{1, 2, 3, 4, 5, 6, 7, 8}
e := c[:3:5] //从0到3为长度,5是定义容量。如果c[1:3:5],则长度为2，容量为4
fmt.Println(e, len(e), cap(e)) //[1 2 3] 3 5
```

#### 类型断言
```
//如果将一个接口类型变量断言成一个指针类型的变量，在断言成功的前提下，两个变量将共享内存空间
package main

import "fmt"

func main() {
    var a = 34
    var i interface{} = &a

    o := i.(*int)

    fmt.Println(i, o)
}

0xc0000120a8 0xc0000120a8
```

#### 反射
- Type是类型,Kind是类别,type和kind可能相同也可能不同
    + var num int = 10 num的type是int，kind也是int
    + var stu Student(student是struct) stu的type是包名.Student,Kind是 struct
