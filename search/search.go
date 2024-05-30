package search

import (
	"slices"
	"strings"
)

// Params represents the search parameters.
type Params struct {
	Include []string
	Exclude []string
}

// ByDescription searches emojis by their description.
func ByDescription(params Params) (result []Emoji) {
	for _, emo := range emjis {
		if shouldExclude(emo, params.Exclude) {
			continue
		}
		for _, include := range params.Include {
			include = strings.ToLower(include)
			if strings.Contains(emo.Label, include) || slices.Contains(emo.Tags, include) {
				result = append(result, emo)
			}
		}
	}
	return
}

// ByTags search for emojis with specified tags
func ByTags(tags ...string) (result []Emoji) {
	for _, emo := range emjis {
		for _, tag := range tags {
			if slices.Contains(emo.Tags, tag) {
				result = append(result, emo)
			}
		}
	}

	return
}

// the specified emoji icon
func EmojiLike(emoji string) (result []Emoji) {
	// find the emoji
	var found Emoji
	for _, emo := range emjis {
		if emo.Emoji == emoji {
			found = emo
			break
		}
	}

	if found.Emoji == "" {
		return
	}

	// find emoji with similar traits
	var includes []string
	includes = append(includes, found.Tags...)
	includes = append(includes, found.Label)

	return ByDescription(Params{Include: includes})
}

// shouldExclude checks if an emoji should be excluded based on the exclude criteria.
func shouldExclude(emo Emoji, excludes []string) bool {
	for _, exclude := range excludes {
		exclude = strings.ToLower(exclude)
		if strings.Contains(emo.Label, exclude) || slices.Contains(emo.Tags, exclude) {
			return true
		}
	}
	return false
}
