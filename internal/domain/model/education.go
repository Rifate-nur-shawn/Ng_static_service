// EducationLevel, Qualification
package model

type Qualification struct {
	ID           int64  `json:"id"`
	LevelID      int64  `json:"level_id"`
	StreamID     int64  `json:"stream_id"` // Denormalized for easier lookup
	Name         string `json:"name"`
	CommonName   string `json:"common_name"`
	EquivalentTo *int64 `json:"equivalent_to"`
}
