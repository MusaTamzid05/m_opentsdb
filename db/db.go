package db

import (
	"log"
	"time"

	"github.com/bluebreezecf/opentsdb-goclient/client"
	"github.com/bluebreezecf/opentsdb-goclient/config"
)

type OpenTSDB struct {
	client client.Client
}

func (o *OpenTSDB) init(address string) error {

	opentsdbCfg := config.OpenTSDBConfig{
		OpentsdbHost: address,
	}

	var err error

	o.client, err = client.NewClient(opentsdbCfg)

	if err != nil {
		return err
	}

	return nil
}

func (o *OpenTSDB) Insert(tags map[string]string, dataList map[string]interface{}) {

	datapoints := o.makeDataPointsFrom(tags, dataList)

	if len(datapoints) == 0 {
		log.Println("No data found!!")
		return
	}

	resp, err := o.client.Put(datapoints, "details")

	if err != nil {

		log.Println("Error inserting data!!!")
		log.Println(err)
		return
	}

	log.Println(resp.String())
}

func (*OpenTSDB) makeDataPointsFrom(tags map[string]string, dataList map[string]interface{}) []client.DataPoint {

	datapoints := make([]client.DataPoint, 0)

	for key, value := range dataList {

		dataPoint := client.DataPoint{
			Metric:    key,
			Timestamp: time.Now().Unix(),
			Value:     value,
		}

		dataPoint.Tags = tags
		datapoints = append(datapoints, dataPoint)
	}

	return datapoints

}

func NewOpenTSDB(host string) (*OpenTSDB, error) {

	opentsdb := OpenTSDB{}
	err := opentsdb.init(host)

	if err != nil {
		return nil, err
	}

	return &opentsdb, nil
}

func (o *OpenTSDB) Search(query Query) (map[string][]*client.DataPoint, error) {

	queryParam := client.QueryParam{
		Start: query.StartTime,
		End:   query.EndTime,
	}

	subqueries := make([]client.SubQuery, 0)

	subQuery := client.SubQuery{
		Aggregator: "sum",
		Metric:     query.MatrixName,
		Tags:       query.Tags,
	}

	subqueries = append(subqueries, subQuery)
	queryParam.Queries = subqueries

	queryRes, err := o.client.Query(queryParam)

	if err != nil {
		return nil, err
	}

	results := make(map[string][]*client.DataPoint)

	for _, res := range queryRes.QueryRespCnts {
		results[res.Metric] = append(results[res.Metric], res.GetDataPoints()...)

	}

	return results, nil

}
