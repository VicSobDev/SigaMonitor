package bot

import "net/http"

func (b *Bot) GetHeaders() http.Header {
	return b.Headers.Clone()
}
