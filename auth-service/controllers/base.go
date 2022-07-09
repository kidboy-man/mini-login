package controllers

import (
	"auth-service/datatransfers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/pagination"
)

type JSONResponse struct {
	Success     bool        `json:"success"`
	Status      int         `json:"status"` // http status code
	Data        interface{} `json:"data"`
	Error       error       `json:"error"`
	CurrentPage int         `json:"currentPage"`
	TotalPages  int         `json:"totalPages"`
	DataPerPage int         `json:"dataPerPage"`
	HasNextPage bool        `json:"hasNextPage"`
	HasPrevPage bool        `json:"hasPrevPage"`
}

func doReturnOK(response *JSONResponse, obj interface{}) {
	response.Success = true
	response.Status = http.StatusOK
	json, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	response.Data = string(json)
	return
}

func doReturnNotOK(response *JSONResponse, err error) {
	response.Success = false
	response.Status = http.StatusInternalServerError
	if v, ok := err.(*datatransfers.CustomError); ok {
		response.Error = v
		response.Status = v.Status
		return
	}

	response.Error = err
	return
}

func (j *JSONResponse) ReturnJSONResponse(obj interface{}, err error) (response *JSONResponse) {
	if err != nil {
		doReturnNotOK(j, err)
		return
	}

	doReturnOK(j, obj)
	return
}

func (j *JSONResponse) SetPagination(ctx *context.Context, totalData int64, limit, page int) {
	log.Println("totalData", totalData)
	log.Println("limit", limit)
	paginator := pagination.SetPaginator(ctx, limit, totalData)

	j.CurrentPage = page
	j.TotalPages = paginator.PageNums()
	log.Println("response", j.TotalPages)
	j.DataPerPage = limit
	j.HasNextPage = paginator.HasNext()
	j.HasPrevPage = paginator.HasPrev()
	log.Println("response", j)
	return
}
