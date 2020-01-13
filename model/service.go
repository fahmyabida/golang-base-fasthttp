package model

//Properties is used to map config file from json
type Properties struct {
	ServiceName string `json:"service_name"`
	IP          string `json:"ip_database"`
	Port        string `json:"port_database"`
	User        string `json:"user_database"`
	Password    string `json:"password_database"`
	DBName      string `json:"name_database"`
	Config      string `json:"config_name"`
	LogPath     string `json:"log_path"`
	TimeOut     string `json:"timeout"`
	PortServer  string `json:"port_server"`
}
