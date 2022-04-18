/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the last time an event was logged.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		event := args[0]
		getLastEvent(event)
	},
	Args: cobra.ExactArgs(1),
}

func getLastEvent(event string) {

	if _, err := os.Stat(DATA_FILE); err == nil {
		f, err := os.Open(DATA_FILE)

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		output := fmt.Sprintf(`No "%s" events found.`, event)
		for scanner.Scan() {
			lineDate, lineEvent := readLine(scanner.Text())
			if lineEvent == event {
				output = lineDate
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(output)

	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No events data available. Please record some events first.")

	} else {
		fmt.Println(err)

	}

}

func readLine(line string) (string, string) {
	splitText := strings.Split(line, DELIMITER)
	// date, event
	return splitText[0], splitText[1]
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
