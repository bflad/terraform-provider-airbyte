// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceMicrosoftOnedriveResourceModel) ToSharedSourceMicrosoftOnedriveCreateRequest() *shared.SourceMicrosoftOnedriveCreateRequest {
	var credentials shared.SourceMicrosoftOnedriveAuthentication
	var sourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth *shared.SourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth
	if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.ValueString()
		tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.TenantID.ValueString()
		sourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth = &shared.SourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
			TenantID:     tenantID,
		}
	}
	if sourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth != nil {
		credentials = shared.SourceMicrosoftOnedriveAuthentication{
			SourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth: sourceMicrosoftOnedriveAuthenticateViaMicrosoftOAuth,
		}
	}
	var sourceMicrosoftOnedriveServiceKeyAuthentication *shared.SourceMicrosoftOnedriveServiceKeyAuthentication
	if r.Configuration.Credentials.ServiceKeyAuthentication != nil {
		clientId1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientID.ValueString()
		clientSecret1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientSecret.ValueString()
		tenantId1 := r.Configuration.Credentials.ServiceKeyAuthentication.TenantID.ValueString()
		userPrincipalName := r.Configuration.Credentials.ServiceKeyAuthentication.UserPrincipalName.ValueString()
		sourceMicrosoftOnedriveServiceKeyAuthentication = &shared.SourceMicrosoftOnedriveServiceKeyAuthentication{
			ClientID:          clientId1,
			ClientSecret:      clientSecret1,
			TenantID:          tenantId1,
			UserPrincipalName: userPrincipalName,
		}
	}
	if sourceMicrosoftOnedriveServiceKeyAuthentication != nil {
		credentials = shared.SourceMicrosoftOnedriveAuthentication{
			SourceMicrosoftOnedriveServiceKeyAuthentication: sourceMicrosoftOnedriveServiceKeyAuthentication,
		}
	}
	driveName := new(string)
	if !r.Configuration.DriveName.IsUnknown() && !r.Configuration.DriveName.IsNull() {
		*driveName = r.Configuration.DriveName.ValueString()
	} else {
		driveName = nil
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	searchScope := new(shared.SourceMicrosoftOnedriveSearchScope)
	if !r.Configuration.SearchScope.IsUnknown() && !r.Configuration.SearchScope.IsNull() {
		*searchScope = shared.SourceMicrosoftOnedriveSearchScope(r.Configuration.SearchScope.ValueString())
	} else {
		searchScope = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceMicrosoftOnedriveFileBasedStreamConfig = []shared.SourceMicrosoftOnedriveFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceMicrosoftOnedriveFormat
		var sourceMicrosoftOnedriveAvroFormat *shared.SourceMicrosoftOnedriveAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceMicrosoftOnedriveAvroFormat = &shared.SourceMicrosoftOnedriveAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceMicrosoftOnedriveAvroFormat != nil {
			format = shared.SourceMicrosoftOnedriveFormat{
				SourceMicrosoftOnedriveAvroFormat: sourceMicrosoftOnedriveAvroFormat,
			}
		}
		var sourceMicrosoftOnedriveCSVFormat *shared.SourceMicrosoftOnedriveCSVFormat
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
			var headerDefinition *shared.SourceMicrosoftOnedriveCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceMicrosoftOnedriveFromCSV *shared.SourceMicrosoftOnedriveFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceMicrosoftOnedriveFromCSV = &shared.SourceMicrosoftOnedriveFromCSV{}
				}
				if sourceMicrosoftOnedriveFromCSV != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveCSVHeaderDefinition{
						SourceMicrosoftOnedriveFromCSV: sourceMicrosoftOnedriveFromCSV,
					}
				}
				var sourceMicrosoftOnedriveAutogenerated *shared.SourceMicrosoftOnedriveAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceMicrosoftOnedriveAutogenerated = &shared.SourceMicrosoftOnedriveAutogenerated{}
				}
				if sourceMicrosoftOnedriveAutogenerated != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveCSVHeaderDefinition{
						SourceMicrosoftOnedriveAutogenerated: sourceMicrosoftOnedriveAutogenerated,
					}
				}
				var sourceMicrosoftOnedriveUserProvided *shared.SourceMicrosoftOnedriveUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceMicrosoftOnedriveUserProvided = &shared.SourceMicrosoftOnedriveUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceMicrosoftOnedriveUserProvided != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveCSVHeaderDefinition{
						SourceMicrosoftOnedriveUserProvided: sourceMicrosoftOnedriveUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
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
			sourceMicrosoftOnedriveCSVFormat = &shared.SourceMicrosoftOnedriveCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceMicrosoftOnedriveCSVFormat != nil {
			format = shared.SourceMicrosoftOnedriveFormat{
				SourceMicrosoftOnedriveCSVFormat: sourceMicrosoftOnedriveCSVFormat,
			}
		}
		var sourceMicrosoftOnedriveJsonlFormat *shared.SourceMicrosoftOnedriveJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceMicrosoftOnedriveJsonlFormat = &shared.SourceMicrosoftOnedriveJsonlFormat{}
		}
		if sourceMicrosoftOnedriveJsonlFormat != nil {
			format = shared.SourceMicrosoftOnedriveFormat{
				SourceMicrosoftOnedriveJsonlFormat: sourceMicrosoftOnedriveJsonlFormat,
			}
		}
		var sourceMicrosoftOnedriveParquetFormat *shared.SourceMicrosoftOnedriveParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceMicrosoftOnedriveParquetFormat = &shared.SourceMicrosoftOnedriveParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceMicrosoftOnedriveParquetFormat != nil {
			format = shared.SourceMicrosoftOnedriveFormat{
				SourceMicrosoftOnedriveParquetFormat: sourceMicrosoftOnedriveParquetFormat,
			}
		}
		var sourceMicrosoftOnedriveUnstructuredDocumentFormat *shared.SourceMicrosoftOnedriveUnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceMicrosoftOnedriveProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceMicrosoftOnedriveLocal *shared.SourceMicrosoftOnedriveLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceMicrosoftOnedriveLocal = &shared.SourceMicrosoftOnedriveLocal{}
				}
				if sourceMicrosoftOnedriveLocal != nil {
					processing = &shared.SourceMicrosoftOnedriveProcessing{
						SourceMicrosoftOnedriveLocal: sourceMicrosoftOnedriveLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceMicrosoftOnedriveParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceMicrosoftOnedriveParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceMicrosoftOnedriveUnstructuredDocumentFormat = &shared.SourceMicrosoftOnedriveUnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceMicrosoftOnedriveUnstructuredDocumentFormat != nil {
			format = shared.SourceMicrosoftOnedriveFormat{
				SourceMicrosoftOnedriveUnstructuredDocumentFormat: sourceMicrosoftOnedriveUnstructuredDocumentFormat,
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
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceMicrosoftOnedriveValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceMicrosoftOnedriveValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceMicrosoftOnedriveFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			Name:                      name,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceMicrosoftOnedrive{
		Credentials: credentials,
		DriveName:   driveName,
		FolderPath:  folderPath,
		SearchScope: searchScope,
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
	out := shared.SourceMicrosoftOnedriveCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMicrosoftOnedriveResourceModel) RefreshFromSharedSourceResponse(resp *shared.SourceResponse) {
	if resp != nil {
		r.Name = types.StringValue(resp.Name)
		r.SourceID = types.StringValue(resp.SourceID)
		r.SourceType = types.StringValue(resp.SourceType)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *SourceMicrosoftOnedriveResourceModel) ToSharedSourceMicrosoftOnedrivePutRequest() *shared.SourceMicrosoftOnedrivePutRequest {
	var credentials shared.SourceMicrosoftOnedriveUpdateAuthentication
	var authenticateViaMicrosoftOAuth *shared.AuthenticateViaMicrosoftOAuth
	if r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth != nil {
		clientID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.RefreshToken.ValueString()
		tenantID := r.Configuration.Credentials.AuthenticateViaMicrosoftOAuth.TenantID.ValueString()
		authenticateViaMicrosoftOAuth = &shared.AuthenticateViaMicrosoftOAuth{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RefreshToken: refreshToken,
			TenantID:     tenantID,
		}
	}
	if authenticateViaMicrosoftOAuth != nil {
		credentials = shared.SourceMicrosoftOnedriveUpdateAuthentication{
			AuthenticateViaMicrosoftOAuth: authenticateViaMicrosoftOAuth,
		}
	}
	var serviceKeyAuthentication *shared.ServiceKeyAuthentication
	if r.Configuration.Credentials.ServiceKeyAuthentication != nil {
		clientId1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientID.ValueString()
		clientSecret1 := r.Configuration.Credentials.ServiceKeyAuthentication.ClientSecret.ValueString()
		tenantId1 := r.Configuration.Credentials.ServiceKeyAuthentication.TenantID.ValueString()
		userPrincipalName := r.Configuration.Credentials.ServiceKeyAuthentication.UserPrincipalName.ValueString()
		serviceKeyAuthentication = &shared.ServiceKeyAuthentication{
			ClientID:          clientId1,
			ClientSecret:      clientSecret1,
			TenantID:          tenantId1,
			UserPrincipalName: userPrincipalName,
		}
	}
	if serviceKeyAuthentication != nil {
		credentials = shared.SourceMicrosoftOnedriveUpdateAuthentication{
			ServiceKeyAuthentication: serviceKeyAuthentication,
		}
	}
	driveName := new(string)
	if !r.Configuration.DriveName.IsUnknown() && !r.Configuration.DriveName.IsNull() {
		*driveName = r.Configuration.DriveName.ValueString()
	} else {
		driveName = nil
	}
	folderPath := new(string)
	if !r.Configuration.FolderPath.IsUnknown() && !r.Configuration.FolderPath.IsNull() {
		*folderPath = r.Configuration.FolderPath.ValueString()
	} else {
		folderPath = nil
	}
	searchScope := new(shared.SearchScope)
	if !r.Configuration.SearchScope.IsUnknown() && !r.Configuration.SearchScope.IsNull() {
		*searchScope = shared.SearchScope(r.Configuration.SearchScope.ValueString())
	} else {
		searchScope = nil
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	var streams []shared.SourceMicrosoftOnedriveUpdateFileBasedStreamConfig = []shared.SourceMicrosoftOnedriveUpdateFileBasedStreamConfig{}
	for _, streamsItem := range r.Configuration.Streams {
		daysToSyncIfHistoryIsFull := new(int64)
		if !streamsItem.DaysToSyncIfHistoryIsFull.IsUnknown() && !streamsItem.DaysToSyncIfHistoryIsFull.IsNull() {
			*daysToSyncIfHistoryIsFull = streamsItem.DaysToSyncIfHistoryIsFull.ValueInt64()
		} else {
			daysToSyncIfHistoryIsFull = nil
		}
		var format shared.SourceMicrosoftOnedriveUpdateFormat
		var sourceMicrosoftOnedriveUpdateAvroFormat *shared.SourceMicrosoftOnedriveUpdateAvroFormat
		if streamsItem.Format.AvroFormat != nil {
			doubleAsString := new(bool)
			if !streamsItem.Format.AvroFormat.DoubleAsString.IsUnknown() && !streamsItem.Format.AvroFormat.DoubleAsString.IsNull() {
				*doubleAsString = streamsItem.Format.AvroFormat.DoubleAsString.ValueBool()
			} else {
				doubleAsString = nil
			}
			sourceMicrosoftOnedriveUpdateAvroFormat = &shared.SourceMicrosoftOnedriveUpdateAvroFormat{
				DoubleAsString: doubleAsString,
			}
		}
		if sourceMicrosoftOnedriveUpdateAvroFormat != nil {
			format = shared.SourceMicrosoftOnedriveUpdateFormat{
				SourceMicrosoftOnedriveUpdateAvroFormat: sourceMicrosoftOnedriveUpdateAvroFormat,
			}
		}
		var sourceMicrosoftOnedriveUpdateCSVFormat *shared.SourceMicrosoftOnedriveUpdateCSVFormat
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
			var headerDefinition *shared.SourceMicrosoftOnedriveUpdateCSVHeaderDefinition
			if streamsItem.Format.CSVFormat.HeaderDefinition != nil {
				var sourceMicrosoftOnedriveUpdateFromCSV *shared.SourceMicrosoftOnedriveUpdateFromCSV
				if streamsItem.Format.CSVFormat.HeaderDefinition.FromCSV != nil {
					sourceMicrosoftOnedriveUpdateFromCSV = &shared.SourceMicrosoftOnedriveUpdateFromCSV{}
				}
				if sourceMicrosoftOnedriveUpdateFromCSV != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveUpdateCSVHeaderDefinition{
						SourceMicrosoftOnedriveUpdateFromCSV: sourceMicrosoftOnedriveUpdateFromCSV,
					}
				}
				var sourceMicrosoftOnedriveUpdateAutogenerated *shared.SourceMicrosoftOnedriveUpdateAutogenerated
				if streamsItem.Format.CSVFormat.HeaderDefinition.Autogenerated != nil {
					sourceMicrosoftOnedriveUpdateAutogenerated = &shared.SourceMicrosoftOnedriveUpdateAutogenerated{}
				}
				if sourceMicrosoftOnedriveUpdateAutogenerated != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveUpdateCSVHeaderDefinition{
						SourceMicrosoftOnedriveUpdateAutogenerated: sourceMicrosoftOnedriveUpdateAutogenerated,
					}
				}
				var sourceMicrosoftOnedriveUpdateUserProvided *shared.SourceMicrosoftOnedriveUpdateUserProvided
				if streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided != nil {
					var columnNames []string = []string{}
					for _, columnNamesItem := range streamsItem.Format.CSVFormat.HeaderDefinition.UserProvided.ColumnNames {
						columnNames = append(columnNames, columnNamesItem.ValueString())
					}
					sourceMicrosoftOnedriveUpdateUserProvided = &shared.SourceMicrosoftOnedriveUpdateUserProvided{
						ColumnNames: columnNames,
					}
				}
				if sourceMicrosoftOnedriveUpdateUserProvided != nil {
					headerDefinition = &shared.SourceMicrosoftOnedriveUpdateCSVHeaderDefinition{
						SourceMicrosoftOnedriveUpdateUserProvided: sourceMicrosoftOnedriveUpdateUserProvided,
					}
				}
			}
			ignoreErrorsOnFieldsMismatch := new(bool)
			if !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsUnknown() && !streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.IsNull() {
				*ignoreErrorsOnFieldsMismatch = streamsItem.Format.CSVFormat.IgnoreErrorsOnFieldsMismatch.ValueBool()
			} else {
				ignoreErrorsOnFieldsMismatch = nil
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
			sourceMicrosoftOnedriveUpdateCSVFormat = &shared.SourceMicrosoftOnedriveUpdateCSVFormat{
				Delimiter:                    delimiter,
				DoubleQuote:                  doubleQuote,
				Encoding:                     encoding,
				EscapeChar:                   escapeChar,
				FalseValues:                  falseValues,
				HeaderDefinition:             headerDefinition,
				IgnoreErrorsOnFieldsMismatch: ignoreErrorsOnFieldsMismatch,
				NullValues:                   nullValues,
				QuoteChar:                    quoteChar,
				SkipRowsAfterHeader:          skipRowsAfterHeader,
				SkipRowsBeforeHeader:         skipRowsBeforeHeader,
				StringsCanBeNull:             stringsCanBeNull,
				TrueValues:                   trueValues,
			}
		}
		if sourceMicrosoftOnedriveUpdateCSVFormat != nil {
			format = shared.SourceMicrosoftOnedriveUpdateFormat{
				SourceMicrosoftOnedriveUpdateCSVFormat: sourceMicrosoftOnedriveUpdateCSVFormat,
			}
		}
		var sourceMicrosoftOnedriveUpdateJsonlFormat *shared.SourceMicrosoftOnedriveUpdateJsonlFormat
		if streamsItem.Format.JsonlFormat != nil {
			sourceMicrosoftOnedriveUpdateJsonlFormat = &shared.SourceMicrosoftOnedriveUpdateJsonlFormat{}
		}
		if sourceMicrosoftOnedriveUpdateJsonlFormat != nil {
			format = shared.SourceMicrosoftOnedriveUpdateFormat{
				SourceMicrosoftOnedriveUpdateJsonlFormat: sourceMicrosoftOnedriveUpdateJsonlFormat,
			}
		}
		var sourceMicrosoftOnedriveUpdateParquetFormat *shared.SourceMicrosoftOnedriveUpdateParquetFormat
		if streamsItem.Format.ParquetFormat != nil {
			decimalAsFloat := new(bool)
			if !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsUnknown() && !streamsItem.Format.ParquetFormat.DecimalAsFloat.IsNull() {
				*decimalAsFloat = streamsItem.Format.ParquetFormat.DecimalAsFloat.ValueBool()
			} else {
				decimalAsFloat = nil
			}
			sourceMicrosoftOnedriveUpdateParquetFormat = &shared.SourceMicrosoftOnedriveUpdateParquetFormat{
				DecimalAsFloat: decimalAsFloat,
			}
		}
		if sourceMicrosoftOnedriveUpdateParquetFormat != nil {
			format = shared.SourceMicrosoftOnedriveUpdateFormat{
				SourceMicrosoftOnedriveUpdateParquetFormat: sourceMicrosoftOnedriveUpdateParquetFormat,
			}
		}
		var sourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat *shared.SourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat
		if streamsItem.Format.UnstructuredDocumentFormat != nil {
			var processing *shared.SourceMicrosoftOnedriveUpdateProcessing
			if streamsItem.Format.UnstructuredDocumentFormat.Processing != nil {
				var sourceMicrosoftOnedriveUpdateLocal *shared.SourceMicrosoftOnedriveUpdateLocal
				if streamsItem.Format.UnstructuredDocumentFormat.Processing.Local != nil {
					sourceMicrosoftOnedriveUpdateLocal = &shared.SourceMicrosoftOnedriveUpdateLocal{}
				}
				if sourceMicrosoftOnedriveUpdateLocal != nil {
					processing = &shared.SourceMicrosoftOnedriveUpdateProcessing{
						SourceMicrosoftOnedriveUpdateLocal: sourceMicrosoftOnedriveUpdateLocal,
					}
				}
			}
			skipUnprocessableFiles := new(bool)
			if !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.IsNull() {
				*skipUnprocessableFiles = streamsItem.Format.UnstructuredDocumentFormat.SkipUnprocessableFiles.ValueBool()
			} else {
				skipUnprocessableFiles = nil
			}
			strategy := new(shared.SourceMicrosoftOnedriveUpdateParsingStrategy)
			if !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsUnknown() && !streamsItem.Format.UnstructuredDocumentFormat.Strategy.IsNull() {
				*strategy = shared.SourceMicrosoftOnedriveUpdateParsingStrategy(streamsItem.Format.UnstructuredDocumentFormat.Strategy.ValueString())
			} else {
				strategy = nil
			}
			sourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat = &shared.SourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat{
				Processing:             processing,
				SkipUnprocessableFiles: skipUnprocessableFiles,
				Strategy:               strategy,
			}
		}
		if sourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat != nil {
			format = shared.SourceMicrosoftOnedriveUpdateFormat{
				SourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat: sourceMicrosoftOnedriveUpdateUnstructuredDocumentFormat,
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
		schemaless := new(bool)
		if !streamsItem.Schemaless.IsUnknown() && !streamsItem.Schemaless.IsNull() {
			*schemaless = streamsItem.Schemaless.ValueBool()
		} else {
			schemaless = nil
		}
		validationPolicy := new(shared.SourceMicrosoftOnedriveUpdateValidationPolicy)
		if !streamsItem.ValidationPolicy.IsUnknown() && !streamsItem.ValidationPolicy.IsNull() {
			*validationPolicy = shared.SourceMicrosoftOnedriveUpdateValidationPolicy(streamsItem.ValidationPolicy.ValueString())
		} else {
			validationPolicy = nil
		}
		streams = append(streams, shared.SourceMicrosoftOnedriveUpdateFileBasedStreamConfig{
			DaysToSyncIfHistoryIsFull: daysToSyncIfHistoryIsFull,
			Format:                    format,
			Globs:                     globs,
			InputSchema:               inputSchema,
			Name:                      name,
			Schemaless:                schemaless,
			ValidationPolicy:          validationPolicy,
		})
	}
	configuration := shared.SourceMicrosoftOnedriveUpdate{
		Credentials: credentials,
		DriveName:   driveName,
		FolderPath:  folderPath,
		SearchScope: searchScope,
		StartDate:   startDate,
		Streams:     streams,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMicrosoftOnedrivePutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}
