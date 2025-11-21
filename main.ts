/**
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview "Valid order: equal or more nuts and washers than bolts."
 */

const BOLT_COST: number = 10;
const NUT_COST: number = 5;
const WASHER_COST: number = 3;

//INPUT

let bolts: number = Number(prompt("How many bolts would you like to purchase? "));
let nuts: number = Number(prompt("How many nuts would you like to purchase? "));
let washers: number = Number(prompt("How many washers would you like to purchase? "));

// Output input values
console.log("Number of bolts: " + bolts);
console.log("Number of nuts: " + nuts);
console.log("Number of washers: " + washers);

//Check Order
if (nuts < bolts) {
  console.log("Check the Order - not enough nuts for those bolts you purchased.")
}

if (washers < bolts) {
  console.log("Check the order - not enough washers for those bolts you purchased.")
}

if (nuts>= bolts && washers>= bolts) {
  console.log("Order is OK.")
}

//Calculate cost
let totalCost: number = (bolts * BOLT_COST) + (nuts * NUT_COST) + (washers * WASHER_COST)

//DISPLAY TOTAL COST
console.log("Your total cost of the order is " + totalCost +"¢")

console.log("\nDone.")
