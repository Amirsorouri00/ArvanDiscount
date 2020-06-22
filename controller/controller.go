package controller

import (
	"log"
	"time"
	"net/http"
	"math/rand"

	"github.com/go-pg/pg/v9"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	orm "github.com/go-pg/pg/v9/orm"
)


type DiscountManager struct {
	Id	 			string 		`json:"id"`
	Code			string		`json:"code"`
	DiscountGift	bool		`json:"discount_gift"` // [false = discount, true = gift]
	DiscountId		string		// Belongs to
	GiftId			string		// Belongs to
	StreamId		string		// Belongs to	
	CreatedAt 		time.Time 		
	UpdatedAt 		time.Time
}

type Discount struct {
	Id					string				`json:"id"`
	Percent				int					`json:"percent"`
	Amount				float64					`json:"amount"`
	PercentAmount		bool				`json:"percent_amount"` // [false = percent, true = amount]
	DiscountManager		*DiscountManager	// Has one relationship
	CreatedAt 			time.Time 			
	UpdatedAt 			time.Time			
}

type Gift struct {
	Id					string				`json:"id"`
	Amount				float64				`json:"amount"`
	Used				int					`json:"used"`	// Default must be 0
	Capacity			int					`json:"capacity"`
	DiscountManager		*DiscountManager	// Has one relationship
	CreatedAt 			time.Time 		
	UpdatedAt 			time.Time
}

// Just for test
type Stream struct {
	Id					string				`json:"id"`
	Name				string				`json:"stream_name"`
	Start 				time.Time
	Finish				time.Time			// start.Add(time.Minute * 10)
	Status				string			
	DiscountManagers	[]*DiscountManager	// Has many relationship
	CreatedAt 			time.Time 		
	UpdatedAt 			time.Time
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

var seededRand *rand.Rand
func InitiateSeed() {
	seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))	
}

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" + "!@#$&*"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandString(length int) string {
	return StringWithCharset(length, charset)
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


// Get All Discounts
func GetAllDiscounts(c *gin.Context) {
	var discounts []Discount
	err := dbConnect.Model(&discounts).Relation("DiscountManager").Select()

	if err != nil {
		log.Printf("Error while getting all discounts, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Discounts",
		"data": discounts,
	})
	return
}


type AddDiscountType struct {
	Amount        float64 `json:"amount"`
	Percent       int     `json:"percent"`
	PercentAmount bool    `json:"percent_amount"`
	DiscountGift  bool    `json:"discount_gift"`
	StreamId      string  `json:"stream_id"`
}

// Add Discount OK
func AddDiscount(c *gin.Context) {
	var req AddDiscountType
	c.BindJSON(&req)
	uuid := guuid.New().String()
	insertError := dbConnect.Insert(&Discount{
		Id: uuid,
		Percent: req.Percent,
		Amount: req.Amount,
		PercentAmount: req.PercentAmount,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new discount into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	var stream Stream	// Check if the streamId exist or not
	err := dbConnect.Model(&stream).Where("id = ?", req.StreamId).Select()
	if err != nil {
		log.Printf("AddDiscount: Error while getting stream by name from db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	insertError2 := dbConnect.Insert(&DiscountManager{
		Id: guuid.New().String(),
		Code: RandString(8),
		DiscountGift: req.DiscountGift,
		DiscountId: uuid,
		StreamId: stream.Id,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if insertError2 != nil {
		log.Printf("AddDiscount: Error while inserting new DiscountManager into db, Reason: %v\n", insertError2)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Discount created Successfully",
	})
	return
}


// Get All Streams OK
func GetAllStreams(c *gin.Context) {
	var streams []Stream
	err := dbConnect.Model(&streams).Relation("DiscountManagers").Select()
	if err != nil {
		log.Printf("Error while getting all streams, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Streams",
		"data": streams,
	})
	return
}

// Add Stream OK
func AddStream(c *gin.Context) {
	var stream Stream
	c.BindJSON(&stream)
	start := time.Now().Local().Add(time.Minute * -45)
	finish := time.Now().Local().Add(time.Minute * 45)
	insertError := dbConnect.Insert(&Stream{
		Id: guuid.New().String(),
		Name: stream.Name,
		Start: start,
		Finish: finish,
		Status: stream.Status,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new stream into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Stream created Successfully",
	})
	return
}