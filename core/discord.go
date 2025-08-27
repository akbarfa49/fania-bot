package core

import "github.com/bwmarrin/discordgo"

func GetDiscordManagableGuild(discord *discordgo.Session) ([]discordgo.UserGuild, error) {

	unfilteredGuilds, err := discord.UserGuilds(0, "", "", false, discordgo.WithRetryOnRatelimit(true))
	if err != nil {
		return []discordgo.UserGuild{}, err
	}
	out := make([]discordgo.UserGuild, 0, len(unfilteredGuilds))
	for _, v := range unfilteredGuilds {
		if (v.Permissions & discordgo.PermissionManageServer) == 0 {
			continue
		}
		out = append(out, *v)
	}
	return out, err
}

func GetDiscordChanelList(discord *discordgo.Session, guildID string) ([]*discordgo.Channel, error) {
	return discord.GuildChannels(guildID)
}
