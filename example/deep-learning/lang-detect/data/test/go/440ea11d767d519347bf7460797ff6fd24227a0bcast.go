// DO NOT EDIT!
// This file was generated by generate/cast.go

// +build amd64

package up

func convertToInt8(a interface{}) int8 {
	switch x := a.(type) {
	case int8:
		return x
	case uint8:
		return int8(x)
	case int16:
		return int8(x)
	case uint16:
		return int8(x)
	case int32:
		return int8(x)
	case int:
		return int8(x)
	case uint32:
		return int8(x)
	case uint:
		return int8(x)
	case int64:
		return int8(x)
	case uint64:
		return int8(x)
	case float32:
		return int8(x)
	case float64:
		return int8(x)
	}
	panic("cannot convert to type int8")
}

func convertToUint8(a interface{}) uint8 {
	switch x := a.(type) {
	case int8:
		return uint8(x)
	case uint8:
		return x
	case int16:
		return uint8(x)
	case uint16:
		return uint8(x)
	case int32:
		return uint8(x)
	case int:
		return uint8(x)
	case uint32:
		return uint8(x)
	case uint:
		return uint8(x)
	case int64:
		return uint8(x)
	case uint64:
		return uint8(x)
	case float32:
		return uint8(x)
	case float64:
		return uint8(x)
	}
	panic("cannot convert to type uint8")
}

func convertToInt16(a interface{}) int16 {
	switch x := a.(type) {
	case int8:
		return int16(x)
	case uint8:
		return int16(x)
	case int16:
		return x
	case uint16:
		return int16(x)
	case int32:
		return int16(x)
	case int:
		return int16(x)
	case uint32:
		return int16(x)
	case uint:
		return int16(x)
	case int64:
		return int16(x)
	case uint64:
		return int16(x)
	case float32:
		return int16(x)
	case float64:
		return int16(x)
	}
	panic("cannot convert to type int16")
}

func convertToUint16(a interface{}) uint16 {
	switch x := a.(type) {
	case int8:
		return uint16(x)
	case uint8:
		return uint16(x)
	case int16:
		return uint16(x)
	case uint16:
		return x
	case int32:
		return uint16(x)
	case int:
		return uint16(x)
	case uint32:
		return uint16(x)
	case uint:
		return uint16(x)
	case int64:
		return uint16(x)
	case uint64:
		return uint16(x)
	case float32:
		return uint16(x)
	case float64:
		return uint16(x)
	}
	panic("cannot convert to type uint16")
}

func convertToInt32(a interface{}) int32 {
	switch x := a.(type) {
	case int8:
		return int32(x)
	case uint8:
		return int32(x)
	case int16:
		return int32(x)
	case uint16:
		return int32(x)
	case int32:
		return x
	case int:
		return int32(x)
	case uint32:
		return int32(x)
	case uint:
		return int32(x)
	case int64:
		return int32(x)
	case uint64:
		return int32(x)
	case float32:
		return int32(x)
	case float64:
		return int32(x)
	}
	panic("cannot convert to type int32")
}

func convertToInt(a interface{}) int {
	switch x := a.(type) {
	case int8:
		return int(x)
	case uint8:
		return int(x)
	case int16:
		return int(x)
	case uint16:
		return int(x)
	case int32:
		return int(x)
	case int:
		return x
	case uint32:
		return int(x)
	case uint:
		return int(x)
	case int64:
		return int(x)
	case uint64:
		return int(x)
	case float32:
		return int(x)
	case float64:
		return int(x)
	}
	panic("cannot convert to type int")
}

func convertToUint32(a interface{}) uint32 {
	switch x := a.(type) {
	case int8:
		return uint32(x)
	case uint8:
		return uint32(x)
	case int16:
		return uint32(x)
	case uint16:
		return uint32(x)
	case int32:
		return uint32(x)
	case int:
		return uint32(x)
	case uint32:
		return x
	case uint:
		return uint32(x)
	case int64:
		return uint32(x)
	case uint64:
		return uint32(x)
	case float32:
		return uint32(x)
	case float64:
		return uint32(x)
	}
	panic("cannot convert to type uint32")
}

