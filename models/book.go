package models

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var id int = 1                //for id increase +1 when added new item Book
var stock_code int = 10000000 //for stock_code increase +1 when added new item Book
//Book contains Author infos
type Book struct {
	Id, StockCode, ISBN, StockNo, PagesNo int
	Price                                 float64
	Name                                  string
	IsDeleted                             bool
	Author
}

//Author
type Author struct {
	//id            int
	name, surname string
	isDeleted     bool
}

//constructor for Author
func (auth *Author) Init(name string, surname string) {
	auth.name = name
	auth.surname = surname
	auth.isDeleted = false
}

//Author print func brings together author infos
func (auth *Author) Print() string {
	return auth.name + " " + auth.surname
}

//constructor for Book
func (book *Book) Init(newName string, writer Author) {
	book.Name = newName
	book.PagesNo = randomInt(50, 500) //random
	book.Price = randFloat(15, 250)   //random
	book.StockNo = randomInt(15, 20)  //random
	book.Author = writer
	book.ISBN = randomInt(1000000000, 9999999999) //random
	book.Id = id
	book.StockCode = stock_code
	book.IsDeleted = false

	id += 1
	stock_code += 1
}

type deletable interface {
	delete(item int) Book
}

//delete func for Book returns changed book
func (book *Book) delete(item int) Book {
	if book.Id == item && !book.isDeleted {
		book.IsDeleted = true
		//fmt.Println(book)
		//fmt.Println(book.IsDeleted)
	} else {
		fmt.Println("The book is already deleted.")
		os.Exit(3)
	}
	return *book
}

//this func for deletable interface to call delete func generator
func DeleteInterface(d deletable, item int) {
	d.delete(item)
}

//Buy func comes buyying count and process on stock number returns changed book infos
func (book *Book) Buy(count int) Book {
	if book.StockNo < count {
		fmt.Println("Given count is higher than stock number")
		os.Exit(3)
	} else {
		book.StockNo -= count
	}
	return *book
}

//Ramdom int generates for Pages, ISBN, Stock number
func randomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	rnd := rand.Intn(max-min) + min
	return rnd
}

//Random floats genarates for Price
func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	res := min + rand.Float64()*(max-min)
	return res
}
