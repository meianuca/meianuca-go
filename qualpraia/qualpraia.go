package qualpraia

import "log"

func FindByWindAndSwell(wind string, swell string) []string {
	log.Printf("Searching spots by: wind=%s & swell=%s", wind, swell)
	var spots []string

	// TODO Move this logic to service?
	if swell == "ne" { // Nordeste
		switch wind {
		case "n", "ne":
			spots = append(spots, "lagoinha-do-leste", "galheta-norte", "galheta", "joaquina", "mocambique", "mole", "santinho")
		case "s", "se":
			spots = append(spots, "barra-da-lagoa")
		case "so":
			spots = append(spots, "santinho", "caldeirao", "barra-da-lagoa")
		}
	} else if swell == "se" { // Sudeste
		switch wind {
		case "n", "ne":
			spots = append(spots, "morro-das-pedras", "novo-campeche", "lomba-do-sabao", "galheta", "galheta-norte", "joaquina", "mocambique", "mole", "pico-da-cruz", "santinho", "igrejinha", "lagoinha-do-leste")
		case "se":
			spots = append(spots, "barra-da-lagoa")
		case "s", "so":
			spots = append(spots, "riozinho", "barra-da-lagoa")
		case "o":
			spots = append(spots, "riozinho")
		case "no":
			spots = append(spots, "morro-das-pedras")
		}
	} else if swell == "s" { // Sul
		switch wind {
		case "n", "ne":
			spots = append(spots, "acores", "joaquina", "mocambique", "santinho", "mole", "solidao", "naufragados", "galheta-norte", "lagoinha-do-leste")
		case "no": // TODO Aqui não seria NO?
			spots = append(spots, "morro-das-pedras")
		case "so", "o":
			spots = append(spots, "campeche")
		}
	} else if swell == "l" { // Leste
		switch wind {
		case "n", "ne":
			spots = append(spots, "pico-da-cruz", "igrejinha", "novo-campeche", "lomba-do-sabao", "santinho")
		case "s":
			spots = append(spots, "brava", "ingleses", "matadeiro")
		case "so":
			spots = append(spots, "armacao", "brava", "caldeirao", "riozinho", "ingleses", "matadeiro")
		case "o":
			spots = append(spots, "riozinho")
		case "no":
			spots = append(spots, "morro-das-pedras")
		}
	} else if swell == "ne" { // Nordeste
		switch wind {
		case "n", "ne":
			spots = append(spots, "lagoinha-do-leste", "galheta-norte", "galheta", "joaquina", "mocambique", "mole", "santinho")
		case "s", "se":
			spots = append(spots, "barra-da-lagoa")
		case "so":
			spots = append(spots, "santinho", "caldeirao", "barra-da-lagoa")
		}
	}

	// TODO check for empty array for -> "Sei não"
	return spots
}
