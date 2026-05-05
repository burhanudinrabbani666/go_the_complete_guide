package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type CompanyData struct {
	Revenue  float64
	Expenses float64
	TaxRate  float64
}

type ResultCalculateFinancial struct {
	Ebt    float64
	Profit float64
	Ratio  float64
}

func main() {
	CompanyData := CompanyData{}

	CompanyData.Revenue = getUserInput("Revenue: ")
	CompanyData.Expenses = getUserInput("Expense: ")
	CompanyData.TaxRate = getUserInput("Tax Rate: ")

	fmt.Print("\n")

	result := calculateFinancial(CompanyData)
	fmt.Printf("EBT: %.2f\n", result.Ebt)
	fmt.Printf("Profit: %.2f\n", result.Profit)
	fmt.Printf("Ratio: %.2f\n", result.Ratio)

	bytes, _ := json.Marshal(result)
	resultString := string(bytes)

	os.WriteFile("resultCalculateFinancial.json", []byte(resultString), 0666)
}

func getUserInput(question string) (userInput float64) {
	fmt.Print(question)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		panic("Input Not Valid. input should be number and positive")
	}

	return userInput
}

func calculateFinancial(data CompanyData) (ResultCalculateFinancial ResultCalculateFinancial) {
	ResultCalculateFinancial.Ebt = data.Revenue - data.Expenses
	ResultCalculateFinancial.Profit = ResultCalculateFinancial.Ebt * (1 - (data.TaxRate / 100))
	ResultCalculateFinancial.Ratio = ResultCalculateFinancial.Ebt / ResultCalculateFinancial.Profit

	return ResultCalculateFinancial
}
