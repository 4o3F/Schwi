package def

import jsoniter "github.com/json-iterator/go"

type Danmu struct {
	Time  float64 `json:"time"`
	Type  int     `json:"type"`
	Color int     `json:"color"`
	Uid   int     `json:"uid"`
	Text  string  `json:"text"`
}

type DanmuGet struct {
	Time  float64 `json:"time"`
	Type  int     `json:"type"`
	Color int     `json:"color"`
	Uid   int     `json:"uid"`
	Text  string  `json:"text"`
}

func (d *DanmuGet) MarshalJSON() (data []byte, err error) {
	return jsoniter.Marshal([]interface{}{d.Time, d.Type, d.Color, d.Uid, d.Text})
}
