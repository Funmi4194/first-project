package main

import "fmt"

func status(name string) {
	grades := map[string]int{"Alma": 0, "Robert": 66, "Funmi": 78, "Precious": 46}
	if grade, ok := grades[name]; !ok {
		fmt.Printf("%q has not been assigned a score yet\n", name)
	} else if grade < 60{
		fmt.Printf("%q with %v marks has failed, Kindly retake course!\n", name, grades[name])
	}else{
		fmt.Printf("%q has passed with %v marks\n", name, grades[name])
	}
}

func main(){
	status("Alma")
	status("Funmi")
	status("Damilola")
	status("Precious")
	status("Robert")

}

//// A code to find a given target number
// package main
// import "fmt"
// func getTarget(number[]int, target int)[]int{
// 	var result = []int{}

// 	for i:=0; i<len(number); i++{
// 		for j:= i+1; j<len(number); j++{
// 			sum := number[i] +number[j]

// 			if sum == target{
// 				result = append(result, number[i])
// 				result = append(result, number[j])
// 				return result
// 			}
// 		}
// 	}
// 	return result
// }

// func main(){
// 	fmt.Println(getTarget([]int{2, 7, 6, 8, 5}, 9))

// }