package productos_test

import (
	"testing"

	p "workplace/digital_academy_1/go-tdd/productos"
)

var scenarios = []struct {
	Escenario string
	Id        int32
	Esperado  string
}{
	{"Para Id 1 debe responder Producto1", 1, "Producto1"},
	{"Para Id 2 debe responder Producto2", 2, "Producto2"},
	{"Para Id 3 debe responder Producto3", 3, "Producto3"},
	{"Para Id 15 debe responder Producto15", 15, "Producto15"},
	{"Para Id 11 debe responder 'No encontrado'", 11, "No encontrado"},
}

var prices = []struct {
	Escenario string
	Id        int32
	Esperado  float32
}{
	{"Para Id 1 debe responder 1.0", 1, 1.0},
	{"Para Id 2 debe responder 1.5", 2, 1.5},
	{"Para Id 3 debe responder 1.8", 3, 1.8},
	{"Para Id 15 debe responder 8.5", 15, 8.5},
	{"Para Id 11 debe responder 'No encontrado'", 11, -1},
}

func TestProductosPrices(t *testing.T) {
	//TODO: Implementar test de acuerdo al requerimiento

	productMap := make(map[int32]*p.Product)

	productMap[1] = &p.Product{Name: "Producto1", Price: 1.0}
	productMap[2] = &p.Product{Name: "Producto2", Price: 1.5}
	productMap[3] = &p.Product{Name: "Producto3", Price: 1.8}
	productMap[15] = &p.Product{Name: "Producto15", Price: 8.5}

	for _, scenario := range prices {
		t.Run(scenario.Escenario, func(t *testing.T) {
			price := p.Findproductprice(productMap, scenario.Id)

			if price != scenario.Esperado {
				t.Error("Resultado no es el esperado: ", price)
			}
		})
	}

}

func TestProductosName(t *testing.T) {
	//TODO: Implementar test de acuerdo al requerimiento

	productMap := make(map[int32]*p.Product)

	productMap[1] = &p.Product{Name: "Producto1", Price: 1.0}
	productMap[2] = &p.Product{Name: "Producto2", Price: 1.5}
	productMap[3] = &p.Product{Name: "Producto3", Price: 1.8}
	productMap[15] = &p.Product{Name: "Producto15", Price: 8.5}

	for _, scenario := range scenarios {
		t.Run(scenario.Escenario, func(t *testing.T) {
			nombre := p.Findproductname(productMap, scenario.Id)

			if nombre != scenario.Esperado {
				t.Error("Resultado no es el esperado: " + nombre)
			}
		})
	}

}
