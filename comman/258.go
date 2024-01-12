package main

func addDigits(num int) int {
	// var remainder int
	// var sum int

	// for num%10 != num{
	//     for num > 0{
	//         remainder=num%10
	//         sum+=remainder
	//         num=num/10
	//     }
	//     num=sum
	//     sum=0
	// }
	// return num

	if num < 10 {
		return num
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
