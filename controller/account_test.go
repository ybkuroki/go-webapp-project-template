package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ybkuroki/go-webapp-project-template/model"
	"github.com/ybkuroki/go-webapp-project-template/test"
)

func TestGetLoginStatus(t *testing.T) {
	router := test.Prepare()
	router.GET(APIAccountLoginStatus, GetLoginStatus())

	req := httptest.NewRequest("GET", APIAccountLoginStatus, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, "true", rec.Body.String())
}

func TestGetLoginAccount(t *testing.T) {
	router := test.Prepare()
	router.GET(APIAccountLoginAccount, GetLoginAccount())

	req := httptest.NewRequest("GET", APIAccountLoginAccount, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	account := model.NewAccountWithPlainPassword("test", "test", model.NewAuthority("Admin"))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(account), rec.Body.String())
}
