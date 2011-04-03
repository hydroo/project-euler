package main

import (
	"fmt"
)


func setSum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}


//0 = nowhere, 1 = first set, 2 = second set
func indexCombinations(n int) [][]int {
	resultMap := make(map[int][]int)
	nextResultMapKey := 0
	firstCombination := make([]int, n)

	indexCombinationsRecursive(n, firstCombination, 0, resultMap, &nextResultMapKey)

	ret := make([][]int, len(resultMap))
	for k, v := range resultMap {
		ret[k] = v
	}

	return ret
}
func indexCombinationsRecursive(n int, combination []int, currentIndex int, results map[int][]int, nextResultMapKey *int) {
	if currentIndex == n {

		has1 := false
		has2 := false

		for _, v := range combination {
			if v == 1 {
				has1 = true
			} else if v == 2 {
				has2 = true
			}
		}

		if has1 == true && has2 == true {
			results[*nextResultMapKey] = combination
			*nextResultMapKey += 1
		}

		return
	}

	for i := 0; i <= 2; i += 1 {
		combinationCopy := make([]int,len(combination))
		copy(combinationCopy, combination)
		combinationCopy[currentIndex] = i

		indexCombinationsRecursive(n, combinationCopy, currentIndex + 1, results, nextResultMapKey)
	}
}


func setSumsDifferent(s []int, combinations [][]int) bool {

	for _, combination := range combinations {

		sum1 := 0
		sum2 := 0

		c1 := 0
		c2 := 0

		for k, v := range combination {

			if v == 1 {
				c1 += 1
				sum1 += s[k]
			} else if v == 2{
				c2 += 1
				sum2 += s[k]
			}

		}

		if (c1 == c2 && sum1 == sum2) ||
				(c1 > c2 && sum1 <= sum2) ||
				(c1 < c2 && sum1 >= sum2) {
			return false
		}

	}

	return true
}


