package fileprocessing

import (
	"bytes"
	"code.sajari.com/docconv"
	"context"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gen2brain/go-fitz"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"strings"
)

// GetTextFromFile extracts text from a file
func GetTextFromFile(ctx context.Context, f io.Reader) (string, error) {
	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	mime := mimetype.Detect(content)
	contentType := mime.String()

	var text string

	switch contentType {
	case "application/msword": // .doc
		text, _, err = docconv.ConvertDoc(bytes.NewReader(content))
		if err != nil {
			return "", fmt.Errorf("error converting .doc file")
		}
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document": // .docx
		text, _, err = docconv.ConvertDocx(bytes.NewReader(content))
		if err != nil {
			return "", fmt.Errorf("error converting .docx file: %v", err)
		}
	case "application/zip": // .pages
		text, _, err = docconv.ConvertPages(bytes.NewReader(content))
		if err != nil {
			return "", fmt.Errorf("error converting .pages file: %v", err)
		}
		text = strings.TrimSpace(text)
	case "application/epub+zip": // .epub
		fitzDoc, err := fitz.NewFromReader(bytes.NewReader(content))
		if err != nil {
			return "", fmt.Errorf("error reading .epub file: %v", err)
		}
		defer fitzDoc.Close()
		for i := 0; i < fitzDoc.NumPage(); i++ {
			pageText, err := fitzDoc.Text(i)
			if err != nil {
				return "", fmt.Errorf("error getting text from page %d: %v", i, err)
			}
			// Preprocess the text by replacing newline characters with spaces
			pageText = strings.ReplaceAll(pageText, "\n", " ")
			text += pageText
		}
	default: // Assume plain text
		detector := chardet.NewTextDetector()
		result, err := detector.DetectBest(content)
		if err != nil {
			return "", fmt.Errorf("error detecting encoding: %v", err)
		}

		if strings.ToLower(result.Charset) == "utf-8" {
			text = string(content)
		} else {
			var enc encoding.Encoding
			switch strings.ToLower(result.Charset) {
			case "iso-8859-1":
				enc = charmap.ISO8859_1
			case "windows-1252":
				enc = charmap.Windows1252
			// Add more encodings here as needed
			default:
				return "", fmt.Errorf("unsupported encoding: %s", result.Charset)
			}

			text, _, err = transform.String(enc.NewDecoder(), string(content))
			if err != nil {
				return "", fmt.Errorf("error decoding content: %v", err)
			}
		}
	}

	return text, nil
}

// ExtractTextFromPDF extract a human-readable text from a given pdf with support for spaces/whitespace.
func ExtractTextFromPDF(ctx context.Context, f io.Reader) (string, error) {
	// Convert the uploaded file to a human-readable text
	bodyResult, _, err := docconv.ConvertPDF(f)
	if err != nil {
		return "", err
	}

	// Remove extra space and newlines
	text := strings.TrimSpace(bodyResult)
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", " ")
	return text, nil
}
