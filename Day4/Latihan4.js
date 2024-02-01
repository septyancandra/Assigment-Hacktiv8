function penentu(nilai){
    if (typeof nilai !== 'number' || isNaN(nilai)) {
        console.log('invalid data')
    }
    else{
        if(nilai % 2 == 0){
            console.log('Genap')
        } 
        else{
            console.log('Ganjil')
        }
    }
}

penentu(2)
penentu(3)
penentu(20)
penentu(21)
penentu("bahs") //cek salah