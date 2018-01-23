# go-steam
Go-Steam is a repository for Steam development in Golang and is planned to contain multiple packages to make developing Steam-related projects as simple and practicable and simple as possible.

### Installation
In your command line:

	> go get -u github.com/Acidic9/go-steam/...

### Usage
**Converting an ID** [*(go-steam/steamaid)*](steamid)
```go
id   := steamid.NewID("STEAM_0:0:86173181")
id64 := id.To64() // 76561198132612090
id32 := id.To32() // 172346362
id3  := id.To3()  // [U:1:172346362]

fmt.Printf("ID: %s\nID64: %d\nID32: %d\nID3: %s\n", id, id64, id32, id3)
```

Output
```bash
ID:   STEAM_0:0:86173181
ID64: 76561198132612090
ID32: 172346362
ID3:  [U:1:172346362]
```

**Retrieving Player Summary** [*(go-steam/steamapi)*](steamapi)
```go
key := steamapi.NewKey("xxxxxxxxxx") // https://steamcommunity.com/dev/apikey

playerSummaries, err := key.GetSinglePlayerSummaries(76561198132612090)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Hello %s [SteamID 64: %v]", playerSummaries.PersonaName, playerSummaries.SteamID)
```

Output
```bash
Hello Acidic9 [SteamID: 76561198132612090]
```

### About
Hopefully this project can be a collaborative project with as many contributers as possible.

So far there are two packages, [steamid](steamid) and [steamapi](steamapi). The [steamid](steamid) package is very stable and contains functions for converting a Steam ID to and from the following types: 
 - Community ID
 - ID64
 - ID32
 - ID3

The [steamapi](steamapi) package aims to contain functions for every existing Steam API method available. It still has a little way to go.

### Contributing
The official guide [Contributing to Open Source on GitHub](https://guides.github.com/activities/contributing-to-open-source/#contributing) explains in detail how you can contribute to a project.

A quick explination:

1. Fork it
2. Create your feature branch (`git checkout -b new-feature`)
3. Commit your changes (`git commit -am 'Fix x code'`)
4. Push to the branch (`git push origin new-feature`)
5. Create new Pull Request
