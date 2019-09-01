package botcommands

import "log"

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func (B *Bot) manageRoles(args []string) {

	// Check if user has given role
	// If user doesn't have it and is avalaible to any user -> add role to member
	// If user tagged then remove it

	member, _ := B.discordSession.GuildMember(B.discordMessageCreate.GuildID, B.discordMessageCreate.Author.ID)

	log.Println(member.Roles)

	rolesList := []string{}
	rolesListRaw, _ := B.discordSession.GuildRoles(B.discordMessageCreate.GuildID)
	for _, role := range rolesListRaw {
		rolesList = append(rolesList, role.Name)
	}

	if contains(rolesList, args[0]) == false {
		// Ne pas oublier de rajouter un check pour les roles d'admins
		// A ajouter une liste des roles modifiables
		log.Println("No such role : " + args[0])
		return
	}

}
