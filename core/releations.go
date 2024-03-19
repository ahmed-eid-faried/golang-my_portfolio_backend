package core

import (
	Services "main/features/Services"
	home_details "main/features/home_details"
	projects "main/features/projects_list"
	social_media "main/features/social_media"
)

///releations///

func InitDataBase() {

	home_details.InitData()
	social_media.InitData()
	projects.InitData()
	Services.InitData()

	// sqldb.AddRelation("users", "id", "addresses", "user_id")

}
