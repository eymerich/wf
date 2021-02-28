// Code generated by "stringer -output=zdatatype_strings.go -type=DataType -trimprefix=DataType"; DO NOT EDIT.

package wf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DataTypeEmpty-0]
	_ = x[DataTypeUint8-1]
	_ = x[DataTypeUint16-2]
	_ = x[DataTypeUint32-3]
	_ = x[DataTypeUint64-4]
	_ = x[DataTypeInt8-5]
	_ = x[DataTypeInt16-6]
	_ = x[DataTypeInt32-7]
	_ = x[DataTypeInt64-8]
	_ = x[DataTypeFloat-9]
	_ = x[DataTypeDouble-10]
	_ = x[DataTypeByteArray16-11]
	_ = x[DataTypeByteBlob-12]
	_ = x[DataTypeSID-13]
	_ = x[DataTypeSecurityDescriptor-14]
	_ = x[DataTypeTokenInformation-15]
	_ = x[DataTypeTokenAccessInformation-16]
	_ = x[DataTypeUnicodeString-17]
}

const _DataType_name = "EmptyUint8Uint16Uint32Uint64Int8Int16Int32Int64FloatDoubleByteArray16ByteBlobSIDSecurityDescriptorTokenInformationTokenAccessInformationUnicodeString"

var _DataType_index = [...]uint8{0, 5, 10, 16, 22, 28, 32, 37, 42, 47, 52, 58, 69, 77, 80, 98, 114, 136, 149}

func (i DataType) String() string {
	if i >= DataType(len(_DataType_index)-1) {
		return "DataType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DataType_name[_DataType_index[i]:_DataType_index[i+1]]
}
