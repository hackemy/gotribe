package models

type UUID string

type SDKUUID struct {
	SDKType string `json:"_sdkType"`
	UUID    string `json:"uuid"`
}

func (u UUID) String() string {
	return string(u)
}
