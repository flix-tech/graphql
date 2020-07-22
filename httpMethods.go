package graphql

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/context/ctxhttp"
)

// GetWithQueryString sends an http get request with the query and variables as a query string
func GetWithQueryString(ctx context.Context, client *http.Client, graphqlURL string, query string, variables map[string]interface{}) (*http.Response, error) {
	queryString := url.QueryEscape(query)
	variableBytes, err := json.Marshal(variables)
	if err != nil {
		return &http.Response{}, err
	}
	variableString := url.QueryEscape(string(variableBytes))
	resp, err := ctxhttp.Get(ctx, client, graphqlURL+`?query=`+queryString+`&variables=`+variableString)
	return resp, err
}

func PostWithBearerToken(ctx context.Context, client *http.Client, url string, bodyType string, body io.Reader, token string) (*http.Response, error) {
	var resp *http.Response
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return resp, err
	}
	req.Header.Set("Content-Type", "application/json")
	if token != `` {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err = ctxhttp.Do(ctx, client, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
