package hydrantclient

import (
	"encoding/json"
	"time"
)

type DimensionType string

const (
	DimensionTypeString DimensionType = "string"
	DimensionTypeEpoch                = "epoch"
)

type DimensionId string

type Dimension struct {
	Id                      DimensionId   `json:"id"`
	Description             string        `json:"description"`
	Group                   string        `json:"group"`
	Type                    DimensionType `json":type`
	AggregationFunctionName *string       `json:"aggregationFunctionName,omitempty"`
	IndexValueHashes        bool          `json:"indexValueHashes"`
}

type Value struct {
	Dimension *Dimension
	Value     interface{}
}

type AggregationInstance struct {
	Dimension     *Dimension `json:"dimension"`
	CategoryValue string     `json:"categoryValue"`
	FunctionName  string     `json:"functionName"`
}

type EventTypeId string

type EventType struct {
	Id          EventTypeId `json:"id"`
	Dimensions  []Dimension `json:"dimensions"`
	Description string      `json:"description"`
}

type Schema struct {
	Id         string       `json:"id"`
	EventTypes []*EventType `json:"eventTypes"`
}

type GroupByMode string

const (
	GroupByModeRow   GroupByMode = "row"
	GroupByModeNest              = "nest"
	GroupByModePivot             = "pivot"
)

type GroupBy struct {
	Dimension string      `json:"dimension"`
	Mode      GroupByMode `json:"mode,omitempty"`
}

type FilterOperator string

type Filter struct {
	Dimension string         `json:"dimension"`
	Operator  FilterOperator `json:"operator"`
	Value     interface{}    `json:"value"`
}

type Timeframe string

const (
	TimeframeAllTime    Timeframe = "all_time"
	TimeframeLastHour             = "last_hour"
	TimeframeToday                = "today"
	TimeframeYesterday            = "yesterday"
	TimeframeThisWeek             = "this_week"
	TimeframeLastWeek             = "last_week"
	TimeframeLast2Weeks           = "last_2_weeks"
	TimeframeLast4Weeks           = "last_4_weeks"
	TimeframeThisMonth            = "this_month"
	TimeframeLastMonth            = "last_month"
	TimeframeThisYear             = "this_year"
	TimeframeLastYear             = "last_year"
)

// TODO: In current API, this is an int
type SortDirection string

const (
	SortDirectionAscending  = "asc"
	SortDirectionDescending = "desc"
)

type TimeInterval string

const (
	TimeInterval1y TimeInterval = "1y"
	TimeInterval1M              = "1M"
	TimeInterval1w              = "1w"
	TimeInterval1d              = "1d"
	TimeInterval1h              = "1h"
)

type TimeRange struct {
	Start *time.Time `json:"start,omitempty"`
	End   *time.Time `json:"end,omitempty"`
}

func (timeRange *TimeRange) MarshalJSON() ([]byte, error) {
	// TODO: API currently uses epoch, we should change that
	var epoch struct {
		Start *float64 `json:"start"`
		Stop  *float64 `json:"stop"`
	}
	if timeRange.Start != nil {
		start := timeRange.Start.Unix()
		epoch.Start = &start
	}
	if timeRange.Stop != nil {
		stop := timeRange.Stop.Unix()
		epoch.Stop = &stop
	}
	return json.Marshal(epoch)
}

type Query struct {
	Schema                  string         `json:"schema"`
	EventTypes              []EventTypeId  `json:"eventTypes,omitempty"`
	Filters                 []*Filter      `json:"filters,omitempty"`
	GroupByItems            []*GroupBy     `json:"groupByItems,omitempty"`
	Timeframe               *Timeframe     `json:"timeframe,omitempty"`
	TimeRange               *TimeRange     `json:"timeRange,omitempty"`
	TimeInterval            *TimeInterval  `json:"timeInterval,omitempty"`
	SortDimension           *DimensionId   `json:"sortDimension,omitempty"`
	SortDirection           *SortDirection `json:"sortDirection,omitempty"`
	Limit                   *int64         `json:"limit,omitempty"`
	TimeZone                *string        `json:"timeZone,omitempty"`
	AggregationFunctionName *string        `json:"aggregationFunctionName,omitempty"`
}

type ResultRow struct {
	Values map[DimensionId]interface{} `json:"values"`
}

type ResultSet struct {
	Rows []ResultRow `json:"rows"`
}
