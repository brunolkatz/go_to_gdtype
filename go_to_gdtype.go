package go_to_gdtype

import (
    "bytes"
    "encoding/binary"
    "errors"
)

/**
Godot serialize data like: https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html
*/

/**
The packet is designed to be always padded to 4 bytes. All values are little-endian-encoded.
All packets have a 4-byte header representing an integer

TODO: Make the 64 compiled format too
*/

var (
    ErrorGDTypeNotDefined = errors.New("GDType not defined")
)

/**
v is the value to convert
gdType is the returned type to buffer
 */
func Var2GodotBytes(value interface{}) ([]byte, error){
    // Verify type of interface

    if _v, ok := value.(GDTypeBool); ok {
        b, err := toBool(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDTypeInteger32); ok {
        b, err := toInteger32(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDTypeFloat); ok {
        b, err := toFloat(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(string); ok {
        b, err := toString(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDVector2); ok {
        b, err := toVector2(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDRect2); ok {
        b, err := toRect2(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDVector3); ok {
        b, err := toVector3(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDTransform2d); ok {
        b, err := toTransform2d(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDPlane); ok {
        b, err := toPlane(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDQuat); ok {
        b, err := toQuat(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDAabb); ok {
        b, err := toAABB(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDBasis); ok {
        b, err := toBasis(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDTransform); ok {
        b, err := toTransform(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDColor); ok {
        b, err := toColor(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDNodePath); ok {
        b, err := toNodePath(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDDictionary); ok {
        b, err := toDictionary(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDArray); ok {
        b, err := toArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDRawArray); ok {
        b, err := toRawArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDIntArray); ok {
        b, err := toIntArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDRealArray); ok {
        b, err := toRealArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDStringArray); ok {
        b, err := toStringArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDVector2Array); ok {
        b, err := toVector2Array(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDVector3Array); ok {
        b, err := toVector3Array(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    if _v, ok := value.(GDColorArray); ok {
        b, err := toColorArray(_v)
        if err != nil {
            return nil, err
        }
        return b, nil
    }

    return nil, ErrorGDTypeNotDefined
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#bool
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeBool)
  4   |  4  | Integer | 0 for False, 1 for True
*/
func toBool(v GDTypeBool) ([]byte, error){

    maxLen := int32(4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeBool); err != nil {
        return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v); err != nil {
        return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#int
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeInteger)
  4   |  8  | Integer | 64-bit signed integer
*/
func toInteger32(v GDTypeInteger32) ([]byte, error){

    maxLen := int32(4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeInteger); err != nil {
        return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v); err != nil {
        return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#float
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeFloat)
  4   |  4  | Float   | IEE 754 32-Bits Float
*/
func toFloat(v GDTypeFloat) ([]byte, error){

    maxLen := int32(4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeBool); err != nil {
        return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v); err != nil {
        return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#string
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeString)
  4	  |  4	| Integer |	String length (in bytes)
  8	  |  X	| Bytes	  | UTF-8 encoded string
 */
func toString(v string) ([]byte, error){

    l := int32(len([]byte(v)))
    maxLen := int32(4+4+len([]byte(v)))
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeString); err != nil {
       return nil, err
    }
    if err := binary.Write(headDataBuf, binary.LittleEndian, l); err != nil {
       return nil, err
    }
    _, _ = headDataBuf.WriteString(v)

    rLen := 4 + l
    for  {
        if rLen%4 <= 0 {
            break
        }
        rLen++
        headDataBuf.WriteByte(byte(0))
    }
    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#vector2
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeString)
  4	  |  4  | Float	  | X coordinate
  8	  |  4  | Float	  | Y coordinate
*/
func toVector2(v GDVector2) ([]byte, error){

    maxLen := int32(4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeVector2); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.X); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.Y); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#rect2
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeRect2)
  4	  |  4	| Float   | X coordinate
  8	  |  4	| Float   | Y coordinate
  12  |  4	| Float   | X size (width)
  16  |  4	| Float   | Y size (height)
 */
func toRect2(v GDRect2) ([]byte, error){

    maxLen := int32(4+4+4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeRect2); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XCor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YCor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XSize); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YSize); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#vector3
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeVector3)
  4   |  4  | Float   | X coordinate
  8   |  4  | Float   | Y coordinate
  12  |  4  | Float   | Z coordinate
 */
func toVector3(v GDVector3) ([]byte, error){

    maxLen := int32(4+4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeVector3); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.X); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.Y); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.Z); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#transform2d
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeTransform2d)
  4   |  4  | Float   | The X component of the X column vector, accessed via [0][0]
  8   |  4  | Float   | The Y component of the X column vector, accessed via [0][1]
  12  |  4  | Float   | The X component of the Y column vector, accessed via [1][0]
  16  |  4  | Float   | The Y component of the Y column vector, accessed via [1][1]
  20  |  4  | Float   | The X component of the origin vector, accessed via [2][0]
  24  |  4  | Float   | The Y component of the origin vector, accessed via [2][1]
 */
func toTransform2d(v GDTransform2d) ([]byte, error){

    maxLen := int32(4+4+4+4+4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeTransform2d); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XOrigin); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YOrigin); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#plane
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypePlane)
  4   |  4  | Float   | Normal X
  8   |  4  | Float   | Normal Y
  12  |  4  | Float   | Normal Z
  16  |  4  | Float   | Distance
 */
func toPlane(v GDPlane) ([]byte, error){

    maxLen := int32(4+4+4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypePlane); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.NormalX); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.NormalY); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.NormalZ); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.Distance); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#quat
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeQuat)
  4   |  4  | Float   | Imaginary X
  8   |  4  | Float   | Imaginary Y
  12  |  4  | Float   | Imaginary Z
  16  |  4  | Float   | Real W
 */
func toQuat(v GDQuat) ([]byte, error){

    maxLen := int32(4+4+4+4+4)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeQuat); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ImaginaryX); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ImaginaryY); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ImaginaryZ); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.RealW); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#aabb
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeAabb)
  4   |  4  | Float   | X coordinate
  8   |  4  | Float   | Y coordinate
  12  |  4  | Float   | Z coordinate
  16  |  4  | Float   | X size
  20  |  4  | Float   | Y size
  24  |  4  | Float   | Z size
 */
func toAABB(v GDAabb) ([]byte, error){

    maxLen := int32(4*6)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeAabb); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XCor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YCor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZCor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XSize); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YSize); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZSize); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#basis
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeBasis)
  4   |  4  | Float   | The X component of the X column vector, accessed via [0][0]
  8   |  4  | Float   | The Y component of the X column vector, accessed via [0][1]
  12  |  4  | Float   | The Z component of the X column vector, accessed via [0][2]
  16  |  4  | Float   | The X component of the Y column vector, accessed via [1][0]
  20  |  4  | Float   | The Y component of the Y column vector, accessed via [1][1]
  24  |  4  | Float   | The Z component of the Y column vector, accessed via [1][2]
  28  |  4  | Float   | The X component of the Z column vector, accessed via [2][0]
  32  |  4  | Float   | The Y component of the Z column vector, accessed via [2][1]
  36  |  4  | Float   | The Z component of the Z column vector, accessed via [2][2]
 */
