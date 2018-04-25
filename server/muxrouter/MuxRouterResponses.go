package muxrouter

import (
	"net/http"
	"encoding/json"
	"io"
)

type JSON map[string]interface{}

type JSONResponseWriter struct {
	http.ResponseWriter
	b []byte
}

func (j JSONResponseWriter) Send() {
	if string(j.b) != "" {
		io.WriteString(j.ResponseWriter, string(j.b))
	}
}

func errorsToJSON(err error) JSON {
	return JSON{"errors": JSON{"global": err.Error()}}
}

// Success => 200
func (j JSONResponseWriter) Success(i JSON) {
	j.WriteHeader(http.StatusOK)
	j.b, _ = json.Marshal(i)
	j.Send()
}

/*
response error = {
	"errors": {
		"global": `error`
	}
}
 */
// Not found 404
func (j JSONResponseWriter) Error404(err error) {
	j.WriteHeader(http.StatusNotFound)
	j.b, _ = json.Marshal(errorsToJSON(err))
	j.Send()
}

// Unauthorized
func (j JSONResponseWriter) Error401(err error) {
	j.WriteHeader(http.StatusUnauthorized)
	j.b, _ = json.Marshal(errorsToJSON(err))
	j.Send()
}
