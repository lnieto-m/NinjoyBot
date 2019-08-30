package botcommands

import "github.com/bwmarrin/discordgo"

func OnServerJoin(s *discordgo.Session, m *discordgo.GuildMemberAdd) {

	s.ChannelMessageSend("614145482780049441", "Wesh alors <@"+m.User.ID+"> <#614150901229289527> .")
}
