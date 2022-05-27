package talk_test

import (
	"github.com/stretchr/testify/assert"
	talk "github.com/xiaoxuan6/ding-talk"
	"testing"
)

var secret = "SEC5cf5fef599c86f7fab78ef25c0b6b1b6956a195e43170f69af9b504bbd7fc8c7"
var accessToken = "969f329a4faa0644baa590eb8c284e874d6c9ad6efdc24e381d944f56d07392d"

var robot = talk.NewRobot(accessToken)

func TestSendText(t *testing.T) {

	content := "我就是我, golang 是不一样的烟火"

	err := robot.SetSecret(secret).SendText(content, []string{}, []string{}, false)

	assert.Nil(t, err)
}

func TestSendLink(t *testing.T) {

	text := "golang 这个即将发布的新版本，创始人xx称它为红树林。而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是红树林"
	title := "时代的火车向前开"
	picUrl := ""
	messageUrl := "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"

	err := robot.SetSecret(secret).SendLink(text, title, picUrl, messageUrl)

	assert.Nil(t, err)
}

func TestSendMarkdown(t *testing.T) {

	title := "杭州天气"
	text := "#### 杭州天气  \n > 9度，@1825718XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) "
	atMobiles := []string{"1825718XXXX"}
	atUserIds := []string{}

	err := robot.SetSecret(secret).SendMarkdown(title, text, atMobiles, atUserIds, false)

	assert.Nil(t, err)
}

func TestSendActionCard(t *testing.T) {

	title := "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
	text := "![screenshot](@lADOpwk3K80C0M0FoA) \n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
	singleTitle := "阅读全文"
	singleURL := "https://www.dingtalk.com/"
	btnOrientation := "0"

	err := robot.SetSecret(secret).SendActionCard(title, text, singleTitle, singleURL, btnOrientation)

	assert.Nil(t, err)
}

func TestSendActionCard2(t *testing.T) {

	title := "我 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身"
	text := "![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png) \n\n #### 乔布斯 20 年前想打造的苹果咖啡厅 \n\n Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划"
	btnOrientation := "0" // 0：按钮竖直排列, 1：按钮横向排列

	btns := make([]talk.Btns, 0)
	btn1 := talk.Btns{
		Title:     "内容不错",
		ActionURL: "https://www.dingtalk.com/",
	}
	btns = append(btns, btn1)
	btn2 := talk.Btns{
		Title:     "不感兴趣",
		ActionURL: "https://www.dingtalk.com/",
	}
	btns = append(btns, btn2)

	err := robot.SetSecret(secret).SendActionCard2(title, text, btnOrientation, btns)

	assert.Nil(t, err)
}

func TestSendFeedCard(t *testing.T) {

	links := make([]talk.Links, 0)
	link1 := talk.Links{
		Title:      "时代的火车向前开1",
		MessageURL: "https://www.dingtalk.com/",
		PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
	}
	links = append(links, link1)
	link2 := talk.Links{
		Title:      "时代的火车向前开2",
		MessageURL: "https://www.dingtalk.com/",
		PicURL:     "https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png",
	}
	links = append(links, link2)

	err := robot.SetSecret(secret).SendFeedCard(links)

	assert.Nil(t, err)
}
