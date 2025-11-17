/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-15
 * @fileoverview This program checks if a contestant is allowed in the pie eating contest
 * based on their weight.
 */

// Constants
const MIN_WEIGHT: number = 77.0;
const MAX_WEIGHT: number = 105.0;

// Input
const weight: number = parseFloat(prompt("How much do you weigh?") || "0");

// Decision
if (weight >= MIN_WEIGHT && weight <= MAX_WEIGHT) {
  console.log("You may enter the contest.");
} else {
  console.log("You may NOT enter the contest.");
}

// Done message
console.log("\nDone.");
