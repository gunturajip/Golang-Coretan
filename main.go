package main

import (
	"context"
	"errors"
	"fmt"
	"golang-coretan/helpers"
	"math"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	a1 = iota + 7
	a2
	a3
)

func main() {
	// chapter 1 - printing hello world
	fmt.Println("Hello world")

	// chapter 2 - variables
	var nama string = "guntur"
	fmt.Println("Nama kamu adalah", nama)
	umur := 20
	fmt.Printf("%d\n", umur)
	fmt.Printf("%d\n", 21)

	// chapter 3 - iota auto increment
	fmt.Println(a1, a2, a3)
	fmt.Println(umur == 17)

	// chapter 4 - if else conditions
	var currentYear = 2023
	age := currentYear - 2002
	if age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else if age == 17 {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu harusnya sudah membuat kartu sim")
	}

	// chapter 5 - switch case
	switch age {
	case 17:
		fmt.Println("Umur kamu 17 tahun")
	case 16:
		fmt.Println("Umur kamu 16 tahun")
	case 15:
		fmt.Println("Umur kamu 15 tahun")
	}

	switch {
	case (age > 17) && (umur >= 20):
		fmt.Println("Umur kamu gajelas")
		fallthrough
	case (age == 17) && (umur == 17):
		fmt.Println("Umur kamu terlalu muda")
	}

	// chapter 6 - for loop
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	for age < 25 {
		fmt.Println("Angka", age)
		age++
	}

	for {
		fmt.Println("Angka", age)
		age++
		if age == 26 {
			break
		}
	}

	for {
		if age%2 == 0 {
			age++
			continue
		}
		if age > 28 {
			age++
			break
		}
		fmt.Println("Angka", age)
		age++
	}

outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke -", i)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Println("Angka", j)
		}
	}

	// chapter 7 - array
	var list2 = [3]string{"guntur", "aji", "pratama"}
	fmt.Printf("%#v\n", list2)

	for _, v := range list2 {
		fmt.Printf("Value: %s\n", v)
	}
	for i := 0; i < len(list2); i++ {
		fmt.Printf("Value: %s\n", list2[i])
	}

	var list3 = [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Printf("%#v\n", list3)

	for i, v := range list3 {
		fmt.Printf("%d: %v\n", i, v)
		for _, w := range list3[0] {
			fmt.Printf("%d ", w)
		}
		fmt.Println()
	}

	// chapter 8 - slice
	var listz = []string{"a", "b", "c"}
	fmt.Printf("%#v\n", listz)

	var listzz = make([]string, 3)
	for i := 0; i < 3; i++ {
		listzz[i] = "melon"
	}
	listzz = append(listzz, "mango")
	fmt.Printf("%#v\n", listzz)

	listz = append(listz, listzz...)
	fmt.Printf("%#v\n", listz)

	newlist := copy(listz, listzz)
	fmt.Printf("%#v\n", newlist)
	fmt.Printf("%#v\n", listz)

	fmt.Printf("%v\n", listz[:3])

	listz = append(listz[:3], listz[2:]...)
	listz[2] = "durian"
	fmt.Printf("%#v\n", listz)

	// chapter 9 - backing array
	var fruits1 = []string{"durian", "mango", "banana"}
	var fruits2 = fruits1[1:3]
	fruits2[0] = "pisang"
	fmt.Printf("%#v\n", fruits1)
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = []string{"durian", "mango", "banana"}
	var fruits4 = []string{}
	fruits4 = append(fruits4, fruits1[1:3]...)
	fruits4[0] = "pisang"
	fmt.Printf("%#v\n", fruits3)
	fmt.Printf("%#v\n", fruits4)

	// chapter 10 - map
	var person map[string]string = map[string]string{}
	person["nama"] = "guntur"
	person["umur"] = "20"
	person["asal"] = "surabaya"
	for key, value := range person {
		fmt.Printf("%s : %s\n", key, value)
	}

	value, exist := person["nama"]
	if exist {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Println("value doesn't exist")
	}
	delete(person, "nama")
	value, exist = person["nama"]
	if exist {
		fmt.Printf("%s\n", value)
	} else {
		fmt.Println("value doesn't exist")
	}

	var persons = []map[string]int{
		{"nama": 10, "umur": 11},
		{"nama": 12, "umur": 13},
		{"nama": 14, "umur": 15},
	}
	for i, person := range persons {
		fmt.Printf("Index: %d, nama: %d, umur: %d\n", i, person["nama"], person["umur"])
	}

	// chapter 11 - aliases
	type hours = uint
	var hour hours = 12
	fmt.Printf("%T\n", hour)

	// chapter 12 - function
	fmt.Println(greet("guntur", "surabaya", 20, 50, "halo", "hola"))

	// chapter 13 - closure
	var evenNumbers = func(numbers ...int) []int {
		var result []int
		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}
	var numbers = []int{1, 2, 3, 4}
	fmt.Println(evenNumbers(numbers...))

	// chapter 14 - iife
	var oddNumbers = func(numbers ...int) []int {
		var result []int
		for _, v := range numbers {
			if v%2 != 0 {
				result = append(result, v)
			}
		}
		return result
	}([]int{1, 2, 3, 4}...)
	fmt.Println(oddNumbers)

	// chapter 15 - callback
	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println("Total odd numbers:", find)

	// chapter 16 - pointer
	var a int = 2
	var b *int = &a
	fmt.Println(a, &a)
	fmt.Println(*b, b)
	*b = 3
	fmt.Println(a, &a)
	fmt.Println(*b, b)

	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)

	// chapter 17 - struct
	var employee1 = Employee{}
	employee1.person.name = "guntur"
	employee1.person.age = 20
	employee1.division = "data scientist"
	fmt.Println("Employee name:", employee1.person.name)
	fmt.Println("Employee age:", employee1.person.age)
	fmt.Println("Employee division:", employee1.division)

	// chapter 18 - anonymous struct
	var employee2 = struct {
		person   Person
		division string
	}{}
	employee2.person.name = "aji"
	employee2.person.age = 20
	employee2.division = "web developer"
	fmt.Println("Employee name:", employee2.person.name)
	fmt.Println("Employee age:", employee2.person.age)
	fmt.Println("Employee division:", employee2.division)

	// chapter 19 - slice of struct
	var people = []Person{
		{name: "pratama", age: 20},
		{name: "raihan", age: 20},
		{name: "farhan", age: 20},
	}
	for _, v := range people {
		fmt.Println("Employee:", v)
	}

	// chapter 20 - method
	var person2 = Person{name: "guntur", age: 20}
	fmt.Println(person2.Introduce("Hello everyone"))
	person2.reviseName("Guntur")
	fmt.Println(person2.Introduce("Hello everyone"))

	// chapter 21 - reflect
	var angka = 20
	var reflectAngka = reflect.ValueOf(angka)
	fmt.Println("Tipe variabel angka:", reflectAngka.Type())
	if reflectAngka.Kind() == reflect.Int {
		fmt.Println("Nilai variabel angka:", reflectAngka.Int())
	}

	// chapter 22 - export package
	fmt.Println("Export dari iniMethod:", helpers.IniMethod("guntur"))

	// chapter 23 - interface
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 3, height: 5}
	fmt.Printf("Type of c1: %T", c1)
	fmt.Println("Circle area:", c1.area())
	fmt.Println("Circle perimeter:", c1.perimeter())
	fmt.Println("Circle volume:", c1.(circle).volume())
	fmt.Printf("Type of r1: %T", r1)
	fmt.Println("Rectangle area:", r1.area())
	fmt.Println("Rectangle perimeter:", r1.perimeter())

	// chapter 24 - empty interface
	var randomValue interface{}
	_ = randomValue
	randomValue = "Jalan Sudirman"
	randomValue = 20
	if value, ok := randomValue.(int); ok {
		randomValue = value * 2
	}
	randomValue = true
	randomValue = []string{"Airell", "Nanda"}
	// var randomSlice []interface{}
	// randomSlice = append(randomSlice, "halo")
	// randomSlice = append(randomSlice, 4)

	// challenge buat looping tek kotek
	first_str := "kotek kotek kotek\nAnak ayam turun berkotek\nTek kotek kotek kotek\nAnak ayam turun berkotek\n\n"

	last_str := "Anak ayam turunlah 1\nPergi 1 tinggal indukya"

	input := 22

	for i := input; i > 1; i-- {
		str := fmt.Sprintf("Anak ayam turunlah %d\nPergi 1 tinggallah %d\n\n", i, i-1)
		first_str += str
	}

	first_str += last_str

	fmt.Printf("%s\n", first_str)

	// chapter 25 - concurrency & goroutine
	fmt.Println("main execution started")
	go firstProcess(100)
	go secondProcess(100)
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())
	time.Sleep(100 * time.Millisecond)
	fmt.Println("main execution ended")

	// goroutine sync - waitgroup
	fruits := []string{"apple", "manggo", "durian", "rambutan"}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg, &mutex)
	}
	wg.Wait()

	// chapter 26 - channels
	c := make(chan string)
	students := []string{"Airell", "Nanda", "Mailo"}
	for _, student := range students {
		go introduce(student, c)
		cResult := <-c
		go greeting(cResult, c)
		fmt.Println(<-c)
	}
	close(c)

	// chapter 27 - unbuffered channel
	c2 := make(chan int, 3)
	go func(c2 chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c2 <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c2)
	}(c2)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(2 * time.Second)

	for v := range c2 { // v = <- c2
		fmt.Println("main goroutine received data:", v)
	}

	// chapter 28 - channel select
	c3 := make(chan string)
	c4 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "Hello!"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		c4 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c3:
			fmt.Println("Received", msg1)
		case msg2 := <-c4:
			fmt.Println("Received", msg2)
		}
	}

	// phase 2 - go programming microservices into the specialization - chapter 30 - fizz buzz
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scanln(&n)
	res := []string{}
	for i := 1; i <= n; i++ {
		out := ""
		if i%3 == 0 {
			out += "Fizz"
		}
		if i%5 == 0 {
			out += "Buzz"
		}
		if out == "" {
			res = append(res, fmt.Sprintf("%d", i))
		} else {
			res = append(res, out)
		}
	}
	fmt.Println(res)

	// chapter 31- defer - first int last out
	callDeferFunc()
	fmt.Println("hai everyone!!")

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	// chapter 32 - exit
	fmt.Println("sebelum exit")
	// os.Exit(1)
	fmt.Println("sesudah exit")

	// chapter 33 - error, panic, recover
	var number int
	var err error
	number, err = strconv.Atoi("123GH")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}
	number, err = strconv.Atoi("123")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(number)
	}

	// custom error
	var password string
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&password)
	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}

	// panic
	var pwd string
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&pwd)
	if valid, err := validPassword(pwd); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

	// recover
	defer catchErr()
	var pw string
	fmt.Print("Masukkan password: ")
	fmt.Scanln(&pw)
	if valid, err := validPassword(pw); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

	// chapter 34 - context withValue (immutable dan menganut konsep parent child)
	ctx1 := context.Background()
	fmt.Println(ctx1)
	contextA := context.WithValue(ctx1, "key_a", "value_a")
	fmt.Println(contextA)
	valA := contextA.Value("key-a")
	fmt.Println(valA)

	ctx2 := context.TODO()
	fmt.Println(ctx2)

	/////// context withCancel (mengirimkan sinyal pembatalan)

	// cek jumlah goroutine yg sedang berjalan
	fmt.Println("Total initiate goroutine", runtime.NumGoroutine())

	// inisiasi context baru
	ctx3 := context.Background()

	// membuat context dengan cancel
	ctxWithCancel, cancel := context.WithCancel(ctx3)

	// menjalankan goroutine
	destination1 := createCounter1(ctxWithCancel)

	// cek jumlah goroutine yg sedang berjalan setelah memanggil goroutine
	fmt.Println("Total goroutine after run", runtime.NumGoroutine())

	// print counter dari goroutine sampai 10
	for n := range destination1 {
		fmt.Println("counter", n)
		if n == 10 {
			break
		}
	}

	// mengirimkan sinyal cancel
	cancel()

	time.Sleep(3 * time.Second)

	// cek jumlah goroutine setelah dilakukan cancel
	fmt.Println("Total goroutine after cancel", runtime.NumGoroutine())

	/////// context withTimeout (memberikan timeout berupa lama waktu pengerjaan)

	// cek jumlah goroutine yg sedang berjalan
	fmt.Println("Total initiate goroutine", runtime.NumGoroutine())

	// inisiasi context baru
	ctx4 := context.Background()

	// membuat context dengan timeout 5 detik
	ctxWithTimeout, cancel := context.WithTimeout(ctx4, 5*time.Second)
	defer cancel()

	// menjalankan goroutine
	destination2 := createCounter2(ctxWithTimeout)

	// cek jumlah goroutine yg sedang berjalan setelah memanggil goroutine
	fmt.Println("Total goroutine after run", runtime.NumGoroutine())

	// print counter dari goroutine sampai timeout
	for n := range destination2 {
		fmt.Println("counter", n)
	}

	time.Sleep(3 * time.Second)

	// cek jumlah goroutine setelah dilakukan cancel
	fmt.Println("Total goroutine after cancel", runtime.NumGoroutine())

	/////// context withDead (memberikan deadline maksimal waktu pengerjaan)

	// cek jumlah goroutine yg sedang berjalan
	fmt.Println("Total initiate goroutine", runtime.NumGoroutine())

	// inisiasi context baru
	ctx5 := context.Background()

	// membuat context dengan deadline 5 detik dari program dijalankan
	ctxWithDeadline, cancel := context.WithDeadline(ctx5, time.Now().Add(5*time.Second))
	defer cancel()

	// menjalankan goroutine
	destination3 := createCounter2(ctxWithDeadline)

	// cek jumlah goroutine yg sedang berjalan setelah memanggil goroutine
	fmt.Println("Total goroutine after run", runtime.NumGoroutine())

	// print counter dari goroutine sampai timeout
	for n := range destination3 {
		fmt.Println("counter", n)
	}

	time.Sleep(3 * time.Second)

	// cek jumlah goroutine setelah dilakukan cancel
	fmt.Println("Total goroutine after cancel", runtime.NumGoroutine())
}

