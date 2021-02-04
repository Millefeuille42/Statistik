package main

import "time"

func extractDesiredBest(ret ScoreSort, bestMap map[string]float64) ScoreSort {
	ret.Best["runheadshots"] = bestMap["runheadshots"]
	ret.Best["roomscleared"] = bestMap["roomscleared"]
	ret.Best["runtriplehits"] = bestMap["runtriplehits"]
	ret.Best["rundoublehits"] = bestMap["rundoublehits"]
	ret.Best["runheadshotrate"] = bestMap["runheadshotrate"]
	ret.Best["runperfectreload"] = bestMap["runperfectreload"]
	ret.Best["runcriticalhits"] = bestMap["runcriticalhits"]
	ret.Best["runpropsdamage"] = bestMap["runpropsdamage"]
	ret.Best["rundamageinc"] = bestMap["rundamageinc"]
	ret.Best["runreloadfails"] = bestMap["runreloadfails"]
	ret.Best["runhealing"] = bestMap["runhealing"]
	ret.Best["runhits"] = bestMap["runhits"]
	ret.Best["runexperiencetotal"] = bestMap["runexperiencetotal"]
	ret.Best["runscore"] = bestMap["runscore"]
	ret.Best["runitemsused"] = bestMap["runitemsused"]
	ret.Best["runcards"] = bestMap["runcards"]
	ret.Best["runbesthit"] = bestMap["runbesthit"]
	ret.Best["runaccuracy"] = bestMap["runaccuracy"]
	ret.Best["runperfectreloadrate"] = bestMap["runperfectreloadrate"]
	ret.Best["rundamagedone"] = bestMap["rundamagedone"]
	ret.Best["runobjectsinteracted"] = bestMap["runobjectsinteracted"]
	ret.Best["runitemsdropped"] = bestMap["runitemsdropped"]
	ret.Best["runkills"] = bestMap["runkills"]
	ret.Best["runcurrency"] = bestMap["runcurrency"]
	ret.Best["rundatatotal"] = bestMap["rundatatotal"]
	ret.Best["runcratesdestroyed"] = bestMap["runcratesdestroyed"]
	ret.Best["runshots"] = bestMap["runshots"]

	return ret
}

func extractDesiredTotal(ret ScoreSort, totalMap map[string]float64) ScoreSort {
	ret.Total["weaponsdropped"] = totalMap["weaponsdropped"]
	ret.Total["perfectreload"] = totalMap["perfectreload"]
	ret.Total["roomscleared"] = totalMap["roomscleared"]
	ret.Total["bulletshit"] = totalMap["bulletshit"]
	ret.Total["looproomscleared"] = totalMap["looproomscleared"]
	ret.Total["questscompleted"] = totalMap["questscompleted"]
	ret.Total["itemsdropped"] = totalMap["itemsdropped"]
	ret.Total["reloads"] = totalMap["reloads"]
	ret.Total["headshots"] = totalMap["headshots"]
	ret.Total["runamount"] = totalMap["runamount"]
	ret.Total["itemsused"] = totalMap["itemsused"]
	ret.Total["damagetaken"] = totalMap["damagetaken"]
	ret.Total["propsdamage"] = totalMap["propsdamage"]
	ret.Total["currency"] = totalMap["currency"]
	ret.Total["keycardsfound"] = totalMap["keycardsfound"]
	ret.Total["endbosskills"] = totalMap["endbosskills"]
	ret.Total["playerdeaths"] = totalMap["playerdeaths"]
	ret.Total["triplehits"] = totalMap["triplehits"]
	ret.Total["criticals"] = totalMap["criticals"]
	ret.Total["experience"] = totalMap["experience"]
	ret.Total["shotsfired"] = totalMap["shotsfired"]
	ret.Total["data"] = totalMap["data"]
	ret.Total["reloadfails"] = totalMap["reloadfails"]
	ret.Total["damagedone"] = totalMap["damagedone"]
	ret.Total["bosskills"] = totalMap["bosskills"]
	ret.Total["objectsinteracted"] = totalMap["objectsinteracted"]
	ret.Total["healing"] = totalMap["healing"]
	ret.Total["besthit"] = totalMap["besthit"]
	ret.Total["enemykills"] = totalMap["enemykills"]

	return ret
}

func sortScore(bestMap, totalMap map[string]float64) ScoreSort {
	ret := ScoreSort{}
	ret.Best = make(map[string]float64)
	ret.Total = make(map[string]float64)
	ret.RunDate = time.Now()

	ret = extractDesiredBest(ret, bestMap)
	ret = extractDesiredTotal(ret, totalMap)
	return ret
}