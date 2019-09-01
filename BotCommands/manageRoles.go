package botcommands

import "log"

var PingableRoles = []string{
	"Newsletter",
}

func contains(a []string, x string) (bool, int) {
	for index, n := range a {
		if x == n {
			return true, index
		}
	}
	return false, -1
}

func (B *Bot) manageRoles(args []string) {

	if len(args) == 0 {
		B.help(args)
		return
	}

	// Check if user has given role
	// If user doesn't have it and is avalaible to any user -> add role to member
	// If user tagged then remove it

	member, _ := B.discordSession.GuildMember(B.discordMessageCreate.GuildID, B.discordMessageCreate.Author.ID)

	log.Println(member.Roles)

	rolesList, roleIDList := []string{}, []string{}
	rolesListRaw, _ := B.discordSession.GuildRoles(B.discordMessageCreate.GuildID)
	for _, role := range rolesListRaw {
		// log.Println(role.Name, role.)
		rolesList = append(rolesList, role.Name)
		roleIDList = append(roleIDList, role.ID)
	}

	inRoleList, roleIDIndex := contains(rolesList, args[0])
	pingable, _ := contains(PingableRoles, args[0])

	if inRoleList == false || pingable == false {
		// Ne pas oublier de rajouter un check pour les roles d'admins
		// A ajouter une liste des roles modifiables
		log.Println("No such role : " + args[0])
		return
	}

	inMemberRoleList, _ := contains(member.Roles, roleIDList[roleIDIndex])
	log.Println("in member list", inMemberRoleList)

	if inMemberRoleList == true {
		B.discordSession.GuildMemberRoleRemove(B.discordMessageCreate.GuildID, member.User.ID, roleIDList[roleIDIndex])
		B.discordSession.ChannelMessageSend(B.discordMessageCreate.ChannelID, "Tag removed.")
	} else {
		B.discordSession.GuildMemberRoleAdd(B.discordMessageCreate.GuildID, member.User.ID, roleIDList[roleIDIndex])
		B.discordSession.ChannelMessageSend(B.discordMessageCreate.ChannelID, "Tag added.")
	}

}
