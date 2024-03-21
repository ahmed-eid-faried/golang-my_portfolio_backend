// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/data": {
            "get": {
                "description": "Retrieve Get All Data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "AllData"
                ],
                "summary": "Get all data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/AllData.KData"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all Data from MongoDB",
                "tags": [
                    "AllData"
                ],
                "summary": "Delete all Data",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/data/init": {
            "get": {
                "description": "init all Data from MongoDB",
                "tags": [
                    "AllData"
                ],
                "summary": "Put All Data",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/home_details": {
            "get": {
                "description": "Retrieve all home details from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Get all home details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/HomeDetails.HomeDetails"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new home detail in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Create a new home detail",
                "parameters": [
                    {
                        "description": "Home detail object",
                        "name": "homeDetail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/HomeDetails.HomeDetails"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ID of the created home detail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all home details from MongoDB",
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Delete all home details",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/home_details/search": {
            "get": {
                "description": "Search home details in MongoDB based on a keyword",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Search home details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/HomeDetails.HomeDetails"
                            }
                        }
                    }
                }
            }
        },
        "/home_details/{id}": {
            "get": {
                "description": "Retrieve a home detail by its ID from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Get a home detail by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Home detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/HomeDetails.HomeDetails"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing home detail in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Update a home detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Home detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated home detail object",
                        "name": "homeDetail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/HomeDetails.HomeDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a home detail by its ID from MongoDB",
                "tags": [
                    "HomeDetails"
                ],
                "summary": "Delete a home detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Home detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "description": "Retrieve all projects from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get all projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Projects.Project"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new project in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "description": "Project object",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Projects.Project"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ID of the created project",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all projects from MongoDB",
                "tags": [
                    "Projects"
                ],
                "summary": "Delete all projects",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/projects/search": {
            "get": {
                "description": "Search projects in MongoDB based on a keyword",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Search projects",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Projects.Project"
                            }
                        }
                    }
                }
            }
        },
        "/projects/{id}": {
            "get": {
                "description": "Retrieve a project by its ID from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Get a project by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Projects.Project"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing project in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Projects"
                ],
                "summary": "Update a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated project object",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Projects.Project"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a project by its ID from MongoDB",
                "tags": [
                    "Projects"
                ],
                "summary": "Delete a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/services": {
            "get": {
                "description": "Get all Services",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Get All Services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Services.Service"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a Service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Create a Service",
                "parameters": [
                    {
                        "description": "Service object",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Services.Service"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            },
            "delete": {
                "description": "Remove all Services from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Remove all Services",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/services/search": {
            "get": {
                "description": "Search services",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Search services",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Services.Service"
                            }
                        }
                    }
                }
            }
        },
        "/services/{id}": {
            "get": {
                "description": "Get a Service by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Get a Service by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Services.Service"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a Service by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Services"
                ],
                "summary": "Update a Service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service object",
                        "name": "service",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Services.Service"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a Service by ID",
                "tags": [
                    "Services"
                ],
                "summary": "Delete a Service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Service ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/social_media": {
            "get": {
                "description": "Retrieve all social media details from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Get all social media details",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/SocialMedia.SocialMedia"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new social media detail in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Create a new social media detail",
                "parameters": [
                    {
                        "description": "Social media detail object",
                        "name": "socialMediaDetail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SocialMedia.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "ID of the created social media detail",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all social media details from MongoDB",
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Delete all social media details",
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/social_media/search": {
            "get": {
                "description": "Search social media details in MongoDB based on a keyword",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Search social media details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search keyword",
                        "name": "keyword",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/SocialMedia.SocialMedia"
                            }
                        }
                    }
                }
            }
        },
        "/social_media/{id}": {
            "get": {
                "description": "Retrieve a social media detail by its ID from MongoDB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Get a social media detail by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Social media detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SocialMedia.SocialMedia"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing social media detail in MongoDB",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Update a social media detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Social media detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated social media detail object",
                        "name": "socialMediaDetail",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SocialMedia.SocialMedia"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "Delete a social media detail by its ID from MongoDB",
                "tags": [
                    "SocialMediaDetails"
                ],
                "summary": "Delete a social media detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Social media detail ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "AllData.KData": {
            "type": "object",
            "properties": {
                "home_detials": {},
                "projects_list": {},
                "services": {},
                "social_media": {}
            }
        },
        "HomeDetails.HomeDetails": {
            "type": "object",
            "required": [
                "hd_aboutimage",
                "hd_aboutmedesc",
                "hd_aboutmename",
                "hd_cv",
                "hd_desc",
                "hd_image",
                "hd_name"
            ],
            "properties": {
                "hd_aboutimage": {
                    "type": "string",
                    "example": "http://example.com/image.jpg"
                },
                "hd_aboutmedesc": {
                    "type": "string",
                    "example": "I am a software engineer."
                },
                "hd_aboutmename": {
                    "type": "string",
                    "example": "About Me"
                },
                "hd_cv": {
                    "type": "string",
                    "example": "http://example.com/cv.pdf"
                },
                "hd_desc": {
                    "type": "string",
                    "example": "Description of my home details"
                },
                "hd_id": {
                    "type": "string"
                },
                "hd_image": {
                    "type": "string",
                    "example": "http://example.com/image.jpg"
                },
                "hd_name": {
                    "type": "string",
                    "example": "My Home Details"
                }
            }
        },
        "Projects.Project": {
            "type": "object",
            "required": [
                "pl_appstore",
                "pl_body",
                "pl_cli",
                "pl_doc",
                "pl_embedded",
                "pl_github",
                "pl_googleplay",
                "pl_image",
                "pl_linux",
                "pl_macos",
                "pl_package",
                "pl_title",
                "pl_web",
                "pl_windows"
            ],
            "properties": {
                "pl_appstore": {
                    "type": "string",
                    "example": "http://example.com/appstore"
                },
                "pl_body": {
                    "type": "string",
                    "example": "Description of my project"
                },
                "pl_cli": {
                    "type": "string",
                    "example": "http://example.com/cli"
                },
                "pl_doc": {
                    "type": "string",
                    "example": "http://example.com/doc"
                },
                "pl_embedded": {
                    "type": "string",
                    "example": "http://example.com/embedded"
                },
                "pl_github": {
                    "type": "string",
                    "example": "http://example.com/github"
                },
                "pl_googleplay": {
                    "type": "string",
                    "example": "http://example.com/googleplay"
                },
                "pl_id": {
                    "type": "string"
                },
                "pl_image": {
                    "type": "string",
                    "example": ""
                },
                "pl_linux": {
                    "type": "string",
                    "example": "http://example.com/linux"
                },
                "pl_macos": {
                    "type": "string",
                    "example": "http://example.com/macos"
                },
                "pl_package": {
                    "type": "string",
                    "example": "http://example.com/package"
                },
                "pl_title": {
                    "type": "string",
                    "example": "My Project"
                },
                "pl_web": {
                    "type": "string",
                    "example": "http://example.com/web"
                },
                "pl_windows": {
                    "type": "string",
                    "example": "http://example.com/windows"
                }
            }
        },
        "Services.Service": {
            "type": "object",
            "required": [
                "services_assets",
                "services_body",
                "services_title",
                "services_type"
            ],
            "properties": {
                "services_assets": {
                    "type": "string",
                    "example": "http://example.com/assets/service.jpg"
                },
                "services_body": {
                    "type": "string",
                    "example": "We specialize in building responsive and user-friendly websites."
                },
                "services_id": {
                    "type": "string"
                },
                "services_title": {
                    "type": "string",
                    "example": "Web Development"
                },
                "services_type": {
                    "type": "string",
                    "example": "Development"
                }
            }
        },
        "SocialMedia.SocialMedia": {
            "type": "object",
            "required": [
                "sm_cv",
                "sm_email",
                "sm_facebook",
                "sm_github",
                "sm_instagram",
                "sm_linkedin",
                "sm_medium",
                "sm_twitter",
                "sm_whatsapp"
            ],
            "properties": {
                "sm_cv": {
                    "type": "string",
                    "example": "ahmed_mady_cv.pdf"
                },
                "sm_email": {
                    "type": "string",
                    "example": "ahmed.mady@example.com"
                },
                "sm_facebook": {
                    "type": "string",
                    "example": "ahmed.mady.facebook"
                },
                "sm_github": {
                    "type": "string",
                    "example": "ahmed_mady_github"
                },
                "sm_id": {
                    "type": "string"
                },
                "sm_instagram": {
                    "type": "string",
                    "example": "ahmed_mady_instagram"
                },
                "sm_linkedin": {
                    "type": "string",
                    "example": "john-doe-linkedin"
                },
                "sm_medium": {
                    "type": "string",
                    "example": "@ahmed_mady_medium"
                },
                "sm_twitter": {
                    "type": "string",
                    "example": "@ahmed_mady_twitter"
                },
                "sm_whatsapp": {
                    "type": "string",
                    "example": "123456789"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "golang-my-portfolio-backend.onrender.com",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "User API",
	Description:      "API for user management",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
