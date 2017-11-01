package spaceapiStruct

type Events struct {
	Extra string  `json:"extra,omitempty"`
	Name  string  `json:"name"`
	T     float64 `json:"t"`
	Type  string  `json:"type"`
}

type SpaceAPI08 struct {
	Address    string   `json:"address,omitempty"`
	Api        string   `json:"api"`
	Cam        []string `json:"cam,omitempty"`
	Events     []Events `json:"events,omitempty"`
	Lastchange float64  `json:"lastchange,omitempty"`
	Lat        float64  `json:"lat,omitempty"`
	Logo       string   `json:"logo"`
	Lon        float64  `json:"lon,omitempty"`
	Open       bool     `json:"open"`
	Phone      string   `json:"phone,omitempty"`
	Space      string   `json:"space"`
	Status     string   `json:"status,omitempty"`
	Stream     *Stream  `json:"stream,omitempty"`
	Url        string   `json:"url"`
}

type Stream struct {
}
