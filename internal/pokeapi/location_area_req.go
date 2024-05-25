package internal

import (
	"fmt"
)

func (l LocationAreasResp) GetLocationNames() []byte {
	locationNames := ""
	for _, location := range l.Results {
		fmt.Println(location.Name)
		locationNames += location.Name + "\n"
	}
	locationNames = locationNames[:len(locationNames)-1]
	return []byte(locationNames)
}
