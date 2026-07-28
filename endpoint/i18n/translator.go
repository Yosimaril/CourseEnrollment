package i18n

import (
	"encoding/json"
	"os"
)

var languages = map[string]map[string]string{}

func LoadLanguage(lang string, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var messages map[string]string

	err = json.Unmarshal(data, &messages)
	if err != nil {
		return err
	}

	languages[lang] = messages

	return nil
}

func T(lang, key string) string {
	if messages, ok := languages[lang]; ok {

		if value, ok := messages[key]; ok {
			return value
		}
	}

	return key
}
