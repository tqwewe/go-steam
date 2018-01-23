package steamid

import "testing"

var (
	expectedID = steamID("STEAM_0:0:86173181")
	expected64 = steamID64(76561198132612090)
	expected32 = steamID32(172346362)
	expected3  = steamID3("[U:1:172346362]")
)

func TestIDTo64(t *testing.T) {
	id64 := expectedID.To64()
	if id64 != expected64 {
		t.Errorf("ID -> ID64 wrong result. Expected: %v, Got: %v", expected64, id64)
	}
}

func TestIDTo32(t *testing.T) {
	id32 := expectedID.To32()
	if id32 != expected32 {
		t.Errorf("ID -> ID32 wrong result. Expected: %v, Got: %v", expected32, id32)
	}
}

func TestIDTo3(t *testing.T) {
	id3 := expectedID.To3()
	if id3 != expected3 {
		t.Errorf("ID -> ID3 wrong result. Expected: %v, Got: %v", expected3, id3)
	}
}

func TestID64ToID(t *testing.T) {
	id := expected64.ToID()
	if id != expectedID {
		t.Errorf("ID64 -> ID wrong result. Expected: %v, Got: %v", expectedID, id)
	}
}

func TestID64To32(t *testing.T) {
	id32 := expected64.To32()
	if id32 != expected32 {
		t.Errorf("ID64 -> ID32 wrong result. Expected: %v, Got: %v", expected32, id32)
	}
}

func TestID64To3(t *testing.T) {
	id3 := expected64.To3()
	if id3 != expected3 {
		t.Errorf("ID64 -> ID3 wrong result. Expected: %v, Got: %v", expected3, id3)
	}
}

func TestID32ToID(t *testing.T) {
	id := expected32.ToID()
	if id != expectedID {
		t.Errorf("ID32 -> ID wrong result. Expected: %v, Got: %v", expectedID, id)
	}
}

func TestID32To64(t *testing.T) {
	id64 := expected32.To64()
	if id64 != expected64 {
		t.Errorf("ID32 -> ID64 wrong result. Expected: %v, Got: %v", expected64, id64)
	}
}

func TestID32To3(t *testing.T) {
	id3 := expected32.To3()
	if id3 != expected3 {
		t.Errorf("ID32 -> ID3 wrong result. Expected: %v, Got: %v", expected3, id3)
	}
}

func TestID3ToID(t *testing.T) {
	id := expected3.ToID()
	if id != expectedID {
		t.Errorf("ID3 -> ID wrong result. Expected: %v, Got: %v", expectedID, id)
	}
}

func TestID3To64(t *testing.T) {
	id64 := expected3.To64()
	if id64 != expected64 {
		t.Errorf("ID3 -> ID64 wrong result. Expected: %v, Got: %v", expected64, id64)
	}
}

func TestID3To32(t *testing.T) {
	id32 := expected3.To32()
	if id32 != expected32 {
		t.Errorf("ID3 -> ID32 wrong result. Expected: %v, Got: %v", expected32, id32)
	}
}
