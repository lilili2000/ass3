package main

import (
	"fmt"

	// mysql connector
	_ "github.com/go-sql-driver/mysql"
	sqlx "github.com/jmoiron/sqlx"
)

const (
	User     = "root"
	Password = ""
	DBName   = "ass3"
)

type Library struct {
	db *sqlx.DB
}

func (lib *Library) ConnectDB() {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", User, Password, DBName))
	if err != nil {
		panic(err)
	}
	lib.db = db
}

func exesqls(db *sqlx.DB, SQLs []string){
	for _, s := range SQLs{
		_, err := db.Exec(s)
		if err != nil{
			panic(err)
		}
	}
}



// CreateTables created the tables in MySQL
func (lib *Library) CreateTables() error {

	exesqls(lib.db, []string{
		"drop table borrow",
		"drop table student",
		"drop table book",
		`create table if not exists book(
			ISBN char(20),
			title char(20),
			author char(20),
			avail int,
			reason char(20),
			primary key(ISBN))`,
		`create table if not exists student(
			id int,
			pri bool,
			primary key(id))`,
		`create table if not exists borrow(
			student_id int,
			book_id char(20),
			bordata int,
			duedata int,
			ret int,
			extimes int,
			primary key(student_id,book_id),
			foreign key(student_id) references student(id),
			foreign key(book_id) references book(ISBN))`,

	})

	return nil
}

// AddBook add a book into the library
func (lib *Library) AddBook(title, author, ISBN string) error {

	exesqls(lib.db, []string{
		fmt.Sprintf("insert into book values('%s','%s','%s',1,'')",ISBN,title,author),
	})
	fmt.Println("The book has been added.")
	return nil
}

// etc...

func (lib *Library) RemoveBook(ISBN,rea string) error{
	exesqls(lib.db, []string{
		fmt.Sprintf("update book set avail=0,reason='%s' where ISBN='%s'",rea,ISBN),
	})
	fmt.Println("the book has been removed.")
	return nil
}

func (lib *Library) AddAccount(stuacc int) error{

	exesqls(lib.db, []string{
		fmt.Sprintf("insert into student values(%d,1)",stuacc),
	})
	fmt.Println("the account has been added.")
	return nil
}

func (lib *Library) QueryBook(value ,pattern string) error{
	rows, err := lib.db.Query(fmt.Sprintf("select ISBN,title,author from book where avail=1 and author='%s'",value))
	if err != nil{
		panic(err)
		fmt.Println("error")
		return nil
	}
	var ISBN,title,author string

	for rows.Next(){
		rows.Scan(&ISBN,&title,&author)
		fmt.Printf( "ISBN = %s, title = %s, author = %s ",ISBN,title,author)
	}

	return nil
}

func (lib *Library) BorrowBook(ISBN string,id int) error{
	rows, err := lib.db.Query(fmt.Sprintf("select * from student where id=%d and pri=1",id))
	if err != nil{
		panic(err)
	}
	if rows.Next(){
		rows, err := lib.db.Query(fmt.Sprintf(
			`select *
					from borrow
					where student_id = %d and ret=0 and duedata<%d `,id,20200501))
		if err != nil{
			panic(err)
		}
		cnt := 0
		for rows.Next(){
			cnt++
		}
		if cnt>=3 {
			fmt.Println("Your have more than 3 overdue books, you can not borrow the book.")
		} else{
			rows, err := lib.db.Query(fmt.Sprintf("select avail from book where ISBN='%s'",ISBN))
			if err!= nil{
				panic(err)
			}
			if rows.Next(){
				ava := 0
				rows.Scan(&ava)
				if ava==0{
					fmt.Println("The book is not available now")
				} else{
					exesqls(lib.db,[]string{fmt.Sprintf("insert into borrow values(%d,'%s',%d,%d,0,0)",id,ISBN,20200501,20200501+30),
					})
				}
			} else{
				fmt.Println("The book does not exist.")
			}
		}
	}
	return nil
}


func (lib *Library) QueryHistory(id int) error{
	rows, err := lib.db.Query(fmt.Sprintf("select student_id, book_id, bordata from borrow where student_id = %d",id))
	if err != nil{
		panic(err)
	}
	cnt := 0
	for rows.Next(){
		cnt ++
		var student_id,bordata int
		var book_id string
		rows.Scan(&student_id,&book_id,&bordata)
		fmt.Println(fmt.Sprintf("studentid = %d, bookid = %s, bordata = %d",student_id,book_id,bordata))
	}
	if cnt==0{
		fmt.Println("No borrowed books.")
	}
	return nil
}

func (lib *Library) QueryNotReturn(id int) error{
	rows, err := lib.db.Query(fmt.Sprintf("select book_id from borrow where student_id = %d and ret = 0",id))
	if err != nil{
		panic(err)
	}
	cnt := 0
	for rows.Next(){
		cnt ++
		var book_id string
		rows.Scan(&book_id)
		fmt.Println(fmt.Sprintf("bookid = %s",book_id))
	}
	if cnt==0{
		fmt.Println("No such book")
	}
	return nil
}

func (lib *Library) QueryDuedata(id int,ISBN string) error{
	rows, err := lib.db.Query(fmt.Sprintf("select duedata from borrow where student_id = %d and book_id = '%s'",id,ISBN))
	if err != nil{
		panic(err)
	}
	cnt :=0
	for rows.Next(){
		cnt ++
		var duedata int
		rows.Scan(&duedata)
		fmt.Println(fmt.Sprintf("duedata = %d",duedata))
	}
	if cnt==0{
		fmt.Println("No such book")
	}
	return nil
}

func (lib *Library) ExtendDueData(id int,ISBN string) error{
	rows, err := lib.db.Query(fmt.Sprintf("select extimes from borrow where student_id = %d and book_id = '%s'",id,ISBN))
	if err != nil{
		panic(err)
	}
	cnt := 0
	for rows.Next(){
		cnt ++
		var times int
		rows.Scan(&times)
		if times==3{
			fmt.Println("Can not be extended")
		} else{
			exesqls(lib.db,[]string{
				fmt.Sprintf("update borrow set duedata = duedata+10,extimes = extimes+1 where student_id = %d and book_id = %s",id,ISBN),
			})
			fmt.Println("The duedata has been extended")
		}
	}
	if cnt==0{
		fmt.Println("No such record")
	}
	return nil
}

func (lib *Library) CheckOverdue(id int) error{
	rows, err := lib.db.Query(fmt.Sprintf("select count(*) from borrow where student_id = %d and ret=0 and duedata<%d",id,20200501))
	if err != nil{
		panic(err)
	}
	for rows.Next(){
		var num int
		rows.Scan(&num)
		fmt.Println(fmt.Sprintf("There %d book(s) needed to be returned.",num))
	}
	return nil
}

func (lib *Library) ReturnBook(id int,ISBN string) error{
	rows, err := lib.db.Query(fmt.Sprintf("select * from borrow where student_id = %d and book_id = '%s'",id,ISBN))
	if err != nil{
		panic(err)
	}
	for rows.Next(){
		exesqls(lib.db,[]string{
			fmt.Sprintf("update borrow set ret=1,extimes=0 where student_id = %d and book_id = '%s'",id,ISBN),
		})
	}
	return nil
}

func main() {
	fmt.Println("Welcome to the Library Management System!")
}