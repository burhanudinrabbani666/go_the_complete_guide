package main

import (
	"fmt"
)

type CompanyData struct {
	revenue  float64
	expenses float64
	taxRate  float64
}

type ResultCalculateFinancial struct {
	ebt    float64
	profit float64
	ratio  float64
}

func main() {
	CompanyData := CompanyData{}

	CompanyData.revenue = getUserInput("Revenue: ")
	CompanyData.expenses = getUserInput("Expense: ")
	CompanyData.taxRate = getUserInput("Tax Rate: ")

	fmt.Print("\n")

	ResultCalculateFinancial := calculateFinancial(CompanyData)
	fmt.Printf("EBT: %.2f\n", ResultCalculateFinancial.ebt)
	fmt.Printf("Profit: %.2f\n", ResultCalculateFinancial.profit)
	fmt.Printf("Ratio: %.2f\n", ResultCalculateFinancial.ratio)
}

func getUserInput(question string) (userInput float64) {
	fmt.Print(question)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancial(data CompanyData) (ResultCalculateFinancial ResultCalculateFinancial) {
	ResultCalculateFinancial.ebt = data.revenue - data.expenses
	ResultCalculateFinancial.profit = ResultCalculateFinancial.ebt * (1 - (data.taxRate / 100))
	ResultCalculateFinancial.ratio = ResultCalculateFinancial.ebt / ResultCalculateFinancial.profit

	return ResultCalculateFinancial
}
