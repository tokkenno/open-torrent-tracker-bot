package language

type translate struct {
	Lang string
	Text string
}

type message []translate

func Localize(langCode string, msg message) string {
	for _, translate := range msg {
		if langCode == translate.Lang {
			return translate.Text
		}
	}

	return msg[0].Text
}