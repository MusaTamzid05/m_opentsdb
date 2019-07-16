package db

import (
	"strconv"
)

type Query struct {
	StartTime  int64
	EndTime    int64
	Tags       map[string]string
	MatrixName string
}

func (q *Query) String() string {

	startTime := strconv.Itoa(int(q.StartTime))
	endTime := strconv.Itoa(int(q.EndTime))

	str := "\n"
	str += "Start Time : " + startTime + "\n"
	str += "End Time : " + endTime + "\n"
	str += "Tags =>  "

	for i, tag := range q.Tags {
		str += string(i) + ":" + tag + ","
	}

	str += "\n"

	str += "Matrix : " + q.MatrixName + "\n"

	return str
}

func MakeQuery(startTime, endTime int64, matrixName string, tags map[string]string) Query {
	return Query{StartTime: startTime, EndTime: endTime, MatrixName: matrixName, Tags: tags}

}
