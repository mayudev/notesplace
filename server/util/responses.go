package util

var NotFound = "not found"
var NotebookCreated = "notebook created"
var RequestInvalid = "invalid request"
var InvalidProtectionLevel = "invalid protection level"
var InvalidTitle = "invalid title"
var Forbidden = "forbidden"
var Unauthorized = "unauthorized"
var InternalServerError = "internal server error"
var PasswordTooLong = "password too long"
var PasswordTooShort = "password too short"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
