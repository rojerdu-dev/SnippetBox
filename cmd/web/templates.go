package main

import "github.com/rojerdu-dev/SnippetBox/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
