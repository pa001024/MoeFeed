package service

import (
	"log"
	"testing"
)

func TestSendMailSync(t *testing.T) {
	err := Mail.SendMailSync("pa001024@qq.com", "hello world", "<p>context</p>", 2)
	if err != nil {
		log.Println(err)
	}
}
