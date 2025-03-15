package helper

import "time"

type Frontmatter struct {
	title     string
	tags      []string
	createdAt string
	updatedAt string
}

func NewFrontmatter(title string, tags []string) *Frontmatter {
	return &Frontmatter{
		title:     title,
		tags:      tags,
		createdAt: time.Now().Format("2006-01-02 15:04:05"),
		updatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

func (f Frontmatter) Render() string {
	tags := ""
	for _, tag := range f.tags {
		tags += ", " + tag
	}
	return "---\n" +
		"title: " + f.title + "\n" +
		"tags:" + tags + "\n" +
		"created_at: " + f.createdAt + "\n" +
		"---\n"
}
