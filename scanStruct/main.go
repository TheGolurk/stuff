// Author: Christian Hernandez @TheGolurk
// Description: When someone is working with SQL with GO there are some problems with scan a huge
// fields in a struct for example... if err = rows.Scan(&struct.field1, &struct.fieldn...); err != nil {...
// With this implementation you can read as a pointer rows.Scan(&interfacePointer...) and then use
// this funcion to append the result to an array of struct of the same type of the struct passed in
// the funcion err = scanToStruct(mymodel, mydata)... myModels = append(myModels, myModel) and then
// you can convert easy to JSON: json, err := json.Marshal(myModels)
package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Info struct {
	Name  string
	Phone int64
}

func main() {
	a := Info{}
	a.Name = "asdasdkj"
	a.Phone = 1231
	b := []string{"Carlos", "123131231"}
	err := scanToStruct(&a, b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After reflect struct:", a)
}

// Expected struct and string with the data as a array
func scanToStruct(Model interface{}, Data []string) error {
	items := reflect.ValueOf(Model).Elem()

	if !items.CanAddr() {
		return errors.New("It must be a pointer")
	}

	if items.Kind() != reflect.Struct {
		return errors.New("Not a struct")
	}

	value := reflect.Indirect(items) // Pointer to item slice

	if value.NumField() != len(Data) {
		return errors.New("It should be the same lenght")
	}

	for j := 0; j < value.NumField(); j++ {
		fieldName := value.Type().Field(j).Name
		actual := value.FieldByName(fieldName)
		if !actual.CanSet() {
			return errors.New("Cannot set to field")
		}
		set := "string"
		vals := Data[j]
		var vali int64
		var err error
		if actual.Type().String() == "int64" { // Add values if type is int32 int63, etc...
			vali, err = strconv.ParseInt(vals, 10, 64)
			if err != nil {
				return errors.New(fmt.Sprintf("Cannot parse to int, reason: %v", err))
			}
			set = "int"
			vali = int64(vali)
		}

		switch set {
		case "string":
			actual.SetString(vals)
		case "int":
			actual.SetInt(vali)
		}
	}

	return nil
}
