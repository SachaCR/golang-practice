package permissions

import "golang-practice/pkg/actor"

// Todo service is autonomous to set its internal permissions.

func Verify(anActor actor.Actor) bool {
	var isAuthorized bool = false

	if anActor.GetType() != actor.User {
		return isAuthorized
	}

	for _, role := range anActor.GetRoles() {
		if role == "basic-user" {
			isAuthorized = true
		}
	}

	return isAuthorized
}
