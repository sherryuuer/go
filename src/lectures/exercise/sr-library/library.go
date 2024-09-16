//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string // 这叫做type alias，是对type的一种别名表达
type Name string  // member name

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct { // 借书人的结构体
	name  Name                // 这里使用了类型别名
	books map[Title]LendAudit // 这里表示借的书的一个标题和时间戳的map结构，用于记录
}

type BookEntry struct {
	total  int // total owned by the library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member     // 感想：这种type的别名挺好的
	books   map[Title]BookEntry // 这里是书的情报
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books { // 注意这里遍历的是member的books
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		// 这里使用了member的name属性
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), ":", returnTime)
	}
}

func printMembersAudit(library *Library) { // 这里使用复数体，用上面的函数，打印所有的member的信息
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println("Title:", title, "/total:", book.total, "/lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	// 检查书是否属于图书馆
	if !found {
		fmt.Println("The book is not part of the library")
		return false
	}
	// 检查书是否还有剩余
	if book.lended == book.total {
		fmt.Println("No more book to be lend, all has been lended out")
		return false
	}

	book.lended += 1
	library.books[title] = book // 这里因为更新了book的数量信息，所以要更新library的信息
	member.books[title] = LendAudit{checkOut: time.Now()}

	return true
}

func returnBook(Library *Library, title Title, member *Member) bool {
	book, found := Library.books[title]
	// 检查书是否属于图书馆目录
	if !found {
		fmt.Println("The book is not part of the library")
		return false
	}
	// 检查用户是否借走过这本书
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not checkout this book")
		return false
	}

	book.lended -= 1
	Library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit // 如果这里不重写，checkIn上就会报错：unused write to checkIn

	return true

}

func main() {
	library := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}

	library.books["WebApps in Go"] = BookEntry{
		total: 4, lended: 0,
	}
	library.books["Happy Java"] = BookEntry{
		total: 5, lended: 0,
	}
	library.books["Go and Microservices"] = BookEntry{
		total: 3, lended: 0,
	}
	library.books["Python is all you need"] = BookEntry{
		total: 6, lended: 0,
	}

	library.members["Tracy"] = Member{"Tracy", make(map[Title]LendAudit)}
	library.members["Herry"] = Member{"Herry", make(map[Title]LendAudit)}
	library.members["Sally"] = Member{"Sally", make(map[Title]LendAudit)}

	fmt.Println("\n---Initial---")
	printMembersAudit(&library)
	printLibraryBooks(&library)

	member := library.members["Tracy"]
	checkedOut := checkoutBook(&library, "Happy Java", &member)
	fmt.Println("\n---Checked out a book---")
	if checkedOut {
		printMembersAudit(&library)
		printLibraryBooks(&library)
	}

	returned := returnBook(&library, "Happy Java", &member)
	fmt.Println("\n---Returned a book---")
	if returned {
		printMembersAudit(&library)
		printLibraryBooks(&library)
	}

}
