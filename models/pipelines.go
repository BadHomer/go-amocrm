package models

type Pipeline struct {
	ID           int              `json:"id,omitempty"`
	Name         string           `json:"name,omitempty"`
	Sort         int              `json:"sort,omitempty"`
	IsMain       bool             `json:"is_main,omitempty"`
	IsArchive    bool             `json:"is_archive,omitempty"`
	IsUnsortedOn bool             `json:"is_unsorted_on,omitempty"`
	AccountID    int              `json:"account_id,omitempty"`
	Embedded     PipelineEmbedded `json:"_embedded,omitempty"`
}

type PipelineEmbedded struct {
	Statuses []Status `json:"statuses,omitempty"`
}
