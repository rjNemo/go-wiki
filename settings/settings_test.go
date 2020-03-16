package settings

import (
	"os"
	"testing"
)

func TestReadEnv(t *testing.T) {
	e, err := os.Create(".envTest")
	if err != nil {
		t.Errorf("Error creating file %s", err)
	}
	e.Write([]byte("PORT=6543"))
	defer e.Close()

	// ans := NewParams(".envTest")
	// if ans.Port() != "6543" {
	// 	t.Errorf("readEnv('.envTest') = %s ; want 6543", ans)
	// }
}
