package main

import (
	"log"
	"net/http"
	"os"
	"reflect"
)

var errorMap = make(map[string]int)

type ErrorBadRequest struct {
	Message string `json:"message"`
}

func (e ErrorBadRequest) Error() string {
	return e.Message
}

type ErrorUnauthorized struct {
	Message string `json:"message"`
}

func (e ErrorUnauthorized) Error() string {
	return e.Message
}

type ErrorPaymentRequired struct {
	Message string `json:"message"`
}

func (e ErrorPaymentRequired) Error() string {
	return e.Message
}

type ErrorForbidden struct {
	Message string `json:"message"`
}

func (e ErrorForbidden) Error() string {
	return e.Message
}

type ErrorNotFound struct {
	Message string `json:"message"`
}

func (e ErrorNotFound) Error() string {
	return e.Message
}

type ErrorMethodNotAllowed struct {
	Message string `json:"message"`
}

func (e ErrorMethodNotAllowed) Error() string {
	return e.Message
}

type ErrorNotAcceptable struct {
	Message string `json:"message"`
}

func (e ErrorNotAcceptable) Error() string {
	return e.Message
}

type ErrorProxyAuthRequired struct {
	Message string `json:"message"`
}

func (e ErrorProxyAuthRequired) Error() string {
	return e.Message
}

type ErrorRequestTimeout struct {
	Message string `json:"message"`
}

func (e ErrorRequestTimeout) Error() string {
	return e.Message
}

type ErrorConflict struct {
	Message string `json:"message"`
}

func (e ErrorConflict) Error() string {
	return e.Message
}

type ErrorGone struct {
	Message string `json:"message"`
}

func (e ErrorGone) Error() string {
	return e.Message
}

type ErrorLengthRequired struct {
	Message string `json:"message"`
}

func (e ErrorLengthRequired) Error() string {
	return e.Message
}

type ErrorPreconditionFailed struct {
	Message string `json:"message"`
}

func (e ErrorPreconditionFailed) Error() string {
	return e.Message
}

type ErrorRequestEntityTooLarge struct {
	Message string `json:"message"`
}

func (e ErrorRequestEntityTooLarge) Error() string {
	return e.Message
}

type ErrorRequestURITooLong struct {
	Message string `json:"message"`
}

func (e ErrorRequestURITooLong) Error() string {
	return e.Message
}

type ErrorUnsupportedMediaType struct {
	Message string `json:"message"`
}

func (e ErrorUnsupportedMediaType) Error() string {
	return e.Message
}

type ErrorRequestedRangeNotSatisfiable struct {
	Message string `json:"message"`
}

func (e ErrorRequestedRangeNotSatisfiable) Error() string {
	return e.Message
}

type ErrorExpectationFailed struct {
	Message string `json:"message"`
}

func (e ErrorExpectationFailed) Error() string {
	return e.Message
}

type ErrorTeapot struct {
	Message string `json:"message"`
}

func (e ErrorTeapot) Error() string {
	return e.Message
}

type ErrorMisdirectedRequest struct {
	Message string `json:"message"`
}

func (e ErrorMisdirectedRequest) Error() string {
	return e.Message
}

type ErrorUnprocessableEntity struct {
	Message string `json:"message"`
}

func (e ErrorUnprocessableEntity) Error() string {
	return e.Message
}

type ErrorLocked struct {
	Message string `json:"message"`
}

func (e ErrorLocked) Error() string {
	return e.Message
}

type ErrorTooEarly struct {
	Message string `json:"message"`
}

func (e ErrorTooEarly) Error() string {
	return e.Message
}

type ErrorTooManyRequests struct {
	Message string `json:"message"`
}

func (e ErrorTooManyRequests) Error() string {
	return e.Message
}

type ErrorUpgradeRequired struct {
	Message string `json:"message"`
}

func (e ErrorUpgradeRequired) Error() string {
	return e.Message
}

type ErrorPreconditionRequired struct {
	Message string `json:"message"`
}

func (e ErrorPreconditionRequired) Error() string {
	return e.Message
}

type ErrorRequestHeaderFieldsTooLarge struct {
	Message string `json:"message"`
}

func (e ErrorRequestHeaderFieldsTooLarge) Error() string {
	return e.Message
}

type ErrorUnavailableForLegalReasons struct {
	Message string `json:"message"`
}

func (e ErrorUnavailableForLegalReasons) Error() string {
	return e.Message
}

