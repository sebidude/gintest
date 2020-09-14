package gintest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
)

/* ############################################## */
/* ###               Gin Contexts             ### */
/* ############################################## */

// PrepareEmptyContext prepares a gin testing context with initialized header.
func PrepareEmptyContext() *gin.Context {
	c, _ := PrepareEmptyRecordingContext()
	return c
}

// PrepareEmptyRecordingContext prepares a gin testing context with initialized header.
func PrepareEmptyRecordingContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = &http.Request{ContentLength: 0, Header: make(map[string][]string)}
	return c, recorder
}

func makeURL(path string) *url.URL {
	urlObj, err := url.Parse(path)
	if err != nil {
		panic(err)
	}
	return urlObj
}

// PrepareRoute sets up the given request in route in a gin context.
func PrepareRoute(c *gin.Context, path, method string) {
	c.Request.URL = makeURL(path)
	c.Request.Method = method
}

/* ############################################## */
/* ###              HTTP Requests             ### */
/* ############################################## */

// GetStatusCode performs a HTTP request to the given url using the given method and returns the HTTP response code.
func GetStatusCode(url, method string) int {
	var response *http.Response
	var err error
	if method == "POST" {
		response, err = http.Post(url, "text/plain", bytes.NewBuffer([]byte("")))
	} else {
		response, err = http.Get(url)
	}
	if err != nil {
		panic(err)
	}
	return response.StatusCode
}
