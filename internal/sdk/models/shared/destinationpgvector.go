// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type Pgvector string

const (
	PgvectorPgvector Pgvector = "pgvector"
)

func (e Pgvector) ToPointer() *Pgvector {
	return &e
}
func (e *Pgvector) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pgvector":
		*e = Pgvector(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Pgvector: %v", v)
	}
}

type DestinationPgvectorSchemasEmbeddingEmbedding5Mode string

const (
	DestinationPgvectorSchemasEmbeddingEmbedding5ModeOpenaiCompatible DestinationPgvectorSchemasEmbeddingEmbedding5Mode = "openai_compatible"
)

func (e DestinationPgvectorSchemasEmbeddingEmbedding5Mode) ToPointer() *DestinationPgvectorSchemasEmbeddingEmbedding5Mode {
	return &e
}
func (e *DestinationPgvectorSchemasEmbeddingEmbedding5Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai_compatible":
		*e = DestinationPgvectorSchemasEmbeddingEmbedding5Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasEmbeddingEmbedding5Mode: %v", v)
	}
}

// DestinationPgvectorOpenAICompatible - Use a service that's compatible with the OpenAI API to embed text.
type DestinationPgvectorOpenAICompatible struct {
	APIKey *string `default:"" json:"api_key"`
	// The base URL for your OpenAI-compatible service
	BaseURL string `json:"base_url"`
	// The number of dimensions the embedding model is generating
	Dimensions int64                                              `json:"dimensions"`
	mode       *DestinationPgvectorSchemasEmbeddingEmbedding5Mode `const:"openai_compatible" json:"mode"`
	// The name of the model to use for embedding
	ModelName *string `default:"text-embedding-ada-002" json:"model_name"`
}

func (d DestinationPgvectorOpenAICompatible) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorOpenAICompatible) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorOpenAICompatible) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *DestinationPgvectorOpenAICompatible) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *DestinationPgvectorOpenAICompatible) GetDimensions() int64 {
	if o == nil {
		return 0
	}
	return o.Dimensions
}

func (o *DestinationPgvectorOpenAICompatible) GetMode() *DestinationPgvectorSchemasEmbeddingEmbedding5Mode {
	return DestinationPgvectorSchemasEmbeddingEmbedding5ModeOpenaiCompatible.ToPointer()
}

func (o *DestinationPgvectorOpenAICompatible) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

type DestinationPgvectorSchemasEmbeddingEmbeddingMode string

const (
	DestinationPgvectorSchemasEmbeddingEmbeddingModeAzureOpenai DestinationPgvectorSchemasEmbeddingEmbeddingMode = "azure_openai"
)

func (e DestinationPgvectorSchemasEmbeddingEmbeddingMode) ToPointer() *DestinationPgvectorSchemasEmbeddingEmbeddingMode {
	return &e
}
func (e *DestinationPgvectorSchemasEmbeddingEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_openai":
		*e = DestinationPgvectorSchemasEmbeddingEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasEmbeddingEmbeddingMode: %v", v)
	}
}

// DestinationPgvectorAzureOpenAI - Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPgvectorAzureOpenAI struct {
	// The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	APIBase string `json:"api_base"`
	// The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	Deployment string                                            `json:"deployment"`
	mode       *DestinationPgvectorSchemasEmbeddingEmbeddingMode `const:"azure_openai" json:"mode"`
	// The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	OpenaiKey string `json:"openai_key"`
}

func (d DestinationPgvectorAzureOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorAzureOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorAzureOpenAI) GetAPIBase() string {
	if o == nil {
		return ""
	}
	return o.APIBase
}

func (o *DestinationPgvectorAzureOpenAI) GetDeployment() string {
	if o == nil {
		return ""
	}
	return o.Deployment
}

func (o *DestinationPgvectorAzureOpenAI) GetMode() *DestinationPgvectorSchemasEmbeddingEmbeddingMode {
	return DestinationPgvectorSchemasEmbeddingEmbeddingModeAzureOpenai.ToPointer()
}

func (o *DestinationPgvectorAzureOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationPgvectorSchemasEmbeddingMode string

const (
	DestinationPgvectorSchemasEmbeddingModeFake DestinationPgvectorSchemasEmbeddingMode = "fake"
)

func (e DestinationPgvectorSchemasEmbeddingMode) ToPointer() *DestinationPgvectorSchemasEmbeddingMode {
	return &e
}
func (e *DestinationPgvectorSchemasEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationPgvectorSchemasEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasEmbeddingMode: %v", v)
	}
}

// DestinationPgvectorFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationPgvectorFake struct {
	mode *DestinationPgvectorSchemasEmbeddingMode `const:"fake" json:"mode"`
}

func (d DestinationPgvectorFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorFake) GetMode() *DestinationPgvectorSchemasEmbeddingMode {
	return DestinationPgvectorSchemasEmbeddingModeFake.ToPointer()
}

type DestinationPgvectorSchemasMode string

const (
	DestinationPgvectorSchemasModeCohere DestinationPgvectorSchemasMode = "cohere"
)

func (e DestinationPgvectorSchemasMode) ToPointer() *DestinationPgvectorSchemasMode {
	return &e
}
func (e *DestinationPgvectorSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationPgvectorSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasMode: %v", v)
	}
}

// DestinationPgvectorCohere - Use the Cohere API to embed text.
type DestinationPgvectorCohere struct {
	CohereKey string                          `json:"cohere_key"`
	mode      *DestinationPgvectorSchemasMode `const:"cohere" json:"mode"`
}

func (d DestinationPgvectorCohere) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorCohere) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorCohere) GetCohereKey() string {
	if o == nil {
		return ""
	}
	return o.CohereKey
}

func (o *DestinationPgvectorCohere) GetMode() *DestinationPgvectorSchemasMode {
	return DestinationPgvectorSchemasModeCohere.ToPointer()
}

type DestinationPgvectorMode string

const (
	DestinationPgvectorModeOpenai DestinationPgvectorMode = "openai"
)

func (e DestinationPgvectorMode) ToPointer() *DestinationPgvectorMode {
	return &e
}
func (e *DestinationPgvectorMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationPgvectorMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorMode: %v", v)
	}
}

// DestinationPgvectorOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPgvectorOpenAI struct {
	mode      *DestinationPgvectorMode `const:"openai" json:"mode"`
	OpenaiKey string                   `json:"openai_key"`
}

func (d DestinationPgvectorOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorOpenAI) GetMode() *DestinationPgvectorMode {
	return DestinationPgvectorModeOpenai.ToPointer()
}

