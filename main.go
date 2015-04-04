// cgotypes provides a means to output information about how Cgo operates on your platform.
//
//
// Consider examining the intermediate output further:
//
//     $ go tool cgo main.go
//     $ ls _obj
package main

/*
#include <stdbool.h>

struct Example {
    // Scalars
    char _char;
    signed char signed_char;
    unsigned char unsigned_char;

    short _short;
    short int short_int;
    signed short signed_short;
    signed short int signed_short_int;

    unsigned short unsigned_short;
    unsigned short int unsigned_short_int;

    int _int;
    signed int signed_int;

    unsigned _unsigned;
    unsigned int unsigned_int;

    long _long;
    long int long_int;
    signed long signed_long;
    signed long int signed_long_int;

    unsigned long unsigned_long;
    unsigned long int unsigned_long_int;

    long long long_long;
    long long int long_long_int;
    signed long long signed_long_long;
    signed long long int signed_long_long_int;

    float _float;

    double _double;

    // long double long_double;  // Impossible!

    bool _bool;

    union {
        int first;
        struct {
            long long nested;
        } second;
    } _union;

    struct {
        int first_member;
        long long second_member;
    } _struct;

    // Fixed Arrays
    char fixed_ary_char[1];
    signed char fixed_ary_signed_char[1];
    unsigned char fixed_ary_unsigned_char[1];

    short fixed_ary_short[1];
    short int fixed_ary_short_int[1];
    signed short fixed_ary_signed_short[1];
    signed short int fixed_ary_signed_short_int[1];

    unsigned short fixed_ary_unsigned_short[1];
    unsigned short int fixed_ary_unsigned_short_int[1];

    int fixed_ary_int[1];
    signed int fixed_ary_signed_int[1];

    unsigned fixed_ary_unsigned[1];
    unsigned int fixed_ary_unsigned_int[1];

    long fixed_ary_long[1];
    long int fixed_ary_long_int[1];
    signed long fixed_ary_signed_long[1];
    signed long int fixed_ary_signed_long_int[1];

    unsigned long fixed_ary_unsigned_long[1];
    unsigned long int fixed_ary_unsigned_long_int[1];

    long long fixed_ary_long_long[1];
    long long int fixed_ary_long_long_int[1];
    signed long long fixed_ary_signed_long_long[1];
    signed long long int fixed_ary_signed_long_long_int[1];

    float fixed_ary_float[1];

    double fixed_ary_double[1];

    // long double fixed_ary_long_double[1];  // Impossible!

    bool fixed_ary_bool[1];

    union {
        int first;
        struct {
            long long nested;
        } second;
    } fixed_ary_union[1];

    struct {
        int first_member;
        long long second_member;
    } fixed_ary_struct[1];

    // Pointers
    char *ptr_char;
    signed char *ptr_signed_char;
    unsigned char *ptr_unsigned_char;

    short *ptr_short;
    short int *ptr_short_int;
    signed short *ptr_signed_short;
    signed short int *ptr_signed_short_int;

    unsigned short *ptr_unsigned_short;
    unsigned short int *ptr_unsigned_short_int;

    int *ptr_int;
    signed int *ptr_signed_int;

    unsigned *ptr_unsigned;
    unsigned int *ptr_unsigned_int;

    long *ptr_long;
    long int *ptr_long_int;
    signed long *ptr_signed_long;
    signed long int *ptr_signed_long_int;

    unsigned long *ptr_unsigned_long;
    unsigned long int *ptr_unsigned_long_int;

    long long *ptr_long_long;
    long long int *ptr_long_long_int;
    signed long long *ptr_signed_long_long;
    signed long long int *ptr_signed_long_long_int;

    float *ptr_float;

    double *ptr_double;

    // long double *ptr_long_double;  // Impossible!

    bool *ptr_bool;

    union {
        int first;
        struct {
            long long nested;
        } second;
    } *ptr_union;

    struct {
        int first_member;
        long long second_member;
    } *ptr_struct;

    void (*ptr_func)();

    void *void_ptr;
};

typedef char ary_char[];
typedef signed char ary_signed_char[];
typedef unsigned char ary_unsigned_char[];

typedef short ary_short[];
typedef short int ary_short_int[];
typedef signed short ary_signed_short[];
typedef signed short int ary_signed_short_int[];

typedef unsigned short ary_unsigned_short[];
typedef unsigned short int ary_unsigned_short_int[];

typedef int ary_int[];
typedef signed int ary_signed_int[];

typedef unsigned ary_unsigned[];
typedef unsigned int ary_unsigned_int[];

typedef long ary_long[];
typedef long int ary_long_int[];
typedef signed long ary_signed_long[];
typedef signed long int ary_signed_long_int[];

typedef unsigned long ary_unsigned_long[];
typedef unsigned long int ary_unsigned_long_int[];

typedef long long ary_long_long[];
typedef long long int ary_long_long_int[];
typedef signed long long ary_signed_long_long[];
typedef signed long long int ary_signed_long_long_int[];

typedef float ary_float[];

typedef double ary_double[];

// typedef long double ary_long_double[];  // Impossible!

typedef bool ary_bool[];

typedef union {
    int first;
    struct {
        long long nested;
    } second;
} ary_union[];

typedef struct {
    int first_member;
    long long second_member;
} ary_struct[];
*/
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
)