type ErrorInternalServer struct {
	Message string `json:"message"`
}

func (e ErrorInternalServer) Error() string {
	return e.Message
}

type ErrorNotImplemented struct {
	Message string `json:"message"`
}

func (e ErrorNotImplemented) Error() string {
	return e.Message
}

func Init() {

	errorMap[reflect.TypeOf(ErrorBadRequest{}).String()] = http.StatusBadRequest
	errorMap[reflect.TypeOf(ErrorUnauthorized{}).String()] = http.StatusUnauthorized
	errorMap[reflect.TypeOf(ErrorPaymentRequired{}).String()] = http.StatusPaymentRequired
	errorMap[reflect.TypeOf(ErrorForbidden{}).String()] = http.StatusForbidden
	errorMap[reflect.TypeOf(ErrorNotFound{}).String()] = http.StatusNotFound
	errorMap[reflect.TypeOf(ErrorMethodNotAllowed{}).String()] = http.StatusMethodNotAllowed
	errorMap[reflect.TypeOf(ErrorNotAcceptable{}).String()] = http.StatusNotAcceptable
	errorMap[reflect.TypeOf(ErrorProxyAuthRequired{}).String()] = http.StatusProxyAuthRequired
	errorMap[reflect.TypeOf(ErrorRequestTimeout{}).String()] = http.StatusRequestTimeout
	errorMap[reflect.TypeOf(ErrorConflict{}).String()] = http.StatusConflict
	errorMap[reflect.TypeOf(ErrorGone{}).String()] = http.StatusGone
	errorMap[reflect.TypeOf(ErrorLengthRequired{}).String()] = http.StatusLengthRequired
	errorMap[reflect.TypeOf(ErrorPreconditionFailed{}).String()] = http.StatusPreconditionFailed
	errorMap[reflect.TypeOf(ErrorRequestEntityTooLarge{}).String()] = http.StatusRequestEntityTooLarge
	errorMap[reflect.TypeOf(ErrorRequestURITooLong{}).String()] = http.StatusRequestURITooLong
	errorMap[reflect.TypeOf(ErrorUnsupportedMediaType{}).String()] = http.StatusUnsupportedMediaType
	errorMap[reflect.TypeOf(ErrorRequestedRangeNotSatisfiable{}).String()] = http.StatusRequestedRangeNotSatisfiable
	errorMap[reflect.TypeOf(ErrorExpectationFailed{}).String()] = http.StatusExpectationFailed
	errorMap[reflect.TypeOf(ErrorTeapot{}).String()] = http.StatusTeapot
	errorMap[reflect.TypeOf(ErrorMisdirectedRequest{}).String()] = http.StatusMisdirectedRequest
	errorMap[reflect.TypeOf(ErrorUnprocessableEntity{}).String()] = http.StatusUnprocessableEntity
	errorMap[reflect.TypeOf(ErrorLocked{}).String()] = http.StatusLocked
	errorMap[reflect.TypeOf(ErrorTooEarly{}).String()] = http.StatusTooEarly
	errorMap[reflect.TypeOf(ErrorUpgradeRequired{}).String()] = http.StatusUpgradeRequired
	errorMap[reflect.TypeOf(ErrorPreconditionRequired{}).String()] = http.StatusPreconditionRequired
	errorMap[reflect.TypeOf(ErrorTooManyRequests{}).String()] = http.StatusTooManyRequests
	errorMap[reflect.TypeOf(ErrorPreconditionRequired{}).String()] = http.StatusPreconditionRequired
	errorMap[reflect.TypeOf(ErrorRequestHeaderFieldsTooLarge{}).String()] = http.StatusRequestHeaderFieldsTooLarge
	errorMap[reflect.TypeOf(ErrorUnavailableForLegalReasons{}).String()] = http.StatusUnavailableForLegalReasons
	errorMap[reflect.TypeOf(ErrorInternalServer{}).String()] = http.StatusInternalServerError
	errorMap[reflect.TypeOf(ErrorInternalServer{}).String()] = http.StatusNotImplemented

}

func DetermineError(error interface{}) (int, interface{}) {
	file, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	logger := log.New(file, "", log.LstdFlags)
	errType := reflect.TypeOf(error).String()
	status, exists := errorMap[errType]
	if !exists {
		status = http.StatusInternalServerError
	}
	logger.Printf("Error: %v", error)
	return status, error
}
