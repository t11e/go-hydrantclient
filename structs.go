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
	ID                      DimensionId   `json:"id"`
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
	ID          EventTypeId `json:"id"`
	Dimensions  []Dimension `json:"dimensions"`
	Description string      `json:"description"`
}

type Schema struct {
	ID         DimensionId  `json:"id"`
	EventTypes []*EventType `json:"eventTypes"`
}

type Grouping struct {
	// ID is the ID of the dimension.
	ID DimensionId `json:"id"`
}

// TableGroupingMode defines the mode of tabular grouping.
type TableGroupingMode string

const (
	TableGroupingModeRow   TableGroupingMode = "row"
	TableGroupingModeNest                    = "nest"
	TableGroupingModePivot                   = "pivot"
)

// TableGrouping specifies how to group a table.
type TableGrouping struct {
	// ID is the ID of the dimension.
	ID DimensionId `json:"id"`

	// Mode is the mode.
	Mode TableGroupingMode `json:"mode,omitempty"`
}

type FilterOperator string

const (
	FilterEq          FilterOperator = "eq"
	FilterNotEq                      = "not_eq"
	FilterIn                         = "in"
	FilterNotIn                      = "not_in"
	FilterIsNull                     = "is_null"
	FilterIsNotNull                  = "is_not_null"
	FilterContainsAny                = "contains_any"
	FilterContainsAll                = "contains_all"
)

type Filter struct {
	ID       DimensionId    `json:"id"`
	Operator FilterOperator `json:"operator"`
	Value    interface{}    `json:"value"`
}

// Timeframe is a named timeframe.
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

type SortDirection string

const (
	SortDirectionAscending  = "asc"
	SortDirectionDescending = "desc"
)

// TimeInterval is the name of a time interval.
type TimeInterval string

const (
	TimeInterval1y TimeInterval = "1y"
	TimeInterval1M              = "1M"
	TimeInterval1w              = "1w"
	TimeInterval1d              = "1d"
	TimeInterval1h              = "1h"
)

// TimeRange defines a start/end time range.
type TimeRange struct {
	Start *time.Time `json:"start,omitempty"`
	End   *time.Time `json:"end,omitempty"`
}

func (timeRange *TimeRange) MarshalJSON() ([]byte, error) {
	// TODO: API currently uses epoch, we should change that
	var epoch struct {
		Start *float64 `json:"start"`
		End   *float64 `json:"end"`
	}
	if timeRange.Start != nil {
		start := float64(timeRange.Start.Unix())
		epoch.Start = &start
	}
	if timeRange.End != nil {
		stop := float64(timeRange.End.Unix())
		epoch.End = &stop
	}
	return json.Marshal(epoch)
}

// Query is an aggregation query.
type Query struct {
	Schema                  string        `json:"schema"`
	EventTypes              []EventTypeId `json:"eventTypes,omitempty"`
	Filters                 []*Filter     `json:"filters,omitempty"`
	Timeframe               *Timeframe    `json:"timeframe,omitempty"`
	TimeRange               *TimeRange    `json:"timeRange,omitempty"`
	TimeInterval            *TimeInterval `json:"timeInterval,omitempty"`
	Limit                   *int64        `json:"limit,omitempty"`
	TimeZone                *string       `json:"timeZone,omitempty"`
	AggregationFunctionName *string       `json:"aggregationFunctionName,omitempty"`
	Groupings               []*Grouping   `json:"grouping,omitempty"`
}

// ResultBucket is a bucket in a result set.
type ResultBucket struct {
	// Key is the grouping value that this bucket represents.
	Key interface{} `json:"key"`

	// Count is the number of matches.
	Count int64 `json:"count"`

	// Buckets contain nested aggregation results.
	Buckets []*ResultBucket `json:"buckets"`
}

// ResultSet is the results of a query.
type ResultSet struct {
	Buckets []ResultBucket `json:"buckets"`
}

// TableQuery is a tabular query that supports pivoting and nesting.
type TableQuery struct {
	Query
	TableGroupings []*TableGrouping `json:"tableGroupings,omitempty"`
	SortDimension  *DimensionId     `json:"sortDimension,omitempty"`
	SortDirection  *SortDirection   `json:"sortDirection,omitempty"`
}

// TableRow is a table row. Its value can be primitives or rows.
type TableRow struct {
	Values map[DimensionId]interface{} `json:"values"`
}

// TableResultSet is a tabular result set.
type TableResultSet struct {
	Rows []TableRow `json:"rows"`
}
