package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpResponse struct {
	Writer http.ResponseWriter
}

func (h *HttpResponse) Success(data interface{}, statusCode int) {
	type successResponse struct {
		Success bool `json:"success"`
		Data interface{} `json:"data"`
	}

	response, err := json.Marshal(successResponse{true, data})

	if err != nil {
		log.Printf("Invalid response data. Data: %s --- Error: %s\n", data, err)
		return
	}

	h.write(response, statusCode)
}

func (h *HttpResponse) Error(message string, statusCode int) {
	type errorResponse struct {
		Success bool `json:"success"`
		Error string `json:"error"`
	}

	if statusCode > 499 {
		log.Println("Internal Server Error", message)
	}

	response, _ := json.Marshal(errorResponse{false, message})
	h.write(response, statusCode)
}

func (h *HttpResponse) write(data []byte, statusCode int) {
	h.Writer.Header().Add("Content-Type", "application/json")
	h.Writer.WriteHeader(statusCode)
	h.Writer.Write(data)
}
