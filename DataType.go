/*---------------------------------------------------------------------------------
// Basic Types

bool
string
int8       int16   int32   int64
uint8      uint16  uint32  uint64
int        //alias for int32 or int64
uint       //alias for uint32 or uint64
byte       //alias for uint8
rune       //alias for int32
float32    float64
complex64  complex128

#---------------------------------------------------------------------------------
Integers Signed:

int      Both int and uint contain same size, either 32 or 64 bit.
int8     8 bits/1 byte
int16    16 bits/2 byte
int32    32 bits/4 byte
int64    64 bits/8 byte
#---------------------------------------------------------------------------------
Integers UnSigned:

uint      Both int and uint contain same size, either 32 or 64 bit.
uint8     8 bits/1 byte
uint16    16 bits/2 byte
uint32    32 bits/4 byte
uint64    64 bits/8 byte
#---------------------------------------------------------------------------------
Floats:

float32	   32 bits/4 byte
float64    64 bits/8 byte
#---------------------------------------------------------------------------------
Complex Numbers:

complex64     Both real and imaginary part are float32
complex128    Both real and imaginary part are float64
#---------------------------------------------------------------------------------
Booleans:

True
False
#---------------------------------------------------------------------------------
String
#---------------------------------------------------------------------------------
Rune
Byte

1- Golang has integer types called byte and rune that are aliases for uint8 and int32 data types, respectively.
2- In Go, the byte and rune data types are used to distinguish characters from integer values.

3- In Go, there is no char data type. It uses byte and rune to represent character values.
4- The byte data type represents ASCII characters.
5- the rune data type represents a more broader set of Unicode characters that are encoded in UTF-8 format.
6- The default type for character values is rune.
7-

#---------------------------------------------------------------------------------
Value Types:

Arrays
Structs
#---------------------------------------------------------------------------------
Reference Types:

Slices
Channels
Maps
Pointers
Functions
#---------------------------------------------------------------------------------
Aggregate Types:

Array
Struct
#---------------------------------------------------------------------------------
Composite Type:

Arrays
Slices
Maps
Structs
#---------------------------------------------------------------------------------