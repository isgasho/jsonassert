package internal

import (
	"testing"
)

func TestStringComparison(t *testing.T) {
	t.Run("degenerate case", func(st *testing.T) {
		tp, a := setup()
		a.checkString("$", "", "")
		if got := len(tp.messages); got != 0 {
			st.Errorf("expect no printed messages but there were %d", got)
		}
	})

	t.Run("unequal case", func(st *testing.T) {
		tp, a := setup()
		a.checkString("$", "Foo", "Bar")
		if len(tp.messages) != 1 {
			st.Errorf("expect exactly one printed message but there were %d", len(tp.messages))
		}
		exp, got := "expected value at '$' to be 'Bar' but was 'Foo'", tp.messages[0]
		if exp != got {
			st.Errorf("Expected error message '%s' but got '%s'", exp, got)
		}
	})

	t.Run("deeper level and case sensitive testing", func(st *testing.T) {
		tp, a := setup()
		a.checkString("$.here.wat", "foo", "Foo")
		if len(tp.messages) != 1 {
			st.Errorf("expect exactly one printed message but there were %d", len(tp.messages))
		}
		exp, got := "expected value at '$.here.wat' to be 'Foo' but was 'foo'", tp.messages[0]
		if exp != got {
			st.Errorf("Expected error message '%s' but got '%s'", exp, got)
		}
	})
}
