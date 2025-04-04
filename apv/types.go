package apv

/*
Config is the configuration struct for APV
It contains the username, password, url and skip-verify flag.
The username and password are used for authentication.
The url is the base URL for the APV API.
The skip-verify flag is used to skip TLS verification.
*/
type Config struct {
	Username   string `yaml:"username" mapstructure:"username"`
	Password   string `yaml:"password" mapstructure:"password"`
	URL        string `yaml:"url" mapstructure:"url"`
	SkipVerify bool   `yaml:"skip-verify" mapstructure:"skip-verify"`
	Debug      bool   `yaml:"debug" mapstructure:"debug"`
}

type Member struct {
	RealService  string `json:"real_service"`
	Weight       int    `json:"weight"`
	Priority     int    `json:"priority"`
	ActiveStatus bool   `json:"active_status"`
	ActiveReason string `json:"active_reason"`
}

type Group struct {
	InstanceId           string   `json:"instance_id"`
	Group                string   `json:"group"`
	RealService          string   `json:"real_service"`
	GroupName            string   `json:"group_name"`
	Method               string   `json:"method"`
	Activation           int      `json:"activation"`
	Failvoer             int      `json:"failover"`
	PriorityMode         bool     `json:"priority_mode"`
	Enable               bool     `json:"enable"`
	Protocol             string   `json:"protocol"`
	ProxyProtocol        bool     `json:"proxy_protocol"`
	Members              []Member `json:"members"`
	HealthRelation       string   `json:"health_relation"`
	HcTcpTempalte        []string `json:"hc_tcp_tempalte"`
	HcHttpTempalte       []string `json:"hc_http_tempalte"`
	GroupPolicyScopeName []string `json:"group_policy_scope_name"`
}

type SimpleGroup struct {
	InstanceId string   `json:"instance_id"`
	Members    []Member `json:"members"`
}

/*
GroupRequest is the request struct for adding a member to a group
*/
type GroupRequest struct {
	RealService string `json:"real_service"`
}

type GroupResponse struct {
	Group Group `json:"Group"`
}

type CliRequest struct {
	Cmd string `json:"cmd"`
}

type CliResponse struct {
	Contents string `json:"contents"`
}

type BatchCliResponse struct {
	Output string `json:"CLI Output"`
}
