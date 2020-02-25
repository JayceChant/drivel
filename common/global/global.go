package global

import (
	"fmt"

)

var (
	// CurrentVersion means literally
	CurrentVersion = &Version{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}
	// type check
	_ fmt.Stringer = (*Version)(nil)
)

// Version mark a application version
type Version struct {
	Major uint
	Minor uint
	Patch uint
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
