package day12

func Assignment2() {
	runAssignment1("day12/input.txt", func(record []rune, sizes []int) ([]rune, []int) {

		next_record := make([]rune, 0)
		next_sizes := make([]int, 0)

		repeat := 5

		for i := 0; i < repeat; i++ {
			next_record = append(next_record, record...)
			if i != repeat-1 {
				next_record = append(next_record, '?')
			}

			next_sizes = append(next_sizes, sizes...)
		}

		return next_record, next_sizes
	})
}
