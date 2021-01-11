package go_to_gdtype

import (
    "errors"
    "strings"
)

type GodotType uint32

const (
    GodotTypeNull GodotType = iota // unsupported

    // Atomic Types
    GodotTypeBool    GodotType = iota
    GodotTypeInteger GodotType = iota
    GodotTypeFloat   GodotType = iota
    GodotTypeString  GodotType = iota

    // math types
    GodotTypeVector2     GodotType = iota
    GodotTypeRect2       GodotType = iota
    GodotTypeVector3     GodotType = iota
    GodotTypeTransform2d GodotType = iota
    GodotTypePlane       GodotType = iota
    GodotTypeQuat        GodotType = iota
    GodotTypeAabb        GodotType = iota
    GodotTypeBasis       GodotType = iota
    GodotTypeTransform   GodotType = iota

    // misc types
    GodotTypeColor      GodotType = iota
    GodotTypeNodePath   GodotType = iota
    GodotTypeRid        GodotType = iota // unsupported
    GodotTypeObject     GodotType = iota // unsupported
    GodotTypeDictionary GodotType = iota
    GodotTypeArray      GodotType = iota

    // typed arrays
    GodotTypeRawArray     GodotType = iota
    GodotTypeIntArray     GodotType = iota
    GodotTypeRealArray    GodotType = iota
    GodotTypeStringArray  GodotType = iota
    GodotTypeVector2Array GodotType = iota
    GodotTypeVector3Array GodotType = iota
    GodotTypeColorArray   GodotType = iota

    GodotTypeMax GodotType = iota // unsupported
)

type GDTypeBool bool        // 0 for False, 1 for True
type GDTypeInteger32 uint32 // 32-bit signed integer
type GDTypeFloat float32    // IEE 754 32-Bits Float
type GDTypeString string    // UTF-8 encoded string
type GDTypeMax int32        // unsupported and unused

type GDString struct {
    Length GDTypeInteger32
    Data   GDTypeString // 4 + 4 + ... bytes
}
type GDEnTypeString GDString

type GDTypeByte uint8 // (0...255)

type GDVector2 struct {
    X GDTypeFloat
    Y GDTypeFloat
}
type GDTypeVector2 GDVector2

type GDRect2 struct {
    XCor  GDTypeFloat
    YCor  GDTypeFloat
    XSize GDTypeFloat
    YSize GDTypeFloat
}
type GDTypeRect2 GDRect2

type GDVector3 struct {
    X GDTypeFloat
    Y GDTypeFloat
    Z GDTypeFloat
}
type GDTypeVector3 GDVector3

type GDTransform2d struct {
    XXColumn GDTypeFloat
    YXColumn GDTypeFloat
    XYColumn GDTypeFloat
    YYColumn GDTypeFloat
    XOrigin  GDTypeFloat
    YOrigin  GDTypeFloat
}
type GDTypeTransform2d GDTransform2d

type GDPlane struct {
    NormalX  GDTypeFloat
    NormalY  GDTypeFloat
    NormalZ  GDTypeFloat
    Distance GDTypeFloat
}
type GDTypePlane GDPlane

type GDQuat struct {
    ImaginaryX GDTypeFloat
    ImaginaryY GDTypeFloat
    ImaginaryZ GDTypeFloat
    RealW      GDTypeFloat
}
type GDTypeQuat GDQuat

type GDAabb struct {
    XCor  GDTypeFloat
    YCor  GDTypeFloat
    ZCor  GDTypeFloat
    XSize GDTypeFloat
    YSize GDTypeFloat
    ZSize GDTypeFloat
}
type GDTypeAabb GDAabb // https://docs.godotengine.org/en/stable/classes/class_aabb.html#class-aabb

type GDBasis struct {
    XXColumn GDTypeFloat
    YXColumn GDTypeFloat
    ZXColumn GDTypeFloat

    XYColumn GDTypeFloat
    YYColumn GDTypeFloat
    ZYColumn GDTypeFloat

    XZColumn GDTypeFloat
    YZColumn GDTypeFloat
    ZZColumn GDTypeFloat
}
type GDTypeBasis GDBasis

type GDTransform struct {
    XXColumn GDTypeFloat
    YXColumn GDTypeFloat
    ZXColumn GDTypeFloat

    XYColumn GDTypeFloat
    YYColumn GDTypeFloat
    ZYColumn GDTypeFloat

    XZColumn GDTypeFloat
    YZColumn GDTypeFloat
    ZZColumn GDTypeFloat

    XOrigin GDTypeFloat
    YOrigin GDTypeFloat
    ZOrigin GDTypeFloat
}
type GDTypeTransform GDTransform

