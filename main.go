package main

import (
	botcommands "NinjoyBot/BotCommands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot NjE0NTIzMjQ1NDk2MTA3MDQx.XWAuBA.fQx4_llAhlk8SwjUZFSKseRb0_Y")
	if err != nil {
		log.Println("Error creating Discord session, ", err)
		return
	}

	discord.AddHandler(botcommands.OnServerJoin)
	discord.AddHandler(botcommands.OnMessageCreate)

	err = discord.Open()
	if err != nil {
		log.Println("Error opening connection, ", err)
		return
	}

	log.Println("Sakamoto at your service.")

	// go tweetHand.getTwitterFeed(discord)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill, syscall.SIGSEGV)
	<-signals

	// tweetHand.stream.Stop()
	discord.Close()
	log.Println("See you later.")
}
