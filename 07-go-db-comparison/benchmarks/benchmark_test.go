package benchmarks

import (
	"fmt"
	"testing"
)

func Benchmark(b *testing.B) {
	setup()
	defer clear()

	// Benchmark goes in here
	limits := []int{
		1,
		10,
		100,
		1000,
		10000,
		15000,
	}

	for _, lim := range limits { // Fetch varying number of rows

		fmt.Printf("================================== BENCHMARKING %d RECORDS ======================================\n", lim)
		// Benchmark Database/sql
		b.Run(fmt.Sprintf("Database/sql limit:%d ", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DbSqlQueryStudentWithLimit(lim)
			}
		})

		// Benchmark Sqlx
		b.Run(fmt.Sprintf("Sqlx limit:%d ", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SqlxQueryStudentWithLimit(lim)
			}
		})

		// Benchmark Sqlc
		b.Run(fmt.Sprintf("Sqlc limit:%d ", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SqlcQueryStudentWithLimit(lim)
			}
		})

		// Benchmark GORM
		b.Run(fmt.Sprintf("GORM limit:%d ", lim), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GormQueryStudentWithLimit(lim)
			}
		})

		fmt.Println("=================================================================================================")
	}
}
