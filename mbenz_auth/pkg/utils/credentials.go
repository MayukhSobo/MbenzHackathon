package utils

import (
	"fmt"
	"mbenz_auth/pkg/constants"
)

// GetCredentialsByRole gets access credentials based on the role
func GetCredentialsByRole(role string) ([]string, error){
	var credentials []string

	switch role {
	case constants.AdminRoleName:
		// admin get all the access
		credentials = []string{
			constants.GetAllUsersCred,
			constants.GetUserCred,
			constants.DeleteUserCred,
			constants.AccessPOCApiCred,
			constants.AccessPlannerAPICred,
		}

	case constants.OwnerRoleName:
		credentials = []string{
			constants.GetUserCred,
			constants.DeleteUserCred,
			constants.AccessPOCApiCred,
			constants.AccessPlannerAPICred,
		}
	case constants.ChauffeurRoleName:
		credentials = []string{
			constants.AccessPOCApiCred,
			constants.AccessPlannerAPICred,
		}

	default:
		return nil, fmt.Errorf("role %v does't exist", role)
	}
	return credentials, nil
}