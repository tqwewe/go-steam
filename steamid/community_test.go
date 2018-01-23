package steamid

import "testing"

func TestResolveID(t *testing.T) {
	queries := []string{"STEAM_1:0:86173181", "76561198132612090", "172346362", "[U:1:172346362]"}
	expects := steamID64(76561198132612090)

	for _, query := range queries {
		result, _ := ResolveID(query, "E1FFB15B2C79FD99EFCE478B86B9E25A")
		if result != expects {
			t.Errorf("ResolveID failed. Query: %s", query)
		}
	}
}
