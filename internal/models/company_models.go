package models

type Company struct {
	RESULTS       []CompanyResult
	TOTAL_RESULTS int `json:"total_results"`
	PAGE          int `json:"page"`
	TOTAL_PAGES   int `json:"total_pages"`
}

type CompanyResult struct {
	SIREN                         string `json:"siren"`
	NOM_COMPLET                   string `json:"nom_complet"`
	NOM_RAISON_SOCIALE            string `json:"nom_raison_sociale"`
	SIGLE                         string `json:"sigle"`
	NOMBRE_ETABLISSEMENTS         int    `json:"nombre_etablissements"`
	NOMBRE_ETABLISSEMENTS_OUVERTS int    `json:"nombre_etablissements_ouverts"`
	SIEGE                         Siege
	CATEGORIE_ENTREPRISE          string `json:"categorie_entreprise"`
	DATE_CREATION                 string `json:"date_creation"`
	DATE_DEBUT_ACTIVITE           string `json:"date_debut_activite"`
	DATE_FERMETURE                string `json:"date_fermeture"`
	DATE_MISE_A_JOUR              string `json:"date_mise_a_jour"`
	DATE_MISE_A_JOUR_INSEE        string `json:"date_mise_a_jour_insee"`
	Dirigeants                    []Dirigeant
}

type Siege struct {
	ACTIVITE_PRINCIPALE string `json:"activite_principale"`
	ADRESSE             string `json:"adresse"`
	CODE_POSTAL         string `json:"code_postal"`
	COMMUNE             string `json:"commune"`
	COMPLEMENT_ADRESSE  string `json:"complement_adresse"`
	COORDONNEES         string `json:"coordonnees"`
	DEPARTEMENT         string `json:"departement"`
	EPCI                string `json:"epci"`
	ETAT_ADMINISTRATIF  string `json:"etat_administratif"`
	GEO_ADRESSE         string `json:"geo_adresse"`
	GEO_ID              string `json:"geo_id"`
	NOM_COMMERCIAL      string `json:"nom_commercial"`
	NUMERO_VOIE         string `json:"numero_voie"`
	REGION              string `json:"region"`
	SIRET               string `json:"siret"`
}

type Dirigeant struct {
	NOM                string `json:"nom"`
	PRENOMS            string `json:"prenoms"`
	ANNEE_DE_NAISSANCE string `json:"annee_de_naissnace"`
	DATE_DE_NAISSANCE  string `json:"date_de_naissance"`
	QUALITE            string `json:"qualite"`
	NATIONALITE        string `json:"nationalite"`
	TYPE_DIRIGEANT     string `json:"type_dirigeant"`
}
