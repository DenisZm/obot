/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math"
	"net"
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

		// Handler for /start command
		obot.Handle("/start", func(m telebot.Context) error {
			log.Printf("Received command: /start")
			welcomeMessage := fmt.Sprintf("👋 Welcome to Obot %s!\n\n", appVersion)
			welcomeMessage += "I'm a DevOps Helper Telegram Bot that can provide CIDR subnet calculations.\n\n"
			welcomeMessage += "📋 *Available Commands:*\n"
			welcomeMessage += "• `/subnet <IP>/<CIDR>` - Calculate subnet information\n\n"
			welcomeMessage += "Example: `/subnet 192.168.0.1/24`"
			return m.Send(welcomeMessage, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		})

		// Handler for /subnet command
		obot.Handle("/subnet", func(m telebot.Context) error {
			log.Printf("Received command: /subnet with args: %s", m.Message().Payload)

			// Get the CIDR input from the payload
			cidrInput := m.Message().Payload
			if cidrInput == "" {
				return m.Send("❌ Please provide an IP address with CIDR notation.\nExample: `/subnet 192.168.0.1/24`")
			}

			// Calculate subnet information
			result, err := calculateSubnetInfo(cidrInput)
			if err != nil {
				log.Printf("Error calculating subnet info: %v", err)
				return m.Send(fmt.Sprintf("❌ %v", err))
			}

			return m.Send(result, &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		})

		// Handler for hello command (for testing)
		obot.Handle("/hello", func(m telebot.Context) error {
			log.Printf("Received command: /hello")
			return m.Send(fmt.Sprintf("Hello, I'm Obot %s!", appVersion))
		})

		obot.Start()

	},
}

// calculateSubnetInfo calculates subnet information based on CIDR notation
func calculateSubnetInfo(cidrInput string) (string, error) {
	// Parse CIDR notation
	ip, ipNet, err := net.ParseCIDR(cidrInput)
	if err != nil {
		return "", fmt.Errorf("Invalid IP/CIDR format: %v", err)
	}

	// Calculate subnet mask in dotted decimal format
	mask := ipNet.Mask
	maskStr := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])

	// Calculate network address (already provided by ipNet.IP)
	network := ipNet.IP.String()

	// Calculate CIDR prefix length (number of bits in the mask)
	ones, bits := mask.Size()

	// Calculate number of hosts
	var usableHosts float64
	if ones <= 30 {
		// Standard case: 2^(32-prefix) - 2 (network and broadcast addresses)
		usableHosts = math.Pow(2, float64(bits-ones)) - 2
	} else if ones == 31 {
		// /31 special case: RFC 3021 allows for 2 usable hosts (no network/broadcast reservation)
		usableHosts = 2
	} else { // ones == 32
		// /32 case: Single host address
		usableHosts = 1
	}

	// Calculate first and last usable IP addresses
	var firstIP, lastIP string
	if ones <= 30 {
		// Convert network address to 4-byte representation
		ip4 := ipNet.IP.To4()
		if ip4 == nil {
			return "", fmt.Errorf("Invalid IP address format")
		}

		// First usable IP is network address + 1
		firstIPBytes := make(net.IP, len(ip4))
		copy(firstIPBytes, ip4)
		firstIPBytes[3]++
		firstIP = firstIPBytes.String()

		// Last usable IP is broadcast address - 1
		// Calculate broadcast by setting host bits to 1
		lastIPBytes := make(net.IP, len(ip4))
		copy(lastIPBytes, ip4)
		for i := 0; i < len(mask); i++ {
			lastIPBytes[i] |= ^mask[i]
		}
		lastIPBytes[3]-- // Subtract 1 from last byte to get last usable
		lastIP = lastIPBytes.String()
	} else if ones == 31 {
		// For /31 networks, both IPs are usable (RFC 3021)
		ip4 := ipNet.IP.To4()

		// First IP is the network address
		firstIP = ip4.String()

		// Last IP is the network address with last bit set to 1
		lastIPBytes := make(net.IP, len(ip4))
		copy(lastIPBytes, ip4)
		lastIPBytes[3] |= 1
		lastIP = lastIPBytes.String()
	} else { // ones == 32
		// For /32, there's only one IP address
		firstIP = ip.String()
		lastIP = ip.String()
	}

	// Format the response
	result := fmt.Sprintf("📊 *Subnet Information for %s*\n\n", cidrInput)
	result += fmt.Sprintf("*Subnet Address:* %s/%d\n", network, ones)
	result += fmt.Sprintf("*Subnet Mask:* %s\n", maskStr)
	result += fmt.Sprintf("*IP Range:* %s - %s\n", firstIP, lastIP)
	result += fmt.Sprintf("*Usable Hosts:* %.0f", usableHosts)

	return result, nil
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
