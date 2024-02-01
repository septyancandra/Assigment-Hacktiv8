function compareNumbers(firstNumber, secondNumebr) {
    if (firstNumber > secondNumebr) {
        return true
    } else if (firstNumber == secondNumebr) {
        return -1
    } else {
        return false
    }

}

console.log(compareNumbers(5, 8));
console.log(compareNumbers(5, 3));
console.log(compareNumbers(4, 4));
console.log(compareNumbers(3, 3));
console.log(compareNumbers(17, 2));