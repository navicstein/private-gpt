package vectorstore

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/vectorstores"
	"navicstein/private-gpt/internal/config"
	customWeaviate "navicstein/private-gpt/internal/gpt/weaviate"
)

// GetDefaultCollectionName TODO: make this configurable
func GetDefaultCollectionName() string {
	return "DefaultCollection"
}

// GetVectorStore returns a vector store based on name
func GetVectorStore(ctx context.Context, uniqueCollectionName string) (vectorstores.VectorStore, error) {
	var (
		vectorStore vectorstores.VectorStore
		err         error
		cfg         = config.GetConfig()
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create embedder: %w", err)
	}

	log.Debug().Str("collection", uniqueCollectionName).Str("vector_store", cfg.VectorStore).Msgf("getting vector store: %s", cfg.VectorStore)

	switch {
	case cfg.VectorStore == "weaviate":
		vectorStore, err = customWeaviate.New(
			customWeaviate.WithScheme(cfg.Weaviate.Scheme),
			customWeaviate.WithHost(cfg.Weaviate.Host),
			customWeaviate.WithIndexName(uniqueCollectionName),
			// weaviate has a flat namespace, so we need to add a namespace key
			// as the collection ID to each document owned by the collection
			customWeaviate.WithNameSpaceKey("collectionId"),

			// use unique collection name as namespace
			customWeaviate.WithNameSpace(uniqueCollectionName),
			customWeaviate.WithQueryAttrs([]string{"documentId"}),
		)
		//vectorStore, err = weaviate.New(
		//	weaviate.WithScheme(config.Config("WEAVIATE_SCHEME")),
		//	weaviate.WithHost(config.Config("WEAVIATE_HOST")),
		//	weaviate.WithIndexName(uniqueCollectionName),
		//	weaviate.WithNameSpaceKey("collectionId"),
		//	weaviate.WithEmbedder(embedder),
		//	weaviate.WithNameSpace(uniqueCollectionName),
		//	weaviate.WithQueryAttrs([]string{"documentId", "collectionId", "userId"}),
		//)
	default:
		return nil, errors.New("vector store not supported at the moment")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to create vector store: %w", err)
	}
	return vectorStore, nil
}
