package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	req, _ := http.NewRequest("POST", "/pubsub", nil)
	rr := httptest.NewRecorder()

	e := echo.New()
	ctx := e.NewContext(req, rr)

	if assert.NoError(t, pubsubHandler(ctx)) {
		assert.Equal(t, rr.Code, http.StatusOK)
	}
}
