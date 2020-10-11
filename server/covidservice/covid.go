package covidservice

import (
	"context"
	pb "covid_server/pb/proto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var resposneData []*pb.CovidCasesResponse
var url = "https://nepalcorona.info/api/v1/data/world"
var methodGet = "GET"

// Server initializes new covid server.
type Server struct{}

//NewCovidServiceServer creates new server.
func NewCovidServiceServer() *Server {
	return &Server{}
}

// GetCurrentCovidInfo get covid data based on the search srting.
func (css *Server) GetCurrentCovidInfo(ctx context.Context, req *pb.CovidCasesRequest) (*pb.CovidCasesResponse, error) {

	client := &http.Client{}
	request, err := http.NewRequest(methodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &resposneData); err != nil {
		return nil, fmt.Errorf("Cannot un marshal body:%v", err)
	}
	filteredCovidData := filterDataByCountry(resposneData, req.SearchString)

	return filteredCovidData, nil
}

// filterDataByCountry filter the covid data by search string
func filterDataByCountry(data []*pb.CovidCasesResponse, search string) *pb.CovidCasesResponse {
	for _, k := range resposneData {
		data2 := k

		if data2.Country == search {
			return data2
		}
	}

	return nil
}
