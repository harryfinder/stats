package stats_test

import (
	"fmt"
	"github.com/harryfinder/bank/v3/pkg/bank/types"
	"github.com/harryfinder/stats/pkg/stats"
	"reflect"
	"testing"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1234,
			Amount:   2_000_00,
			Category: "restaurnt",
			Status:   "FAIL",
		},
		{
			ID:       1111,
			Amount:   3_000_00,
			Category: "chemist's",
			Status:   "InProgress",
		},
		{
			ID:       7000,
			Amount:   7_000_00,
			Category: "auto",
			Status:   "OK",
		},
	}
	fmt.Println(stats.Avg(payments))

	//output:
	//500000
}
func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       1234,
			Amount:   2_000_00,
			Category: "restaurnt",
			Status:   "FAIL",
		},
		{
			ID:       1111,
			Amount:   3_000_00,
			Category: "auto",
			Status:   "InProgress",
		},
		{
			ID:       7000,
			Amount:   7_000_00,
			Category: "auto",
			Status:   "OK",
		},
	}

	summa := stats.TotalInCategory(payments, "auto")
	fmt.Println(summa)

	//output:
	//1000000
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	}
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := stats.TotalInCategory(payments, "auto")

	if reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{Amount: 1, Category: "auto"},
		{Amount: 2, Category: "food"},
		{Amount: 3, Category: "auto"},
		{Amount: 4, Category: "auto"},
		{Amount: 5, Category: "food"},
	}

	expected := map[types.Category]types.Money{
		"auto": 2,
		"food": 3,
	}

	result := stats.CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}
