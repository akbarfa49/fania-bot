package core

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fania-bot/model"
	"fania-bot/platforms/tiktok"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

type streamWatcherData struct {
	Platform            string
	UserIdentifier      string
	UserName            string
	UniqueID            string
	StreamIdentifier    string
	StreamURL           string
	StreamThumbnailURL  string
	ViewerCount         int
	Title               string
	ProfileURL          string
	ProfileThumbnailURL string
	StreamCategory      string
}

func (core *Core) runStreamWatcher(ctx context.Context, ch chan<- streamWatcherData) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan<- streamWatcherData, ctx context.Context) {
		defer wg.Done()
		log.Logger.Println("fetching streamer data")
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
	mainloop:
		for {
			follow, _, err := core.TiktokClient.CheckFollowingLive()
			if err != nil {
				log.Err(err).Str("location", "Core.runStreamWatcher").Msg("failed to check following live")
				return
			}
			for _, v := range follow.Data {
				if len(v.Data.Owner.OwnRoom.RoomIdsStr) == 0 {
					continue
				}
				ch <- streamWatcherData{
					Platform:         "tiktok",
					UserIdentifier:   v.Data.Owner.IDStr,
					UserName:         v.Data.Owner.Nickname,
					StreamIdentifier: v.Data.Owner.OwnRoom.RoomIdsStr[0],
					UniqueID:         v.Data.Owner.DisplayID,
					StreamURL:        "https://tiktok.com/@" + v.Data.Owner.DisplayID + "/live",
				}
			}
			select {
			case <-ticker.C:
			case <-ctx.Done():
				break mainloop
			}
		}

	}(ch, ctx)
	wg.Wait()
	close(ch)
}

func (core *Core) RunStreamNotifier() {
	ctx := core.Context
	ch := make(chan streamWatcherData, 10)
	// use mode guest for now because it's impossible to follow everyone 1 by 1 since tiktok make user cant follow other for bot prevention purpose
	// go core.runStreamWatcher(ctx, ch)
	go core.streamChecker(ctx, ch)
	for v := range ch {

		history, isExist, err := core.StreamRepo.FindLatestHistoryByID(ctx, v.Platform, v.UserIdentifier)
		if err != nil {
			log.Err(err).Msg("failed to find stream history by id")
			continue
		}
		if isExist && history.StreamIdentifier == v.StreamIdentifier {
			continue
		}
		if !isExist {
			history.Platform = v.Platform
			history.UserIdentifier = v.UserIdentifier
		}
		history.CreatedAt = time.Now()
		history.StreamIdentifier = v.StreamIdentifier
		notifData, err := core.StreamRepo.FindActiveNotificationByUserIdentifierAndStreamPlatform(ctx, v.UserIdentifier, v.Platform)
		if err != nil {
			log.Err(err).Msg("failed to find notification by user identifier")
			continue
		}
		if len(notifData) == 0 {
			continue
		}
		if err := core.StreamRepo.InsertHistory(ctx, history); err != nil {
			log.Err(err).Msg("failed to insert stream history")
			// continue
		}
		// if err := core.StreamRepo.CreateNotifDelay(ctx, v.Platform+":"+v.UserIdentifier, time.Now().Add(10*time.Minute)); err != nil {
		// 	log.Err(err).Msg("failed to insert notif delay")
		// 	continue
		// }
		go core.sendStreamNotifier(notifData, v)
	}
}

