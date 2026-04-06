package main

import (
	"fmt"
	"helper/helper"
	"strings"
)

func main() {

	fmt.Println("In this Session You're using a CLI converter for (HEX), and (BIN)")
MainMenu:
	fmt.Println("This converts user inputs from (Hex)/(Bin) to a decimal figure")
	fmt.Println("MAIN MENU")
	fmt.Println("Choose either")
	fmt.Println("[1] for (Hex) to decimal")
	fmt.Println("[2] for (BIN) to decimal")
	fmt.Println("[3] for Dec to (Hex)")
	fmt.Println("[4] for Dec to (Bin)")
	fmt.Println("---->NOTE: Enter Help for user guide and support")
	fmt.Println("---->NOTE: Enter Quit to exit the app")

	var UserIput_operation string
	fmt.Scan(&UserIput_operation)

	if UserIput_operation == "1" {

		fmt.Println("You will have to enter the (Hex) value")
		fmt.Print("Enter (Hex) you want to convert: ")

		var userFindvalue string
		fmt.Scan(&userFindvalue)

		// fmt.Print(userFindvalue, " is:", "in HEX")

		fmt.Println(helper.HexConv(userFindvalue), "is the result.")

		fmt.Println(("What would you like to do next?"))
		fmt.Println("Enter M for Homepage/Main Menu, Enter Q to Quit the App")
		var OptCare string
		fmt.Scan(&OptCare)

		if OptCare == "M" || OptCare == "m" {
			fmt.Println("Ok, You're Back on Homepage")
			main()

		}

		if OptCare == "Q" || OptCare == "q" {
			fmt.Println("Logging out session")
			fmt.Println("Goodbye")
			return
		}

	}

	if UserIput_operation == "2" {
		fmt.Println("You will have to enter the (Bin) value")
		fmt.Print("Enter (Bin) you want to convert: ")

		var userFindvalue string
		fmt.Scan(&userFindvalue)

		fmt.Println(helper.BinConv(userFindvalue), "is the result.")
		fmt.Println(("What would you like to do next?"))
		fmt.Println("Enter M for Homepage/Main Menu, Enter Q to Quit the App")
		var OptCare string
		fmt.Scan(&OptCare)

		if OptCare == "M" || OptCare == "m" {
			fmt.Println("Ok, You're Back on Homepage")
			main()

		}

		if OptCare == "Q" || OptCare == "q" {
			fmt.Println("Logging out session")
			fmt.Println("Goodbye")
			return
		}

	}
	if UserIput_operation == "3" {
		fmt.Println("Enter Decimal value to convert to (Hex)")
		fmt.Print("Enter Value Here: ")

		var userFindvalue string
		fmt.Scan(&userFindvalue)

		fmt.Println(helper.DecToHex(userFindvalue), "is the result.")

		fmt.Println(("What would you like to do next?"))
		fmt.Println("Enter M for Homepage/Main Menu, Enter Q to Quit the App")
		var OptCare string
		fmt.Scan(&OptCare)

		if OptCare == "M" || OptCare == "m" {
			fmt.Println("Ok, You're Back on Homepage")
			main()

		}

		if OptCare == "Q" || OptCare == "q" || OptCare == "quit" || OptCare == "QUIT" || OptCare == "Quit" {
			fmt.Println("Logging out session")
			fmt.Println("Goodbye")
			return
		} else {
			fmt.Println("Please Choose of the available options")
			goto MainMenu
		}

	}

	if UserIput_operation == "4" {
		fmt.Println("Enter Decimal value to convert to (BIN)")
		fmt.Print("Enter Value Here: ")

		var userFindvalue string
		fmt.Scan(&userFindvalue)

		fmt.Println(helper.DecToBin(userFindvalue), "is the result.")

		fmt.Println(("What would you like to do next?"))
		fmt.Println("Enter M for Homepage/Main Menu, Enter Q to Quit the App")
		var OptCare string
		fmt.Scan(&OptCare)

		if OptCare == "M" || OptCare == "m" {
			fmt.Println("Ok, You're Back on Homepage")
			main()

		}

		if OptCare == "Q" || OptCare == "q" || OptCare == "quit" || OptCare == "QUIT" || OptCare == "Quit" {
			fmt.Println("Logging out session")
			fmt.Println("Goodbye")
			return
		} else {
			fmt.Println("Please Choose of the available options")
			goto MainMenu
		}

	}

	if UserIput_operation == "help" || UserIput_operation == "HELP" || UserIput_operation == "Help" {
		fmt.Println("This is the help page for the users and support")
		fmt.Println(("What would you like to do next?"))
		fmt.Println("Enter M for Homepage/Main Menu, Enter Q to Quit the App")
		var OptCare string
		fmt.Scan(&OptCare)

		if OptCare == "M" || OptCare == "m" {
			fmt.Println("Ok, You're Back on Homepage")
			main()

		}

		if OptCare == "Q" || OptCare == "q" {
			fmt.Println("Logging out session")
			fmt.Println("Goodbye")
			return
		}

	}

	if UserIput_operation == "QUIT" || UserIput_operation == "quit" || UserIput_operation == "Quit" || UserIput_operation == "q" || UserIput_operation == "Q" {
		fmt.Println("Are you sure you want to quit?")
		fmt.Println("Choose Option 1 or 2")
		fmt.Println("1. Yes or 2. No ")

		var optOption string
		fmt.Scan(&optOption)
		lowOptOption := strings.ToLower(optOption)

		if lowOptOption == "yes" || lowOptOption == "1" {
			fmt.Println("Goodbye")
			return
		} else {
			if lowOptOption == "no" || lowOptOption == "2" {
				goto MainMenu
			}
		}
	}

	fmt.Printf("Please Input a valid option here: ")
	goto MainMenu

	// var UserIput_operation string
	// fmt.Print("Enter option here (enter 1 0r 2): ")
	// fmt.Scan(&UserIput_operation)

	// fmt.Println(CLiConv(UserIput_operation))
}
