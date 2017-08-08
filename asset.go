// +build deployed

package engine

import (
	"fmt"
)

func LoadAsset(d string, f string) ([]byte, error) {
	buf, err := Asset(fmt.Sprintf("%v/%v", d, f))
	if err != nil {
		return nil, err
	}
	return buf, nil
}