func (o *DestinationPgvectorOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationPgvectorEmbeddingType string

const (
	DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAI           DestinationPgvectorEmbeddingType = "destination-pgvector_OpenAI"
	DestinationPgvectorEmbeddingTypeDestinationPgvectorCohere           DestinationPgvectorEmbeddingType = "destination-pgvector_Cohere"
	DestinationPgvectorEmbeddingTypeDestinationPgvectorFake             DestinationPgvectorEmbeddingType = "destination-pgvector_Fake"
	DestinationPgvectorEmbeddingTypeDestinationPgvectorAzureOpenAI      DestinationPgvectorEmbeddingType = "destination-pgvector_Azure OpenAI"
	DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAICompatible DestinationPgvectorEmbeddingType = "destination-pgvector_OpenAI-compatible"
)

// DestinationPgvectorEmbedding - Embedding configuration
type DestinationPgvectorEmbedding struct {
	DestinationPgvectorOpenAI           *DestinationPgvectorOpenAI
	DestinationPgvectorCohere           *DestinationPgvectorCohere
	DestinationPgvectorFake             *DestinationPgvectorFake
	DestinationPgvectorAzureOpenAI      *DestinationPgvectorAzureOpenAI
	DestinationPgvectorOpenAICompatible *DestinationPgvectorOpenAICompatible

	Type DestinationPgvectorEmbeddingType
}

func CreateDestinationPgvectorEmbeddingDestinationPgvectorOpenAI(destinationPgvectorOpenAI DestinationPgvectorOpenAI) DestinationPgvectorEmbedding {
	typ := DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAI

	return DestinationPgvectorEmbedding{
		DestinationPgvectorOpenAI: &destinationPgvectorOpenAI,
		Type:                      typ,
	}
}

func CreateDestinationPgvectorEmbeddingDestinationPgvectorCohere(destinationPgvectorCohere DestinationPgvectorCohere) DestinationPgvectorEmbedding {
	typ := DestinationPgvectorEmbeddingTypeDestinationPgvectorCohere

	return DestinationPgvectorEmbedding{
		DestinationPgvectorCohere: &destinationPgvectorCohere,
		Type:                      typ,
	}
}

func CreateDestinationPgvectorEmbeddingDestinationPgvectorFake(destinationPgvectorFake DestinationPgvectorFake) DestinationPgvectorEmbedding {
	typ := DestinationPgvectorEmbeddingTypeDestinationPgvectorFake

	return DestinationPgvectorEmbedding{
		DestinationPgvectorFake: &destinationPgvectorFake,
		Type:                    typ,
	}
}

func CreateDestinationPgvectorEmbeddingDestinationPgvectorAzureOpenAI(destinationPgvectorAzureOpenAI DestinationPgvectorAzureOpenAI) DestinationPgvectorEmbedding {
	typ := DestinationPgvectorEmbeddingTypeDestinationPgvectorAzureOpenAI

	return DestinationPgvectorEmbedding{
		DestinationPgvectorAzureOpenAI: &destinationPgvectorAzureOpenAI,
		Type:                           typ,
	}
}

func CreateDestinationPgvectorEmbeddingDestinationPgvectorOpenAICompatible(destinationPgvectorOpenAICompatible DestinationPgvectorOpenAICompatible) DestinationPgvectorEmbedding {
	typ := DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAICompatible

	return DestinationPgvectorEmbedding{
		DestinationPgvectorOpenAICompatible: &destinationPgvectorOpenAICompatible,
		Type:                                typ,
	}
}

func (u *DestinationPgvectorEmbedding) UnmarshalJSON(data []byte) error {

	var destinationPgvectorFake DestinationPgvectorFake = DestinationPgvectorFake{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorFake, "", true, true); err == nil {
		u.DestinationPgvectorFake = &destinationPgvectorFake
		u.Type = DestinationPgvectorEmbeddingTypeDestinationPgvectorFake
		return nil
	}

	var destinationPgvectorOpenAI DestinationPgvectorOpenAI = DestinationPgvectorOpenAI{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorOpenAI, "", true, true); err == nil {
		u.DestinationPgvectorOpenAI = &destinationPgvectorOpenAI
		u.Type = DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAI
		return nil
	}

	var destinationPgvectorCohere DestinationPgvectorCohere = DestinationPgvectorCohere{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorCohere, "", true, true); err == nil {
		u.DestinationPgvectorCohere = &destinationPgvectorCohere
		u.Type = DestinationPgvectorEmbeddingTypeDestinationPgvectorCohere
		return nil
	}

	var destinationPgvectorAzureOpenAI DestinationPgvectorAzureOpenAI = DestinationPgvectorAzureOpenAI{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorAzureOpenAI, "", true, true); err == nil {
		u.DestinationPgvectorAzureOpenAI = &destinationPgvectorAzureOpenAI
		u.Type = DestinationPgvectorEmbeddingTypeDestinationPgvectorAzureOpenAI
		return nil
	}

	var destinationPgvectorOpenAICompatible DestinationPgvectorOpenAICompatible = DestinationPgvectorOpenAICompatible{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorOpenAICompatible, "", true, true); err == nil {
		u.DestinationPgvectorOpenAICompatible = &destinationPgvectorOpenAICompatible
		u.Type = DestinationPgvectorEmbeddingTypeDestinationPgvectorOpenAICompatible
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationPgvectorEmbedding", string(data))
}

func (u DestinationPgvectorEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationPgvectorOpenAI != nil {
		return utils.MarshalJSON(u.DestinationPgvectorOpenAI, "", true)
	}

	if u.DestinationPgvectorCohere != nil {
		return utils.MarshalJSON(u.DestinationPgvectorCohere, "", true)
	}

	if u.DestinationPgvectorFake != nil {
		return utils.MarshalJSON(u.DestinationPgvectorFake, "", true)
	}

	if u.DestinationPgvectorAzureOpenAI != nil {
		return utils.MarshalJSON(u.DestinationPgvectorAzureOpenAI, "", true)
	}

	if u.DestinationPgvectorOpenAICompatible != nil {
		return utils.MarshalJSON(u.DestinationPgvectorOpenAICompatible, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationPgvectorEmbedding: all fields are null")
}

type DestinationPgvectorCredentials struct {
	// Enter the password you want to use to access the database
	Password string `json:"password"`
}

func (o *DestinationPgvectorCredentials) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

// DestinationPgvectorPostgresConnection - Postgres can be used to store vector data and retrieve embeddings.
type DestinationPgvectorPostgresConnection struct {
	Credentials DestinationPgvectorCredentials `json:"credentials"`
	// Enter the name of the database that you want to sync data into
	Database string `json:"database"`
	// Enter the name of the default schema
	DefaultSchema string `json:"default_schema"`
	// Enter the account name you want to use to access the database.
	Host string `json:"host"`
	// Enter the port you want to use to access the database
	Port int64 `json:"port"`
	// Enter the name of the user you want to use to access the database
	Username string `json:"username"`
}

func (o *DestinationPgvectorPostgresConnection) GetCredentials() DestinationPgvectorCredentials {
	if o == nil {
		return DestinationPgvectorCredentials{}
	}
	return o.Credentials
}

func (o *DestinationPgvectorPostgresConnection) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationPgvectorPostgresConnection) GetDefaultSchema() string {
	if o == nil {
		return ""
	}
	return o.DefaultSchema
}

func (o *DestinationPgvectorPostgresConnection) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationPgvectorPostgresConnection) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *DestinationPgvectorPostgresConnection) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type DestinationPgvectorFieldNameMappingConfigModel struct {
	// The field name in the source
	FromField string `json:"from_field"`
	// The field name to use in the destination
	ToField string `json:"to_field"`
}

func (o *DestinationPgvectorFieldNameMappingConfigModel) GetFromField() string {
	if o == nil {
		return ""
	}
	return o.FromField
}

func (o *DestinationPgvectorFieldNameMappingConfigModel) GetToField() string {
	if o == nil {
		return ""
	}
	return o.ToField
}

// DestinationPgvectorLanguage - Split code in suitable places based on the programming language
type DestinationPgvectorLanguage string

const (
	DestinationPgvectorLanguageCpp      DestinationPgvectorLanguage = "cpp"
	DestinationPgvectorLanguageGo       DestinationPgvectorLanguage = "go"
	DestinationPgvectorLanguageJava     DestinationPgvectorLanguage = "java"
	DestinationPgvectorLanguageJs       DestinationPgvectorLanguage = "js"
	DestinationPgvectorLanguagePhp      DestinationPgvectorLanguage = "php"
	DestinationPgvectorLanguageProto    DestinationPgvectorLanguage = "proto"
	DestinationPgvectorLanguagePython   DestinationPgvectorLanguage = "python"
	DestinationPgvectorLanguageRst      DestinationPgvectorLanguage = "rst"
	DestinationPgvectorLanguageRuby     DestinationPgvectorLanguage = "ruby"
	DestinationPgvectorLanguageRust     DestinationPgvectorLanguage = "rust"
	DestinationPgvectorLanguageScala    DestinationPgvectorLanguage = "scala"
	DestinationPgvectorLanguageSwift    DestinationPgvectorLanguage = "swift"
	DestinationPgvectorLanguageMarkdown DestinationPgvectorLanguage = "markdown"
	DestinationPgvectorLanguageLatex    DestinationPgvectorLanguage = "latex"
	DestinationPgvectorLanguageHTML     DestinationPgvectorLanguage = "html"
	DestinationPgvectorLanguageSol      DestinationPgvectorLanguage = "sol"
)

func (e DestinationPgvectorLanguage) ToPointer() *DestinationPgvectorLanguage {
	return &e
}
func (e *DestinationPgvectorLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cpp":
		fallthrough
	case "go":
		fallthrough
	case "java":
		fallthrough
	case "js":
		fallthrough
	case "php":
		fallthrough
	case "proto":
		fallthrough
	case "python":
		fallthrough
	case "rst":
		fallthrough
	case "ruby":
		fallthrough
	case "rust":
		fallthrough
	case "scala":
		fallthrough
	case "swift":
		fallthrough
	case "markdown":
		fallthrough
	case "latex":
		fallthrough
	case "html":
		fallthrough
	case "sol":
		*e = DestinationPgvectorLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorLanguage: %v", v)
	}
}

type DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode string

const (
	DestinationPgvectorSchemasProcessingTextSplitterTextSplitterModeCode DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode = "code"
)

func (e DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode) ToPointer() *DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode {
	return &e
}
func (e *DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code":
		*e = DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode: %v", v)
	}
}

// DestinationPgvectorByProgrammingLanguage - Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.
type DestinationPgvectorByProgrammingLanguage struct {
	// Split code in suitable places based on the programming language
	Language DestinationPgvectorLanguage                                       `json:"language"`
	mode     *DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode `const:"code" json:"mode"`
}

func (d DestinationPgvectorByProgrammingLanguage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorByProgrammingLanguage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorByProgrammingLanguage) GetLanguage() DestinationPgvectorLanguage {
	if o == nil {
		return DestinationPgvectorLanguage("")
	}
	return o.Language
}

func (o *DestinationPgvectorByProgrammingLanguage) GetMode() *DestinationPgvectorSchemasProcessingTextSplitterTextSplitterMode {
	return DestinationPgvectorSchemasProcessingTextSplitterTextSplitterModeCode.ToPointer()
}

type DestinationPgvectorSchemasProcessingTextSplitterMode string

const (
	DestinationPgvectorSchemasProcessingTextSplitterModeMarkdown DestinationPgvectorSchemasProcessingTextSplitterMode = "markdown"
)

