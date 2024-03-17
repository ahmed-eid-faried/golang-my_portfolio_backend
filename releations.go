package main

import (
	// sqldb "main/core/db/sql"
	mongodb "main/core/db/monodb"
	// sqldb "main/core/db/sql"
	Services "main/features/Services"
	home_details "main/features/home_details"
	projects "main/features/projects_list"
	social_media "main/features/social_media"
)

///releations///

func InitDataBase() error {
	// sqldb.Init()
	mongodb.InitMongoDB()
	home_details.InitData()
	projects.InitData()
	social_media.InitData()
	Services.InitData()

	// sqldb.AddRelation("users", "id", "addresses", "user_id")

	return nil
}
