package spaceapi

type Flickr struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type Barometer struct {
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
}

type Humidity struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type Wind struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Properties *Properties `json:"properties"`
}

type Gust struct {
  Value float64 `json:"value"`
  Unit string `json:"unit"`
}

type Cache struct {
  Schedule string `json:"schedule"`
}

type Contact struct {
  Phone string `json:"phone,omitempty"`
  Sip string `json:"sip,omitempty"`
  Email string `json:"email,omitempty"`
  Foursquare string `json:"foursquare,omitempty"`
  Google *Google `json:"google,omitempty"`
  Identica string `json:"identica,omitempty"`
  Keymasters []Keymasters `json:"keymasters,omitempty"`
  Twitter string `json:"twitter,omitempty"`
  Facebook string `json:"facebook,omitempty"`
  Irc string `json:"irc,omitempty"`
  IssueMail string `json:"issue_mail,omitempty"`
  Jabber string `json:"jabber,omitempty"`
  Ml string `json:"ml,omitempty"`
}

type Google struct {
  Plus string `json:"plus,omitempty"`
}

type Blog struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type BeverageSupply struct {
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
}

type DoorLocked struct {
  Name string `json:"name,omitempty"`
  Value bool `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
}

type Radiation struct {
  Gamma []Gamma `json:"gamma,omitempty"`
  Alpha []Alpha `json:"alpha,omitempty"`
  Beta []Beta `json:"beta,omitempty"`
  BetaGamma []BetaGamma `json:"beta_gamma,omitempty"`
}

type SpaceAPI013 struct {
  Api string `json:"api"`
  Cache *Cache `json:"cache,omitempty"`
  Cam []string `json:"cam,omitempty"`
  Space string `json:"space"`
  Stream *Stream `json:"stream,omitempty"`
  Events []Events `json:"events,omitempty"`
  Projects []string `json:"projects,omitempty"`
  Sensors *Sensors `json:"sensors,omitempty"`
  Spacefed *Spacefed `json:"spacefed,omitempty"`
  State *State `json:"state"`
  Feeds *Feeds `json:"feeds,omitempty"`
  IssueReportChannels []string `json:"issue_report_channels"`
  Logo string `json:"logo"`
  RadioShow []RadioShow `json:"radio_show,omitempty"`
  Url string `json:"url"`
  Contact *Contact `json:"contact"`
  Location *Location `json:"location"`
}

type RadioShow struct {
  Start string `json:"start"`
  Type string `json:"type"`
  Url string `json:"url"`
  End string `json:"end"`
  Name string `json:"name"`
}

type AccountBalance struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
}

type BetaGamma struct {
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
}

type Properties struct {
  Direction *Direction `json:"direction"`
  Elevation *Elevation `json:"elevation"`
  Gust *Gust `json:"gust"`
  Speed *Speed `json:"speed"`
}

type Calendar struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type Wiki struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type Location struct {
  Lat float64 `json:"lat"`
  Lon float64 `json:"lon"`
  Address string `json:"address,omitempty"`
}

type Spacefed struct {
  Spacesaml bool `json:"spacesaml"`
  Spacenet bool `json:"spacenet"`
  Spacephone bool `json:"spacephone"`
}

type Beta struct {
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
}

type Temperature struct {
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
}

type Direction struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type Feeds struct {
  Blog *Blog `json:"blog,omitempty"`
  Calendar *Calendar `json:"calendar,omitempty"`
  Flickr *Flickr `json:"flickr,omitempty"`
  Wiki *Wiki `json:"wiki,omitempty"`
}

type NetworkConnections struct {
  Machines []Machines `json:"machines,omitempty"`
  Name string `json:"name,omitempty"`
  Type string `json:"type,omitempty"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
}

type PeopleNowPresent struct {
  Names []string `json:"names,omitempty"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
}

type Machines struct {
  Mac string `json:"mac"`
  Name string `json:"name,omitempty"`
}

type PowerConsumption struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
}

type Alpha struct {
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
}

type Gamma struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
}

type Elevation struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type Keymasters struct {
  Email string `json:"email,omitempty"`
  IrcNick string `json:"irc_nick,omitempty"`
  Name string `json:"name,omitempty"`
  Phone string `json:"phone,omitempty"`
  Twitter string `json:"twitter,omitempty"`
}

type Events struct {
  Extra string `json:"extra,omitempty"`
  Name string `json:"name"`
  Timestamp float64 `json:"timestamp"`
  Type string `json:"type"`
}

type Sensors struct {
  AccountBalance []AccountBalance `json:"account_balance,omitempty"`
  BeverageSupply []BeverageSupply `json:"beverage_supply,omitempty"`
  DoorLocked []DoorLocked `json:"door_locked,omitempty"`
  Humidity []Humidity `json:"humidity,omitempty"`
  PowerConsumption []PowerConsumption `json:"power_consumption,omitempty"`
  Temperature []Temperature `json:"temperature,omitempty"`
  TotalMemberCount []TotalMemberCount `json:"total_member_count,omitempty"`
  Barometer []Barometer `json:"barometer,omitempty"`
  NetworkConnections []NetworkConnections `json:"network_connections,omitempty"`
  PeopleNowPresent []PeopleNowPresent `json:"people_now_present,omitempty"`
  Radiation *Radiation `json:"radiation,omitempty"`
  Wind []Wind `json:"wind,omitempty"`
}

type Stream struct {
  M4 string `json:"m4,omitempty"`
  Mjpeg string `json:"mjpeg,omitempty"`
  Ustream string `json:"ustream,omitempty"`
}

type Speed struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type State struct {
  TriggerPerson string `json:"trigger_person,omitempty"`
  Icon *Icon `json:"icon,omitempty"`
  Lastchange float64 `json:"lastchange,omitempty"`
  Message string `json:"message,omitempty"`
  Open bool `json:"open"`
}

type Icon struct {
  Closed string `json:"closed"`
  Open string `json:"open"`
}

type TotalMemberCount struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Value float64 `json:"value"`
}
