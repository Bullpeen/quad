package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	xpost "github.com/jirwin/xpost-quadlek/pkg"
	"go.uber.org/zap"

	"github.com/Bullpeen/infobot"
	"github.com/Bullpeen/stox"
	gifs "github.com/jirwin/gifs-quadlek/src"
	"github.com/jirwin/quadlek/plugins/comics"
	"github.com/jirwin/quadlek/plugins/karma"
	"github.com/jirwin/quadlek/plugins/random"
	"github.com/jirwin/quadlek/plugins/spotify"
	"github.com/jirwin/quadlek/plugins/twitter"
	"github.com/jirwin/quadlek/quadlek"

	"github.com/urfave/cli"
)

const Version = "0.0.1"

func run(c *cli.Context) error {
	var apiToken string
	if c.IsSet("api-key") {
		apiToken = c.String("api-key")
	} else {
		err := cli.ShowAppHelp(c)
		if err != nil {
			return err
		}
		return cli.NewExitError("Missing --api-key arg.", 1)
	}

	var verificationToken string
	if c.IsSet("verification-token") {
		verificationToken = c.String("verification-token")
	} else {
		err := cli.ShowAppHelp(c)
		if err != nil {
			return err
		}
		return cli.NewExitError("Missing --verification-token arg.", 1)
	}

	dbPath := c.String("db-path")

	bot, err := quadlek.NewBot(context.Background(), apiToken, verificationToken, dbPath)
	if err != nil {
		zap.L().Error("error creating bot", zap.Error(err))
		return nil
	}

	err = bot.RegisterPlugin(karma.Register())
	if err != nil {
		fmt.Printf("error registering karma plugin: %s\n", err.Error())
		return nil
	}

	err = bot.RegisterPlugin(random.Register())
	if err != nil {
		fmt.Printf("error registering random plugin: %s\n", err.Error())
		return nil
	}

	err = bot.RegisterPlugin(spotify.Register())
	if err != nil {
		fmt.Printf("error registering spotify plugin: %s\n", err.Error())
		return nil
	}

	// TVDB not working
	//if c.IsSet("tvdb-key") {
	//	tvdbKey := c.String("tvdb-key")
	//
	//	err = bot.RegisterPlugin(nextep.Register(tvdbKey))
	//	if err != nil {
	//		fmt.Printf("error registering nextep plugin: %s\n", err.Error())
	//		return nil
	//	}
	//}

	err = bot.RegisterPlugin(infobot.Register())
	if err != nil {
		fmt.Printf("error registering infobot plugin: %s\n", err.Error())
		return nil
	}

	err = bot.RegisterPlugin(twitter.Register(
		c.String("twitter-consumer-key"),
		c.String("twitter-consumer-secret"),
		c.String("twitter-access-token"),
		c.String("twitter-access-secret"),
		// These must be twitter user ids, not names. https://tweeterid.com/ for easy conversion between the two.
		map[string]string{
			"138203134":           "politics-feed",  // @AOC
			"783792992":           "politics-feed",  // @IlhanMN
			"31013444":            "politics-feed",  // @AyannaPressley
			"1079769536730140672": "politics-feed",  // @RepRashida
			"1081222837459996672": "politics-feed",  // @RepKatiePorter
			"151444950":           "politics-feed",  // @DaviSusan
			"2315512764":          "politics-feed",  // @bellingcat
			"976366106561490944":  "artfolio",       // @DrawnDavidsOff
			"921111554371682304":  "artfolio",       // @DrawnDavidson
			"1581511":             "wwdc",           // @macrumorslive
			"29472803":            "final-frontier", // @NASAWebb
			"17217640":            "final-frontier", // @SpaceFlightNow
			"14091091":            "final-frontier", // @NASAHubble
			"18831926":            "covid19-feed",   // @DrEricDing
			"88589013":            "covid19-feed",   //@fitterhappierAJ
			"920943710254313472":  "covid19-feed",   //@dgurdasani1
			"29431996":            "covid19-feed",   //@patricklsimpson

			//added for retweet testing
			"778682": "quadlek-chat", // @jirwin

		},
	))

	//coinbasePlugin := cointip.Register(
	//	c.String("coinbase-key"),
	//	c.String("coinbase-secret"),
	//	c.String("coinbase-account"),
	//)
	//if coinbasePlugin != nil {
	//	err = bot.RegisterPlugin(coinbasePlugin)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	err = bot.RegisterPlugin(comics.Register(
		c.String("imgur-client-id"),
		"/opt/quad-assets/Arial.ttf",
	))
	if err != nil {
		fmt.Printf("error registering comic plugin: %s\n", err.Error())
		return err
	}
	//esPlugin, err := eslogs.Register(c.String("es-endpoint"), c.String("es-index"))
	//if err != nil {
	//	fmt.Printf("Error creating eslogs plugin: %s\n", err.Error())
	//	return err
	//}
	//err = bot.RegisterPlugin(esPlugin)
	//if err != nil {
	//	fmt.Printf("Error registering eslogs plugin: %s\n", err.Error())
	//	return err
	//}

	xpostPlugin := xpost.Register()
	err = bot.RegisterPlugin(xpostPlugin)
	if err != nil {
		fmt.Printf("error registering xpost plugin: %s\n", err.Error())
		return err
	}

	stoxPlugin := stox.Register(c.String("av-key"))
	err = bot.RegisterPlugin(stoxPlugin)
	if err != nil {
		fmt.Printf("Error registering stox plugin: %s\n", err.Error())
		return err
	}

	gifPlugin := gifs.Register(c.String("giphy-api-key"))
	err = bot.RegisterPlugin(gifPlugin)
	if err != nil {
		fmt.Printf("Error registering gifs plugin: %s\n", err.Error())
		return err
	}

	// Twitch not working
	//twitchFollows := []*twitch.TwitchFollow{
	//	{
	//		SlackChannels: []string{"kitboga"},
	//		TwitchUser:    "kitboga",
	//		WatchStream:   true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "telewyn",
	//		SlackUser:     "telewyn",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "khryo72",
	//		SlackUser:     "morgabra",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "mekilek",
	//		SlackUser:     "jirwin",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "sonicdm",
	//		SlackUser:     "sonicdm",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "sakatana",
	//		SlackUser:     "saka",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//	{
	//		SlackChannels: []string{"vidyagames"},
	//		TwitchUser:    "pountaar",
	//		SlackUser:     "purdyk",
	//		WatchStream:   true,
	//		WatchFollows:  true,
	//	},
	//}

	//twitchPlugin := twitch.Register(c.String("twitch-oauth-client-id"), "", "https://bullpeen-quadlek.quadlek.dev/slack/plugin/twitch", false, twitchFollows)
	//if twitchPlugin != nil {
	//	err = bot.RegisterPlugin(twitchPlugin)
	//	if err != nil {
	//		fmt.Printf("error registering twitch plugin: %s\n", err.Error())
	//		return err
	//	}
	//}
	//
	//err = bot.RegisterPlugin(slices.Register([]string{
	//	"76561197980107683", // sonicdm
	//	"76561197976367183", // morgabra
	//	"76561198057633471", // greenjeans
	//	"76561198002272597", // newsomr
	//	"76561197974723967", // purdyk
	//	"1967806609958425",  // bh
	//	"1964439323566513",  // schonstal
	//}))
	//if err != nil {
	//	fmt.Printf("error registering slices plugin: %s\n", err.Error())
	//	return nil
	//}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	bot.Start()
	<-signals
	bot.Stop()

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "quadlek"
	app.Version = Version
	app.Usage = "a slack bot"
	app.Action = run
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key",
			Usage:  "The slack api token for the bot",
			EnvVar: "QUADLEK_API_TOKEN",
		},
		cli.StringFlag{
			Name:   "verification-token",
			Usage:  "The slack webhook verification token.",
			EnvVar: "QUADLEK_VERIFICATION_TOKEN",
		},
		cli.StringFlag{
			Name:   "db-path",
			Usage:  "The path where the database is stored.",
			Value:  "quadlek.db",
			EnvVar: "QUADLEK_DB_PATH",
		},
		cli.StringFlag{
			Name:   "tvdb-key",
			Usage:  "The TVDB api key for the bot, used by the nextep command",
			EnvVar: "QUADLEK_TVDB_TOKEN",
		},
		cli.StringFlag{
			Name:   "twitter-consumer-key",
			Usage:  "The consumer key for the twitter api",
			EnvVar: "QUADLEK_TWITTER_CONSUMER_KEY",
		},
		cli.StringFlag{
			Name:   "twitter-consumer-secret",
			Usage:  "The consumer secret for the twitter api",
			EnvVar: "QUADLEK_TWITTER_CONSUMER_SECRET",
		},
		cli.StringFlag{
			Name:   "twitter-access-token",
			Usage:  "The access key for the twitter api",
			EnvVar: "QUADLEK_TWITTER_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "twitter-access-secret",
			Usage:  "The access secret for the twitter api",
			EnvVar: "QUADLEK_TWITTER_ACCESS_SECRET",
		},
		cli.StringFlag{
			Name:   "coinbase-key",
			Usage:  "The access key for the coinbase api",
			EnvVar: "QUADLEK_COINBASE_KEY",
		},
		cli.StringFlag{
			Name:   "coinbase-secret",
			Usage:  "The access secret for the coinbase api",
			EnvVar: "QUADLEK_COINBASE_SECRET",
		},
		cli.StringFlag{
			Name:   "coinbase-account",
			Usage:  "The bank account for the coinbase api",
			EnvVar: "QUADLEK_COINBASE_BANK_ACCOUNT",
		},
		cli.StringFlag{
			Name:   "imgur-client-id",
			Usage:  "The imgur client id",
			EnvVar: "QUADLEK_IMGUR_CLIENT_ID",
		},
		cli.StringFlag{
			Name:   "es-endpoint",
			Usage:  "The elastic search endpoint",
			EnvVar: "QUADLEK_ES_ENDPOINT",
		},
		cli.StringFlag{
			Name:   "es-index",
			Usage:  "The elastic search index",
			EnvVar: "QUADLEK_ES_INDEX",
		},
		cli.StringFlag{
			Name:   "av-key",
			Usage:  "Alpha Vantage Api Key",
			EnvVar: "QUADLEK_AV_KEY",
		},
		cli.StringFlag{
			Name:   "giphy-api-key",
			Usage:  "Giphy API Key",
			EnvVar: "QUADLEK_GIPHY_KEY",
		},
		cli.StringFlag{
			Name:   "twitch-oauth-client-id",
			Usage:  "OAuth application client id for twitch API",
			EnvVar: "QUADLEK_TWITCH_OAUTH_CLIENT_ID",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
