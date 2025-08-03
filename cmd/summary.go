/*
Copyright © 2025 NAME HERE <maxime.crespo@protonmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var month string

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Give a sum of all monthly expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := loadExpenses("expenses.json")
		if err != nil {
			log.Fatal("Can't load expenses file")
		}

		var total float64
		for _, expense := range expenses {
			if month == "" || month == expense.Date.Month().String() {
				total += expense.Amount
			}
		}

		fmt.Printf("Total expenses: %.0f$\n", total)

	},
}

func init() {
	summaryCmd.Flags().StringVarP(&month, "month", "m", "", "Summary of expense for the current month")
	rootCmd.AddCommand(summaryCmd)
}
