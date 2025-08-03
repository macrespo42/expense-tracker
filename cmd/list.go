/*
Copyright © 2025 NAME HERE <maxime.crespo@protonmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/macrespo42/expense-tracker/internal/expenses"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all expenses",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := expenses.LoadExpenses()
		if err != nil {
			log.Fatal("Can't load expenses file")
		}

		fmt.Printf("# %-3s %-10s %-20s %8s\n", "ID", "Date", "Description", "Amount")
		for _, expense := range expenses {
			fmt.Printf("  %-3d %-10s %-20s %7.0f$\n", expense.ID, expense.Date.Format("2006-01-02"), expense.Description, expense.Amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
