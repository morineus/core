package model

type Category struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type Series struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Tag struct {
	ID   int    `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type PostListItem struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Slug      string    `json:"slug"`
	Locale    string    `json:"locale"`
	Type      string    `json:"type"`
	Series    *Series   `json:"series,omitempty"`
	Category  *Category `json:"category,omitempty"`
	Tags      []*Tag    `json:"tags,omitempty"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
}

type PostDetail struct {
	PostListItem
	Content string `json:"content"`
}
