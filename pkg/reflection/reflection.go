package reflection

import (
	"fmt"
	"reflect"
)

type orders struct {
	orderId    int
	customerId int
}

func Reflection() {
	var greeting interface{} = 1
	i, ok := greeting.(int)
	fmt.Println(i, ok)

	o := orders{
		orderId:    111,
		customerId: 343434,
	}
	oValue := reflect.ValueOf(o)
	fmt.Println("type:", reflect.TypeOf(o))
	fmt.Println("struct name:", reflect.TypeOf(o).Name())
	if oValue.Kind() == reflect.Struct {
		numField := oValue.NumField()
		for i := 0; i < numField; i++ {
			fmt.Printf("field no: %d, type: %T, value: %v\n", i, oValue.Field(i).Int(), oValue.Field(i))
		}
	}

}
