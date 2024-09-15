package ideatag_http_handler

import "net/http"

type IdeaTagHTTPHanlder interface {
	GetTags(w http.ResponseWriter, r *http.Request)
	GetTagById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}
