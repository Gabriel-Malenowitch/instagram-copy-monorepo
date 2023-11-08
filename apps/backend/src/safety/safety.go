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

func IsValidFollowPayload(user model.Followers) bool {
	// Validate here
	return true
}

func IsValidUnfollowPayload(user model.Followers) bool {
	// Validate here
	return true
}

func IsValidLoginPayload(user model.Login) bool {
	// Validate here
	return true
}
