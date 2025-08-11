package data

import (
	"context"
	"fmt"
	"io"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
)

type Loader struct {
}

// PDFLoader is a function that loads PDF documents from a reader.
func (l *Loader) PDFLoader(ctx context.Context, reader io.ReaderAt, size int64, options *documentloaders.PDFOptions) ([]schema.Document, error) {
	if reader == nil {
		return nil, fmt.Errorf("reader cannot be nil")
	}
	loader := documentloaders.NewPDF(reader, size, *options)
	return loader.Load(ctx)
}

// TextLoader is a function that loads text documents from a file path.
func (l *Loader) TextLoader(ctx context.Context, reader io.Reader) ([]schema.Document, error) {
	if reader == nil {
		return nil, fmt.Errorf("reader cannot be nil")
	}
	loader := documentloaders.NewText(reader)
	return loader.Load(ctx)
}
