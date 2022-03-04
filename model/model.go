package model

type Country struct {
	Name           string                   `bson:"name"`
	Iso3           string                   `bson:"iso3"`
	Iso2           string                   `bson:"iso2"`
	NumberCode     string                   `bson:"number_code"`
	PhoneCode      string                   `bson:"phone_code"`
	Capital        string                   `bson:"capital"`
	Currency       string                   `bson:"currency"`
	CurrencyName   string                   `bson:"currency_name"`
	CurrencySymbol string                   `bson:"currency_symbol"`
	Tld            string                   `bson:"tld"`
	Native         string                   `bson:"native"`
	Region         string                   `bson:"region"`
	Subregion      string                   `bson:"subregion"`
	Timezones      []map[string]interface{} `bson:"timezones"`
	Latitude       float64                  `bson:"latitude"`
	Longitude      float64                  `bson:"longitude"`
	Emoji          string                   `bson:"emoji"`
	EmojiU         string                   `bson:"emojiU"`
}

type State struct {
	Name        string  `bson:"name"`
	CountryCode string  `bson:"country_code"`
	StateCode   string  `bson:"state_code"`
	Type        string  `bson:"type"`
	Latitude    float64 `bson:"latitude"`
	Longitude   float64 `bson:"longitude"`
}

type City struct {
	Name        string  `bson:"name"`
	CountryCode string  `bson:"country_code"`
	StateCode   string  `bson:"state_code"`
	Latitude    float64 `bson:"latitude"`
	Longitude   float64 `bson:"longitude"`
}
