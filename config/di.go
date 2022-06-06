//go:build wireinject
// +build wireinject

package config

import (
	"github.com/google/wire"
	"klik/controllers"
	"klik/repository"
	"klik/service"
	"gorm.io/gorm"
)

func InjectUserController(db *gorm.DB) controllers.UserController {
	wire.Build(
		controllers.NewUserController,
		service.NewUserService,
		repository.NewUserRepository,
	)
	return controllers.UserController{}
}

func InjectDashboardController(db *gorm.DB) controllers.DashboardController {
	wire.Build(
		controllers.NewDashboardController,
		service.NewDashboardService,
		repository.NewDashboardRepository,
	)
	return controllers.DashboardController{}
}

func InjectAuthController(db *gorm.DB) controllers.AuthController {
	wire.Build(
		controllers.NewAuthController,
		service.NewAuthService,
		repository.NewUserRepository,
	)
	return controllers.AuthController{}
}

func InjectCustomerController(db *gorm.DB) controllers.CustomerController {
	wire.Build(
		controllers.NewCustomerController,
		service.NewCustomerService,
		repository.NewCustomerRepository,
	)
	return controllers.CustomerController{}
}

func InjectRoleController(db *gorm.DB) controllers.RoleController {
	wire.Build(
		controllers.NewRoleController,
		service.NewRoleService,
		repository.NewRoleRepository,
	)
	return controllers.RoleController{}
}

func InjectPermissionController(db *gorm.DB) controllers.PermissionController {
	wire.Build(
		controllers.NewPermissionController,
		service.NewPermissionService,
		repository.NewPermissionRepository,
	)
	return controllers.PermissionController{}
}

func InjectMenuController(db *gorm.DB) controllers.MenuController {
	wire.Build(
		controllers.NewMenuController,
		service.NewMenuService,
		repository.NewMenuRepository,
	)
	return controllers.MenuController{}
}

func InjectSubMenuController(db *gorm.DB) controllers.SubMenuController {
	wire.Build(
		controllers.NewSubMenuController,
		service.NewMenuService,
		repository.NewMenuRepository,
	)
	return controllers.SubMenuController{}
}