func greet(name, address string, age, weight int, list ...string) string {
	result := fmt.Sprintf("Nama saya %s, alamat saya %s, umur saya %d, berat saya %d ", name, address, age, weight)
	result += strings.Join(list, " ")
	return result
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

func changeValue(number *int) {
	*number = 20
}

type Person struct {
	name string
	age  int
}

func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s, my name is %s and i'm %d years old", msg, p.name, p.age)
}

func (p *Person) reviseName(name string) {
	p.name = name
}

type Employee struct {
	division string
	person   Person
}

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process ended")
}

func printFruit(index int, fruit string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	mutex.Unlock()
	wg.Done()
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hai, nama saya %s", student)
	c <- result
}

func greeting(cResult string, c chan string) {
	greeting := fmt.Sprintf("%s Happy Weekend", cResult)
	c <- greeting
}

func callDeferFunc() {
	defer deferFunc()
	fmt.Println("hai defer")
}

func deferFunc() {
	fmt.Println("ini defer")
}

func validPassword(password string) (string, error) {
	pl := len(password)
	if pl < 5 {
		return "", errors.New("password has to be more than 4 characters")
	}
	return "Valid password", nil
}

func catchErr() {
	r := recover()
	if r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func createCounter1(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func createCounter2(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				// memberikan jeda 1 detik untuk simulasi response lambat
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}
