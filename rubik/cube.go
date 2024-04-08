package rubik

import (
	"reflect"
)

type Cube struct {
	T [][]byte // Top face
	B [][]byte // Back face
	L [][]byte // Left face
	F [][]byte // Front face
	R [][]byte // Right face
	D [][]byte // Down face
}

// Create + fill
func (cube *Cube) InitCube(size int) {
	cube.CreateCube(size)
	cube.FillCube(size)
}

func (cube *Cube) CreateCube(size int) {
	// 'Reflect' elements for iterating over struct fields.
	// Likely a bit overengineered, mainly to avoid repeating 'make' blocks e.g.
	// c.T = make([][]byte, size)
	// c.B = make([][]byte, size)
	// c.L = make([][]byte, size)
	// c.F = make([][]byte, size)
	// c.R = make([][]byte, size)
	// c.D = make([][]byte, size)
	//
	// for i := 0; i < size; i++ {
	// 	c.T[i] = make([]byte, size)
	// 	c.B[i] = make([]byte, size)
	// 	c.L[i] = make([]byte, size)
	// 	c.F[i] = make([]byte, size)
	// 	c.R[i] = make([]byte, size)
	// 	c.D[i] = make([]byte, size)
	// }

	cubeValues := reflect.ValueOf(cube).Elem()
	sliceType := reflect.SliceOf(reflect.TypeOf([]byte(nil)))

	// Iterate over the fields of the struct (cube faces)
	for i := 0; i < cubeValues.NumField(); i++ {
		faceValues := cubeValues.Field(i)
		// Create slices of specified size for current cube face
		faceValues.Set(reflect.MakeSlice(sliceType, size, size))
		// Iterate over face rows
		for j := 0; j < size; j++ {
			row := faceValues.Index(j)
			// Create inner slices of given size for current face row
			row.Set(reflect.MakeSlice(reflect.TypeOf([]byte(nil)), size, size))
		}
	}
}

