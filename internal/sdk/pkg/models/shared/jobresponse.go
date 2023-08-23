// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JobResponse - Provides details of a single job.
type JobResponse struct {
	BytesSynced  *int64 `json:"bytesSynced,omitempty"`
	ConnectionID string `json:"connectionId"`
	// Duration of a sync in ISO_8601 format
	Duration *string `json:"duration,omitempty"`
	JobID    int64   `json:"jobId"`
	// Enum that describes the different types of jobs that the platform runs.
	JobType       JobTypeEnum   `json:"jobType"`
	LastUpdatedAt *string       `json:"lastUpdatedAt,omitempty"`
	RowsSynced    *int64        `json:"rowsSynced,omitempty"`
	StartTime     string        `json:"startTime"`
	Status        JobStatusEnum `json:"status"`
}
