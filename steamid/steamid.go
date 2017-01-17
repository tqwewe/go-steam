package steamid

// ID is a regular SteamID. STEAM_0:0:86173181
type ID string

// ID64 is a SteamID 64bit. 76561198132612090
type ID64 uint64

// ID32 is a SteamID 32bit. 172346362
type ID32 uint32

// ID3 is a SteamID v3. [U:1:172346362]
type ID3 string
