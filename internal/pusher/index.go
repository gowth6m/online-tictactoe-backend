package pusher

import (
	"github.com/pusher/pusher-http-go/v5"
	"online-tictactoe/internal/config"
)

var Client *pusher.Client

func Init() {
	Client = &pusher.Client{
		AppID:   config.AppConfig().Pusher.AppID,
		Key:     config.AppConfig().Pusher.Key,
		Secret:  config.AppConfig().Pusher.Secret,
		Cluster: config.AppConfig().Pusher.Cluster,
		Secure:  true,
	}
}
