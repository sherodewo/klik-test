package controllers

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go-checkin/dto"
	"go-checkin/models"
	"go-checkin/service"
	"go-checkin/utils/session"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RoleController struct {
	BaseBackendController
	service *service.RoleService
}

func NewRoleController(service *service.RoleService) RoleController {
	return RoleController{
		BaseBackendController: BaseBackendController{
			Menu:        "Roles",
			BreadCrumbs: []map[string]interface{}{},
		},
		service: service,
	}
}
func (c *RoleController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "List Role",
		"link": "/check/admin/role",
	}
	return Render(ctx, "Roles List", "role/index", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *RoleController) List(ctx echo.Context) error {

	draw, err := strconv.Atoi(ctx.Request().URL.Query().Get("draw"))
	search := ctx.Request().URL.Query().Get("search[value]")
	start, err := strconv.Atoi(ctx.Request().URL.Query().Get("start"))
	length, err := strconv.Atoi(ctx.Request().URL.Query().Get("length"))
	order, err := strconv.Atoi(ctx.Request().URL.Query().Get("order[0][column]"))
	orderName := ctx.Request().URL.Query().Get("columns[" + strconv.Itoa(order) + "][name]")
	//orderAscDesc := ctx.Request().URL.Query().Get("order[0][dir]")

	recordTotal, recordFiltered, data, err := c.service.QueryDatatable(search, "desc", orderName, length, start)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	var action string

	//var role string
	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {

		action = `<div class="btn-group open">
					<button class="btn btn-xs dropdown-toggle" type="button" data-toggle="dropdown" aria-expanded="true"> Actions</button>
                      <ul class="dropdown-menu" role="menu">
                      <li>
                      <a href="/check/admin/role/edit/` + v.ID + `" style="text-decoration: none;font-weight: 400; color: #333;">
                      <i class="fa fa-edit"></i>Edit </a>
                      </li>
                      <li>
                      <a href="/check/admin/role/detail/` + v.ID + `"style="text-decoration: none;font-weight: 400; color: #333;">
                      <i class="fa fa-user"></i>Detail </a>
                      </li>
                      <li>
                      <a href="javascript:;" onclick="Delete('` + v.ID + `')" data-toggle="tooltip" data-placement="right" title="Delete" style="text-decoration: none;font-weight: 400; color: #333;">
                      <i class="fa fa-trash" style="color: #ff4d65;"></i> Delete </a>
                      </li>
                      </ul>
                      </div>`

		listOfData[k] = map[string]interface{}{
			"id":     v.ID,
			"name":   v.Name,
			"description": v.Description,
			"action": action,
		}
	}

	result := models.ResponseDatatable{
		Draw:            draw,
		RecordsTotal:    recordTotal,
		RecordsFiltered: recordFiltered,
		Data:            listOfData,
	}
	return ctx.JSON(http.StatusOK, &result)
}

func (c *RoleController) Add(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Add",
		"link": "/check/admin/role/add",
	}
	return Render(ctx, "Add Role", "role/add", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *RoleController) Store(ctx echo.Context) error {
	var roleDto dto.RoleDto
	if err := ctx.Bind(&roleDto); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	if err := ctx.Validate(&roleDto); err != nil {
		var validationErrors []models.ValidationError
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = models.WrapValidationErrors(errs)
		}
		return ctx.JSON(400, echo.Map{"message": "error validation", "errors": validationErrors})
	}

	result, err := c.service.SaveRole(roleDto)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error save data user"})
	}

	session.SetFlashMessage(ctx, "save data success", "success", nil)
	return ctx.JSON(200, echo.Map{"message": "data has been saved", "data": result})
}

func (c *RoleController) Edit(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := c.service.FindUserById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			session.SetFlashMessage(ctx, err.Error(), "error", nil)
			return ctx.Redirect(302, "/check/admin/role")
		}
		session.SetFlashMessage(ctx, err.Error(), "error", nil)
		return ctx.Redirect(302, "/check/admin/role")
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Edit",
		"link": "/check/admin/users/edit",
	}

	dataRole := models.UserRole{
		ID:     data.ID,
		Name:       data.Name,
		Description:      data.Description,
	}
	return Render(ctx, "Edit Role", "role/edit", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), dataRole)
}

func (c *RoleController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.service.DeleteRole(id)
	if err != nil {
		return ctx.JSON(500, echo.Map{"message": "error when trying delete data"})
	}
	return ctx.JSON(200, echo.Map{"message": "delete data has been deleted"})
}

func (c *RoleController) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var roleDto dto.RoleDto
	if err := ctx.Bind(&roleDto); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}

	if err := ctx.Validate(&roleDto); err != nil {
		var validationErrors []models.ValidationError
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = models.WrapValidationErrors(errs)
		}
		return ctx.JSON(400, echo.Map{"message": "error validation", "errors": validationErrors})
	}
	result, err := c.service.UpdateRole(id, roleDto)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error update data user"})
	}
	session.SetFlashMessage(ctx, "update data success", "success", nil)
	return ctx.JSON(200, echo.Map{"message": "data has been updated", "data": result})
}

func (c *RoleController) View(ctx echo.Context) error {
	id := ctx.Param("id")
	var data models.UserRole
	err := c.service.GetDbInstance().First(&data, "id =?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			session.SetFlashMessage(ctx, err.Error(), "error", nil)
			return ctx.Redirect(302, "/check/admin/role")
		}
		session.SetFlashMessage(ctx, err.Error(), "error", nil)
		return ctx.Redirect(302, "/check/admin/role")
	}
	breadCrumbs := map[string]interface{}{
		"menu": "Detail Role",
		"link": "/check/admin/role/detail",
	}
	return Render(ctx, "Detail Role ", "role/view", c.Menu, session.FlashMessage{},
		append(c.BreadCrumbs, breadCrumbs), echo.Map{"Role": data})
}