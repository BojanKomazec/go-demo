package elliotchancephpserializedemo

import (
	"log"

	"github.com/elliotchance/phpserialize"
)

func deserializePHPIntArray(serializedPHPArray string) ([]string, error) {
	if serializedPHPArray == "" {
		return make([]string, 0), nil
	}

	var dest []interface{}
	if err := phpserialize.Unmarshal([]byte(serializedPHPArray), &dest); err != nil {
		return nil, err
	}

	stringSlice := make([]string, len(dest))
	for i, v := range dest {
		stringSlice[i] = v.(string)
	}

	return stringSlice, nil
}

func demo() {
	serializedPHPArray := "a:1:{i:0;s:3:\"289\";}"
	elements, err := deserializePHPIntArray(serializedPHPArray)
	if err != nil {
		log.Println("err = ", err.Error())
	} else {
		log.Println("elements = ", elements)
	}
}

// ShowDemo func
func ShowDemo() {
	demo()
}
