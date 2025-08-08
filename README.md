# [Expense-tracker](https://roadmap.sh/projects/expense-tracker)

Expense tracker as a CLI

## Requirements

go installed

## installation

To install or run locally, clone the repo and make the script executable:

```bash
git clone https://github.com/macrespo42/expense-tracker
cd expense-tracker
go run . <command> <arguments>

```

or to installed it globally

```bash
go install https://github.com/macrespo42/expense-tracker@latest
# if $GOPATH isn't in your .zshrc/.bashrc
export PATH=$PATH:$(go env GOPATH)/bin > ~/.zshrc
source ~/.zshrc
# Then run the program from anywhere
expense-tracker <command> <arguments>
```

## Commands

```bash
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expense-tracker list
# ID  Date       Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expense-tracker summary
# Total expenses: $30

$ expense-tracker delete --id 2
# Expense deleted successfully

$ expense-tracker summary
# Total expenses: $20

$ expense-tracker summary --month 8
# Total expenses for August: $20
```
