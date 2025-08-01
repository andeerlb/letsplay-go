package locale

import (
	"context"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const LocalizerKey = "localizer"
const DefaultLanguage = "pt"

func T(ctx context.Context, messageID string, templateData map[string]interface{}) string {
	localizer := localizerFromContext(ctx)

	return localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
	})
}

func Msg(ctx context.Context, messageID string) string {
	return T(ctx, messageID, nil)
}

func localizerFromContext(ctx context.Context) *i18n.Localizer {
	loc, ok := ctx.Value(LocalizerKey).(*i18n.Localizer)
	if !ok {
		return i18n.NewLocalizer(Bundle, DefaultLanguage)
	}
	return loc
}
