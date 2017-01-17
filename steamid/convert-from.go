package steamid

import (
	"math/big"
	"strconv"
	"strings"
)

// FromIDTo64 converts a SteamID to a SteamID 64bit.
//     STEAM_0:0:86173181 -> 76561198132612090
func FromIDTo64(id string) uint64 {
	idParts := strings.Split(id, ":")
	magic, _ := new(big.Int).SetString("76561197960265728", 10)
	steam64, _ := new(big.Int).SetString(idParts[2], 10)
	steam64 = steam64.Mul(steam64, big.NewInt(2))
	steam64 = steam64.Add(steam64, magic)
	auth, _ := new(big.Int).SetString(idParts[1], 10)
	return steam64.Add(steam64, auth).Uint64()
}

// FromIDTo32 converts a SteamID to a SteamID 32bit.
//     STEAM_0:0:86173181 -> 172346362
func FromIDTo32(id string) uint32 {
	return From64To32(FromIDTo64(id))
}

// FromIDTo3 converts a SteamID to a SteamID 3.
//     STEAM_0:0:86173181 -> [U:1:172346362]
func FromIDTo3(id string) string {
	idParts := strings.Split(id, ":")
	idLastPart, err := strconv.ParseUint(idParts[len(idParts)-1], 10, 64)
	if err != nil {
		return ""
	}
	return "[U:1:" + strconv.FormatUint(idLastPart*2, 10) + "]"
}

// From64ToID converts a SteamID 64bit to a SteamID.
//     76561198132612090 -> STEAM_0:0:86173181
func From64ToID(id64 uint64) string {
	id := new(big.Int).SetInt64(int64(id64))
	magic, _ := new(big.Int).SetString("76561197960265728", 10)
	id = id.Sub(id, magic)
	isServer := new(big.Int).And(id, big.NewInt(1))
	id = id.Sub(id, isServer)
	id = id.Div(id, big.NewInt(2))
	return "STEAM_0:" + isServer.String() + ":" + id.String()
}

// From64To32 converts a SteamID 64bit to a SteamID 32bit.
//     76561198132612090 -> 172346362
func From64To32(id64 uint64) uint32 {
	id64Str := strconv.FormatUint(id64, 10)
	if len(id64Str) < 3 {
		return 0
	}
	id32, err := strconv.ParseInt(id64Str[3:], 10, 64)
	if err != nil {
		return 0
	}
	return uint32(id32 - 61197960265728)
}

// From64To3 converts a SteamID 64bit to a SteamID 3.
//     76561198132612090 -> [U:1:172346362]
func From64To3(id64 uint64) string {
	return FromIDTo3(From64ToID(id64))
}

// From32ToID converts a SteamID 32bit to a SteamID.
//     172346362 -> STEAM_0:0:86173181
func From32ToID(id32 uint32) string {
	return From64ToID(From32To64(id32))
}

// From32To64 converts a SteamID 32bit to a SteamID 64 bit.
//     172346362 -> 76561198132612090
func From32To64(id32 uint32) uint64 {
	idLong := strconv.FormatUint(uint64(id32)+61197960265728, 10)
	id64, err := strconv.ParseInt("765"+idLong, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(id64)
}

// From32To3 converts a SteamID 32bit to a SteamID 3.
//     172346362 -> [U:1:172346362]
func From32To3(id32 uint32) string {
	id := From32ToID(id32)
	return FromIDTo3(id)
}

// From3ToID converts a SteamID 3 to a SteamID.
//     172346362 -> [U:1:172346362]
func From3ToID(id3 string) string {
	id32 := From3To32(id3)
	if id32 == 0 {
		return ""
	}
	return From32ToID(id32)
}

// From3To64 converts a SteamID 3 to a SteamID 64bit.
//     [U:1:172346362] -> 76561198132612090
func From3To64(id3 string) uint64 {
	return From32To64(From3To32(id3))
}

// From3To32 converts a SteamID 3 to a SteamID 32bit.
//     [U:1:172346362] -> 172346362
func From3To32(id3 string) uint32 {
	id3Parts := strings.Split(id3, ":")
	id32Str := id3Parts[len(id3Parts)-1]
	if len(id32Str) <= 0 {
		return 0
	}
	if id32Str[len(id32Str)-1:] == "]" {
		id32Str = id32Str[:len(id32Str)-1]
	}
	id32, err := strconv.ParseUint(id32Str, 10, 64)
	if err != nil {
		return 0
	}
	return uint32(id32)
}
