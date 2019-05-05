package main

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/FracKenA/op5beat/beater"
)

func main() {
	beat.Run("op5beat", "", beater.New)
}
