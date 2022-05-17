package util

var NotebookNotFound = "notebook not found"
var NotebookCreated = "notebook created"
var RequestInvalid = "invalid request"
var InvalidProtectionLevel = "invalid protection level"
var InvalidTitle = "invalid title"
var Forbidden = "forbidden"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
