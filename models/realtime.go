package models

type RealTime struct {
	UID      string  `json:"uid"`
	Nama     *string `json:"nama"`
	DeviceId string  `json:"device_id"`
	Waktu    string  `json:"waktu"`
}
