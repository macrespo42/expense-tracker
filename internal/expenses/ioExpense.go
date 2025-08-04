package expenses

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func getExpensePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Can't retrieve home dire")
	}
	return path.Join(homeDir, "expenses.json")
}

func LoadExpenses() ([]Expense, error) {
	var expenses []Expense
	expensePath := getExpensePath()
	if _, err := os.Stat(expensePath); os.IsNotExist(err) {
		err := os.WriteFile(expensePath, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	data, err := os.ReadFile(expensePath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func SaveExpenses(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(getExpensePath(), data, 0644)
}
