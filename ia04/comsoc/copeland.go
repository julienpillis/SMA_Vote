package comsoc

func CountIsPref(p Profile) map[AltTuple]int {
	// Compte le nombre de fois que les alternatives du profil battent les autre alternatives
	// La fonction retourne un dictionnaire dont les clés sont des tuples (a,b) (" a bat b "), et les valeurs le nombre de fois que cela arrive.
	win := make(map[AltTuple]int) // enregistre le nombre de fois où a bat b
	for _, pref := range p {
		for index, alt := range pref {
			if alt == pref[len(pref)-1] {
				// On stop si on arrive à la dernière valeur (inutile de l'étudier)
				break
			} else {
				for _, alt2 := range pref[index+1:] {
					tuple := AltTuple{alt, alt2}
					_, exist := win[tuple]
					if exist {
						win[tuple]++
					} else {
						win[tuple] = 1
					}
				}
			}
		}
	}
	return win
}

func CopelandSWF(p Profile) (count Count, err error) {
	err = CheckProfile(p) // à voir si on utilise CheckProfileAlternative()
	if err != nil {
		return nil, err
	}
	count = make(Count)
	count_duels := CountIsPref(p)

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

func CopelandSCF(p Profile) (bestAlts []Alternative, err error) {
	var count Count
	count, err = CopelandSWF(p)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
