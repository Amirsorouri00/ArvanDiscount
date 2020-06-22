package controller

import (
	// "log"
	// "net/http"
	// "time"

	// "github.com/gin-gonic/gin"
	// orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	// guuid "github.com/google/uuid"
)


// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}


