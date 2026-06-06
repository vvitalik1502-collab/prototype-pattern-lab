from abc import ABC, abstractmethod
import copy

# 1. Интерфейс Прототипа
class ProductPrototype(ABC):
    @abstractmethod
    def clone(self):
        pass

# 2. Конкретный прототип — Товар в интернет-магазине
class Product(ProductPrototype):
    def __init__(self, name: str, category: str, price: float, size: str, color: str):
        self.name = name
        self.category = category
        self.price = price
        self.size = size
        self.color = color

    def clone(self):
        """Создаёт точную копию объекта"""
        return copy.deepcopy(self)

    def __str__(self):
        return f"Product({self.name}, {self.category}, {self.price}₽, {self.size}, {self.color})"


# 3. Реестр прототипов (хранилище товаров-шаблонов)
class ProductRegistry:
    def __init__(self):
        self._prototypes = {}

    def register(self, key: str, product: Product):
        self._prototypes[key] = product

    def create_copy(self, key: str, **changes):
        """Клонирует прототип и изменяет указанные поля"""
        if key not in self._prototypes:
            raise ValueError(f"Прототип '{key}' не найден")
        
        cloned = self._prototypes[key].clone()
        for attr, value in changes.items():
            if hasattr(cloned, attr):
                setattr(cloned, attr, value)
        return cloned
