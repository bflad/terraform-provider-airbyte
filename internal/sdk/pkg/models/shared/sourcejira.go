// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
	"time"
)

type Jira string

const (
	JiraJira Jira = "jira"
)

func (e Jira) ToPointer() *Jira {
	return &e
}

func (e *Jira) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "jira":
		*e = Jira(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Jira: %v", v)
	}
}

type SourceJira struct {
	// Jira API Token. See the <a href="https://docs.airbyte.com/integrations/sources/jira">docs</a> for more information on how to generate this key. API Token is used for Authorization to your account by BasicAuth.
	APIToken string `json:"api_token"`
	// The Domain for your Jira account, e.g. airbyteio.atlassian.net, airbyteio.jira.com, jira.your-domain.com
	Domain string `json:"domain"`
	// The user email for your Jira account which you used to generate the API token. This field is used for Authorization to your account by BasicAuth.
	Email string `json:"email"`
	// Allow the use of experimental streams which rely on undocumented Jira API endpoints. See https://docs.airbyte.com/integrations/sources/jira#experimental-tables for more info.
	EnableExperimentalStreams *bool `default:"false" json:"enable_experimental_streams"`
	// Expand the changelog when replicating issues.
	ExpandIssueChangelog *bool `default:"false" json:"expand_issue_changelog"`
	// List of Jira project keys to replicate data for, or leave it empty if you want to replicate data for all projects.
	Projects []string `json:"projects,omitempty"`
	// Render issue fields in HTML format in addition to Jira JSON-like format.
	RenderFields *bool `default:"false" json:"render_fields"`
	sourceType   Jira  `const:"jira" json:"sourceType"`
	// The date from which you want to replicate data from Jira, use the format YYYY-MM-DDT00:00:00Z. Note that this field only applies to certain streams, and only data generated on or after the start date will be replicated. Or leave it empty if you want to replicate all data. For more information, refer to the <a href="https://docs.airbyte.com/integrations/sources/jira/">documentation</a>.
	StartDate *time.Time `json:"start_date,omitempty"`
}

func (s SourceJira) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceJira) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceJira) GetAPIToken() string {
	if o == nil {
		return ""
	}
	return o.APIToken
}

func (o *SourceJira) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *SourceJira) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *SourceJira) GetEnableExperimentalStreams() *bool {
	if o == nil {
		return nil
	}
	return o.EnableExperimentalStreams
}

func (o *SourceJira) GetExpandIssueChangelog() *bool {
	if o == nil {
		return nil
	}
	return o.ExpandIssueChangelog
}

func (o *SourceJira) GetProjects() []string {
	if o == nil {
		return nil
	}
	return o.Projects
}

func (o *SourceJira) GetRenderFields() *bool {
	if o == nil {
		return nil
	}
	return o.RenderFields
}

func (o *SourceJira) GetSourceType() Jira {
	return JiraJira
}

func (o *SourceJira) GetStartDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartDate
}