/*
GDColor need to be initialized with the A attribute with 1, because Godot put default value 1 when they not declared explicit.
Because Go has the default 0 values in structs always you use a GDColor struct, you will need to declare if the A value will be a 0..1 value
*/
type GDColor struct {
    R GDTypeFloat // Red (typically 0..1, can be above 1 for overbright colors)
    G GDTypeFloat // Green (typically 0..1, can be above 1 for overbright colors)
    B GDTypeFloat // Blue (typically 0..1, can be above 1 for overbright colors)
    A GDTypeFloat // Alpha (0..1) (default 1)
}
type GDTypeColor GDColor

type GDNodePath struct {
    Path GDTypeString // String length, or new format (val&0x80000000!=0 and NameCount=val&0x7FFFFFFF)
}
type GDTypeNodePath int32

func (g *GDNodePath) getNameCount() uint32 {
    if g.isAbsolute() {
        // remove first "/"
        path := strings.TrimLeft(string(g.Path), "/")
        return uint32(len(strings.Split(path, "/")))
    }
    return uint32(len(strings.Split(string(g.Path), "/")))
}

// verify if path start with /
func (g *GDNodePath) isAbsolute() (is bool) {
    if len(g.Path) > 0 {
        if g.Path[0] == '/' {
            is = true
        }
    }
    return
}

/*
Gets the number of resource or property names ("subnames") in the path. Each subname is
listed after a colon character (:) in the node path

For example, "Path2D/PathFollow2D/Sprite:texture:load_path" has 2 subnames.
*/
func (g *GDNodePath) getSubNameCount() uint32 {
    idx := strings.IndexRune(string(g.Path), ':')
    if idx > -1 {
        l := len(strings.Split(string(g.Path[idx+1:]), ":"))
        if l == 1 {
            if strings.Split(string(g.Path[idx+1:]), ":")[0] == "" {
                return 0
            }
        }
        return uint32(l)
    }
    return 0
}

func (g *GDNodePath) getName(idx uint32) (string, error) {
    path := g.Path
    if g.isAbsolute() {
        path = GDTypeString(strings.TrimLeft(string(g.Path), "/"))
    }
    s := strings.Split(string(path), "/")
    if len(s) > 0 {
        str := s[idx]
        if ok := strings.ContainsRune(str, ':'); ok {
            str = str[:strings.IndexRune(str, ':')]
        }
        return str, nil
    }
    return "", errors.New("index in nodePath not exists")
}

func (g *GDNodePath) getSubName(sIdx uint32) (string, error) {
    idx := strings.IndexRune(string(g.Path), ':')
    if idx > -1 {
        subNames := strings.Split(string(g.Path[idx+1:]), ":")
        if len(subNames) == 1 {
            if strings.Split(string(g.Path[idx+1:]), ":")[0] == "" {
                return "", errors.New("index in nodePath not exists")
            }
        }
        return subNames[sIdx], nil
    }
    return "", errors.New("index in nodePath not exists")
}

type GDTypeRid int32    // unsupported
type GDTypeObject int32 // unsupported

/**
Then what follows is, for amount of "elements", pairs of key and value,
one after the other, using this same format.
*/
type GDDictionary struct {
    Values map[string]interface{}
}
type GDTypeDictionary GDDictionary

/**
Then what follows is, for amount of "elements", values one after the other, using this same format.
*/
type GDArray struct {
    Values []interface{}
}
type GDTypeArray GDArray

// "PoolByteArray"
type GDRawArray struct {
    Values []byte
}
type GDTypeRawArray int32 // or "PoolByteArray"

type GDIntArray struct {
    Values []GDTypeInteger32
}
type GDTypeIntArray GDIntArray

type GDRealArray struct {
    Values []GDTypeFloat
}
type GDTypeRealArray GDRealArray

// Every string is padded to 4 bytes.
type GDStringArray struct {
    Values []GDTypeString
}
type GDTypeStringArray GDStringArray

type GDVector2Array struct {
    Values []GDVector2
}
type GDTypeVector2Array GDVector2Array

type GDVector3Array struct {
    Values []GDVector3
}
type GDTypeVector3Array GDVector3Array

type GDColorArray struct {
    Colors []GDColor
}
type GDTypeColorArray GDColorArray
