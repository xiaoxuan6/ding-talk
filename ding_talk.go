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

	val, ok := limit.Load(LimitName)
	if ok {
		LimitRate := val.(*LimitRate)
		if ok = LimitRate.Allow(); ok {
			err := send(r, msg)
			if err != nil {
				LimitRate.Reduce()
			}
			return err
		}
		return errors.New("请求超出最大限制次数")
	}

	return send(r, msg)
}

func send(r *Robot, msg interface{}) error {

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