func (e DestinationPgvectorSchemasProcessingTextSplitterMode) ToPointer() *DestinationPgvectorSchemasProcessingTextSplitterMode {
	return &e
}
func (e *DestinationPgvectorSchemasProcessingTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "markdown":
		*e = DestinationPgvectorSchemasProcessingTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasProcessingTextSplitterMode: %v", v)
	}
}

// DestinationPgvectorByMarkdownHeader - Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.
type DestinationPgvectorByMarkdownHeader struct {
	mode *DestinationPgvectorSchemasProcessingTextSplitterMode `const:"markdown" json:"mode"`
	// Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points
	SplitLevel *int64 `default:"1" json:"split_level"`
}

func (d DestinationPgvectorByMarkdownHeader) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorByMarkdownHeader) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorByMarkdownHeader) GetMode() *DestinationPgvectorSchemasProcessingTextSplitterMode {
	return DestinationPgvectorSchemasProcessingTextSplitterModeMarkdown.ToPointer()
}

func (o *DestinationPgvectorByMarkdownHeader) GetSplitLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.SplitLevel
}

type DestinationPgvectorSchemasProcessingMode string

const (
	DestinationPgvectorSchemasProcessingModeSeparator DestinationPgvectorSchemasProcessingMode = "separator"
)

func (e DestinationPgvectorSchemasProcessingMode) ToPointer() *DestinationPgvectorSchemasProcessingMode {
	return &e
}
func (e *DestinationPgvectorSchemasProcessingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "separator":
		*e = DestinationPgvectorSchemasProcessingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPgvectorSchemasProcessingMode: %v", v)
	}
}

// DestinationPgvectorBySeparator - Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.
type DestinationPgvectorBySeparator struct {
	// Whether to keep the separator in the resulting chunks
	KeepSeparator *bool                                     `default:"false" json:"keep_separator"`
	mode          *DestinationPgvectorSchemasProcessingMode `const:"separator" json:"mode"`
	// List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".
	Separators []string `json:"separators,omitempty"`
}

