package utilities

import (
	"swapper/package/globals"
	"runtime"
	"os/exec"
	"strings"
	"time"
	"fmt"
	"os"
)

const (
	RED     = "\u001b[31m"
	GREEN   = "\u001b[32m"
	YELLOW  = "\u001b[33m"
	MAGENTA = "\u001b[35m"
	BLACK   = "\u001b[30m"
	RESET   = "\u001b[0m"
)

func UpdateOnScreenCounters(updateCounter *bool) {
	for *updateCounter {
		fmt.Printf("\rRequests: %d %s|%s Ratelimits: %d %s|%s Errors: %d %s|%s Elapsed (ms): %d", globals.Requests, MAGENTA, RESET, globals.Ratelimits, MAGENTA, RESET, globals.Errors, MAGENTA, RESET, globals.TimeElapsed)
		time.Sleep(1 * time.Millisecond)
	}
}

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
	} else {
		cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
	}
}

func PrintLogo() {
	fmt.Printf("%s\n\n", strings.ReplaceAll(strings.ReplaceAll(`
 
    ░██████╗░██╗░░░░░░░██╗░█████╗░██████╗░██████╗░███████╗██████╗░
    ██╔════╝░██║░░██╗░░██║██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
    ╚█████╗░░╚██╗████╗██╔╝███████║██████╔╝██████╔╝█████╗░░██████╔╝
    ░╚═══██╗░░████╔═████║░██╔══██║██╔═══╝░██╔═══╝░██╔══╝░░██╔══██╗
    ██████╔╝░░╚██╔╝░╚██╔╝░██║░░██║██║░░░░░██║░░░░░███████╗██║░░██║
    ╚═════╝░░░░╚═╝░░░╚═╝░░╚═╝░░╚═╝╚═╝░░░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝`, "░", BLACK + "░" + RESET), "█", MAGENTA + "█" + RESET))
}

func Debug(message string, debugLevel byte) {
	parsedMessage := parseMessage(message, debugLevel)

	fmt.Print(parsedMessage)
}

func parseMessage(message string, debugLevel byte) string {
	switch debugLevel {
	case '+':
		return fmt.Sprintf("%s[%s+%s]%s %s\n", MAGENTA, GREEN, MAGENTA, RESET, message)

	case '-':
		return fmt.Sprintf("%s[%s-%s]%s %s\n", MAGENTA, RED, MAGENTA, RESET, message)

	case '*':
		return fmt.Sprintf("%s[%s*%s]%s %s\n", MAGENTA, YELLOW, MAGENTA, RESET, message)

	default:
		return fmt.Sprintf("%s[%s-%s]%s INVALID DEBUG LEVEL\n", MAGENTA, RED, MAGENTA, RESET)
	}
}
