package models

// Timestamps holds all timestamps
type Timestamps struct {
	ResizeTimestamps []Timestamp `json:"resizeTimestamps"`
	TileTimestamps   []Timestamp `json:"tileTimestamps"`
}

// AllTimestamps global
var instance Timestamps

// Timestamp struct holding info for each request
type Timestamp struct {
	T1 int64 `json:"t1"`
	T2 int64 `json:"t2"`
	T3 int64 `json:"t3"`
	T4 int64 `json:"t4"`
	T5 int64 `json:"t5"`
}

// AddResizeTimestamp add a reset timestamp
func AddResizeTimestamp(timestamp Timestamp) error {
	instance.ResizeTimestamps = append(instance.ResizeTimestamps, timestamp)
	return nil
}

// AddTileTimestamp adds a tile timestamp
func AddTileTimestamp(timestamp Timestamp) error {
	instance.TileTimestamps = append(instance.TileTimestamps, timestamp)
	return nil
}

// Reset all timestamps
func Reset() error {
	instance.ResizeTimestamps = nil
	instance.TileTimestamps = nil
	return nil
}

// GetTimestamps gets all the timestamps
func GetTimestamps() Timestamps {
	return instance
}
