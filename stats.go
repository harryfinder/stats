package stats

import "github.com/harryfinder/bank/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	var length int
	var sum int64
	for i, operations := range payments {
		sum += int64(operations.Amount)
		length += i
	}

	average := types.Money(sum / int64(length))

	return average
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, operations := range payments {
		if category == operations.Category {
			sum += operations.Amount
		}
	}
	return sum
}