func (d DestinationPgvectorBySeparator) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorBySeparator) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorBySeparator) GetKeepSeparator() *bool {
	if o == nil {
		return nil
	}
	return o.KeepSeparator
}

func (o *DestinationPgvectorBySeparator) GetMode() *DestinationPgvectorSchemasProcessingMode {
	return DestinationPgvectorSchemasProcessingModeSeparator.ToPointer()
}

func (o *DestinationPgvectorBySeparator) GetSeparators() []string {
	if o == nil {
		return nil
	}
	return o.Separators
}

type DestinationPgvectorTextSplitterType string

const (
	DestinationPgvectorTextSplitterTypeDestinationPgvectorBySeparator           DestinationPgvectorTextSplitterType = "destination-pgvector_By Separator"
	DestinationPgvectorTextSplitterTypeDestinationPgvectorByMarkdownHeader      DestinationPgvectorTextSplitterType = "destination-pgvector_By Markdown header"
	DestinationPgvectorTextSplitterTypeDestinationPgvectorByProgrammingLanguage DestinationPgvectorTextSplitterType = "destination-pgvector_By Programming Language"
)

// DestinationPgvectorTextSplitter - Split text fields into chunks based on the specified method.
type DestinationPgvectorTextSplitter struct {
	DestinationPgvectorBySeparator           *DestinationPgvectorBySeparator
	DestinationPgvectorByMarkdownHeader      *DestinationPgvectorByMarkdownHeader
	DestinationPgvectorByProgrammingLanguage *DestinationPgvectorByProgrammingLanguage

	Type DestinationPgvectorTextSplitterType
}

func CreateDestinationPgvectorTextSplitterDestinationPgvectorBySeparator(destinationPgvectorBySeparator DestinationPgvectorBySeparator) DestinationPgvectorTextSplitter {
	typ := DestinationPgvectorTextSplitterTypeDestinationPgvectorBySeparator

	return DestinationPgvectorTextSplitter{
		DestinationPgvectorBySeparator: &destinationPgvectorBySeparator,
		Type:                           typ,
	}
}

func CreateDestinationPgvectorTextSplitterDestinationPgvectorByMarkdownHeader(destinationPgvectorByMarkdownHeader DestinationPgvectorByMarkdownHeader) DestinationPgvectorTextSplitter {
	typ := DestinationPgvectorTextSplitterTypeDestinationPgvectorByMarkdownHeader

	return DestinationPgvectorTextSplitter{
		DestinationPgvectorByMarkdownHeader: &destinationPgvectorByMarkdownHeader,
		Type:                                typ,
	}
}

func CreateDestinationPgvectorTextSplitterDestinationPgvectorByProgrammingLanguage(destinationPgvectorByProgrammingLanguage DestinationPgvectorByProgrammingLanguage) DestinationPgvectorTextSplitter {
	typ := DestinationPgvectorTextSplitterTypeDestinationPgvectorByProgrammingLanguage

	return DestinationPgvectorTextSplitter{
		DestinationPgvectorByProgrammingLanguage: &destinationPgvectorByProgrammingLanguage,
		Type:                                     typ,
	}
}

