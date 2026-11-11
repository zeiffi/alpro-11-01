package main

import "fmt"

func main() {
	var current, prev int
	var count int
	var sum int
	var max, min int
	var hasData bool

	for {
		_, err := fmt.Scan(&current)
		if err != nil {
			break
		}

		if hasData && current == prev {
			break
		}

		if !hasData {
			max = current
			min = current
			hasData = true
		} else {
			if current > max {
				max = current
			}
			if current < min {
				min = current
			}
		}

		sum += current
		count++
		prev = current
	}

	if count > 0 {
		avg := float64(sum) / float64(count)
		fmt.Printf("%d %d %.2f\n", max, min, avg)
	}
}