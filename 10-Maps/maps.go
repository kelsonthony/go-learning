package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// as chaves são todas do mesmo tipo, e os valores também
	usuario := map[string]string{
		"nome":      "Mchilanny",
		"sobrenome": "Bussinguer",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Mchilanny",
			"ultimo":   "Bussinguer",
		},
		"curso": {
			"nome":   "T.O",
			"campus": "Campus 2",
		},
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2["curso"])

	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Gemeos",
	}
	fmt.Println(usuario2)
}
