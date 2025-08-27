package core

import (
	"context"
	"fania-bot/platforms/tiktok"
	"fania-bot/repository"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Core struct {
	TiktokClient *tiktok.Client
	StreamRepo   *repository.Streamer
	DiscordBot   *discordgo.Session
	TextTemplate *sync.Map
	cancel       context.CancelFunc
	Context      context.Context
}

func New() *Core {
	ctx, cancel := context.WithCancel(context.Background())
	dsc, _ := discordgo.New("Bot " + os.Getenv("discord_bot_token"))
	dsc.StateEnabled = true
	pool, err := pgxpool.New(ctx, os.Getenv("db_url"))
	if err != nil {
		panic(err)
	}
	core := &Core{
		TiktokClient: tiktok.New(os.Getenv("tt_cookie"), os.Getenv("tt_device_id")),
		DiscordBot:   dsc,
		TextTemplate: &sync.Map{},
		StreamRepo:   repository.NewStreamer(pool),
		cancel:       cancel,
		Context:      ctx,
	}
	return core
}

func (c *Core) Shutdown() {
	c.cancel()
}
