package utils

import (
	"fmt"
	"mbenz_auth/pkg/constants"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role string) (string, error) {
	// Switch given role.
	switch role {
	case constants.AdminRoleName:
		// Nothing to do, verified successfully.
	//case constants.EngineerRoleName:
	//	// Nothing to do, verified successfully.
	case constants.OwnerRoleName:
		// Nothing to do, verified successfully.
	case constants.ChauffeurRoleName:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}
	return role, nil
}
