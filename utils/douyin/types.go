package douyin

type Header struct {
	Accept          string
	AcceptEncoding  string
	AcceptLanguage  string
	CacheControl    string
	Cookie          string
	Pragma          string
	Referer         string
	SecChUa         string
	SecChUaMobile   string
	SecChUaPlatform string
	SecFetchDest    string
	SecFetchMode    string
	SecFetchSite    string
	UserAgent       string
	Host            string
	Connection      string
}

type QueryValue struct {
	DevicePlatform  string
	Aid             string
	Channel         string
	PcClientType    string
	VersionCode     string
	VersionName     string
	CookieEnabled   string
	ScreenWidth     string
	ScreenHeight    string
	BrowserLanguage string
	BrowserPlatform string
	BrowserName     string
	BrowserVersion  string
	BrowserOnline   string
	EngineName      string
	EngineVersion   string
	OsName          string
	OsVersion       string
	CpuCoreNum      string
	DeviceMemory    string
	Platform        string
	Downlink        string
	EffectiveType   string
	RoundTripTime   string
	Webid           string
	MsToken         string
	XBogus          string
}
