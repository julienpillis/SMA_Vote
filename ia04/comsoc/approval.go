package comsoc

import (
	"errors"
)

func ApprovalSWF(p Profile, thresholds []int) (count Count, err error) {
	if len(p) == 0 {
		return nil, errors.New("profil is empty")
	}
	err = CheckProfileAlternative(p, p[0])
	if err != nil {
		return nil, err
	}

	if len(thresholds) != len(p) {
		return nil, errors.New("the number of thresholds doesn't match the profile")
	}

	count = make(map[Alternative]int)
	// Initialisation des décomptes à 0
	for _, alt := range p[0] {
		count[alt] = 0
	}

	for index_profile, pref := range p {
		if thresholds[index_profile] > len(pref) {
			return nil, errors.New("the thresholds exceeds the preference length")
		}
		for _, key := range pref[:thresholds[index_profile]] {
			// On itère uniquement entre l'indice 0 et le seuil associé (indice exclu)
			_, exist := count[key]
			if exist {
				count[key]++
			} else {
				count[key] = 1
			}
		}
	}
	return count, nil
}
func ApprovalSCF(p Profile, thresholds []int) (bestAlts []Alternative, err error) {
	var count Count
	count, err = ApprovalSWF(p, thresholds)
	if err != nil {
		return nil, err
	}
	return maxCount(count), err
}
