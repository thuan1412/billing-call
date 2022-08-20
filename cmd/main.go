package main

import (
	"calling-bill/helpers"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "seedDb",
	Short: "Create data for the `users` table",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := helpers.GetDb()
		helpers.PanicErr(err)
		SeedDb(client)
	},
}

func main() {
	err := godotenv.Load()
	helpers.PanicErr(err)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
