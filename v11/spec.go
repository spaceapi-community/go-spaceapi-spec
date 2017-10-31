package spaceapi

type SpaceAPI011 struct {
  Status string `json:"status,omitempty"`
  Url string `json:"url"`
  Address string `json:"address,omitempty"`
  Icon *Icon `json:"icon"`
  Lat float64 `json:"lat,omitempty"`
  Events []Events `json:"events,omitempty"`
  Api string `json:"api"`
  Cam []string `json:"cam,omitempty"`
  Contact *Contact `json:"contact,omitempty"`
  Open bool `json:"open"`
  Space string `json:"space"`
  Stream *Stream `json:"stream,omitempty"`
  Lastchange float64 `json:"lastchange,omitempty"`
  Logo string `json:"logo"`
  Lon float64 `json:"lon,omitempty"`
}

type Contact struct {
  Keymaster []string `json:"keymaster,omitempty"`
  Ml string `json:"ml,omitempty"`
  Phone string `json:"phone,omitempty"`
  Sip string `json:"sip,omitempty"`
  Twitter string `json:"twitter,omitempty"`
  Email string `json:"email,omitempty"`
  Irc string `json:"irc,omitempty"`
  Jabber string `json:"jabber,omitempty"`
}

type Events struct {
  Extra string `json:"extra,omitempty"`
  Name string `json:"name"`
  T float64 `json:"t"`
  Type string `json:"type"`
}

type Icon struct {
  Closed string `json:"closed"`
  Open string `json:"open"`
}

type Stream struct {
}