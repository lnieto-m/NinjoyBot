package botcommands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// OnServerJoin mention the user that just joined the server
func OnServerJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	s.ChannelMessageSend("614145482780049441", "Bienvenue <@"+m.User.ID+">. <#614150901229289527> avant toute chose. Tu peux aussi utiliser la commande `!help` dans <#614151630820212736> pour plus d'informations sur le bot.")
}

// OnMessageCreate handle the commands received by the bot
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	bot := start(s, m)

	// Ignore self message
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!") {
		bot.execute(m.Content[1:])
	}
}
