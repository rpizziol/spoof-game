package main // Declare main package

/*func getSum(num1 int, num2 int) int {
	return num1 + num2
}*/

//func main() {  // main() function executes by default while running 'main' package
/*    fmt.Println("Hello, World!")
    name, email := "Brad", "brad@gmail.com"
    size := 1.3
    var age int32 = 37
    const isCool = true
    fmt.Println(name, age, isCool, size, email)
    fmt.Printf("%T\n", name)
    fmt.Println(math.Floor(2.7))
    fmt.Println(util.Reverse("pippo"))
    fmt.Println(util.Greeting(name))
	fmt.Println(getSum(2,3))*/

/*var fruitArr [2]string

fruitArr[0] = "apple"
fruitArr[1] = "orange"

fmt.Println(fruitArr[1])*/

/*fruitArr := []string{"Apple", "Orange", "Banana"}
fmt.Println(fruitArr)
fmt.Println(fruitArr[1:3])*/
/*
	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}*/

//color := "red"

/*if color == "red" {
	fmt.Println("coor is red")
} else if color == "blue" {
	fmt.Println("coor is blue")
} else {
	fmt.Println("coor is not blue or red")
}
*/

/*switch color {
case "red":
	fmt.Println("color is red")
case "blue":
	fmt.Println("color is blue")
default:
	fmt.Println("color is not blue or red")
}*/

/*i := 1
for i <= 10 {
	fmt.Println(i)
	i++
}*/

/*for i := 1; i <= 100; i++ {
	if i % 15 == 0 {
		fmt.Println("FizzBuzz")
	} else if i % 3 == 0 {
		fmt.Println("Fizz")
	} else if i % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
	}
}*/

/*emails := make(map[string]string)

emails["Bob"] = "bob@gmail.com"
emails["Kate"] = "kate@gmail.com"
emails["Mike"] = "mike@gmail.com"*/

/*emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}

fmt.Println(emails)

ids := []int{22,43,5,66,89}

for _, id := range ids {
	fmt.Printf("ID: %d\n", id)
}

sum := 0
for _, id := range ids {
	sum += id
}

fmt.Println("Sum", sum)

// Range with map

for k, v := range emails {
	fmt.Printf("%s: %s\n", k, v)
}

for k := range emails {
	fmt.Println("name: " + k)
}
*/
/*a := 5
b := &a

fmt.Println(a, b)

fmt.Printf("%T\n", a)

fmt.Println(*b)

*b = 10
fmt.Println(a)*/

/*	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}*/
//}

/*func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}*/

/*type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}*/

/*type Person struct {
	firstName, lastName, city, gender string
	age                               int
}
*/
/*func main() {
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Roberto", "Pizziol", "Roma", "m", 28}
	//fmt.Println(person1.firstName, person2.lastName)

	person1.age++

	fmt.Println(person1, person2)
	fmt.Println(person1.greet())
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}*/
/*
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + ", and I am " + strconv.Itoa(p.age) + "."
}

func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}*/
