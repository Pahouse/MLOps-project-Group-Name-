//PLACEHOLDER FUNCTION UNTIL WE SORT THE NOTEBOOK FILES INTO PYTHON ONES


package main

import (
	"dagger/pipeline/internal/dagger"
)

type Pipeline struct{}

// Env builds the Python environment with the project's dependencies installed
func (m *Pipeline) Env(
	// +defaultPath="/"
	// +ignore=[".git", ".venv", "**/__pycache__", "dagger"]
	source *dagger.Directory,
) *dagger.Container {
	return dag.Container().
		From("python:3.10-slim").
		WithMountedCache("/root/.cache/pip", dag.CacheVolume("pip-cache")). // faster re-runs
		WithDirectory("/app", source).
		WithWorkdir("/app").
		WithExec([]string{"pip", "install", "-r", "requirements.txt"})
}

// Train runs the training script and returns the trained model file
func (m *Pipeline) Train(
	// +defaultPath="/"
	// +ignore=[".git", ".venv", "**/__pycache__", "dagger"]
	source *dagger.Directory,
) *dagger.File {
	return m.Env(source).
		WithExec([]string{"python", "./models/placeholder_train.py"}).
		File("/app/models/model.pkl")
}