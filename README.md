# Lab 0. 准备与预热


## 介绍
本lab将带你上手体验社区版Mosn（Mosn Community Edition，即MosnCE），以对其有更直观、更hands-on的认识。在此过程中，你将首先使用Mosn对「**标准HTTP协议**」请求进行代理；然后，你将进一步使用「**插件动态加载机制**」，通过插件的形式对「**第三方非标准协议**」请求进行代理。

## 开始
在开始之前，请先安装`Go`语言开发环境，以及`Git`。其中，**本lab对`Go`语言最低的版本要求为`1.14`，最高为`1.17.3`。**

环境准备好后，首先获取MosnCE源码：
```shell
~ mkdir -p $GOPATH/src/mosn.io/ && cd $GOPATH/src/mosn.io/
~ git clone https://github.com/mosn/mosn.git
```

在本Lab的所有阶段，我们将使用同一个指定的MosnCE版本：`v0.26.0`：
```shell
~ mkdir -p $GOPATH/src/mosn.io/ && cd $GOPATH/src/mosn.io/
~ git clone https://github.com/mosn/mosn.git
```

然后，请获取本Lab的工程代码：
```shell
~ go get -u github.com/fdingiit/mpl

```
并切换到`lab0`分支：
```shell
~ cd $GOPATH/src/github.com/fdingiit/mpl && git checkout lab0
Switched to branch 'lab0'
```

## 你的任务
接下来，请简要地浏览MosnCE的整体源码结构，找到并理解工程的`Makefile`文件和`main`函数；同时，请找出其所依赖的`mosn.io/pkg`和`mosn.io/api`版本，并简要浏览其源码。
> *提示：
> `main`函数源码文件的位置在`cmd/mosn`目录下。*

### 任务A：编译（简单）
| 任务单 |
| ------------- |
| 使用MosnCE源码提供的`Makefile`文件编译得到二进制文件，将其放置于Lab工程的固定目录`build/mosn`下，并确保二进制的文件名为`mosnd`。然后，在Lab工程下运行命令`make lab0-task-a`，查验是否能够通过所有测试。 |

> *提示：
> 为了简化测试，我们在本Lab的所有阶段都将使用固定的目录来放置产出的Mosn二进制、协议插件、配置文件等。如有需要，建议使用软链接的方式对他们进行更新、管理。除非特别声明，我们默认将把`$GOPATH/src/github.com/fdingiit/mpl`作为当前工作目录。如`build/mosn`是指`$GOPATH/src/github.com/fdingiit/mpl/build/mosn`*

### 任务B：代理HTTP请求（简单）
在编译之后，你可以在`$GOPATH/src/mosn.io/mosn/build/bundles/v0.26.0/binary`目录下找到标准示例配置文件`mosn_config.json`。请确保你能够理解此配置文件每一行的内容，并回答下列问题：  
- 使用这份配置启动的mosn的角色是什么（服务端sidecar、客户端sidecar、其他）？
- 这份配置所配置的协议是什么？如果实际接收的数据并不是配置的协议，会发生什么？
- 包括客户端、服务端应用、mosn在内的整个数据链路是怎样的（精确到端口）？

如果无法回答上述问题或不确定答案是否正确，请仔细复习理论课中的相关章节，结合MosnCE源码并动手实验后再继续后面的内容。
| 任务单 |
| ------------- |
| 给出一份完整的mosn配置，作为`server`端代理`HTTP1`请求，并监听在`12046`端口。其所代理的应用于`localhost:1080/`提供服务。配置请放置于Lab工程的固定目录`build/mosn`下，并命名为`mosn_config_lab0_taskb_server.json`。给出另一份完整的mosn配置，作为`client`端代理`HTTP1`请求，并监听在`12045`端口。对于目标path为`/`的请求，将其代理到本地`server`端的sidecar。配置请放置于Lab工程的固定目录`build/mosn`下，并命名为`mosn_config_lab0_taskb_client.json`。然后，在Lab工程下运行命令`make lab0-task-b`，查验是否能够通过所有测试。 |


### 任务C：使用插件代理非标协议请求（中等）
在本lab的最后，我们将使用Go-native-plugin作为编解码器，对非标准协议请求进行代理。在`pkg/protocol/demo`目录下，你可以找到一份完整的非标协议编解码器示例源码、一份不完整的编译脚本、以及两份不完整的mosn配置文件。特别的，请确保自己能够充分掌握理论课中关于Go-plugin的内容，然后再继续。
| 任务单 |
| ------------- |
| 请补充实现编译协议插件的脚本：`pkg/protocol/demo/make_codec.sh`中的函数`make_so`，测试程序将在脚本所在目录调用脚本生成插件。同时，结合编译脚本，将两个mosn配置文件：`pkg/protocol/demo/client_config.json`和`pkg/protocol/demo/server_config.json`的内容补充完整（注：不要修改已有内容），并放置于`build/mosn`目录下，并分别命名为`mosn_config_lab0_taskc_client.json`和`mosn_config_lab0_taskc_server.json`。然后，在Lab工程下运行命令`make lab0-task-c`，查验是否能够通过所有测试。|

在使用`mosnd start -c`命令启动mosn时看到类似如下的标准输出内容，则可以说明插件被成功加载：
```
2022-03-30 13:38:11,455 [INFO] [mosn] [init codec] loading protocol [codec] from third part codec
2022-03-30 13:38:11,455 [INFO] [network] [ register pool factory] register protocol: demo factory
2022-03-30 13:38:11,455 [INFO] [mosn] [init codec] load go plugin codec succeed: /Users/dingfei/Go/src/github.com/fdingiit/mpl/build/mosn/codec.so
```

> *提示：
> 为了能够正确加载插件，可能需要修改MosnCE源码的Makefile。*

你的代码应该能够通过此lab的所有测试用例（测试case可能会不断增加，make命令输出以实际为准）：
```
~ make lab0-task
make lab0-make-plugin
cd ./pkg/protocol/demo && bash make_codec.sh
cd ./test && GO111MODULE=on go test -v -run Task
=== RUN   Test_TaskA
lab0_test.go:48: [pass] correct mosn executable
lab0_test.go:55: [pass] correct mosn version
--- PASS: Test_TaskA (0.17s)
=== RUN   Test_TaskB
lab0_test.go:65: [pass] mosn server config exists
lab0_test.go:73: [pass] mosn client config exists
lab0_test.go:85: [pass] mosn server stared
lab0_test.go:94: [pass] mosn client stared
lab0_test.go:123: [pass] correct response
--- PASS: Test_TaskB (2.03s)
=== RUN   Test_TaskC
lab0_test.go:133: [pass] mosn server config exists
lab0_test.go:141: [pass] mosn client config exists
lab0_test.go:153: [pass] mosn server stared
lab0_test.go:162: [pass] mosn client stared
Hello World
Hello, I am server
lab0_test.go:214: [pass] correct response
--- PASS: Test_TaskC (1.01s)
PASS
ok  	github.com/fdingiit/mpl/test	4.005s
```

**请将你的源码github repo link发送到指定邮箱。This completes the lab.**

