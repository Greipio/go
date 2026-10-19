package greip

// ? Greip represents the Greip client
type Greip struct {
	token   string
	BaseURL string
	test    bool
}

type LookupIP struct {
	Address string `json:"address"`
	Version int    `json:"version"`
	Numeric int    `json:"numeric"`
	Bogon   bool   `json:"bogon"`
}

type LookupGeo struct {
	Continent          string  `json:"continent"`
	ContinentCode      string  `json:"continent_code"`
	ContinentGeoNameID int     `json:"continent_geo_name_id"`
	Country            string  `json:"country"`
	CountryCode        string  `json:"country_code"`
	CountryGeoNameID   int     `json:"country_geo_name_id"`
	Region             string  `json:"region"`
	City               string  `json:"city"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Capital            string  `json:"capital"`
	Population         int     `json:"population"`
	Phone              string  `json:"phone"`
	IsEU               bool    `json:"is_eu"`
	Neighbours         string  `json:"neighbours"`
	TLD                string  `json:"tld"`
}

type LookupASN struct {
	Number       string `json:"number"`
	Name         string `json:"name"`
	Organization string `json:"org"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Domain       string `json:"domain"`
	Created      string `json:"created"`
	Type         string `json:"type"`
}

type LookupLocation struct {
	Capital           string         `json:"capital"`
	Population        int            `json:"population"`
	Language          LookupLanguage `json:"language"`
	Flag              LookupFlag     `json:"flag"`
	PhoneCode         string         `json:"phoneCode"`
	CountryIsEU       bool           `json:"countryIsEU"`
	CountryNeighbours string         `json:"countryNeighbours"`
	TLD               string         `json:"tld"`
}

type LookupLanguage struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	NativeName string `json:"native"`
}

type LookupFlag struct {
	Emoji   string        `json:"emoji"`
	Unicode string        `json:"unicode"`
	PNG     LookupFlagPNG `json:"png"`
	SVG     string        `json:"svg"`
}

type LookupFlagPNG struct {
	Large  string `json:"1000px"`
	Medium string `json:"250px"`
	Small  string `json:"100px"`
}

type LookupTimezone struct {
	Name            string    `json:"name"`
	Abbreviation    string    `json:"abbreviation"`
	Offset          int       `json:"offset"`
	CurrentTime     string    `json:"currentTime"`
	CurrentTimeUnix int       `json:"currentTimestamp"`
	IsDST           bool      `json:"isDST"`
	Sun             LookupSun `json:"sunInfo"`
}

type LookupSun struct {
	Sunset                    string `json:"sunset"`
	Sunrise                   string `json:"sunrise"`
	Transit                   string `json:"transit"`
	CivilTwilightBegin        string `json:"civilTwilightBegin"`
	CivilTwilightEnd          string `json:"civilTwilightEnd"`
	NauticalTwilightBegin     string `json:"nauticalTwilightBegin"`
	NauticalTwilightEnd       string `json:"nauticalTwilightEnd"`
	AstronomicalTwilightBegin string `json:"astronomicalTwilightBegin"`
	AstronomicalTwilightEnd   string `json:"astronomicalTwilightEnd"`
	DayLength                 string `json:"dayLength"`
}

type LookupSecurity struct {
	IsProxy   bool   `json:"isProxy"`
	ProxyType string `json:"proxyType"`
	IsTor     bool   `json:"isTor"`
	IsBot     bool   `json:"isBot"`
	IsRelay   bool   `json:"isRelay"`
	IsHosting bool   `json:"isHosting"`
}

type LookupDevice struct {
	IsMobile bool          `json:"isMobile"`
	Type     string        `json:"type"`
	OS       LookupOS      `json:"OS"`
	Browser  LookupBrowser `json:"browser"`
}

type LookupOS struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Version   string `json:"version"`
	Title     string `json:"title"`
	BitMode64 string `json:"64bits_mode"`
}

type LookupBrowser struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	VersionMajor string `json:"versionMajor"`
	Title        string `json:"title"`
	UserAgent    string `json:"userAgent"`
}

type ResponseLookup struct {
	IP       LookupIP       `json:"ip"`
	Geo      LookupGeo      `json:"geo"`
	Location LookupLocation `json:"location"`
	ASN      LookupASN      `json:"asn"`
	Timezone LookupTimezone `json:"timezone"`
	Security LookupSecurity `json:"security"`
	Device   LookupDevice   `json:"device"`
}

