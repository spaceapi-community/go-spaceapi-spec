package spaceapiStruct

// AccountBalance
type AccountBalance struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// Alpha
type Alpha struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

// Barometer
type Barometer struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// Beta
type Beta struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

// BetaGamma
type BetaGamma struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

// BeverageSupply
type BeverageSupply struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// Blog
type Blog struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

// Cache Specifies options about caching of your SpaceAPI endpoint. Use this if you want to avoid hundreds/thousands of application instances crawling your status.
type Cache struct {
	Schedule string `json:"schedule"`
}

// Calendar
type Calendar struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

// Contact Contact information about your space. You must define at least one which is in the list of allowed values of the issue_report_channels field.
type Contact struct {
	Email      string       `json:"email,omitempty"`
	Facebook   string       `json:"facebook,omitempty"`
	Foursquare string       `json:"foursquare,omitempty"`
	Google     *Google      `json:"google,omitempty"`
	Identica   string       `json:"identica,omitempty"`
	Irc        string       `json:"irc,omitempty"`
	IssueMail  string       `json:"issue_mail,omitempty"`
	Jabber     string       `json:"jabber,omitempty"`
	Keymasters []Keymasters `json:"keymasters,omitempty"`
	Ml         string       `json:"ml,omitempty"`
	Phone      string       `json:"phone,omitempty"`
	Sip        string       `json:"sip,omitempty"`
	Twitter    string       `json:"twitter,omitempty"`
}

// Direction The wind direction in degrees. Use this <a href="https://github.com/slopjong/OpenSpaceLint/issues/80" target="_blank_">mapping</a> to convert the degrees into a string.
type Direction struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

// DoorLocked
type DoorLocked struct {
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	Name        string `json:"name,omitempty"`
	Value       bool   `json:"value"`
}

// Elevation Height above mean sea level.
type Elevation struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

// Events
type Events struct {
	Extra     string  `json:"extra,omitempty"`
	Name      string  `json:"name"`
	Timestamp float64 `json:"timestamp"`
	Type      string  `json:"type"`
}

// Feeds Feeds where users can get updates of your space
type Feeds struct {
	Blog     *Blog     `json:"blog,omitempty"`
	Calendar *Calendar `json:"calendar,omitempty"`
	Flickr   *Flickr   `json:"flickr,omitempty"`
	Wiki     *Wiki     `json:"wiki,omitempty"`
}

// Flickr
type Flickr struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

// Gamma
type Gamma struct {
	ConversionFactor float64 `json:"conversion_factor,omitempty"`
	DeadTime         float64 `json:"dead_time,omitempty"`
	Description      string  `json:"description,omitempty"`
	Location         string  `json:"location,omitempty"`
	Name             string  `json:"name,omitempty"`
	Unit             string  `json:"unit"`
	Value            float64 `json:"value"`
}

// Google Google services.
type Google struct {
	Plus string `json:"plus,omitempty"`
}

// Gust
type Gust struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

// Humidity
type Humidity struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// Icon Icons that show the status graphically
type Icon struct {
	Closed string `json:"closed"`
	Open   string `json:"open"`
}

