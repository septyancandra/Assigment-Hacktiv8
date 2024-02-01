var input = ['a', 'b', 'a', 1, 1, 2, 1, '1'];

function count(arr) {
    var n = arr.length;
    var freq = [];

    for (var i = 0; i < n; i++) {
        freq.push(0);
    }
    for (var i = 0; i < n; i++) {
        for (var j = 0; j < n; j++) {
            if (arr[i] === arr[j]) {
                freq[i]++;

            }
        }
    }
    var maks = 0;
    var maksdimana = 0;
    for (i = 0; i < n; i++) {
        if (freq[i] > maks) {
            maksdimana = i;
            maks = freq[i];
        }
    }
    // return (input[maksdimana],maks);
    console.log(arr[maksdimana], maks);
}

count(input);