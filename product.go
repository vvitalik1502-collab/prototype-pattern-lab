package main

import "fmt"

type Prototype interface {
	Clone() Prototype
	Show()
}

type Product struct {
	Name     string
	Category string
	Price    float64
	Size     string
	Color    string
}

func (p *Product) Clone() Prototype {
	clone := *p
	return &clone
}

func (p *Product) Show() {
	fmt.Println("Товар:", p.Name)
	fmt.Println("  Категория:", p.Category)
	fmt.Println("  Цена:", p.Price, "руб")
	fmt.Println("  Размер:", p.Size)
	fmt.Println("  Цвет:", p.Color)
	fmt.Println("---")
}

type ProductRegistry struct {
	prototypes map[string]Prototype
}

func NewProductRegistry() *ProductRegistry {
	return &ProductRegistry{
		prototypes: make(map[string]Prototype),
	}
}

func (r *ProductRegistry) Register(key string, p Prototype) {
	r.prototypes[key] = p
}

func (r *ProductRegistry) CreateCopy(key string) (Prototype, error) {
	if p, exists := r.prototypes[key]; exists {
		fmt.Println("Клонируем прототип с ключом:", key)
		return p.Clone(), nil
	}
	return nil, fmt.Errorf("ошибка: прототип '%s' не найден", key)
}
