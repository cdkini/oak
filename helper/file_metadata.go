package helper

import (
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

// FileMetadata represents important metadata for a note file.
// It is rendered and parsed as frontmatter.
type FileMetadata struct {
	Title     string `yaml:"title"`
	Tags      string `yaml:"tags"`
	CreatedAt string `yaml:"created_at"`
	UpdatedAt *string
}

var frontMatterRegex = regexp.MustCompile(`(?s)^---\n(.*?)\n---\n(.*)`)

func NewFileMetadata(title string, tags []string) *FileMetadata {
	return &FileMetadata{
		Title:     title,
		Tags:      strings.Join(tags, ", "),
		CreatedAt: time.Now().Format("2006-01-02 15:04"),
		UpdatedAt: nil, // Determined dynamically when parsing
	}
}

func ParseFileMetadata(path string) (*FileMetadata, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	metadata := &FileMetadata{}
	_, err = frontmatter.MustParse(f, metadata)

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	updatedAt := stat.ModTime().Format("2006-01-02 15:04")
	metadata.UpdatedAt = &updatedAt

	return metadata, err
}

func (f *FileMetadata) Render() string {
	return "---" + "\n" + "title: " + f.Title + "\n" + "tags: " + f.Tags + "\n" + "created_at: " + f.CreatedAt + "\n" + "---" + "\n\n\n"
}
