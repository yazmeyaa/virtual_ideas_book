package ideatag_http_handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/yazmeyaa/virtual_ideas_book/internal/common"
	ideatag "github.com/yazmeyaa/virtual_ideas_book/internal/repository/idea_tag"
	ideatag_service "github.com/yazmeyaa/virtual_ideas_book/internal/service/idea_tag"
	idea_request "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea_tag/dto/request"
	tag_response "github.com/yazmeyaa/virtual_ideas_book/internal/transport/http/idea_tag/dto/response"
)

type ideaTagHTTPHanlder struct {
	tagService ideatag_service.TagService
	log        *slog.Logger
}

// Create implements IdeaTagHTTPHanlder.
func (i *ideaTagHTTPHanlder) Create(w http.ResponseWriter, r *http.Request) {
	const fnCreateTag = "IdeaTagHTTPHanlder.Create"
	data := idea_request.CreateTagRequestDTO{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, common.UNREACHABLE_BODY_ERROR)
		return
	}

	tag, err := i.tagService.CreateTag(ideatag.CreateTagDTO{
		Name:        data.Name,
		Description: data.Description,
	})

	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok {
			if err.Code == pgerrcode.UniqueViolation {
				common.WriteError(w, http.StatusBadRequest, "Idea tag with same name already exist")
			}
		} else {
			common.WriteError(w, http.StatusBadRequest, common.UNEXPECTED_ERROR)
		}

		i.log.Error(fmt.Sprintf("[%s]: %s", fnCreateTag, err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag_response.CreateTagResponseDTO{
		Name:        tag.Name,
		Description: tag.Description,
	})
}

// Delete implements IdeaTagHTTPHanlder.
func (i *ideaTagHTTPHanlder) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// GetTagById implements IdeaTagHTTPHanlder.
func (i *ideaTagHTTPHanlder) GetTagById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam, has := vars["id"]
	if !has {
		common.WriteError(w, http.StatusInternalServerError, common.UNEXPECTED_ERROR)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		common.WriteError(w, http.StatusBadRequest, "Unexpected value in `id` param.")
		return
	}

	tag, err := i.tagService.GetTagByID(uint(id))
	if err != nil {
		common.WriteError(w, http.StatusNotFound, "Tag with provided id was not found.")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tag_response.GetTtagByIdResponseDTO{
		Name:        tag.Name,
		Description: tag.Description,
	})
}

// GetTags implements IdeaTagHTTPHanlder.
func (i *ideaTagHTTPHanlder) GetTags(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// Update implements IdeaTagHTTPHanlder.
func (i *ideaTagHTTPHanlder) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

func New(
	tagService ideatag_service.TagService,
	log *slog.Logger,
) IdeaTagHTTPHanlder {
	return &ideaTagHTTPHanlder{
		tagService: tagService,
		log:        log,
	}
}
