package main

import (
	"fmt"

	"github.com/gabrielssssssssss/kestrel/internal/repository"
)

func main() {
	test, err := repository.FetchRechercheEntreprise("https://recherche-entreprises.api.gouv.fr/search?q=SBSI&departement=22")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(test.RESULTS[0].SIEGE.ADRESSE)
}