func (core *Core) streamChecker(ctx context.Context, ch chan streamWatcherData) {
	secondTicker := 60
	if second, err := strconv.Atoi(os.Getenv("stream_checker_ticker")); err == nil {
		secondTicker = second
	}
	ticker := time.NewTicker(time.Duration(secondTicker) * time.Second)
	defer ticker.Stop()
mainloop:
	for {
		notifData, err := core.StreamRepo.FindActiveNotificationStreamPlatformAndUserUniqueIDGroupByUserIdentifier(ctx)
		if err != nil {
			log.Err(err).Msg("failed to find active notification group by user identifier")
			continue
		}
		wg := sync.WaitGroup{}
		for _, v := range notifData {
			wg.Add(1)

			go func(ctx context.Context, notifData model.StreamNotificationTable, ch chan streamWatcherData, wg *sync.WaitGroup) {
				defer wg.Done()
				// if core.StreamRepo.NotifDelayIsActive(ctx, notifData.StreamPlatform+":"+notifData.UserIdentifier) {
				// 	return
				// }
				minSleep := 1  // in seconds
				maxSleep := 50 // in seconds
				sleepDuration := time.Duration(rand.Intn(maxSleep-minSleep+1)+minSleep) * time.Millisecond
				time.Sleep(sleepDuration)

				switch notifData.StreamPlatform {
				case "tiktok":
					data, b, err := core.TiktokClient.GetWebcastPreloadRoom(ctx, notifData.UserIdentifier)
					if err != nil {
						if errors.Is(err, tiktok.ErrTiktokInternalError) {
							fmt.Println(string(b))
							return
						}
						// fmt.Println(userUniqueID)
						// fmt.Println(notifData.UserIdentifier)
						// fmt.Println(notifData.UserUniqueID)
						log.Err(err).Str("location", "Core.streamChecker").Msg("failed to get user detail by unique id")
						return
					}

					if data.Extra.Reason != "" {
						// fmt.Println(notifData.UserUniqueID, data.Extra.Reason)
						return
					}
					roomID := ""
					if len(data.Data.Owner.OwnRoom.RoomIdsStr) > 0 {
						roomID = data.Data.Owner.OwnRoom.RoomIdsStr[0]
					} else if len(data.Data.Owner.OwnRoom.RoomIds) > 0 {
						roomID = strconv.FormatInt(data.Data.Owner.OwnRoom.RoomIds[0], 10)
					}
					if roomID == "" {
						return
					}
					title := "Live Now!"
					streamThumbnailURL := data.Data.Owner.AvatarLarge.URLList[0]
					roomCompleteData, _, err := core.TiktokClient.WebcastEnterRoom(ctx, roomID)
					if roomCompleteData.Extra.Reason == "" && err == nil {
						// fmt.Println(userUniqueID)
						title = roomCompleteData.Data.Title
						if len(roomCompleteData.Data.RectangleCoverImg.URLList) > 0 {
							streamThumbnailURL = roomCompleteData.Data.RectangleCoverImg.URLList[0]
						} else if len(roomCompleteData.Data.Cover.URLList) > 0 {
							streamThumbnailURL = roomCompleteData.Data.Cover.URLList[0]
						} else {
							fmt.Printf("failed to get stream thumbnail for platform %s : %s(%s)\n", notifData.StreamPlatform, notifData.UserUniqueID, notifData.UserIdentifier)
						}
					}
					category := ""
					switch roomCompleteData.Data.Hashtag.ID {
					case 5:
						category = roomCompleteData.Data.GameTagDetail.DisplayName
					default:
						category = roomCompleteData.Data.Hashtag.Title
					}
					if category == "" {
						category = "Unknown"
					}

					ch <- streamWatcherData{
						Platform:            "tiktok",
						UserIdentifier:      data.Data.Owner.IDStr,
						UserName:            data.Data.Owner.Nickname,
						StreamIdentifier:    data.Data.Owner.OwnRoom.RoomIdsStr[0],
						UniqueID:            data.Data.Owner.DisplayID,
						StreamURL:           "https://tiktok.com/@" + data.Data.Owner.DisplayID + "/live",
						ViewerCount:         data.Data.UserCount,
						Title:               title,
						StreamThumbnailURL:  streamThumbnailURL,
						ProfileURL:          "https://tiktok.com/@" + data.Data.Owner.DisplayID,
						ProfileThumbnailURL: data.Data.Owner.AvatarThumb.URLList[0],
						StreamCategory:      category,
					}
				default:
				}
			}(ctx, v, ch, &wg)

		}
		wg.Wait()
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break mainloop
		}
	}
}
func (core *Core) sendStreamNotifier(data []model.StreamNotificationTable, streamData streamWatcherData) {
	for _, v := range data {
		tmplIdentifier := v.UserIdentifier + v.Guild + v.Channel
		// tmplt, ok := core.TextTemplate.Load(tmplIdentifier)

		// if !ok {
		// 	tmplt = template.Must(template.New(tmplIdentifier).Parse(v.Message))

		// 	core.TextTemplate.Store(tmplIdentifier, tmplt)
		// }
		// tmpl := tmplt.(*template.Template)

		tmpl, _ := template.New(tmplIdentifier).Parse(v.Message)
		var sb strings.Builder
		tmpl.Execute(&sb, streamData)
		switch v.Platform {
		case "discord":
			disc := model.DiscordMetadata{}
			if v.Metadata.Valid {
				json.Unmarshal([]byte(v.Metadata.String), &disc)
				if disc.AvatarURI == "" {
					disc.AvatarURI = streamData.ProfileThumbnailURL
				}
				if disc.Username == "" {
					disc.Username = streamData.UserName
				}
			}
			field := make([]*discordgo.MessageEmbedField, 0, 1)
			viewerCount := strconv.Itoa(streamData.ViewerCount)
			field = append(field, &discordgo.MessageEmbedField{
				Name:   "Viewers",
				Value:  viewerCount,
				Inline: true,
			})
			field = append(field, &discordgo.MessageEmbedField{
				Name:   "Category",
				Value:  streamData.StreamCategory,
				Inline: true,
			})
			author := &discordgo.MessageEmbedAuthor{
				Name:    streamData.UserName,
				URL:     streamData.StreamURL,
				IconURL: streamData.ProfileThumbnailURL,
			}
			// if disc.Username != "" {
			// 	author.Name = disc.Username
			// }
			// if disc.AvatarURI != "" {
			// 	author.IconURL = disc.AvatarURI
			// }
			image := &discordgo.MessageEmbedImage{
				URL:    streamData.StreamThumbnailURL,
				Width:  720,
				Height: 720,
			}
			embed := &discordgo.MessageEmbed{
				Title:  streamData.Title,
				Author: author,
				Fields: field,
				URL:    streamData.StreamURL,
				Image:  image,
				Color:  0x7700ff,
				Footer: &discordgo.MessageEmbedFooter{Text: "Fantie"},
			}

			button := discordgo.Button{
				Label: "Watch Stream",
				Style: discordgo.LinkButton,
				URL:   streamData.StreamURL,
			}

			param := SendDiscordParam{
				Webhook: struct {
					ID        string
					Token     string
					Username  string
					AvatarURI string
				}{
					ID:        disc.WebhookID,
					Token:     disc.WebhookToken,
					Username:  disc.Username,
					AvatarURI: disc.AvatarURI,
				},
				GuildID:     v.Guild,
				ChannelID:   v.Channel,
				Message:     sb.String(),
				TagEveryone: true,
				Embed:       []*discordgo.MessageEmbed{embed},
				Component:   []discordgo.MessageComponent{discordgo.ActionsRow{Components: []discordgo.MessageComponent{button}}},
			}
			if err := core.SendTextDiscord(param); err != nil {
				log.Err(err).Str("guild", v.Guild).Str("channel", v.Channel).Msg("failed to send message")
				continue
			}
		default:
		}
	}
}

