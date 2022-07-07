package main

import (
	"fmt"
	"reflect"
)

// Не запуская, скачать что выведет код?
func main() {
	numbers := make([]*int, 0, 10)
	for _, value := range []int{1, 2, 3, 4} {
		numbers = append(numbers, &value)
	}

	multiply(&numbers)

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
}

func multiply(numbers *[]*int) {
	var sum int
	dic := make(map[*int]struct{})
	for _, n := range *numbers {
		if reflect.DeepEqual(dic[n], reflect.Struct) {
			*n = (*n) * 2
		}
		sum += *n
		dic[n] = struct{}{}
	}

	*numbers = append(*numbers, &sum)
}