// Keymasters
type Keymasters struct {
	Email   string `json:"email,omitempty"`
	IrcNick string `json:"irc_nick,omitempty"`
	Name    string `json:"name,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Twitter string `json:"twitter,omitempty"`
}

// Location Position data such as a postal address or geographic coordinates
type Location struct {
	Address string  `json:"address,omitempty"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

// Machines
type Machines struct {
	Mac  string `json:"mac"`
	Name string `json:"name,omitempty"`
}

// NetworkConnections
type NetworkConnections struct {
	Description string     `json:"description,omitempty"`
	Location    string     `json:"location,omitempty"`
	Machines    []Machines `json:"machines,omitempty"`
	Name        string     `json:"name,omitempty"`
	Type        string     `json:"type,omitempty"`
	Value       float64    `json:"value"`
}

// PeopleNowPresent
type PeopleNowPresent struct {
	Description string   `json:"description,omitempty"`
	Location    string   `json:"location,omitempty"`
	Name        string   `json:"name,omitempty"`
	Names       []string `json:"names,omitempty"`
	Value       float64  `json:"value"`
}

// PowerConsumption
type PowerConsumption struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// Properties
type Properties struct {
	Direction *Direction `json:"direction"`
	Elevation *Elevation `json:"elevation"`
	Gust      *Gust      `json:"gust"`
	Speed     *Speed     `json:"speed"`
}

// Radiation Compound radiation sensor. Check this <a rel="nofollow" href="https://sites.google.com/site/diygeigercounter/gm-tubes-supported" target="_blank">resource</a>.
type Radiation struct {
	Alpha     []Alpha     `json:"alpha,omitempty"`
	Beta      []Beta      `json:"beta,omitempty"`
	BetaGamma []BetaGamma `json:"beta_gamma,omitempty"`
	Gamma     []Gamma     `json:"gamma,omitempty"`
}

// RadioShow
type RadioShow struct {
	End   string `json:"end"`
	Name  string `json:"name"`
	Start string `json:"start"`
	Type  string `json:"type"`
	Url   string `json:"url"`
}

// Sensors Data of various sensors in your space (e.g. temperature, humidity, amount of Club-Mate left, …). The only canonical property is the <em>temp</em> property, additional sensor types may be defined by you. In this case, you are requested to share your definition for inclusion in this specification.
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

// SpaceAPI013 SpaceAPI 0.13
type SpaceAPI013 struct {
	Api                 string      `json:"api"`
	Cache               *Cache      `json:"cache,omitempty"`
	Cam                 []string    `json:"cam,omitempty"`
	Contact             *Contact    `json:"contact"`
	Events              []Events    `json:"events,omitempty"`
	Feeds               *Feeds      `json:"feeds,omitempty"`
	IssueReportChannels []string    `json:"issue_report_channels"`
	Location            *Location   `json:"location"`
	Logo                string      `json:"logo"`
	Projects            []string    `json:"projects,omitempty"`
	RadioShow           []RadioShow `json:"radio_show,omitempty"`
	Sensors             *Sensors    `json:"sensors,omitempty"`
	Space               string      `json:"space"`
	Spacefed            *Spacefed   `json:"spacefed,omitempty"`
	State               *State      `json:"state"`
	Stream              *Stream     `json:"stream,omitempty"`
	Url                 string      `json:"url"`
}

// Spacefed A flag indicating if the hackerspace uses SpaceFED, a federated login scheme so that visiting hackers can use the space WiFi with their home space credentials.
type Spacefed struct {
	Spacenet   bool `json:"spacenet"`
	Spacephone bool `json:"spacephone"`
	Spacesaml  bool `json:"spacesaml"`
}

// Speed
type Speed struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

// State A collection of status-related data: actual open/closed status, icons, last change timestamp etc.
type State struct {
	Icon          *Icon   `json:"icon,omitempty"`
	Lastchange    float64 `json:"lastchange,omitempty"`
	Message       string  `json:"message,omitempty"`
	Open          bool    `json:"open"`
	TriggerPerson string  `json:"trigger_person,omitempty"`
}

// Stream A mapping of stream types to stream URLs.If you use other stream types make a <a href="add-your-space" target="_blank">change request</a> or prefix yours with <samp>ext_</samp> .
type Stream struct {
	M4      string `json:"m4,omitempty"`
	Mjpeg   string `json:"mjpeg,omitempty"`
	Ustream string `json:"ustream,omitempty"`
}

// Temperature
type Temperature struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location"`
	Name        string  `json:"name,omitempty"`
	Unit        string  `json:"unit"`
	Value       float64 `json:"value"`
}

// TotalMemberCount
type TotalMemberCount struct {
	Description string  `json:"description,omitempty"`
	Location    string  `json:"location,omitempty"`
	Name        string  `json:"name,omitempty"`
	Value       float64 `json:"value"`
}

// Wiki
type Wiki struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url"`
}

// Wind
type Wind struct {
	Description string      `json:"description,omitempty"`
	Location    string      `json:"location"`
	Name        string      `json:"name,omitempty"`
	Properties  *Properties `json:"properties"`
}
