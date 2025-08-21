package main

import "fmt"

func main() {
	var num int64 = 5

	settedSecondToOne := setToOne(num, 1)
	fmt.Printf("setted 2 bit to one; bits of num = %b, num = %v, bits of result = %b, result = %v\n",
		num, num,
		settedSecondToOne, settedSecondToOne,
	)

	settedFirstToZero := setToZero(num, 0)
	fmt.Printf("setted 1 bit to zero; bits of num = %b, num = %v, bits of result = %b, result = %v\n",
		num, num,
		settedFirstToZero, settedFirstToZero,
	)
}

func setToOne(num int64, i int64) int64 {
	return num | (1 << i)
}

func setToZero(num int64, i int64) int64 {
	return num &^ (1 << i)
}
