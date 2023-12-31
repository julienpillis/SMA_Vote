package main

import (
	"fmt"
	"ia04/comsoc"
)

func main() {
	/*
			// Test
			tab := make([]comsoc.Alternative, 10)
			for i, _ := range tab {
				tab[i] = comsoc.Alternative(i + 1)
			}
			fmt.Println(tab)
			fmt.Print(comsoc.Rank(comsoc.Alternative(5), tab))
			mp := make(map[comsoc.Alternative]int)
			for _, k := range tab {
				mp[k] = rand.Intn(10)
			}

			fmt.Println(mp)
			fmt.Println(comsoc.MaxCount(mp))

		// TEST CHECK PROFILE
		pref1 := make([]comsoc.Alternative, 11)
		pref2 := make([]comsoc.Alternative, 11)
		pref3 := make([]comsoc.Alternative, 11)
		pref4 := make([]comsoc.Alternative, 11)
		alts := make([]comsoc.Alternative, 11)
		for i, _ := range alts {
			alts[i] = comsoc.Alternative(i + 1)

		}

		for i, _ := range pref1 {
			pref1[i] = comsoc.Alternative(11 - i)

		}
		for i, _ := range pref2 {
			pref2[i] = comsoc.Alternative(i + 1)

		}
		for i, _ := range pref3 {
			pref3[i] = comsoc.Alternative(i + 1)

		}
		for i, _ := range pref4 {
			pref4[i] = comsoc.Alternative(11 - i)
		}
		//pref4[0] = 10

		profile := make([][]comsoc.Alternative, 4)
		profile[0] = pref1
		profile[1] = pref2
		profile[2] = pref3
		profile[3] = pref4

		fmt.Println(alts)
		fmt.Println(pref1)
		fmt.Println(pref2)
		fmt.Println(pref3)
		fmt.Println(pref4)

		fmt.Println(comsoc.CheckProfileAlternative(profile, alts))

		fmt.Println(comsoc.MajoritySWF(profile))
		fmt.Println(comsoc.MajoritySCF(profile))
		fmt.Println(comsoc.BordaSCF(profile))
	*/

	/*
		// TEST BORDA
			prefs := [][]comsoc.Alternative{
				{1, 2, 3},
				{1, 2, 3},
				{2, 3, 1},
				{2, 3, 1},
			}

			res, _ := comsoc.BordaSWF(prefs)

			fmt.Println(res)


	*/
	/*
		// TEST SWF

		prefs := [][]comsoc.Alternative{
			{2, 1,3,4,5,6},
			{5,4,2,3,1,6},
			{5,2,3,4,1,6},
			{2,1,3,4,5,6},
			{2,4,3,5,1,6},
			{5,2,3,6,4,1},
		}
		//thresholds := []int{2, 1, 2, 3}

		res, _ := comsoc.SWFFactory(comsoc.BordaSWF, comsoc.TieBreakFactory([]comsoc.Alternative{1,5,4,6,3,2}))(prefs)

		fmt.Println(res)

		res2, _ := comsoc.SCFFactory(comsoc.BordaSCF, comsoc.TieBreakFactory([]comsoc.Alternative{1,5,4,6,3,2}))(prefs)
		fmt.Println(res2)
	*/

	/*
		// TEST APPROVAL

		prefs := [][]comsoc.Alternative{
			{1, 3, 2},
			{1, 2, 3},
			{2, 3, 1},
			{3, 1, 2},
		}
		thresholds := []int{2, 1, 2, 3}
		tieb := make([]comsoc.Alternative,3)
		tieb[0] = comsoc.Alternative(1)
		tieb[1] = comsoc.Alternative(2)
		tieb[2] = comsoc.Alternative(3)
		res, err := comsoc.SCFFactoryOptions(comsoc.ApprovalSCF,comsoc.TieBreakFactory(tieb))(prefs,thresholds)

		if err != nil {
			fmt.Println(err)
		}
		//if len(res) != 1 || res[0] != 1 {
		//	fmt.Println("error, 1 should be the only best Alternative")
		//}

		fmt.Println(res)
	*/
	/*
		// TEST	CONDORCET

		prefs := [][]comsoc.Alternative{
			{1, 2, 3},
			{1, 3, 2},
			{3, 2, 1},
		}

		res, err := comsoc.CondorcetWinner(prefs)

		if err != nil {
			fmt.Println(err)
		}
		//if len(res) != 1 || res[0] != 1 {
		//	fmt.Println("error, 1 should be the only best Alternative")
		//}

		fmt.Println(res)
	*/
	/*
		// TEST COPELAND
		p := [][]comsoc.Alternative{
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{1, 2, 3, 4},
			{2, 3, 4, 1},
			{2, 3, 4, 1},
			{2, 3, 4, 1},
			{2, 3, 4, 1},
			{4, 3, 1, 2},
			{4, 3, 1, 2},
			{4, 3, 1, 2},
		}

		fmt.Println(comsoc.CopelandSWF(p))
	*/
	/*
		// TEST STV
		p := [][]comsoc.Alternative{
			{1, 3, 2},
			{1, 2, 3},
			{2, 3, 1},
			{3, 1, 2},
		}
		tieb := make([]comsoc.Alternative,3)
		tieb[0] = comsoc.Alternative(1)
		tieb[1] = comsoc.Alternative(2)
		tieb[2] = comsoc.Alternative(3)
		res, _ := comsoc.SWFFactoryOptions(comsoc.STV_SWF,comsoc.TieBreakFactory(tieb))(p,tieb)

		fmt.Println(res)
	*/

	/* TEST Kramer
	p := [][]comsoc.Alternative{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
	}

	fmt.Println(comsoc.SWFFactory(comsoc.KramerSimpson_SWF, comsoc.TieBreakFactory([]comsoc.Alternative{1,2,3,4}))(p))
	fmt.Println(comsoc.SCFFactory(comsoc.KramerSimpson_SCF, comsoc.TieBreakFactory([]comsoc.Alternative{1,2,3,4}))(p))

	//fmt.Println(comsoc.KramerSimpson_SCF(p))
	*/
	// One_flip(pref []Alternative) [][]Alternative

	/* TEST DODGSON
	p := [][]comsoc.Alternative{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
	}
	//pref := p[0]
	tieb := make([]comsoc.Alternative, 4)
	tieb[0] = comsoc.Alternative(1)
	tieb[1] = comsoc.Alternative(2)
	tieb[2] = comsoc.Alternative(3)
	tieb[3] = comsoc.Alternative(4)
	fmt.Println(comsoc.Dodgson(p, tieb))
	*/

	/* TEST KEMENY */
	p := [][]comsoc.Alternative{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{2, 3, 4, 1},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
		{4, 3, 1, 2},
	}
	//pref := p[0]
	tieb := make([]comsoc.Alternative, 4)
	tieb[0] = comsoc.Alternative(1)
	tieb[1] = comsoc.Alternative(2)
	tieb[2] = comsoc.Alternative(3)
	tieb[3] = comsoc.Alternative(4)
	fmt.Println(comsoc.Kemeny(p, tieb))

}
