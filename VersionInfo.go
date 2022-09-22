package contracts

import (
	"go.dolittle.io/contracts/v7/versioning"
)

// GetCurrentVersion returns the current version of the Contracts.
func GetCurrentVersion() versioning.Version {
	return versioning.Version{
		Major:            7,
		Minor:            5,
		Patch:            0,
		PreReleaseString: "arwen.1",
	}
}
