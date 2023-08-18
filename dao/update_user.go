package dao

import (
	"context"
	"lambdarouting/dbConfig"
	"lambdarouting/dto"

	"gopkg.in/mgo.v2/bson"
)

func UpdateUser(ctx context.Context, application dto.User) error {

	/////////////////////////////////MongoDb///////////////////////////
	filter := bson.M{"username": application.Username}
	update := bson.M{
		"$set": bson.M{
			"email":    application.Email,
			"password": application.Password,
		},
	}
	_, err := dbConfig.DATABASE.Collection("User").UpdateOne(ctx, filter, update)
	///////////////////////MongoDb///////////////////////////////

	///////////////////////////Firebase-firestore////////////////////////
	// query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", application.Username).Limit(1)
	// docSnap, err := query.Documents(ctx).Next()
	///////////////////////////Firebase-firestore////////////////////////

	///////////////////////////////Common//////////////////////////////
	if err != nil {
		return err
	}
	///////////////////////////////Common//////////////////////////////

	///////////////////////////Firebase-firestore////////////////////////
	// docID := docSnap.Ref.ID
	// _, err = dbConfig.DATABASE_FireBase.Collection("users").Doc(docID).Set(ctx, application)
	// if err != nil {
	// 	return err
	// }
	///////////////////////////Firebase-firestore////////////////////////

	///////////////////////////////Common//////////////////////////////
	return nil
	///////////////////////////////Common//////////////////////////////

}

func UpdateUser_Firebase(ctx context.Context, application dto.User) error {
	query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", application.Username).Limit(1)
	docSnap, err := query.Documents(ctx).Next()
	if err != nil {
		return err
	}
	docID := docSnap.Ref.ID
	_, err = dbConfig.DATABASE_FireBase.Collection("users").Doc(docID).Set(ctx, application)
	if err != nil {
		return err
	}
	return err
}

func UpdateUser_sql(ctx context.Context, request dto.User) error {
	result := dbConfig.DATABASE_Sql.Model(&dto.User{}).Where("username = ?", request.Username).Updates(&request)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
