package spaceapi

type SpaceAPI09 struct {
  Events []Events `json:"events,omitempty"`
  Lat float64 `json:"lat,omitempty"`
  Space string `json:"space"`
  Api string `json:"api"`
  Url string `json:"url"`
  Address string `json:"address,omitempty"`
  Open bool `json:"open"`
  Cam []string `json:"cam,omitempty"`
  Lastchange float64 `json:"lastchange,omitempty"`
  Logo string `json:"logo"`
  Lon float64 `json:"lon,omitempty"`
  Status string `json:"status,omitempty"`
  Stream *Stream `json:"stream,omitempty"`
  Contact *Contact `json:"contact,omitempty"`
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

type Stream struct {
}
