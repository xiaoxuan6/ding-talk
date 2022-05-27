package talk_test

import (
	"github.com/stretchr/testify/assert"
	talk "github.com/xiaoxuan6/ding-talk"
	"testing"
)

func init() {
	talk.Init()
}

func TestText(t *testing.T) {
	for i := 0; i < 5; i++ {
		r := talk.NewRobot("xxx")

		err := r.SetSecret("xxx").SendText("test limit send", []string{}, []string{}, false)
		if err != nil {
			assert.Errorf(t, err, err.Error())
		} else {
			assert.Nil(t, err)
		}
	}
}
