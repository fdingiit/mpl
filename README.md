# Lab 2. XProtocol API

## 介绍
终于，我们要开始为`MosnCE`做扩展了。从本章开始，我们将会基于`XProtocol API`框架再一次实现Lab1中引入的`Simple`协议，并最终由`sidecar`代理`SDBS`流量。在开始之前，请先保证能够理解理论课1.1、1.2和1.3的内容再继续。在本章中，你可能需要精读部分mosn core代码。

## 开始
首先请确保你已经完成`lab1`并保存好你的代码，然后切换到`lab2`分支：
```shell
~ cd $GOPATH/src/github.com/fdingiit/mpl && git checkout lab2
Switched to branch 'lab2'
```
不要忘记使用`git pull`命令拉取最新代码。你应当可以在`pkg/plugin/simple`目录下找到一些新增的源文件：
```shell
~ cd $GOPATH/src/github.com/fdingiit/mpl/pkg/plugin/simple
~ tree
.
├── codec
│   ├── codec.go
│   ├── const.go
│   └── protocol.go
└── codec.go

1 directory, 4 files
```

## 你的任务
### 任务A：XProtocol API（简单）
回顾理论课的知识，MosnCE框架围绕`XProtocol API`，通过多套API对协议扩展进行支撑。首先，`XProtocol API`给出了一套接口，用来定义协议文法；其次，`XProtocolCodec API`给出另一套接口，将`XProtocol`作为其中一个部分，以编解码器的概念对外露出；最后，基于上面两套接口，我们通过声明、调用编解码器加载函数的方式，从侧面建立起「协议插件」和「sidecar主程序」之间的关系。
如果你无法快速理解上面的表述，请回顾1.3的内容并阅读相关源码后再继续。
| 任务单 |
| ------------- |
|源码文件`pkg/plugin/simple/protocol.go`中定义了一个空的数据结构`XProtocolSimple`，请使用它来实现`XProtocol API`；源码文件`pkg/plugin/simple/codec.go`中定义了另一个空的数据结构`CodecSimple`，请使用它来实现`XProtocolCodec API。`同时，请识别并实现加载编解码的方法。然后，在Lab工程下运行命令`make lab2-task-a`，查验是否能够通过所有测试。|

> *提示：
> 请注意，对于任何无法理解的接口方法，你都可以暂时不用真正实现其逻辑。*

### 任务B：协议名与协议嗅探（简单）
当边车与应用建立连接并收到流量后，首先他需要从收到的数据中判断具体的协议。在`MosnCE`架构中，即建立起「数据流量」-「协议名」-「协议编解码器」之间的关系。而每个协议编解码器都提供了协议嗅探的方法，`MosnCE`使用它们对数据流量进行协议识别。
| 任务单 |
| ------------- |
|请首先识别，然后实现`XProtocol API`和`XProtocolCodec API`（即`XProtocolSimple`和`CodecSimple`）中，与「协议名」和「协议嗅探」相关的所有方法。然后，在Lab工程下运行命令`make lab2-task-b`，查验是否能够通过所有测试。|

> *提示：
> `pkg/plugin/simple/const.go:4`定义了`simple`协议名的字符串表示，如需使用请对齐。*

### 任务C：编码与解码（中等）
你已经在`lab1`中实现了`simple`协议的编解码器。但是，这个编解码器工作在业务应用程序端，而`mosn`作为`sidecar`，同样需要有一个针对具体协议的编解码器。建议首先回忆`mosn`四层架构，以及1.1.3部分的内容后再继续。在本任务中，可能需要你花一些时间去仔细阅读开源`mosn`相关部分代码才能顺利完成。

#### XFrame
在`mosn`中，无论请求还是应答，所有的协议数据都会以「帧」的形式存在。「帧」由一套名为`XFrame`的API支撑。⚠️注意，在`mosn core`代码`pkg/stream/xprotocol/conn.go#L120`，`mosn`把从链接中读到的数据`Decode`成为了`frame`，并在`pkg/stream/xprotocol/conn.go#L142`做了数据类型校验，然后交给后续逻辑模块处理；在`pkg/stream/xprotocol/stream.go#L139`，`frame`被`Encode`成数据，最终由`mosn`转发出去。这里是协议编解码器被调用的重要节点之一，强烈建议仔细阅读理解这部分代码逻辑。

