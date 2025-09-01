package main

import (
	"context"
	"dagger/textgen/internal/dagger"
)

type Textgen struct{}

func (t *Textgen) GenerateClient(
	ctx context.Context,
	modSource *dagger.ModuleSource,
	introspectionJSON *dagger.File,
	outputDir string,
) (*dagger.Directory, error) {
	return dag.Directory().WithDirectory(outputDir, dag.Directory().WithNewFile("hello.txt", "hello world")), nil
}
