package botcommands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var maxIssuesOpen = 10

// Bot : Base struct
type Bot struct {
	commandList          map[string]interface{}     // command list [commandName] - > corresponding function
	discordSession       *discordgo.Session         // Current discord session
	discordMessageCreate *discordgo.MessageCreate   // Stored message create event
	voiceConn            *discordgo.VoiceConnection // Stored current voice connexion
}

// Stores current session, message event and avalaible commands then return a Sakamoto object
func start(s *discordgo.Session, m *discordgo.MessageCreate) Bot {
	B := Bot{}
	B.discordSession = s
	B.discordMessageCreate = m
	B.commandList = map[string]interface{}{
		"help": func(args []string) { B.help(args) },
		"tag":  func(args []string) { B.manageRoles(args) },
	}
	return B
}

// Inspect the command map to get the corresponding function else does nothing
func (B *Bot) execute(commandInput string) {
	commandList := strings.Split(commandInput, " ")
	args := []string{}
	if len(commandList[1:]) > 0 {
		args = commandList[1:]
	}
	if command, ok := B.commandList[commandList[0]]; ok {
		command.(func([]string))(args)
	}
}
