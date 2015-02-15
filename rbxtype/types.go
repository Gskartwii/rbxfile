// The rbxtype package implements in-memory representations of Roblox data
// types.
package rbxtype

import (
	"errors"
	"github.com/robloxapi/rbxfile"
	"strconv"
	"strings"
)

var ErrUnknownType = errors.New("unknown type")

// Type holds a value of a particular Roblox type.
type Type interface {
	// TypeString returns the name of the type.
	TypeString() string

	// String returns a string representation of the type's current value.
	String() string

	// Copy returns a copy of the value.
	Copy() Type
}

func joinstr(a ...string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}
	n := 0
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := 0
	for _, s := range a {
		bp += copy(b[bp:], s)
	}
	return string(b)
}

////////////////////////////////////////////////////////////////
// Types

type String []byte

func (String) TypeString() string {
	return "string"
}
func (t String) String() string {
	return string(t)
}
func (t String) Copy() Type {
	c := make(String, len(t))
	copy(c, t)
	return c
}

////////////////

type BinaryString []byte

func (BinaryString) TypeString() string {
	return "BinaryString"
}
func (t BinaryString) String() string {
	return string(t)
}
func (t BinaryString) Copy() Type {
	c := make(BinaryString, len(t))
	copy(c, t)
	return c
}

////////////////

type ProtectedString []byte

func (ProtectedString) TypeString() string {
	return "ProtectedString"
}
func (t ProtectedString) String() string {
	return string(t)
}
func (t ProtectedString) Copy() Type {
	c := make(ProtectedString, len(t))
	copy(c, t)
	return c
}

////////////////

type Content []byte

func (Content) TypeString() string {
	return "Content"
}
func (t Content) String() string {
	return string(t)
}
func (t Content) Copy() Type {
	c := make(Content, len(t))
	copy(c, t)
	return c
}

////////////////

type Bool bool

func (Bool) TypeString() string {
	return "bool"
}
func (t Bool) String() string {
	if t {
		return "true"
	} else {
		return "false"
	}
}
func (t Bool) Copy() Type {
	return t
}

////////////////

type Int int32

func (Int) TypeString() string {
	return "int"
}
func (t Int) String() string {
	return strconv.FormatInt(int64(t), 10)
}
func (t Int) Copy() Type {
	return t
}

////////////////

type Float float32

func (Float) TypeString() string {
	return "float"
}
func (t Float) String() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 32)
}
func (t Float) Copy() Type {
	return t
}

////////////////

type Double float64

func (Double) TypeString() string {
	return "double"
}
func (t Double) String() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
func (t Double) Copy() Type {
	return t
}

////////////////

type UDim struct {
	Scale  float32
	Offset int32
}

func (UDim) TypeString() string {
	return "UDim"
}
func (t UDim) String() string {
	return joinstr(
		strconv.FormatFloat(float64(t.Scale), 'f', -1, 32),
		", ",
		strconv.FormatInt(int64(t.Offset), 10),
	)
}
func (t UDim) Copy() Type {
	return t
}

////////////////

type UDim2 struct {
	X, Y UDim
}

func (UDim2) TypeString() string {
	return "UDim2"
}
func (t UDim2) String() string {
	return joinstr(
		"{",
		t.X.String(),
		"}, {",
		t.Y.String(),
		"}",
	)
}
func (t UDim2) Copy() Type {
	return t
}

////////////////

type Ray struct {
	Origin, Direction Vector3
}

func (Ray) TypeString() string {
	return "Ray"
}
func (t Ray) String() string {
	return joinstr(
		"{",
		t.Origin.String(),
		"}, {",
		t.Direction.String(),
		"}",
	)
}
func (t Ray) Copy() Type {
	return t
}

////////////////

type Faces struct {
	Right, Top, Back, Left, Bottom, Front bool
}

func (Faces) TypeString() string {
	return "Faces"
}
func (t Faces) String() string {
	s := make([]string, 6)
	if t.Front {
		s = append(s, "Front")
	}
	if t.Bottom {
		s = append(s, "Bottom")
	}
	if t.Left {
		s = append(s, "Left")
	}
	if t.Back {
		s = append(s, "Back")
	}
	if t.Top {
		s = append(s, "Top")
	}
	if t.Right {
		s = append(s, "Right")
	}

	return strings.Join(s, ", ")
}
func (t Faces) Copy() Type {
	return t
}

////////////////

type Axes struct {
	X, Y, Z bool
}

func (Axes) TypeString() string {
	return "Axes"
}
func (t Axes) String() string {
	s := make([]string, 3)
	if t.X {
		s = append(s, "X")
	}
	if t.Y {
		s = append(s, "Y")
	}
	if t.Z {
		s = append(s, "Z")
	}

	return strings.Join(s, ", ")
}
func (t Axes) Copy() Type {
	return t
}

