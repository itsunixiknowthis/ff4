package ff4

import "testing"
import "github.com/itsunixiknowthis/ff2/v3"
import "github.com/itsunixiknowthis/ff3/v2"

func TestA(t *testing.T) {
    emb := Emb{ff2.T2{},ff3.T1d{}}
    if "T1.F" != emb.F() { t.Errorf("!") }
}

