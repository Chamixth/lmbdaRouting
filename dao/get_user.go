package dao

import (
	"context"
	"lambdarouting/dbConfig"
	"lambdarouting/dto"
	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

func GetUserbyUserName(ctx context.Context, userName string) (*dto.User, error) {
	var request dto.User

	///////////////////////////////MongoDb//////////////////////////////

	filter := bson.M{"username": userName}
	err := dbConfig.DATABASE.Collection("User").FindOne(ctx, filter).Decode(&request)

	///////////////////////////////MongoDb//////////////////////////////

	//////////////////////////////////Firebase-firestore////////////////////////////

	// query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", userName).Limit(1)
	// docSnap, err := query.Documents(ctx).Next()

	//////////////////////////////////Firebase-firestore////////////////////////////

	////////////////////////////////////Common//////////////////////////////////////
	if err != nil {
		return nil, err
	}
	////////////////////////////////////Common//////////////////////////////////////

	//////////////////////////////////Firebase-firestore////////////////////////////
	// err = docSnap.DataTo(&request)
	// if err != nil {
	// 	return nil, err
	// }
	//////////////////////////////////Firebase-firestore////////////////////////////

	////////////////////////////////////Common//////////////////////////////////////
	return &request, nil
	////////////////////////////////////Common//////////////////////////////////////
}

func GetUserByUserName_Firebase(ctx context.Context, userName string) (*dto.User, error) {
	var request dto.User
	query := dbConfig.DATABASE_FireBase.Collection("users").Where("Username", "==", userName).Limit(1)
	docSnap, err := query.Documents(ctx).Next()
	if err != nil {
		return nil, err
	}
	err = docSnap.DataTo(&request)
	if err != nil {
		return nil, err
	}
	return &request, nil
}

func GetUser_Sql(ctx context.Context, userName string) (*dto.User, error) {
	var user dto.User
	result := dbConfig.DATABASE_Sql.Where("username = ?", userName).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil

}