func (cube *Cube) FillCube(size int) {
	// 'Reflect' elements for iterating over struct fields.
	// Likely a bit overengineered, mainly to avoid
	// separate loops for each struct field.

	cubeValues := reflect.ValueOf(cube).Elem()
	types := cubeValues.Type()

	// Iterate over the fields of the struct (cube faces)
	for i := 0; i < cubeValues.NumField(); i++ {
		faceValues := cubeValues.Field(i)
		faceName := types.Field(i).Name
		// Change fill color value based on current face
		var value byte
		switch faceName {
		case "T":
			value = 'y'
		case "B":
			value = 'b'
		case "L":
			value = 'r'
		case "F":
			value = 'g'
		case "R":
			value = 'o'
		case "D":
			value = 'w'
		}
		// Iterate over face rows
		for j := 0; j < size; j++ {
			row := faceValues.Index(j)
			// Fill the current row with the corresponding value
			for k := 0; k < size; k++ {
				row.Index(k).Set(reflect.ValueOf(value))
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Rotating section
// // Rotating cube faces in place should be implemented differently
// // to support bigger cubes - hardcoded for 3x3.
// // Even then - it's probably the section with biggest improvement potential.
// ----------------------------------------------------------------------------
func (cube *Cube) Rotate_R2() {
	cube.Rotate_R()
	cube.Rotate_R()
}

func (cube *Cube) Rotate_R() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.T[length-1][i]
		cube.T[length-1][i] = cube.F[length-1][i]
		cube.F[length-1][i] = cube.D[length-1][i]
		cube.D[length-1][i] = cube.B[0][length-1-i]
		cube.B[0][length-1-i] = temp[i]
	}
	// Rotating right face clockwise
	temp[0], temp[1], temp[2] = cube.R[0][0], cube.R[0][1], cube.R[0][2]
	cube.R[0][0], cube.R[0][1], cube.R[0][2] = cube.R[0][2], cube.R[1][2], cube.R[2][2]
	cube.R[0][2], cube.R[1][2], cube.R[2][2] = cube.R[2][2], cube.R[2][1], cube.R[2][0]
	cube.R[2][2], cube.R[2][1], cube.R[2][0] = cube.R[2][0], cube.R[1][0], cube.R[0][0]
	cube.R[2][0], cube.R[1][0], cube.R[0][0] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_Rc() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.T[length-1][i]
		cube.T[length-1][i] = cube.B[0][length-1-i]
		cube.B[0][length-1-i] = cube.D[length-1][i]
		cube.D[length-1][i] = cube.F[length-1][i]
		cube.F[length-1][i] = temp[i]
	}
	// Rotating right face counterclockwise
	temp[0], temp[1], temp[2] = cube.R[0][0], cube.R[0][1], cube.R[0][2]
	cube.R[0][0], cube.R[0][1], cube.R[0][2] = cube.R[2][0], cube.R[1][0], cube.R[0][0]
	cube.R[2][0], cube.R[1][0], cube.R[0][0] = cube.R[2][2], cube.R[2][1], cube.R[2][0]
	cube.R[2][2], cube.R[2][1], cube.R[2][0] = cube.R[0][2], cube.R[1][2], cube.R[2][2]
	cube.R[0][2], cube.R[1][2], cube.R[2][2] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_L2() {
	cube.Rotate_L()
	cube.Rotate_L()
}

func (cube *Cube) Rotate_Lc() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.T[0][i]
		cube.T[0][i] = cube.F[0][i]
		cube.F[0][i] = cube.D[0][i]
		cube.D[0][i] = cube.B[length-1][length-1-i]
		cube.B[length-1][length-1-i] = temp[i]
	}
	// Rotating left face clockwise
	temp[0], temp[1], temp[2] = cube.L[0][0], cube.L[0][1], cube.L[0][2]
	cube.L[0][0], cube.L[0][1], cube.L[0][2] = cube.L[2][0], cube.L[1][0], cube.L[0][0]
	cube.L[2][0], cube.L[1][0], cube.L[0][0] = cube.L[2][2], cube.L[2][1], cube.L[2][0]
	cube.L[2][2], cube.L[2][1], cube.L[2][0] = cube.L[0][2], cube.L[1][2], cube.L[2][2]
	cube.L[0][2], cube.L[1][2], cube.L[2][2] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_L() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.D[0][i]
		cube.D[0][i] = cube.F[0][i]
		cube.F[0][i] = cube.T[0][i]
		cube.T[0][i] = cube.B[length-1][length-1-i]
		cube.B[length-1][length-1-i] = temp[i]
	}
	// Rotating left face counterclockwise
	temp[0], temp[1], temp[2] = cube.L[0][0], cube.L[0][1], cube.L[0][2]
	cube.L[0][0], cube.L[0][1], cube.L[0][2] = cube.L[0][2], cube.L[1][2], cube.L[2][2]
	cube.L[0][2], cube.L[1][2], cube.L[2][2] = cube.L[2][2], cube.L[2][1], cube.L[2][0]
	cube.L[2][2], cube.L[2][1], cube.L[2][0] = cube.L[2][0], cube.L[1][0], cube.L[0][0]
	cube.L[2][0], cube.L[1][0], cube.L[0][0] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_U2() {
	cube.Rotate_U()
	cube.Rotate_U()
}

func (cube *Cube) Rotate_U() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.F[i][0]
		cube.F[i][0] = cube.R[i][0]
		cube.R[i][0] = cube.B[i][0]
		cube.B[i][0] = cube.L[i][0]
		cube.L[i][0] = temp[i]
	}
	// Rotating top face clockwise
	temp[0], temp[1], temp[2] = cube.T[0][0], cube.T[0][1], cube.T[0][2]
	cube.T[0][0], cube.T[0][1], cube.T[0][2] = cube.T[0][2], cube.T[1][2], cube.T[2][2]
	cube.T[0][2], cube.T[1][2], cube.T[2][2] = cube.T[2][2], cube.T[2][1], cube.T[2][0]
	cube.T[2][2], cube.T[2][1], cube.T[2][0] = cube.T[2][0], cube.T[1][0], cube.T[0][0]
	cube.T[2][0], cube.T[1][0], cube.T[0][0] = temp[0], temp[1], temp[2]

}

