package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"timezone-modifier/internal"
)

func main() {
	fmt.Println("Timezone modifier service started. Enter commands (help for list):")
	fmt.Println("Enter help for list")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		switch command {
		case "list":
			internal.ListTimezones()
		case "set":
			if len(args) < 2 {
				fmt.Println("Error: timezone not specified")
				printHelp()
				continue
			}
			internal.SetTimezone(args[1])
		case "current":
			internal.GetCurrentTimezone()
		case "exit":
			os.Exit(0)
		case "help":
			printHelp()
		default:
			fmt.Printf("Unknown command: %s\n", command)
			printHelp()
		}

		fmt.Print("\n> ") // Приглашение для следующей команды
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Commands:")
	fmt.Println("  list                - List all available timezones")
	fmt.Println("  set <timezone>      - Set the system timezone")
	fmt.Println("  current             - Show current timezone")
	fmt.Println("  help                - Show this help message")
}
