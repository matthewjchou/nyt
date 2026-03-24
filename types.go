package nyt

type Article struct {
	Section           string       `json:"section"`
	Subsection        string       `json:"subsection"`
	Title             string       `json:"title"`
	Abstract          string       `json:"abstract"`
	URL               string       `json:"url"`
	URI               string       `json:"uri"`
	Byline            string       `json:"byline"`
	ItemType          string       `json:"item_type"`
	UpdatedDate       string       `json:"updated_date"`
	CreatedDate       string       `json:"created_date"`
	PublishedDate     string       `json:"published_date"`
	MaterialTypeFacet string       `json:"material_type_facet"`
	Kicker            string       `json:"kicker"`
	DesFacet          []string     `json:"des_facet"`
	OrgFacet          []string     `json:"org_facet"`
	PerFacet          []string     `json:"per_facet"`
	GeoFacet          []string     `json:"geo_facet"`
	Multimedia        []Multimedia `json:"multimedia"`
	ShortURL          string       `json:"short_url"`
}

type Multimedia struct {
	URL       string `json:"url"`
	Format    string `json:"format"`
	Height    int    `json:"height"`
	Width     int    `json:"width"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Caption   string `json:"caption"`
	Copyright string `json:"copyright"`
}

