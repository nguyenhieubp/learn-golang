package main

import (
	"fmt"
)

type User struct {
    name string
}

type UserOop interface {
	getName() string
	setName() string
	setName2() string
}

func (user *User) getName() {
	fmt.Println(user.name)
}

func (user *User) setName(newName string) string{
	user.name = newName;
	return user.name
}

func (user User) setName2(newName string) string{
	user.name = newName;
	fmt.Println(user.name)
	return user.name
	
}

func main() {
   //SCANF
	// %s string 
	// %d int
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)


	fmt.Println("Hello", name, ", Age:", age)


	//MAKE
   testSlice := make([]int,3)
   testSlice2 := []int{}
   fmt.Println(testSlice)
   fmt.Println(testSlice2)
   
   testMap := make(map[string]int,3)
   testMap2 := map[string]int{}
   fmt.Println(testMap)
   fmt.Println(testMap2)

   //IF ELSE
   var value int = 10
   if value % 2 == 0 {
	fmt.Println("true")
    } else{
		fmt.Println("false")
   }

   //SWITCH CASE
   var ageUser int = 18
   switch ageUser  {
    	case 16:
	    	fmt.Println("low")
		case 18: 
			fmt.Println("as")
		default:
			fmt.Println("high")
   } 
   


   //STRUCT - INTERFACE
   var hieu User = User{name: "hieu"}
   hieu.getName()
   hieu.setName2("ok")
   hieu.getName()
   hieu.setName("hieu siuuu")
   hieu.getName()
   

    //MAP
    var mp map[string]int = map[string]int{
		"apple": 1,
		"banana": 2,
	}
        //get
        fmt.Println(mp["apple"]) 
	    
		//add
        mp["orange"] = 3
		
		//set
		mp["apple"] = 9

		//delete
		delete(mp,"banana")


	fmt.Println(mp)

	

    //SLICE
	//SLICE != ARRAY 
	// array --> numbers := [3]int{6,7,8}  OR numbers := [...]int{6,7,8}(Đây là mảng động)
	// slice --> numbers := []int{6,7,8}
	numbers := []int{2,3,4,5,6,7,8}
	
	
	//ADD FIRST
	numbers = append([]int{1}, numbers...)
	fmt.Println("add first",numbers)

	//ADD
	numbers = append(numbers, 1)
	fmt.Println("add last",numbers)
	
	//DELETE LAST
	numbers = numbers[:len(numbers)-1]
	fmt.Println("delete last",numbers)
	
	//DELETE FIRST
	numbers = numbers[1:]
	fmt.Println("delete first",numbers)
	
	//SPLICE
	numbers = append(numbers[:2], append([]int{6, 7}, numbers[4:]...)...)
	fmt.Println("splice",numbers)


	//ARRAY
	array := [5]int{1,2,3,4,5}
	//FUNCTION
	fmt.Println("Sum: ",sum(array))
	
	
	//FOR
	fmt.Println("For")
	for i:=0;  i< len(array) ; i++ {
		fmt.Println(i)
	}
	
	//FOR RANGE
	fmt.Println("range")
	for i,v := range numbers {
		fmt.Println(i,v)
	}
}



//FUNCTION
func sum(array [5]int) int{
    total := 0
	for _, v := range array {
	   total += v
	};
	return total
}