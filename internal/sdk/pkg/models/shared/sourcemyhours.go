// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type MyHours string

const (
	MyHoursMyHours MyHours = "my-hours"
)

func (e MyHours) ToPointer() *MyHours {
	return &e
}

func (e *MyHours) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "my-hours":
		*e = MyHours(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MyHours: %v", v)
	}
}

type SourceMyHours struct {
	// Your My Hours username
	Email string `json:"email"`
	// Pagination size used for retrieving logs in days
	LogsBatchSize *int64 `default:"30" json:"logs_batch_size"`
	// The password associated to the username
	Password   string  `json:"password"`
	sourceType MyHours `const:"my-hours" json:"sourceType"`
	// Start date for collecting time logs
	StartDate string `json:"start_date"`
}

func (s SourceMyHours) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMyHours) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMyHours) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceMyHours) GetLogsBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.LogsBatchSize
}

func (o *SourceMyHours) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SourceMyHours) GetSourceType() MyHours {
	return MyHoursMyHours
}

func (o *SourceMyHours) GetStartDate() string {
	if o == nil {
		return ""
	}
	return o.StartDate
}
