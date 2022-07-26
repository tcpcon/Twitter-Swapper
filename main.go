package main

import (
	"swapper/package/utilities"
	"swapper/package/twitter"
	"math/rand"
	"strconv"
	"time"
	"fmt"
	"os"
)

var (
	updateCounter bool = true
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
	swapStatistics := twitter.Swap(rareUsername, config.Accounts.AccountToRecieveUsername, config.Accounts.AccountWithRareUsername)

	updateCounter = false
	utilities.ClearScreen()
	utilities.PrintLogo()

	utilities.Debug("Successfully swapped " + rareUsername + " => " + swapStatistics.RareAccountNewName + ", " + receivingUsername + " => " + swapStatistics.ReceivingAccountNewName + " in " + strconv.Itoa(swapStatistics.TimeElapsed) + "ms", '+')
	fmt.Println()
	utilities.Debug("Press [ENTER] to exit swapper...", '*')
	fmt.Scanln()
}
