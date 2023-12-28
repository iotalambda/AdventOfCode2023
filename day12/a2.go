package day12

func Assignment2() {
	runAssignment1("day12/input.txt", func(record []rune, sizes []int) ([]rune, []int) {

		nextRecord := make([]rune, 0)
		nextSizes := make([]int, 0)

		repeat := 5

		for i := 0; i < repeat; i++ {
			nextRecord = append(nextRecord, record...)
			if i != repeat-1 {
				nextRecord = append(nextRecord, '?')
			}

			nextSizes = append(nextSizes, sizes...)
		}

		return nextRecord, nextSizes
	})
}
