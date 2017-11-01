package spaceapiStruct

// Contact Contact information about your space
type Contact struct {
	Email     string   `json:"email,omitempty"`
	Irc       string   `json:"irc,omitempty"`
	Jabber    string   `json:"jabber,omitempty"`
	Keymaster []string `json:"keymaster,omitempty"`
	Ml        string   `json:"ml,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	Sip       string   `json:"sip,omitempty"`
	Twitter   string   `json:"twitter,omitempty"`
}

// Events
type Events struct {
	Extra string  `json:"extra,omitempty"`
	Name  string  `json:"name"`
	T     float64 `json:"t"`
	Type  string  `json:"type"`
}

// Feeds
type Feeds struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

// Icon Icons that show the status graphically
type Icon struct {
	Closed string `json:"closed"`
	Open   string `json:"open"`
}

// Sensors
type Sensors struct {
	Temp *Temp `json:"temp,omitempty"`
}

// SpaceAPI012 SpaceAPI 0.12
type SpaceAPI012 struct {
	Address    string    `json:"address,omitempty"`
	Api        string    `json:"api"`
	Cam        []string  `json:"cam,omitempty"`
	Contact    *Contact  `json:"contact,omitempty"`
	Events     []Events  `json:"events,omitempty"`
	Feeds      []Feeds   `json:"feeds,omitempty"`
	Icon       *Icon     `json:"icon"`
	Lastchange float64   `json:"lastchange,omitempty"`
	Lat        float64   `json:"lat,omitempty"`
	Logo       string    `json:"logo"`
	Lon        float64   `json:"lon,omitempty"`
	Open       bool      `json:"open"`
	Sensors    []Sensors `json:"sensors,omitempty"`
	Space      string    `json:"space"`
	Status     string    `json:"status,omitempty"`
	Stream     *Stream   `json:"stream,omitempty"`
	Url        string    `json:"url"`
}

// Stream A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>
type Stream struct {
}

// Temp A mapping of measuring locations to temperature values. The values should match the basic regular expression <code>^[+-]?[0-9]+(\.[0-9]+)?[FCK]$</code>. Example: <samp>{'kitchen':'48F', 'storage':'-273.1K'}</samp>
type Temp struct {
}
