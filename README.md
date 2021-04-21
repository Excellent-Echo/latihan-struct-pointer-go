**# latihan-struct-pointer-go**

soal latihan struct dan pointer



LATIHAN STRUCT DAN POINTER

latihan struct

Buatlah sebuah fungsi yang menerima parameter berupa data variadic, dan kembalian data berupa struct dengan properti antara lain "status" bertipe boolean , score bertipe slice dan menampung tipe data integer, statisctic bertipe data struct.



score diisi dengan data yang dimasukkan dari parameter variadic

statistic memiliki 3 properti:

average (int) didapat dari score rata-rata

minimum (int) didapat dari score terendah

maximum (int) didapat dari score tertinggi

status didapat jika statistic rata-rata melebihi 50 maka sisi dengan true, jika dibawah atau sama dengan 50 isi dengan false

Test case



GetStatistic(76,80,50,50,60,70)

GetStatistic(50,50,50)

GetStatistic(90,90,90,90)



// expected if result data in JSON



{

​    status : true, 

​    score : [76, 80, 50, 50, 60, 70], 

​    statistic : {

​        avereage : 64,

​        minimum : 50,

​        maximum : 80,

​    }

}



{

​    status : false, 

​    score : [50, 50, 50], 

​    statistic : {

​        avereage : 50,

​        minimum : 50,

​        maximum : 50,

​    }

}



{

​    status : true, 

​    score : [90, 90, 90, 90], 

​     statistic : {

​        avereage : 90,

​        minimum : 90,

​        maximum : 90,

​    }

}

latihan pointer

buatlah sebuah fungsi yang dapat menambah nilai dan menghitung rata - rata hasil dari data yang diinput ditambah dengan data yang sudah ada :



Test case



number = [80,80,90,90]



avg := GetAverage(number , 70,70)  // 80+80+90+90+70+70 / 6 = 80

avg2 := GetAverage(number, 60,60) // 80+80+90+90+70+70+60+60 / 8 = 75

avg3 := GetAverage(number, 50,40,80) // 80+80+90+90+70+70+60+60+50+40+80 / 11 = 70