package botcommands

import (
	issuesdatabase "NinjoyBot/IssuesDatabase"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func (B *Bot) closeIssue(args []string) {

	db, err := gorm.Open("sqlite3", "issues.db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	var issue issuesdatabase.Issue
	db.Where("channel = ?", B.discordMessageCreate.ChannelID).First(&issue)

	if issue.Channel != "" {
		B.discordSession.ChannelDelete(B.discordMessageCreate.ChannelID)
		db.Delete(&issue)
	}
}

func (B *Bot) modmail() {
	var count int

	db, err := gorm.Open("sqlite3", "issues.db")
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer db.Close()

	db.Model(&issuesdatabase.Issue{}).Count(&count)
	if count > maxIssuesOpen {
		var issue issuesdatabase.Issue
		db.Where("sender = ?", B.discordMessageCreate.Author.ID).First(&issue)

		// if not found create
	} else {
		var issue issuesdatabase.Issue
		db.Where("sender = ?", B.discordMessageCreate.Author.ID).First(&issue)

		// if found get channel ID and send message to this channel
		if issue.Sender != "" {
			B.discordSession.ChannelMessageSend(issue.Channel, "From "+B.discordMessageCreate.Author.Username+"\n"+B.discordMessageCreate.Content)
		} else {
			// B.discordMessageCreate.Author.Username+B.discordMessageCreate.Author.Discriminator
			data := discordgo.GuildChannelCreateData{
				Name:     B.discordMessageCreate.Author.Username + B.discordMessageCreate.Author.Discriminator,
				NSFW:     false,
				ParentID: "625375153622351912",
				Type:     0,
			}
			channelCreated, err := B.discordSession.GuildChannelCreateComplex("614145482780049439", data)
			if err != nil {
				log.Println(err.Error())
				return
			}
			log.Println(channelCreated.ID)
			newIssue := issuesdatabase.Issue{Sender: B.discordMessageCreate.Author.ID, Channel: channelCreated.ID}
			db.Create(&newIssue)
			B.discordSession.ChannelMessageSend(newIssue.Channel, "From "+B.discordMessageCreate.Author.Username+"\n"+B.discordMessageCreate.Content)
		}
	}
}
