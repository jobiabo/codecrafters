package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Welcome to Cli Verse Calculator")

	fmt.Println("Choose Math Operation")
	fmt.Println("Choose option 1-5")
	fmt.Println("You can Type Quit to exit the Calculator")

OPERATIONERROR:
	fmt.Println("[1] Addition")
	fmt.Println("[2] Subtraction")
	fmt.Println("[3] Multiplication")
	fmt.Println("[4] Division")
	fmt.Println("Help")
	fmt.Println("Quit")
	fmt.Println("NOTE: You will have to enter each number at a time")
	fmt.Print("Enter [1]-[5] or Help here: ")

	var user_INputOperator string
	fmt.Scan(&user_INputOperator)

	if user_INputOperator == "1" {

	firstadd:

		fmt.Println("Enter single number at a time to add(eg. 2 enter key 9)")
		fmt.Print("Enter first number: ")
		// fmt.Println(a + b)
		var storeUserIntI string
		fmt.Scan(&storeUserIntI)

		fmt.Print("Enter second number: ")
		var storeUserIntII string
		fmt.Scan(&storeUserIntII)

		numI, err := strconv.ParseFloat(storeUserIntI, 64)
		if err == nil {

			// storeUserIntI = numI
		} else {
			fmt.Println("math error!")
			fmt.Println("Enter a valid integer")
			goto firstadd
		}

		numII, err := strconv.ParseFloat(storeUserIntII, 64)
		if err == nil {
			// storeUserIntII = numII
		} else {
			fmt.Println("math error!")
			fmt.Println("Enter a valid int")
			goto firstadd
		}
		var toTalSum = numI + numII
		fmt.Println("result", toTalSum)
		fmt.Println("would you love to do anything else?")
		fmt.Println("[1] yes [2] no")

		var Recons string
		fmt.Scan(&Recons)

		if Recons == "1" {
			fmt.Println("Welcome back to Cli verse calculator")
			main()
		} else if Recons == "2" {
			fmt.Println("would you love to quit now?")
			fmt.Println("[1] yes [2] no")

			var quitOpt string
			fmt.Scan(&quitOpt)

			if quitOpt == "1" {
				fmt.Println("Good bye")
				return
			} else {
				main()
			}

		}

	}

	if user_INputOperator == "2" {

	Subtraction:

		fmt.Println("Enter single number at a time to subtract(eg. 2 enter key 9)")
		fmt.Print("Enter first number: ")

		// fmt.Println(a + b)
		var storeUserIntI string
		fmt.Scan(&storeUserIntI)

		fmt.Print("Enter second number: ")
		var storeUserIntII string
		fmt.Scan(&storeUserIntII)

		numI, err := strconv.ParseFloat(storeUserIntI, 64)
		if err == nil {

			// storeUserIntI = numI
		} else {
			fmt.Println("math error!")
			fmt.Println("Enter a valid integer")
			goto Subtraction
		}

		numII, err := strconv.ParseFloat(storeUserIntII, 64)
		if err == nil {
			// storeUserIntII = numII
		} else {
			fmt.Println("math error!")
			fmt.Println("Enter a valid int")
			goto Subtraction
		}
		var toTalSum = numI - numII
		fmt.Println("result", toTalSum)
		fmt.Println("would you love to do anything else?")
		fmt.Println("[1] yes [2] no")
		var Recons string
		fmt.Scan(&Recons)

		if Recons == "1" {
			fmt.Println("Welcome back to Cli verse calculator")
			main()
		} else if Recons == "2" {
			fmt.Println("would you love to quit now?")
			fmt.Println("[1] yes [2] no")
		}

		var quitOpt string
		fmt.Scan(&quitOpt)
		if quitOpt == "1" {
			fmt.Println("Good bye")
			return
		} else {
			main()
		}

	}

	if user_INputOperator == "3" {

	StepMultiplication:

		fmt.Println("Enter single number at a time to multiply(eg. 2 enter key 9):")
		fmt.Print("Enter first number: ")

		// fmt.Println(a + b)
		var storeUserIntI string
		fmt.Scan(&storeUserIntI)

		fmt.Print("Enter second number: ")
		var storeUserIntII string
		fmt.Scan(&storeUserIntII)

		numI, err := strconv.ParseFloat(storeUserIntI, 64)
		if err == nil {

			// storeUserIntI = numI
		} else {
			fmt.Println("math error!")
			fmt.Print("Enter a valid integer: ")
			goto StepMultiplication
		}

		numII, err := strconv.ParseFloat(storeUserIntII, 64)
		if err == nil {
			// storeUserIntII = numII
		} else {
			fmt.Println("math error!")
			fmt.Print("Enter a valid int:")
			goto StepMultiplication
		}
		var toTalSum = numI * numII
		fmt.Println("result", toTalSum)

		fmt.Println("would you love to do anything else?")
		fmt.Println("[1] yes [2] no")
		var Recons string
		fmt.Scan(&Recons)

		if Recons == "1" {
			fmt.Println("Welcome back to Cli verse calculator")
			main()
		} else if Recons == "2" {
			fmt.Println("would you love to quit now?")
			fmt.Println("[1] yes [2] no")
		}

		var quitOpt string
		fmt.Scan(&quitOpt)
		if quitOpt == "1" {
			fmt.Println("Good bye")
			return
		} else {
			main()
		}

	}

	if user_INputOperator == "4" {

	StepDivision:

		fmt.Println("Enter single number at a time to divide(eg. 2 enter key 9):")
		fmt.Print("Enter first number:")
		// fmt.Println(a + b)
		var storeUserIntI string
		fmt.Scan(&storeUserIntI)

		fmt.Print("Enter second number:")
		var storeUserIntII string
		fmt.Scan(&storeUserIntII)

		numI, err := strconv.ParseFloat(storeUserIntI, 64)
		if err == nil {

			// storeUserIntI = numI
		} else {
			fmt.Println("math error!")
			fmt.Print("Enter a valid integer: ")
			goto StepDivision
		}

		numII, err := strconv.ParseFloat(storeUserIntII, 64)
		if err == nil {
			// storeUserIntII = numII
		} else {
			fmt.Println("math error!")
			fmt.Print("Enter a valid int: ")
			goto StepDivision
		}

		if numII < 1 {
			fmt.Println("ERROR! math undefined")

			goto StepDivision
		}

		var toTalSum = numI / numII
		fmt.Println("result", toTalSum)

		fmt.Println("would you love to do anything else?")
		fmt.Println("[1] yes [2] no")
		var Recons string
		fmt.Scan(&Recons)

		if Recons == "1" {
			fmt.Println("Welcome back to Cli verse calculator")
			main()
		} else if Recons == "2" {
			fmt.Println("would you love to quit now?")
			fmt.Println("[1] yes [2] no")
		}

		var quitOpt string
		fmt.Scan(&quitOpt)
		if quitOpt == "1" {
			fmt.Println("Good bye")
			return
		} else {
			main()
		}

	}

	if user_INputOperator == "help" || user_INputOperator == "Help" {
		fmt.Println("This page contains userguide for this calculator")
		fmt.Println("case> Add e.g 2+2")
		fmt.Println("case> Subtract e.g 3-1")
		fmt.Println("case> Multiply e.g 2*1")
		fmt.Println("case> Divide e.g 2/2")
		fmt.Println("NOTE: Every math operation interger should be separated by a math operator e.g 2+2")
		fmt.Println("would you love to do anything else?")
		fmt.Println("[1] yes [2] no")
		var Recons string
		fmt.Scan(&Recons)

		if Recons == "1" {
			fmt.Println("Welcome back to Cli verse calculator")
			main()
		} else if Recons == "2" {
			fmt.Println("would you love to quit now?")
			fmt.Println("[1] yes [2] no")
		}

		var quitOpt string
		fmt.Scan(&quitOpt)
		if quitOpt == "1" {
			fmt.Println("Goodbye see you soon")
			return
		} else {
			main()
		}

	}

	if user_INputOperator == "Quit" || user_INputOperator == "QUIT" || user_INputOperator == "quit" {
		fmt.Println("would you love to quit now?")
		fmt.Println("[1] yes [2] no")

	quitPart:
		var quitOpt string
		fmt.Scan(&quitOpt)
		if quitOpt == "1" {
			fmt.Println("Goodbye, see you soon")
			return
		} else {
			if quitOpt == "2" {
				goto OPERATIONERROR
			} else {
				goto quitPart

			}
		}

	}

	if user_INputOperator != "1" || user_INputOperator != "2" || user_INputOperator != "3" || user_INputOperator != "4" || user_INputOperator != "Help" || user_INputOperator != "help" || user_INputOperator != "HELP" {
		// fmt.Println("INVALID INPUT OPERATION")
		fmt.Println("Please choose from the available options [1]-[5] & Help what you want to do next")
		goto OPERATIONERROR

	}
}
