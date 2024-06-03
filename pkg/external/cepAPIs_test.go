package external

import (
	"reflect"
	"strings"
	"testing"
)

func TestViaCep(t *testing.T) {
	cep := "20541155"
	result, err := ViaCep(cep)
	if err != nil {
		t.Errorf("ViaCep() returned an error: %v", err)
	}

	if strings.ReplaceAll(result.Cep, "-", "") != strings.ReplaceAll(cep, "-", "") {
		t.Errorf("ViaCep() returned an invalid CEP: %v", result.Cep)
	}

	fields := []string{
		"Cep", "State", "City", "Neighborhood", "Street", "Source",
	}

	viaCepVal := reflect.ValueOf(result)
	for _, field := range fields {
		val := viaCepVal.FieldByName(field)
		if !val.IsValid() {
			t.Errorf("ViaCep() did not return a Marine struct with the field %s", field)
		}
	}
}

func TestViaCepInvalidZip(t *testing.T) {
	cep := "90541155"
	_, err := ViaCep(cep)
	if err == nil {
		t.Fatalf("ViaCep() returned a value instead of an err: %v", err)
	}
	if err.Error() != "invalid zipcode" {
		t.Errorf("ViaCep() did not return an invalid zipcode error")
	}

}

func TestViaCepInvalidFormat(t *testing.T) {
	cep := "905411551"
	_, err := ViaCep(cep)
	if err == nil {
		t.Fatalf("ViaCep() returned a value instead of an err: %v", err)
	}
	if err.Error() != "can not find zipcode" {
		t.Errorf("ViaCep() did not return an invalid zipcode error")
	}
}

func TestBrasilApiCep(t *testing.T) {
	cep := "20541155"
	result, err := BrasilApiCep(cep)
	if err != nil {
		t.Errorf("BrasilApiCep() returned an error: %v", err)
	}

	if strings.ReplaceAll(result.Cep, "-", "") != strings.ReplaceAll(cep, "-", "") {
		t.Errorf("BrasilApiCep() returned an invalid CEP: %v", result.Cep)
	}

	fields := []string{
		"Cep", "State", "City", "Neighborhood", "Street", "Source",
	}

	viaCepVal := reflect.ValueOf(result)
	for _, field := range fields {
		val := viaCepVal.FieldByName(field)
		if !val.IsValid() {
			t.Errorf("BrasilApiCep() did not return a Marine struct with the field %s", field)
		}
	}
}

func TestBrasilApiCepInvalidZip(t *testing.T) {
	cep := "90541155"
	_, err := BrasilApiCep(cep)
	if err == nil {
		t.Fatalf("BrasilApi() returned a value instead of an err: %v", err)
	}
	if err.Error() != "invalid zipcode" {
		t.Errorf("BrasilApi() did not return an invalid zipcode error")
	}

}

func TestBrasilApiCepInvalidFormat(t *testing.T) {
	cep := "90541A155"
	_, err := BrasilApiCep(cep)
	if err == nil {
		t.Fatalf("ViaCep() returned a value instead of an err: %v", err)
	}
	if err.Error() != "can not find zipcode" {
		t.Errorf("ViaCep() did not return an invalid zipcode error")
	}
}
