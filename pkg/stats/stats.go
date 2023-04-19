package stats

import (
	"github.com/harryfinder/bank/v3/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	var length int
	var sum int64
	for i, operations := range payments {
		if operations.Status != types.StatusFail {
			sum += int64(operations.Amount)
			length = i
		}
	}

	average := types.Money(sum / int64(length))

	return average
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, operations := range payments {
		if category == operations.Category && operations.Status != types.StatusFail {
			sum += operations.Amount
		}
	}
	return sum
}

func FindPaymentByID(payments []types.Payment, category types.Category) []types.Payment {
	var paymentsAny []types.Payment
	for _, operations := range payments {
		if category == operations.Category {
			paymentsAny = append(paymentsAny, operations)
		}
	}
	return paymentsAny
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := make(map[types.Category]types.Money)
	counts := make(map[types.Category]int)

	for _, payment := range payments {
		category := payment.Category
		categories[category] += payment.Amount
		counts[category]++
	}

	for category, total := range categories {
		count := counts[category]
		avg := total / types.Money(count)
		categories[category] = avg
	}

	return categories
}