func (u *DestinationPgvectorTextSplitter) UnmarshalJSON(data []byte) error {

	var destinationPgvectorByMarkdownHeader DestinationPgvectorByMarkdownHeader = DestinationPgvectorByMarkdownHeader{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorByMarkdownHeader, "", true, true); err == nil {
		u.DestinationPgvectorByMarkdownHeader = &destinationPgvectorByMarkdownHeader
		u.Type = DestinationPgvectorTextSplitterTypeDestinationPgvectorByMarkdownHeader
		return nil
	}

	var destinationPgvectorByProgrammingLanguage DestinationPgvectorByProgrammingLanguage = DestinationPgvectorByProgrammingLanguage{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorByProgrammingLanguage, "", true, true); err == nil {
		u.DestinationPgvectorByProgrammingLanguage = &destinationPgvectorByProgrammingLanguage
		u.Type = DestinationPgvectorTextSplitterTypeDestinationPgvectorByProgrammingLanguage
		return nil
	}

	var destinationPgvectorBySeparator DestinationPgvectorBySeparator = DestinationPgvectorBySeparator{}
	if err := utils.UnmarshalJSON(data, &destinationPgvectorBySeparator, "", true, true); err == nil {
		u.DestinationPgvectorBySeparator = &destinationPgvectorBySeparator
		u.Type = DestinationPgvectorTextSplitterTypeDestinationPgvectorBySeparator
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for DestinationPgvectorTextSplitter", string(data))
}

func (u DestinationPgvectorTextSplitter) MarshalJSON() ([]byte, error) {
	if u.DestinationPgvectorBySeparator != nil {
		return utils.MarshalJSON(u.DestinationPgvectorBySeparator, "", true)
	}

	if u.DestinationPgvectorByMarkdownHeader != nil {
		return utils.MarshalJSON(u.DestinationPgvectorByMarkdownHeader, "", true)
	}

	if u.DestinationPgvectorByProgrammingLanguage != nil {
		return utils.MarshalJSON(u.DestinationPgvectorByProgrammingLanguage, "", true)
	}

	return nil, errors.New("could not marshal union type DestinationPgvectorTextSplitter: all fields are null")
}

type DestinationPgvectorProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.
	FieldNameMappings []DestinationPgvectorFieldNameMappingConfigModel `json:"field_name_mappings,omitempty"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// Split text fields into chunks based on the specified method.
	TextSplitter *DestinationPgvectorTextSplitter `json:"text_splitter,omitempty"`
}

func (d DestinationPgvectorProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvectorProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvectorProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationPgvectorProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationPgvectorProcessingConfigModel) GetFieldNameMappings() []DestinationPgvectorFieldNameMappingConfigModel {
	if o == nil {
		return nil
	}
	return o.FieldNameMappings
}

func (o *DestinationPgvectorProcessingConfigModel) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationPgvectorProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *DestinationPgvectorProcessingConfigModel) GetTextSplitter() *DestinationPgvectorTextSplitter {
	if o == nil {
		return nil
	}
	return o.TextSplitter
}

// DestinationPgvector - The configuration model for the Vector DB based destinations. This model is used to generate the UI for the destination configuration,
// as well as to provide type safety for the configuration passed to the destination.
//
// The configuration model is composed of four parts:
// * Processing configuration
// * Embedding configuration
// * Indexing configuration
// * Advanced configuration
//
// Processing, embedding and advanced configuration are provided by this base class, while the indexing configuration is provided by the destination connector in the sub class.
type DestinationPgvector struct {
	destinationType Pgvector `const:"pgvector" json:"destinationType"`
	// Embedding configuration
	Embedding DestinationPgvectorEmbedding `json:"embedding"`
	// Postgres can be used to store vector data and retrieve embeddings.
	Indexing DestinationPgvectorPostgresConnection `json:"indexing"`
	// Do not store the text that gets embedded along with the vector and the metadata in the destination. If set to true, only the vector and the metadata will be stored - in this case raw text for LLM use cases needs to be retrieved from another source.
	OmitRawText *bool                                    `default:"false" json:"omit_raw_text"`
	Processing  DestinationPgvectorProcessingConfigModel `json:"processing"`
}

func (d DestinationPgvector) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPgvector) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPgvector) GetDestinationType() Pgvector {
	return PgvectorPgvector
}

func (o *DestinationPgvector) GetEmbedding() DestinationPgvectorEmbedding {
	if o == nil {
		return DestinationPgvectorEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationPgvector) GetIndexing() DestinationPgvectorPostgresConnection {
	if o == nil {
		return DestinationPgvectorPostgresConnection{}
	}
	return o.Indexing
}

func (o *DestinationPgvector) GetOmitRawText() *bool {
	if o == nil {
		return nil
	}
	return o.OmitRawText
}

func (o *DestinationPgvector) GetProcessing() DestinationPgvectorProcessingConfigModel {
	if o == nil {
		return DestinationPgvectorProcessingConfigModel{}
	}
	return o.Processing
}
