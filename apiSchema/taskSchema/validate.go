package taskSchema

import (
	"strings"

	"ThisProject/apiSchema/commonSchema"
	"ThisProject/statics/constants/status"
	"ThisProject/statics/customErr"
)

func (req *CreateRequest) Validate(validateExtraData commonSchema.ValidateExtraData) (string, int, error) {
	req.Title = strings.TrimSpace(req.Title)
	if req.Title == "" {
		return "03", status.StatusBadRequest, customErr.InvalidTitle
	}

	_ = validateExtraData
	return "", status.StatusOK, nil
}

func (req *ListRequest) Validate(validateExtraData commonSchema.ValidateExtraData) (string, int, error) {
	if req.Page < 1 {
		return "06", status.StatusBadRequest, customErr.InvalidPage
	}

	if req.PerPage < 1 || req.PerPage > 100 {
		return "09", status.StatusBadRequest, customErr.InvalidPerPage
	}

	_ = validateExtraData
	return "", status.StatusOK, nil
}

func (req *UpdateRequest) Validate(validateExtraData commonSchema.ValidateExtraData) (string, int, error) {
	if req.TaskID < 1 {
		return "03", status.StatusBadRequest, customErr.InvalidTaskID
	}

	if req.Title == nil && req.Description == nil {
		return "06", status.StatusBadRequest, customErr.InvalidUpdate
	}

	if req.Title != nil {
		trimmed := strings.TrimSpace(*req.Title)
		if trimmed == "" {
			return "09", status.StatusBadRequest, customErr.InvalidTitle
		}
		req.Title = &trimmed
	}

	if req.Description != nil {
		trimmed := strings.TrimSpace(*req.Description)
		req.Description = &trimmed
	}

	_ = validateExtraData
	return "", status.StatusOK, nil
}

func (req *DeleteRequest) Validate(validateExtraData commonSchema.ValidateExtraData) (string, int, error) {
	if req.TaskID < 1 {
		return "03", status.StatusBadRequest, customErr.InvalidTaskID
	}

	_ = validateExtraData
	return "", status.StatusOK, nil
}