func (cube *Cube) Rotate_Uc() {
	length := len(cube.T)
	temp := make([]byte, length)
	for i := 0; i < length; i++ {
		temp[i] = cube.F[i][0]
		cube.F[i][0] = cube.L[i][0]
		cube.L[i][0] = cube.B[i][0]
		cube.B[i][0] = cube.R[i][0]
		cube.R[i][0] = temp[i]
	}
	// Rotating top face counterclockwise
	temp[0], temp[1], temp[2] = cube.T[0][0], cube.T[1][0], cube.T[2][0]
	cube.T[0][0], cube.T[1][0], cube.T[2][0] = cube.T[2][0], cube.T[2][1], cube.T[2][2]
	cube.T[2][0], cube.T[2][1], cube.T[2][2] = cube.T[2][2], cube.T[1][2], cube.T[0][2]
	cube.T[2][2], cube.T[1][2], cube.T[0][2] = cube.T[0][2], cube.T[0][1], cube.T[0][0]
	cube.T[0][2], cube.T[0][1], cube.T[0][0] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_DownCube() {
	tempFace := cube.D
	cube.D = cube.F
	cube.F = cube.T
	cube.T = cube.B
	cube.B = tempFace

	// Adjust orientation of back and top faces
	length := len(cube.D)
	for i := 0; i < length/2; i++ {
		cube.B[i], cube.B[length-1-i] = cube.B[length-1-i], cube.B[i]
		cube.T[i], cube.T[length-1-i] = cube.T[length-1-i], cube.T[i]
	}
	length = len(cube.T)
	temp := make([]byte, length)
	// Rotating right face counterclockwise
	temp[0], temp[1], temp[2] = cube.R[0][0], cube.R[0][1], cube.R[0][2]
	cube.R[0][0], cube.R[0][1], cube.R[0][2] = cube.R[2][0], cube.R[1][0], cube.R[0][0]
	cube.R[2][0], cube.R[1][0], cube.R[0][0] = cube.R[2][2], cube.R[2][1], cube.R[2][0]
	cube.R[2][2], cube.R[2][1], cube.R[2][0] = cube.R[0][2], cube.R[1][2], cube.R[2][2]
	cube.R[0][2], cube.R[1][2], cube.R[2][2] = temp[0], temp[1], temp[2]
	// Rotating left face clockwise
	temp[0], temp[1], temp[2] = cube.L[0][0], cube.L[0][1], cube.L[0][2]
	cube.L[0][0], cube.L[0][1], cube.L[0][2] = cube.L[0][2], cube.L[1][2], cube.L[2][2]
	cube.L[0][2], cube.L[1][2], cube.L[2][2] = cube.L[2][2], cube.L[2][1], cube.L[2][0]
	cube.L[2][2], cube.L[2][1], cube.L[2][0] = cube.L[2][0], cube.L[1][0], cube.L[0][0]
	cube.L[2][0], cube.L[1][0], cube.L[0][0] = temp[0], temp[1], temp[2]
}

func (cube *Cube) Rotate_UpCube() {
	tempFace := cube.T
	cube.T = cube.F
	cube.F = cube.D
	cube.D = cube.B
	cube.B = tempFace

	// Adjust orientation of down and back faces
	length := len(cube.D)
	for i := 0; i < length/2; i++ {
		cube.D[i], cube.D[length-1-i] = cube.D[length-1-i], cube.D[i]
		cube.B[i], cube.B[length-1-i] = cube.B[length-1-i], cube.B[i]
	}
	length = len(cube.T)
	temp := make([]byte, length)
	// Rotating right face clockwise
	temp[0], temp[1], temp[2] = cube.R[0][0], cube.R[0][1], cube.R[0][2]
	cube.R[0][0], cube.R[0][1], cube.R[0][2] = cube.R[0][2], cube.R[1][2], cube.R[2][2]
	cube.R[0][2], cube.R[1][2], cube.R[2][2] = cube.R[2][2], cube.R[2][1], cube.R[2][0]
	cube.R[2][2], cube.R[2][1], cube.R[2][0] = cube.R[2][0], cube.R[1][0], cube.R[0][0]
	cube.R[2][0], cube.R[1][0], cube.R[0][0] = temp[0], temp[1], temp[2]
	// Rotating left face counterclockwise
	temp[0], temp[1], temp[2] = cube.L[0][0], cube.L[0][1], cube.L[0][2]
	cube.L[0][0], cube.L[0][1], cube.L[0][2] = cube.L[2][0], cube.L[1][0], cube.L[0][0]
	cube.L[2][0], cube.L[1][0], cube.L[0][0] = cube.L[2][2], cube.L[2][1], cube.L[2][0]
	cube.L[2][2], cube.L[2][1], cube.L[2][0] = cube.L[0][2], cube.L[1][2], cube.L[2][2]
	cube.L[0][2], cube.L[1][2], cube.L[2][2] = temp[0], temp[1], temp[2]
}
