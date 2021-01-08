package controller

import (
	"encoding/json"
	"github.com/SKFE396/search-service/service"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"sync"
)

type ServerResponse struct {
	Code   int8        `json:"code"`
	Result interface{} `json:"result"`
}

type SearchController struct {
	searchService service.SearchService
}

func (s *SearchController) SetSearchService(searchService service.SearchService) {
	s.searchService = searchService
}

func (s *SearchController) Init(group *sync.WaitGroup, searchService service.SearchService) (server *http.Server) {
	s.searchService = searchService
	err := s.searchService.Init(nil, nil)
	if err != nil {
		log.Info(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/searchQuestions", s.SearchQuestions)
	mux.HandleFunc("/searchAnswers", s.SearchAnswers)
	mux.HandleFunc("/searchUsers", s.SearchUsers)
	mux.HandleFunc("/hotlist", s.HotList)
	mux.HandleFunc("/search", s.Search)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		Debug:            true,
	})
	handler := c.Handler(mux)
	go func() {
		defer group.Done()
		if err := http.ListenAndServe(":9094", handler); err != http.ErrServerClosed {
			log.Info(err)
		}
	}()
	return server
}

func (s *SearchController) Destruct() {
	s.searchService.Destruct()
}

func (s *SearchController) SearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 32)
			text := r.FormValue("text")
			token := r.Header.Get("Authorization")
			if pageErr == nil {
				code, result := s.searchService.SearchQuestions(token, page, text)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (s *SearchController) SearchAnswers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 32)
			text := r.FormValue("text")
			token := r.Header.Get("Authorization")
			if pageErr == nil {
				code, result := s.searchService.SearchAnswers(token, page, text)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (s *SearchController) SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			page, pageErr := strconv.ParseInt(r.FormValue("page"), 10, 32)
			text := r.FormValue("text")
			token := r.Header.Get("Authorization")
			if pageErr == nil {
				code, result := s.searchService.SearchUsers(token, page, text)
				response.Code = code
				response.Result = result
				object, _ := json.Marshal(response)
				_, _ = w.Write(object)
				return
			}
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (s *SearchController) HotList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	if r.Method == "GET" {
		token := r.Header.Get("Authorization")
		code, result := s.searchService.HotList(token)
		response.Code = code
		response.Result = result
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}

func (s *SearchController) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var response ServerResponse
	var err error
	if r.Method == "GET" {
		err = r.ParseForm()
		if err == nil {
			text := r.FormValue("text")
			token := r.Header.Get("Authorization")
			code, result := s.searchService.Search(token, text)
			response.Code = code
			response.Result = result
			object, _ := json.Marshal(response)
			_, _ = w.Write(object)
			return
		}
		log.Info(err)
		response.Code = 1
		object, _ := json.Marshal(response)
		_, _ = w.Write(object)
		return
	}
}
