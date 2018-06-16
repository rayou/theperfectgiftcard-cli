package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	theperfectgiftcard "github.com/rayou/go-theperfectgiftcard"
	"github.com/rayou/theperfectgiftcard-cli/card"
	"github.com/urfave/cli"
)

func main() {
	var json bool
	var c card.Card
	app := cli.NewApp()
	app.Name = "The Perfect Gift Card cli"
	app.Usage = "Display card summary and transaction history"
	app.Version = "1.0.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ray Ou",
			Email: "yuhung.ou@live.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "cardno, n",
			Usage: "Gift card number",
		},
		cli.StringFlag{
			Name:  "pin, p",
			Usage: "Gift card pin",
		},
		cli.BoolFlag{
			Name:        "json",
			Usage:       "Display as JSON",
			Destination: &json,
		},
	}

	app.Action = cli.ActionFunc(func(ctx *cli.Context) error {
		if ctx.NumFlags() == 0 {
			cli.ShowAppHelp(ctx)
			return nil
		}

		if !ctx.IsSet("cardno") {
			return errors.New("Gift card number is required")
		}

		if !ctx.IsSet("pin") {
			return errors.New("Gift card pin is required")
		}

		client, _ := theperfectgiftcard.NewClient()
		if json {
			c = card.NewCardJSON(client, os.Stdout)
		} else {
			tableWriter := tablewriter.NewWriter(os.Stdout)
			c = card.NewCardTable(client, os.Stdout, tableWriter)
		}
		c.SetCardNo(ctx.String("cardno"))
		c.SetPin(ctx.String("pin"))
		return c.Print()
	})

	err := app.Run(os.Args)
	if err != nil {
		if c != nil {
			c.PrintErr(err)
		} else {
			fmt.Printf("\033[31mERR\033[0m: %v\n", err)
		}
	}
}
