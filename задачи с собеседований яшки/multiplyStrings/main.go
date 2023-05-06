package multiplyStrings

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	nums := make([]byte, len(num1)+len(num2))
	for i := len(num1)-1; i >= 0; i-- {
		for j := len(num2)-1; j >= 0; j-- {
			nums[i+j+1] += (num1[i]-'0')*(num2[j]-'0')
			if nums[i+j+1] >= 10 {
				nums[i+j] += nums[i+j+1]/10
				nums[i+j+1] %= 10
			}
		}
	}

	if nums[0] == 0 {
		nums = nums[1:]
	}
	for i := range nums {
		nums[i] += '0'
	}

	return string(nums)
}
