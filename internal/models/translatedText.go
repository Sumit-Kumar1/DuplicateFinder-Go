package models

const (
	LocaleEN = "en-US"
)

type Locale string

type TranslatedText struct {
	Text   string `json:"text"`
	Locale Locale `json:"locale"`
}
