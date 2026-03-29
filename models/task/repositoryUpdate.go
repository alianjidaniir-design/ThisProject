package task

import (
	"context"

	"ThisProject/apiSchema/commonSchema"
	"ThisProject/apiSchema/taskSchema"
	"ThisProject/statics/constants/status"
	"ThisProject/statics/customErr"
)

func (repo *Repository) Update(ctx context.Context, req commonSchema.BaseRequest[taskSchema.UpdateRequest]) (res taskSchema.UpdateResponse, errStr string, code int, err error) {
	if repo.initErr != nil {
		return taskSchema.UpdateResponse{}, "03", status.StatusInternalServerError, repo.initErr
	}

	updatedTask, found, err := repo.db().UpdateTask(ctx, req.Body)
	if err != nil {
		return taskSchema.UpdateResponse{}, "04", status.StatusInternalServerError, err
	}
	if !found {
		return taskSchema.UpdateResponse{}, "12", status.StatusBadRequest, customErr.TaskNotFound
	}

	repo.cache().InvalidateList()
	return taskSchema.UpdateResponse{Task: updatedTask}, "", status.StatusOK, nil
}
