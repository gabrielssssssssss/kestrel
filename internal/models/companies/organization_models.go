package models

type Organization struct {
	Siren                       string        `json:"siren"`
	NomComplet                  string        `json:"nom_complet"`
	NomRaisonSociale            string        `json:"nom_raison_sociale"`
	Sigle                       *string       `json:"sigle"`
	NombreEtablissements        int           `json:"nombre_etablissements"`
	NombreEtablissementsOuverts int           `json:"nombre_etablissements_ouverts"`
	Siege                       Siege         `json:"siege"`
	ActivitePrincipale          string        `json:"activite_principale"`
	CategorieEntreprise         *string       `json:"categorie_entreprise"`
	CaractereEmployeur          *string       `json:"caractere_employeur"`
	AnneeCategorieEntreprise    *string       `json:"annee_categorie_entreprise"`
	DateCreation                string        `json:"date_creation"`
	DateDebutActivite           string        `json:"date_debut_activite"`
	DateFermeture               *string       `json:"date_fermeture"`
	DateMiseAJour               string        `json:"date_mise_a_jour"`
	DateMiseAJourInsee          string        `json:"date_mise_a_jour_insee"`
	DateMiseAJourRNE            *string       `json:"date_mise_a_jour_rne"`
	Dirigeants                  []Dirigeant   `json:"dirigeants"`
	EtatAdministratif           string        `json:"etat_administratif"`
	NatureJuridique             string        `json:"nature_juridique"`
	SectionActivitePrincipale   string        `json:"section_activite_principale"`
	TrancheEffectifSalarie      string        `json:"tranche_effectif_salarie"`
	AnneeTrancheEffectifSalarie *string       `json:"annee_tranche_effectif_salarie"`
	StatutDiffusion             string        `json:"statut_diffusion"`
	MatchingEtablissements      []interface{} `json:"matching_etablissements"`
	Finances                    interface{}   `json:"finances"`
	Complements                 Complements   `json:"complements"`
}

type Siege struct {
	ActivitePrincipale               string      `json:"activite_principale"`
	ActivitePrincipaleRegistreMetier *string     `json:"activite_principale_registre_metier"`
	AnneeTrancheEffectifSalarie      *string     `json:"annee_tranche_effectif_salarie"`
	Adresse                          string      `json:"adresse"`
	CaractereEmployeur               string      `json:"caractere_employeur"`
	Cedex                            *string     `json:"cedex"`
	CodePaysEtranger                 *string     `json:"code_pays_etranger"`
	CodePostal                       string      `json:"code_postal"`
	Commune                          string      `json:"commune"`
	ComplementAdresse                *string     `json:"complement_adresse"`
	Coordonnees                      string      `json:"coordonnees"`
	DateCreation                     string      `json:"date_creation"`
	DateDebutActivite                string      `json:"date_debut_activite"`
	DateFermeture                    *string     `json:"date_fermeture"`
	DateMiseAJour                    *string     `json:"date_mise_a_jour"`
	DateMiseAJourInsee               string      `json:"date_mise_a_jour_insee"`
	Departement                      string      `json:"departement"`
	DistributionSpeciale             *string     `json:"distribution_speciale"`
	Epci                             string      `json:"epci"`
	EstSiege                         bool        `json:"est_siege"`
	EtatAdministratif                string      `json:"etat_administratif"`
	GeoAdresse                       string      `json:"geo_adresse"`
	GeoId                            string      `json:"geo_id"`
	IndiceRepetition                 string      `json:"indice_repetition"`
	Latitude                         string      `json:"latitude"`
	LibelleCedex                     *string     `json:"libelle_cedex"`
	LibelleCommune                   string      `json:"libelle_commune"`
	LibelleCommuneEtranger           *string     `json:"libelle_commune_etranger"`
	LibellePaysEtranger              *string     `json:"libelle_pays_etranger"`
	LibelleVoie                      string      `json:"libelle_voie"`
	ListeEnseignes                   []string    `json:"liste_enseignes"`
	ListeFiness                      interface{} `json:"liste_finess"`
	ListeIdBio                       interface{} `json:"liste_id_bio"`
	ListeIdCC                        interface{} `json:"liste_idcc"`
	ListeIdOrganismeFormation        interface{} `json:"liste_id_organisme_formation"`
	ListeRGE                         interface{} `json:"liste_rge"`
	ListeUAI                         interface{} `json:"liste_uai"`
	Longitude                        string      `json:"longitude"`
	NomCommercial                    *string     `json:"nom_commercial"`
	NumeroVoie                       string      `json:"numero_voie"`
	Region                           string      `json:"region"`
	Siret                            string      `json:"siret"`
	StatutDiffusionEtablissement     string      `json:"statut_diffusion_etablissement"`
	TrancheEffectifSalarie           string      `json:"tranche_effectif_salarie"`
	TypeVoie                         string      `json:"type_voie"`
}

type Dirigeant struct {
	Nom              string  `json:"nom"`
	Prenoms          string  `json:"prenoms"`
	AnneeDeNaissance string  `json:"annee_de_naissance"`
	DateDeNaissance  string  `json:"date_de_naissance"`
	Qualite          string  `json:"qualite"`
	Nationalite      *string `json:"nationalite"`
	TypeDirigeant    string  `json:"type_dirigeant"`
}

type Complements struct {
	CollectiviteTerritoriale       *string     `json:"collectivite_territoriale"`
	ConventionCollectiveRenseignee bool        `json:"convention_collective_renseignee"`
	ListeIDCC                      interface{} `json:"liste_idcc"`
	EgaproRenseignee               bool        `json:"egapro_renseignee"`
	EstAchatsResponsables          bool        `json:"est_achats_responsables"`
	EstAlimConfiance               bool        `json:"est_alim_confiance"`
	EstAssociation                 bool        `json:"est_association"`
	EstBio                         bool        `json:"est_bio"`
	EstEntrepreneurIndividuel      bool        `json:"est_entrepreneur_individuel"`
	EstEntrepreneurSpectacle       bool        `json:"est_entrepreneur_spectacle"`
	EstESS                         bool        `json:"est_ess"`
	EstFiness                      bool        `json:"est_finess"`
	EstOrganismeFormation          bool        `json:"est_organisme_formation"`
	EstQualiopi                    bool        `json:"est_qualiopi"`
	ListeIDOrganismeFormation      interface{} `json:"liste_id_organisme_formation"`
	EstRGE                         bool        `json:"est_rge"`
	EstServicePublic               bool        `json:"est_service_public"`
	EstL1003                       bool        `json:"est_l100_3"`
	EstSIAE                        bool        `json:"est_siae"`
	EstSocieteMission              bool        `json:"est_societe_mission"`
	EstUAI                         bool        `json:"est_uai"`
	EstPatrimoineVivant            bool        `json:"est_patrimoine_vivant"`
	BilanGESRenseigne              bool        `json:"bilan_ges_renseigne"`
	IdentifiantAssociation         *string     `json:"identifiant_association"`
	StatutEntrepreneurSpectacle    *string     `json:"statut_entrepreneur_spectacle"`
	TypeSIAE                       *string     `json:"type_siae"`
}
