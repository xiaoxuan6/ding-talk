package talk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"net/url"
	"time"
)

const webhook = "https://oapi.dingtalk.com/robot/send?access_token="

const (
	msgTypeText       = "text"
	msgTypeLink       = "link"
	msgTypeMarkdown   = "markdown"
	msgTypeActionCard = "actionCard"
	msgTypeFeedCard   = "feedCard"
)

// Robot represents a dingtalk custom robot that can send messages to groups.
type Robot struct {
	accessToken string
	secret      string
}

// NewRobot returns a roboter that can send messages.
func NewRobot(accessToken string) *Robot {
	return &Robot{
		accessToken: accessToken,
	}
}

// SetSecret set the secret to add additional signature when send request
func (r *Robot) SetSecret(secret string) *Robot {
	r.secret = secret
	return r
}

type textMessage struct {
	Msgtype string `json:"msgtype"`
	Text    text   `json:"text"`
	At      at     `json:"at"`
}

type text struct {
	Content string `json:"content" describe:"消息内容"`
}

type at struct {
	AtMobiles []string `json:"atMobiles" describe:"被@人的手机号"`
	AtUserIds []string `json:"atUserIds" describe:"被@人的用户userid。"`
	IsAtAll   bool     `json:"isAtAll" describe:"是否@所有人。"`
}

// SendText send a text type message.
func (r *Robot) SendText(content string, atMobiles, atUserIds []string, isAtAll bool) error {
	return r.send(&textMessage{
		Msgtype: msgTypeText,
		Text: text{
			Content: content,
		},
		At: at{
			AtMobiles: atMobiles,
			AtUserIds: atUserIds,
			IsAtAll:   isAtAll,
		},
	})
}

type linkMessage struct {
	Msgtype string `json:"msgtype"`
	Link    link   `json:"link"`
}

type link struct {
	Text       string `json:"text" describe:"消息内容。如果太长只会部分展示。"`
	Title      string `json:"title" describe:"消息标题"`
	PicUrl     string `json:"picUrl" describe:"图片URL。"`
	MessageUrl string `json:"messageUrl" describe:"点击消息跳转的URL，打开方式如下：【移动端：在钉钉客户端内打开、PC端:默认侧边栏打开】"`
}

// SendLink send a link type message.
func (r *Robot) SendLink(text, title, picUrl, messageUrl string) error {
	return r.send(&linkMessage{
		Msgtype: msgTypeLink,
		Link: link{
			Text:       text,
			Title:      title,
			PicUrl:     picUrl,
			MessageUrl: messageUrl,
		},
	})
}

type markdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
	At       at       `json:"at"`
}

type markdown struct {
	Title string `json:"title" describe:"首屏会话透出的展示内容。"`
	Text  string `json:"text" describe:"markdown格式的消息。"`
}

// SendMarkdown send a markdown type message.
func (r *Robot) SendMarkdown(title, text string, atMobiles, atUserIds []string, isAtAll bool) error {
	return r.send(&markdownMessage{
		Msgtype: msgTypeMarkdown,
		Markdown: markdown{
			Text:  text,
			Title: title,
		},
		At: at{
			AtMobiles: atMobiles,
			AtUserIds: atUserIds,
			IsAtAll:   isAtAll,
		},
	})
}

type actionCardMessage struct {
	Msgtype    string     `json:"msgtype"`
	ActionCard actionCard `json:"actionCard"`
}

type actionCard struct {
	Title          string `json:"title" describe:"首屏会话透出的展示内容"`
	Text           string `json:"text" describe:"markdown格式的消息。"`
	BtnOrientation string `json:"btnOrientation" describe:"0：按钮竖直排列、1：按钮横向排列"`
	SingleTitle    string `json:"singleTitle" describe:"单个按钮的标题"`
	SingleURL      string `json:"singleURL" describe:"点击消息跳转的URL"`
}

// 整体跳转ActionCard类型
// SendActionCard send a action card type message.
func (r *Robot) SendActionCard(title, text, singleTitle, singleURL, btnOrientation string) error {
	return r.send(&actionCardMessage{
		Msgtype: msgTypeActionCard,
		ActionCard: actionCard{
			Title:          title,
			Text:           text,
			BtnOrientation: btnOrientation,
			SingleTitle:    singleTitle,
			SingleURL:      singleURL,
		},
	})
}

type actionCard2Message struct {
	Msgtype     string      `json:"msgtype"`
	ActionCard2 actionCard2 `json:"actionCard"`
}

type actionCard2 struct {
	Title          string `json:"title" describe:"首屏会话透出的展示内容。"`
	Text           string `json:"text" describe:"markdown格式的消息。"`
	BtnOrientation string `json:"btnOrientation" describe:"0：按钮竖直排列、1：按钮横向排列"`
	Btns           []Btns `json:"btns" describe:"按钮。"`
}

type Btns struct {
	Title     string `json:"title" describe:"按钮标题。"`
	ActionURL string `json:"actionURL" describe:"点击按钮触发的URL"`
}

// 独立跳转ActionCard类型
// SendActionCard2 send a action card type message.
func (r *Robot) SendActionCard2(title, text, btnOrientation string, btns []Btns) error {
	return r.send(&actionCard2Message{
		Msgtype: msgTypeActionCard,
		ActionCard2: actionCard2{
			Title:          title,
			Text:           text,
			BtnOrientation: btnOrientation,
			Btns:           btns,
		},
	})
}

type feedCardMessage struct {
	Msgtype  string   `json:"msgtype"`
	FeedCard feedCard `json:"feedCard"`
}

type feedCard struct {
	Links []Links `json:"links"`
}

type Links struct {
	Title      string `json:"title" describe:"单条信息文本。"`
	MessageURL string `json:"messageURL" describe:"点击单条信息到跳转链接。"`
	PicURL     string `json:"picURL" describe:"单条信息后面图片的URL。"`
}

// SendFeedCard send a feed card type message.
func (r *Robot) SendFeedCard(links []Links) error {
	return r.send(&feedCardMessage{
		Msgtype: msgTypeFeedCard,
		FeedCard: feedCard{
			Links: links,
		},
	})
}

func (r *Robot) send(msg interface{}) error {

	body, err := json.Marshal(msg)
	if err != nil {
		return errors.New("json 格式化数据失败")
	}

	var uri = fmt.Sprintf("%s%s", webhook, r.accessToken)
	if len(r.secret) != 0 {
		uri += genSignedByHmacSHA256(r.secret)
	}

	res, err := resty.New().R().
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetBody(string(body)).
		Post(uri)

	if err != nil {
		return err
	}

	var item = make(map[string]interface{})
	json.Unmarshal(res.Body(), &item)

	if item["errcode"] == float64(0) {
		return nil
	}

	return errors.New(item["errmsg"].(string))
}

func genSignedByHmacSHA256(secret string) string {

	timeStr := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	sign := fmt.Sprintf("%s\n%s", timeStr, secret)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(sign))
	signData := base64.StdEncoding.EncodeToString(h.Sum(nil))
	encodeURL := url.QueryEscape(signData)

	return fmt.Sprintf("&timestamp=%s&sign=%s", timeStr, encodeURL)
}
