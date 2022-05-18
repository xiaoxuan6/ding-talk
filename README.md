![Dingtalk_20220518174235](https://user-images.githubusercontent.com/32762220/169010157-188f2101-b747-4c6b-a333-00b718777e4a.jpg)

# DingTalk

钉钉自定义机器人 Go API.

### 限流限制
由于消息发送太频繁会严重影响群成员的使用体验，因此钉钉开放平台对自定义机器人发送消息的频率作出以下限制：

每个机器人每分钟最多发送20条消息到群里，如果超过20条，会限流10分钟。

# 支持的消息类型：

* text 类型
* link 类型
* markdown 类型
* actionCard 类型
* feedCard 类型

# Installation

    go get github.com/xiaoxuan6/ding-talk

# Usage

<details>
<summary><b>text 类型</b></summary>

```
    robot := talk.NewRobot(accessToken)

    content := "我就是我, @XXX 是不一样的烟火"
    atMobiles := []string{}
    atUserIds := []string{}
    isAtAll := false
    
    err := robot.SendText(content, atMobiles, atUserIds, isAtAll)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

<details>
<summary><b>link 类型</b></summary>

```
    robot := talk.NewRobot(accessToken)

    text := "这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林"
    title := "时代的火车向前开"
    picUrl := ""
    messageUrl := "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
    
    err := robot.SendLink(text, title, picUrl, messageUrl)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

<details>
<summary><b>markdown 类型</b></summary>

```
    robot := talk.NewRobot(accessToken)

    title := "杭州天气"
    text := "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) "
    atMobiles := []string{"1825718XXXX"}
    atUserIds := []string{}
    isAtAll := false
	
    err := robot.SendMarkdown(title, text, atMobiles, atUserIds, isAtAll)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

<details>
<summary><b>整体跳转ActionCard类型</b></summary>

```
    robot := talk.NewRobot(accessToken)

    title := "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
    text := "![screenshot](@lADOpwk3K80C0M0FoA) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
    singleTitle := "阅读全文"
    singleURL := "https://www.dingtalk.com/"
    btnOrientation := "0" // 0：按钮竖直排列, 1：按钮横向排列
	
    err := robot.SendActionCard(title, text, singleTitle, singleURL, btnOrientation)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

<details>
<summary><b>独立跳转ActionCard类型</b></summary>

```
    robot := talk.NewRobot(accessToken)

    title := "我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
    text := "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
    btnOrientation := "0" // 0：按钮竖直排列, 1：按钮横向排列

    btns := make([]talk.Btns, 0)
    btn1 := talk.Btns{
        Title: "内容不错",		
        ActionURL: "https://www.dingtalk.com/",
    }
    btns = append(btns, btn1)
    btn2 := talk.Btns{
        Title: "不感兴趣",		
        ActionURL: "https://www.dingtalk.com/",
    }
    btns = append(btns, btn2)
	
    err := robot.SendActionCard2(title, text, btnOrientation, btns)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

<details>
<summary><b>FeedCard类型</b></summary>

```
    robot := talk.NewRobot(accessToken)
    
    links := make([]talk.Links, 0)
    link1 := talk.Links{
    	Title: "时代的火车向前开1",
    	MessageURL: "https://www.dingtalk.com/",
    	PicURL: "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
    }
    links = append(links, link1)
    link2 := talk.Links{
    	Title: "时代的火车向前开2",
    	MessageURL: "https://www.dingtalk.com/",
    	PicURL: "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
    }
    links = append(links, link2)
	
    err := robot.SendFeedCard(links)
    if err != nil {
        log.Fatal(err)
    }
```

</details>

# 安全设置为加密模式

    robot := talk.NewRobot(accessToken)

    content := "我就是我, @XXX 是不一样的烟火"
    atMobiles := []string{}
    atUserIds := []string{}
    isAtAll := false

    robot.SetSecret("xxx").SendText(content, atMobiles, atUserIds, isAtAll)
