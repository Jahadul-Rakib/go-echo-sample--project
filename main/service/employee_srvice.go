package service

import (
	common "com/rakib/banking/main/common"
	"github.com/labstack/echo/v4"
	"log"
	"com/rakib/banking/main/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var databaseCon = database.GetDmManager()
var conn = databaseCon.Database.Collection("employee")

type Employee struct {
	Name    string `json:"name" xml:"name" bson:"name"`
	City    string `json:"current_city" xml:"city" bson:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code" bson:"zip_code"`
}

func (e Employee) CreateEmployee(context echo.Context) error {
	employeeData := new(Employee)

	if err := context.Bind(employeeData); err != nil {
		log.Println(common.DataBindingError, err.Error())
		return common.ErrorResponse(context, err.Error(), common.FailedToParsePayload)
	}
	data, errs := conn.InsertOne(databaseCon.Ctx, Employee{
		Name:    employeeData.Name,
		City:    employeeData.City,
		ZipCode: employeeData.ZipCode,
	})
	if errs != nil {
		return common.ErrorResponse(context, errs.Error(), common.DataInsertError)
	}
	return common.SuccessResponse(context, common.EmployeeFetchSuccessfully, data)
}

func (e Employee) GetEmployeeById(context echo.Context) error {
	employeeDTO := new(EmployeeDTO)
	id := context.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	err := conn.FindOne(databaseCon.Ctx, bson.M{"_id": objectId}).Decode(employeeDTO)
	if err != nil {
		return common.ErrorResponse(context, err.Error(), common.DataDecodeError)
	}
	return common.SuccessResponse(context, common.EmployeeGetById, employeeDTO)
}

func (e Employee) GetAllEmployee(context echo.Context) error {
	var data []EmployeeDTO
	cursor, err := conn.Find(databaseCon.Ctx, bson.M{})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), common.DataGetError)
	}
	for cursor.Next(databaseCon.Ctx) {
		emp := new(EmployeeDTO)
		er := cursor.Decode(emp)
		if er != nil {
			return common.ErrorResponse(context, er.Error(), common.DataDecodeError)
		}
		data = append(data, *emp)
	}
	return common.SuccessResponse(context, common.EmployeeFetchSuccessfully, data)
}

func (e Employee) DeleteEmployeeById(context echo.Context) error {
	id := context.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)
	_, err := conn.DeleteOne(databaseCon.Ctx, bson.M{"_id": objectID})
	if err != nil {
		return common.ErrorResponse(context, err.Error(), common.DeleteFailed)
	}
	return common.SuccessResponse(context, common.DeleteSuccessful, true)
}

func (e Employee) UpdateEmployeeById(context echo.Context) error {
	id := context.Param("id")
	objectID, _ := primitive.ObjectIDFromHex(id)

	empData := new(Employee)
	if err := context.Bind(empData); err != nil {
		return common.ErrorResponse(context, err.Error(), common.DataDecodeError)
	}
	filter := bson.M{"_id": objectID}
	updateModel := bson.M{
		"$set": &empData,
	}
	empDataDTO := new(EmployeeDTO)
	updateError := conn.FindOneAndUpdate(databaseCon.Ctx, filter, updateModel).Decode(empDataDTO)
	if updateError != nil {
		return common.ErrorResponse(context, updateError.Error(), common.DataDecodeError)
	}
	return common.SuccessResponse(context, common.UpdateSuccessful, empDataDTO)
}
