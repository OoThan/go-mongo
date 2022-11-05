package model

type Role struct {
	RoleId   string `json:"_id" bson:"_id"`
	RoleName string `json:"role_name" bson:"role"`
}

type CreateRole struct {
	RoleName string `json:"role_name" bson:"role"`
}

type UpdateRole struct {
	RoleName string `json:"role_name" bson:"role"`
}
