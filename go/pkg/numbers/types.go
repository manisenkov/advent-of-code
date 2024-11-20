package numbers

// SignedInt constraint represents any signed integer type
type SignedInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInt constraint represents any unsigned integer type
type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// AnyInt constraint represents any integer tyoe, both signed and usigned
type AnyInt interface {
	SignedInt | UnsignedInt
}

// Float constraint represents any float type
type Float interface {
	~float32 | ~float64
}

// Number constract represents any number type
type Number interface {
	AnyInt | Float
}
