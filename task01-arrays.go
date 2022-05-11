package homework

func average(input [15]float32) (result float32) {

	var sum float32
	var num int
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			num++
		}
		sum += input[i]
	}
	return sum / float32(num)
}

// Second solution
/*
	for _, el := range input {
		result += el
	}
	tmp := float32(len(input))
	result = result / tmp
	return result
*/

/*
// For testing
func main() {
	// declaring an array of values
	array := [6]int{1, 2, 3, 4, 5, 6}

	// size of the array
	// n := 6

	// declaring a variable
	// to store the sum
	sum := 0

	// traversing through the
	// array using for loop
	for i := 0; i < len(array); i++ {

		// adding the values of
		// array to the variable sum
		sum += (array[i])
	}

	// declaring a variable
	// avg to find the average
	avg := (float32(sum)) / (float32(len(array)))

	// typecast all values to float
	// to get the correct result
	fmt.Println("Sum = ", sum, "\nAverage = ", avg)
}
*/
