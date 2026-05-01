package main

// Release represents a release from the Discogs collection
type Release struct {
	ID       int      `json:"id"`
	Title    string   `json:"title"`
	Artists  []Artist `json:"artists"`
	Year     int      `json:"year"`
	URI      string   `json:"uri"`
	Resource string   `json:"resource_url"`
}

// Artist represents an artist in a release
type Artist struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// CollectionItem represents an item in the Discogs collection
type CollectionItem struct {
	ID          int     `json:"id"`
	InstanceID  int     `json:"instance_id"`
	DateAdded   string  `json:"date_added"`
	BasicInfo   Release `json:"basic_information"`
	RatingAvg   float64 `json:"rating"`
	URI         string  `json:"uri"`
	ResourceURL string  `json:"resource_url"`
}

// CollectionResponse represents the paginated response from Discogs collection API
type CollectionResponse struct {
	Pagination struct {
		Page    int `json:"page"`
		Pages   int `json:"pages"`
		PerPage int `json:"per_page"`
		Total   int `json:"total"`
		URLs    struct {
			Last string `json:"last"`
			Next string `json:"next"`
		} `json:"urls"`
	} `json:"pagination"`
	Releases []CollectionItem `json:"releases"`
}
