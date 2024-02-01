let str = 'nanan';

let newString = '';
for (let i = str.length - 1; i >= 0; i--) {
    newString += str[i];
}

if (str == newString) {
    console.log("Palindrome")
} else {
    console.log(newString)
}

