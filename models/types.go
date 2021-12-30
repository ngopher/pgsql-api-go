package models

import (
	"encoding/json"
	"time"
)

type PaginationRequest struct {
	Limit, Offset int
}

type DBResponse struct {
	Query      string    `json:"query"`
	Calls      int       `json:"calls"`
	TotalTime  time.Time `json:"total_time"`
	MinTime    time.Time `json:"min_time"`
	MaxTime    time.Time `json:"max_time"`
	MeanTime   time.Time `json:"mean_time"`
	StddevTime time.Time `json:"stddev_time"`
	Rows       int       `json:"rows"`
}

func (D *DBResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, D)
}

func (D *DBResponse) MarshalBinary() (data []byte, err error) {
	return json.Marshal(data)
}

//APIRequest ...
type APIRequest struct {
	Filter     string
	Pagination *PaginationRequest
}
