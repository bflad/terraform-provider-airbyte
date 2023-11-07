// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type Pypi string

const (
	PypiPypi Pypi = "pypi"
)

func (e Pypi) ToPointer() *Pypi {
	return &e
}

func (e *Pypi) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pypi":
		*e = Pypi(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pypi: %v", v)
	}
}

type SourcePypi struct {
	// Name of the project/package. Can only be in lowercase with hyphen. This is the name used using pip command for installing the package.
	ProjectName string `json:"project_name"`
	sourceType  Pypi   `const:"pypi" json:"sourceType"`
	// Version of the project/package.  Use it to find a particular release instead of all releases.
	Version *string `json:"version,omitempty"`
}

func (s SourcePypi) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourcePypi) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourcePypi) GetProjectName() string {
	if o == nil {
		return ""
	}
	return o.ProjectName
}

func (o *SourcePypi) GetSourceType() Pypi {
	return PypiPypi
}

func (o *SourcePypi) GetVersion() *string {
	if o == nil {
		return nil
	}
	return o.Version
}
