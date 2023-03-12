package main

import (
	"fmt"
	"golang-coretan/helpers"
	"reflect"
	"strings"
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

	var str = "САШАРВО"
	for pos, char := range str {
		fmt.Printf("%U %c %d\n", char, char, pos)
	}
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
