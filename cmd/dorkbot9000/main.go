package main

import (
	"fmt"
	"github.com/stevehnh/dorkbot9000/internal/commands"
	"github.com/stevehnh/dorkbot9000/internal/config"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
)

// CommandMsg is a message from a user to the bot
type CommandMsg struct {
	Command  string
	Argument []string
}

var validCommands = []string{
	"!swear",
	"!quote",
}

func findInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Exec(command *CommandMsg) (string, error) {
	switch command.Command {
	case "!swear":
		// return the help dialogue for the !swear command
		if len(command.Argument) <= 0 || command.Argument[0] == "help" {
			return commands.HelpSwear(), nil
		}
		response, err := commands.Swear(command.Argument)
		if err != nil {
			return "", err
		}
		return response, nil
		
	case "!quote":
		return "I'm not a bot, I'm a human", nil
	default:
		return "", fmt.Errorf("Invalid command")
	}
}

func ReadMessages(conf *config.BotConfig, client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		v := strings.Fields(message.Message)
		_, found := findInSlice(validCommands, v[0])
		if !found {
			client.Say(conf.Channel, "Invalid command: "+v[0])
		} else {
			command := &CommandMsg{
				Command:  v[0],
				Argument: v[1:],
			}
			response, err := Exec(command)
			if err != nil {
				client.Say(conf.Channel, "Error executing command")
				fmt.Printf("Error executing command: %v \n", err)
			} else {
				client.Say(conf.Channel, response)
			}
		}
	})
}

func main() {

	// Load configuration
	conf := config.Load()

	// Connect to Twitch IRC
	client := twitch.NewClient(conf.Username, conf.OAuth)

	client.OnConnect(func() {
		fmt.Println("Connected to Twitch IRC")
	})

	client.Join(conf.Channel)

	ReadMessages(conf, client)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