////////////////

type BrickColor uint32

func (BrickColor) TypeString() string {
	return "BrickColor"
}
func (t BrickColor) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

//
func (bc BrickColor) Name() string {
	name, ok := brickColorNames[bc]
	if !ok {
		return brickColorNames[194]
	}

	return name
}

func (bc BrickColor) Color() Color3 {
	color, ok := brickColorColors[bc]
	if !ok {
		return brickColorColors[194]
	}

	return color
}

func (bc BrickColor) Palette() int {
	for i, n := range brickColorPalette {
		if bc == n {
			return i
		}
	}
	return -1
}
func (t BrickColor) Copy() Type {
	return t
}

////////////////

type Color3 struct {
	R, G, B float32
}

func (Color3) TypeString() string {
	return "Color3"
}
func (t Color3) String() string {
	return joinstr(
		strconv.FormatFloat(float64(t.R), 'f', -1, 32),
		", ",
		strconv.FormatFloat(float64(t.G), 'f', -1, 32),
		", ",
		strconv.FormatFloat(float64(t.B), 'f', -1, 32),
	)
}
func (t Color3) Copy() Type {
	return t
}

////////////////

type Vector2 struct {
	X, Y float32
}

func (Vector2) TypeString() string {
	return "Vector2"
}
func (t Vector2) String() string {
	return joinstr(
		strconv.FormatFloat(float64(t.X), 'f', -1, 32),
		", ",
		strconv.FormatFloat(float64(t.Y), 'f', -1, 32),
	)
}
func (t Vector2) Copy() Type {
	return t
}

////////////////

type Vector3 struct {
	X, Y, Z float32
}

func (Vector3) TypeString() string {
	return "Vector3"
}
func (t Vector3) String() string {
	return joinstr(
		strconv.FormatFloat(float64(t.X), 'f', -1, 32),
		", ",
		strconv.FormatFloat(float64(t.Y), 'f', -1, 32),
		", ",
		strconv.FormatFloat(float64(t.Z), 'f', -1, 32),
	)
}
func (t Vector3) Copy() Type {
	return t
}

////////////////

type CFrame struct {
	X, Y, Z float32
	R       [9]float32
}

func (CFrame) TypeString() string {
	return "CoordinateFrame"
}
func (t CFrame) String() string {
	s := make([]string, 12)
	s[0] = strconv.FormatFloat(float64(t.X), 'f', -1, 32)
	s[1] = strconv.FormatFloat(float64(t.Y), 'f', -1, 32)
	s[2] = strconv.FormatFloat(float64(t.Z), 'f', -1, 32)
	for i, f := range t.R {
		s[i+3] = strconv.FormatFloat(float64(f), 'f', -1, 32)
	}
	return strings.Join(s, ", ")
}
func (t CFrame) Copy() Type {
	return t
}

////////////////

type Token int32

func (Token) TypeString() string {
	return "token"
}
func (t Token) String() string {
	return strconv.FormatInt(int64(t), 10)
}
func (t Token) Copy() Type {
	return t
}

////////////////

type Reference *rbxfile.Instance

func (Reference) TypeString() string {
	return "Ref"
}
func (t Reference) String() string {
	return *rbxfile.Instance(t).Name()
}
func (t Reference) Copy() Type {
	return t
}

////////////////

type Vector3int16 struct {
	X, Y, Z int16
}

func (Vector3int16) TypeString() string {
	return "Vector3int16"
}
func (t Vector3int16) String() string {
	return joinstr(
		strconv.FormatInt(int64(t.X), 10),
		", ",
		strconv.FormatInt(int64(t.Y), 10),
		", ",
		strconv.FormatInt(int64(t.Z), 10),
	)
}
func (t Vector3int16) Copy() Type {
	return t
}

////////////////

type Vector2int16 struct {
	X, Y int16
}

func (Vector2int16) TypeString() string {
	return "Vector2int16"
}
func (t Vector2int16) String() string {
	return joinstr(
		strconv.FormatInt(int64(t.X), 10),
		", ",
		strconv.FormatInt(int64(t.Y), 10),
	)
}
func (t Vector2int16) Copy() Type {
	return t
}

////////////////

type Region3 struct {
	CFrame CFrame
	Size   Vector3
}

func (Region3) TypeString() string {
	return "Region3"
}
func (t Region3) String() string {
	return joinstr(
		t.CFrame.String(),
		"; ",
		t.Size.String(),
	)
}
func (t Region3) Copy() Type {
	return t
}

////////////////

type Region3int16 struct {
	Max, Min Vector3int16
}

func (Region3int16) TypeString() string {
	return "Region3int16"
}
func (t Region3int16) String() string {
	return joinstr(
		t.Min.String(),
		"; ",
		t.Max.String(),
	)
}
func (t Region3int16) Copy() Type {
	return t
}

////////////////