在理论课中我们曾经提到，与`InternetProtocol`一样，`mosn`采用`header-body`的形式设计协议数据结构。而`XFrameapi`的其中一个方法：`GetHeader`的返回值便是这个`header`。请特别注意，虽然在API层面并没有对`header`的数据结构做更多的约束，但是在数据流层，`mosn core`会在这里将`header`强制数据类型转换为`XFrame`，以拿到更多的信息。因此，一般来说，`GetHeader` 接口的返回值，是包含了`header`的完整`frame`，而不仅仅只是`header`。

| 任务单 |
| ------------- |
|`pkg/plugin/simple/codec/protocol.go`定义了数据结构：`RequestFrame`和`ResponseFrame`，并已经为它们设计好了部分field。请仔细阅读源码，用这两个数据结构实现`XFrameAPI`。然后，在Lab工程下运行命令`make lab2-task-c-xframe`，查验是否能够通过所有测试。|

>*提示：
> 我们将会在`lab3`中处理`XFrame`接口所继承的`Multiplexing`和`HeartbeatPredicate`接口，如果有困难，你可以暂时不用去关心它们。*

#### Codec
| 任务单 |
| ------------- |
|你已经在`lab1`中实现了`simple`协议的编解码器。请思考`XProtocol API`中关于编解码器接口与前者的关系，尤其注意`XFrame`在此处的作用，并实现相应的函数（`XProtocolSimple.Encode`、`XProtocolSimple.Decode`）。然后，在Lab工程下运行命令`make lab2-task-c-codec`，查验是否能够通过所有测试。|

>*提示：
>你应该能够复用`lab1`中的绝大部分代码。如果你在编码上花费超过30分钟，请暂时停止并先回头思考任务单中的问题。*

你的代码应该能够通过此lab的所有测试用例（测试case可能会不断增加，make命令输出以实际为准）：
```
~ make lab2-task
cd ./test && GO111MODULE=on go test -v -run Lab2
=== RUN   Test_Lab2_TaskA_Interface
    lab2_test.go:21: codec: {}
    lab2_test.go:22: xprotocol: &{}
--- PASS: Test_Lab2_TaskA_Interface (0.00s)
=== RUN   Test_Lab2_TaskA_LoadCodec
    lab2_test.go:29: codec: &{}
--- PASS: Test_Lab2_TaskA_LoadCodec (0.00s)
=== RUN   Test_Lab2_TaskB_ProtocolName
=== RUN   Test_Lab2_TaskB_ProtocolName/#00
--- PASS: Test_Lab2_TaskB_ProtocolName (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolName/#00 (0.00s)
=== RUN   Test_Lab2_TaskB_ProtocolMatch
=== RUN   Test_Lab2_TaskB_ProtocolMatch/#00
=== RUN   Test_Lab2_TaskB_ProtocolMatch/notenough
=== RUN   Test_Lab2_TaskB_ProtocolMatch/#01
=== RUN   Test_Lab2_TaskB_ProtocolMatch/#02
=== RUN   Test_Lab2_TaskB_ProtocolMatch/#03
--- PASS: Test_Lab2_TaskB_ProtocolMatch (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolMatch/#00 (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolMatch/notenough (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolMatch/#01 (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolMatch/#02 (0.00s)
    --- PASS: Test_Lab2_TaskB_ProtocolMatch/#03 (0.00s)
=== RUN   Test_Lab2_TaskC_Encode
=== RUN   Test_Lab2_TaskC_Encode/#00
=== RUN   Test_Lab2_TaskC_Encode/#01
--- PASS: Test_Lab2_TaskC_Encode (0.00s)
    --- PASS: Test_Lab2_TaskC_Encode/#00 (0.00s)
    --- PASS: Test_Lab2_TaskC_Encode/#01 (0.00s)
=== RUN   Test_Lab2_TaskC_Decode
=== RUN   Test_Lab2_TaskC_Decode/#00
=== RUN   Test_Lab2_TaskC_Decode/#01
--- PASS: Test_Lab2_TaskC_Decode (0.00s)
    --- PASS: Test_Lab2_TaskC_Decode/#00 (0.00s)
    --- PASS: Test_Lab2_TaskC_Decode/#01 (0.00s)
PASS
ok  	github.com/fdingiit/mpl/test	0.610s
```

**请将你的源码github repo link发送到指定邮箱。This completes the lab.**
