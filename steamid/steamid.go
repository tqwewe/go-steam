/*
	Package steamid is used for general SteamID conversion to and from any existing SteamID.
	It also currently supports resolving a SteamID 64 from a search query.


	Converting a SteamID

	For simplicity, each SteamID type has the prefix "steam" removed resulting in each
	type begging with "ID".

	To convert a SteamID 64 bit (ID64) to a regular SteamID (ID) for example, you would
	use the following code:

		// Define type the SteamID 64 bit as type ID64.
		id64 := steamid.ID64(76561198132612090)

		// Convert id64 to a regular SteamID.
		id := id64.ToID()

	Although the code above is the preferred method for SteamID conversion (defining the
	SteamID type and using it's conversion function), you may also use other functions to
	convert using built-in types such as uint64 and string. You can accomplish this with
	the following code:

		// Convert a SteamID 64 bit to a regular SteamID.
		id := steamid.FromIDTo64(76561198132612090)


	Finding an ID64 with a Search Query

	ResolveID is a function available in this package that attempts to resolve a SteamID 64 from a
	search query. It resolves queries by first identifying what the query is (eg. a vanity url) and
	completes the required tasks to successfully resolve the query.

	You can use this function with a Steam web API key which is used to resolve the query
	as a vanity URL, or, if you do not specify a valid API key (eg. an empty string), then in
	most cases you will still be able to find the ID64 as long as the search query is a
	SteamID in another format.

	Here is an example of the ResolveID function in use:

		steamid.ResolveID("http://steamcommunity.com/id/gabelogannewell", "xxx")

	Which would return 76561197960287930 as an ID64.
*/
package steamid

// ID is a regular SteamID. STEAM_0:0:86173181
type ID string

// ID64 is a SteamID 64 bit. 76561198132612090
type ID64 uint64

// ID32 is a SteamID 32 bit. 172346362
type ID32 uint32

// ID3 is a SteamID v3. [U:1:172346362]
type ID3 string
