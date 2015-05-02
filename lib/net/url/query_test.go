package url

import (
	"net/url"
	"testing"

	"github.com/cosiner/gohper/lib/test"
)

func TestQueryEscape(t *testing.T) {
	tt := test.Wrap(t)
	qs := map[string]string{
		"A": "DD",
		"Z": "DD",
	}
	q, _ := EscapeEncode(qs, nil)
	s := string(q)
	tt.NE(url.QueryEscape("A=DD&Z=DD") == s, url.QueryEscape("Z=DD&A=DD") == s)
}

func TestParamString(t *testing.T) {
	tt := test.Wrap(t)
	tt.Eq(ParamString(1), "1")
	tt.Eq(ParamString("ddd"), "ddd")
	tt.Eq(ParamString([]int{1, 2, 3}), "1,2,3")
	tt.Eq(ParamString([]string{"1", "2", "3"}), "1,2,3")
	tt.Eq(ParamString([]byte("ddd")), "ddd")

	defer tt.Recover()

	ParamString(int64(2))
}