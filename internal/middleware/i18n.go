package middleware

import (
	"strings"

	"letsplay-microservice/internal/locale"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const LocalizerKey = "localizer"
const defaultLanguage = "pt"

func I18nMiddleware(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = defaultLanguage
	}

	lang = strings.Split(lang, ",")[0]
	lang = strings.Split(lang, "-")[0]

	localizer := i18n.NewLocalizer(locale.Bundle, lang)

	c.Set(LocalizerKey, localizer)
	c.Next()
}

func LocalizerFromContext(c *gin.Context) *i18n.Localizer {
	loc, exists := c.Get(LocalizerKey)
	if !exists {
		return i18n.NewLocalizer(locale.Bundle, defaultLanguage)
	}
	localizer, ok := loc.(*i18n.Localizer)
	if !ok {
		return i18n.NewLocalizer(locale.Bundle, defaultLanguage)
	}
	return localizer
}
