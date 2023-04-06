package main

import (
	log "log"
	time "time"
	http "net/http"
	fmt "fmt"

	tele "gopkg.in/telebot.v3"

	_ "github.com/sakirsensoy/genv/dotenv/autoload"
	config "cattpbot/config"
)

func main() {
	pref := tele.Settings{
		Token: config.BOT_TOKEN,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", startHander)

	b.Handle(tele.OnText, textHandler)
	b.Handle(tele.OnQuery, inlineQueryHandler)

	b.Start()
}

func startHander(ctx tele.Context) error {
	text := "Send pictures of cats instead of HTTP status codes (see http.cat)."
	text += "\n\n"
	text += "Send any message to start!"
	text += "\n\n"
	text += "Original idea from https://github.com/KnorpelSenf/cattpbot (Author: @delta_alted)"

	return ctx.Send(text, tele.NoPreview)
}

func textHandler(ctx tele.Context) error {
	menu := &tele.ReplyMarkup{}
	menu.Inline(
		menu.Row(menu.Query("Search here", "")),
		menu.Row(menu.QueryChat("Share cat", "")),
	)

	return ctx.Send("I can search for HTTP cat images inline.", menu)
}

func inlineQueryHandler(ctx tele.Context) error {
	results := make(tele.Results, 1)
	url := fmt.Sprintf("https://http.cat/%s.jpg", ctx.Query().Text)

	res, _ := http.Get(url)
	if res.StatusCode != 200 {
		results[0] = &tele.PhotoResult{
			URL: "https://http.cat/404.jpg",
			ThumbURL: "https://http.cat/404.jpg",
		}
	} else {
		results[0] = &tele.PhotoResult{
			URL: url,
			ThumbURL: url,
		}
	}

	return ctx.Answer(&tele.QueryResponse{
		Results: results,
	})
}