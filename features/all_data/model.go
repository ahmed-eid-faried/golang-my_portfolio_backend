package AllData

type KData struct {
	SocialMedia  interface{} `json:"social_media"`
	ProjectsList interface{} `json:"projects_list"`
	HomeDetails  interface{} `json:"home_detials"`
	Services     interface{} `json:"services"`
	// SocialMedia  []social_media.SocialMedia `json:"social_media"`
	// ProjectsList []projects.Project         `json:"projects_list"`
	// HomeDetails  []home_details.HomeDetails `json:"home_detials"`
	// Services     []services.Service         `json:"services"`
}
