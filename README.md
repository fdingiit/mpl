# Lab 1. Simple协议设计


## 介绍
协议是应用节点之间数据交互的基础。在通信框架流行之前，开发人员一般需要和上下游应用商讨出一份协议规范，然后开发对应的编、解码器，最终通过操作系统提供的`socket api`进行数据的收发。久而久之，定义协议及开发编解码器成为了阻碍业务快速发展的瓶颈，而一些优秀的协议设计被开发者广泛使用，并最终逐渐成为`de-facto standard`，如`HTTP`，`gRPC`等。
在微服务框架风靡互联网企业的今天，开发者一般不需要关注协议的编解码及网络层读写。但在例如银行、保险、政府机关等传统企业依旧大量存在着使用私有协议的业务应用，这些协议设计各不相同，但其目的都是一样的：为数据结构作出规范。本lab将带你设计一个简单的通信协议（名为`Simple`协议），并实现其编解码器。

## 开始
首先请确保你已经完成`lab0`并保存好你的代码，然后切换到`lab1`分支：
```shell
~ cd $GOPATH/src/github.com/fdingiit/mpl && git checkout lab1
Switched to branch 'lab1'
```

不要忘记使用`git pull`命令拉取最新代码。你应当可以在`pkg/simple`和`pkg/sdbs`目录下找到一些源文件：
```shell
~ cd $GOPATH/src/github.com/fdingiit/mpl/pkg/simple
~ tree
.
└── protocol.go
~ cd $GOPATH/src/github.com/fdingiit/mpl/pkg/sdbs
~ tree
.
├── gateway
│   ├── start_gateway.sh
│   └── stop_gateway.sh
└── server
    ├── main.go
    └── pkg
        └── server.go
```

## 你的任务
下面三个表格给出了`Simple`协议报文的完整定义，并给出了请求报文和应答报文的例子。请仔细阅读理解后再继续。

### 公共报文头 
| 字段名称 | 类型	| 长度（字节） | 必填 |	备注  |
| ------------- | ------------- | ------------- | ------------- | ------------- |
| 报文总长度  | int  | 8 | 是  | 不满8位左补0  |
| 报文类型	| char	| 2	| 是	| 请求：RQ；应答：RS |
| 公共报文头 | 	报文总长度 | 	int	| 8	| 是	|
| 报文类型	| char| 	2	| 是| 	请求：RQ；应答：RS| 
| 翻页标志	| int| 	1| 	是| 	首页：0；翻页：1| 
| 校验码	| char | 	32	| 是	| 自由选择校验码生成算法；不满32位右补空格| 
| 服务码	| int| 	8| 	是| 	不满8位左补0| 
| 保留字段	| int| 	1| 	是| 	默认为0| 

### 请求
| 字段名称	| 类型	| Tag	| 必填	| 备注|
| ------------- | ------------- | ------------- | ------------- | ------------- |
| 公共报文头	| -	| -	| 是	| | 
| Unix时间戳	| int	| timestamp	| 是	| 不含根tag的非标准xml格式，utf-8编码；报文长度不超过4kb| 
| 流水号	| int	| serial_no | 	是| 同上| 
| 币种	| int	| currency	| 是|  同上| 
| 转账金额	| int	| amount	| 是|  同上| 
| 金额单位	| int	| unit	| 是| 同上| 
| 转出账户id	| int| 	out_account_id	| 是|  同上| 
| 转出银行id	| int	| out_bank_id	| 是|  同上| 
| 转入账户id	| int	| in_account_id| 	是|  同上| 
| 转入银行id	| int	| in_bank_id| 	是|  同上| 
| 备注	| string| 	notes| 	否| 同上| 

### 应答

| 字段名称	| 类型	| Tag	| 必填	| 备注|
| ------------- | ------------- | ------------- | ------------- | ------------- |
| 公共报文头	| -	| -	| 是	| |  
| Unix时间戳	| int	| timestamp	| 是	| 不含根tag的非标准xml格式，utf-8编码；报文长度不超过4kb| 
| 流水号	| int	| serial_no| 	是|  同上| 
| 错误码	| int	| err_code	| 是|  同上| 
| 信息	| string	| message| 	否|  同上| 

