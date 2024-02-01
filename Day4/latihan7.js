function urutHuruf (text) {
    return text.split('').sort().join('');
}

console.log(urutHuruf('hallo'))
console.log(urutHuruf('qwerty'))