func toBasis(v GDBasis) ([]byte, error){

    maxLen := int32(4*9)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeBasis); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XZColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YZColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZZColumn); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#transform
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeTransform)
  4   |  4  | Float   | The X component of the X column vector, accessed via [0][0]
  8   |  4  | Float   | The Y component of the X column vector, accessed via [0][1]
  12  |  4  | Float   | The Z component of the X column vector, accessed via [0][2]
  16  |  4  | Float   | The X component of the Y column vector, accessed via [1][0]
  20  |  4  | Float   | The Y component of the Y column vector, accessed via [1][1]
  24  |  4  | Float   | The Z component of the Y column vector, accessed via [1][2]
  28  |  4  | Float   | The X component of the Z column vector, accessed via [2][0]
  32  |  4  | Float   | The Y component of the Z column vector, accessed via [2][1]
  36  |  4  | Float   | The Z component of the Z column vector, accessed via [2][2]
  40  |  4  | Float   | The X component of the origin vector, accessed via [3][0]
  44  |  4  | Float   | The Y component of the origin vector, accessed via [3][1]
  48  |  4  | Float   | The Z component of the origin vector, accessed via [3][2]
 */
func toTransform(v GDTransform) ([]byte, error){

    maxLen := int32(4*12)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeTransform); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XZColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YZColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZXColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZYColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZZColumn); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.XOrigin); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.YOrigin); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.ZOrigin); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#transform
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeTransform)
  4   |  4  | Float   | Red (typically 0..1, can be above 1 for overbright colors) (default 0)
  8   |  4  | Float   | Green (typically 0..1, can be above 1 for overbright colors) (default 0)
  12  |  4  | Float   | Blue (typically 0..1, can be above 1 for overbright colors) (default 0)
  16  |  4  | Float   | Alpha (0..1) (default 1)
 */
