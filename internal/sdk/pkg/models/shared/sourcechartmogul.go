// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

// SourceChartmogulInterval - Some APIs such as <a href="https://dev.chartmogul.com/reference/endpoint-overview-metrics-api">Metrics</a> require intervals to cluster data.
type SourceChartmogulInterval string

const (
	SourceChartmogulIntervalDay     SourceChartmogulInterval = "day"
	SourceChartmogulIntervalWeek    SourceChartmogulInterval = "week"
	SourceChartmogulIntervalMonth   SourceChartmogulInterval = "month"
	SourceChartmogulIntervalQuarter SourceChartmogulInterval = "quarter"
)

func (e SourceChartmogulInterval) ToPointer() *SourceChartmogulInterval {
	return &e
}

func (e *SourceChartmogulInterval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "day":
		fallthrough
	case "week":
		fallthrough
	case "month":
		fallthrough
	case "quarter":
		*e = SourceChartmogulInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceChartmogulInterval: %v", v)
	}
}

type Chartmogul string

const (
	ChartmogulChartmogul Chartmogul = "chartmogul"
)

func (e Chartmogul) ToPointer() *Chartmogul {
	return &e
}

func (e *Chartmogul) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "chartmogul":
		*e = Chartmogul(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Chartmogul: %v", v)
	}
}

type SourceChartmogul struct {
	// Your Chartmogul API key. See <a href="https://help.chartmogul.com/hc/en-us/articles/4407796325906-Creating-and-Managing-API-keys#creating-an-api-key"> the docs </a> for info on how to obtain this.
	APIKey string `json:"api_key"`
	// Some APIs such as <a href="https://dev.chartmogul.com/reference/endpoint-overview-metrics-api">Metrics</a> require intervals to cluster data.
	Interval   *SourceChartmogulInterval `default:"month" json:"interval"`
	sourceType Chartmogul                `const:"chartmogul" json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. When feasible, any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}

func (s SourceChartmogul) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceChartmogul) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceChartmogul) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

func (o *SourceChartmogul) GetInterval() *SourceChartmogulInterval {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *SourceChartmogul) GetSourceType() Chartmogul {
	return ChartmogulChartmogul
}

func (o *SourceChartmogul) GetStartDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartDate
}
