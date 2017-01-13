package main

import (
	"fmt"
	"strconv"
	"strings"
)

//ParseForCommands parses input for Commands, returns message if no command specified, else return is empty
func ParseForCommands(line string) string {
	//One Key Commands
	switch line {
	case ":g":
		SelectGuild()
		line = ""
	case ":c":
		SelectChannel()
		line = ""
	case ":p":
		SelectPrivate()
		line = ""
	case ":a":
		AddUserChannel()
		line = ""
	case ":d":
		Msg(TextMsg, "Not yet fully implemented.\nResults may vary.\n")
		SelectDeletePrivate()
		line = ""
	default:
		// Nothing
	}

	//Argument Commands
	if strings.HasPrefix(line, ":m") {
		AmountStr := strings.Split(line, " ")
		if len(AmountStr) < 2 {
			Msg(ErrorMsg, "[:m] No Arguments \n")
			return ""
		}

		Amount, err := strconv.Atoi(AmountStr[1])
		if err != nil {
			Msg(ErrorMsg, "[:m] Argument Error: %s \n", err)
			return ""
		}

		Msg(InfoMsg, "Printing last %d messages!\n", Amount)
		State.RetrieveMessages(Amount)
		PrintMessages(Amount)
		line = ""
	}

	return line
}

//SelectGuild selects a new Guild
func SelectGuild() {
	State.Enabled = false
	SelectGuildMenu()
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}

//AddUserChannel moves a user to a private channel with another user.
func AddUserChannel() {
	State.Enabled = false
	AddUserChannelMenu()
	State.Enabled = true
	ShowContent()
}

//AddUserChannelMenu takes a user from the current guild and adds them to a private message. WILL RETURN ERROR IF IN USER CHANNEL.
func AddUserChannelMenu() {
	if State.Channel.IsPrivate {
		Msg(ErrorMsg, "Currently in a user channel, move to a guild with :g\n")
	} else {
		SelectMap := make(map[int]string)
	Start:
		SelectID := 0
		for _, Member := range State.Members {
			SelectMap[SelectID] = Member.User.ID
			Msg(TextMsg, "[%d] %s\n", SelectID, Member.User.Username)
			SelectID++
		}
		var response string
		fmt.Scanf("%s\n", &response)

		if response == "b" {
			return
		}

		ResponseInteger, err := strconv.Atoi(response)
		if err != nil {
			Msg(ErrorMsg, "(CH) Conversion Error: %s\n", err)
			goto Start
		}

		if ResponseInteger > SelectID-1 || ResponseInteger < 0 {
			Msg(ErrorMsg, "(CH) Error: ID is out of bound\n")
			goto Start
		}
		Chan, err := Session.DiscordGo.UserChannelCreate(SelectMap[ResponseInteger])
		State.Channel = Chan
	}
}

//SelectChannel selects a new Channel
func SelectChannel() {
	State.Enabled = false
	SelectChannelMenu()
	State.Enabled = true
	ShowContent()
}

//SelectPrivate a private channel
func SelectPrivate() {
	State.Enabled = false
	SelectPrivateMenu()
	State.Enabled = true
	ShowContent()
}

//SelectDeletePrivate a private channel
func SelectDeletePrivate() {
	State.Enabled = false
	SelectDeletePrivateMenu()
	State.Enabled = true
	ShowContent()
}
