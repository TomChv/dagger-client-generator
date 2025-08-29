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
) (*dagger.Directory, error) {
  return dag.Directory().WithNewFile("hello.txt", "hello world"), nil
}
