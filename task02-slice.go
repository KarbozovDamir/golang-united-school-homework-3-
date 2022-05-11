package homework

func reverse(input []int64) (result []int64) {
	//Place your code here
	result = make([]int64, 0, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])

	}
	return result

}

/*
// Second solution
// for i := len(input) - 1; i >= len(input)/2; i-- {
// 	tmp := len(input) - 1 - i
// 	input[i], input[tmp] = input[tmp], input[i]
// }
// return
*/

/*
// For check
// func main() {
// 	a := []int64{1, 2, 5, 15}
// 	fmt.Println(a, reverse(a))
// }
*/
