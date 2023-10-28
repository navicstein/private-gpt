package gpt

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"navicstein/private-gpt/internal/config"
	"navicstein/private-gpt/internal/ffmepg_encoder"
	"os"
	"path/filepath"
)

// TranscribeMedia texts text transcript from a file
func TranscribeMedia(ctx context.Context, path string) (string, error) {
	var (
		cfg       = config.GetConfig()
		tmpDir    = os.TempDir()
		audioPath = fmt.Sprintf("%s/%s__.mp3", tmpDir, filepath.Base(path))
	)

	if err := os.MkdirAll(tmpDir, os.ModePerm); err != nil {
		return "", err
	}

	if err := encoder.Encode(ctx, path, audioPath); err != nil {
		return "", fmt.Errorf("failed to run ffmpeg command: %w", err)
	}

	log.Info().Msgf("starting transcription: %s", path)
	openaiCfg := openai.DefaultConfig(cfg.OpenAI.APIKey)
	openaiCfg.BaseURL = cfg.OpenAI.BaseURL

	c := openai.NewClientWithConfig(openaiCfg)

	req := openai.AudioRequest{
		Model:       openai.Whisper1,
		FilePath:    audioPath,
		Temperature: 1,
		Language:    "en",
		Format:      openai.AudioResponseFormatText,
	}

	resp, err := c.CreateTranscription(ctx, req)
	if err != nil {
		return "", fmt.Errorf("transcription error: %w", err)
	}

	_ = os.Remove(audioPath)
	return resp.Text, nil
}
