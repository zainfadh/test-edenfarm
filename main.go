package main

import (
	"errors"
	"fmt"
)

func main() {
	var arr1 = []int{1, 2, 3, 8, 9, 3, 2, 1}
	var arr2 = []int{1, 2, 1, 2, 2, 1}
	var arr3 = []int{7, 1, 2, 9, 7, 2, 1}

	value1, err := checkHighest(arr1)
	fmt.Println(value1, err)

	value2, err := checkHighest(arr2)
	fmt.Println(value2, err)

	value3, err := checkHighest(arr3)
	fmt.Println(value3, err)
}

func checkHighest(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array kosong")
	}
	deret, err := findDeret(arr)
	if err != nil {
		return 0, err
	}
	fmt.Println(deret)
	deretReverse, err := findDeretReverse(arr)
	if err != nil {
		return 0, err
	}
	fmt.Println(deretReverse)
	if len(deret) == len(deretReverse) && deret[len(deret)-1] == deret[len(deretReverse)-1] {
		return deret[len(deret)-1], nil
	} else {
		return 0, errors.New("tidak ada deret angka yang sesuai")
	}

}

func findDeret(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return []int{}, errors.New("array kosong")
	}
	deret := []int{}
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 {
			if arr[i+1]-arr[i] == 1 {
				deret = append(deret, arr[i])
			}
			if len(deret) != 0 && arr[i+1]-arr[i] != 1 {
				deret = append(deret, arr[i])
				break
			}
		}
	}
	if len(deret) == 0 {
		return []int{}, errors.New("tidak ada angka yang berurutan dalam array")
	}
	return deret, nil
}

func findDeretReverse(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return []int{}, errors.New("array kosong")
	}
	deret := []int{}
	for i := len(arr) - 1; i > 0; i-- {
		if i != 1 {
			if arr[i-1]-arr[i] == 1 {
				deret = append(deret, arr[i])
			}
			if len(deret) != 0 && arr[i-1]-arr[i] != 1 {
				deret = append(deret, arr[i])
				break
			}
		}
	}
	if len(deret) == 0 {
		return []int{}, errors.New("tidak ada angka yang berurutan dalam array")
	}
	return deret, nil
}
