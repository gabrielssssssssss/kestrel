package main

import (
	"fmt"

	"github.com/gabrielssssssssss/kestrel/internal/handlers"
)

func main() {
	handler := handlers.NewCompanyHandler()
	test, err := handler.HandleCompanyRequest("https://recherche-entreprises.api.gouv.fr/search?q=SBSI&departement=22")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test)
}
