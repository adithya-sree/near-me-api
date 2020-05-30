package db

//Location location db entry
type Location struct {
	Username   string `json:"username"`
	Created    string `json:"created"`
	LastUpdate string `json:"lastUpdate"`
	Location   string `json:"location"`
	EntryID    int    `json:"entryID"`
}
