package variables

import "fmt"

func Variables() {
	fmt.Println("Hello GoLang!")

	//Assigning Variables
	var var1 int                //just initialize
	var1 = 1                    //assign
	var var2 int = 2            //Initialize and assign
	var3 := 3                   //explicit assignment
	var name, age = "sapan", 28 //define and assign multiple variables together
	var (
		a = 1
		b = 2
		c = "string"
	)

	fmt.Println(var1, var2, var3, name, age, a, b, c)

}
