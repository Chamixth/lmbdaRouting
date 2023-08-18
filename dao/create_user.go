package dao

import (
	"context"
	"log"
	"lambdarouting/dbConfig"
	"lambdarouting/dto"
)

func CreateUser(ctx context.Context, application dto.User) error {

	//////////////////////////////////MongoDb///////////////////////////////////////////

	_, err := dbConfig.DATABASE.Collection("User").InsertOne(ctx, application)

	//////////////////////////////////MongoDb///////////////////////////////////////////

	/////////////////////////////////////////Firebase-firestore///////////////////////////////

	// _, _, err := dbConfig.DATABASE_FireBase.Collection("users").Add(ctx, application)

	/////////////////////////////////////////Firebase-firestore///////////////////////////////

	/////////////////////////////////////////SQL/////////////////////////////////////////////

	// stmt, err := dbConfig.DATABASE_Sql.Prepare("INSERT INTO User2 (Username, Email, Password) VALUES (?, ?, ?)")
	// if err != nil {
	// 	return 0, err
	// }
	//////////////////////////////////////////////Common///////////////////////////////////////////

	if err != nil {
		return err
	}

	//////////////////////////////////////////////Common//////////////////////////////////////////

	return nil
}

func CreateUser_Firebase(ctx context.Context, application dto.User) error {
	_, _, err := dbConfig.DATABASE_FireBase.Collection("users").Add(ctx, application)
	if err != nil {
		return err
	}
	return nil
}

func CreateUser_Sql(ctx context.Context, request dto.User) error {
	// AutoMigrate will create the "users" table if it does not exist
	err := dbConfig.DATABASE_Sql.AutoMigrate(&dto.User{})
	if err != nil {
		log.Fatalf("Failed to migrate table: %v", err)
	}
	result := dbConfig.DATABASE_Sql.Create(&request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
