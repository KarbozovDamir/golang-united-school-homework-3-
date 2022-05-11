package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	//Place your code here
	keys := make([]int, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		result = append(result, input[k])
	}
	return
}

/*
// Second solution
// values := []string{}
	// keys := []int{}
	// for k := range input {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)
	// for _, k := range keys {
	// 	values = append(values, input[k])
	// }
	// return values
*/

/*
// For check
func main() {
	// basket := map[int]string{2: "a", 0: "b", 1: "c"}
	basket := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println(sortMapValues(basket))
}
*/
