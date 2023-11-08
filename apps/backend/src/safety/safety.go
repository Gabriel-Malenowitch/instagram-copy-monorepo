package safety

import (
	"model"
)

func IsValidInsertUserPayload(user model.UserInsertPayload) bool {
	// Validate here
	return true
}

func IsValidUpdateUserPayload(user model.UserUpdatePayload) bool {
	// Validate here
	return true
}

func IsValidUpdateUserConfigurationPayload(user model.UserConfigurationPayload) bool {
	// Validate here
	return true
}
