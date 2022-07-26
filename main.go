package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"swapper/package/globals"
	"swapper/package/twitter"
	"swapper/package/utilities"
	"time"
)

var (
	updateCounter      bool = true
	rareAccountNewName string
)

func main() {
	rand.Seed(time.Now().Unix())

	utilities.ClearScreen()

	config := utilities.ReadConfig("./config.json")

	utilities.PrintLogo()
	
	receivingUsername, err := twitter.VerifyAccount(config.Accounts.AccountToRecieveUsername)
	if err == nil && receivingUsername != "" {
		utilities.Debug("Account to recieve username is verified | @" + receivingUsername, '+')
	} else {
		utilities.Debug("Could not verify account, quitting...", '-')
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

	rareUsername, err := twitter.VerifyAccount(config.Accounts.AccountWithRareUsername)
	if err == nil && rareUsername != "" {
		utilities.Debug("Account with rare username is verified | @" + rareUsername, '+')
	} else {
		utilities.Debug("Could not verify account, quitting...", '-')
		time.Sleep(1 * time.Second)
		os.Exit(1)
	}

	if twitter.CheckRatelimit(config.Accounts.AccountToRecieveUsername) {
		utilities.Debug("@" + receivingUsername + " is being ratelimited, try again later for a faster swap", '-')
	} else {
		utilities.Debug("@" + receivingUsername + " is not being ratelimited", '+')
	}

	if twitter.CheckRatelimit(config.Accounts.AccountWithRareUsername) {
		utilities.Debug("@" + rareUsername + " is being ratelimited, try again later for a faster swap", '-')
	} else {
		utilities.Debug("@" + rareUsername + " is not being ratelimited", '+')
	}
	
	fmt.Println()
	utilities.Debug("Press [ENTER] to release username...", '*')
	fmt.Scanln()
	utilities.ClearScreen()

	go utilities.UpdateOnScreenCounters(&updateCounter)
	twitter.Swap(rareUsername, config.Accounts.AccountToRecieveUsername, config.Accounts.AccountWithRareUsername, &rareAccountNewName)

	updateCounter = false
	utilities.ClearScreen()
	utilities.PrintLogo()

	utilities.Debug("Successfully swapped " + rareUsername + " => " + rareAccountNewName + ", " + receivingUsername + " => " + rareUsername + " in " + strconv.Itoa(globals.TimeElapsed) + "ms", '+')
	fmt.Println()
	utilities.Debug("Press [ENTER] to exit swapper...", '*')
	fmt.Scanln()
}
