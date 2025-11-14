// Division, District, Upazila structs
package model

type Division struct {
	ID     int64  `json:"id"`
	NameEn string `json:"name_en"`
	NameBn string `json:"name_bn"`
}

type District struct {
	ID         int64  `json:"id"`
	DivisionID int64  `json:"division_id"`
	NameEn     string `json:"name_en"`
	NameBn     string `json:"name_bn"`
}

type Upazila struct {
	ID         int64  `json:"id"`
	DistrictID int64  `json:"district_id"`
	NameEn     string `json:"name_en"`
	NameBn     string `json:"name_bn"`
}
