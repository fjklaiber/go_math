package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("zero length edge is not allowed")
}

// vMAJOR.MINOR.PATCH
// example: v1.2.3: major 1, minor 2, patch 3
