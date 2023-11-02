package comsoc


func SWFFactory(swf func(p Profile) (Count, error), tieb func([]Alternative) (Alternative, error)) func(Profile) ([]Alternative, error) {
	/* Idée : 
	* On récupère le maxCount
	* On tri en fonction de orderedAlts
	* On ajoute les alts triées dans un nouveau tableau
	* On les supprime de l'ancien tableau
	* On recommence sur le tableau restant
	* à la fin on obtient un tableau des alternatives triées par l'ordre donné, en cas d'égalite.
	 */
	f_swf := func(p Profile) ([]Alternative, error) {
		// Construction de la fonction avec application tiebreak
		count, err := swf(p)
		// Récupération du décompte
		if err != nil {
			return nil, err
		}

		sorted_alts := make([]Alternative, 0)
		for len(count) > 0 {
			// Récupération des meilleurs alternatives
			alts := maxCount(count)
			for len(alts) > 0 {
				// Tri des meilleures alternatives en fonction du tiebreak
				alt, err := tieb(alts)
				if err != nil {
					return nil, err
				}
				index := rank(alt, alts)
				alts = Remove(alts, index)
				sorted_alts = append(sorted_alts, alt)

				// suppression des alts étudiées de count
				delete(count,alt)
			}
		}
		return sorted_alts, nil
	}
	return f_swf
}

func SCFFactory(scf func(p Profile) ([]Alternative, error), tieb func([]Alternative) (Alternative, error)) func(Profile) (Alternative, error) {
	f_scf := func(p Profile) (Alternative, error) {
		// Construction de la fonction avec application tiebreak
		alts, err := scf(p)
		// Récupération du décompte
		if err != nil {
			return 0, err
		}
		alt, err := tieb(alts)
		// Application du tiebreak sur les objets ayant le plus de voix
		if err != nil {
			return 0, err
		}
		return alt, nil
	}
	return f_scf
}



// Même fonctionnement mais lorsque swf a besoin d'un paramètre en plus (ex : STV, approval)
func SWFFactoryOptions[T any](swf func(p Profile,options []T) (Count, error), tieb func([]Alternative) (Alternative, error)) func(Profile,[]T) ([]Alternative, error) {
	/* Idée : 
	* On récupère le maxCount
	* On tri en fonction de orderedAlts
	* On ajoute les alts triées dans un nouveau tableau
	* On les supprime de l'ancien tableau
	* On recommence sur le tableau restant
	* à la fin on obtient un tableau des alternatives triées par l'ordre donné, en cas d'égalite.
	 */
	f_swf := func(p Profile,options []T) ([]Alternative, error) {
		// Construction de la fonction avec application tiebreak
		count, err := swf(p,options)
		// Récupération du décompte
		if err != nil {
			return nil, err
		}

		sorted_alts := make([]Alternative, 0)
		for len(count) > 0 {
			// Récupération des meilleurs alternatives
			alts := maxCount(count)
			for len(alts) > 0 {
				// Tri des meilleures alternatives en fonction du tiebreak
				alt, err := tieb(alts)
				if err != nil {
					return nil, err
				}
				index := rank(alt, alts)
				alts = Remove(alts, index)
				sorted_alts = append(sorted_alts, alt)

				// suppression des alts étudiées de count
				delete(count,alt)
			}
		}
		return sorted_alts, nil
	}
	return f_swf
}

func SCFFactoryOptions[T any](scf func(p Profile,options []T) ([]Alternative, error), tieb func([]Alternative) (Alternative, error)) func(Profile,[]T) (Alternative, error) {
	f_scf := func(p Profile,options []T) (Alternative, error) {
		// Construction de la fonction avec application tiebreak
		alts, err := scf(p,options)
		// Récupération du décompte
		if err != nil {
			return 0, err
		}
		alt, err := tieb(alts)
		// Application du tiebreak sur les objets ayant le plus de voix
		if err != nil {
			return 0, err
		}
		return alt, nil
	}
	return f_scf
}