---
title: go-go-micro机器人bot
categories: Go
tags: [go, 微服务, bot, 机器人, slack]
date: 2019-10-10 19:30:53
comments: false
---

> go-go-micro机器人bot

<!-- more -->

---

###  前篇

- micro bot - https://www.bookstack.cn/read/go-micro/72414279e3575e70.md
- https://medium.com/microhq/the-micro-bot-chatops-for-microservices-546ecc1a9ed8



**micro bot** 是一只藏在微服务中的小马蝇，有了它，我们可以在 Slack、HipChat、XMPP 等等聊天程序中与它对话，通过它来操控服务。

我们把消息发送给它，它基于这些消息模仿执行CLI，触发指定的接口功能。

![](http://yxbl.itengshe.com/20191010193315-1.png)



是用很简单. Slack 为例

1. 获得 Slack 的 token. 如: `xoxb-123-123-asdasd`

2. 启动 bot

    ```json
    $ micro bot --inputs=slack --slack_token=xoxb-123-123-asdasd
    2019/10/10 19:25:06 [bot] starting
    2019/10/10 19:25:06 [bot] starting input slack
    2019/10/10 19:25:07 [bot][loop] starting slack
    2019/10/10 19:25:07 [bot][loop] connecting to slack
    2019/10/10 19:25:07 Transport [http] Listening on [::]:64084
    2019/10/10 19:25:07 Broker [http] Connected to [::]:64085
    2019/10/10 19:25:07 Registry [mdns] Registering node: go.micro.bot-538d4b38-d9a4-4acd-8210-6a46bf400f6e
    ```

3. 在 Slack 就可以输入命令与机器人交互

    ![](http://yxbl.itengshe.com/20191010193742-1.png)



---

### Slack

- 使用Go开发一个 Slack 运维机器人 - https://colobu.com/2015/11/04/create-a-slack-bot-with-golang/



#### 创建流程

1. 你首先创建一个Team，并且加入到这个Team中。这是使用Slack的第一步。以后你可以直接访问 http://.slack.com 登录到你的team中。

2. 新建一个 [bot user integration](https://my.slack.com/services/new/bot)。你需要为你的机器人起一个名字. 这里我起了 *xiaoyang*.

    可以为它指定头像，slack会为它生成一个 API Token。 这个API Token很重要， 以后访问slack API需要传入这个token。如: `xoxb-123-123-asdasd`

    你也可以为你的普通登录用户生成full-access token，网址是: https://api.slack.com/web。

3. 将你创建的 bot 加入到一个 Apps 中.

    ![](http://yxbl.itengshe.com/20191010194354-1.png)

4. 在 Slack 就可以输入命令与机器人交互

    <img src="http://yxbl.itengshe.com/20191010193742-1.png" style="zoom:50%;" />



---

### 增加命令

#### 方式一: 启动一个命令服务 (推荐) 

- Commands as Services - https://micro.mu/docs/bot.html#commands-as-services

参考官方 examples

1. 启动服务

    ```json
    f:\a_link_workspace\go\GoWinEnv_MicroExamples\src\github.com\micro\examples\command (master -> origin)
    $ go run main.go
    ```

2. 刷新一下 slack, 输入 *help* 指令可以看到新加的指令 *command*.

    ![](http://yxbl.itengshe.com/20191011112647-1.png)



##### 添加多个命令服务

官方的说明一个 命令服务职能执行一个命令. 多个命令的话需要多个 命令服务.

- How does it work? - https://micro.mu/docs/bot.html#how-does-it-work

在 `go.micro.bot` 这个命名空间下的都可以添加到可执行的 命令列表.

实践测试: *test_cmd_service*

1. 添加两个命令

    ```go
    // cmd1
    	rsp.Usage = "wilker"
    	rsp.Description = "This is an example bot command as a micro service wilker"
    
    		micro.Name("go.micro.bot.wilker"), // rsp.Usage 与 go.micro.bot. 后缀一致
    
    // cmd2
    	rsp.Usage = "yun"
    	rsp.Description = "This is an example bot command as a micro service yun"
    		micro.Name("go.micro.bot.yun"),
    ```

2. 分别启动两个命令

    ```json
    f:\a_link_workspace\go\GoWinEnv_new\src\GoMicro\test_cmd\test_cmd_service\cmd1 (master -> origin)
    $ go run main.go
    2019/10/11 11:37:37 Registry [mdns] Registering node: go.micro.bot.wilker-fd7da5e9-5339-402c-9dc0-b7cecb2d93e2
    
    f:\a_link_workspace\go\GoWinEnv_new\src\GoMicro\test_cmd\test_cmd_service\cmd2 (master -> origin)
    $ go run main.go
    2019/10/11 11:37:59 Registry [mdns] Registering node: go.micro.bot.yun-6a8f9cee-5815-48aa-939e-6763022e97f6
    ```

3. 刷新一下 slack, 输入 *help* 指令可以看到新加的指令 *command*.

    ![](http://yxbl.itengshe.com/20191011114704-1.png)



---

#### 方式二: 重新编译 micro

参考: *test_cmd_compile*

- 增加命令 - https://www.bookstack.cn/read/go-micro/72414279e3575e70.md#%E5%A2%9E%E5%8A%A0%E5%91%BD%E4%BB%A4



使用流程

1. cd 到 *src/github.com/micro/micro* 目录下, 新增一个自定义命令文件 *my-cmd.go*

    增加一个 *yang* 的指令

    ```go
    package main
    
    import (
    	"github.com/micro/go-micro/agent/command"
    )
    
    func Ping() command.Command {
    	usage := "yang"
    	description := "hello wilker!!"
    	return command.NewCommand("ping", usage, description, func(args ...string) ([]byte, error) {
    		return []byte("Returns xuan 666"), nil
    	})
    }
    
    func init() {
    	command.Commands["^yang$"] = Ping()
    }
    ```

2. 加入 *my-cmd.go* 重新编译 micro 生成可执行文件 *myMicro.exe*

    ```json
    F:\a_link_workspace\go\GoWinEnv_Test01
    $ cd src\github.com\micro\micro
    
    F:\a_link_workspace\go\GoWinEnv_Test01\src\github.com\micro\micro (master -> origin)
    $ go build -o myMicro.exe main.go my-cmd.go
    ```

3. 启动机器人

    ```json
    $ myMicro.exe bot --inputs=slack --slack_token=xoxb-123-123-asdasd
    2019/10/10 20:17:56 [bot] starting
    2019/10/10 20:17:56 [bot] starting input slack
    2019/10/10 20:17:57 [bot][loop] starting slack
    2019/10/10 20:17:57 [bot][loop] connecting to slack
    ```

4. 访问 https://app.slack.com/client, 输入 *help* 指令与 *yang* 指令可以看到增加的命令

    ![](http://yxbl.itengshe.com/20191010202819-1.png)



---

### 自定义机器人交互插件

- 增加新的输入源 - https://www.bookstack.cn/read/go-micro/72414279e3575e70.md#%E5%A2%9E%E5%8A%A0%E6%96%B0%E7%9A%84%E8%BE%93%E5%85%A5%E6%BA%90