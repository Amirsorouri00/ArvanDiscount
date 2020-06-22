package controller

import (
	"log"
	// "net/http"
	"time"

	// "github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	// guuid "github.com/google/uuid"
)


type DiscountManager struct {
	Id	 			string 		`json:"id"`
	Code			string		`json:"code"`
	Discount_Gift	bool		`json:"discount_gift"` // [false = discount, true = gift]
	DiscountId		string		// Belongs to
	GiftId			string		// Belongs to
	StreamId		string		// Belongs to	
	CreatedAt 		time.Time 		
	UpdatedAt 		time.Time
}

type Discount struct {
	Id					string				`json:"id"`
	Percent				int					`json:"percent"`
	Amount				int					`json:"amount"`
	Percent_Amount		bool				`json:"percent_amount"` // [false = percent, true = amount]
	DiscountManager		*DiscountManager	// Has one relationship
	CreatedAt 			time.Time 			
	UpdatedAt 			time.Time			
}

type Gift struct {
	Id					string				`json:"id"`
	Amount				int					`json:"amount"`
	Used				int					`json:"used"`	// Default must be 0
	DiscountManager		*DiscountManager	// Has one relationship
	CreatedAt 			time.Time 		
	UpdatedAt 			time.Time
}

// Just for test
type Stream struct {
	Id					string				`json:"id"`
	Name				string				`json:"stream_name"`
	DiscountManagers	[]*DiscountManager	// Has many relationship
	CreatedAt 			time.Time 		
	UpdatedAt 			time.Time
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

// Create Tables OK
func CreateTables(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createError1 := db.CreateTable(&DiscountManager{}, opts)
	createError2 := db.CreateTable(&Discount{}, opts)
	createError3 := db.CreateTable(&Gift{}, opts)
	createError4 := db.CreateTable(&Stream{}, opts)
	if createError1 != nil {
		log.Printf("Error while creating Discount Manager table, Reason: %v\n", createError1)
		return createError1
	}
	if createError2 != nil {
		log.Printf("Error while creating Discount table, Reason: %v\n", createError2)
		return createError2
	}
	if createError3 != nil {
		log.Printf("Error while creating Gift table, Reason: %v\n", createError3)
		return createError3
	}
	if createError4 != nil {
		log.Printf("Error while creating Stream table, Reason: %v\n", createError4)
		return createError4
	}

	table_names := [4]string {"DiscountManager", "Discount", "Gift", "Stream"}
	for i := 0; i < 4; i++ {
		log.Printf("%s table created.", table_names[i])
	}
	return nil
}



