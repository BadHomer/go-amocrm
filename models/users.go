package models

type User struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Lang              string `json:"lang"`
	Rights            string `json:"rights"`
	CatalogRights     string `json:"catalog_rights"`
	CustomFieldRights string `json:"custom_field_rights"`
	IsAdmin           bool   `json:"is_admin"`
	IsFree            bool   `json:"is_free"`
	IsActive          bool   `json:"is_active"`
	GroupID           int    `json:"group_id"`
	RoleID            int    `json:"role_id"`
}
