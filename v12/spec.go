package spaceapi

type Icon struct {
  Closed string `json:"closed"`
  Open string `json:"open"`
}

type Sensors struct {
  Temp *Temp `json:"temp,omitempty"`
}

type Temp struct {
}

type Stream struct {
}

type SpaceAPI012 struct {
  Address string `json:"address,omitempty"`
  Api string `json:"api"`
  Cam []string `json:"cam,omitempty"`
  Lon float64 `json:"lon,omitempty"`
  Events []Events `json:"events,omitempty"`
  Icon *Icon `json:"icon"`
  Open bool `json:"open"`
  Space string `json:"space"`
  Status string `json:"status,omitempty"`
  Stream *Stream `json:"stream,omitempty"`
  Feeds []Feeds `json:"feeds,omitempty"`
  Lastchange float64 `json:"lastchange,omitempty"`
  Logo string `json:"logo"`
  Sensors []Sensors `json:"sensors,omitempty"`
  Contact *Contact `json:"contact,omitempty"`
  Lat float64 `json:"lat,omitempty"`
  Url string `json:"url"`
}

type Contact struct {
  Jabber string `json:"jabber,omitempty"`
  Keymaster []string `json:"keymaster,omitempty"`
  Ml string `json:"ml,omitempty"`
  Phone string `json:"phone,omitempty"`
  Sip string `json:"sip,omitempty"`
  Twitter string `json:"twitter,omitempty"`
  Email string `json:"email,omitempty"`
  Irc string `json:"irc,omitempty"`
}

type Events struct {
  Type string `json:"type"`
  Extra string `json:"extra,omitempty"`
  Name string `json:"name"`
  T float64 `json:"t"`
}

type Feeds struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
  Name string `json:"name"`
}
