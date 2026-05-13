// Package types contains all types used by the scraper service.
package types

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"isbn"`
	Genre     string `json:"genre"`
	Publisher string `json:"publisher"`
	StoreLink string `json:"link"`

	Year        int    `json:"year,omitempty"`
	Language    string `json:"language,omitempty"`
	CoverType   string `json:"cover_type,omitempty"`
	PageCount   int    `json:"page_count,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUUID   string `json:"image_uuid,omitempty"`

	ImageLink string `json:"-"`
}

// TempBook holds all scraped fields as raw text for easier processing by quirks.
type TempBook struct {
	Title,
	Author,
	ISBN,
	Genre,
	Publisher,
	StoreLink,
	Year,
	Language,
	CoverType,
	PageCount,
	Description,
	ImageUUID,
	ImageLink string
}
