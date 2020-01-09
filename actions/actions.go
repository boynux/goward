package actions

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Action func(int32) error

func CreateTurnOnTVAction(key string) Action {
	return func(m int32) error {
		fmt.Println(key)
		_, err := http.Get("https://maker.ifttt.com/trigger/socket:on/with/key/" + key)
		if err != nil {
			log.Fatal("Failed to turn on TV", err)
		}

		_, err = http.PostForm("https://maker.ifttt.com/trigger/screen-time/with/key/" + key,
		url.Values{"value1": {fmt.Sprintf("%d", m)}})

		return err
	}
}
