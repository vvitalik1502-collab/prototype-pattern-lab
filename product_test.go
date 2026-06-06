package main

import "testing"

func TestCloneWorks(t *testing.T) {
	original := &Product{Name: "Футболка", Price: 999}
	cloned := original.Clone()

	if original == cloned {
		t.Error("Клон и оригинал - один объект!")
	}
}

func TestRegistryWorks(t *testing.T) {
	registry := NewProductRegistry()
	registry.Register("test", &Product{Name: "Тест"})

	copied, err := registry.CreateCopy("test")
	if err != nil {
		t.Error("Ошибка клонирования:", err)
	}
	if copied == nil {
		t.Error("Копия не создана!")
	}
}
