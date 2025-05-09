/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v4"
)

var (
	// TeleToken is the token for the Telegram bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// obotCmd represents the obot command
var obotCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"obot"},
	Short:   "Start the DevOps Helper Telegram Bot",
	Long:    `The "start" command initializes and starts the DevOps Helper Telegram Bot.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("obot %s started\n", appVersion)

		obot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
		}

		obot.Handle("/hello", func(m telebot.Context) error {
			log.Printf("Received command: /hello")
			return m.Send(fmt.Sprintf("Hello, I'm Obot %s!", appVersion))
		})

		obot.Start()

	},
}

func init() {
	rootCmd.AddCommand(obotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// obotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// obotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
