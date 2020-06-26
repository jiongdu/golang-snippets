package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type ChatSetting struct {
	ConsumeTotal  string          `mapstructure:"consume_total"`
	FansClubLevel map[string]bool `mapstructure:"fans_club_level"`
	UserLevel     string          `mapstructure:"user_level"`
	OnlySeeFans   string          `mapstructure:"only_see_fans"`
}

func main() {

	// mapstructure.Decode
	fansClubLevel := make(map[string]bool)
	fansClubLevel["300"] = true
	m := make(map[string]interface{})
	m["consume_total"] = "300"
	m["fans_club_level"] = fansClubLevel
	m["user_level"] = "100"
	m["only_see_fans"] = true

	chatSetting := ChatSetting{}
	mapstructure.Decode(m, &chatSetting)
	fmt.Println("value:", chatSetting.FansClubLevel)

}
