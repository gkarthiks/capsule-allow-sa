package utils

import (
	"github.com/clastix/capsule/pkg/utils"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"strings"
)

func IsCapsuleUser(req admission.Request, userGroups []string, allowGitOpsSA string) bool {
	// if the user is a ServiceAccount belonging to the kube-system namespace and is used for GitOps engine
	// then Capsule should definitely honor the request for creation of Namespace and Tenant
	if strings.Compare(req.UserInfo.Username, allowGitOpsSA) == 0 {
		return true
	}
	groupList := utils.NewUserGroupList(req.UserInfo.Groups)
	// if the user is a ServiceAccount belonging to the kube-system namespace, definitely, it's not a Capsule user
	// and we can skip the check in case of Capsule user group assigned to system:authenticated
	// (ref: https://github.com/clastix/capsule/issues/234)
	if groupList.Find("system:serviceaccounts:kube-system") {
		return false
	}

	for _, group := range userGroups {
		if groupList.Find(group) {
			return true
		}
	}

	return false
}
