package resp

import "strconv"

const (
	STRING 	= '+'
	ERROR 	= '-'
	INTEGER = ':'
	BULK 		= '$'
	ARRAY 	= '*'
)

type Value struct {
	Typ		string	// Used to determine the data type
	Str		string	// Holds the value of the string received from simple strings
	Num		int			// Holds the value of the integer received from the integers
	Bulk	string	// Used to store the string received from bulk strings
	Array	[]Value	// Holds all the values received from arrays
}

// Marshalling Value to bytes
func (v Value) Marshall() []byte {
	switch v.Typ {
		case "array":
			return v.marshallArray()
		
		case "bulk":
			return v.marshallBulk()

		case "string":
			return v.marshallString()

		case "null":
			return v.marshallNull()

		case "error":
			return v.marshallError()

		default:
			return []byte{}
	}
}

func (v Value) marshallString() []byte {
	var bytes []byte

	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallBulk() []byte {
	var bytes []byte

	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...) // Appending the length of bulk
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallArray() []byte {
	len := len(v.Array)
	var bytes []byte

	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(len)...)
	bytes = append(bytes, '\r', '\n')

	for i:= 0; i < len; i++ {
		bytes = append(bytes, v.Array[i].Marshall()...)
	}

	return bytes
}

func (v Value) marshallError() []byte {
	var bytes []byte

	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')

	return bytes
}

func (v Value) marshallNull() []byte {
	return []byte("$-1\r\n")
}