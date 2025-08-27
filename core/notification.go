package core

import "github.com/bwmarrin/discordgo"

type SendDiscordParam struct {
	Mode    string
	Webhook struct {
		ID        string
		Token     string
		Username  string
		AvatarURI string
	}
	GuildID     string
	ChannelID   string
	Message     string
	TagEveryone bool
	TagRole     []string
	TagUser     []string
	Component   []discordgo.MessageComponent
	Embed       []*discordgo.MessageEmbed
}

func (s SendDiscordParam) Transform() *discordgo.MessageSend {
	msg := &discordgo.MessageSend{
		Content:         s.Message,
		AllowedMentions: &discordgo.MessageAllowedMentions{Parse: make([]discordgo.AllowedMentionType, 0, 3)},
		Embeds:          s.Embed,
		Components:      s.Component,
	}
	// allow mention section
	if s.TagEveryone {
		msg.AllowedMentions.Parse = append(msg.AllowedMentions.Parse, discordgo.AllowedMentionTypeEveryone)
	}
	if len(s.TagRole) > 0 {
		msg.AllowedMentions.Parse = append(msg.AllowedMentions.Parse, discordgo.AllowedMentionTypeRoles)
	}
	if len(s.TagUser) > 0 {
		msg.AllowedMentions.Parse = append(msg.AllowedMentions.Parse, discordgo.AllowedMentionTypeUsers)
	}
	return msg
}

func (c *Core) SendTextDiscord(param SendDiscordParam) error {
	if param.Webhook.ID != "" {
		return c.SendWebhookMessage(param)
	}
	_, err := c.DiscordBot.ChannelMessageSendComplex(param.ChannelID, param.Transform())
	if err != nil {
		return err
	}
	return nil
}

func (c *Core) SendWebhookMessage(param SendDiscordParam) error {
	msg := param.Transform()
	_, err := c.DiscordBot.WebhookExecute(param.Webhook.ID, param.Webhook.Token, true, &discordgo.WebhookParams{
		Content:         msg.Content,
		Username:        param.Webhook.Username,
		AvatarURL:       param.Webhook.AvatarURI,
		TTS:             msg.TTS,
		Files:           msg.Files,
		Components:      msg.Components,
		Embeds:          msg.Embeds,
		AllowedMentions: msg.AllowedMentions,
	})
	return err
}
