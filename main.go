/*
 * @author Marko Skorupan
 * @version 1.0.0
 *@date 2025-11-21
 * @fileoverview " Valid order: equal or more nuts and washers than bolts."
 */

package main

import "fmt"

func main () {
	// Constants
	const BOLT_COST int = 10
	const NUT_COST int = 5
	const WASHER_COST int = 3

	//Variables
	var bolts int
	var nuts int
	var washers int

	//INPUT
	//Ask user for ammounts
fmt.Print("How many bolts would you like to purchase? ")
fmt.Scan(&bolts)

fmt.Print("How many nuts would you like to purchase? ")
fmt.Scan(&nuts)

fmt.Print("How many washers would you like to purchase? ")
fmt.Scan(&washers)

//PROCESS

totalCost:= (bolts * BOLT_COST) + (nuts * NUT_COST) + (washers * WASHER_COST)

fmt.Println("Number of bolts: ", bolts)
fmt.Println("Number of nuts: ", nuts)
fmt.Println("Number of washers: ", washers)

if nuts < bolts {
fmt.Println("Check the Order - not enough nuts for those bolts you purchased.")
}

if washers < bolts {
fmt.Println("Check the Order - not enough washers for those bolts you purchased.")
}

if nuts>= bolts && washers>= bolts {
	fmt.Println("Order is OK.")
}
//OUTPUT
//DISPLAY TOTAL COST
fmt.Println("Your total cost of the order is ",totalCost,"¢")
}
