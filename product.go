package main

import "fmt"

// 1. Интерфейс Прототипа
type Prototype interface {
	Clone() Prototype
	Show()
}

// 2. Конкретный прототип — Товар
type Product struct {
	Name     string
	Category string
	Price    float64
	Size     string
	Color    string
}

// Clone - создаёт копию объекта
func (p *Product) Clone() Prototype {
	clone := *p
	return &clone
}

// Show - выводит информацию о товаре
func (p *Product) Show() {
	fmt.Println("Товар:", p.Name)
	fmt.Println("  Категория:", p.Category)
	fmt.Println("  Цена:", p.Price, "руб")
	fmt.Println("  Размер:", p.Size)
	fmt.Println("  Цвет:", p.Color)
	fmt.Println("---")
}

// 3. Реестр прототипов
type ProductRegistry struct {
	prototypes map[string]Prototype
}

// Новый реестр
func NewProductRegistry() *ProductRegistry {
	return &ProductRegistry{
		prototypes: make(map[string]Prototype),
	}
}

// Добавить прототип
func (r *ProductRegistry) Register(key string, p Prototype) {
	r.prototypes[key] = p
}

// Скопировать прототип
func (r *ProductRegistry) CreateCopy(key string) (Prototype, error) {
	if p, exists := r.prototypes[key]; exists {
		fmt.Println("Клонируем прототип с ключом:", key)
		return p.Clone(), nil
	}
	return nil, fmt.Errorf("ошибка: прототип '%s' не найден", key)
}