type WatchStreamer_In struct {
	StreamPlatform        string `json:"stream_platform"`
	StreamerUniqueID      string `json:"streamer_unique_id"`
	Message               string `json:"message"`
	SendPlatform          string `json:"send_platform"`
	Guild                 string `json:"guild"`
	Channel               string `json:"channel"`
	CustomSenderUsername  string `json:"custom_sender_username"`
	CustomSenderAvatarURI string `json:"custom_sender_avatar_uri"`
	WithPersonalization   bool   `json:"with_personalization"`
}

type WatchStreamer_Out struct {
	Error model.Errors `json:"error"`
}

func (core *Core) WatchStreamer(ctx context.Context, in WatchStreamer_In) WatchStreamer_Out {
	out := WatchStreamer_Out{}
	userID := ""
	uniqueID := ""
	switch in.StreamPlatform {
	case "tiktok":
		user, _, err := core.TiktokClient.GetUserDetailByUniqueID(ctx, in.StreamerUniqueID)
		if err != nil {
			if err == tiktok.ErrUserDoesntExist {
				out.Error.ErrorCode = "BAD_REQUEST"
				out.Error.ErrorMessage = "User Doesnt Exist"
				return out
			}
			log.Err(err).Msg("failed to get use by unique id")
			out.Error.ErrorCode = "INTERNAL_ERROR"
			out.Error.ErrorMessage = "Internal Error"
			return out
		}
		userID = user.UserInfo.User.ID
		uniqueID = user.UserInfo.User.UniqueID
	default:
		out.Error.ErrorCode = "BAD_REQUEST"
		out.Error.ErrorMessage = "User Doesnt Exist"
		return out
	}
	metadata := sql.NullString{}
	switch in.SendPlatform {
	case "discord":
		if !in.WithPersonalization {
			break
		}
		webhookList, err := core.DiscordBot.ChannelWebhooks(in.Channel)
		if err != nil {
			return out
		}
		webhookID := ""
		webhookToken := ""
		for _, v := range webhookList {
			if v.Name == os.Getenv("discord_webhook_name") {
				webhookID = v.ID
				webhookToken = v.Token
				break
			}
		}
		if webhookID == "" {
			wh, err := core.DiscordBot.WebhookCreate(in.Channel, os.Getenv("discord_webhook_name"), "")
			if err != nil {
				return out
			}
			webhookID = wh.ID
			webhookToken = wh.Token
		}
		data := model.DiscordMetadata{}
		data.Username = in.CustomSenderUsername
		data.AvatarURI = in.CustomSenderAvatarURI
		data.WebhookID = webhookID
		data.WebhookToken = webhookToken
		b, _ := json.Marshal(data)
		metadata.String = string(b)
		metadata.Valid = true
	default:
		out.Error.ErrorCode = "BAD_REQUEST"
		out.Error.ErrorMessage = "Platform Doesnt Supported"
		return out
	}

	if err := core.StreamRepo.CreateNotification(ctx, model.StreamNotificationTable{
		StreamPlatform: in.StreamPlatform,
		Platform:       in.SendPlatform,
		Guild:          in.Guild,
		Channel:        in.Channel,
		UserIdentifier: userID,
		UserUniqueID:   uniqueID,
		Message:        in.Message,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		Metadata:       metadata,
	}); err != nil {
		log.Err(err).Msg("failed to create notification")
		out.Error.ErrorCode = "INTERNAL ERROR"
		out.Error.ErrorMessage = "Internal Error"
		return out
	}
	return out
}
