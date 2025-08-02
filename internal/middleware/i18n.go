package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"letsplay-microservice/internal/locale"
)

func I18nMiddleware(c *gin.Context) {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = locale.DefaultLanguage
	}

	localizer := i18n.NewLocalizer(locale.Bundle, lang)

	c.Set(locale.LocalizerKey, localizer)
	c.Next()
}
