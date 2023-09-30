//****************

// package main

// import "fmt"

// func main() {
// 	var des[60] string
// 	des[0] =
// 	des[49]
// }


// ************************
// package main

// import "fmt"

// func main() {

// 	var des = [5] string {"gg1", "gg2", "gg3", "gg4", "gg5"}

// 	fmt.Println(des[0])
// 	fmt.Println(des[1])
// 	fmt.Println(des[2])
// 	fmt.Println(des[3])
// 	fmt.Println(des[4])
// }

// **************************
// package main

// import "fmt"

// func main() {
// 	var des = [5] string {"gg1", "gg2", "gg3", "gg4", "gg5"}

// 	for i := 0; i < 5; i++ {
// 		fmt.Println(des[i])
// 	}
// }


// ************************************
// package main

// import "fmt"

// func main() {
// 	var des = [5] string {"gg1", "gg2", "gg3", "gg4", "gg5"}
	
// 	var choice int
// 	fmt.Println("choice 1-4")
// 	fmt.Scan(&choice)

// 	fmt.Println(des[choice])

// }

// **************************************************
// package main

// import "fmt"

// func main() {
// 	var student [5] string 
	
// 	for i := 0; i < 5; i++ {
// 		fmt.Print("enter name: ")
// 		fmt.Scan(&student[i])	
// 	}

// 	fmt.Println(student)
// 	fmt.Println(len(student))

// }

// ***********************slice*******************

package main

import "fmt"

func main() {
	var student [] string //slice name 'student'
	var studentName string

	var NumofStudent int
	fmt.Printf("how many entry? (1- *): ")
	fmt.Scan(&NumofStudent)

	for i := 0; i < NumofStudent; i++ {
		fmt.Print("enter name: ")
		fmt.Scan(&studentName)
		student = append(student, studentName)	
	}

	fmt.Println(student)
	fmt.Println(len(student))

}









