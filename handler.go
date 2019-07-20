package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/emicklei/go-restful"
)

type restHandler interface {
	all(req *restful.Request, res *restful.Response)
	create(req *restful.Request, res *restful.Response)
	get(req *restful.Request, res *restful.Response)
	update(req *restful.Request, res *restful.Response)
	delete(req *restful.Request, res *restful.Response)

	regsiter(container *restful.Container)
}

type bookHandler struct {
	store bookStore
}

func newBookHandler(store bookStore) (handler restHandler) {
	handler = &bookHandler{store: store}
	return
}

func (handler *bookHandler) all(req *restful.Request, res *restful.Response) {
	books, e := handler.store.all()
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusInternalServerError, "Internal handler error")
		return
	}
	res.WriteEntity(books)
}

func (handler *bookHandler) create(req *restful.Request, res *restful.Response) {
	d := json.NewDecoder(req.Request.Body)
	var book Book
	e := d.Decode(&book)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusBadRequest, "Bad Request")
		return
	}
	id, e := handler.store.create(book)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusInternalServerError, "Internal handler error")
		return
	}
	book.ID = id
	res.WriteHeaderAndEntity(http.StatusCreated, book)
}

func (handler *bookHandler) get(req *restful.Request, res *restful.Response) {
	var id int64
	i, e := strconv.Atoi(req.PathParameter("id"))
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusBadRequest, "Bad Request")
		return
	}
	id = int64(i)
	book, e := handler.store.get(id)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusInternalServerError, "Internal handler error")
		return
	}

	res.WriteHeaderAndEntity(http.StatusOK, book)
}

func (handler *bookHandler) update(req *restful.Request, res *restful.Response) {
	d := json.NewDecoder(req.Request.Body)
	var book Book
	e := d.Decode(&book)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusBadRequest, "Bad Request")
		return
	}
	i, e := strconv.Atoi(req.PathParameter("id"))
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusBadRequest, "Bad Request")
		return
	}
	book.ID = int64(i)
	e = handler.store.update(book)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusInternalServerError, "Internal handler error")
		return
	}
	res.WriteHeaderAndEntity(http.StatusOK, book)
}

func (handler *bookHandler) delete(req *restful.Request, res *restful.Response) {
	var id int64
	i, e := strconv.Atoi(req.PathParameter("id"))
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusBadRequest, "Bad Request")
		return
	}
	id = int64(i)
	e = handler.store.delete(id)
	if e != nil {
		log.Println(e)
		res.WriteErrorString(http.StatusInternalServerError, "Internal handler error")
		return
	}

	res.WriteHeader(http.StatusOK)
}

func (handler *bookHandler) regsiter(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/books").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("").To(handler.all))
	ws.Route(ws.GET("/{id}").To(handler.get))
	ws.Route(ws.POST("").To(handler.create))
	ws.Route(ws.PUT("/{id}").To(handler.update))
	ws.Route(ws.DELETE("/{id}").To(handler.delete))

	container.Add(ws)
}
