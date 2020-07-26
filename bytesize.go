package tilemap8x8x256x256

// ByteSize represents the number of bytes of the tile-map — i.e., how many uint8 are in the tile-map.
const ByteSize = (MapWidth * MapHeight) * ((TileWidth * TileHeight) * Depth)
