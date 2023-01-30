// 共用資料模型
package models

import "encoding/json"

type Game struct {
	VendorCode string          `json:"vendor_code"`
	GameCode   string          `json:"game_code"`
	GameID     string          `json:"game_id"`
	Name       json.RawMessage `json:"name"`
	GameType   string          `json:"game_type"`
	Status     string          `json:"status"`
}

type GameList []Game

func (m GameList) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *GameList) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, m)
}

type Postgres struct {

}