package comsoc

import (
	"errors"
)

/*
======================================

	  @brief :
	  'Fonction de calcul du classement (SWF) de la méthode de vote Copeland.'
	  @params :
		- 'p' : profile sur lequel appliquer la méthode
	  @returned :
	    -  'count' : le décompte des points
		- 'err' : erreur (nil si aucune erreur)

======================================
*/
func CopelandSWF(p Profile) (count Count, err error) {
	if len(p) == 0 {
		return nil, errors.New("profil is empty")
	}
	err = CheckProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	}
	count = make(Count)
	count_duels := countIsPref(p)

	for len(count_duels) > 0 {

		var tuple AltTuple
		// récupération du 1er élément de la map
		for first, _ := range count_duels {
			tuple = first
			break
		}

		value := count_duels[tuple]

		// Construction du tuple inverse
		invert_tuple := AltTuple{tuple.Second(), tuple.First()}
		_, exist := count_duels[invert_tuple]

		// Valeur si égalité
		default_first := 0
		default_second := 0
		add_first := 0
		add_second := 0

		// On vérifie si l'alternative 2 bat l'alternative 1 au moins une fois
		if !exist || value > count_duels[invert_tuple] {
			// Si l'alternative 1 bat l'alternative 2 plus de fois que l'inverse ou s'il n'existe pas, on incrémente la valeur de a
			default_first = 1
			default_second = -1
			add_first = 1
			add_second = -1

		} else if value < count_duels[invert_tuple] {
			// cas inverse
			default_first = -1
			default_second = 1
			add_first = -1
			add_second = 1
		}

		// Suppression des tuples déjà comptabilisés
		delete(count_duels, tuple)
		if exist {
			delete(count_duels, AltTuple{tuple.Second(), tuple.First()})
		}

		// On met à jour les compteurs
		_, exist = count[tuple.First()]
		if exist {
			count[tuple.First()] += add_first
		} else {
			count[tuple.First()] = default_first
		}
		_, exist = count[tuple.Second()]
		if exist {
			count[tuple.Second()] += add_second
		} else {
			count[tuple.Second()] = default_second
		}
	}

	return count, nil
}

/*
======================================

	  @brief :
	  'Fonction de calcul du gagnant(SCF) de la méthode de vote de Copeland.'
	  @params :
		- 'p' : profile sur lequel appliquer la méthode
	  @returned :
	    -  'bestAlt' : le gagnant (vide si erreur)
		- 'err' : erreur (nil si aucune erreur)

======================================
*/
func CopelandSCF(p Profile) (bestAlt []Alternative, err error) {
	var count Count
	count, err = CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
