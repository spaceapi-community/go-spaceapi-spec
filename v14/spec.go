package spaceapi_struct

type AccountBalance struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type Alpha struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

type Barometer struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type Beta struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

type BetaGamma struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

type BeverageSupply struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type Blog struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

type Cache struct {
	Schedule string `json:"schedule"`
}

type Calendar struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

type Contact struct {
	Email      string       `json:"email,omitempty"`
	Facebook   string       `json:"facebook,omitempty"`
	Foursquare string       `json:"foursquare,omitempty"`
	Google     *Google      `json:"google,omitempty"`
	Identica   string       `json:"identica,omitempty"`
	Irc        string       `json:"irc,omitempty"`
	IssueMail  string       `json:"issue_mail,omitempty"`
	Keymasters []Keymasters `json:"keymasters,omitempty"`
	Ml         string       `json:"ml,omitempty"`
	Phone      string       `json:"phone,omitempty"`
	Sip        string       `json:"sip,omitempty"`
	Twitter    string       `json:"twitter,omitempty"`
	Xmpp       string       `json:"xmpp,omitempty"`
}

type Direction struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type DoorLocked struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	Name        string `json:"name,omitempty"`
	Value       bool   `json:"value"`
}

type Elevation struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type Events struct {
	Extra     string  `json:"extra,omitempty"`
	Name      string  `json:"name"`
	Timestamp float64 `json:"timestamp"`
	Type      string  `json:"type"`
}

type Feeds struct {
	Blog     *Blog     `json:"blog,omitempty"`
	Calendar *Calendar `json:"calendar,omitempty"`
	Flickr   *Flickr   `json:"flickr,omitempty"`
	Wiki     *Wiki     `json:"wiki,omitempty"`
}

type Flickr struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

type Gamma struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

type Google struct {
	Plus string `json:"plus,omitempty"`
}

type Gust struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type Humidity struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type Icon struct {
	Closed string `json:"closed"`
	Open   string `json:"open"`
}

type Keymasters struct {
	Email   string `json:"email,omitempty"`
	IrcNick string `json:"irc_nick,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Twitter string `json:"twitter,omitempty"`
	Xmpp    string `json:"xmpp,omitempty"`
}

type Location struct {
	Address  string  `json:"address,omitempty"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone,omitempty"`
}

type Machines struct {
	Mac  string `json:"mac"`
	Name string `json:"name,omitempty"`
}

type MembershipPlans struct {
	Currency        string  `json:"currency"`
	Description     string  `json:"description,omitempty"`
	Name            string  `json:"name"`
	RenewalInterval string  `json:"renewal_interval"`
	Value           float64 `json:"value"`
}

type NetworkConnections struct {
	Description string     `json:"description,omitempty"`
	Location    string     `json:"location,omitempty"`
	Machines    []Machines `json:"machines,omitempty"`
	Name        string     `json:"name,omitempty"`
	Type        string     `json:"type,omitempty"`
	Value       float64    `json:"value"`
}

type PeopleNowPresent struct {
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
	Name        string   `json:"name,omitempty"`
	Names       []string `json:"names,omitempty"`
	Value       float64  `json:"value"`
}

type PowerConsumption struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type Properties struct {
	Direction *Direction `json:"direction"`
	Elevation *Elevation `json:"elevation"`
	Gust      *Gust      `json:"gust"`
	Speed     *Speed     `json:"speed"`
}

type Radiation struct {
	Alpha     []Alpha     `json:"alpha,omitempty"`
	Beta      []Beta      `json:"beta,omitempty"`
	BetaGamma []BetaGamma `json:"beta_gamma,omitempty"`
	Gamma     []Gamma     `json:"gamma,omitempty"`
}

type RadioShow struct {
	End   string `json:"end"`
	Name  string `json:"name"`
	Start string `json:"start"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

type Sensors struct {
	AccountBalance     []AccountBalance     `json:"account_balance,omitempty"`
	Barometer          []Barometer          `json:"barometer,omitempty"`
	BeverageSupply     []BeverageSupply     `json:"beverage_supply,omitempty"`
	DoorLocked         []DoorLocked         `json:"door_locked,omitempty"`
	Humidity           []Humidity           `json:"humidity,omitempty"`
	NetworkConnections []NetworkConnections `json:"network_connections,omitempty"`
	PeopleNowPresent   []PeopleNowPresent   `json:"people_now_present,omitempty"`
	PowerConsumption   []PowerConsumption   `json:"power_consumption,omitempty"`
	Radiation          *Radiation           `json:"radiation,omitempty"`
	Temperature        []Temperature        `json:"temperature,omitempty"`
	TotalMemberCount   []TotalMemberCount   `json:"total_member_count,omitempty"`
	Wind               []Wind               `json:"wind,omitempty"`
}

type SpaceAPI014 struct {
	Api                 string            `json:"api"`
	Cache               *Cache            `json:"cache,omitempty"`
	Cam                 []string          `json:"cam,omitempty"`
	Contact             *Contact          `json:"contact"`
	Events              []Events          `json:"events,omitempty"`
	Feeds               *Feeds            `json:"feeds,omitempty"`
	IssueReportChannels []string          `json:"issue_report_channels"`
	Location            *Location         `json:"location"`
	Logo                string            `json:"logo"`
	MembershipPlans     []MembershipPlans `json:"membership_plans,omitempty"`
	Projects            []string          `json:"projects,omitempty"`
	RadioShow           []RadioShow       `json:"radio_show,omitempty"`
	Sensors             *Sensors          `json:"sensors,omitempty"`
	Space               string            `json:"space"`
	Spacefed            *Spacefed         `json:"spacefed,omitempty"`
	State               *State            `json:"state"`
	Stream              *Stream           `json:"stream,omitempty"`
	Url                 string            `json:"url"`
}

type Spacefed struct {
	Spacenet   bool `json:"spacenet"`
	Spacephone bool `json:"spacephone"`
	Spacesaml  bool `json:"spacesaml"`
}

type Speed struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type State struct {
	Icon          *Icon   `json:"icon,omitempty"`
	Lastchange    float64 `json:"lastchange,omitempty"`
	Message       string  `json:"message,omitempty"`
	Open          bool    `json:"open"`
	TriggerPerson string  `json:"trigger_person,omitempty"`
}

type Stream struct {
	M4      string `json:"m4,omitempty"`
	Mjpeg   string `json:"mjpeg,omitempty"`
	Ustream string `json:"ustream,omitempty"`
}

type Temperature struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

type TotalMemberCount struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Value       float64 `json:"value"`
}

type Wiki struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

type Wind struct {
	Description string      `json:"description,omitempty"`
	Location    string      `json:"location"`
	Name        string      `json:"name,omitempty"`
	Properties  *Properties `json:"properties"`
}
