package productos

//Product estructura representativa de un producto
type Product struct {
	Name  string
	Price float32
}

//Findproductname función que busca si existe un producto
func Findproductname(productos map[int32]*Product, id int32) string {
	if productos[id] == nil {
		return "No encontrado"
	}

	return productos[id].Name

}

//Findproductprice funcion que busca si existe precio de un producto
func Findproductprice(productos map[int32]*Product, id int32) float32 {
	if productos[id] == nil {
		return -1
	}

	return productos[id].Price

}
