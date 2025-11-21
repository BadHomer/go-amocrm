package models

type User struct {
	ID                int          `json:"id,omitempty"`
	Name              string       `json:"name,omitempty"`
	Email             string       `json:"email,omitempty"`
	Lang              string       `json:"lang,omitempty"`
	Rights            string       `json:"rights,omitempty"`
	CatalogRights     string       `json:"catalog_rights,omitempty"`
	CustomFieldRights string       `json:"custom_field_rights,omitempty"`
	IsAdmin           bool         `json:"is_admin,omitempty"`
	IsFree            bool         `json:"is_free,omitempty"`
	IsActive          bool         `json:"is_active,omitempty"`
	GroupID           int          `json:"group_id,omitempty"`
	RoleID            int          `json:"role_id,omitempty"`
	Embedded          UserEmbedded `json:"_embedded,omitempty"`
}

type UserEmbedded struct {
	Groups []Group `json:"groups,omitempty"`
}
