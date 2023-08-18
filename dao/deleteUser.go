package dao

import (
	"context"
	"lambdarouting/dbConfig"
	"lambdarouting/dto"
	"gopkg.in/mgo.v2/bson"
)

func DeleteUser(ctx context.Context, userName string) error {

	/////////////////////////////////////MongoDb///////////////////////////////////////

	filter := bson.M{"username": userName}
	_, err := dbConfig.DATABASE.Collection("User").DeleteOne(ctx, filter)

	/////////////////////////////////////MongoDb///////////////////////////////////////

	////////////////////////////////////Firebase-firestore/////////////////////////////

	// query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", userName).Limit(1)
	// docSnap, err := query.Documents(ctx).Next()

	////////////////////////////////////Firebase-firestore/////////////////////////////

	////////////////////////////////////Common////////////////////////////////////////

	if err != nil {
		return err
	}

	////////////////////////////////////Common////////////////////////////////////////

	////////////////////////////////////Firebase-firestore/////////////////////////////

	// docID := docSnap.Ref.ID
	// _, err = dbConfig.DATABASE_FireBase.Collection("users").Doc(docID).Delete(ctx)

	// if err != nil {
	// 	return err
	// }

	////////////////////////////////////Firebase-firestore/////////////////////////////

	///////////////////////////////////////Common///////////////////////////////////////

	return nil

	///////////////////////////////////////Common///////////////////////////////////////
}

func DeleteUser_Firebase(ctx context.Context, userName string) error {
	query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", userName).Limit(1)
	docSnap, err := query.Documents(ctx).Next()

	if err != nil {
		return err
	}
	docID := docSnap.Ref.ID
	_, err = dbConfig.DATABASE_FireBase.Collection("users").Doc(docID).Delete(ctx)

	if err != nil {
		return err
	}
	return nil
}

func DeleteUser_Sql(ctx context.Context, userName string) error {
	result := dbConfig.DATABASE_Sql.Where("username = ?", userName).Delete(&dto.User{})
	if result.RowsAffected < 1 {
		return result.Error
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
