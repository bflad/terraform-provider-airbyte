// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationPineconeResourceModel) ToSharedDestinationPineconeCreateRequest() *shared.DestinationPineconeCreateRequest {
	var embedding shared.DestinationPineconeEmbedding
	var destinationPineconeOpenAI *shared.DestinationPineconeOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationPineconeOpenAI = &shared.DestinationPineconeOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationPineconeOpenAI != nil {
		embedding = shared.DestinationPineconeEmbedding{
			DestinationPineconeOpenAI: destinationPineconeOpenAI,
		}
	}
	var destinationPineconeCohere *shared.DestinationPineconeCohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		destinationPineconeCohere = &shared.DestinationPineconeCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationPineconeCohere != nil {
		embedding = shared.DestinationPineconeEmbedding{
			DestinationPineconeCohere: destinationPineconeCohere,
		}
	}
	var destinationPineconeFake *shared.DestinationPineconeFake
	if r.Configuration.Embedding.Fake != nil {
		destinationPineconeFake = &shared.DestinationPineconeFake{}
	}
	if destinationPineconeFake != nil {
		embedding = shared.DestinationPineconeEmbedding{
			DestinationPineconeFake: destinationPineconeFake,
		}
	}
	var destinationPineconeAzureOpenAI *shared.DestinationPineconeAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		destinationPineconeAzureOpenAI = &shared.DestinationPineconeAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationPineconeAzureOpenAI != nil {
		embedding = shared.DestinationPineconeEmbedding{
			DestinationPineconeAzureOpenAI: destinationPineconeAzureOpenAI,
		}
	}
	var destinationPineconeOpenAICompatible *shared.DestinationPineconeOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationPineconeOpenAICompatible = &shared.DestinationPineconeOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if destinationPineconeOpenAICompatible != nil {
		embedding = shared.DestinationPineconeEmbedding{
			DestinationPineconeOpenAICompatible: destinationPineconeOpenAICompatible,
		}
	}
	index := r.Configuration.Indexing.Index.ValueString()
	pineconeEnvironment := r.Configuration.Indexing.PineconeEnvironment.ValueString()
	pineconeKey := r.Configuration.Indexing.PineconeKey.ValueString()
	indexing := shared.DestinationPineconeIndexing{
		Index:               index,
		PineconeEnvironment: pineconeEnvironment,
		PineconeKey:         pineconeKey,
	}
	omitRawText := new(bool)
	if !r.Configuration.OmitRawText.IsUnknown() && !r.Configuration.OmitRawText.IsNull() {
		*omitRawText = r.Configuration.OmitRawText.ValueBool()
	} else {
		omitRawText = nil
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.DestinationPineconeFieldNameMappingConfigModel = []shared.DestinationPineconeFieldNameMappingConfigModel{}
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.DestinationPineconeFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = []string{}
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = []string{}
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationPineconeTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationPineconeBySeparator *shared.DestinationPineconeBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = []string{}
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationPineconeBySeparator = &shared.DestinationPineconeBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationPineconeBySeparator != nil {
			textSplitter = &shared.DestinationPineconeTextSplitter{
				DestinationPineconeBySeparator: destinationPineconeBySeparator,
			}
		}
		var destinationPineconeByMarkdownHeader *shared.DestinationPineconeByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationPineconeByMarkdownHeader = &shared.DestinationPineconeByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationPineconeByMarkdownHeader != nil {
			textSplitter = &shared.DestinationPineconeTextSplitter{
				DestinationPineconeByMarkdownHeader: destinationPineconeByMarkdownHeader,
			}
		}
		var destinationPineconeByProgrammingLanguage *shared.DestinationPineconeByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationPineconeLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationPineconeByProgrammingLanguage = &shared.DestinationPineconeByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationPineconeByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationPineconeTextSplitter{
				DestinationPineconeByProgrammingLanguage: destinationPineconeByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationPineconeProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationPinecone{
		Embedding:   embedding,
		Indexing:    indexing,
		OmitRawText: omitRawText,
		Processing:  processing,
	}
	definitionID := new(string)
	if !r.DefinitionID.IsUnknown() && !r.DefinitionID.IsNull() {
		*definitionID = r.DefinitionID.ValueString()
	} else {
		definitionID = nil
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPineconeCreateRequest{
		Configuration: configuration,
		DefinitionID:  definitionID,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPineconeResourceModel) RefreshFromSharedDestinationResponse(resp *shared.DestinationResponse) {
	if resp != nil {
		r.DestinationID = types.StringValue(resp.DestinationID)
		r.DestinationType = types.StringValue(resp.DestinationType)
		r.Name = types.StringValue(resp.Name)
		r.WorkspaceID = types.StringValue(resp.WorkspaceID)
	}
}

func (r *DestinationPineconeResourceModel) ToSharedDestinationPineconePutRequest() *shared.DestinationPineconePutRequest {
	var embedding shared.DestinationPineconeUpdateEmbedding
	var destinationPineconeUpdateOpenAI *shared.DestinationPineconeUpdateOpenAI
	if r.Configuration.Embedding.OpenAI != nil {
		openaiKey := r.Configuration.Embedding.OpenAI.OpenaiKey.ValueString()
		destinationPineconeUpdateOpenAI = &shared.DestinationPineconeUpdateOpenAI{
			OpenaiKey: openaiKey,
		}
	}
	if destinationPineconeUpdateOpenAI != nil {
		embedding = shared.DestinationPineconeUpdateEmbedding{
			DestinationPineconeUpdateOpenAI: destinationPineconeUpdateOpenAI,
		}
	}
	var destinationPineconeUpdateCohere *shared.DestinationPineconeUpdateCohere
	if r.Configuration.Embedding.Cohere != nil {
		cohereKey := r.Configuration.Embedding.Cohere.CohereKey.ValueString()
		destinationPineconeUpdateCohere = &shared.DestinationPineconeUpdateCohere{
			CohereKey: cohereKey,
		}
	}
	if destinationPineconeUpdateCohere != nil {
		embedding = shared.DestinationPineconeUpdateEmbedding{
			DestinationPineconeUpdateCohere: destinationPineconeUpdateCohere,
		}
	}
	var destinationPineconeUpdateFake *shared.DestinationPineconeUpdateFake
	if r.Configuration.Embedding.Fake != nil {
		destinationPineconeUpdateFake = &shared.DestinationPineconeUpdateFake{}
	}
	if destinationPineconeUpdateFake != nil {
		embedding = shared.DestinationPineconeUpdateEmbedding{
			DestinationPineconeUpdateFake: destinationPineconeUpdateFake,
		}
	}
	var destinationPineconeUpdateAzureOpenAI *shared.DestinationPineconeUpdateAzureOpenAI
	if r.Configuration.Embedding.AzureOpenAI != nil {
		apiBase := r.Configuration.Embedding.AzureOpenAI.APIBase.ValueString()
		deployment := r.Configuration.Embedding.AzureOpenAI.Deployment.ValueString()
		openaiKey1 := r.Configuration.Embedding.AzureOpenAI.OpenaiKey.ValueString()
		destinationPineconeUpdateAzureOpenAI = &shared.DestinationPineconeUpdateAzureOpenAI{
			APIBase:    apiBase,
			Deployment: deployment,
			OpenaiKey:  openaiKey1,
		}
	}
	if destinationPineconeUpdateAzureOpenAI != nil {
		embedding = shared.DestinationPineconeUpdateEmbedding{
			DestinationPineconeUpdateAzureOpenAI: destinationPineconeUpdateAzureOpenAI,
		}
	}
	var destinationPineconeUpdateOpenAICompatible *shared.DestinationPineconeUpdateOpenAICompatible
	if r.Configuration.Embedding.OpenAICompatible != nil {
		apiKey := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.APIKey.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.APIKey.IsNull() {
			*apiKey = r.Configuration.Embedding.OpenAICompatible.APIKey.ValueString()
		} else {
			apiKey = nil
		}
		baseURL := r.Configuration.Embedding.OpenAICompatible.BaseURL.ValueString()
		dimensions := r.Configuration.Embedding.OpenAICompatible.Dimensions.ValueInt64()
		modelName := new(string)
		if !r.Configuration.Embedding.OpenAICompatible.ModelName.IsUnknown() && !r.Configuration.Embedding.OpenAICompatible.ModelName.IsNull() {
			*modelName = r.Configuration.Embedding.OpenAICompatible.ModelName.ValueString()
		} else {
			modelName = nil
		}
		destinationPineconeUpdateOpenAICompatible = &shared.DestinationPineconeUpdateOpenAICompatible{
			APIKey:     apiKey,
			BaseURL:    baseURL,
			Dimensions: dimensions,
			ModelName:  modelName,
		}
	}
	if destinationPineconeUpdateOpenAICompatible != nil {
		embedding = shared.DestinationPineconeUpdateEmbedding{
			DestinationPineconeUpdateOpenAICompatible: destinationPineconeUpdateOpenAICompatible,
		}
	}
	index := r.Configuration.Indexing.Index.ValueString()
	pineconeEnvironment := r.Configuration.Indexing.PineconeEnvironment.ValueString()
	pineconeKey := r.Configuration.Indexing.PineconeKey.ValueString()
	indexing := shared.DestinationPineconeUpdateIndexing{
		Index:               index,
		PineconeEnvironment: pineconeEnvironment,
		PineconeKey:         pineconeKey,
	}
	omitRawText := new(bool)
	if !r.Configuration.OmitRawText.IsUnknown() && !r.Configuration.OmitRawText.IsNull() {
		*omitRawText = r.Configuration.OmitRawText.ValueBool()
	} else {
		omitRawText = nil
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var fieldNameMappings []shared.DestinationPineconeUpdateFieldNameMappingConfigModel = []shared.DestinationPineconeUpdateFieldNameMappingConfigModel{}
	for _, fieldNameMappingsItem := range r.Configuration.Processing.FieldNameMappings {
		fromField := fieldNameMappingsItem.FromField.ValueString()
		toField := fieldNameMappingsItem.ToField.ValueString()
		fieldNameMappings = append(fieldNameMappings, shared.DestinationPineconeUpdateFieldNameMappingConfigModel{
			FromField: fromField,
			ToField:   toField,
		})
	}
	var metadataFields []string = []string{}
	for _, metadataFieldsItem := range r.Configuration.Processing.MetadataFields {
		metadataFields = append(metadataFields, metadataFieldsItem.ValueString())
	}
	var textFields []string = []string{}
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	var textSplitter *shared.DestinationPineconeUpdateTextSplitter
	if r.Configuration.Processing.TextSplitter != nil {
		var destinationPineconeUpdateBySeparator *shared.DestinationPineconeUpdateBySeparator
		if r.Configuration.Processing.TextSplitter.BySeparator != nil {
			keepSeparator := new(bool)
			if !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsUnknown() && !r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.IsNull() {
				*keepSeparator = r.Configuration.Processing.TextSplitter.BySeparator.KeepSeparator.ValueBool()
			} else {
				keepSeparator = nil
			}
			var separators []string = []string{}
			for _, separatorsItem := range r.Configuration.Processing.TextSplitter.BySeparator.Separators {
				separators = append(separators, separatorsItem.ValueString())
			}
			destinationPineconeUpdateBySeparator = &shared.DestinationPineconeUpdateBySeparator{
				KeepSeparator: keepSeparator,
				Separators:    separators,
			}
		}
		if destinationPineconeUpdateBySeparator != nil {
			textSplitter = &shared.DestinationPineconeUpdateTextSplitter{
				DestinationPineconeUpdateBySeparator: destinationPineconeUpdateBySeparator,
			}
		}
		var destinationPineconeUpdateByMarkdownHeader *shared.DestinationPineconeUpdateByMarkdownHeader
		if r.Configuration.Processing.TextSplitter.ByMarkdownHeader != nil {
			splitLevel := new(int64)
			if !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsUnknown() && !r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.IsNull() {
				*splitLevel = r.Configuration.Processing.TextSplitter.ByMarkdownHeader.SplitLevel.ValueInt64()
			} else {
				splitLevel = nil
			}
			destinationPineconeUpdateByMarkdownHeader = &shared.DestinationPineconeUpdateByMarkdownHeader{
				SplitLevel: splitLevel,
			}
		}
		if destinationPineconeUpdateByMarkdownHeader != nil {
			textSplitter = &shared.DestinationPineconeUpdateTextSplitter{
				DestinationPineconeUpdateByMarkdownHeader: destinationPineconeUpdateByMarkdownHeader,
			}
		}
		var destinationPineconeUpdateByProgrammingLanguage *shared.DestinationPineconeUpdateByProgrammingLanguage
		if r.Configuration.Processing.TextSplitter.ByProgrammingLanguage != nil {
			language := shared.DestinationPineconeUpdateLanguage(r.Configuration.Processing.TextSplitter.ByProgrammingLanguage.Language.ValueString())
			destinationPineconeUpdateByProgrammingLanguage = &shared.DestinationPineconeUpdateByProgrammingLanguage{
				Language: language,
			}
		}
		if destinationPineconeUpdateByProgrammingLanguage != nil {
			textSplitter = &shared.DestinationPineconeUpdateTextSplitter{
				DestinationPineconeUpdateByProgrammingLanguage: destinationPineconeUpdateByProgrammingLanguage,
			}
		}
	}
	processing := shared.DestinationPineconeUpdateProcessingConfigModel{
		ChunkOverlap:      chunkOverlap,
		ChunkSize:         chunkSize,
		FieldNameMappings: fieldNameMappings,
		MetadataFields:    metadataFields,
		TextFields:        textFields,
		TextSplitter:      textSplitter,
	}
	configuration := shared.DestinationPineconeUpdate{
		Embedding:   embedding,
		Indexing:    indexing,
		OmitRawText: omitRawText,
		Processing:  processing,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPineconePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}