type Threats struct {
	IsProxy   bool   `json:"isProxy"`
	ProxyType string `json:"proxyType"`
	IsTor     bool   `json:"isTor"`
	IsBot     bool   `json:"isBot"`
	IsRelay   bool   `json:"isRelay"`
	IsHosting bool   `json:"isHosting"`
}

type ThreatsSecurity struct {
	VpnOrProxy     bool   `json:"vpn_or_proxy"`
	VpnOrProxyType string `json:"vpn_or_proxy_type"`
	Tor            bool   `json:"tor"`
	Bot            bool   `json:"bot"`
	Relay          bool   `json:"relay"`
	Hosting        bool   `json:"hosting"`
	IsBlacklisted  bool   `json:"is_blacklisted"`
}

type ResponseThreats struct {
	IP             LookupIP        `json:"ip"`
	Score          int             `json:"score"`
	RiskLevel      string          `json:"risk_level"`
	Recommendation string          `json:"recommendation"`
	Security       ThreatsSecurity `json:"security"`
}

type CountryCurrency struct {
	Name   string `json:"currencyName"`
	Code   string `json:"currencyCode"`
	Symbol string `json:"currencySymbol"`
}

type ResponseCountry struct {
	Geo      LookupGeo       `json:"geo"`
	Language LookupLanguage  `json:"language"`
	Flag     LookupFlag      `json:"flag"`
	Currency CountryCurrency `json:"currency"`
	Timezone LookupTimezone  `json:"timezone"`
}

type ResponseProfanity struct {
	Text              string `json:"text"`
	TotalProfaneWords int    `json:"total_bad_words"`
	RiskScore         int    `json:"risk_score"`
	RiskLevel         string `json:"risk_level"`
	Recommendation    string `json:"recommendation"`
}

type ASNIPv4 struct {
	Total int `json:"total"`
}

type ASNIPv6 struct {
	Total int `json:"total"`
}

type ResponseASN struct {
	Number       string  `json:"number"`
	ASN          string  `json:"asn"`
	Name         string  `json:"name"`
	Organization string  `json:"org"`
	Phone        string  `json:"phone"`
	Email        string  `json:"email"`
	Domain       string  `json:"domain"`
	Created      string  `json:"created"`
	Type         string  `json:"type"`
	Registry     string  `json:"registry"`
	TotalIPs     int     `json:"totalIPs"`
	IPv4         ASNIPv4 `json:"IPv4"`
	IPv6         ASNIPv6 `json:"IPv6"`
}

type ResponseEmail struct {
	Score   int    `json:"score"`
	Reason  string `json:"reason"`
	IsValid bool   `json:"is_valid"`
	Email   string `json:"email"`
}

type ResponsePhone struct {
	Carrier     string `json:"carrier"`
	Reason      string `json:"reason"`
	IsValid     bool   `json:"is_valid"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country_code"`
}

type IBANFormats struct {
	Machine    string `json:"machine"`
	Human      string `json:"human"`
	Obfuscated string `json:"obfuscated"`
}

type IBANCountry struct {
	Name          string                 `json:"name"`
	IANA          string                 `json:"IANA"`
	ISO3166       string                 `json:"ISO3166"`
	Currency      string                 `json:"currency"`
	CentralBank   IBANCountryCentralBank `json:"centralBank"`
	Membership    string                 `json:"membership"`
	IsEU          bool                   `json:"isEU"`
	Length        string                 `json:"length"`
	SampleIBAN    string                 `json:"example_iban"`
	IsSepa        bool                   `json:"isSEPA"`
	SwiftOfficial bool                   `json:"swiftOfficial"`
}

type IBANCountryCentralBank struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type IBANBank struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	ShortName  string `json:"short_name"`
	BranchCode string `json:"branch_code"`
	BBAN       string `json:"bban"`
}

type ResponseIBAN struct {
	IsValid bool        `json:"is_valid"`
	IBAN    string      `json:"iban"`
	Formats IBANFormats `json:"formats"`
	Country IBANCountry `json:"country"`
	Bank    IBANBank    `json:"bank"`
}

type PaymentRule struct {
	Id          string `json:"id"`
	Description string `json:"description"`
}

type ResponsePayment struct {
	Score              float64       `json:"score"`
	Rules              []PaymentRule `json:"rules"`
	TotalRulesChecked  int           `json:"rulesChecked"`
	TotalRulesDetected int           `json:"rulesDetected"`
}
