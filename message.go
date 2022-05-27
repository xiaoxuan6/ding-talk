package talk

// textMessage type text
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

// linkMessage type link
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

// markdownMessage type mark
type markdownMessage struct {
	Msgtype  string   `json:"msgtype"`
	Markdown markdown `json:"markdown"`
	At       at       `json:"at"`
}

type markdown struct {
	Title string `json:"title" describe:"首屏会话透出的展示内容。"`
	Text  string `json:"text" describe:"markdown格式的消息。"`
}

// actionCardMessage type actionCard
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

// actionCard2Message type actionCard
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

// feedCardMessage type feedCard
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
