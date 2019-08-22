package main

import (
	"github.com/yext/go-elasticsearch/internal/cmd/generate/commands"
	_ "github.com/yext/go-elasticsearch/internal/cmd/generate/commands/gensource"
	_ "github.com/yext/go-elasticsearch/internal/cmd/generate/commands/genstruct"
	_ "github.com/yext/go-elasticsearch/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
