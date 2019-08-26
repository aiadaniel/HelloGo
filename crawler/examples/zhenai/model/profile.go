package model

import "encoding/json"

type Profile struct {
	Name         string
	Gender       int
	GenderString string
	Age          int
	Height       int
	Weight       int
	Income       string
	Marriage     string
	Education    string
	Occupation   string
	Hukou        string
	XingZuo      string
	House        string
	Car          string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	bytes, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(bytes, &profile)
	return profile, err
}
