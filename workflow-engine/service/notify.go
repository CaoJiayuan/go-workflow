package service

import (
	config "github.com/CaoJiayuan/go-workflow/workflow-config"
	"github.com/enorith/feather"
)

var conf = *config.Config

func Notify(data interface{}) {
	c := feather.NewClient()

	c.Post(conf.NotifyUrl, feather.RequestOptions{Json: data})
}
