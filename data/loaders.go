package data

import (
	"context"
	"io"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
)

type Loader struct {
}

// PDFLoader is a function that loads PDF documents from a reader.
func (l *Loader) PDFLoader(ctx context.Context, reader io.ReaderAt, size int64) ([]schema.Document, error) {
	loader := documentloaders.NewPDF(reader, size, nil) // Using 'nil' for options for simplicity
	return loader.Load(ctx)
}

// TextLoader is a function that loads text documents from a file path.
func (l *Loader) TextLoader(ctx context.Context, path io.Reader) ([]schema.Document, error) {
	loader := documentloaders.NewText(path)
	return loader.Load(ctx)
}