func convertToUint(a interface{}) uint {
	switch x := a.(type) {
	case int8:
		return uint(x)
	case uint8:
		return uint(x)
	case int16:
		return uint(x)
	case uint16:
		return uint(x)
	case int32:
		return uint(x)
	case int:
		return uint(x)
	case uint32:
		return uint(x)
	case uint:
		return x
	case int64:
		return uint(x)
	case uint64:
		return uint(x)
	case float32:
		return uint(x)
	case float64:
		return uint(x)
	}
	panic("cannot convert to type uint")
}

func convertToInt64(a interface{}) int64 {
	switch x := a.(type) {
	case int8:
		return int64(x)
	case uint8:
		return int64(x)
	case int16:
		return int64(x)
	case uint16:
		return int64(x)
	case int32:
		return int64(x)
	case int:
		return int64(x)
	case uint32:
		return int64(x)
	case uint:
		return int64(x)
	case int64:
		return x
	case uint64:
		return int64(x)
	case float32:
		return int64(x)
	case float64:
		return int64(x)
	}
	panic("cannot convert to type int64")
}

func convertToUint64(a interface{}) uint64 {
	switch x := a.(type) {
	case int8:
		return uint64(x)
	case uint8:
		return uint64(x)
	case int16:
		return uint64(x)
	case uint16:
		return uint64(x)
	case int32:
		return uint64(x)
	case int:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint:
		return uint64(x)
	case int64:
		return uint64(x)
	case uint64:
		return x
	case float32:
		return uint64(x)
	case float64:
		return uint64(x)
	}
	panic("cannot convert to type uint64")
}

func convertToFloat32(a interface{}) float32 {
	switch x := a.(type) {
	case int8:
		return float32(x)
	case uint8:
		return float32(x)
	case int16:
		return float32(x)
	case uint16:
		return float32(x)
	case int32:
		return float32(x)
	case int:
		return float32(x)
	case uint32:
		return float32(x)
	case uint:
		return float32(x)
	case int64:
		return float32(x)
	case uint64:
		return float32(x)
	case float32:
		return x
	case float64:
		return float32(x)
	}
	panic("cannot convert to type float32")
}

func convertToFloat64(a interface{}) float64 {
	switch x := a.(type) {
	case int8:
		return float64(x)
	case uint8:
		return float64(x)
	case int16:
		return float64(x)
	case uint16:
		return float64(x)
	case int32:
		return float64(x)
	case int:
		return float64(x)
	case uint32:
		return float64(x)
	case uint:
		return float64(x)
	case int64:
		return float64(x)
	case uint64:
		return float64(x)
	case float32:
		return float64(x)
	case float64:
		return x
	}
	panic("cannot convert to type float64")
}

