// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationLangchainResourceModel) ToCreateSDKType() *shared.DestinationLangchainCreateRequest {
	destinationType := shared.DestinationLangchainLangchain(r.Configuration.DestinationType.ValueString())
	var embedding shared.DestinationLangchainEmbedding
	var destinationLangchainEmbeddingOpenAI *shared.DestinationLangchainEmbeddingOpenAI
	if r.Configuration.Embedding.DestinationLangchainEmbeddingOpenAI != nil {
		mode := new(shared.DestinationLangchainEmbeddingOpenAIMode)
		if !r.Configuration.Embedding.DestinationLangchainEmbeddingOpenAI.Mode.IsUnknown() && !r.Configuration.Embedding.DestinationLangchainEmbeddingOpenAI.Mode.IsNull() {
			*mode = shared.DestinationLangchainEmbeddingOpenAIMode(r.Configuration.Embedding.DestinationLangchainEmbeddingOpenAI.Mode.ValueString())
		} else {
			mode = nil
		}
		openaiKey := r.Configuration.Embedding.DestinationLangchainEmbeddingOpenAI.OpenaiKey.ValueString()
		destinationLangchainEmbeddingOpenAI = &shared.DestinationLangchainEmbeddingOpenAI{
			Mode:      mode,
			OpenaiKey: openaiKey,
		}
	}
	if destinationLangchainEmbeddingOpenAI != nil {
		embedding = shared.DestinationLangchainEmbedding{
			DestinationLangchainEmbeddingOpenAI: destinationLangchainEmbeddingOpenAI,
		}
	}
	var destinationLangchainEmbeddingFake *shared.DestinationLangchainEmbeddingFake
	if r.Configuration.Embedding.DestinationLangchainEmbeddingFake != nil {
		mode1 := new(shared.DestinationLangchainEmbeddingFakeMode)
		if !r.Configuration.Embedding.DestinationLangchainEmbeddingFake.Mode.IsUnknown() && !r.Configuration.Embedding.DestinationLangchainEmbeddingFake.Mode.IsNull() {
			*mode1 = shared.DestinationLangchainEmbeddingFakeMode(r.Configuration.Embedding.DestinationLangchainEmbeddingFake.Mode.ValueString())
		} else {
			mode1 = nil
		}
		destinationLangchainEmbeddingFake = &shared.DestinationLangchainEmbeddingFake{
			Mode: mode1,
		}
	}
	if destinationLangchainEmbeddingFake != nil {
		embedding = shared.DestinationLangchainEmbedding{
			DestinationLangchainEmbeddingFake: destinationLangchainEmbeddingFake,
		}
	}
	var indexing shared.DestinationLangchainIndexing
	var destinationLangchainIndexingPinecone *shared.DestinationLangchainIndexingPinecone
	if r.Configuration.Indexing.DestinationLangchainIndexingPinecone != nil {
		index := r.Configuration.Indexing.DestinationLangchainIndexingPinecone.Index.ValueString()
		mode2 := new(shared.DestinationLangchainIndexingPineconeMode)
		if !r.Configuration.Indexing.DestinationLangchainIndexingPinecone.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainIndexingPinecone.Mode.IsNull() {
			*mode2 = shared.DestinationLangchainIndexingPineconeMode(r.Configuration.Indexing.DestinationLangchainIndexingPinecone.Mode.ValueString())
		} else {
			mode2 = nil
		}
		pineconeEnvironment := r.Configuration.Indexing.DestinationLangchainIndexingPinecone.PineconeEnvironment.ValueString()
		pineconeKey := r.Configuration.Indexing.DestinationLangchainIndexingPinecone.PineconeKey.ValueString()
		destinationLangchainIndexingPinecone = &shared.DestinationLangchainIndexingPinecone{
			Index:               index,
			Mode:                mode2,
			PineconeEnvironment: pineconeEnvironment,
			PineconeKey:         pineconeKey,
		}
	}
	if destinationLangchainIndexingPinecone != nil {
		indexing = shared.DestinationLangchainIndexing{
			DestinationLangchainIndexingPinecone: destinationLangchainIndexingPinecone,
		}
	}
	var destinationLangchainIndexingDocArrayHnswSearch *shared.DestinationLangchainIndexingDocArrayHnswSearch
	if r.Configuration.Indexing.DestinationLangchainIndexingDocArrayHnswSearch != nil {
		destinationPath := r.Configuration.Indexing.DestinationLangchainIndexingDocArrayHnswSearch.DestinationPath.ValueString()
		mode3 := new(shared.DestinationLangchainIndexingDocArrayHnswSearchMode)
		if !r.Configuration.Indexing.DestinationLangchainIndexingDocArrayHnswSearch.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainIndexingDocArrayHnswSearch.Mode.IsNull() {
			*mode3 = shared.DestinationLangchainIndexingDocArrayHnswSearchMode(r.Configuration.Indexing.DestinationLangchainIndexingDocArrayHnswSearch.Mode.ValueString())
		} else {
			mode3 = nil
		}
		destinationLangchainIndexingDocArrayHnswSearch = &shared.DestinationLangchainIndexingDocArrayHnswSearch{
			DestinationPath: destinationPath,
			Mode:            mode3,
		}
	}
	if destinationLangchainIndexingDocArrayHnswSearch != nil {
		indexing = shared.DestinationLangchainIndexing{
			DestinationLangchainIndexingDocArrayHnswSearch: destinationLangchainIndexingDocArrayHnswSearch,
		}
	}
	var destinationLangchainIndexingChromaLocalPersistance *shared.DestinationLangchainIndexingChromaLocalPersistance
	if r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance != nil {
		collectionName := new(string)
		if !r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.CollectionName.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.CollectionName.IsNull() {
			*collectionName = r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.CollectionName.ValueString()
		} else {
			collectionName = nil
		}
		destinationPath1 := r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.DestinationPath.ValueString()
		mode4 := new(shared.DestinationLangchainIndexingChromaLocalPersistanceMode)
		if !r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.Mode.IsNull() {
			*mode4 = shared.DestinationLangchainIndexingChromaLocalPersistanceMode(r.Configuration.Indexing.DestinationLangchainIndexingChromaLocalPersistance.Mode.ValueString())
		} else {
			mode4 = nil
		}
		destinationLangchainIndexingChromaLocalPersistance = &shared.DestinationLangchainIndexingChromaLocalPersistance{
			CollectionName:  collectionName,
			DestinationPath: destinationPath1,
			Mode:            mode4,
		}
	}
	if destinationLangchainIndexingChromaLocalPersistance != nil {
		indexing = shared.DestinationLangchainIndexing{
			DestinationLangchainIndexingChromaLocalPersistance: destinationLangchainIndexingChromaLocalPersistance,
		}
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	processing := shared.DestinationLangchainProcessingConfigModel{
		ChunkOverlap: chunkOverlap,
		ChunkSize:    chunkSize,
		TextFields:   textFields,
	}
	configuration := shared.DestinationLangchain{
		DestinationType: destinationType,
		Embedding:       embedding,
		Indexing:        indexing,
		Processing:      processing,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationLangchainCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationLangchainResourceModel) ToGetSDKType() *shared.DestinationLangchainCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationLangchainResourceModel) ToUpdateSDKType() *shared.DestinationLangchainPutRequest {
	var embedding shared.DestinationLangchainUpdateEmbedding
	var destinationLangchainUpdateEmbeddingOpenAI *shared.DestinationLangchainUpdateEmbeddingOpenAI
	if r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingOpenAI != nil {
		mode := new(shared.DestinationLangchainUpdateEmbeddingOpenAIMode)
		if !r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingOpenAI.Mode.IsUnknown() && !r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingOpenAI.Mode.IsNull() {
			*mode = shared.DestinationLangchainUpdateEmbeddingOpenAIMode(r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingOpenAI.Mode.ValueString())
		} else {
			mode = nil
		}
		openaiKey := r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingOpenAI.OpenaiKey.ValueString()
		destinationLangchainUpdateEmbeddingOpenAI = &shared.DestinationLangchainUpdateEmbeddingOpenAI{
			Mode:      mode,
			OpenaiKey: openaiKey,
		}
	}
	if destinationLangchainUpdateEmbeddingOpenAI != nil {
		embedding = shared.DestinationLangchainUpdateEmbedding{
			DestinationLangchainUpdateEmbeddingOpenAI: destinationLangchainUpdateEmbeddingOpenAI,
		}
	}
	var destinationLangchainUpdateEmbeddingFake *shared.DestinationLangchainUpdateEmbeddingFake
	if r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingFake != nil {
		mode1 := new(shared.DestinationLangchainUpdateEmbeddingFakeMode)
		if !r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingFake.Mode.IsUnknown() && !r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingFake.Mode.IsNull() {
			*mode1 = shared.DestinationLangchainUpdateEmbeddingFakeMode(r.Configuration.Embedding.DestinationLangchainUpdateEmbeddingFake.Mode.ValueString())
		} else {
			mode1 = nil
		}
		destinationLangchainUpdateEmbeddingFake = &shared.DestinationLangchainUpdateEmbeddingFake{
			Mode: mode1,
		}
	}
	if destinationLangchainUpdateEmbeddingFake != nil {
		embedding = shared.DestinationLangchainUpdateEmbedding{
			DestinationLangchainUpdateEmbeddingFake: destinationLangchainUpdateEmbeddingFake,
		}
	}
	var indexing shared.DestinationLangchainUpdateIndexing
	var destinationLangchainUpdateIndexingPinecone *shared.DestinationLangchainUpdateIndexingPinecone
	if r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone != nil {
		index := r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.Index.ValueString()
		mode2 := new(shared.DestinationLangchainUpdateIndexingPineconeMode)
		if !r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.Mode.IsNull() {
			*mode2 = shared.DestinationLangchainUpdateIndexingPineconeMode(r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.Mode.ValueString())
		} else {
			mode2 = nil
		}
		pineconeEnvironment := r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.PineconeEnvironment.ValueString()
		pineconeKey := r.Configuration.Indexing.DestinationLangchainUpdateIndexingPinecone.PineconeKey.ValueString()
		destinationLangchainUpdateIndexingPinecone = &shared.DestinationLangchainUpdateIndexingPinecone{
			Index:               index,
			Mode:                mode2,
			PineconeEnvironment: pineconeEnvironment,
			PineconeKey:         pineconeKey,
		}
	}
	if destinationLangchainUpdateIndexingPinecone != nil {
		indexing = shared.DestinationLangchainUpdateIndexing{
			DestinationLangchainUpdateIndexingPinecone: destinationLangchainUpdateIndexingPinecone,
		}
	}
	var destinationLangchainUpdateIndexingDocArrayHnswSearch *shared.DestinationLangchainUpdateIndexingDocArrayHnswSearch
	if r.Configuration.Indexing.DestinationLangchainUpdateIndexingDocArrayHnswSearch != nil {
		destinationPath := r.Configuration.Indexing.DestinationLangchainUpdateIndexingDocArrayHnswSearch.DestinationPath.ValueString()
		mode3 := new(shared.DestinationLangchainUpdateIndexingDocArrayHnswSearchMode)
		if !r.Configuration.Indexing.DestinationLangchainUpdateIndexingDocArrayHnswSearch.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainUpdateIndexingDocArrayHnswSearch.Mode.IsNull() {
			*mode3 = shared.DestinationLangchainUpdateIndexingDocArrayHnswSearchMode(r.Configuration.Indexing.DestinationLangchainUpdateIndexingDocArrayHnswSearch.Mode.ValueString())
		} else {
			mode3 = nil
		}
		destinationLangchainUpdateIndexingDocArrayHnswSearch = &shared.DestinationLangchainUpdateIndexingDocArrayHnswSearch{
			DestinationPath: destinationPath,
			Mode:            mode3,
		}
	}
	if destinationLangchainUpdateIndexingDocArrayHnswSearch != nil {
		indexing = shared.DestinationLangchainUpdateIndexing{
			DestinationLangchainUpdateIndexingDocArrayHnswSearch: destinationLangchainUpdateIndexingDocArrayHnswSearch,
		}
	}
	var destinationLangchainUpdateIndexingChromaLocalPersistance *shared.DestinationLangchainUpdateIndexingChromaLocalPersistance
	if r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance != nil {
		collectionName := new(string)
		if !r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.CollectionName.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.CollectionName.IsNull() {
			*collectionName = r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.CollectionName.ValueString()
		} else {
			collectionName = nil
		}
		destinationPath1 := r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.DestinationPath.ValueString()
		mode4 := new(shared.DestinationLangchainUpdateIndexingChromaLocalPersistanceMode)
		if !r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.Mode.IsUnknown() && !r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.Mode.IsNull() {
			*mode4 = shared.DestinationLangchainUpdateIndexingChromaLocalPersistanceMode(r.Configuration.Indexing.DestinationLangchainUpdateIndexingChromaLocalPersistance.Mode.ValueString())
		} else {
			mode4 = nil
		}
		destinationLangchainUpdateIndexingChromaLocalPersistance = &shared.DestinationLangchainUpdateIndexingChromaLocalPersistance{
			CollectionName:  collectionName,
			DestinationPath: destinationPath1,
			Mode:            mode4,
		}
	}
	if destinationLangchainUpdateIndexingChromaLocalPersistance != nil {
		indexing = shared.DestinationLangchainUpdateIndexing{
			DestinationLangchainUpdateIndexingChromaLocalPersistance: destinationLangchainUpdateIndexingChromaLocalPersistance,
		}
	}
	chunkOverlap := new(int64)
	if !r.Configuration.Processing.ChunkOverlap.IsUnknown() && !r.Configuration.Processing.ChunkOverlap.IsNull() {
		*chunkOverlap = r.Configuration.Processing.ChunkOverlap.ValueInt64()
	} else {
		chunkOverlap = nil
	}
	chunkSize := r.Configuration.Processing.ChunkSize.ValueInt64()
	var textFields []string = nil
	for _, textFieldsItem := range r.Configuration.Processing.TextFields {
		textFields = append(textFields, textFieldsItem.ValueString())
	}
	processing := shared.DestinationLangchainUpdateProcessingConfigModel{
		ChunkOverlap: chunkOverlap,
		ChunkSize:    chunkSize,
		TextFields:   textFields,
	}
	configuration := shared.DestinationLangchainUpdate{
		Embedding:  embedding,
		Indexing:   indexing,
		Processing: processing,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationLangchainPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationLangchainResourceModel) ToDeleteSDKType() *shared.DestinationLangchainCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationLangchainResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationLangchainResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
