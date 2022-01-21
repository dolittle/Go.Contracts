package contracts

import (
	"go.dolittle.io/contracts/v6/versioning"
)

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() versioning.Version {
	return versioning.Version{
		Major:            6,
		Minor:            5,
		Patch:            0,
		PreReleaseString: "",
	}
}
