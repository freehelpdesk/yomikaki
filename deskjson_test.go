package deskjson

import (
	"log"
	"testing"
)

func TestDirectWrite(t *testing.T) {
	jsonObj := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"value": "original value",
			},
		},
	}
	log.Println(jsonObj)

	m := DirectWrite("foo->bar->aaa->eee", jsonObj, 12)

	log.Println(m)
}

func TestDirectRead(t *testing.T) {
	jsonObj := map[string]interface{}{
		"foo": map[string]interface{}{
			"bar": map[string]interface{}{
				"value": "original value",
			},
		},
	}

	read, err := DirectRead("foo->bar", jsonObj)
	if err != nil {
		t.Fatal(err)
	}

	log.Println(read)
}
