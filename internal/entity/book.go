package model

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define Errors
var (
	ERR_BOOK_NOT_FOUND                = errors.New("Book not found")
	ERR_BOOK_NOT_AVAILABLE            = errors.New("Book not available")
	ERR_BOOK_RESERVE_ALREADY_EXTENDED = errors.New("Book reservation already extended")
	ERR_BOOK_RESERVED_BY_OTHER        = errors.New("Book reserved by other")
	ERR_BOOK_NOT_BORROWED_BY_USER     = errors.New("Book not borrowed by user")
)

type Book struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	ISBN         string             `json:"isbn" bson:"isbn"`
	Name         string             `json:"name" bson:"name"`
	Genre        string             `json:"genre" bson:"genre"`
	Description  string             `json:"description" bson:"description"`
	Author       string             `json:"author" bson:"author"`
	Publisher    string             `json:"publisher" bson:"publisher"`
	PublishDate  time.Time          `json:"publishdate" bson:"publishdate"`
	TotalPages   int                `json:"totalpages" bson:"totalpages"`
	Condition    bool               `json:"condition" bson:"condition"`
	Availability int                `json:"availability" bson:"availability"` // availiblity enum (0,1,2,3) 0:available 1:unavailable 2:reserved 3:borrowed
	Location     string             `json:"location" bson:"location"`
	ImageURL     string             `json:"image_url" bson:"image_url"`
	StartDate    time.Time          `json:"startdate" bson:"startdate"`
	EndDate      time.Time          `json:"enddate" bson:"enddate"`
	ExtendedDate time.Time          `json:"extendeddate" bson:"extendeddate"`
}

type BookReplacementList struct {
	TimeReplace       time.Time `json:"TimeReplace" bson:"TimeReplace"`
	ManagerID         int       `json:"ManagerID" bson:"ManagerID"`
	LibrarianID       int       `json:"LibrarianID" bson:"LibrarianID"`
	UniqueCodeReplace string    `json:"UniqueCodeReplace" bson:"UniqueCodeReplace"`
}

type BookActivityList struct {
	BorrowDate         time.Time `json:"borrowdate" bson:"borrowdate"`
	ExtendDate         time.Time `json:"ExtendDate" bson:"ExtendDate"`
	ReturnDate         time.Time `json:"returndate" bson:"returndate"`
	UniqueCodeActivity string    `json:"UniqueCodeActivity" bson:"UniqueCodeActivity"`
	ManagerID          int       `json:"ManagerID" bson:"ManagerID"`
	LibrarianID        int       `json:"LibrarianID" bson:"LibrarianID"`
}

type SearchBook struct {
	UniqueCodeSearch string `json:"UniqueCodeSearch" bson:"UniqueCodeSearch"`
	SearcherID       int    `json:"SearcherID" bson:"SearcherID"`
}

type ReserveBook struct {
	UniqueCodeReserve string `json:"UniqueCodeReserve" bson:"UniqueCodeReserve"`
	ReserverID        int    `json:"ReserverID" bson:"ReserverID"`
}

//	type BorrowBook struct {
//		BorrowDate   time.Time          `json:"BorrowDate" bson:"BorrowDate"`
//	 ReturnDate   time.Time			`json:"returndate" bson:"returndate"`
//	}

type BookName struct {
	Name string `json:"name" bson:"name"`
}
