package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type PersonalFinanceCategory struct {
    Detailed string `json:"detailed"`
}

type Transaction struct {
    Date     string     `json:"date"`
    Name     string     `json:"name"`
    Amount   float64    `json:"amount"`
    Category []string   `json:"category"`
    PersonalFinanceCategory PersonalFinanceCategory `json:"personal_finance_category"`
}

type Response struct {
    LatestTransactions []Transaction `json:"latest_transactions"`
}

func main() {
    data, err := os.ReadFile("pretty.json")
    if err != nil {
        panic(err)
    }

    var resp Response
    if err := json.Unmarshal(data, &resp); err != nil {
        panic(err)
    }

    for _, t := range resp.LatestTransactions {
        if t.PersonalFinanceCategory.Detailed == "INCOME_SALARY" {
			fmt.Println(t.Amount)
            fmt.Println(t.Date)
		}
	}
}
