package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const minecraftPort = 25565

var (
	red   = "\033[1;31m"
	reset = "\033[0m"
)

func main() {
	printWelcomeMessage()

	startIP := getUserInput("Enter the starting IP address: ", "\033[1;34m")
	endIP := getUserInput("Enter the ending IP address: ", "\033[1;34m")

	startAddr := net.ParseIP(startIP)
	endAddr := net.ParseIP(endIP)

	if startAddr == nil || endAddr == nil {
		fmt.Println("Invalid IP address format.")
		return
	}

	fmt.Printf("\n%s Scanning IP Range %s to %s %s\n", red+strings.Repeat("-", 120)+reset, startIP, endIP, red+strings.Repeat("-", 120)+reset)

	for ip := startAddr; !ip.Equal(endAddr); incrementIP(ip) {
		ipStr := ip.String()

		if isPortOpen(ipStr, minecraftPort) {
			fmt.Printf("Minecraft port (25565) is OPEN on %s\n", ipStr)
		} else {
			fmt.Printf("Minecraft port (25565) is CLOSED on %s\n", ipStr)
		}
	}
}

func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func isPortOpen(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func printWelcomeMessage() {
	red := "\033[1;31m"
	reset := "\033[0m"

	fmt.Printf(`
%s-----------------------------------------------------------------
|	░██████╗░░█████╗░██████╗░██╗░░██╗███████╗██████╗░	|
|	██╔════╝░██╔══██╗██╔══██╗██║░░██║██╔════╝██╔══██╗	|
|	██║░░██╗░██║░░██║██████╔╝███████║█████╗░░██████╔╝	|
|	██║░░╚██╗██║░░██║██╔═══╝░██╔══██║██╔══╝░░██╔══██╗	|
|	╚██████╔╝╚█████╔╝██║░░░░░██║░░██║███████╗██║░░██║	|
|	░╚═════╝░░╚════╝░╚═╝░░░░░╚═╝░░╚═╝╚══════╝╚═╝░░╚═╝	|
-----------------------------------------------------------------
	`, red)

	fmt.Printf("\n%s >> Written by Xnrrrrrr %s\n", red, reset)
	fmt.Printf("\n%s %s %s\n\n", red+strings.Repeat("-", 120)+reset, "HAPPY HUNTING", red+strings.Repeat("-", 120)+reset)
}

func getUserInput(prompt string, color string) string {
	var userInput string
	fmt.Print(color, prompt, reset)
	fmt.Scanln(&userInput)
	return userInput
}
