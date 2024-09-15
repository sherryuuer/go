package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		// 这里不需要冒号是因为var的声明作用
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 4
	fmt.Println(heidi)

	casey.Boarded = true
	if casey.Boarded {
		fmt.Println(casey.Name, "is boarded")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")

}

// Go 语言中的 `struct` 和传统的面向对象编程中的 `class` 有相似之处，都是用来封装数据的，但 Go 语言不直接提供类（`class`）的概念。相反，Go 使用结构体（`struct`）和方法（`method`）来实现类似类的功能。

// ### 1. **Go 中的结构体（struct）**
// `struct` 是一种自定义类型，允许你将不同的数据类型组合在一起，类似于类的字段。你可以用它来表示一个对象的属性。

// #### 示例：
type Person struct {
	Name string
	Age  int
}

// 这里 `Person` 是一个结构体，包含两个字段：`Name` 和 `Age`，类似于类的属性。

// ### 2. **Go 中的方法（method）**
// 在 Go 中，**方法**是定义在某个类型（比如 `struct`）上的函数。虽然 Go 没有类的概念，但你可以为 `struct` 定义方法，从而实现类似类的行为。

// 方法是通过在函数前面指定一个**接收者（receiver）**来定义的。接收者指定了方法与哪个类型（通常是结构体）关联。

// #### 方法示例：
// 定义方法，接收者是 Person
func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s, and I am %d years old.\n", p.Name, p.Age)
}

func main2() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet() // 调用方法
}

// **方法的关键点**：
// - **接收者（Receiver）**：在上面的例子中，`Person` 类型是方法 `Greet` 的接收者，方法是通过 `p.Greet()` 调用的，`p` 是 `Person` 类型的实例。
// - 通过将方法与结构体关联，你可以为某个结构体类型实现行为逻辑，就像类中的成员方法一样。

// ### 3. **值接收者 vs 指针接收者**

// Go 的方法可以有两种接收者类型：
// - **值接收者（Value Receiver）**：方法接收一个结构体的副本，不能修改原结构体的字段。
// - **指针接收者（Pointer Receiver）**：方法接收一个结构体的指针，可以修改原结构体的字段。

// #### 值接收者示例：
func (p Person) ChangeName(newName string) {
	// p.Name = newName // 这是在修改副本, 因为这里修改的是副本不会被使用，编译发生unused错误，所以我comment掉了
}

// 调用时：
func main3() {
	p := Person{Name: "Alice", Age: 30}
	p.ChangeName("Bob")
	fmt.Println(p.Name) // 输出依然是 "Alice"，因为 p 是一个副本
}

// #### 指针接收者示例：
func (p *Person) ChangeName2(newName string) {
	p.Name = newName // 直接修改原结构体，这里就没事
}

// 调用时：
func main4() {
	p := Person{Name: "Alice", Age: 30}
	p.ChangeName2("Bob")
	fmt.Println(p.Name) // 输出 "Bob"，因为 p 是指针，直接修改了原结构体

}

// ### 4. **struct 和 class 的区别**

// 虽然 `struct` 和 `class` 有相似之处，但 Go 没有提供类继承的概念。Go 的设计理念是使用**组合**（composition）代替继承，通过嵌套结构体实现代码的复用和扩展。Go 语言采用的是一种**接口**和**组合**的方式来实现类似面向对象的行为。

// #### 主要区别：
// - **没有继承**：Go 没有传统的类继承机制。
// - **没有构造函数**：Go 没有类的构造函数（constructor），但是你可以通过工厂函数的方式创建结构体实例。
// - **接口（Interface）**：Go 提供了接口来实现多态，而不是通过继承。

// ### 总结

// - Go 的 `struct` 类似于类的属性部分，用于定义对象的属性。
// - Go 中可以通过给结构体定义**方法**，实现类似类的方法行为。
// - Go 没有类继承，但通过组合和接口可以实现灵活的代码重用和多态。

// 这使得 Go 语言的结构体和方法体系既简单又强大，非常适合构建高效和可扩展的系统。
// cool!
