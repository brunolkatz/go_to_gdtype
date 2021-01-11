Project: go_to_gdtype
------------------------

This package is used to binary serialize Go variables to GDScript (Godot). Can be used to transport tcp packets and unserialize with ```get_var``` function
in Godot GDScript native function.

Supported variables:

Type | Value | Struct |
-----|-------|--------|
0  | null    | **Serialize unsupported**
1  | bool | GDTypeBool
2  | integer | GDTypeInteger32
3  | float | GDTypeFloat
4  | string | GDTypeString / string
5  | vector2 | GDString
6  | rect2 | GDRect2
7  | vector3 | GDVector3
8  | transform2d | GDTransform2d
9  | plane | GDPlane
10 | quat | GDQuat
11 | aabb | GDAabb
12 | basis | GDBasis
13 | transform | GDTransform
14 | color | GDColor
15 | node path | GDNodePath
16 | rid | GDTypeRid - **Serialize unsupported**
17 | object | GDTypeObject - **Serialize unsupported**
18 | dictionary | GDDictionary
19 | array | GDArray
20 | raw array | GDRawArray
21 | int array | GDIntArray
22 | real array | GDRealArray
23 | string array | GDStringArray
24 | vector2 array | GDVector2Array
25 | vector3 array | GDVector3Array
26 | color array | GDColorArray
27 | max | **Serialize unsupported**
