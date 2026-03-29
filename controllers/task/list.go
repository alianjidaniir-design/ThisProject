package task

import (
	"ThisProject/apiSchema/commonSchema"
	"ThisProject/apiSchema/taskSchema"
	"ThisProject/controllers/mainController"
	"ThisProject/models/repositories"
	"ThisProject/statics/constants/controllerBaseErrCode"
	"github.com/gofiber/fiber/v2"
)

func List(ctx *fiber.Ctx) error {
	spanCtx := mainController.InitAPI(ctx, "12")

	defer mainController.FinishAPISpan(ctx)

	queryReq := taskSchema.ListRequest{}
	errStr, code, err := mainController.ParseQuery(ctx, &queryReq)
	if err != nil {
		return mainController.Error(ctx, controllerBaseErrCode.TaskErrCode, "01", errStr, code, err)
	}

	req := commonSchema.BaseRequest[taskSchema.ListRequest]{Body: queryReq}
	res, errStr, code, err := repositories.TaskRepo.List(spanCtx, req)
	if err != nil {
		return mainController.Error(ctx, controllerBaseErrCode.TaskErrCode, "02", errStr, code, err)
	}

	return mainController.Response(ctx, res)
}
