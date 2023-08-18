package dao

import (
	"context"
	"errors"
	"log"
	"lambdarouting/dbConfig"
	"lambdarouting/dto"

	"gopkg.in/mgo.v2/bson"
	"gorm.io/gorm"
)

func GetAllUser(ctx context.Context) ([]dto.User, error) {
	var results []dto.User
	cursor, err := dbConfig.DATABASE.Collection("User").Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var user dto.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, errors.New("Error When Decoding db Operation")
		}
		results = append(results, user)
	}

	return results, nil
}

func GetAllUser_Sql(ctx context.Context) (*[]dto.User, error) {
	var user []dto.User
	result := dbConfig.DATABASE_Sql.Find(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		return nil, result.Error
	}
	return &user, nil

}
