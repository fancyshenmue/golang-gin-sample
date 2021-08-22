package app

import (
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	// viper
	V              = viper.New()
	URLList        []string
	SlackTokenID   string
	SlackChannelID string
)

type domain struct {
	Domain string
	Status bool
}
