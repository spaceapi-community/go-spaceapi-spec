package spaceapi

type Contact struct {
  Twitter string `json:"twitter,omitempty"`
  Email string `json:"email,omitempty"`
  Foursquare string `json:"foursquare,omitempty"`
  IssueMail string `json:"issue_mail,omitempty"`
  Keymasters []Keymasters `json:"keymasters,omitempty"`
  Ml string `json:"ml,omitempty"`
  Phone string `json:"phone,omitempty"`
  Sip string `json:"sip,omitempty"`
  Facebook string `json:"facebook,omitempty"`
  Google *Google `json:"google,omitempty"`
  Identica string `json:"identica,omitempty"`
  Irc string `json:"irc,omitempty"`
  Xmpp string `json:"xmpp,omitempty"`
}

type Feeds struct {
  Blog *Blog `json:"blog,omitempty"`
  Calendar *Calendar `json:"calendar,omitempty"`
  Flickr *Flickr `json:"flickr,omitempty"`
  Wiki *Wiki `json:"wiki,omitempty"`
}

type Location struct {
  Timezone string `json:"timezone,omitempty"`
  Address string `json:"address,omitempty"`
  Lat float64 `json:"lat"`
  Lon float64 `json:"lon"`
}

type MembershipPlans struct {
  Description string `json:"description,omitempty"`
  Name string `json:"name"`
  RenewalInterval string `json:"renewal_interval"`
  Value float64 `json:"value"`
  Currency string `json:"currency"`
}

type Sensors struct {
  AccountBalance []AccountBalance `json:"account_balance,omitempty"`
  DoorLocked []DoorLocked `json:"door_locked,omitempty"`
  PowerConsumption []PowerConsumption `json:"power_consumption,omitempty"`
  Radiation *Radiation `json:"radiation,omitempty"`
  TotalMemberCount []TotalMemberCount `json:"total_member_count,omitempty"`
  Wind []Wind `json:"wind,omitempty"`
  Barometer []Barometer `json:"barometer,omitempty"`
  BeverageSupply []BeverageSupply `json:"beverage_supply,omitempty"`
  Humidity []Humidity `json:"humidity,omitempty"`
  NetworkConnections []NetworkConnections `json:"network_connections,omitempty"`
  PeopleNowPresent []PeopleNowPresent `json:"people_now_present,omitempty"`
  Temperature []Temperature `json:"temperature,omitempty"`
}

type BeverageSupply struct {
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
}

type Machines struct {
  Mac string `json:"mac"`
  Name string `json:"name,omitempty"`
}

type SpaceAPI014 struct {
  Feeds *Feeds `json:"feeds,omitempty"`
  Spacefed *Spacefed `json:"spacefed,omitempty"`
  State *State `json:"state"`
  Api string `json:"api"`
  Cache *Cache `json:"cache,omitempty"`
  Cam []string `json:"cam,omitempty"`
  Contact *Contact `json:"contact"`
  Events []Events `json:"events,omitempty"`
  Location *Location `json:"location"`
  Logo string `json:"logo"`
  Projects []string `json:"projects,omitempty"`
  Url string `json:"url"`
  MembershipPlans []MembershipPlans `json:"membership_plans,omitempty"`
  Sensors *Sensors `json:"sensors,omitempty"`
  IssueReportChannels []string `json:"issue_report_channels"`
  RadioShow []RadioShow `json:"radio_show,omitempty"`
  Space string `json:"space"`
  Stream *Stream `json:"stream,omitempty"`
}

type Wind struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Properties *Properties `json:"properties"`
}

type Speed struct {
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type TotalMemberCount struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Value float64 `json:"value"`
}

type Events struct {
  Extra string `json:"extra,omitempty"`
  Name string `json:"name"`
  Timestamp float64 `json:"timestamp"`
  Type string `json:"type"`
}

type AccountBalance struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type NetworkConnections struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Machines []Machines `json:"machines,omitempty"`
  Name string `json:"name,omitempty"`
  Type string `json:"type,omitempty"`
  Value float64 `json:"value"`
}

type Properties struct {
  Direction *Direction `json:"direction"`
  Elevation *Elevation `json:"elevation"`
  Gust *Gust `json:"gust"`
  Speed *Speed `json:"speed"`
}

type Elevation struct {
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

type Google struct {
  Plus string `json:"plus,omitempty"`
}

type Beta struct {
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type Direction struct {
  Value float64 `json:"value"`
  Unit string `json:"unit"`
}

type Spacefed struct {
  Spacenet bool `json:"spacenet"`
  Spacephone bool `json:"spacephone"`
  Spacesaml bool `json:"spacesaml"`
}

type Wiki struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type PowerConsumption struct {
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
}

type Radiation struct {
  Alpha []Alpha `json:"alpha,omitempty"`
  Beta []Beta `json:"beta,omitempty"`
  BetaGamma []BetaGamma `json:"beta_gamma,omitempty"`
  Gamma []Gamma `json:"gamma,omitempty"`
}

type Stream struct {
  Mjpeg string `json:"mjpeg,omitempty"`
  Ustream string `json:"ustream,omitempty"`
  M4 string `json:"m4,omitempty"`
}

type PeopleNowPresent struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Names []string `json:"names,omitempty"`
  Value float64 `json:"value"`
}

type Barometer struct {
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
}

type Alpha struct {
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
}

type Icon struct {
  Closed string `json:"closed"`
  Open string `json:"open"`
}

type Keymasters struct {
  Phone string `json:"phone,omitempty"`
  Twitter string `json:"twitter,omitempty"`
  Xmpp string `json:"xmpp,omitempty"`
  Email string `json:"email,omitempty"`
  IrcNick string `json:"irc_nick,omitempty"`
  Name string `json:"name,omitempty"`
}

type Gust struct {
  Value float64 `json:"value"`
  Unit string `json:"unit"`
}

type Flickr struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type RadioShow struct {
  Type string `json:"type"`
  Url string `json:"url"`
  End string `json:"end"`
  Name string `json:"name"`
  Start string `json:"start"`
}

type DoorLocked struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Value bool `json:"value"`
}

type Blog struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type Calendar struct {
  Type string `json:"type,omitempty"`
  Url string `json:"url"`
}

type Humidity struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type BetaGamma struct {
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
}

type Gamma struct {
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
  ConversionFactor float64 `json:"conversion_factor,omitempty"`
  DeadTime float64 `json:"dead_time,omitempty"`
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
}

type Temperature struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location"`
  Name string `json:"name,omitempty"`
  Unit string `json:"unit"`
  Value float64 `json:"value"`
}

type Cache struct {
  Schedule string `json:"schedule"`
}
