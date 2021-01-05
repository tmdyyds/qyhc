#### 常用命令
- docker network create -d bridge test-net 新建网络 多个容器互相连接
    + docker run -itd --name test1 --network test-net ubuntu /bin/bash
    + docker run -itd --name test2 --network test-net ubuntu /bin/bash
    + d：参数指定 Docker 网络类型，有 bridge、overlay
- docker run ubuntu:15.10 /bin/echo "Hello world"
    + docker: Docker 的二进制执行文件
    + run: 与前面的 docker 组合来运行一个容器
    + ubuntu:15.10 指定要运行的镜像，Docker 首先从本地主机上查找镜像是否存在，如果不存在，Docker 就会从镜像仓库 Docker Hub 下载公共镜像
    + /bin/echo "Hello world": 在启动的容器里执行的命令
    + **Docker 以 ubuntu15.10 镜像创建一个新容器，然后在容器里执行 bin/echo "Hello world"**
- docker run -i -t (-d) ubuntu:15.10 /bin/bash(进行交互)
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
    + 构建dockerfile命令
        * RUN 
            - shell格式 RUN <命令行命令(等同于，在终端操作的 shell 命令)>
            - exec格式 RUN ["可执行文件", "参数1", "参数2"]
                + 例如 RUN ["./test.php", "dev", "offline"] 等价于 RUN ./test.php dev offline
-  docker tag 镜像添加一个新的标签
    +  docker tag 860c279d2fec runoob/centos:dev(dev新标签名)
    +  修改docker tag ubuntu:15.10(源) runoob/ubuntu:v3(v3新标签)（将镜像ubuntu:15.10标记为 runoob/ubuntu:v3 镜像）

#### 指令解析
- COPY
    + COPY [--chown=<user>:<group>] <源路径1>...  <目标路径>
    + COPY [--chown=<user>:<group>] ["<源路径1>",...  "<目标路径>"]
        * [--chown=<user>:<group>]：可选参数，用户改变复制到容器内文件的拥有者和属组
        * <源路径>：源文件或者源目录，这里可以是通配符表达式，其通配符规则要满足 Go 的 filepath.Match 规则
            - 例如 COPY hom* /mydir/  |  COPY hom?.txt /mydir/
        * <目标路径>：容器内的指定路径，该路径不用事先建好，路径不存在的话，会自动创建
- ADD
    + ADD 指令和 COPY 的使用格式一致（同样需求下，官方推荐使用 COPY）。功能也类似，不同之处如下
        * ADD 的优点：在执行 <源文件> 为 tar 压缩文件的话，压缩格式为 gzip, bzip2 以及 xz 的情况下，会自动复制并解压到 <目标路径>
        * ADD 的缺点：在不解压的前提下，无法复制 tar 压缩文件。会令镜像构建缓存失效，从而可能会令镜像构建变得比较缓慢。具体是否使用，可以根据是否需要自动解压来决定
- CMD(为启动的容器指定默认要运行的程序，程序运行结束，容器也就结束。CMD 指令指定的程序可被 docker run 命令行参数中指定要运行的程序所覆盖。如果 Dockerfile 中如果存在多个 CMD 指令，仅最后一个生效)
    + 格式
        * CMD <shell 命令> 
        * CMD ["<可执行文件或命令>","<param1>","<param2>",...] (推荐使用)
        * CMD ["<param1>","<param2>",...]  # 该写法是为 ENTRYPOINT 指令指定的程序提供默认参数
    + 类似于 RUN 指令，用于运行程序，但二者运行的时间点不同
        * CMD 在docker run 时运行
        * RUN 是在 docker build
- ENTRYPOINT(entrypoint， 类似于 CMD 指令，但其不会被 docker run 的命令行参数指定的指令所覆盖，而且这些命令行参数会被当作参数送给 ENTRYPOINT 指令指定的程序。但是, 如果运行 docker run 时使用了 --entrypoint 选项，此选项的参数可当作要运行的程序覆盖 ENTRYPOINT 指令指定的程序)
    + 格式
        * NTRYPOINT ["<executeable>","<param1>","<param2>",...]
- ENV(置环境变量，定义了环境变量，那么在后续的指令中，就可以使用这个环境变量)
    + 格式
        * ENV <key> <value>
        * ENV <key1>=<value1> <key2>=<value2>...
- ARG(构建参数，与 ENV 作用一至。不过作用域不一样。ARG 设置的环境变量仅对 Dockerfile 内有效，也就是说只有 docker build 的过程中有效，构建好的镜像内不存在此环境变量。构建命令 docker build 中可以用 --build-arg <参数名>=<值> 来覆盖)
    + 格式
        * ARG <参数名>[=<默认值>]
- VOLUME(volume，定义匿名数据卷。在启动容器时忘记挂载数据卷，会自动挂载到匿名卷)
    + 格式
        * VOLUME ["<路径1>", "<路径2>"...]
        * VOLUME <路径>
    + 作用
        * 避免重要的数据，因容器重启而丢失，这是非常致命的
        * 避免容器不断变大
    + 在启动容器 docker run 的时候，我们可以通过 -v 参数修改挂载点
- EXPOSE(仅仅只是声明端口)
    + 格式
        * EXPOSE <端口1> [<端口2>...]
    + 作用
        * 帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射
        * 运行时使用随机端口映射时，也就是 docker run -P 时，会自动随机映射 EXPOSE 的端口
- WORKDIR(指定工作目录.用 WORKDIR 指定的工作目录，会在构建镜像的每一层中都存在。（WORKDIR 指定的工作目录，必须是提前创建好的）。docker build 构建镜像过程中的，每一个 RUN 命令都是新建的一层。只有通过 WORKDIR 创建的目录才会一直存在)
    + 格式
        * WORKDIR <工作目录路径>
- USER(用于指定执行后续命令的用户和用户组，这边只是切换后续命令执行的用户（用户和用户组必须提前已经存在）)
    + USER <用户名>[:<用户组>]
- HEALTHCHECK(healthcheck, 用于指定某个程序或者指令来监控 docker 容器服务的运行状态)
    + 格式
        * HEALTHCHECK [选项] CMD <命令>：设置检查容器健康状况的命令
        * HEALTHCHECK NONE：如果基础镜像有健康检查指令，使用这行可以屏蔽掉其健康检查指令
        * HEALTHCHECK [选项] CMD <命令> : 这边 CMD 后面跟随的命令使用，可以参考 CMD 的用法
- ONBUILD(onbutld, 用于延迟构建命令的执行。简单的说，就是 Dockerfile 里用 ONBUILD 指定的命令，在本次构建镜像的过程中不会执行（假设镜像为 test-build）。当有新的 Dockerfile 使用了之前构建的镜像 FROM test-build ，这是执行新镜像的 Dockerfile 构建时候，会执行 test-build 的 Dockerfile 里的 ONBUILD 指定的命令)
    + 格式
        * ONBUILD <其它指令>