```
# Simple协议请求示例
00000328RQ0tPK6UhVeIHb2hrsedxXMJHw         010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><currency>2</currency><amount>100</amount><unit>0</unit><out_bank_id>2</out_bank_id><out_account_id>1234567899321</out_account_id><in_bank_id>2</in_bank_id><in_account_id>3211541298661</in_account_id><notes></notes>
```
```
# Simple协议请求示例
00000156RS0tPK6UhVeIHb2hrsedxXMJHw         010005010<timestamp>1648811583</timestamp><serial_no>12345</serial_no><err_code>0</err_code><message>ok</message>
```

## 任务A：公共报文头（简单）
| 任务单 |
| ------------- |
| 源码文件`pkg/simple/protocol.go`给出了公共报文头的数据结构定义，并给出了一种编码器实现。找到他们，并补充实现其解码器方法。然后，在Lab工程下运行命令`make lab1-task-a`，查验是否能够通过所有测试。|

> *提示：
> 表1里定义了报文总长度，请思考这个数值分别在什么阶段（编码/解码；收数据/写数据）由谁（编解码器/业务逻辑）来对其做什么操作（读/写）。另外，如果出现了实际数据长度与这个值不一致的情况，该如何处理。*

## 任务B：请求体和应答体（简单）
| 任务单 |
| ------------- |
| 源码文件`pkg/simple/protocol.go`给出了请求和应答的数据结构定义，并给出了一种应答的解码器实现。找到他们，并补充实现其他缺失的编解码器方法。特别的，在本lab中，我们统一使用`MD5 Hash`算法对报文体进行校验码生成。然后，在Lab工程下运行命令`make lab1-task-b`，查验是否能够通过所有测试。|

> *提示：
> 如果请求体/应答体里的字段是随机乱序的话，你的代码还能够正常work吗？如果业务方需要对字段进行增改，你的实现能做到无需修改吗？*

## 任务C：SDBS系统（简单）

你已经对如何开发协议编解码器有了初步的认识，但这还不够。现在，我们将基于上述协议建设一个简单的银行数字化信息系统：`Simple Dummy Banking System（or SDBS）`。`SDBS`的代码在目录`pkg/sdbs`下。`pkg/sdbs/server/server.go`调用了你所实现的编解码器方法，并作为一个`SDBS`中的一个应用，在`9999`端口对外提供服务。
| 任务单 |
| ------------- |
| 请复用任务A/B的代码实现一个`SDBS API Gateway`服务。这个服务在`:80/transfer`端口接收请求，并根据其内容向某个`SDBS`服务实例发起调用。`SDBS API Gateway`具体的接口定义见下文`SDBS API Gateway OpenAPI`文档。请把`gateway`的源码放置在`pkg/sdbs/gateway`目录下，并分别补全脚本`pkg/sdbs/gateway/start_gateway.sh`和`pkg/sdbs/gateway/stop_gateway.sh`中的内容，以启停你的`SDBS API Gateway`。然后，在Lab工程下运行命令`make lab1-task-c`，查验是否能够通过所有测试。|

