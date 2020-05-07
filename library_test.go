package main

import (
	"testing"
)

func TestCreateTables(t *testing.T) {
	lib := Library{}
	lib.ConnectDB()
	err := lib.CreateTables()
	if err != nil {
		t.Errorf("can't create tables")
	}
}

func TestAddBook(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.AddBook("book1","David","w1234567")
	if err != nil{
		t.Errorf("can't add the book")
	}
}

func TestRemoveBook(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.RemoveBook("w1234567","lost")
	if err != nil{
		t.Errorf("can't remove the book")
	}
}

func TestAddAccount(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.AddAccount(1)
	if err != nil{
		t.Errorf("can't add the account")
	}
}

func TestQueryBook(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.QueryBook("David","author")
	if err != nil{
		t.Errorf("can't query the book")
	}
}

func TestBorrowBook(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.BorrowBook("w1234567",1)
	if err != nil{
		t.Errorf("can't borrow the book")
	}
}

func TestQueryHistory(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.QueryHistory(1)
	if err != nil{
		t.Errorf("can't borrow the history")
	}
}

func TestQueryNotReturn(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.QueryNotReturn(1)
	if err != nil{
		t.Errorf("can't query")
	}
}

func TestQueryDuedata(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.QueryDuedata(1,"w1234567")
	if err != nil{
		t.Errorf("can't query")
	}
}

func TestExtendDuedata(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.ExtendDueData(1,"w1234567")
	if err != nil{
		t.Errorf("can't query")
	}
}

func TestCheckOverdue(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.CheckOverdue(1)
	if err != nil{
		t.Errorf("can't query")
	}
}

func TestReturnBook(t *testing.T){
	lib := Library{}
	lib.ConnectDB()
	err := lib.ReturnBook(1,"w1234567")
	if err != nil{
		t.Errorf("can't return the book")
	}
}