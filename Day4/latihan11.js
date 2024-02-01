function isPrime (numbers) {
    if (numbers <= 1) 
            return false; 
  
        // Check from 2 to n-1 
        for (let i = 2; i < numbers; i++) 
            if (numbers % i == 0) 
                return false; 
  
        return true; 
}

console.log(isPrime(3));
console.log(isPrime(7));
console.log(isPrime(6));
console.log(isPrime(23));
console.log(isPrime(33));