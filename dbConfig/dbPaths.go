package dbConfig

import (
	firestore "cloud.google.com/go/firestore"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var DATABASE *mongo.Database
var DATABASE_FireBase *firestore.Client
var DATABASE_Sql *gorm.DB

const DATABASE_URL = "mongodb+srv://chamith:123@cluster0.ujlq82i.mongodb.net/?retryWrites=true&w=majority"

//const DATABASE_URL = "mongodb+srv://pasinduruwantha12:1234@cgaasui.amademj.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME = "AWSLambda"
