package v1

import "superTools-background/internal/service"

/**
* @Author: super
* @Date: 2021-02-13 19:47
* @Description:
**/

type SpAttributeController struct {
	Service service.ISpAttributeService
}

func NewSpAttributeController(service service.ISpAttributeService) SpAttributeController {
	return SpAttributeController{Service: service}
}
