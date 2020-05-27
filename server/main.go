package main

import (
	"github.com/kindaidensan/D-Chat/infrastructure"
)

func main() {
	infrastructure.Router.Run(":3000")
}
