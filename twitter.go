/*
	Basic twitter api setup
	Not in use at the momment
*/
package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TwitterHandler struct {
	stream *twitter.Stream
}

func (T *TwitterHandler) getTwitterFeed(discord *discordgo.Session) {
	userID := int64(349591688)
	log.Print("Connecting to Twitter...")

	config := oauth1.NewConfig("nFAyRc7404OlFnGrzNyglJq1d", "RYnMhEG49dH0O2Pgz0Jw4IIqViBL2rZ0DbkLnSTpETjZT12ENN")
	token := oauth1.NewToken("349591688-uYVZ4ricRGjjoa4VAHB29S52MGG0J9Orat6gA7Lx", "pKtZrpPjMojJYBr8OhN1edZA7UwE8ONDAZZIcHM56iA5h")
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		log.Println("Tweet received:", tweet.Text, "SAUCE:", tweet.User.ID)
		if tweet.User.ID == userID {
			messageText := tweet.Text + "\n\n [View tweet](https://twitter.com/" + tweet.User.ScreenName + "/status/" + tweet.IDStr + ")"
			discord.ChannelMessageSend("614513898191847475", messageText)
		}
	}
	demux.Warning = func(warning *twitter.StallWarning) {
		log.Println("Warning received:", warning.Code, warning.Message)
	}

	log.Println("Done.")

	params := &twitter.StreamFilterParams{
		Follow:        []string{"349591688"},
		StallWarnings: twitter.Bool(true),
	}

	stream, err := client.Streams.Filter(params)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("Starting stream...")
	go demux.HandleChan(stream.Messages)
	T.stream = stream
}
