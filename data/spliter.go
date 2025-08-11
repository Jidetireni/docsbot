package data

import (
	"context"
	"os"

	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

type DocumentSplitterConfig struct {
	PdfPaths     []string
	TextPaths    []string
	ChunckSize   int
	ChunkOverlap int
}

func (l *Loader) SplitDocuments(ctx context.Context, config *DocumentSplitterConfig) ([]schema.Document, error) {
	var alldocs []schema.Document
	for _, pdfPath := range config.PdfPaths {
		reader, size, err := loadPDFFile(pdfPath)
		if err != nil {
			return nil, err
		}

		pdfDocs, err := l.PDFLoader(ctx, reader, size)
		if err != nil {
			return nil, err
		}

		alldocs = append(alldocs, pdfDocs...)
	}

	for _, textPath := range config.TextPaths {
		reader, err := os.Open(textPath)
		if err != nil {
			return nil, err
		}
		textDocs, err := l.TextLoader(ctx, reader)
		if err != nil {
			return nil, err
		}

		alldocs = append(alldocs, textDocs...)
	}

	splitter := textsplitter.NewRecursiveCharacter(
		textsplitter.WithChunkSize(config.ChunckSize),
		textsplitter.WithChunkOverlap(config.ChunkOverlap),
	)

	chunks, err := textsplitter.SplitDocuments(splitter, alldocs)
	if err != nil {
		return nil, err
	}

	return chunks, nil
}
