package def

import jsoniter "github.com/json-iterator/go"

type DanmuGetResponse struct {
	Time  float64 `json:"time"`
	Type  int     `json:"type"`
	Color int     `json:"color"`
	Uid   int     `json:"uid"`
	Text  string  `json:"text"`
}

type DanmuSendResponse struct {
	Uid    string     `json:"author"`
	Color  int     `json:"color"`
	Date   int64   `json:"date"`
	Ip     string  `json:"ip"`
	Id     string     `json:"id,omitempty"`
	Player int     `json:"player"`
	Text   string  `json:"text"`
	Time   float64 `json:"time"`
	Token  string  `json:"token,omitempty"`
	Type   int     `json:"type"`
}

type DanmuSaveType struct {
	Uid   int     `json:"uid"`
	Type  int     `json:"type"`
	Time  float64 `json:"time"`
	Color int     `json:"color"`
	Text  string  `json:"text"`
}

func (d *DanmuGetResponse) MarshalJSON() (data []byte, err error) {
	return jsoniter.Marshal([]interface{}{d.Time, d.Type, d.Color, d.Uid, d.Text})
}
