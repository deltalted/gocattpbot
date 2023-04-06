package config

import (
	genv "github.com/sakirsensoy/genv"
)

var BOT_TOKEN string = genv.Key("BOT_TOKEN").String()

