package spaceapi_spec

type Location struct {
	Address	string `json:"address,omitempty"`
	Lat	float64 `json:"lat"`
	Lon	float64 `json:"lon"`
}

type Spacefed struct {
	Spacenet	bool `json:"spacenet"`
	Spacesaml	bool `json:"spacesaml"`
	Spacephone	bool `json:"spacephone"`
}

type Stream struct {
	M4	string `json:"m4,omitempty"`
	Mjpeg	string `json:"mjpeg,omitempty"`
	Ustream	string `json:"ustream,omitempty"`
}

type Icon struct {
	Open	string `json:"open"`
	Closed	string `json:"closed"`
}

type State struct {
	Open          bool `json:"open"`
	Lastchange    int `json:"lastchange,omitempty"`
	TriggerPerson string `json:"trigger_person,omitempty"`
	Message       string `json:"message,omitempty"`
	Icon          *Icon `json:"icon,omitempty"`
}

type Event struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Timestamp int `json:"timestamp"`
	Extra string `json:"extra,omitempty"`
}

type Keymaster struct {
	Name	string `json:"name,omitempty"`
	IrcNick	string `json:"irc_nick,omitempty"`
	Phone	string `json:"phone,omitempty"`
	Email	string `json:"email,omitempty"`
	Twitter	string `json:"twitter,omitempty"`
}

type Google struct {
	Plus	string `json:"plus,omitempty"`
}

type Contact struct {
	Phone      string `json:"phone,omitempty"`
	Sip        string `json:"sip,omitempty"`
	Keymasters []Keymaster `json:"keymasters,omitempty"`
	Irc        string `json:"irc,omitempty"`
	Twitter    string `json:"twitter,omitempty"`
	Facebook   string `json:"facebook,omitempty"`
	Google     *Google `json:"google,omitempty"`
	Identica   string `json:"identica,omitempty"`
	Foursquare string `json:"foursquare,omitempty"`
	Email      string `json:"email,omitempty"`
	Ml         string `json:"ml,omitempty"`
	Jabber     string `json:"jabber,omitempty"`
	IssueMail  string `json:"issue_mail,omitempty"`
}

type Temperature struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type DoorLocked struct {
	Value	bool `json:"value"`
	Location	string `json:"location"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Barometer struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Alpha struct {
	Value float64 `json:"value"`
	Unit	string `json:"unit"`
	DeadTime	float64 `json:"dead_time,omitempty"`
	ConversionFactor	float64 `json:"conversion_factor,omitempty"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Beta struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	DeadTime	float64 `json:"dead_time,omitempty"`
	ConversionFactor	float64 `json:"conversion_factor,omitempty"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Gamma struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	DeadTime	float64 `json:"dead_time,omitempty"`
	ConversionFactor	float64 `json:"conversion_factor,omitempty"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type BetaGamma struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	DeadTime	float64 `json:"dead_time,omitempty"`
	ConversionFactor	float64 `json:"conversion_factor,omitempty"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Radiation struct {
	Alpha	[]Alpha `json:"alpha,omitempty"`
	Beta	[]Beta `json:"beta,omitempty"`
	Gamma	[]Gamma `json:"gamma,omitempty"`
	BetaGamma	[]BetaGamma `json:"beta_gamma,omitempty"`
}

type Humidity struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type BeverageSupply struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type PowerConsumption struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Speed struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
}

type Gust struct {
	Value	float64 `json:"value,omitempty"`
	Unit	string `json:"unit,omitempty"`
}

type Direction struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
}

type Elevation struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
}

type Properties struct {
	Speed     Speed `json:"speed"`
	Gust      Gust `json:"gust"`
	Direction Direction `json:"direction"`
	Elevation Elevation `json:"elevation"`
}

type Wind struct {
	Properties  Properties `json:"properties"`
	Location    string `json:"location"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Machine struct {
	Name	string `json:"name,omitempty"`
	Mac	string `json:"mac"`
}

type NetworkConnection struct {
	Type	string `json:"type,omitempty"`
	Value	float64 `json:"value"`
	Machines	[]Machine `json:"machines,omitempty"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type AccountBalance struct {
	Value	float64 `json:"value"`
	Unit	string `json:"unit"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type TotalMemberCount struct {
	Value	float64 `json:"value"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Description	string `json:"description,omitempty"`
}

type PeopleNowPresent struct {
	Value	float64 `json:"value"`
	Location	string `json:"location,omitempty"`
	Name	string `json:"name,omitempty"`
	Names	[]string `json:"names,omitempty"`
	Description	string `json:"description,omitempty"`
}

type Sensors struct {
	Temperature        []Temperature `json:"temperature,omitempty"`
	DoorLocked         []DoorLocked `json:"door_locked,omitempty"`
	Barometer          []Barometer `json:"barometer,omitempty"`
	Radiation          *Radiation `json:"radiation,omitempty"`
	Humidity           []Humidity `json:"humidity,omitempty"`
	BeverageSupply     []BeverageSupply `json:"beverage_supply,omitempty"`
	PowerConsumption   []PowerConsumption `json:"power_consumption,omitempty"`
	Wind               []Wind `json:"wind,omitempty"`
	NetworkConnections []NetworkConnection `json:"network_connections,omitempty"`
	AccountBalance     []AccountBalance `json:"account_balance,omitempty"`
	TotalMemberCount   []TotalMemberCount `json:"total_member_count,omitempty"`
	PeopleNowPresent   []PeopleNowPresent `json:"people_now_present,omitempty"`
}

type Feed struct {
	Type	string `json:"type,omitempty"`
	Url	string `json:"url"`
}

type Feeds struct {
	Blog     *Feed `json:"blog,omitempty"`
	Wiki     *Feed `json:"wiki,omitempty"`
	Calendar *Feed `json:"calendar,omitempty"`
	Flickr   *Feed `json:"flickr,omitempty"`
}

type Cache struct {
	Schedule	string `json:"schedule"`
}

type RadioShow struct {
	Name	string `json:"name"`
	Url	string `json:"url"`
	Type	string `json:"type"`
	Start	string `json:"start"`
	End	string `json:"end"`
}

type Root struct {
	Api                 string `json:"api"`
	Space               string `json:"space"`
	Logo                string `json:"logo"`
	Url                 string `json:"url"`
	Location            *Location `json:"location"`
	Spacefed            *Spacefed `json:"spacefed,omitempty"`
	Cam                 []string `json:"cam,omitempty"`
	Stream              *Stream `json:"stream,omitempty"`
	State               *State `json:"state,omitempty"`
	Events              []Event `json:"events,omitempty"`
	Contact             *Contact `json:"contact,omitempty"`
	IssueReportChannels []string `json:"issue_report_channels,omitempty"`
	Sensors             *Sensors `json:"sensors,omitempty"`
	Feeds               *Feeds `json:"feeds,omitempty"`
	Cache               *Cache `json:"cache,omitempty"`
	Projects            []string `json:"projects,omitempty"`
	RadioShow           []RadioShow `json:"radio_show,omitempty"`
}