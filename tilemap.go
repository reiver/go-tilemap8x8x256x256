package tilemap8x8x256x256

import (
	"image"
)

// TileMap represents a tile-map.
type TileMap struct {
	Map []uint8

	Tiles func(uint8)image.Image
}
