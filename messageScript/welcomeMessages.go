package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var artMessage = `Welcome to our Art Channel!

In this channel you can share your own created art and in game cinematic screenshots.

**To be able to upload images in this channel you will need the Artist role. DM NinjoyBot with your work (your own creations). If it suits the channel you will be given the role.**

**Please read and follow the guidelines below before posting. Posts that don't follow these rules will be deleted.**

• If you use this channel to bypass the no uploading/embeds in the discussions category to post memes or off topic content you may get muted.
• By posting in this art channel, you understand that other users may repost your art. We cannot control this unfortunately, and encourage you always use a signature or watermark on your drawings to protect your work to the best capacity. If you are going to use someone's art, please be a nice person and credit them for their work.
• Keep art SFW - This means no gore, nudity or likewise
• Be mindful of your judgements and for what others might say. Please remember negativity is not allowed. Creative criticism and negativity are two different things. this is a place for artists to show off, not to be brought down

It you aren't sure sure whether or not content you'd like to share, meets our rules, message  <@614523245496107041> to receive clarification`

var audioMessage = `Welcome to our Audio Channel!

In this channel you can share your own created Musics and Sounds.

**To be able to upload audio in this channel you will need the Musician role. DM NinjoyBot with your work (your own creations). If it suits the channel you will be given the role.**

**Please read and follow the guidelines below before posting. Posts that don't follow these rules will be deleted.**

• If you use this channel to bypass the no uploading/embeds in the discussions category to post memes or off topic content you may get muted.
• By posting in this audio channel, you understand that other users may repost your audio. We cannot control this unfortunately. If you are going to use someone's audio, please be a nice person and credit them for their work.
• Keep audio SFW - This means no gore, nudity or likewise
• Be mindful of your judgements and for what others might say. Please remember negativity is not allowed. Creative criticism and negativity are two different things. this is a place for artists to show off, not to be brought down

It you aren't sure sure whether or not content you'd like to share, meets our rules, message  <@614523245496107041> to receive clarification`

var ruleMessage1 = `Welcome to the **NINJOY** Official Discord!

We are a french video game studio.
We are currently working on MagiColor, a game where you make balls collide.

Currently the game is in Active Development with the Beta version published!
 `

var ruleMessage2 = `
While you wait 10 minutes to be able to talk, please familiarize yourself with our rules below.
**1.** Be nice or leave! 
  1.1. No personal attacks, offensive language, harassment, witchhunting, sexism, racism, hate speech or other disruptive behavior, this also applies to voice chats. Disruptive behaviour in voice includes voice changers, soundboards, extremely loud noises, etc.
  1.2. Do not ping/mention users who are not currently engaged in the chat unless they're okay with it. This rule applies especially for our @Dev Team and Moderators
  1.3. Do not impersonate anyone — including Developers, Moderators, or anyone else
  1.4. Do not act as if you are able to carry out Staff actions if you're not a Staff member.
**2.** No offensive or otherwise inappropriate nicknames or profile pictures
  2.1. This includes blank or invisible names and excessive use of noisy or unusual unicode characters (Names that break these rules will be modified)
**3.** Don't spam
  3.1. Includes excessive amounts of messages, emojis, capital letters, pings/mentions, etc.
  3.2 Do NOT send Discord invite links, they will be removed by our Bot and if you spam them you can be subject to a mute.
**4.** NSFW content is NOT allowed
  4.1. NSFW = Not Safe For Work, i.e. porn, gore, suggestive content, etc.
**5.** Use the appropriate channels for discussions. All off-topic/non NINJOY games related discussion should happen in General
  5.1. Advertising is NOT allowed in any channels.
**6.** English, please
  6.1. Keep all discussion in text channels and general voice channels in English
**7.** No scam links, URL shorteners, IP grabbers, etc.
  7.1 No discussion about ANYTHING illegal, no matter what.`

