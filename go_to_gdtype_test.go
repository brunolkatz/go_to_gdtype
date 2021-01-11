package go_to_gdtype

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "testing"
)

/**
String test conversion
 */
func TestVar2GodotBytesString(t *testing.T) {

    vTest := []struct{
        Test string
        Expect []byte
    }{
        {
            Test:   "testeteste",
            Expect: []byte{4, 0, 0, 0, 10, 0, 0, 0, 116, 101, 115, 116, 101, 116, 101, 115, 116, 101, 0, 0},
        },
        {
            Test:   "teste",
            Expect: []byte{4, 0, 0, 0, 5, 0, 0, 0, 116, 101, 115, 116, 101, 0, 0, 0},
        },
        {
            Test:   "",
            Expect: []byte{4, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   "nil",
            Expect: []byte{4, 0, 0, 0, 3, 0, 0, 0, 110, 105, 108, 0},
        },
        {
            Test:   "a",
            Expect: []byte{4, 0, 0, 0, 1, 0, 0, 0, 97, 0, 0, 0},
        },
        {
            Test:   "cabeça",
            Expect: []byte{4, 0, 0, 0, 7, 0, 0, 0, 99, 97, 98, 101, 195, 167, 97, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + v.Test)
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesVector2(t *testing.T) {

    vTest := []struct{
        Test GDVector2
        Expect []byte
    }{
        {
            Test:   GDVector2{ X: 10, Y: 10 },
            Expect: []byte{5, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65},
        },
        {
            Test:   GDVector2{ X: 10.123, Y: 10.123 },
            Expect: []byte{5, 0, 0, 0, 207, 247, 33, 65, 207, 247, 33, 65},
        },
        {
            Test:   GDVector2{ X: 999.999, Y: 999.999 },
            Expect: []byte{5, 0, 0, 0, 240, 255, 121, 68, 240, 255, 121, 68},
        },
        {
            Test:   GDVector2{ X: 0, Y: 0 },
            Expect: []byte{5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesRect2(t *testing.T) {

    vTest := []struct{
        Test GDRect2
        Expect []byte
    }{
        {
            Test:   GDRect2{ XCor: 10, YCor: 10, YSize: 100, XSize: 100},
            Expect: []byte{6, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 200, 66, 0, 0, 200, 66},
        },
        {
            Test:   GDRect2{ XCor: 999.999, YCor: 999.999, YSize: 999.999, XSize: 999.999},
            Expect: []byte{6, 0, 0, 0, 240, 255, 121, 68, 240, 255, 121, 68, 240, 255, 121, 68, 240, 255, 121, 68},
        },
        {
            Test:   GDRect2{ XCor: 0, YCor: 0, YSize: 0, XSize: 0},
            Expect: []byte{6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDRect2{},
            Expect: []byte{6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesVector3(t *testing.T) {

    vTest := []struct{
        Test GDVector3
        Expect []byte
    }{
        {
            Test:   GDVector3{ X: 10, Y: 10, Z: 10},
            Expect: []byte{7, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65},
        },
        {
            Test:   GDVector3{ X: 999.999, Y: 999.999, Z: 999.999},
            Expect: []byte{7, 0, 0, 0, 240, 255, 121, 68, 240, 255, 121, 68, 240, 255, 121, 68},
        },
        {
            Test:   GDVector3{ X: 0, Y: 0, Z: 0},
            Expect: []byte{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDVector3{},
            Expect: []byte{7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesTransform2d(t *testing.T) {

    vTest := []struct{
        Test GDTransform2d
        Expect []byte
    }{
        {
            Test:   GDTransform2d{
                XXColumn: 10,
                YXColumn: 10,
                XYColumn: 10,
                YYColumn: 10,
                XOrigin:  10,
                YOrigin:  10,
            },
            Expect: []byte{8, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65},
        },
        {
            Test:   GDTransform2d{
                XXColumn: 999.876,
                YXColumn: 999.876,
                XYColumn: 999.876,
                YYColumn: 999.876,
                XOrigin:  999.876,
                YOrigin:  999.876,
            },
            Expect: []byte{8, 0, 0, 0, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68},
        },
        {
            Test:   GDTransform2d{
                XXColumn: 0,
                YXColumn: 0,
                XYColumn: 0,
                YYColumn: 0,
                XOrigin:  0,
                YOrigin:  0,
            },
            Expect: []byte{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDTransform2d{},
            Expect: []byte{8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesPlane(t *testing.T) {

    vTest := []struct{
        Test GDPlane
        Expect []byte
    }{
        {
            Test:   GDPlane{
                NormalX: 10,
                NormalY: 10,
                NormalZ: 10,
                Distance: 10,
            },
            Expect: []byte{9, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65},
        },
        {
            Test:   GDPlane{
                NormalX: 999.876,
                NormalY: 999.876,
                NormalZ: 999.876,
                Distance: 999.876,
            },
            Expect: []byte{9, 0, 0, 0, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68},
        },
        {
            Test:   GDPlane{
                NormalX: 0,
                NormalY: 0,
                NormalZ: 0,
                Distance: 0,
            },
            Expect: []byte{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDPlane{},
            Expect: []byte{9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesQuat(t *testing.T) {

    vTest := []struct{
        Test GDQuat
        Expect []byte
    }{
        {
            Test:   GDQuat{
                ImaginaryX: 10,
                ImaginaryY: 10,
                ImaginaryZ: 10,
                RealW: 10,
            },
            Expect: []byte{10, 0, 0, 0, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65, 0, 0, 32, 65},
        },
        {
            Test:   GDQuat{
                ImaginaryX: 999.876,
                ImaginaryY: 999.876,
                ImaginaryZ: 999.876,
                RealW: 999.876,
            },
            Expect: []byte{10, 0, 0, 0, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68, 16, 248, 121, 68},
        },
        {
            Test:   GDQuat{
                ImaginaryX: 0,
                ImaginaryY: 0,
                ImaginaryZ: 0,
                RealW: 0,
            },
            Expect: []byte{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDQuat{},
            Expect: []byte{10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesBasis(t *testing.T) {

    vTest := []struct{
        Test GDBasis
        Expect []byte
    }{
        {
            Test:   GDBasis{
                XXColumn: 77,
                YXColumn: 77,
                ZXColumn: 77,
                XYColumn: 88,
                YYColumn: 88,
                ZYColumn: 88,
                XZColumn: 99,
                YZColumn: 99,
                ZZColumn: 99,
            },
            Expect: []byte{12, 0, 0, 0, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66},
        },
        {
            Test:   GDBasis{
                XXColumn: 999.876,
                YXColumn: 999.876,
                ZXColumn: 999.876,
                XYColumn: 999.777,
                YYColumn: 999.777,
                ZYColumn: 999.777,
                XZColumn: 999.555,
                YZColumn: 999.555,
                ZZColumn: 999.555,
            },
            Expect: []byte{12, 0, 0, 0, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68},
        },
        {
            Test:   GDBasis{
                XXColumn: 0,
                YXColumn: 0,
                ZXColumn: 0,
                XYColumn: 0,
                YYColumn: 0,
                ZYColumn: 0,
                XZColumn: 0,
                YZColumn: 0,
                ZZColumn: 0,
            },
            Expect: []byte{12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDBasis{},
            Expect: []byte{12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesTrasnform(t *testing.T) {

    vTest := []struct{
        Test GDTransform
        Expect []byte
    }{
        {
            Test:   GDTransform{
                XXColumn: 77,
                YXColumn: 77,
                ZXColumn: 77,
                XYColumn: 88,
                YYColumn: 88,
                ZYColumn: 88,
                XZColumn: 99,
                YZColumn: 99,
                ZZColumn: 99,
                XOrigin: 10,
                YOrigin: 20,
                ZOrigin: 30,
            },
            Expect: []byte{13, 0, 0, 0, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66, 0, 0, 154, 66, 0, 0, 176, 66, 0, 0, 198, 66, 0, 0, 32, 65, 0, 0, 160, 65, 0, 0, 240, 65},
        },
        {
            Test:   GDTransform{
                XXColumn: 999.876,
                YXColumn: 999.876,
                ZXColumn: 999.876,
                XYColumn: 999.777,
                YYColumn: 999.777,
                ZYColumn: 999.777,
                XZColumn: 999.555,
                YZColumn: 999.555,
                ZZColumn: 999.555,
                XOrigin: 10,
                YOrigin: 20,
                ZOrigin: 30,
            },
            Expect: []byte{13, 0, 0, 0, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68, 16, 248, 121, 68, 186, 241, 121, 68, 133, 227, 121, 68, 0, 0, 32, 65, 0, 0, 160, 65, 0, 0, 240, 65},
        },
        {
            Test:   GDTransform{
                XXColumn: 0,
                YXColumn: 0,
                ZXColumn: 0,
                XYColumn: 0,
                YYColumn: 0,
                ZYColumn: 0,
                XZColumn: 0,
                YZColumn: 0,
                ZZColumn: 0,
                XOrigin: 0,
                YOrigin: 0,
                ZOrigin: 0,
            },
            Expect: []byte{13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDTransform{},
            Expect: []byte{13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotBytesColor(t *testing.T) {

    vTest := []struct {
        Test   GDColor
        Expect []byte
    }{
        {
            Test: GDColor{
                R: 1,
                G: 0.5,
                B: 0.6,
                A: 1,
            },
            Expect: []byte{14, 0, 0, 0, 0, 0, 128, 63, 0, 0, 0, 63, 154, 153, 25, 63, 0, 0, 128, 63},
        },
        {
            Test: GDColor{
                R: 1,
                G: 0.5,
                B: 0.6,
                A: 1,
            },
            Expect: []byte{14, 0, 0, 0, 0, 0, 128, 63, 0, 0, 0, 63, 154, 153, 25, 63, 0, 0, 128, 63},
        },
        {
            Test: GDColor{
                R: 1,
                G: 0.5,
                B: 0.6,
                A: 0.2,
            },
            Expect: []byte{14, 0, 0, 0, 0, 0, 128, 63, 0, 0, 0, 63, 154, 153, 25, 63, 205, 204, 76, 62},
        },
        {
            Test: GDColor{
                R: 0,
                G: 0,
                B: 0,
                A: 1,
            },
            Expect: []byte{14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 63},
        },
        {
            Test: GDColor{
                R: 0,
                G: 0,
                B: 0,
                A: 0,
            },
            Expect: []byte{14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
        {
            Test:   GDColor{},
            Expect: []byte{14, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotNodePath(t *testing.T) {

    vTest := []struct {
        Test   GDNodePath
        Expect []byte
    }{
        {
           Test: GDNodePath{
               "/root/Main",
           },
           Expect: []byte{15, 0, 0, 0, 2, 0, 0, 128, 0, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 114, 111, 111, 116, 4, 0, 0, 0, 77, 97, 105, 110},
        },
        {
           Test: GDNodePath{
               "/root:load_path",
           },
           Expect: []byte{15, 0, 0, 0, 1, 0, 0, 128, 1, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 114, 111, 111, 116, 9, 0, 0, 0, 108, 111, 97, 100, 95, 112, 97, 116, 104, 0, 0, 0},
        },
        {
           Test: GDNodePath{
               "A/B",
           },
           Expect: []byte{15, 0, 0, 0, 2, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 65, 0, 0, 0, 1, 0, 0, 0, 66, 0, 0, 0},
        },
        {
         Test: GDNodePath{
             "Path2D/PathFollow2D/Sprite:position:x",
         },
         Expect: []byte{15, 0, 0, 0, 3, 0, 0, 128, 2, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 80, 97, 116, 104, 50, 68, 0, 0, 12, 0, 0, 0, 80, 97, 116, 104, 70, 111, 108, 108, 111, 119, 50, 68, 6, 0, 0, 0, 83, 112, 114, 105, 116, 101, 0, 0, 8, 0, 0, 0, 112, 111, 115, 105, 116, 105, 111, 110, 1, 0, 0, 0, 120, 0, 0, 0},
        },
        {
         Test: GDNodePath{
             "Path2D/PathFollow2D/Sprite:y:x",
         },
         Expect: []byte{15, 0, 0, 0, 3, 0, 0, 128, 2, 0, 0, 0, 0, 0, 0, 0, 6, 0, 0, 0, 80, 97, 116, 104, 50, 68, 0, 0, 12, 0, 0, 0, 80, 97, 116, 104, 70, 111, 108, 108, 111, 119, 50, 68, 6, 0, 0, 0, 83, 112, 114, 105, 116, 101, 0, 0, 1, 0, 0, 0, 121, 0, 0, 0, 1, 0, 0, 0, 120, 0, 0, 0},
        },
        {
         Test: GDNodePath{
             "/root/MyAutoload",
         },
         Expect: []byte{15, 0, 0, 0, 2, 0, 0, 128, 0, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 114, 111, 111, 116, 10, 0, 0, 0, 77, 121, 65, 117, 116, 111, 108, 111, 97, 100, 0, 0},
        },
        {
          Test: GDNodePath{
              "../C",
          },
          Expect: []byte{15, 0, 0, 0, 2, 0, 0, 128, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 46, 46, 0, 0, 1, 0, 0, 0, 67, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotDictionary(t *testing.T) {

    vTest := []struct {
        Test   GDDictionary
        Expect []byte
    }{
        {
           Test: GDDictionary{
               Values: map[string]interface{}{
                   "casa": GDTypeInteger32(1),
               },
           },
           Expect: []byte{18, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 4, 0, 0, 0, 99, 97, 115, 97, 2, 0, 0, 0, 1, 0, 0, 0},
        },
        {
           Test: GDDictionary{
               Values: map[string]interface{}{
                   "casa": GDTypeInteger32(1),
                   "teste2": GDVector2{
                       X: 10,
                       Y: 11,
                   },
               },
           },
           Expect: []byte{18, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 4, 0, 0, 0, 99, 97, 115, 97, 2, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 6, 0, 0, 0, 116, 101, 115, 116, 101, 50, 0, 0, 5, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65},
        },
        {
           Test: GDDictionary{
               Values: map[string]interface{}{
                   "casa": GDDictionary{Values: map[string]interface{}{
                       "teste": GDVector2{
                           X: 10,
                           Y: 11,
                       },
                   }},
               },
           },
           Expect: []byte{18, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 4, 0, 0, 0, 99, 97, 115, 97, 18, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 5, 0, 0, 0, 116, 101, 115, 116, 101, 0, 0, 0, 5, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotArray(t *testing.T) {

    vTest := []struct {
        Test   GDArray
        Expect []byte
    }{
        {
           Test: GDArray{
               Values: []interface{}{
                   "teste", "teste2",
               },
           },
           Expect: []byte{19, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 5, 0, 0, 0, 116, 101, 115, 116, 101, 0, 0, 0, 4, 0, 0, 0, 6, 0, 0, 0, 116, 101, 115, 116, 101, 50, 0, 0},
        },
        {
           Test: GDArray{
               Values: []interface{}{
                   "teste", GDVector2{
                       X: 10,
                       Y: 11,
                   },
               },
           },
           Expect: []byte{19, 0, 0, 0, 2, 0, 0, 0, 4, 0, 0, 0, 5, 0, 0, 0, 116, 101, 115, 116, 101, 0, 0, 0, 5, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotRawArray(t *testing.T) {

    vTest := []struct {
        Test   GDRawArray
        Expect []byte
    }{
        {
           Test: GDRawArray{
               Values: []byte{11, 46, 255},
           },
           Expect: []byte{20, 0, 0, 0, 3, 0, 0, 0, 11, 46, 255, 0},
        },
        {
           Test: GDRawArray{
               Values: []byte{},
           },
           Expect: []byte{20, 0, 0, 0, 0, 0, 0, 0},
        },
        {
           Test: GDRawArray{},
           Expect: []byte{20, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotIntArray(t *testing.T) {

    vTest := []struct {
        Test   GDIntArray
        Expect []byte
    }{
        {
           Test: GDIntArray{
               Values: []GDTypeInteger32{1,2,3},
           },
           Expect: []byte{21, 0, 0, 0, 3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3, 0, 0, 0},
        },
        {
           Test: GDIntArray{
               Values: []GDTypeInteger32{987,123,999},
           },
           Expect: []byte{21, 0, 0, 0, 3, 0, 0, 0, 219, 3, 0, 0, 123, 0, 0, 0, 231, 3, 0, 0},
        },
        {
           Test: GDIntArray{},
           Expect: []byte{21, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotRealArray(t *testing.T) {

    vTest := []struct {
        Test   GDRealArray
        Expect []byte
    }{
        {
           Test: GDRealArray{
               Values: []GDTypeFloat{1.2,2.3,3.4},
           },
           Expect: []byte{22, 0, 0, 0, 3, 0, 0, 0, 154, 153, 153, 63, 51, 51, 19, 64, 154, 153, 89, 64},
        },
        {
           Test: GDRealArray{
               Values: []GDTypeFloat{987.9999,123.444,6534234.5445},
           },
           Expect: []byte{22, 0, 0, 0, 3, 0, 0, 0, 254, 255, 118, 68, 84, 227, 246, 66, 181, 104, 199, 74},
        },
        {
           Test: GDRealArray{},
           Expect: []byte{22, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotStringArray(t *testing.T) {

    vTest := []struct {
        Test   GDStringArray
        Expect []byte
    }{
        {
           Test: GDStringArray{
               Values: []GDTypeString{"teste", "teste2"},
           },
           Expect: []byte{
               23, 0, 0, 0, 2, 0, 0, 0, 6, 0, 0, 0, 116, 101, 115, 116, 101, 0, 0, 0, 7, 0, 0, 0, 116, 101, 115, 116, 101, 50, 0, 0},
        },
        {
           Test: GDStringArray{
               Values: []GDTypeString{"asdf", "asdf1231asdf13241"},
           },
           Expect: []byte{23, 0, 0, 0, 2, 0, 0, 0, 5, 0, 0, 0, 97, 115, 100, 102, 0, 0, 0, 0, 18, 0, 0, 0, 97, 115, 100, 102, 49, 50, 51, 49, 97, 115, 100, 102, 49, 51, 50, 52, 49, 0, 0, 0},
        },
        {
           Test: GDStringArray{
               Values: []GDTypeString{"asdf", "asdf1231asdf13241 cabeça de vento \n"},
           },
           Expect: []byte{23, 0, 0, 0, 2, 0, 0, 0, 5, 0, 0, 0, 97, 115, 100, 102, 0, 0, 0, 0, 37, 0, 0, 0, 97, 115, 100, 102, 49, 50, 51, 49, 97, 115, 100, 102, 49, 51, 50, 52, 49, 32, 99, 97, 98, 101, 195, 167, 97, 32, 100, 101, 32, 118, 101, 110, 116, 111, 32, 10, 0, 0, 0, 0},
        },
        {
           Test: GDStringArray{
               Values: []GDTypeString{"asdf",`
	minha nossa
	quero
	ver
	se
	isso
	funciona
	agora
	`},
           },
           Expect: []byte{23, 0, 0, 0, 2, 0, 0, 0, 5, 0, 0, 0, 97, 115, 100, 102, 0, 0, 0, 0, 55, 0, 0, 0, 10, 9, 109, 105, 110, 104, 97, 32, 110, 111, 115, 115, 97, 10, 9, 113, 117, 101, 114, 111, 10, 9, 118, 101, 114, 10, 9, 115, 101, 10, 9, 105, 115, 115, 111, 10, 9, 102, 117, 110, 99, 105, 111, 110, 97, 10, 9, 97, 103, 111, 114, 97, 10, 9, 0, 0},
        },
        {
           Test: GDStringArray{},
           Expect: []byte{23, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotVector2Array(t *testing.T) {

    vTest := []struct {
        Test   GDVector2Array
        Expect []byte
    }{
        {
            Test: GDVector2Array{
                Values: []GDVector2{
                    {
                        X: 10,
                        Y: 11,
                    },
                },
            },
            Expect: []byte{24, 0, 0, 0, 1, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65},
        },
        {
            Test: GDVector2Array{
                Values: []GDVector2{
                    {
                        X: 10,
                        Y: 11,
                    },
                    {
                        X: 9.999999,
                        Y: 8.8888888,
                    },
                },
            },
            Expect: []byte{24, 0, 0, 0, 2, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65, 255, 255, 31, 65, 227, 56, 14, 65},
        },
        {
           Test: GDVector2Array{},
           Expect: []byte{24, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotVector3Array(t *testing.T) {

    vTest := []struct {
        Test   GDVector3Array
        Expect []byte
    }{
        {
            Test: GDVector3Array{
                Values: []GDVector3{
                    {
                        X: 10,
                        Y: 11,
                        Z: 12,
                    },
                },
            },
            Expect: []byte{25, 0, 0, 0, 1, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65, 0, 0, 64, 65},
        },
        {
            Test: GDVector3Array{
                Values: []GDVector3{
                    {
                        X: 10,
                        Y: 11,
                        Z: 12,
                    },
                    {
                        X: 9.999999,
                        Y: 8.8888888,
                        Z: 7.7777777,
                    },
                },
            },
            Expect: []byte{25, 0, 0, 0, 2, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65, 0, 0, 64, 65, 255, 255, 31, 65, 227, 56, 14, 65, 142, 227, 248, 64},
        },
        {
           Test: GDVector3Array{},
           Expect: []byte{25, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}

func TestVar2GodotColorArray(t *testing.T) {

    vTest := []struct {
        Test   GDColorArray
        Expect []byte
    }{
        {
            Test: GDColorArray{
                Colors: []GDColor{
                    {
                        R: 10,
                        G: 11,
                        B: 12,
                        A: 1,
                    },
                },
            },
            Expect: []byte{26, 0, 0, 0, 1, 0, 0, 0, 0, 0, 32, 65, 0, 0, 48, 65, 0, 0, 64, 65, 0, 0, 128, 63},
        },
        {
            Test: GDColorArray{
                Colors: []GDColor{
                    {
                        R: .2,
                        G: .3,
                        B: .4,
                        A: 1,
                    },
                },
            },
            Expect: []byte{26, 0, 0, 0, 1, 0, 0, 0, 205, 204, 76, 62, 154, 153, 153, 62, 205, 204, 204, 62, 0, 0, 128, 63},
        },
        {
           Test: GDColorArray{},
           Expect: []byte{26, 0, 0, 0, 0, 0, 0, 0},
        },
    }
    for _, v := range vTest {
        fmt.Println(">>>> Testing: " + fmt.Sprintf("%+v", &v.Test))
        gdType, err := Var2GodotBytes(v.Test)
        if err != nil {
            assert.Error(t, err)
        }
        assert.Equal(t, v.Expect, gdType)
    }
}