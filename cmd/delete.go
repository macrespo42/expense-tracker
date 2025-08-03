/*
Copyright © 2025 NAME HERE <maxime.crespo@protonmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var id int

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := loadExpenses("expenses.json")
		if err != nil {
			log.Fatal("Can't load expenses file")
		}

		if id == 0 {
			fmt.Println("Please provide the id of the expense you want to delete with --id")
		}

		var expensesCleared []Expense
		for _, expense := range expenses {
			if id != expense.ID {
				expensesCleared = append(expensesCleared, expense)
			}
		}

		err = saveExpenses("expenses.json", expensesCleared)
		if err != nil {
			log.Fatal("Failed to remove expense")
		}
	},
}

func init() {
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "Id of the expense to delete")
	rootCmd.AddCommand(deleteCmd)
}
