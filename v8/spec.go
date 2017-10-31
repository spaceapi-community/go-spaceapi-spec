package spaceapi

type SpaceAPI08 struct {
  Lon float64 `json:"lon,omitempty"`
  Url string `json:"url"`
  Open bool `json:"open"`
  Phone string `json:"phone,omitempty"`
  Status string `json:"status,omitempty"`
  Stream *Stream `json:"stream,omitempty"`
  Address string `json:"address,omitempty"`
  Lastchange float64 `json:"lastchange,omitempty"`
  Lat float64 `json:"lat,omitempty"`
  Cam []string `json:"cam,omitempty"`
  Logo string `json:"logo"`
  Api string `json:"api"`
  Events []Events `json:"events,omitempty"`
  Space string `json:"space"`
}

type Events struct {
  Extra string `json:"extra,omitempty"`
  Name string `json:"name"`
  T float64 `json:"t"`
  Type string `json:"type"`
}

type Stream struct {
}
