package pulic_type

import (
	"encoding/json"
	"testing"
)

var data = []byte(`{"key1":null,"key2":"22"}`)

type TestStringNullType struct {
	Key1 StringNull `json:"key1,omitempty"`
	Key2 StringNull `json:"key2,omitempty"`
	Key3 StringNull `json:"key3,omitempty"`
}

func TestStringNull(t *testing.T) {
	testData := new(TestStringNullType)
	err := json.Unmarshal(data, testData)
	if err != nil {
		t.Error(err)
	}

	if testData.Key1.Status() != 1 ||
		testData.Key2.Status() != 2 ||
		testData.Key3.Status() != 0 {
		t.Error("stringnull数据当前判断错误")
	}

	testData2 := &TestStringNullType{
		Key1: StringNull{},
	}
	testData2.Key1.Set("111")

	data2, err := json.Marshal(testData2)
	if err != nil {
		t.Error(err)
	}

	if string(data2) != `{"key1":"111","key2":null,"key3":null}` {
		t.Error("stringnull MarshalJSON 错误")
	}
}

type TestBoolNullType struct {
	Key1 BoolNull `json:"key1,omitempty"`
	Key2 BoolNull `json:"key2,omitempty"`
	Key3 BoolNull `json:"key3,omitempty"`
}

var data2 = []byte(`{"key1":null,"key2":true}`)

func TestBoolNull(t *testing.T) {
	testData := new(TestBoolNullType)
	err := json.Unmarshal(data2, testData)
	if err != nil {
		t.Error(err)
	}

	if testData.Key1.Status() != 1 ||
		testData.Key3.Status() != 0 {
		t.Error("BoolNull数据当前判断错误")
	}

	testData2 := &TestBoolNullType{
		Key1: BoolNull{},
		Key2: BoolNull{},
		Key3: BoolNull{},
	}
	testData2.Key1.Set(false)
	testData2.Key2.Set(false)
	testData2.Key2.Null()

	data2, err := json.Marshal(testData2)
	if err != nil {
		t.Error(err)
	}

	if string(data2) != `{"key1":false,"key2":null,"key3":null}` {
		t.Error("BoolNull MarshalJSON 错误")
	}
}
