package filesops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteValueToFile(value float64, fileName string) {
	balanceText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(balanceText), 0666)

	if err != nil {
		panic(err)
	}

}

func GetFloatFromFile(fileName string) (float64, error) {
	bytes, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to Find file.")
	}

	valueText := string(bytes)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to Parse Storage Value.")
	}

	return value, nil
}
