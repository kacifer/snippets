package snippets

import (
	"github.com/stretchr/testify/require"
	"net/url"
	"testing"
)

func TestURLDelQueryParam(t *testing.T) {
	assertions := require.New(t)
	u, err := url.Parse("/test?a=b")
	assertions.Nil(err)
	assertions.Equal("/test", URLDelQueryParam(u, "a"))
}

func TestURLSetQueryParam(t *testing.T) {
	assertions := require.New(t)
	u, err := url.Parse("/test?a=b")
	assertions.Nil(err)
	assertions.Equal("/test?a=b", URLSetQueryParam(u, "a", "b"))
}
