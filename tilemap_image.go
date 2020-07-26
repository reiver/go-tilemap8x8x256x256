package tilemap8x8x256x256

import (
	"image"
	"image/color"
)

const (
	imageWidth  =  MapWidth  * TileWidth
	imageHeight =  MapHeight * TileHeight
)

func (receiver TileMap) At(x, y int) color.Color {
	if x < 0 || imageWidth <= x {
		return nil
	}
	if y < 0 || imageHeight <= y {
		return nil
	}

	var offset int
	{
		xx := x / TileWidth
		yy := y / TileHeight

		offset = receiver.mapOffset(xx,yy)
	}

	index := receiver.Map[offset]

	var c color.Color
	{
		tile := receiver.Tiles(index)

		xx := x % TileWidth
		yy := y % TileHeight

		c = tile.At(xx,yy)
	}

	return c
}

func (receiver TileMap) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	return image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x+imageWidth,
			Y: y+imageHeight,
		},
	}
}

func (receiver TileMap) ColorModel() color.Model {
	const anyTileID = 0

	return receiver.Tiles(anyTileID).ColorModel()
}

func (receiver TileMap) mapOffset(x int, y int) int {
	x = x % MapWidth
	if 0 > x {
		x += MapWidth
	}

	y = y % MapHeight
	if 0 > y {
		y += MapHeight
	}

	return y*MapWidth + x
}
