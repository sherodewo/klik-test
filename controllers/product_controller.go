package controllers

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
	"klik/dto"
	"klik/models"
	"klik/service"
	"klik/utils/session"
	"net/http"
	"strconv"
)

type ProductController struct {
	BaseBackendController
	service *service.ProductService
}

func NewProductController(service *service.ProductService) ProductController {
	return ProductController{
		BaseBackendController: BaseBackendController{
			Menu:        "Product",
			BreadCrumbs: []map[string]interface{}{},
		},
		service: service,
	}
}
func (c *ProductController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "List Product",
		"link": "/klik/admin/product",
	}
	return Render(ctx, "Product List", "product/index", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), nil)
}
func (c *ProductController) List(ctx echo.Context) error {

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

	listOfData := make([]map[string]interface{}, len(data))
	for k, v := range data {

		action = `<div class="btn-group open">
					<button class="btn btn-xs dropdown-toggle" type="button" data-toggle="dropdown" aria-expanded="true"> Actions</button>
                      <ul class="dropdown-menu" role="menu">
                      	<li>
        	 				<a href="JavaScript:void(0);" onclick="Edit('` + v.SKU + `')" data-toggle="modal" data-target="#edit-modal" data-placement="right" title="Set Active"><i class="fa fa-lock-open"></i>Edit</a>
      					</li>
      					<li>
         					<a href="JavaScript:void(0);" onclick="Delete('` + v.SKU + `')" style="text-decoration: none;font-weight: 400; color: #333;" data-toggle="tooltip" data-placement="right" title="Delete"><i class="fa fa-trash" style="color: #ff4d65"></i>Delete</a>
      					</li>
                      </ul>
                      </div>`

		listOfData[k] = map[string]interface{}{
			"sku":          v.SKU,
			"product_name": v.ProductName,
			"action":       action,
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

func (c *ProductController) GetBySKU(ctx echo.Context) error {
	id := ctx.Param("sku")
	data, err := c.service.FindProductBySku(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, data)
}

func (c *ProductController) Store(ctx echo.Context) error {
	var req dto.ProductDto
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}
	if err := ctx.Validate(&req); err != nil {
		var validationErrors []models.ValidationError
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = models.WrapValidationErrors(errs)
		}
		return ctx.JSON(400, echo.Map{"message": "error validation", "errors": validationErrors})
	}

	result, err := c.service.SaveProduct(req)
	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error save data user"})
	}

	session.SetFlashMessage(ctx, "store data success", "success", nil)
	return ctx.JSON(200, echo.Map{"message": "data has been saved", "data": result})
}


func (c *ProductController) Update(ctx echo.Context) error {
	sku := ctx.Param("sku")

	var req dto.ProductDto

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{"message": "error binding data"})
	}

	if err := ctx.Validate(&req); err != nil {
		var validationErrors []models.ValidationError
		if errs, ok := err.(validator.ValidationErrors); ok {
			validationErrors = models.WrapValidationErrors(errs)
		}
		return ctx.JSON(400, echo.Map{"message": "error validation", "errors": validationErrors})
	}

	result, err := c.service.UpdateProduct(sku, req)


	if err != nil {
		return ctx.JSON(400, echo.Map{"message": "error update data"})
	}

	session.SetFlashMessage(ctx, "update data success", "success", nil)
	return ctx.JSON(200, echo.Map{"message": "data has been updated", "data": result})

}

func (c *ProductController) Delete(ctx echo.Context) error {
	sku := ctx.Param("sku")
	err := c.service.DeleteProduct(sku)
	if err != nil {
		return ctx.JSON(500, echo.Map{"message": "error when trying delete data"})
	}
	return ctx.JSON(200, echo.Map{"message": "delete data has been deleted"})
}