package main

import (
	"context"
	"fmt"
	"log"
	//"net/http"
	"os"

	        cli "github.com/urfave/cli/v2"
	lib "github.com/chickenandpork/pagerduty-oncall/lib"
)

const (
	ScheduleID = "schedule_id"
	AuthToken = "auth_token"
)

func main() {
        app := &cli.App{
                Flags: []cli.Flag{
                        &cli.StringFlag{Name: ScheduleID, Aliases: []string{"s"}, Usage: "Schedule ID, ie \"\"", EnvVars: []string{"PD_SCHEDULE_ID"}},
                        &cli.StringFlag{Name: AuthToken, Aliases: []string{"t"}, Usage: "Auth Token, ie \"\"", EnvVars: []string{"PD_AUTHTOKEN"}},
                },

                Name:  "pdoc",
                Usage: "check the current oncall",
                Action: func(c *cli.Context) error {
			if ocs, err := lib.ListCurrentOncall(context.Background(), c.String(AuthToken), c.String(ScheduleID)); err == nil {
				for _, v := range ocs {
					fmt.Printf("oncall: %s\n", v)
				}
			} else {
				return err
			}
                        return nil
                },
	}

        err := app.Run(os.Args)
        if err != nil {
                log.Fatal(err)
        }
}
