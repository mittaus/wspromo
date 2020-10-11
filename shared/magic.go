package shared

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ValidateInputs(obj []interface{}) {
	for i := 0; i < len(obj); i++ {
		for _, v := range obj[i].([]interface{}) { // use type assertion to loop over []interface{}
			fmt.Printf("%v ", v)
		}
	}

}

func ResponseService(status string, code string, message string, data interface{}) string {
	resp := Resp{Error: code, Message: message, Data: data}
	response := Response{Status: status, Resp: resp}
	e, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("%v", string(e))
}

func StringResponseToResponseObj(stringObj string) interface{} {
	byt := []byte(stringObj)
	var responseData map[string]interface{}
	if err := json.Unmarshal(byt, &responseData); err != nil {
		panic(err)
	}
	return responseData

}

func ConvertErrorDvToErrorDbStrcut(err error) ErrorDB {
	res2B, _ := json.Marshal(err)
	var eDv ErrorDB
	json.Unmarshal(res2B, &eDv)
	return eDv
}

func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}