func commonType(a interface{}, b interface{}) interface{} {
	switch x := a.(type) {
	case int8:
		switch y := b.(type) {
		case int8:
			return [2]int8{x, y}
		case uint8:
			return [2]uint8{uint8(x), y}
		case int16:
			return [2]int16{int16(x), y}
		case uint16:
			return [2]uint16{uint16(x), y}
		case int32:
			return [2]int32{int32(x), y}
		case int:
			return [2]int{int(x), y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case uint8:
		switch y := b.(type) {
		case int8:
			return [2]uint8{x, uint8(y)}
		case uint8:
			return [2]uint8{x, y}
		case int16:
			return [2]int16{int16(x), y}
		case uint16:
			return [2]uint16{uint16(x), y}
		case int32:
			return [2]int32{int32(x), y}
		case int:
			return [2]int{int(x), y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case int16:
		switch y := b.(type) {
		case int8:
			return [2]int16{x, int16(y)}
		case uint8:
			return [2]int16{x, int16(y)}
		case int16:
			return [2]int16{x, y}
		case uint16:
			return [2]uint16{uint16(x), y}
		case int32:
			return [2]int32{int32(x), y}
		case int:
			return [2]int{int(x), y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case uint16:
		switch y := b.(type) {
		case int8:
			return [2]uint16{x, uint16(y)}
		case uint8:
			return [2]uint16{x, uint16(y)}
		case int16:
			return [2]uint16{x, uint16(y)}
		case uint16:
			return [2]uint16{x, y}
		case int32:
			return [2]int32{int32(x), y}
		case int:
			return [2]int{int(x), y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case int32:
		switch y := b.(type) {
		case int8:
			return [2]int32{x, int32(y)}
		case uint8:
			return [2]int32{x, int32(y)}
		case int16:
			return [2]int32{x, int32(y)}
		case uint16:
			return [2]int32{x, int32(y)}
		case int32:
			return [2]int32{x, y}
		case int:
			return [2]int{int(x), y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case int:
		switch y := b.(type) {
		case int8:
			return [2]int{x, int(y)}
		case uint8:
			return [2]int{x, int(y)}
		case int16:
			return [2]int{x, int(y)}
		case uint16:
			return [2]int{x, int(y)}
		case int32:
			return [2]int{x, int(y)}
		case int:
			return [2]int{x, y}
		case uint32:
			return [2]uint32{uint32(x), y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case uint32:
		switch y := b.(type) {
		case int8:
			return [2]uint32{x, uint32(y)}
		case uint8:
			return [2]uint32{x, uint32(y)}
		case int16:
			return [2]uint32{x, uint32(y)}
		case uint16:
			return [2]uint32{x, uint32(y)}
		case int32:
			return [2]uint32{x, uint32(y)}
		case int:
			return [2]uint32{x, uint32(y)}
		case uint32:
			return [2]uint32{x, y}
		case uint:
			return [2]uint{uint(x), y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case uint:
		switch y := b.(type) {
		case int8:
			return [2]uint{x, uint(y)}
		case uint8:
			return [2]uint{x, uint(y)}
		case int16:
			return [2]uint{x, uint(y)}
		case uint16:
			return [2]uint{x, uint(y)}
		case int32:
			return [2]uint{x, uint(y)}
		case int:
			return [2]uint{x, uint(y)}
		case uint32:
			return [2]uint{x, uint(y)}
		case uint:
			return [2]uint{x, y}
		case int64:
			return [2]int64{int64(x), y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case int64:
		switch y := b.(type) {
		case int8:
			return [2]int64{x, int64(y)}
		case uint8:
			return [2]int64{x, int64(y)}
		case int16:
			return [2]int64{x, int64(y)}
		case uint16:
			return [2]int64{x, int64(y)}
		case int32:
			return [2]int64{x, int64(y)}
		case int:
			return [2]int64{x, int64(y)}
		case uint32:
			return [2]int64{x, int64(y)}
		case uint:
			return [2]int64{x, int64(y)}
		case int64:
			return [2]int64{x, y}
		case uint64:
			return [2]uint64{uint64(x), y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case uint64:
		switch y := b.(type) {
		case int8:
			return [2]uint64{x, uint64(y)}
		case uint8:
			return [2]uint64{x, uint64(y)}
		case int16:
			return [2]uint64{x, uint64(y)}
		case uint16:
			return [2]uint64{x, uint64(y)}
		case int32:
			return [2]uint64{x, uint64(y)}
		case int:
			return [2]uint64{x, uint64(y)}
		case uint32:
			return [2]uint64{x, uint64(y)}
		case uint:
			return [2]uint64{x, uint64(y)}
		case int64:
			return [2]uint64{x, uint64(y)}
		case uint64:
			return [2]uint64{x, y}
		case float32:
			return [2]float32{float32(x), y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case float32:
		switch y := b.(type) {
		case int8:
			return [2]float32{x, float32(y)}
		case uint8:
			return [2]float32{x, float32(y)}
		case int16:
			return [2]float32{x, float32(y)}
		case uint16:
			return [2]float32{x, float32(y)}
		case int32:
			return [2]float32{x, float32(y)}
		case int:
			return [2]float32{x, float32(y)}
		case uint32:
			return [2]float32{x, float32(y)}
		case uint:
			return [2]float32{x, float32(y)}
		case int64:
			return [2]float32{x, float32(y)}
		case uint64:
			return [2]float32{x, float32(y)}
		case float32:
			return [2]float32{x, y}
		case float64:
			return [2]float64{float64(x), y}
		}
	case float64:
		switch y := b.(type) {
		case int8:
			return [2]float64{x, float64(y)}
		case uint8:
			return [2]float64{x, float64(y)}
		case int16:
			return [2]float64{x, float64(y)}
		case uint16:
			return [2]float64{x, float64(y)}
		case int32:
			return [2]float64{x, float64(y)}
		case int:
			return [2]float64{x, float64(y)}
		case uint32:
			return [2]float64{x, float64(y)}
		case uint:
			return [2]float64{x, float64(y)}
		case int64:
			return [2]float64{x, float64(y)}
		case uint64:
			return [2]float64{x, float64(y)}
		case float32:
			return [2]float64{x, float64(y)}
		case float64:
			return [2]float64{x, y}
		}
	}
	return nil
}