var Arr = [1, 7, 2, 8, 3, 4, 5, 0, 9];

for (let i = 0; i < Arr.length; i++) {
    for (j = 0; j < i; j++) {
        if (Arr[i] < Arr[j]) {
            x = Arr[i];
            Arr[i] = Arr[j];
            Arr[j] = x;
        }
    }
}

for (let i = 0; i < Arr.length; i++) {
    console.log(Arr[i]);
}



