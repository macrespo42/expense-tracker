/*
Copyright © 2025 NAME HERE <maxime.crespo@protonmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/macrespo42/expense-tracker/internal/expenses"
	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" || amount == 0 {
			fmt.Println("Please specify --description and --amount flags")
			return
		}

		expensesList, err := expenses.LoadExpenses()
		if err != nil {
			fmt.Println(err)
			log.Fatal("Can't load expenses file")
		}

		id := 0
		if len(expensesList) > 0 {
			id = expensesList[len(expensesList)-1].ID + 1
		}

		newExpense := expenses.Expense{
			ID:          id,
			Description: description,
			Amount:      amount,
			Date:        time.Now(),
		}

		expensesList = append(expensesList, newExpense)
		err = expenses.SaveExpenses(expensesList)
		if err != nil {
			log.Fatal("Can't save new expense")
		}
	},
}

func init() {
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the expense")
	addCmd.Flags().Float64VarP(&amount, "amount", "a", 0.0, "Amount of the expense")
	rootCmd.AddCommand(addCmd)
}
