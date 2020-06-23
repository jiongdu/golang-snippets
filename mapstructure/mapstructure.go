package mapstructure

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type ChatSetting struct {
	ConsumeTotal  string  `mapstructure:"consume_total"`
	FansClubLevel string  `mapstructure:"fans_club_level"`
	UserLevel     string  `mapstructure:"user_level"`
	OnlySeeFans   string `mapstructure:"only_see_fans"`
}

func main() {
	m := make(map[string]interface{})
	m["consume_total"] = "300"
	m["fans_club_level"] = "200"
	m["user_level"] = "100"
	m["only_see_fans"] = true

	chatSetting := ChatSetting{}
	mapstructure.Decode(m, &chatSetting)
	fmt.Println("value:", chatSetting.ConsumeTotal)
}