func toColor(v GDColor) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeColor); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.R); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.G); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.B); err != nil {
       return nil, err
    }

    if err := binary.Write(headDataBuf, binary.LittleEndian, v.A); err != nil {
       return nil, err
    }

    return headDataBuf.Bytes(), nil
}

/**
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#nodepath
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeNodePath)
  4   |  4  | Integer | String length, or new format (val&0x80000000!=0 and NameCount=val&0x7FFFFFFF)

A NodePath is composed of a list of slash-separated node names (like a filesystem path) and an optional colon-separated list of "subnames" which can be resources or properties.

Some examples of NodePaths include the following:

# No leading slash means it is relative to the current node.
@"A" # Immediate child A
@"A/B" # A's child B
@"." # The current node.
@".." # The parent node.
@"../C" # A sibling node C.
# A leading slash means it is absolute from the SceneTree.
@"/root" # Equivalent to get_tree().get_root().
@"/root/Main" # If your main scene's root node were named "Main".
@"/root/MyAutoload" # If you have an autoloaded node or scene.
 */
func toNodePath(v GDNodePath) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeNodePath); err != nil {
       return nil, err
    }

    sl := v.getNameCount() | 0x80000000
    if err := binary.Write(headDataBuf, binary.LittleEndian, sl); err != nil {
       return nil, err
    }
    if err := binary.Write(headDataBuf, binary.LittleEndian, v.getSubNameCount()); err != nil {
       return nil, err
    }

    npFlags := uint32(0)
    if v.isAbsolute() {
        npFlags |= 1
    }
    if err := binary.Write(headDataBuf, binary.LittleEndian, npFlags); err != nil {
        return nil, err
    }

    total := v.getNameCount() + v.getSubNameCount()
    for i := uint32(0); i < total; i++ {

        str := ""
        if i < v.getNameCount() {
            if s, err := v.getName(i); err != nil {
                return nil, err
            }else{
                str = s
            }
        }else{
            if s, err := v.getSubName(i - v.getNameCount()); err != nil {
                return nil, err
            }else{
                str = s
            }
        }

        l := int32(len([]byte(str)))
        if err := binary.Write(headDataBuf, binary.LittleEndian, l); err != nil {
            return nil, err
        }

        _, _ = headDataBuf.WriteString(str)
        if l%4 > 0 {
            of := 4-l%4
            if of > 0 {
                for i  :=0; i < int(of); i++ {
                    headDataBuf.WriteByte(byte(0))
                }
            }
        }
    }
    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#dictionary
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeDictionary)
  4   |  4  | Integer | val&0x7FFFFFFF = elements, val&0x80000000 = shared (bool)
*/
func toDictionary(v GDDictionary) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeDictionary); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for key, variant := range v.Values {

        b, err := Var2GodotBytes(key)
        if err != nil {
            return nil, err
        }
        headDataBuf.Write(b)

        b, err = Var2GodotBytes(variant)
        if err != nil {
            return nil, err
        }
        headDataBuf.Write(b)
    }


    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#array
Offset|	Len	| Type	  | Description
  0   |  4  | Integer | type of variable (GodotTypeArray)
  4   |  4  | Integer | val&0x7FFFFFFF = elements, val&0x80000000 = shared (bool)
