package main

import (
	"github.com/RayleaBot/RayleaBot/sdk/go/pluginbuild"
	"github.com/RayleaBot/RayleaBot/sdk/go/pluginbuild/buildcmd"
)

func main() {
	buildcmd.Main(buildcmd.Config{
		BackendPackage: "./cmd/game-guide",
		Assets:         []string{"templates"},
		MappedAssets: []pluginbuild.AssetMapping{{
			Source: "internal/assets/characters.json", Destination: "data/characters.json",
		}},
	})
}
