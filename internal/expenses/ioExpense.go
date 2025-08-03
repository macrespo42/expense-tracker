package expenses

import (
	"encoding/json"
	"os"
	"time"
)

const expensePath = "./expenses.json"

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

func LoadExpenses() ([]Expense, error) {
	var expenses []Expense

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

	return os.WriteFile(expensePath, data, 0644)
}
