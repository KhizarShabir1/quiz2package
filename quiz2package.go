package quiz2package

import (
	"fmt"
)

//32542
//1235
//9478 republic of korea

func PrintCovid19Cases() {

	var cases [3]int
	cases[0] = 1235
	cases[1] = 9478
	cases[2] = 32542
	var option int
	for true {
		fmt.Println("Please select an option:\n1 – Print Covid19 cases in Pakistan \n2 – Print Covid19 cases in SouthKorea\n3 – Print Covid19 cases in France\n4 – Print my personalized message to Coronavirus\n0 – Exit")

		option = -1
		// fmt.Println("Enter 1 for transaction : ")
		_, err:= fmt.Scan(&option)
		if err != nil{
		
		  }

		switch option {
		case 1:
			fmt.Printf("No of cases in Pakistan: %d \n", cases[0])
		case 2:
			fmt.Printf("No of cases in South Korea: %d \n", cases[1])
		case 3:
			fmt.Printf("No of cases in France :%d \n", cases[2])
		case 4:
			fmt.Println("Stay Home, Stay Safe")
		case 0:
			return 
		default:
			fmt.Println("please Enter again")
		}
	}
}
