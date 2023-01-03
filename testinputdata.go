package main

func return_keywords_basic() []string {
	A := []string{"class , apa",
		"Apa,Korv, Buss",
		"Apa, Korv, Buss triggers Stick",
		"Apa, Korv, Buss triggers Stick, Stork, Stork",
		"Apa, Korv, Buss relates Stick, Stork, Stork",
		"Apa, Korv, Buss >- Stick, Stork, Stork",
		"Apa, Korv, Buss requires Stick, Stork, Stork",
		"Apa, Korv, Buss -> Stick, Stork, Stork",
		"zoomin Apa",
		"subroutine Apa",
		"class Apa",
		"Arne as Bosse",
		"Arne is Bosse",
		"Arne is Bosse, Nisse, Klas",
		"call Nisse as Arne",
		"Arne -> Bosse -> Nisse >- Klas triggers Lars"}
	return A
}

func return_keywords_advanced() []string {
	A := []string{"call Nisse as Arne",
		"Arne::Nisse >- Klas",
		"Arne::Nisse >- Klas::Gunnar",
		"Arne::Nisse::Kjell, Greta >- Klas::Gunnar",
		"subroutine Arne",
		"#pragma Kalle",
		"Arne, Britt.Gun -> Bertil",
		"Arne, Britt.Gun -> Bertil, Gert.Klas.Kalle_Korv"}
	return A
}

func return_keywords() []string {
	A := []string{"Arne -> Bertil -> Klas",
		"Arne_Anka, Britt.Gun -> Bertil",
		"Arne_Weise, Britt.Gun -> Bertil_Bajskorv, Gert.Klas.Kalle_Korv",
		"Kjell, Lenny, Claes triggers Arne",
		"Sol_Gun -< Sune"}

	return A
}
