/*
Copyright © 2025 NAME HERE <maxime.crespo@protonmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	description string
	amount      float64
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func loadExpenses(expenseFile string) ([]Expense, error) {
	var expenses []Expense

	if _, err := os.Stat(expenseFile); os.IsNotExist(err) {
		err := os.WriteFile(expenseFile, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(expenseFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func saveExpenses(expenseFile string, expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(expenseFile, data, 0644)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an expense",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if description == "" || amount == 0 {
			fmt.Println("Please specify --description and --amount flags")
			return
		}

		expenses, err := loadExpenses("expenses.json")
		if err != nil {
			fmt.Println(err)
			log.Fatal("Can't load expenses file")
		}

		id := 0
		if len(expenses) > 0 {
			id = expenses[len(expenses)-1].ID + 1
		}

		newExpense := Expense{
			ID:          id,
			Description: description,
			Amount:      amount,
			Date:        time.Now(),
		}

		expenses = append(expenses, newExpense)
		err = saveExpenses("expenses.json", expenses)
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
