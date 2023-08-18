package dbConfig

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/option"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMongoDB() {
	fmt.Println("Connecting to mongo cluster")

	// Create a context
	ctx := context.Background()

	// Connect to MongoDB Atlas
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DATABASE_URL))
	if err != nil {
		log.Fatal(err)
	}

	//List Available Databases in the Cluster
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully Connected to mongo cluster")
	}

	DATABASE = client.Database(DATABASE_NAME)

	fmt.Println(databases)

}

func ConnectToFirestore(ctx context.Context, credentialsURL string) {
	// Download the credentials file from S3 and save it to a temporary location
	tmpFile, err := os.Create("/tmp/firebase-adminsdk.json")
	if err != nil {
		log.Fatalf("Failed to create temporary file: %v", err)
	}
	defer tmpFile.Close()

	response, err := http.Get(credentialsURL)
	if err != nil {
		log.Fatalf("Failed to fetch credentials file from S3: %v", err)
	}
	defer response.Body.Close()

	_, err = io.Copy(tmpFile, response.Body)
	if err != nil {
		log.Fatalf("Failed to write credentials file to temporary location: %v", err)
	}

	config := &firebase.Config{
		ProjectID: "go-lambda",
	}

	opt := option.WithCredentialsFile("/tmp/firebase-adminsdk.json")
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	DATABASE_FireBase = client
}

func InitDB() {
	// Replace the following variables with your AWS MySQL database details
	dsn := "admin:12345678@tcp(user.ctqroturhzsd.us-east-1.rds.amazonaws.com:3306)/Users"
	var err error
	DATABASE_Sql, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

}
