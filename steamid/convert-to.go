package steamid

import (
	"math/big"
	"strconv"
	"strings"
)

// To64 converts a SteamID to a SteamID 64bit.
//     STEAM_0:0:86173181 -> 76561198132612090
func (id ID) To64() ID64 {
	idParts := strings.Split(string(id), ":")
	magic, _ := new(big.Int).SetString("76561197960265728", 10)
	steam64, _ := new(big.Int).SetString(idParts[2], 10)
	steam64 = steam64.Mul(steam64, big.NewInt(2))
	steam64 = steam64.Add(steam64, magic)
	auth, _ := new(big.Int).SetString(idParts[1], 10)
	return ID64(steam64.Add(steam64, auth).Uint64())
}

// To32 converts a SteamID to a SteamID 32bit.
//     STEAM_0:0:86173181 -> 172346362
func (id ID) To32() ID32 {
	return id.To64().To32()
}

// To3 converts a SteamID to a SteamID 3.
//     STEAM_0:0:86173181 -> [U:1:172346362]
func (id ID) To3() ID3 {
	idParts := strings.Split(string(id), ":")
	idLastPart, err := strconv.ParseUint(idParts[len(idParts)-1], 10, 64)
	if err != nil {
		return ""
	}
	return ID3("[U:1:" + strconv.FormatUint(idLastPart*2, 10) + "]")
}

// ToID converts a SteamID 64bit to a SteamID.
//     76561198132612090 -> STEAM_0:0:86173181
func (id64 ID64) ToID() ID {
	id := new(big.Int).SetInt64(int64(id64))
	magic, _ := new(big.Int).SetString("76561197960265728", 10)
	id = id.Sub(id, magic)
	isServer := new(big.Int).And(id, big.NewInt(1))
	id = id.Sub(id, isServer)
	id = id.Div(id, big.NewInt(2))
	return ID("STEAM_0:" + isServer.String() + ":" + id.String())
}

// To32 converts a SteamID 64bit to a SteamID 32bit.
//     76561198132612090 -> 172346362
func (id64 ID64) To32() ID32 {
	id64Str := strconv.FormatUint(uint64(id64), 10)
	if len(id64Str) < 3 {
		return 0
	}
	id32, err := strconv.ParseInt(id64Str[3:], 10, 64)
	if err != nil {
		return 0
	}
	return ID32(id32 - 61197960265728)
}

// To3 converts a SteamID 64bit to a SteamID 3.
//     76561198132612090 -> [U:1:172346362]
func (id64 ID64) To3() ID3 {
	return id64.ToID().To3()
}

// ToID converts a SteamID 32bit to a SteamID.
//     172346362 -> STEAM_0:0:86173181
func (id32 ID32) ToID() ID {
	return id32.To64().ToID()
}

// To64 converts a SteamID 32bit to a SteamID 64 bit.
//     172346362 -> 76561198132612090
func (id32 ID32) To64() ID64 {
	idLong := strconv.FormatUint(uint64(id32)+61197960265728, 10)
	id64, err := strconv.ParseInt("765"+idLong, 10, 64)
	if err != nil {
		return 0
	}
	return ID64(id64)
}

// To3 converts a SteamID 32bit to a SteamID 3.
//     172346362 -> [U:1:172346362]
func (id32 ID32) To3() ID3 {
	return id32.ToID().To3()
}

// ToID converts a SteamID 3 to a SteamID.
//     172346362 -> [U:1:172346362]
func (id3 ID3) ToID() ID {
	id32 := id3.To32()
	if id32 == 0 {
		return ID("")
	}
	return id32.ToID()
}

// To64 converts a SteamID 3 to a SteamID 64bit.
//     [U:1:172346362] -> 76561198132612090
func (id3 ID3) To64() ID64 {
	return id3.To32().To64()
}

// To32 converts a SteamID 3 to a SteamID 32bit.
//     [U:1:172346362] -> 172346362
func (id3 ID3) To32() ID32 {
	id3Parts := strings.Split(string(id3), ":")
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
	return ID32(id32)
}
