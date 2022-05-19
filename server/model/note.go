package model

import "time"

// Note Model
type Note struct {
	// The note ID
	//
	// Required. Expected to be an UUID.
	ID string `json:"id"`
	// ID of the notebook the note belongs to
	//
	// Required. Foreign key in database
	NotebookID string `json:"notebook_id"`
	// Note title
	//
	// Required.
	Title string `json:"title"`
	// Note order.
	//
	// Required.
	Order *uint `json:"order"`
	// Note contents
	//
	// Optional. Markdown formatted.
	Content string `json:"content"`
	// Created at timestamp.
	//
	// Required.
	CreatedAt time.Time `json:"created_at"`
	// Updated at timestamp.
	//
	// Required.
	UpdatedAt time.Time `json:"updated_at"`
}
