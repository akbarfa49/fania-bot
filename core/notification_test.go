package core

import (
	"fmt"
	"testing"

	"github.com/bwmarrin/discordgo"
)

// func TestNotif(t *testing.T) {
// 	c := Core{}
// 	dsc, _ := discordgo.New("Bot NzE3NzgyNTQ1MTg5ODMwNzA2.GRj1kK.U-trVRh2MT1zYJdavpiQ8jORIM-sHNR9NYuHaA")
// 	c.DiscordBot = dsc
// 	// c.SendTextDiscord("", "1219793797831987210", "@everyone")
// }

func TestCreateWebhook(t *testing.T) {
	sess, _ := discordgo.New("")
	c := Core{DiscordBot: sess}
	hook, err := c.DiscordBot.WebhookCreate("1219793797831987210", "Fania", "https://images-ext-1.discordapp.net/external/hOwVrTxs2PnuK43x6C_QKAmB3lK_1pKXc9BypKuJ0DI/%3Flk3s%3Da5d48078%26x-expires%3D1713254400%26x-signature%3DYc4cREjTVOggEyUkIo3QMuX8AsY%253D/https/p16-sign-sg.tiktokcdn.com/aweme/720x720/tos-alisg-avt-0068/e602e7e9c884b5723deda6d920946f2d.webp?format=webp")
	if err != nil {
		panic(err)
	}
	fmt.Println(hook.ID, hook.Token)
}

func TestWebhook(t *testing.T) {
	sess, _ := discordgo.New("")
	c := Core{DiscordBot: sess}
	c.DiscordBot.WebhookExecute("1229302013951475732", "", true, &discordgo.WebhookParams{
		Content:   "halo",
		Username:  "BikDig",
		AvatarURL: "https://wallpapercave.com/wp/wp5922165.jpg",
	})
}
