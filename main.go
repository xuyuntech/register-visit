package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
	"github.com/xuyuntech/register-visit/api"
	"github.com/xuyuntech/register-visit/manager"
	"github.com/xuyuntech/register-visit/settings"
	"github.com/xuyuntech/register-visit/version"
)

var flags = []cli.Flag{
	cli.BoolFlag{
		EnvVar: "DEBUG",
		Name:   "debug, D",
		Usage:  "start the server in debug mode",
	},
	cli.StringFlag{
		EnvVar: "LISTEN",
		Name:   "listen, l",
		Usage:  "http listen address",
	},
	cli.StringFlag{
		EnvVar: "DATABASE_DATASOURCE",
		Name:   "database-datasource, ds",
		Usage:  "DATABASE_DATASOURCE",
	},
}

func main() {
	app := cli.NewApp()
	app.Flags = flags
	app.Name = "虚云科技挂号后台"
	app.Version = version.Version.String()
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
		return nil
	}
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func action(c *cli.Context) error {
	settings.InitSettings(c)

	controllerManager, err := manager.NewManager(c)
	if err != nil {
		return err
	}
	server := &api.Api{
		Listen:  c.String("listen"),
		Manager: controllerManager,
	}
	return server.Run()
}
