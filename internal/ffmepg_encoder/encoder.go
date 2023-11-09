package encoder

import (
	"context"
	"github.com/rs/zerolog/log"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// Encode encodes a video
func Encode(ctx context.Context, input, output string) error {
	log.Debug().Msgf("encoding from %s to %s", input, output)

	cmd := ffmpeg.Input(input).
		Output(output).
		OverWriteOutput()

	if err := cmd.ErrorToStdOut().Run(); err != nil {
		return err
	}

	return nil
}
