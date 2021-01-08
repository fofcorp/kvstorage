package storage

import (
	"reflect"
	"testing"
)

func TestInitInMemory(t *testing.T) {
	observed := InitInMemory()
	observedType := reflect.TypeOf(observed)
	expectedType := reflect.TypeOf(&InMemory{})
	if observedType != expectedType {
		t.Error("Expected", expectedType, "got", observedType)
	}
}

func TestPut(t *testing.T) {
	store := InitInMemory()
	err := store.Put("1", "2")
	if err != nil {
		t.Error("Expected nothing, got", err)
	}
	observed, err := store.Get("1")
	if err != nil {
		t.Error("Expected nothing, got", err)
	}
	if observed != "2" {
		t.Error("Expected 2, got", observed)
	}
}

func TestGet(t *testing.T) {
	store := InitInMemory()
	store.Put("1", "2")
	observed, err := store.Get("1")
	if err != nil {
		t.Error("Expected nothing, got", err)
	}
	if observed != "2" {
		t.Error("Expected 2, got", observed)
	}
	observed, err = store.Get("2")
	if err == nil {
		t.Error("Expected no data, got nil")
	}
	if observed != "" {
		t.Error("Expected nothing, got", observed)
	}
}

func TestDelete(t *testing.T) {
	store := InitInMemory()
	store.Put("1", "2")
	observed, err := store.Get("1")
	if err != nil {
		t.Error("Expected nothing, got", err)
	}
	if observed != "2" {
		t.Error("Expected 2, got", observed)
	}
	err = store.Delete("1")
	if err != nil {
		t.Error("Expected nothing, got", err)
	}
	observed, err = store.Get("1")
	if err == nil {
		t.Error("Expected no data, got nil")
	}
	if observed != "" {
		t.Error("Expected nothing, got", observed)
	}
}
