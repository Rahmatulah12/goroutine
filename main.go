package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

/*
	Channel secara default adalah un-buffered, tidak di-buffer di memori.
	Ketika ada goroutine yang mengirimkan data lewat channel,
	harus ada goroutine lain yang bertugas menerima data dari channel yang sama,
	dengan proses serah-terima yang bersifat blocking.
	Maksudnya, baris kode di bagian pengiriman dan penerimaan data,
	tidak akan akan diproses sebelum proses serah-terima-nya selesai.
*/

func main() {
	// runtime.GOMAXPROCS(n) menentukan jumlah prosesor yang aktif, sebanyak n
	// runtime.GOMAXPROCS(2)

	// // implementasi channel sebagai tipe data parameter
	// var messages = make(chan string)

	// var names = []string{"Sakura", "Sasuke", "Kakashi", "Naruto"}

	// for _, each := range names {
	// 	// memanggil function
	// 	go func(who string) {
	// 		// menangkap data who dan simpan ke var data
	// 		var data = fmt.Sprintf("hello %s", who)
	// 		messages <- data
	// 	}(each)
	// }

	// for i := 0; i < len(names); i++ {
	// 	printMessage(messages)
	// }

	// pembuatan GOROUTINE, dilakukan asynchronous
	// go print(5, "Hello World")

	// // dilakukan synchronous
	// print(5, "Apa kabar?")

	// var input string

	// /*
	// 	fmt.Scanln(&input)
	// 	memberhentikan proses aplikasi, sampai user menekan tombol enter
	// 	Hal ini perlu dilakukan karena ada kemungkinan waktu selesainya eksekusi
	// 	goroutine print() lebih lama dibanding waktu selesainya goroutine utama main() ,
	// 	mengingat bahwa keduanya sama-sama asnychronous.
	// */
	// fmt.Scanln(&input)

	// var input1, input2, input3 string

	// // fmt.Scanln(&input1) => Fungsi ini akan meng-capture semua karakter sebelum user menekan tombol enter,
	// // lalu menyimpannya pada variabel.
	// fmt.Scanln(&input1)
	// fmt.Scanln(&input2)
	// fmt.Scanln(&input3)

	// fmt.Println(input1)
	// fmt.Println(input2)
	// fmt.Println(input3)

	// /*
	// 	Channel, merupakan sebuah variabel, dibuat dengan menggunakan keyword make dan chan
	// */
	// // pembuatan channel
	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	var data = fmt.Sprintf("Hello %s", who)
	// 	messages <- data
	// }

	// go sayHelloTo("John Wick")
	// go sayHelloTo("Bertand Antolin")
	// go sayHelloTo("James Bond")

	// var messages1 = <-messages
	// fmt.Println(messages1)
	// var messages2 = <-messages
	// fmt.Println(messages2)
	// var messages3 = <-messages
	// fmt.Println(messages3)

	/*
		Penerapan Buffered Channel
	*/
	runtime.GOMAXPROCS(2)
	// messages := make(chan int, 2) // memiliki jumlah buffer 3 (0, 1, 2)

	// // melakukan penerimaan data dengan goroutine
	// go func() {
	// 	for {
	// 		i := <-messages
	// 		fmt.Println("Receive data", i)
	// 	}
	// }()

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("send data", i)
	// 	messages <- i
	// }
	// // fmt.Scanln()

	// var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	// fmt.Println("Numbers :", numbers)

	// // channel1 untuk menangkap data rata2x
	// channel1 := make(chan float64)
	// go getAverage(numbers, channel1)

	// // channel2 untuk menangkap data yang paling besar
	// channel2 := make(chan int)
	// go getMax(numbers, channel2)

	// for j := 0; j < 2; j++ {
	// 	select {
	// 	case avg := <-channel1:
	// 		fmt.Printf("Average \t: %.2f \n", avg)
	// 	case max := <-channel2:
	// 		fmt.Printf("Max \t: %d \n", max)
	// 	}
	// }

	var message2 = make(chan string)
	go sendMessage(message2)
	receiveAndPrintMessage(message2)

	rand.Seed(time.Now().Unix())
	var messages5 = make(chan int)

	go sendData(messages5)
	retrieveData(messages5)

	/*
		Defer, digunakan untuk mengakhirkan eksekusi sebuah statement.
		Ketika eksekusi sudah sampai pada akhir blok fungsi,
		statement yang di defer baru akan dijalankan.
		Ketika ada banyak statement yang di-defer, maka statement
		tersebut akan dieksekusi di akhir secara berurutan.
	*/

	defer fmt.Println("Hello ini dijalankan diakhir baris program.")

	/*
		Exit, digunakan untuk menghentikan program.
		os.Exit(), berada dalam package os.  Fungsi ini memiliki
		sebuah parameter bertipe numerik yang wajib diisi.
		Angka yang dimasukkan akan muncul sebagai exit status
	*/
	os.Exit(1)
	fmt.Println("selamat datang")
}

/*
	Channel sebagai tipe data parameter
*/
func printMessage(what chan string) {
	fmt.Println(<-what)
}

func getAverage(numbers []int, channel chan float64) {
	var sum = 0
	for _, e := range numbers {
		sum += e
	}
	channel <- float64(sum) / float64(len(numbers))
}

// channel chan int dapat menerima dan mengirim data
func getMax(numbers []int, channel chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	channel <- max
}

//  ch chan<- hanya bisa mengirim data
func sendMessage(ch chan<- string) {
	// data akan dikirimkan sebnyak perulangan
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	// close channel
	close(ch)
}

// ch <-chan hanya bisa menerima data
func receiveAndPrintMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

/*
	Channel Timeout
	Timeout digunakan untuk mengontrol penerimaan data dari channel berdasarkan waktu diterimanya,
	dengan durasi timeout bisa ditentukan sendiri.
	Ketika tidak ada aktivitas penerimaan data selama durasi tersebut,
	akan memicu callback yang isinya juga ditentukan sendiri.
*/
func sendData(ch chan<- int) {
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print("Receive data ", data, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("Timeout. No Activity under 5 second.")
			break loop
		}
	}
}
