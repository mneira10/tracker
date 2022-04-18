package cmd

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log an event you want to track.",
	Long:  `This command logs an event and the time it was triggered.`,
	Run: func(cmd *cobra.Command, args []string) {
		event := args[0]
		appendToFile(event)
	},
	Args: cobra.ExactArgs(1),
}

func appendToFile(event string) {
	f, err := os.OpenFile(DATA_FILE,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	time := time.Now().Format(time.RFC3339)
	textToAppend := time + DELIMITER + event + "\n"

	if _, err := f.WriteString(textToAppend); err != nil {
		log.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