*/
func toArray(v GDArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, variant := range v.Values {

        b, err := Var2GodotBytes(variant)
        if err != nil {
            return nil, err
        }
        headDataBuf.Write(b)
    }


    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolbytearray
Offset          | Len | Type    | Description
  0             |  4  | Integer | type of variable (GodotTypeRawArray) (is the same of poolbytearray)
  8..8+length*4 |  4  | Integer	| Byte (0..255)
*/
func toRawArray(v GDRawArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeRawArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    headDataBuf.Write(v.Values)

    rLen := 4 + uint32(len(v.Values))
    for  {
        if rLen%4 <= 0 {
            break
        }
        rLen++
        headDataBuf.WriteByte(byte(0))
    }
    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolintarray
Offset          | Len | Type    | Description
  0             |  4  | Integer | type of variable (GodotTypeIntArray)
  8..8+length*4 |  4  | Integer	| 32-bit signed integer
 */
func toIntArray(v GDIntArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeIntArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, iV := range v.Values {
        if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(iV)); err != nil {
            return nil, err
        }
    }

    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolrealarray
Offset          | Len | Type    | Description
  0             |  4  | Integer | type of variable (GodotTypeRealArray)
  8..8+length*4 |  4  | Integer	| 32-bits IEEE 754 float
 */
func toRealArray(v GDRealArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeRealArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, iV := range v.Values {
        if err := binary.Write(headDataBuf, binary.LittleEndian, iV); err != nil {
            return nil, err
        }
    }

    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolstringarray
Offset | Len | Type    | Description
  0    |  4  | Integer | type of variable (GodotTypeRealArray)
  4    |  4  | Integer | Array length (Strings)
 */
func toStringArray(v GDStringArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeStringArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, utf8 := range v.Values {

        l := uint32(len([]byte(utf8)))
        // the +1 in length of the string are in marshalls.cpp line 1413: "encode_uint32(utf8.length() + 1, buf);"
        if err := binary.Write(headDataBuf, binary.LittleEndian, l+1); err != nil {
            return nil, err
        }

        headDataBuf.WriteString(string(utf8))

        rLen := 4 + l
        if rLen%4 <= 0 {
            headDataBuf.Write([]byte{0,0,0,0})
        }else{
            for  {
                if rLen%4 <= 0 {
                    break
                }
                rLen++
                headDataBuf.WriteByte(byte(0))
            }
        }
    }

    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolvector2array
Offset           | Len | Type    | Description
  0              |  4  | Integer | type of variable (GodotTypeVector2Array)
  4              |  4  | Integer | Array length
  8..8+length*8  |  4  | Float   | X coordinate
  8..12+length*8 |  4  | Float   | Y coordinate
 */
func toVector2Array(v GDVector2Array) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeVector2Array); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, vector2 := range v.Values {
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.X); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.Y); err != nil {
            return nil, err
        }
    }

    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolvector3array
Offset           | Len | Type    | Description
  0              |  4  | Integer | type of variable (GodotTypeVector3Array)
  4              |  4  | Integer | Array length
  8..8+length*8  |  4  | Float   | X coordinate
  8..12+length*8 |  4  | Float   | Y coordinate
  8..12+length*8 |  4  | Float   | Z coordinate
 */
func toVector3Array(v GDVector3Array) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeVector3Array); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Values))); err != nil {
        return nil, err
    }

    for _, vector2 := range v.Values {
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.X); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.Y); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.Z); err != nil {
            return nil, err
        }
    }

    return headDataBuf.Bytes(), nil
}

/*
https://docs.godotengine.org/en/stable/tutorials/misc/binary_serialization_api.html#poolcolorarray
Offset           | Len | Type    | Description
  0              |  4  | Integer | type of variable (GodotTypeColorArray)
  4              |  4  | Integer | Array length
  8..8+length*8  |  4  | Float   | X coordinate
  8..12+length*8 |  4  | Float   | Y coordinate
  8..12+length*8 |  4  | Float   | Z coordinate
 */
func toColorArray(v GDColorArray) ([]byte, error){

    maxLen := int32(4*5)
    headDataBuf := bytes.NewBuffer(make([]byte,0, maxLen))

    if err := binary.Write(headDataBuf, binary.LittleEndian, GodotTypeColorArray); err != nil {
        return nil, err
    }

    // Add size of dictionary
    if err := binary.Write(headDataBuf, binary.LittleEndian, uint32(len(v.Colors))); err != nil {
        return nil, err
    }

    for _, vector2 := range v.Colors {
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.R); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.G); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.B); err != nil {
            return nil, err
        }
        if err := binary.Write(headDataBuf, binary.LittleEndian, vector2.A); err != nil {
            return nil, err
        }
    }

    return headDataBuf.Bytes(), nil
}