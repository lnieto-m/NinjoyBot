package botcommands

import (
	"github.com/bwmarrin/discordgo"
)

func (B *Bot) help(args []string) {

	if B.discordMessageCreate.ChannelID != "614151630820212736" {
		return
	}

	rolesField := &discordgo.MessageEmbedField{
		Name:  "Roles List",
		Value: "`1`,`2`,`3`",
	}

	fields := []*discordgo.MessageEmbedField{
		rolesField,
	}

	helpMessage := &discordgo.MessageEmbed{
		Title:       "",
		Description: "`s!<base_command> <args>`",
		Fields:      fields,
		Color:       0x98BDF0,
	}

	B.discordSession.ChannelMessageSendEmbed(B.discordMessageCreate.ChannelID, helpMessage)
}
