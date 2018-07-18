package main

import(
	"net/http"
	"testing"
)

type ServiceResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func TestTestService_Call(t *testing.T) {

	testCases := []struct {
		method_used string
		api            string
		expectedCode   int
		expectedStatus string

	}{
		{"GET","/", http.StatusNotFound, "ERROR"},
		{"POST","/UpdateLocation", http.StatusOK, "SUCCESS"},
		{"GET","/GetLocation", http.StatusOK, "SUCCESS"},
		{"GET","/GetLocation?key=1",http.StatusOK,"SUCCESS"},
	}
	for _, tc := range testCases {
		s := TestService{}
		response := s.Call(tc.api, &http.Request{})
		if response.Code != tc.expectedCode {
			t.Errorf("Error in code. Expected: %v. Got: %v", tc.expectedCode, response.Code)
		}
		if response.Status != tc.expectedStatus {
			t.Errorf("Error in code. Expected: %s. Got: %s", tc.expectedStatus, response.Status)
		}
	}
}


type TestService struct{}

func (t TestService) Call(api string, r *http.Request) ServiceResponse {

	switch api {
	case "/GetLocation":
		return ServiceResponse{200, "SUCCESS", "OK", nil}
	case "/UpdateLocation":
		return ServiceResponse{200, "SUCCESS", "OK", nil}
	case "/GetLocation?key=1":
		return ServiceResponse{200,"SUCCESS","OK",nil}
	default:
		return ServiceResponse{404, "ERROR", "This is an invalid API.", nil}
	}
}