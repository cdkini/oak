package helper

import (
	"strings"
	"time"
)

// Frontmatter represents the YAML frontmatter of a note.
// It encapsulates important metadata.
type Frontmatter struct {
	title     string
	tags      []string
	createdAt string
	updatedAt string
}

func NewFrontmatter(title string, tags []string) *Frontmatter {
	timestamp := time.Now().Format("2006-01-02 15:04")
	return &Frontmatter{
		title:     title,
		tags:      tags,
		createdAt: timestamp,
		updatedAt: timestamp,
	}
}

func (f *Frontmatter) Render() string {
	tags := strings.Join(f.tags, ", ")
	return "---" + "\n" + "title: " + f.title + "\n" + "tags: " + tags + "\n" + "created_at: " + f.createdAt + "\n" + "---" + "\n\n\n"
}
