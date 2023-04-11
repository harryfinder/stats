package stats_test

import (
	"fmt"
	"github.com/harryfinder/bank/pkg/bank/types"
	"github.com/harryfinder/stats"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1234,
			Amount:   2_000_00,
			Category: "restaurnt",
		},
		{
			ID:       1111,
			Amount:   3_000_00,
			Category: "chemist's",
		},
		{
			ID:       7000,
			Amount:   7_000_00,
			Category: "auto",
		},
	}
	fmt.Println(stats.Avg(payments))

	//output:
	//400000
}
func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       1234,
			Amount:   2_000_00,
			Category: "restaurnt",
		},
		{
			ID:       1111,
			Amount:   3_000_00,
			Category: "auto",
		},
		{
			ID:       7000,
			Amount:   7_000_00,
			Category: "auto",
		},
	}

	summa := stats.TotalInCategory(payments, "auto")
	fmt.Println(summa)

	//output:
	//1000000
}
