#### 常用命令
-docker network create -d bridge test-net 新建网络 多个容器互相连接
    + docker run -itd --name test1 --network test-net ubuntu /bin/bash
    + docker run -itd --name test2 --network test-net ubuntu /bin/bash
    + d：参数指定 Docker 网络类型，有 bridge、overlay
- docker run ubuntu:15.10 /bin/echo "Hello world"
    + docker: Docker 的二进制执行文件
    + run: 与前面的 docker 组合来运行一个容器
    + ubuntu:15.10 指定要运行的镜像，Docker 首先从本地主机上查找镜像是否存在，如果不存在，Docker 就会从镜像仓库 Docker Hub 下载公共镜像
    + /bin/echo "Hello world": 在启动的容器里执行的命令
    + **Docker 以 ubuntu15.10 镜像创建一个新容器，然后在容器里执行 bin/echo "Hello world"**
- (进行交互)docker run -i -t (-d) ubuntu:15.10 /bin/bash
    + t: 在新容器内指定一个伪终端或终端
    + i: 允许你对容器内的标准输入 (STDIN) 进行交互
    + d: 启动容器（后台模式）
    + P(随机): 容器内部使用的网络端口随机映射到我们使用的主机上
        * docker run -d -P training/webapp python app.py
    + p(固定): docker run -d -p 5000:5000 training/webapp python app.py
        * docker run -d -p 127.0.0.1:5001:5000 training/webapp python app.py(指定IP)
- docker ps OPTIONS
    + 输出列表
        * 查看运行的容器
        * CONTAINER ID: 容器 ID
        IMAGE: 使用的镜像
        COMMAND: 启动容器时运行的命令
        CREATED: 容器的创建时间
        STATUS: 容器状态：created（已创建）、restarting（重启中）、running 或 Up（运行中）、removing（迁移中）、paused（暂停）、exited（停止）、dead（死亡）
        * PORTS: 容器的端口信息和使用的连接类型（tcp\udp）
        * NAMES: 自动分配的容器名称
    + OPTIONS值
        * -a: 查看所有容器
        * -l: 查询最后一次创建的容器
- docker logs 查看容器标准输出
- docker stop 容器名或者容器id 停止容器
- docker start 容器ID 启动容器
- docker restart重启
- docker attach(退出容器后容器停止)/exec(推荐exec,退出容器后容器不会停止) 进入容器
- docker export 导出容器
    + docker export 1e560fca3906 > ubuntu.tar
- docker import 导入容器
    + cat docker/ubuntu.tar | docker import - test/ubuntu:v1(将快照文件 ubuntu.tar 导入到镜像 test/ubuntu:v1)
    + docker import http://example.com/exampleimage.tgz example/imagerepo
- docker rm -f 容器id
- docker container prune (清理掉所有处于终止状态的容器)
- docker prot 使用 docker port 可以查看指定 （ID 或者名字）容器的某个确定端口映射到宿主机的端口号
- docker logs [ID或者名字] 可以查看容器内部的标准输出
    + docker logs -f bf08b7f2cd89
    + 参数
        * -f: 让 docker logs 像使用 tail -f 一样来输出容器内部的标准输出
- docker inspect 查看 Docker 的底层信息
- docker images 本地镜像列表
- docker pull 下载镜像
- docker search搜索镜像
- docker rmi 镜像 删除镜像
- docker push 上传镜像
    + docker tag 2e25d8496557 xxxxx.com/abc-dev/arc:1334(远程仓库)
    + docker tag ubuntu:18.04 username/ubuntu:18.04
        * docker push username/ubuntu:18.04

#### 创建镜像
- 1、从已经创建的容器中更新镜像，并且提交这个镜像
    + 提交修改镜像 docker commit -m="has update" -a="runoob" e218edb10161 runoob/ubuntu:v2
        * -m: 提交的描述信息
        * -a: 指定镜像作者
        * e218edb10161：容器 ID
        * runoob/ubuntu:v2: 指定要创建的目标镜像名
- 2、使用 Dockerfile 指令来创建一个新的镜像
    + docker build -t runoob/centos:6.7 .
        * -t ：指定要创建的目标镜像名
        * . ：Dockerfile 文件所在目录，可以指定Dockerfile 的绝对路径
-  docker tag 镜像添加一个新的标签
    +  docker tag 860c279d2fec runoob/centos:dev(dev新标签名)
    +  修改docker tag ubuntu:15.10(源) runoob/ubuntu:v3(v3新标签)（将镜像ubuntu:15.10标记为 runoob/ubuntu:v3 镜像）

#### 新建