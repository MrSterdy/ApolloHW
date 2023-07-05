package main

import (
	_ "embed"

	"github.com/MrSterdy/ApolloHW/bot"
	"github.com/MrSterdy/ApolloHW/config"
	"github.com/MrSterdy/ApolloHW/db"
	"github.com/MrSterdy/ApolloHW/webapp"
)

//go:generate qtc

//go:embed config.json
var file string

func main() {
	config.Init(file)

	db.InitMySQL()

	bot.Start()
	webapp.Start()
}
