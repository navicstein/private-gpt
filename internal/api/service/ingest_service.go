package service

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/schema"
	"navicstein/private-gpt/internal/database/model"
	"navicstein/private-gpt/internal/gpt"
	"path/filepath"
	"time"
)

type IngestDataParam struct {
	CollectionID  string         `json:"collectionId"`
	Meta          map[string]any `json:"meta"`
	SavedFilePath string         // path to the file saved in fs
}

type IngestDataResult struct {
	IngestedAt time.Time      `json:"ingestedAt"`
	Document   model.Document `json:"document"`
}

func (s *Service) IngestFiles(ctx context.Context, params IngestDataParam) (*IngestDataResult, error) {
	var (
		err error
	)

	// upload to s3 for global storage for transcription
	s3Path := fmt.Sprintf("uploads/%s", filepath.Base(params.SavedFilePath))

	//TODO: create a tunnel for transcription

	// extract relevant params from file
	docParams, err := gpt.ProcessDocumentFile(ctx, params.SavedFilePath)
	if err != nil {
		return nil, err
	}

	// saved s3Path
	docParams.Src = s3Path
	document, err := s.documentRepo.CreateDocument(*docParams)
	if err != nil {
		return nil, fmt.Errorf("failed to create document for collection: %w", err)
	}

	// split a document into chunks
	docChunks, err := gpt.SplitDocument(document)
	if err != nil {
		return nil, fmt.Errorf("failed to split document: %w", err)
	}

	// save chunks to vector store here
	// get only the first two & last two chunks of the document
	var extractedChunks []schema.Document

	// Get the length of the slice
	length := len(docChunks)

	// Calculate 25% of the length
	twentyFivePercentLength := length * 25 / 100

	if length > 50 {
		log.Warn().Msgf("the slice is greater than 50 (%v), truncating ..", length)

		if length > 0 && twentyFivePercentLength > 0 {
			if twentyFivePercentLength > length/2 {
				twentyFivePercentLength = length / 2
			}
			// extracting 25% from the beginning and 25% from the ending
			extractedChunks = append(docChunks[:twentyFivePercentLength], docChunks[length-twentyFivePercentLength:]...)
		} else {
			extractedChunks = append(extractedChunks, docChunks...)
			log.Warn().Err(err).Msg("the slice is empty or 25 percent of the slice length is zero")
		}
	} else {
		extractedChunks = append(extractedChunks, docChunks...)
		log.Warn().Err(err).Msg("using the whole document chunk as its len < 50 docs")
	}

	log.Info().Msgf("using extracted chunks: %v", len(extractedChunks))
	if err := gpt.EmbedDocumentChunks(ctx, extractedChunks); err != nil {
		if err := s.documentRepo.DeleteDocument(document); err != nil {
			log.Error().Err(err).Msgf("failed to delete document after embedding failed: %+v", document)
		}
		return nil, fmt.Errorf("failed to embed document chunks: %w", err)
	}

	result := IngestDataResult{
		IngestedAt: time.Now(),
		Document:   *document,
	}

	log.Info().Msgf("successfully ingested files: %+v", result.Document.ID)
	return &result, nil
}
