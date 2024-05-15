// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceGoogleDriveResourceModel) ToSharedSourceGoogleDriveCreateRequest() *shared.SourceGoogleDriveCreateRequest {
	var credentials shared.SourceGoogleDriveAuthentication
	var sourceGoogleDriveAuthenticateViaGoogleOAuth *shared.SourceGoogleDriveAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.AuthenticateViaGoogleOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleDriveAuthenticateViaGoogleOAuth = &shared.SourceGoogleDriveAuthenticateViaGoogleOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleDriveAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleDriveAuthentication{
			SourceGoogleDriveAuthenticateViaGoogleOAuth: sourceGoogleDriveAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleDriveServiceAccountKeyAuthentication *shared.SourceGoogleDriveServiceAccountKeyAuthentication
	if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
		serviceAccountInfo := r.Configuration.Credentials.ServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleDriveServiceAccountKeyAuthentication = &shared.SourceGoogleDriveServiceAccountKeyAuthentication{
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleDriveServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleDriveAuthentication{
			SourceGoogleDriveServiceAccountKeyAuthentication: sourceGoogleDriveServiceAccountKeyAuthentication,
		}
	}
	folderURL := r.Configuration.FolderURL.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGoogleDriveFileBasedStreamConfig = []shared.SourceGoogleDriveFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGoogleDriveFormat
		var sourceGoogleDriveAvroFormat *shared.SourceGoogleDriveAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceGoogleDriveAvroFormat = &shared.SourceGoogleDriveAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceGoogleDriveAvroFormat != nil {
			format = shared.SourceGoogleDriveFormat{
				SourceGoogleDriveAvroFormat: sourceGoogleDriveAvroFormat,
			}
		}
		var sourceGoogleDriveCSVFormat *shared.SourceGoogleDriveCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGoogleDriveCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGoogleDriveFromCSV *shared.SourceGoogleDriveFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGoogleDriveFromCSV = &shared.SourceGoogleDriveFromCSV{}
				}
				if sourceGoogleDriveFromCSV != nil {
					headerDefinition = &shared.SourceGoogleDriveCSVHeaderDefinition{
						SourceGoogleDriveFromCSV: sourceGoogleDriveFromCSV,
					}
				}
				var sourceGoogleDriveAutogenerated *shared.SourceGoogleDriveAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGoogleDriveAutogenerated = &shared.SourceGoogleDriveAutogenerated{}
				}
				if sourceGoogleDriveAutogenerated != nil {
					headerDefinition = &shared.SourceGoogleDriveCSVHeaderDefinition{
						SourceGoogleDriveAutogenerated: sourceGoogleDriveAutogenerated,
					}
				}
				var sourceGoogleDriveUserProvided *shared.SourceGoogleDriveUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGoogleDriveUserProvided = &shared.SourceGoogleDriveUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGoogleDriveUserProvided != nil {
					headerDefinition = &shared.SourceGoogleDriveCSVHeaderDefinition{
						SourceGoogleDriveUserProvided: sourceGoogleDriveUserProvided,
					}
				}
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGoogleDriveCSVFormat = &shared.SourceGoogleDriveCSVFormat{
				Delimiter:            delimiter,
				DoubleQuote:          doubleQuote,
				Encoding:             encoding,
				EscapeChar:           escapeChar,
				FalseValues:          falseValues,
				HeaderDefinition:     headerDefinition,
				NullValues:           nullValues,
				QuoteChar:            quoteChar,
				SkipRowsAfterHeader:  skipRowsAfterHeader,
				SkipRowsBeforeHeader: skipRowsBeforeHeader,
				StringsCanBeNull:     stringsCanBeNull,
				TrueValues:           trueValues,
			}
		}
		if sourceGoogleDriveCSVFormat != nil {
			format = shared.SourceGoogleDriveFormat{
				SourceGoogleDriveCSVFormat: sourceGoogleDriveCSVFormat,
			}
		}
		var sourceGoogleDriveJsonlFormat *shared.SourceGoogleDriveJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceGoogleDriveJsonlFormat = &shared.SourceGoogleDriveJsonlFormat{}
		}
		if sourceGoogleDriveJsonlFormat != nil {
			format = shared.SourceGoogleDriveFormat{
				SourceGoogleDriveJsonlFormat: sourceGoogleDriveJsonlFormat,
			}
		}
		var sourceGoogleDriveParquetFormat *shared.SourceGoogleDriveParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceGoogleDriveParquetFormat = &shared.SourceGoogleDriveParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceGoogleDriveParquetFormat != nil {
			format = shared.SourceGoogleDriveFormat{
				SourceGoogleDriveParquetFormat: sourceGoogleDriveParquetFormat,
			}
		}
		var sourceGoogleDriveDocumentFileTypeFormatExperimental *shared.SourceGoogleDriveDocumentFileTypeFormatExperimental
		if streamsItem.Format.DocumentFileTypeFormatExperimental != nil {
			var processing *shared.SourceGoogleDriveProcessing
			if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing != nil {
				var sourceGoogleDriveLocal *shared.SourceGoogleDriveLocal
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.Local != nil {
					sourceGoogleDriveLocal = &shared.SourceGoogleDriveLocal{}
				}
				if sourceGoogleDriveLocal != nil {
					processing = &shared.SourceGoogleDriveProcessing{
						SourceGoogleDriveLocal: sourceGoogleDriveLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceGoogleDriveParsingStrategy)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsNull() {
				*strategy = shared.SourceGoogleDriveParsingStrategy(streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceGoogleDriveDocumentFileTypeFormatExperimental = &shared.SourceGoogleDriveDocumentFileTypeFormatExperimental{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceGoogleDriveDocumentFileTypeFormatExperimental != nil {
			format = shared.SourceGoogleDriveFormat{
				SourceGoogleDriveDocumentFileTypeFormatExperimental: sourceGoogleDriveDocumentFileTypeFormatExperimental,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		name := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGoogleDriveValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGoogleDriveValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGoogleDriveFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			Name:                      name,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceGoogleDrive{
		Credentials: credentials,
		FolderURL:   folderURL,
		StartDate:   startDate,
		Streams:     streams,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleDriveCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceGoogleDriveResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceGoogleDriveResourceModel) ToSharedSourceGoogleDrivePutRequest() *shared.SourceGoogleDrivePutRequest {
	var credentials shared.SourceGoogleDriveUpdateAuthentication
	var sourceGoogleDriveUpdateAuthenticateViaGoogleOAuth *shared.SourceGoogleDriveUpdateAuthenticateViaGoogleOAuth
	if r.Configuration.Credentials.AuthenticateViaGoogleOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaGoogleOAuth.RefreshToken.ValueString()
		sourceGoogleDriveUpdateAuthenticateViaGoogleOAuth = &shared.SourceGoogleDriveUpdateAuthenticateViaGoogleOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
		}
	}
	if sourceGoogleDriveUpdateAuthenticateViaGoogleOAuth != nil {
		credentials = shared.SourceGoogleDriveUpdateAuthentication{
			SourceGoogleDriveUpdateAuthenticateViaGoogleOAuth: sourceGoogleDriveUpdateAuthenticateViaGoogleOAuth,
		}
	}
	var sourceGoogleDriveUpdateServiceAccountKeyAuthentication *shared.SourceGoogleDriveUpdateServiceAccountKeyAuthentication
	if r.Configuration.Credentials.ServiceAccountKeyAuthentication != nil {
		serviceAccountInfo := r.Configuration.Credentials.ServiceAccountKeyAuthentication.ServiceAccountInfo.ValueString()
		sourceGoogleDriveUpdateServiceAccountKeyAuthentication = &shared.SourceGoogleDriveUpdateServiceAccountKeyAuthentication{
			ServiceAccountInfo: serviceAccountInfo,
		}
	}
	if sourceGoogleDriveUpdateServiceAccountKeyAuthentication != nil {
		credentials = shared.SourceGoogleDriveUpdateAuthentication{
			SourceGoogleDriveUpdateServiceAccountKeyAuthentication: sourceGoogleDriveUpdateServiceAccountKeyAuthentication,
		}
	}
	folderURL := r.Configuration.FolderURL.ValueString()
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceGoogleDriveUpdateFileBasedStreamConfig = []shared.SourceGoogleDriveUpdateFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceGoogleDriveUpdateFormat
		var sourceGoogleDriveUpdateAvroFormat *shared.SourceGoogleDriveUpdateAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceGoogleDriveUpdateAvroFormat = &shared.SourceGoogleDriveUpdateAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceGoogleDriveUpdateAvroFormat != nil {
			format = shared.SourceGoogleDriveUpdateFormat{
				SourceGoogleDriveUpdateAvroFormat: sourceGoogleDriveUpdateAvroFormat,
			}
		}
		var sourceGoogleDriveUpdateCSVFormat *shared.SourceGoogleDriveUpdateCSVFormat
		if streamsItem.Format.CSVFormat != nil {
			delimiter := new(string)
			if !streamsItem.Format.CSVFormat.Delimiter.IsUnknown() && !streamsItem.Format.CSVFormat.Delimiter.IsNull() {
				*delimiter = streamsItem.Format.CSVFormat.Delimiter.ValueString()
			} else {
				delimiter = nil
			}
			doubleQuote := new(bool)
			if !streamsItem.Format.CSVFormat.DoubleQuote.IsUnknown() && !streamsItem.Format.CSVFormat.DoubleQuote.IsNull() {
				*doubleQuote = streamsItem.Format.CSVFormat.DoubleQuote.ValueBool()
			} else {
				doubleQuote = nil
			}
			encoding := new(string)
			if !streamsItem.Format.CSVFormat.Encoding.IsUnknown() && !streamsItem.Format.CSVFormat.Encoding.IsNull() {
				*encoding = streamsItem.Format.CSVFormat.Encoding.ValueString()
			} else {
				encoding = nil
			}
			escapeChar := new(string)
			if !streamsItem.Format.CSVFormat.EscapeChar.IsUnknown() && !streamsItem.Format.CSVFormat.EscapeChar.IsNull() {
				*escapeChar = streamsItem.Format.CSVFormat.EscapeChar.ValueString()
			} else {
				escapeChar = nil
			}
			var falseValues []string = []string{}
			for _, falseValuesItem := range streamsItem.Format.CSVFormat.FalseValues {
				falseValues = append(falseValues, falseValuesItem.ValueString())
			}
			var headerDefinition *shared.SourceGoogleDriveUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceGoogleDriveUpdateFromCSV *shared.SourceGoogleDriveUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceGoogleDriveUpdateFromCSV = &shared.SourceGoogleDriveUpdateFromCSV{}
				}
				if sourceGoogleDriveUpdateFromCSV != nil {
					headerDefinition = &shared.SourceGoogleDriveUpdateCSVHeaderDefinition{
						SourceGoogleDriveUpdateFromCSV: sourceGoogleDriveUpdateFromCSV,
					}
				}
				var sourceGoogleDriveUpdateAutogenerated *shared.SourceGoogleDriveUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceGoogleDriveUpdateAutogenerated = &shared.SourceGoogleDriveUpdateAutogenerated{}
				}
				if sourceGoogleDriveUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceGoogleDriveUpdateCSVHeaderDefinition{
						SourceGoogleDriveUpdateAutogenerated: sourceGoogleDriveUpdateAutogenerated,
					}
				}
				var sourceGoogleDriveUpdateUserProvided *shared.SourceGoogleDriveUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceGoogleDriveUpdateUserProvided = &shared.SourceGoogleDriveUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceGoogleDriveUpdateUserProvided != nil {
					headerDefinition = &shared.SourceGoogleDriveUpdateCSVHeaderDefinition{
						SourceGoogleDriveUpdateUserProvided: sourceGoogleDriveUpdateUserProvided,
					}
				}
			}
			var nullValues []string = []string{}
			for _, nullValuesItem := range streamsItem.Format.CSVFormat.NullValues {
				nullValues = append(nullValues, nullValuesItem.ValueString())
			}
			quoteChar := new(string)
			if !streamsItem.Format.CSVFormat.QuoteChar.IsUnknown() && !streamsItem.Format.CSVFormat.QuoteChar.IsNull() {
				*quoteChar = streamsItem.Format.CSVFormat.QuoteChar.ValueString()
			} else {
				quoteChar = nil
			}
			skipRowsAfterHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsAfterHeader.IsNull() {
				*skipRowsAfterHeader = streamsItem.Format.CSVFormat.SkipRowsAfterHeader.ValueInt64()
			} else {
				skipRowsAfterHeader = nil
			}
			skipRowsBeforeHeader := new(int64)
			if !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsUnknown() && !streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.IsNull() {
				*skipRowsBeforeHeader = streamsItem.Format.CSVFormat.SkipRowsBeforeHeader.ValueInt64()
			} else {
				skipRowsBeforeHeader = nil
			}
			stringsCanBeNull := new(bool)
			if !streamsItem.Format.CSVFormat.StringsCanBeNull.IsUnknown() && !streamsItem.Format.CSVFormat.StringsCanBeNull.IsNull() {
				*stringsCanBeNull = streamsItem.Format.CSVFormat.StringsCanBeNull.ValueBool()
			} else {
				stringsCanBeNull = nil
			}
			var trueValues []string = []string{}
			for _, trueValuesItem := range streamsItem.Format.CSVFormat.TrueValues {
				trueValues = append(trueValues, trueValuesItem.ValueString())
			}
			sourceGoogleDriveUpdateCSVFormat = &shared.SourceGoogleDriveUpdateCSVFormat{
				Delimiter:            delimiter,
				DoubleQuote:          doubleQuote,
				Encoding:             encoding,
				EscapeChar:           escapeChar,
				FalseValues:          falseValues,
				HeaderDefinition:     headerDefinition,
				NullValues:           nullValues,
				QuoteChar:            quoteChar,
				SkipRowsAfterHeader:  skipRowsAfterHeader,
				SkipRowsBeforeHeader: skipRowsBeforeHeader,
				StringsCanBeNull:     stringsCanBeNull,
				TrueValues:           trueValues,
			}
		}
		if sourceGoogleDriveUpdateCSVFormat != nil {
			format = shared.SourceGoogleDriveUpdateFormat{
				SourceGoogleDriveUpdateCSVFormat: sourceGoogleDriveUpdateCSVFormat,
			}
		}
		var sourceGoogleDriveUpdateJsonlFormat *shared.SourceGoogleDriveUpdateJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceGoogleDriveUpdateJsonlFormat = &shared.SourceGoogleDriveUpdateJsonlFormat{}
		}
		if sourceGoogleDriveUpdateJsonlFormat != nil {
			format = shared.SourceGoogleDriveUpdateFormat{
				SourceGoogleDriveUpdateJsonlFormat: sourceGoogleDriveUpdateJsonlFormat,
			}
		}
		var sourceGoogleDriveUpdateParquetFormat *shared.SourceGoogleDriveUpdateParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceGoogleDriveUpdateParquetFormat = &shared.SourceGoogleDriveUpdateParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceGoogleDriveUpdateParquetFormat != nil {
			format = shared.SourceGoogleDriveUpdateFormat{
				SourceGoogleDriveUpdateParquetFormat: sourceGoogleDriveUpdateParquetFormat,
			}
		}
		var sourceGoogleDriveUpdateDocumentFileTypeFormatExperimental *shared.SourceGoogleDriveUpdateDocumentFileTypeFormatExperimental
		if streamsItem.Format.DocumentFileTypeFormatExperimental != nil {
			var processing *shared.SourceGoogleDriveUpdateProcessing
			if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing != nil {
				var sourceGoogleDriveUpdateLocal *shared.SourceGoogleDriveUpdateLocal
				if streamsItem.Format.DocumentFileTypeFormatExperimental.Processing.Local != nil {
					sourceGoogleDriveUpdateLocal = &shared.SourceGoogleDriveUpdateLocal{}
				}
				if sourceGoogleDriveUpdateLocal != nil {
					processing = &shared.SourceGoogleDriveUpdateProcessing{
						SourceGoogleDriveUpdateLocal: sourceGoogleDriveUpdateLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.DocumentFileTypeFormatExperimental.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceGoogleDriveUpdateParsingStrategy)
			if !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsUnknown() && !streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.IsNull() {
				*strategy = shared.SourceGoogleDriveUpdateParsingStrategy(streamsItem.Format.DocumentFileTypeFormatExperimental.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceGoogleDriveUpdateDocumentFileTypeFormatExperimental = &shared.SourceGoogleDriveUpdateDocumentFileTypeFormatExperimental{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceGoogleDriveUpdateDocumentFileTypeFormatExperimental != nil {
			format = shared.SourceGoogleDriveUpdateFormat{
				SourceGoogleDriveUpdateDocumentFileTypeFormatExperimental: sourceGoogleDriveUpdateDocumentFileTypeFormatExperimental,
			}
		}
		var globs []string = []string{}
		for _, globsItem := range streamsItem.Globs {
			globs = append(globs, globsItem.ValueString())
		}
		inputSchema := new(string)
		if !streamsItem.InputSchema.IsUnknown() && !streamsItem.InputSchema.IsNull() {
			*inputSchema = streamsItem.InputSchema.ValueString()
		} else {
			inputSchema = nil
		}
		name := streamsItem.Name.ValueString()
		primaryKey := new(string)
		if !streamsItem.PrimaryKey.IsUnknown() && !streamsItem.PrimaryKey.IsNull() {
			*primaryKey = streamsItem.PrimaryKey.ValueString()
		} else {
			primaryKey = nil
		}
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceGoogleDriveUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceGoogleDriveUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceGoogleDriveUpdateFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			Name:                      name,
			PrimaryKey:                primaryKey,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceGoogleDriveUpdate{
		Credentials: credentials,
		FolderURL:   folderURL,
		StartDate:   startDate,
		Streams:     streams,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceGoogleDrivePutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}
