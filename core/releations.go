package core

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	mongodb "main/core/db/monodb"
	Services "main/features/Services"
	home_details "main/features/home_details"
	projects "main/features/projects_list"
	social_media "main/features/social_media"
)

///releations///

func InitDataBase() error {
	_, err := mongodb.DB.Collection("home_details").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	////////////////////////*************projects************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("projects_list").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	////////////////////////************services*************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("services").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	////////////////////////*************social media details************\\\\\\\\\\\\\\\\\\\\\\\\\\	////////////////////////*************************\\\\\\\\\\\\\\\\\\\\\\\\\\
	_, err = mongodb.DB.Collection("social_media").DeleteMany(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	home_details.InitData()
	projects.InitData()
	social_media.InitData()
	Services.InitData()

	// sqldb.AddRelation("users", "id", "addresses", "user_id")

	return nil
}
