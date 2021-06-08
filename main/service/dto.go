package service

type EmployeeDTO struct {
	Id      string `json:"id" xml:"id" bson:"_id"`
	Name    string `json:"name" xml:"name" bson:"name"`
	City    string `json:"current_city" xml:"city" bson:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code" bson:"zip_code"`
}