func describeField(num int, f reflect.StructField) {
	typ := f.Type
	fmt.Printf(
		"%d\tname: %s\tkind: %s\ttype: %s\n",
		num, f.Name, typ.Kind(), typ)
}

func describeSingleton(data interface{}) {
	typ := reflect.TypeOf(data)
	fmt.Printf(
		"kind: %s\ttype: %s\n",
		typ.Kind(), typ)
}

func main() {
	fmt.Println("Readout for", runtime.GOARCH, runtime.GOOS)
	var (
		nonAry    C.struct_Example
		nonAryTyp = reflect.TypeOf(nonAry)
	)
	for i := 0; i < nonAryTyp.NumField(); i++ {
		describeField(i, nonAryTyp.Field(i))
	}
	// Non-embedables:
	var (
		ary_char          C.ary_char
		ary_signed_char   C.ary_signed_char
		ary_unsigned_char C.ary_unsigned_char

		ary_short            C.ary_short
		ary_short_int        C.ary_short_int
		ary_signed_short     C.ary_signed_short
		ary_signed_short_int C.ary_signed_short_int

		ary_unsigned_short     C.ary_unsigned_short
		ary_unsigned_short_int C.ary_unsigned_short_int

		ary_int        C.ary_int
		ary_signed_int C.ary_signed_int

		ary_unsigned     C.ary_unsigned
		ary_unsigned_int C.ary_unsigned_int

		ary_long            C.ary_long
		ary_long_int        C.ary_long_int
		ary_signed_long     C.ary_signed_long
		ary_signed_long_int C.ary_signed_long_int

		ary_unsigned_long     C.ary_unsigned_long
		ary_unsigned_long_int C.ary_unsigned_long_int

		ary_long_long            C.ary_long_long
		ary_long_long_int        C.ary_long_long_int
		ary_signed_long_long     C.ary_signed_long_long
		ary_signed_long_long_int C.ary_signed_long_long_int

		ary_float C.ary_float

		ary_double C.ary_double

		// ary_long_double C.ary_long_double  // Impossible!

		ary_bool C.ary_bool

		ary_union C.ary_union

		ary_struct C.ary_struct
	)
	describeSingleton(ary_char)
	describeSingleton(ary_signed_char)
	describeSingleton(ary_unsigned_char)

	describeSingleton(ary_short)
	describeSingleton(ary_short_int)
	describeSingleton(ary_signed_short)
	describeSingleton(ary_signed_short_int)

	describeSingleton(ary_unsigned_short)
	describeSingleton(ary_unsigned_short_int)

	describeSingleton(ary_int)
	describeSingleton(ary_signed_int)

	describeSingleton(ary_unsigned)
	describeSingleton(ary_unsigned_int)

	describeSingleton(ary_long)
	describeSingleton(ary_long_int)
	describeSingleton(ary_signed_long)
	describeSingleton(ary_signed_long_int)

	describeSingleton(ary_unsigned_long)
	describeSingleton(ary_unsigned_long_int)

	describeSingleton(ary_long_long)
	describeSingleton(ary_long_long_int)
	describeSingleton(ary_signed_long_long)
	describeSingleton(ary_signed_long_long_int)

	describeSingleton(ary_float)

	describeSingleton(ary_double)

	// describeSingleton(ary_long_double) // Impossible!

	describeSingleton(ary_bool)

	describeSingleton(ary_union)

	describeSingleton(ary_struct)
}
