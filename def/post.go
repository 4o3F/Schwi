package def

type Post struct {
	Pid      int    `json:"pid,omitempty"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Sort    string `json:"sort"`
	Tag     string `json:"tag,omitempty"`
	Time    string `json:"time"`
	Uid     int    `json:"uid,omitempty"`
	Username   string `json:"username,omitempty"`
}
