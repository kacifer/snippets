package snippets

import "net/url"

func URLDelQueryParam(u *url.URL, key string) string {
	q := u.Query()
	q.Del(key)
	if qs := q.Encode(); qs != "" {
		return u.Path + "?" + qs
	}
	return u.Path
}

func URLSetQueryParam(u *url.URL, key, value string) string {
	q := u.Query()
	q.Set(key, value)
	if qs := q.Encode(); qs != "" {
		return u.Path + "?" + qs
	}
	return u.Path
}
