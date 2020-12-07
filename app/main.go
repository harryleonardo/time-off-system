package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	SharedConfig "github.com/time-off-system/shared/config"
	SharedContainer "github.com/time-off-system/shared/container"
	SharedContext "github.com/time-off-system/shared/context"
	SharedDatabase "github.com/time-off-system/shared/database"
	SharedLog "github.com/time-off-system/shared/log"
	SharedMiddleware "github.com/time-off-system/shared/middleware"
	SharedValidator "github.com/time-off-system/shared/validator"

	authRepository "github.com/time-off-system/domain/auth/repository"
	employeeRepository "github.com/time-off-system/domain/employee/repository"
	leaveRepository "github.com/time-off-system/domain/leave/repository"

	authUsecase "github.com/time-off-system/domain/auth/usecase"
	employeeUsecase "github.com/time-off-system/domain/employee/usecase"
	leaveUsecase "github.com/time-off-system/domain/leave/usecase"

	authHandler "github.com/time-off-system/domain/auth/delivery/http"
	employeeHandler "github.com/time-off-system/domain/employee/delivery/http"
	leaveHandler "github.com/time-off-system/domain/leave/delivery/http"
)

var (
	logger = SharedLog.NewLog()
)

func main() {
	// - initialize echo
	e := echo.New()

	// - get default dependency injection container
	container := SharedContainer.GetDefaultContainer()

	customizeMiddleware := container.MustGet("shared.middleware").(SharedMiddleware.Middleware)
	conf := container.MustGet("shared.config").(SharedConfig.ImmutableConfig)
	mysql := container.MustGet("shared.database").(SharedDatabase.MysqlInterface)

	mysqlSess, err := mysql.OpenMysqlConn()
	if err != nil {
		msgError := fmt.Sprintf("Failed to open mysql connection: %s", err.Error())
		logger.Errorf(msgError)
		panic(msgError)
	}

	defer mysqlSess.Close()

	// - declaring echo context;
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &SharedContext.ApplicationContext{
				Context:      c,
				Container:    container,
				MysqlSession: mysqlSess,
				Config:       conf,
			}
			return h(ac)
		}
	})

	//  register struct validator
	e.Validator = SharedValidator.DefaultValidator()
	e.HTTPErrorHandler = customizeMiddleware.ErrorHandler()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	// - root handler
	e.GET("/api", func(c echo.Context) error {
		ac := c.(*SharedContext.ApplicationContext)
		return ac.WithSuccess("success", 200, map[string]string{
			"time-off-system": "0.0.1",
		})
	})

	authRepo := authRepository.NewAuthRepository()
	employeeRepo := employeeRepository.NewEmployeeRepository()
	leaveRepo := leaveRepository.NewLeaveRepository()

	authUcase := authUsecase.NewAuthUsecase(authRepo)
	employeeUcase := employeeUsecase.NewEmployeeUsecase(employeeRepo)
	leaveUcase := leaveUsecase.NewLeaveUsecase(leaveRepo, employeeRepo)

	authHandler.AuthHandler(e, authUcase)
	employeeHandler.EmployeeHandler(e, employeeUcase)
	leaveHandler.LeaveHandler(e, leaveUcase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.GetPort())))
}
