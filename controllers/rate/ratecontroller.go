package rate

import (
	"github.com/emicklei/go-restful"
	"fiscaluno-ws/database/filter"
	"fiscaluno-ws/models/rate/general"
	"fiscaluno-ws/models/rate/detailed"
	"log"
)

// Gets all ratings by user id
func RatedBy(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	var filterList = [] filter.Filter{*filter.New("RatedBy", id, "=")}
	ratings, err := getRatingByUser(filterList)
	if err != nil {
		response.WriteEntity(err)
	}else {
		response.WriteEntity(ratings)

	}
}

func GeneralRate(request *restful.Request, response *restful.Response) {
	rate := new(general.GeneralRate)
	err := request.ReadEntity(&rate)
	//TODO: need improvement on error checking
	if err == nil {
		err = newGeneralRate(*rate)
		if err != nil {
			response.WriteEntity(err)
		} else {
			response.WriteEntity(response.StatusCode())
		}
	}else {
		log.Panic(err)
	}
}

func NewDetailedRateForInstitution(request *restful.Request, response *restful.Response) {
	rate := new(specific.DetailedRate)
	err := request.ReadEntity(&rate)

	if err == nil {
		err = newDetailedRate(*rate)
		if err != nil {
			response.WriteEntity(err)
		} else {
			response.WriteEntity(response.StatusCode())
		}
	} else {
		log.Panic(err)
	}
}
