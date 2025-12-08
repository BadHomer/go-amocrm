package models

type User struct {
	ID       int          `json:"id,omitempty"`
	Name     string       `json:"name,omitempty"`
	Email    string       `json:"email,omitempty"`
	Lang     string       `json:"lang,omitempty"`
	GroupID  int          `json:"group_id,omitempty"`
	RoleID   int          `json:"role_id,omitempty"`
	Embedded UserEmbedded `json:"_embedded,omitempty"`
	Rights   RightEntity  `json:"rights,omitempty"`
}

type UserEmbedded struct {
	Groups []Group `json:"groups,omitempty"`
}

type RightEntity struct {
	View     string `json:"view,omitempty"`
	Edit     string `json:"edit,omitempty"`
	Add      string `json:"add,omitempty"`
	Delete   string `json:"delete,omitempty"`
	Export   string `json:"export,omitempty"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
	IsFree   bool   `json:"is_free,omitempty"`
	IsActive bool   `json:"is_active,omitempty"`
}