```yaml
openapi: 3.0.0
servers:
  # Added by API Auto Mocking Plugin
  - description: API Gateway for SDBS
    url: https://virtserver.swaggerhub.com/fdingiit/SDBS/1.0.0
info:
  description: This is a gateway for Simple Dummy Banking System
  version: "1.0.0"
  title: API Gateway for SDBS
  contact:
    email: fdingiit@gmail.com
  license:
    name: Apache 2.0
    url: 'http://www.apache.org/licenses/LICENSE-2.0.html'
tags:
  - name: services
    description: API Gateway for Simple Dummy Banking System
paths:
  /transfer:
    post:
      tags:
        - services
      summary: transfer system
      operationId: addInventory
      description: Transfter from one account to another
      parameters: 
        - in: header
          name: X-SDBS-PAGING-MASK
          schema:
            type: integer
            example: 0
        - in: header
          name: X-SDBS-CHECKSUM
          schema:
            type: string
            example: '665db818fa5ef08e9f10ec77d76b9a0e'
      responses:
        '200':
          description: transfer done
          headers:
            X-SDBS-PAGING-MASK:
              schema:
                type: integer
                example: 0
            X-SDBS-CHECKSUM:
              schema:
                type: string
                example: '665db818fa5ef08e9f10ec77d76b9a0e'
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/Response'
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Request'
        description: Inventory item to add
components:
  schemas:
    Request:
      type: object
      required:
        - timestamp
        - serial_no
        - currency
        - amount
        - unit
        - out_bank_id
        - out_account_id
        - in_bank_id
        - in_account_id
      properties:
        timestamp:
          type: integer
          format: int64
          example: 1648811583
        serial_no:
          type: integer
          format: int64
          example: 12345
        currency:
          type: integer
          example: 2
        amount:
          type: integer
          format: int64
          example: 100
        unit:
          type: integer
          example: 0
        out_bank_id:
          type: integer
          format: int64
          example: 2
        out_account_id:
          type: integer
          format: int64
          example: 1234567899321
        in_bank_id:
          type: integer
          format: int64
          example: 2
        in_account_id:
          type: integer
          format: int64
          example: 3211541298661
        notes:
          type: string
          example: 'borrowed to liyundi'
    Response:
      type: object
      required:
        - timestamp
        - serial_no
        - err_code
      properties:
        timestamp:
          type: integer
          format: int64
          example: 1648811583
        serial_no:
          type: integer
          format: int64
          example: 12345
        err_code:
          type: integer
          example: 0
        message:
          type: string
          example: 'ok'
```

> *提示：
> 你可以在`SwaggerHub`中查阅和mock这个api，以获得更直观的理解。*

| 挑战🌟 |
| ------------- |
| 语言/技术栈无关是服务网格作为应用负载节点网络代理的一个重要技术优势。请尝试使用任何你所喜爱的非`Go`语言编写`gateway`服务，并完成任务C。请注意，对于一些有编译、运行环境要求的编程语言，推荐使用容器的形式启动。请不要在`start_gateway.sh`脚本里下载、安装依赖。|

> *提示：
> 你可以使用`Swagger`的`CodeGen`工具，或直接在`SwaggerHub`中快速生成多种编程语言的HTTP Server桩代码。*

| 挑战🌟🌟🌟 |
| ------------- |
| `Simple`协议的数据格式设计并非是一个好的工程实践。请自行搜索阅读相关知识内容或调研业界优秀案例，尽可能多地识别`Simple`协议设计存在的问题，并尝试在兼容现有能力的前提下给出优化后的协议设计及实现。除了源码之外，你还需要提交一份有支撑的（如资料引用、实验数据等）技术报告。|

你的代码应该能够通过此lab的所有测试用例（测试case可能会不断增加，make命令输出以实际为准）：
```
~ make lab1-task
cd ./test && GO111MODULE=on go test -v -run Lab1
=== RUN   Test_Lab1_TaskA
=== RUN   Test_Lab1_TaskA/#00
--- PASS: Test_Lab1_TaskA (0.00s)
    --- PASS: Test_Lab1_TaskA/#00 (0.00s)
=== RUN   Test_Lab1_TaskB_Request
=== RUN   Test_Lab1_TaskB_Request/#00
--- PASS: Test_Lab1_TaskB_Request (0.00s)
    --- PASS: Test_Lab1_TaskB_Request/#00 (0.00s)
=== RUN   Test_Lab1_TaskB_Response
=== RUN   Test_Lab1_TaskB_Response/#00
--- PASS: Test_Lab1_TaskB_Response (0.00s)
    --- PASS: Test_Lab1_TaskB_Response/#00 (0.00s)
=== RUN   Test_Lab1_TaskC
checking port: 9999
checking port: 80
[SDBS] Error reading: EOF
=== RUN   Test_Lab1_TaskC/#00
[SDBS] Rsp:  00000170RS0c32dafd4a53d0a4d04f2b15f9305dd5e001005010<timestamp>1649664561</timestamp><serial_no>5630350334869219902</serial_no><err_code>0</err_code><message>ok</message>
--- PASS: Test_Lab1_TaskC (4.15s)
    --- PASS: Test_Lab1_TaskC/#00 (0.00s)
PASS
ok  	github.com/fdingiit/mpl/test	4.975s
```

**请将你的源码github repo link发送到指定邮箱。This completes the lab.**
