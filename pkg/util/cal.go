package util

func addUpper(num int) int {
	res := 0
	for i := 1; i <= num; i++ {
		res += i
	}
	return res
}

func GetSub(n1, n2 int) int {
	res := n2 - n1 //计划是n1-n2 故意写错为n2-n1
	return res
}
