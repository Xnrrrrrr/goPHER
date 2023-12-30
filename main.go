package main

import (
	"fmt"
	"net"
	"sort"
	"strings"
	"sync"
	"time"
)

const minecraftPort = 25565

var (
	red   = "\033[1;31m"
	green = "\033[1;32m"
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

	var wg sync.WaitGroup
	var ipAddresses []net.IP

	fmt.Printf("\n%s Scanning IP Range %s to %s %s\n", red+strings.Repeat("-", 120)+reset, startIP, endIP, red+strings.Repeat("-", 120)+reset)

	// Generate a slice of IP addresses to scan
	for ip := startAddr; lessThanOrEqual(ip, endAddr); incrementIP(ip) {
		ipCopy := make(net.IP, len(ip))
		copy(ipCopy, ip)
		ipAddresses = append(ipAddresses, ipCopy)
		if ip.Equal(endAddr) {
			break
		}
	}

	// Define a channel to communicate results
	results := make(chan string, len(ipAddresses))

	// Launch goroutines to scan IP addresses
	for _, ip := range ipAddresses {
		wg.Add(1)
		go func(ip net.IP) {
			defer wg.Done()

			ipStr := ip.String()
			if isPortOpen(ipStr, minecraftPort) {
				results <- fmt.Sprintf("%sMinecraft port (25565) is OPEN on %s%s", green, ipStr, reset)
			} else {
				results <- fmt.Sprintf("%sMinecraft port (25565) is CLOSED on %s%s", red, ipStr, reset)
			}
		}(ip)
	}

	// Close the results channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results into a slice
	var scanResults []string
	for result := range results {
		scanResults = append(scanResults, result)
	}

	// Sort results
	sort.Strings(scanResults)

	// Print sorted results
	for _, result := range scanResults {
		fmt.Println(result)
	}
}

// Function to check if IP1 is less than or equal to IP2
func lessThanOrEqual(ip1, ip2 net.IP) bool {
	return ip1.To16().String() <= ip2.To16().String()
}

// Function to increment IP address (both IPv4 and IPv6)
func incrementIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
		// Ensure that the IP address is in IPv4 format
		if len(ip) == net.IPv4len && ip.To4() == nil {
			copy(ip, net.IPv4zero)
		}
	}
}

// Function to check if a port is open on a given IP address
func isPortOpen(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// Function to print the welcome message
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

// Function to get user input
func getUserInput(prompt string, color string) string {
	var userInput string
	fmt.Print(color, prompt, reset)
	fmt.Scanln(&userInput)
	return userInput
}