var ruleMessage6 = `**8.** Spoilers must use spoiler tags and be labelled
  8.1. Generally applies to spoilers regarding recent/upcoming movies, games, series, etc.
  8.2. Example of a labelled spoiler tag: Star Wars spoiler: ||darth vader is luke's father||
  8.3. Unlabelled or mislabelled spoiler tags will be removed
**9.** Listen to server staff
  9.1. If a moderator tells you to stop doing something, stop it
  9.2. Don't argue about mod decisions in chat. If you'd like to discuss or dispute a decision, please message <@614523245496107041>.

**In addition to these rules, the moderation team reserves the right to remove messages and users from the server that are detrimental to the discussion and community. Since this is the first page you are seeing, ignorance of the rules does not excuse breaking them.**

Invite your friends to the server! 
https://discord.gg/jXNaBxk

Need to talk to the mods or want to report something? DM NinjoyBot! When you message, an available mod will be with you shortly. If reporting, please include all relevant information (name, discord name, why reporting, any links) so that we may handle the issue promptly.`

var ruleMessage3 = `
<@&624336644090363925>: Sergio, Lead Developer at NINJOY
<@&614150284712869899> : Developers working at NINJOY
<@&624336801623965706> : Community Manager for Ninjoy communities
<@&614511934342955043> : Content moderators for the Ninjoy Discord
<@&614525663101059102> : Bot that gets you directly in touch with a Moderator. If a Moderator is needed, please DM <@614523245496107041>
<@&624337221767397398> : Vanity role for those that have participated in our Alpha Tests
<@&617703532967231513> : Members of our community who have signed up for Ninjoy Games/Server notifications, self assignable in #role-assignment`

var ruleMessage4 = `<#614150901229289527> - Where you are now, which displays information and rules about NINJOY, it's games and it's Discord
<#614151282370019351> - A channel for Announcements we make to the community! These announcements can include games/server news, or whatever we think you should see!
<#614151449072369699> - Self explanatory
<#614151477757345842> - A channel that contains all of our tweets!
--
<#614145482780049441> - The main discussion channel for anything not related to NINJOY GAME'S
<#614151591959724042> - A channel for members to post suggestions for the NINJOY Discord
<#614530155532124230> - A channel for members to post theirs and others musics and sounds. This channel is STRICTLY for music and sounds! Any and all music and sounds must be accompanied with a source, unless it is your own original works. See channel pins to see guidelines for this channel.
<#614151611958296579> for members to post theirs and others art! This channel is STRICTLY for Art! Any and all art must be accompanied with a source, unless it is your own original works. See channel pins to see guidelines for this channel.
<#614151630820212736> - A channel for bot commands
--
<#614153006434222083> - The main discussion channel for MAGICOLOR, this channel is reserved for MAGICOLOR discussion only.
<#614153058397585439> - A channel for members to post suggestions for MAGICOLOR`

var ruleMessage5 = `Check us out on social media!
Follow us on Twitter! https://twitter.com/ninjoy3 
Like us on Facebook! https://www.facebook.com/ninjoy42 
Follow us on Instagram! https://www.instagram.com/ninjoy42/`

func main() {
	discord, err := discordgo.New("Bot NjE0NTIzMjQ1NDk2MTA3MDQx.XWAuBA.fQx4_llAhlk8SwjUZFSKseRb0_Y")
	if err != nil {
		log.Println("Error creating Discord session, ", err)
		return
	}

	// discord.ChannelMessageDelete("614150901229289527", "624630280103133198")
	// discord.ChannelMessageDelete("614150901229289527", "624625886188273711")
	// discord.ChannelMessageDelete("614150901229289527", "624623029653340170")
	// discord.ChannelMessageDelete("614150901229289527", "624623689576611859")

	// discord.ChannelMessageSend("614150901229289527", ruleMessage1)
	// discord.ChannelMessageSend("614150901229289527", ruleMessage2)
	// discord.ChannelMessageSend("614150901229289527", ruleMessage6)
	// discord.ChannelMessageSend("614150901229289527", ruleMessage3)
	// discord.ChannelMessageSend("614150901229289527", ruleMessage4)
	// discord.ChannelMessageSend("614150901229289527", ruleMessage5)
	discord.ChannelMessageEdit("614150901229289527", "624626239570837504", ruleMessage6)
	// discord.ChannelMessageSend("614151611958296579", artMessage)
	// discord.ChannelMessageSend("614530155532124230", audioMessage)

	// discord.ChannelMessageDelete("614150901229289527", "624615652761665557")
	// 624612992738394112
	discord.Close()
}
