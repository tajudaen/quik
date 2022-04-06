package app

import (
	"os"
	"testing"
)

//TODO:

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}



func TestGetWalletBalance(t *testing.T) {

	// response := httptest.NewRecorder()
	// c, _ := gin.CreateTestContext(response)

	
}

func TestCreditWallet(t *testing.T) {
	// Mock LocationService methods
	getCountryfunc = func(countryId string) (*location.Country, *errors.APIError) {
		return &location.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationService = &locationServiceMock{}

} 

func TestDebitWallet(t *testing.T) {
	// Mock LocationService methods
	getCountryfunc = func(countryId string) (*location.Country, *errors.APIError) {
		return &location.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationService = &locationServiceMock{}


} 