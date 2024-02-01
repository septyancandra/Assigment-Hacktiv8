function isPrime(number){
    if(number <= 1){
        return false;
    }
    for(i = 2; i<number; i++){
        if(number % i == 0){
            return false
        }
    }
    return true;
}
function listPrima(angka1, angka2){
    let temp= [];
    for(let i = angka1; i <= angka2; i++){
        if(isPrime(i)){
            temp.push(i);
        }
    }
    return temp;
}


console.log(listPrima(1,5));
console.log(listPrima(5,10));
console.log(listPrima(10,20));