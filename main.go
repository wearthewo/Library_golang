package main

import (
	"fmt"
	"time"
)

// Author represents a book's author
type Author struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
}

// Book represents a book in the library
type Book struct {
	ID          int
	Title       string
	Author      Author
	Genre       string
	PublishedAt time.Time
	IsAvailable bool 
	Price       int
}

// Member represents a library member
type Member struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	JoinedAt  time.Time
}

// Loan represents a book loan record
type Loan struct {
	ID        int
	Book      Book
	Member    Member
	LoanDate  time.Time
	DueDate   time.Time
	ReturnDate time.Time
	IsReturned bool
} 


// Library represents the library system
type Library struct {
	Books   []Book
	Members []Member
	Loans   []Loan 
}  

func  (l Library) registerMember() []Member  {    
	FirstName := "" 
	LastName := ""
	Email := ""  
	var ID int 
	fmt.Println("Welcome to Library") 
	fmt.Printf("Enter your firstname: ") 
	_,err := fmt.Scanln(&FirstName)
	if err!= nil {
		return l.Members
	}
	fmt.Printf("Enter your lastname: ") 
	_,err = fmt.Scanln(&LastName)
	if err!= nil {
		return l.Members
	}  
	fmt.Printf("Enter your email: ") 
	_,err = fmt.Scanln(&Email)
	if err!= nil {
		return l.Members
	} 
	ID = len(l.Members) + 1  
	newMember := Member{ 
		ID: ID, 
		FirstName: FirstName, 
		LastName: LastName, 
		Email: Email,
	} 
	l.Members = append(l.Members,newMember) 
	return l.Members
}  

func (l Library) availableBooks() int {  
	existedBooks := 0
	for _,book  := range l.Books {
		if(book.IsAvailable) {
			existedBooks ++ 
		} 
	} 
	return existedBooks
}  

func (l Library) countOfBooks(existedBooks int) int {  
	count := existedBooks
	for _,book  := range l.Loans {
		if(!book.IsReturned) {
			count --
		} 
	} 
	return count
}  

func (l Library) priceOfBooks () map[string] int {
	//availableBooksSlice := [] int {}
	availableBooksMap := make(map[string]int)

	for _,book := range(l.Books) {
		if book.IsAvailable {
			//availableBooksSlice = append(availableBooks, book.Title)
			availableBooksMap[book.Title] = book.Price 
			fmt.Printf("The available books you can borrow are %s and it costs %d: ", book.Title, availableBooksMap[book.Title])
		} 
	} 
	return availableBooksMap
}  

func (l Library) discountPolicy (availableBooks map[string] int) float32  {
	var answer string
	var choice1 string 
	var choice2 string 
	var choice3 string  
	var choice4 string
	sum:=0 
	fmt.Println("Our library offers you a discount: If you borrow 2 books the third is 50% cheaper!") 
	fmt.Printf("Do you want the offer YES OR NO ? ", ) 
	fmt.Scanln(&answer) 
	if answer == "YES" {  
			fmt.Printf("Choose three books from the available books %v",availableBooks) 
			fmt.Scan(&choice1, &choice2, &choice3)
			fmt.Printf("You choose the books %s, %s, %s: ", choice1, choice2, choice3) 
			price1 := availableBooks[choice1] 
			price2 := availableBooks[choice2] 
			price3 := availableBooks[choice3] 
			sum = sum + price1 + price2 + (price3/2)
			fmt.Println("Nice choice, the total cost with the discount is:  ", sum) 
			return float32(sum)

		} 
	fmt.Println("You didn't choose our offer ") 
	fmt.Printf("Choose one book from the available books:  %v",availableBooks)  
	fmt .Scan(&choice4) 
	fmt.Printf("You choose the books %s: ", choice4) 
	price4:=availableBooks[choice4] 
	fmt.Println("Nice choice, the total cost of the book tou chose is:  ", price4) 
	return float32(price4)		
	}

