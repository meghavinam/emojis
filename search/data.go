package search

// Emoji represents a structure for storing information about emojis.
type Emoji struct {
	Emoji   string   // Emoji representation
	Label   string   // Label describing the emoji
	Tags    []string // Tags associated with the emoji
	Emoicon string   // Emoicon representation (if any)
}

// emjis is a slice containing Emoji structs, initialized with some sample data.
var emjis = []Emoji{
	{Emoji: "smile", Label: "Smile emoji", Tags: []string{"face", "grin"}, Emoicon: ""},
	{Emoji: "cry", Label: "Cry emoji", Tags: []string{"face", "grin", "eye"}, Emoicon: ""},
	{Emoji: "wonder", Label: "Wonder emoji", Tags: []string{"face", "grin", "mouth"}, Emoicon: ""},
}

