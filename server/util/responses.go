package util

var NotFound = "not found"
var NoteNotFound = "note not found"
var NotebookNotFound = "notebook not found"
var NotebookCreated = "notebook created"
var RequestInvalid = "invalid request"
var InvalidProtectionLevel = "invalid protection level"
var InvalidTitle = "invalid title"
var Forbidden = "forbidden"
var Unauthorized = "unauthorized"
var InternalServerError = "internal server error"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