func readFile() [][]int {

	s := make([][]int, 100)

	s[1]   =  []int{81,88,75,42,87,84,86,65}
	s[2]   =  []int{157,150,164,119,79,159,161,139,158}
	s[3]   =  []int{673,465,569,603,629,592,584,300,601,599,600}
	s[4]   =  []int{90,85,83,84,65,87,76,46}
	s[5]   =  []int{165,168,169,190,162,85,176,167,127}
	s[6]   =  []int{224,275,278,249,277,279,289,295,139}
	s[7]   =  []int{354,370,362,384,359,324,360,180,350,270}
	s[8]   =  []int{599,595,557,298,448,596,577,667,597,588,602}
	s[9]   =  []int{175,199,137,88,187,173,168,171,174}
	s[10]  =  []int{93,187,196,144,185,178,186,202,182}
	s[11]  =  []int{157,155,81,158,119,176,152,167,159}
	s[12]  =  []int{184,165,159,166,163,167,174,124,83}
	s[13]  =  []int{1211,1212,1287,605,1208,1189,1060,1216,1243,1200,908,1210}
	s[14]  =  []int{339,299,153,305,282,304,313,306,302,228}
	s[15]  =  []int{94,104,63,112,80,84,93,96}
	s[16]  =  []int{41,88,82,85,61,74,83,81}
	s[17]  =  []int{90,67,84,83,82,97,86,41}
	s[18]  =  []int{299,303,151,301,291,302,307,377,333,280}
	s[19]  =  []int{55,40,48,44,25,42,41}
	s[20]  =  []int{1038,1188,1255,1184,594,890,1173,1151,1186,1203,1187,1195}
	s[21]  =  []int{76,132,133,144,135,99,128,154}
	s[22]  =  []int{77,46,108,81,85,84,93,83}
	s[23]  =  []int{624,596,391,605,529,610,607,568,604,603,453}
	s[24]  =  []int{83,167,166,189,163,174,160,165,133}
	s[25]  =  []int{308,281,389,292,346,303,302,304,300,173}
	s[26]  =  []int{593,1151,1187,1184,890,1040,1173,1186,1195,1255,1188,1203}
	s[27]  =  []int{68,46,64,33,60,58,65}
	s[28]  =  []int{65,43,88,87,86,99,93,90}
	s[29]  =  []int{83,78,107,48,84,87,96,85}
	s[30]  =  []int{1188,1173,1256,1038,1187,1151,890,1186,1184,1203,594,1195}
	s[31]  =  []int{302,324,280,296,294,160,367,298,264,299}
	s[32]  =  []int{521,760,682,687,646,664,342,698,692,686,672}
	s[33]  =  []int{56,95,86,97,96,89,108,120}
	s[34]  =  []int{344,356,262,343,340,382,337,175,361,330}
	s[35]  =  []int{47,44,42,27,41,40,37}
	s[36]  =  []int{139,155,161,158,118,166,154,156,78}
	s[37]  =  []int{118,157,164,158,161,79,139,150,159}
	s[38]  =  []int{299,292,371,150,300,301,281,303,306,262}
	s[39]  =  []int{85,77,86,84,44,88,91,67}
	s[40]  =  []int{88,85,84,44,65,91,76,86}
	s[41]  =  []int{138,141,127,96,136,154,135,76}
	s[42]  =  []int{292,308,302,346,300,324,304,305,238,166}
	s[43]  =  []int{354,342,341,257,348,343,345,321,170,301}
	s[44]  =  []int{84,178,168,167,131,170,193,166,162}
	s[45]  =  []int{686,701,706,673,694,687,652,343,683,606,518}
	s[46]  =  []int{295,293,301,367,296,279,297,263,323,159}
	s[47]  =  []int{1038,1184,593,890,1188,1173,1187,1186,1195,1150,1203,1255}
	s[48]  =  []int{343,364,388,402,191,383,382,385,288,374}
	s[49]  =  []int{1187,1036,1183,591,1184,1175,888,1197,1182,1219,1115,1167}
	s[50]  =  []int{151,291,307,303,345,238,299,323,301,302}
	s[51]  =  []int{140,151,143,138,99,69,131,137}
	s[52]  =  []int{29,44,42,59,41,36,40}
	s[53]  =  []int{348,329,343,344,338,315,169,359,375,271}
	s[54]  =  []int{48,39,34,37,50,40,41}
	s[55]  =  []int{593,445,595,558,662,602,591,297,610,580,594}
	s[56]  =  []int{686,651,681,342,541,687,691,707,604,675,699}
	s[57]  =  []int{180,99,189,166,194,188,144,187,199}
	s[58]  =  []int{321,349,335,343,377,176,265,356,344,332}
	s[59]  =  []int{1151,1255,1195,1173,1184,1186,1188,1187,1203,593,1038,891}
	s[60]  =  []int{90,88,100,83,62,113,80,89}
	s[61]  =  []int{308,303,238,300,151,304,324,293,346,302}
	s[62]  =  []int{59,38,50,41,42,35,40}
	s[63]  =  []int{352,366,174,355,344,265,343,310,338,331}
	s[64]  =  []int{91,89,93,90,117,85,60,106}
	s[65]  =  []int{146,186,166,175,202,92,184,183,189}
	s[66]  =  []int{82,67,96,44,80,79,88,76}
	s[67]  =  []int{54,50,58,66,31,61,64}
	s[68]  =  []int{343,266,344,172,308,336,364,350,359,333}
	s[69]  =  []int{88,49,87,82,90,98,86,115}
	s[70]  =  []int{20,47,49,51,54,48,40}
	s[71]  =  []int{159,79,177,158,157,152,155,167,118}
	s[72]  =  []int{1219,1183,1182,1115,1035,1186,591,1197,1167,887,1184,1175}
	s[73]  =  []int{611,518,693,343,704,667,686,682,677,687,725}
	s[74]  =  []int{607,599,634,305,677,604,603,580,452,605,591}
	s[75]  =  []int{682,686,635,675,692,730,687,342,517,658,695}
	s[76]  =  []int{662,296,573,598,592,584,553,593,595,443,591}
	s[77]  =  []int{180,185,186,199,187,210,93,177,149}
	s[78]  =  []int{197,136,179,185,156,182,180,178,99}
	s[79]  =  []int{271,298,218,279,285,282,280,238,140}
	s[80]  =  []int{1187,1151,890,593,1194,1188,1184,1173,1038,1186,1255,1203}
	s[81]  =  []int{169,161,177,192,130,165,84,167,168}
	s[82]  =  []int{50,42,43,41,66,39,36}
	s[83]  =  []int{590,669,604,579,448,599,560,299,601,597,598}
	s[84]  =  []int{174,191,206,179,184,142,177,180,90}
	s[85]  =  []int{298,299,297,306,164,285,374,269,329,295}
	s[86]  =  []int{181,172,162,138,170,195,86,169,168}
	s[87]  =  []int{1184,1197,591,1182,1186,889,1167,1219,1183,1033,1115,1175}
	s[88]  =  []int{644,695,691,679,667,687,340,681,770,686,517}
	s[89]  =  []int{606,524,592,576,628,593,591,584,296,444,595}
	s[90]  =  []int{94,127,154,138,135,74,136,141}
	s[91]  =  []int{179,168,172,178,177,89,198,186,137}
	s[92]  =  []int{302,299,291,300,298,149,260,305,280,370}
	s[93]  =  []int{678,517,670,686,682,768,687,648,342,692,702}
	s[94]  =  []int{302,290,304,376,333,303,306,298,279,153}
	s[95]  =  []int{95,102,109,54,96,75,85,97}
	s[96]  =  []int{150,154,146,78,152,151,162,173,119}
	s[97]  =  []int{150,143,157,152,184,112,154,151,132}
	s[98]  =  []int{36,41,54,40,25,44,42}
	s[99]  =  []int{37,48,34,59,39,41,40}
	s[0] =  []int{681,603,638,611,584,303,454,607,606,605,596}

	return s
}


func allCombinations(min, max int) map[int][][]int {
	ret := make(map[int][][]int)

	for i := min; i <= max; i += 1 {
		ret[i] = indexCombinations(i)
	}

	return ret
}




func main() {

	s := readFile()

	c := allCombinations(2, 12)

	sum := int64(0)

	for k, v := range s {

		if len(v) > 12 {
			panic("w00t")
		}

		special := setSumsDifferent(v, c[len(v)])
		setSum := setSum(v)
		fmt.Println(k, special, v,setSum)

		if special == true {
			sum += int64(setSum)
		}
	}

	fmt.Println(sum)
	
}


