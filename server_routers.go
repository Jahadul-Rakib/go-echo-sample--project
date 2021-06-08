package main

import "github.com/labstack/echo/v4"
import (
	service "com/rakib/banking/main/service"
)

func Router(e *echo.Echo) {
	employeeRouterBase := e.Group("api/v1/employees")
	EmployeeRouter(employeeRouterBase)
}

func EmployeeRouter(baseRouter *echo.Group) {
	employee := service.Employee{}
	baseRouter.POST("", employee.CreateEmployee)
	baseRouter.GET("", employee.GetAllEmployee)
	baseRouter.GET("/:id", employee.GetEmployeeById)
	baseRouter.DELETE("/:id", employee.DeleteEmployeeById)
	baseRouter.PUT("/:id", employee.UpdateEmployeeById)
}
