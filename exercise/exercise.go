package exercise

import (
	"fmt"
	"strings"
	"time"

	"github.com/Ko-TTae/study-go/accounts"
	"github.com/Ko-TTae/study-go/mydict"
	"github.com/Ko-TTae/study-go/something"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func lengA(name string) (length int, uppsercase string) {
	length = len(name)
	uppsercase = strings.ToUpper(name)
	return
}

func lengAPl(name string) (length int, uppsercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppsercase = strings.ToUpper(name)
	return
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	case koreanAge > 50:
		return true
	}
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func arryExer() {
	names := []string{"a", "b", "ko"}
	// append 는 새로운 arr 반환 names를 바꾸지 않음
	names = append(names, "ttae")
	fmt.Println(names)
}

func mapExer() {
	kottae := map[string]string{"name": "kottae", "age": "12"}
	fmt.Println(kottae)

	for key, value := range kottae {
		fmt.Println(key, value)
	}
}

type person struct {
	name    string
	age     int
	favFood []string
}

func structureExer() {
	favFood := []string{"빵", "국밥"}
	kottae := person{name: "kottae", age: 28, favFood: favFood}
	fmt.Println(kottae)
}

func bankMo() {
	account := accounts.NewAccount("kottae")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)

	}
	fmt.Println(account.Balance())

	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)
}

func dic() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}

func dicAdd() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func dicUpdate() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
}

func dicDelete() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}

func Exercise() {
	fmt.Println("Hello world!")
	something.SayHello()

	const name string = "ko"
	var name2 string = "asfd"
	fmt.Println(multiply(2, 2))
	fmt.Println(name, name2)

	totalLength, _ := lenAndUpper("shko")
	fmt.Println(totalLength)

	repeatMe("a", "b", "c")

	totalL, up := lengA("kottae")
	fmt.Println(totalL, up)

	totalL2, up2 := lengAPl("kottae")
	fmt.Println(totalL2, up2)

	total := superAdd(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(total)

	fmt.Println(canIDrink(15))

	a := 2
	b := &a
	fmt.Println("a:", &a, "->", a)
	fmt.Println("b:", b, "->", *b)
	a = 5
	fmt.Println("a was changed")
	fmt.Println("a:", &a, "->", a)
	fmt.Println("b:", b, "->", *b)
	*b = 20
	fmt.Println("b* was changed")
	fmt.Println("a:", &a, "->", a)
	fmt.Println("b:", b, "->", *b)

	arryExer()
	mapExer()
	structureExer()

	bankMo()

	dic()
	dicAdd()
	dicUpdate()
	dicDelete()
}

func goRoutineExer() {
	c := make(chan string)
	people := [2]string{"a", "b"}
	for _, p := range people {
		go isCool(p, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("wating", i)
		fmt.Println(<-c)
	}
}

func isCool(p string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- p + " is cool"
}
