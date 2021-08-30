package app

type FindDocumentBody struct {
	Status     string                   `json:"status,omitempty"`
	Collection string                   `json:"collection" binding:"required"`
	Data       []map[string]interface{} `json:"data" binding:"required"`
}

type InsertDocumentBody struct {
	Status     string                 `json:"status,omitempty"`
	Collection string                 `json:"collection" binding:"required"`
	Data       map[string]interface{} `json:"data" binding:"required"`
}

type UpdateDocumentBody struct {
	Status      string                   `json:"status,omitempty"`
	Collection  string                   `json:"collection" binding:"required"`
	FindData    []map[string]interface{} `json:"filter" binding:"required"`
	Upsert      bool                     `json:"upsert"`
	UpdateField string                   `json:"field" binding:"required"`
	UpdateData  interface{}              `json:"data" binding:"required"`
}

type DeleteDocumentBody struct {
	Status     string                   `json:"status,omitempty"`
	Collection string                   `json:"collection" binding:"required"`
	FindData   []map[string]interface{} `json:"filter" binding:"required"`
	Data       []map[string]interface{} `json:"data,omitempty"`
}
