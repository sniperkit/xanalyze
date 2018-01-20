// DO NOT EDIT!
// This file was generated by generate/ops.go

package up

// OpPlus implements plus operator
func OpPlus(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] + v[1]
	case [2]uint8:
		return v[0] + v[1]
	case [2]int16:
		return v[0] + v[1]
	case [2]uint16:
		return v[0] + v[1]
	case [2]int32:
		return v[0] + v[1]
	case [2]int:
		return v[0] + v[1]
	case [2]uint32:
		return v[0] + v[1]
	case [2]uint:
		return v[0] + v[1]
	case [2]int64:
		return v[0] + v[1]
	case [2]uint64:
		return v[0] + v[1]
	case [2]float32:
		return v[0] + v[1]
	case [2]float64:
		return v[0] + v[1]
	}
	panic(`unknown types for operator +`)
}

// OpMinus implements minus opertator
func OpMinus(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] - v[1]
	case [2]uint8:
		return v[0] - v[1]
	case [2]int16:
		return v[0] - v[1]
	case [2]uint16:
		return v[0] - v[1]
	case [2]int32:
		return v[0] - v[1]
	case [2]int:
		return v[0] - v[1]
	case [2]uint32:
		return v[0] - v[1]
	case [2]uint:
		return v[0] - v[1]
	case [2]int64:
		return v[0] - v[1]
	case [2]uint64:
		return v[0] - v[1]
	case [2]float32:
		return v[0] - v[1]
	case [2]float64:
		return v[0] - v[1]
	}
	panic(`unknown types for operator -`)
}

// OpProduct implements product operator
func OpProduct(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] * v[1]
	case [2]uint8:
		return v[0] * v[1]
	case [2]int16:
		return v[0] * v[1]
	case [2]uint16:
		return v[0] * v[1]
	case [2]int32:
		return v[0] * v[1]
	case [2]int:
		return v[0] * v[1]
	case [2]uint32:
		return v[0] * v[1]
	case [2]uint:
		return v[0] * v[1]
	case [2]int64:
		return v[0] * v[1]
	case [2]uint64:
		return v[0] * v[1]
	case [2]float32:
		return v[0] * v[1]
	case [2]float64:
		return v[0] * v[1]
	}
	panic(`unknown types for operator *`)
}

// OpDivide implements division operator
func OpDivide(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] / v[1]
	case [2]uint8:
		return v[0] / v[1]
	case [2]int16:
		return v[0] / v[1]
	case [2]uint16:
		return v[0] / v[1]
	case [2]int32:
		return v[0] / v[1]
	case [2]int:
		return v[0] / v[1]
	case [2]uint32:
		return v[0] / v[1]
	case [2]uint:
		return v[0] / v[1]
	case [2]int64:
		return v[0] / v[1]
	case [2]uint64:
		return v[0] / v[1]
	case [2]float32:
		return v[0] / v[1]
	case [2]float64:
		return v[0] / v[1]
	}
	panic(`unknown types for operator /`)
}

// OpMod implements modulo operator
func OpMod(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] % v[1]
	case [2]uint8:
		return v[0] % v[1]
	case [2]int16:
		return v[0] % v[1]
	case [2]uint16:
		return v[0] % v[1]
	case [2]int32:
		return v[0] % v[1]
	case [2]int:
		return v[0] % v[1]
	case [2]uint32:
		return v[0] % v[1]
	case [2]uint:
		return v[0] % v[1]
	case [2]int64:
		return v[0] % v[1]
	case [2]uint64:
		return v[0] % v[1]
	}
	panic(`unknown types for operator %`)
}

// OpRightShift implements right shift operator
func OpRightShift(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] >> uint8(v[1])
	case [2]uint8:
		return v[0] >> v[1]
	case [2]int16:
		return v[0] >> uint16(v[1])
	case [2]uint16:
		return v[0] >> v[1]
	case [2]int32:
		return v[0] >> uint32(v[1])
	case [2]int:
		return v[0] >> uint(v[1])
	case [2]uint32:
		return v[0] >> v[1]
	case [2]uint:
		return v[0] >> v[1]
	case [2]int64:
		return v[0] >> uint64(v[1])
	case [2]uint64:
		return v[0] >> v[1]
	}
	panic(`unknown types for operator >>`)
}

// OpLeftShift implements left shift operator
func OpLeftShift(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] << uint8(v[1])
	case [2]uint8:
		return v[0] << v[1]
	case [2]int16:
		return v[0] << uint16(v[1])
	case [2]uint16:
		return v[0] << v[1]
	case [2]int32:
		return v[0] << uint32(v[1])
	case [2]int:
		return v[0] << uint(v[1])
	case [2]uint32:
		return v[0] << v[1]
	case [2]uint:
		return v[0] << v[1]
	case [2]int64:
		return v[0] << uint64(v[1])
	case [2]uint64:
		return v[0] << v[1]
	}
	panic(`unknown types for operator <<`)
}

// OpAnd implements and operator
func OpAnd(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] & v[1]
	case [2]uint8:
		return v[0] & v[1]
	case [2]int16:
		return v[0] & v[1]
	case [2]uint16:
		return v[0] & v[1]
	case [2]int32:
		return v[0] & v[1]
	case [2]int:
		return v[0] & v[1]
	case [2]uint32:
		return v[0] & v[1]
	case [2]uint:
		return v[0] & v[1]
	case [2]int64:
		return v[0] & v[1]
	case [2]uint64:
		return v[0] & v[1]
	}
	panic(`unknown types for operator &`)
}

// OpOr implements or operator
func OpOr(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] | v[1]
	case [2]uint8:
		return v[0] | v[1]
	case [2]int16:
		return v[0] | v[1]
	case [2]uint16:
		return v[0] | v[1]
	case [2]int32:
		return v[0] | v[1]
	case [2]int:
		return v[0] | v[1]
	case [2]uint32:
		return v[0] | v[1]
	case [2]uint:
		return v[0] | v[1]
	case [2]int64:
		return v[0] | v[1]
	case [2]uint64:
		return v[0] | v[1]
	}
	panic(`unknown types for operator |`)
}

// OpAndNot implements and not operator
func OpAndNot(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] &^ v[1]
	case [2]uint8:
		return v[0] &^ v[1]
	case [2]int16:
		return v[0] &^ v[1]
	case [2]uint16:
		return v[0] &^ v[1]
	case [2]int32:
		return v[0] &^ v[1]
	case [2]int:
		return v[0] &^ v[1]
	case [2]uint32:
		return v[0] &^ v[1]
	case [2]uint:
		return v[0] &^ v[1]
	case [2]int64:
		return v[0] &^ v[1]
	case [2]uint64:
		return v[0] &^ v[1]
	}
	panic(`unknown types for operator &^`)
}

// OpXor implements xor operator
func OpXor(a interface{}, b interface{}) interface{} {
	c := commonType(a, b)
	switch v := c.(type) {
	case [2]int8:
		return v[0] ^ v[1]
	case [2]uint8:
		return v[0] ^ v[1]
	case [2]int16:
		return v[0] ^ v[1]
	case [2]uint16:
		return v[0] ^ v[1]
	case [2]int32:
		return v[0] ^ v[1]
	case [2]int:
		return v[0] ^ v[1]
	case [2]uint32:
		return v[0] ^ v[1]
	case [2]uint:
		return v[0] ^ v[1]
	case [2]int64:
		return v[0] ^ v[1]
	case [2]uint64:
		return v[0] ^ v[1]
	}
	panic(`unknown types for operator ^`)
}