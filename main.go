package main

func getTarget(number[]int, target int)[]int{
	var result := []int{}

	for i:=0; i<len(number); i++{
		for j:= i+1; j<len(number); j++{
			sum := number[i] +number[j]

			if sum == target{
				result = append(result, number[i])
				result = append(result, number[j])
				return result
			}
		}
	}
	return result
}

func main(){

}