func main() {
	// Example to create and manage the library system
	author1 := Author{
		ID:        1,
		FirstName: "George",
		LastName:  "Orwell",
	}  

	author2 := Author{
		ID:        2,
		FirstName: "Suzzane",
		LastName:  "Collins",
	} 

	author3 := Author{
		ID:        3,
		FirstName: "J.K",
		LastName:  "Rowling",
	}  

	author4 := Author{
		ID:        4,
		FirstName: "J.R.R",
		LastName:  "Tolkien",
	}  

	author5 := Author{
		ID:        5,
		FirstName: "C.S",
		LastName:  "Lewis",
	} 
	author6 := Author{
		ID:        6,
		FirstName: "Gopher",
		LastName:  "GO",
	}

	book1 := Book{
		ID:          1,
		Title:       "1984",
		Author:      author1,
		Genre:       "Dystopian",
		PublishedAt: time.Date(1949, time.June, 8, 0, 0, 0, 0, time.UTC),
		IsAvailable: true,  
		Price: 50,
	} 
	book2 := Book { 
		ID:			2, 
		Title:      "The Hunger Games",
		Author:      author2,                                             //Suzzane Collins 
		Genre:       "Fiction",
		PublishedAt:  time.Date(2008, time.March, 7, 0, 0, 0, 0, time.UTC), 
		IsAvailable:  false, 
		Price: 40,
	} 

	book3 := Book { 
		ID:			3,
		Title:      "Harry Potter and the Order of the Phoenix", 
		Author:     author3,
		Genre:       "Fantasy" , 																	//J.K Rowling
		PublishedAt:  time.Date(2003, time.June, 6, 0, 0, 0, 0, time.UTC),
		IsAvailable:  false, 
		Price: 55,

	}  

	book4 := Book{  
		ID: 4, 
		Title: "The Hobbit and The Lord of the Rings",  
		Author:      author4,
		Genre:       "Fiction" ,  
		PublishedAt: time.Date(1954, time.January, 1, 0,0,0,0, time.UTC), 
		IsAvailable: true,  
		Price: 60,            //J.R.R Tolkien

	}  

	book5 := Book{
		ID:  5, 
		Title: "The Chronicles of Narnia", 
		Author: author5,
		Genre: "Fantasy",               //C.S Lewis
		PublishedAt: time.Date(1956, time.January, 1, 0,0,0,0, time.UTC ), 
		IsAvailable: true,  
		Price: 45,
	} 

	book6:= Book{
		ID:  6, 
		Title: "Golang", 
		Author: author6,
		Genre: "Software",               
		PublishedAt: time.Date(1956, time.January, 1, 0,0,0,0, time.UTC ), 
		IsAvailable: true,  
		Price: 35,
	}


	member1 := Member{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	} 

	member2 := Member{
		ID:        2,
		FirstName: "Mike",
		LastName:  "James",
		Email:     "mike.james@example.com",
	} 
	member3 := Member{
		ID:        3,
		FirstName: "Derrick",
		LastName:  "Rose",
		Email:     "derrick.rose@example.com",
	} 
	member4 := Member{
		ID:        4,
		FirstName: "Leo",
		LastName:  "Martins",
		Email:     "leo.martins@example.com",
	} 

	member5 := Member{
		ID:        5,
		FirstName: "Papu",
		LastName:  "Gomez",
		Email:     "papu.gomez@example.com",
	}

	loan1 := Loan{
		ID:        1,
		Book:      book1,
		Member:    member2,
		LoanDate:  time.Now(),
		DueDate:   time.Now().Add(14 * 24 * time.Hour), // due in 14 days
		IsReturned: false,
	}  

	loan2 := Loan{
		ID:        2,
		Book:      book4,
		Member:    member2,
		LoanDate:  time.Now(),
		DueDate:   time.Now().Add(14 * 24 * time.Hour), // due in 14 days
		IsReturned: true,
	}  



	// Create the library instance and add data
	library := Library{
		Books:   []Book{book1,book2,book3,book4,book5,book6},
		Members: []Member{member1,member2,member3,member4,member5},
		Loans:   []Loan{loan1,loan2},
	}

	// Print library details
	fmt.Println("Library System: ")
	fmt.Println("Books: ", library.Books)
	fmt.Println("Members: ", library.Members)
	//fmt.Println("Loans: ", library.Loans)  

	members := library.registerMember() 
	fmt.Println(members) 
	existedBooks := library.availableBooks() 
	fmt.Println(existedBooks) 
	count:=library.countOfBooks(existedBooks) 
	fmt.Println(count) 
	availableBooks := library.priceOfBooks() 
	fmt.Println(availableBooks) 
	totalDiscountCost := library.discountPolicy(availableBooks) 
	fmt.Println(totalDiscountCost)

}  


/*create logic for: 
1) total books of the library 
2) lending books total -= lending 
3) returned books = returned + lending = total 
4) offers for the books + discounts 
5) METHODS - POINTERS 
6) slice of struct books me times 
7) rwta to xrhsth poio vivlio thelei dinei thn timh kai o xrhsths pliwnei kai oi times sinathrizontai ftiaxnontas ena tameio gia th library.*/


 
