package main

import (
	"context"
)

func init() {
	ctx, cancel := context.WithCancel(context.Background())
	dbCancel = cancel
	initDB(ctx)
}

func main() {
	// file, err := ioutil.ReadFile("tokens.json")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// tokens := &Tokens{}
	// err = json.Unmarshal(file, tokens)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bot, err := tb.NewBot(tb.Settings{
	// 	Token:  tokens.ApiToken,
	// 	Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// b := CreateBot(bot, tokens)

	// b.Init()
}
