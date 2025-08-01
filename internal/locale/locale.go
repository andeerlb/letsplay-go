package locale

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
	"log"
)

var Bundle *i18n.Bundle

func Init() {
	Bundle := i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	_, err := Bundle.LoadMessageFile("locale/locale.en.toml")
	if err != nil {
		log.Fatal(err)
	}
	_, err = Bundle.LoadMessageFile("locale/locale.pt.toml")
	if err != nil {
		log.Fatal(err)
	}
}
