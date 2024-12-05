package seeds

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"
)

const SEEDS = "487758422 524336848 2531594804 27107767 1343486056 124327551 1117929819 93097070 3305050822 442320425 2324984130 87604424 4216329536 45038934 1482842780 224610898 115202033 371332058 2845474954 192579859"

const SEEDS_RANGE = "487758422 524336848|2531594804 27107767|1343486056 124327551|1117929819 93097070|3305050822 442320425|2324984130 87604424|4216329536 45038934|1482842780 224610898|115202033 371332058|2845474954 192579859"

type RangeMap struct {
	DestinationStart int
	SourceStart      int
	Range            int
}

var SEEDS_TO_SOIL = []RangeMap{
	{DestinationStart: 152560994, SourceStart: 173671324, Range: 63296280},
	{DestinationStart: 22185606, SourceStart: 1527272669, Range: 123700133},
	{DestinationStart: 1331635416, SourceStart: 297391996, Range: 25501160},
	{DestinationStart: 2532562923, SourceStart: 236967604, Range: 60424392},
	{DestinationStart: 145885739, SourceStart: 854580877, Range: 6675255},
	{DestinationStart: 696074404, SourceStart: 427174664, Range: 330582271},
	{DestinationStart: 0, SourceStart: 1407203152, Range: 22185606},
	{DestinationStart: 1174684019, SourceStart: 152310608, Range: 21360716},
	{DestinationStart: 2679050525, SourceStart: 3516671293, Range: 482421092},
	{DestinationStart: 1227353908, SourceStart: 322893156, Range: 104281508},
	{DestinationStart: 1077860077, SourceStart: 757756935, Range: 96823942},
	{DestinationStart: 3161471617, SourceStart: 3999092385, Range: 22763615},
	{DestinationStart: 1509447184, SourceStart: 1650972802, Range: 82059456},
	{DestinationStart: 4170512486, SourceStart: 2802241718, Range: 1263617},
	{DestinationStart: 1196044735, SourceStart: 1495963496, Range: 31309173},
	{DestinationStart: 1357136576, SourceStart: 0, Range: 152310608},
	{DestinationStart: 4171776103, SourceStart: 2679050525, Range: 123191193},
	{DestinationStart: 1731875772, SourceStart: 1733032258, Range: 319590386},
	{DestinationStart: 282432012, SourceStart: 1001625264, Range: 405577888},
	{DestinationStart: 215857274, SourceStart: 1429388758, Range: 66574738},
	{DestinationStart: 1026656675, SourceStart: 2541783913, Range: 51203402},
	{DestinationStart: 1591506640, SourceStart: 861256132, Range: 140369132},
	{DestinationStart: 688009900, SourceStart: 2052622644, Range: 8064504},
	{DestinationStart: 3457346528, SourceStart: 2803505335, Range: 713165958},
	{DestinationStart: 3184235232, SourceStart: 4021856000, Range: 273111296},
	{DestinationStart: 2051466158, SourceStart: 2060687148, Range: 481096765},
}

var SOIL_TO_FERTILIZER = []RangeMap{
	{DestinationStart: 2587123207, SourceStart: 1612631011, Range: 14556918},
	{DestinationStart: 425896400, SourceStart: 1627187929, Range: 180219453},
	{DestinationStart: 974395525, SourceStart: 3228073255, Range: 181091940},
	{DestinationStart: 606115853, SourceStart: 2482605187, Range: 15274968},
	{DestinationStart: 621390821, SourceStart: 1843957593, Range: 329096619},
	{DestinationStart: 4269010275, SourceStart: 3749182981, Range: 25957021},
	{DestinationStart: 2601680125, SourceStart: 1807407382, Range: 36550211},
	{DestinationStart: 3749182981, SourceStart: 3775140002, Range: 91697225},
	{DestinationStart: 3337986166, SourceStart: 2411426158, Range: 71179029},
	{DestinationStart: 2306503017, SourceStart: 631361705, Range: 276943999},
	{DestinationStart: 0, SourceStart: 205465305, Range: 425896400},
	{DestinationStart: 950487440, SourceStart: 181557220, Range: 23908085},
	{DestinationStart: 1155487465, SourceStart: 1188132368, Range: 424498643},
	{DestinationStart: 1579986108, SourceStart: 2497880155, Range: 726516909},
	{DestinationStart: 2583447016, SourceStart: 3224397064, Range: 3676191},
	{DestinationStart: 3840880206, SourceStart: 3866837227, Range: 428130069},
	{DestinationStart: 3099614220, SourceStart: 2173054212, Range: 238371946},
	{DestinationStart: 2918057000, SourceStart: 0, Range: 181557220},
	{DestinationStart: 2638230336, SourceStart: 908305704, Range: 279826664},
}

var FERTILIZER_TO_WATER = []RangeMap{
	{DestinationStart: 0, SourceStart: 226390191, Range: 111676682},
	{DestinationStart: 342717440, SourceStart: 10141562, Range: 176981703},
	{DestinationStart: 3507713259, SourceStart: 629378619, Range: 187481170},
	{DestinationStart: 3695194429, SourceStart: 3739990160, Range: 60489319},
	{DestinationStart: 2747106431, SourceStart: 4155961766, Range: 11528550},
	{DestinationStart: 2120906094, SourceStart: 2661985654, Range: 77106422},
	{DestinationStart: 3983043532, SourceStart: 1233406588, Range: 239832885},
	{DestinationStart: 3045388563, SourceStart: 4287081612, Range: 7885684},
	{DestinationStart: 529840705, SourceStart: 187123265, Range: 39266926},
	{DestinationStart: 2198012516, SourceStart: 1102191084, Range: 16615097},
	{DestinationStart: 3053274247, SourceStart: 4167490316, Range: 104130064},
	{DestinationStart: 2564366436, SourceStart: 2756374717, Range: 129809025},
	{DestinationStart: 3331241613, SourceStart: 816859789, Range: 87383675},
	{DestinationStart: 3418625288, SourceStart: 3536806097, Range: 36187146},
	{DestinationStart: 1239632010, SourceStart: 3393454292, Range: 143351805},
	{DestinationStart: 1147966039, SourceStart: 1776147938, Range: 91665971},
	{DestinationStart: 1382983815, SourceStart: 2230578660, Range: 82475007},
	{DestinationStart: 4222876417, SourceStart: 1045170068, Range: 57021016},
	{DestinationStart: 3455914127, SourceStart: 2313053667, Range: 4640764},
	{DestinationStart: 2835800762, SourceStart: 3800479479, Range: 42590884},
	{DestinationStart: 4279897433, SourceStart: 1473239473, Range: 15069863},
	{DestinationStart: 3780709534, SourceStart: 2342720217, Range: 202333998},
	{DestinationStart: 1753297424, SourceStart: 3902953503, Range: 253008263},
	{DestinationStart: 1465458822, SourceStart: 1488309336, Range: 287838602},
	{DestinationStart: 3454812434, SourceStart: 2588893134, Range: 1101693},
	{DestinationStart: 3287402694, SourceStart: 2545054215, Range: 43838919},
	{DestinationStart: 2704024450, SourceStart: 1002088087, Range: 43081981},
	{DestinationStart: 629378619, SourceStart: 2589994827, Range: 71990827},
	{DestinationStart: 3189558071, SourceStart: 904243464, Range: 97844623},
	{DestinationStart: 2214627613, SourceStart: 2886183742, Range: 264407413},
	{DestinationStart: 2006305687, SourceStart: 1118806181, Range: 114600407},
	{DestinationStart: 3460554891, SourceStart: 1867813909, Range: 47158368},
	{DestinationStart: 3755683748, SourceStart: 2317694431, Range: 25025786},
	{DestinationStart: 111676682, SourceStart: 338066873, Range: 231040758},
	{DestinationStart: 519699143, SourceStart: 0, Range: 10141562},
	{DestinationStart: 2479035026, SourceStart: 3313735125, Range: 69870178},
	{DestinationStart: 2548905204, SourceStart: 4271620380, Range: 15461232},
	{DestinationStart: 2775917622, SourceStart: 3843070363, Range: 59883140},
	{DestinationStart: 3157404311, SourceStart: 2198424900, Range: 32153760},
	{DestinationStart: 864513416, SourceStart: 1914972277, Range: 283452623},
	{DestinationStart: 2758634981, SourceStart: 2739092076, Range: 17282641},
	{DestinationStart: 701369446, SourceStart: 3150591155, Range: 163143970},
	{DestinationStart: 2878391646, SourceStart: 3572993243, Range: 166996917},
	{DestinationStart: 2694175461, SourceStart: 3383605303, Range: 9848989},
}

var WATER_TO_LIGHT = []RangeMap{
	{DestinationStart: 1222332482, SourceStart: 2306154207, Range: 322881400},
	{DestinationStart: 3721269109, SourceStart: 3329751112, Range: 30895612},
	{DestinationStart: 4157109606, SourceStart: 3191893422, Range: 137857690},
	{DestinationStart: 2602036554, SourceStart: 3676681423, Range: 255279159},
	{DestinationStart: 2293973174, SourceStart: 3078247260, Range: 113646162},
	{DestinationStart: 1208052724, SourceStart: 2246358310, Range: 14279758},
	{DestinationStart: 1835200232, SourceStart: 1597610395, Range: 302070784},
	{DestinationStart: 3118371905, SourceStart: 1208052724, Range: 389557671},
	{DestinationStart: 2137271016, SourceStart: 3519979265, Range: 156702158},
	{DestinationStart: 2501933666, SourceStart: 1952285487, Range: 54586749},
	{DestinationStart: 3753874345, SourceStart: 1899681179, Range: 52604308},
	{DestinationStart: 3806478653, SourceStart: 3931960582, Range: 7636192},
	{DestinationStart: 3973340148, SourceStart: 2006872236, Range: 24436917},
	{DestinationStart: 3814114845, SourceStart: 2629035607, Range: 159225303},
	{DestinationStart: 3507929576, SourceStart: 2031309153, Range: 213339533},
	{DestinationStart: 3752164721, SourceStart: 2244648686, Range: 1709624},
	{DestinationStart: 2857315713, SourceStart: 3939596774, Range: 261056192},
	{DestinationStart: 3997777065, SourceStart: 3360646724, Range: 159332541},
	{DestinationStart: 2556520415, SourceStart: 2260638068, Range: 45516139},
	{DestinationStart: 2407619336, SourceStart: 4200652966, Range: 94314330},
	{DestinationStart: 1545213882, SourceStart: 2788260910, Range: 289986350},
}

var LIGHT_TO_TEMPERATURE = []RangeMap{
	{DestinationStart: 40645637, SourceStart: 589707204, Range: 89929230},
	{DestinationStart: 2331703372, SourceStart: 3634092968, Range: 247989827},
	{DestinationStart: 375013050, SourceStart: 880330876, Range: 388603146},
	{DestinationStart: 959059361, SourceStart: 861524497, Range: 18806379},
	{DestinationStart: 3081535248, SourceStart: 2703807387, Range: 671456921},
	{DestinationStart: 2080450790, SourceStart: 1678532701, Range: 100545991},
	{DestinationStart: 1520717011, SourceStart: 1522300302, Range: 156232399},
	{DestinationStart: 3801827392, SourceStart: 4285326304, Range: 9640992},
	{DestinationStart: 2834731556, SourceStart: 1990942492, Range: 115059472},
	{DestinationStart: 2949791028, SourceStart: 2566169187, Range: 71005535},
	{DestinationStart: 3020796563, SourceStart: 4224587619, Range: 60738685},
	{DestinationStart: 4159346802, SourceStart: 1855321998, Range: 135620494},
	{DestinationStart: 977865740, SourceStart: 0, Range: 297257755},
	{DestinationStart: 3811468384, SourceStart: 3375264308, Range: 157802817},
	{DestinationStart: 2304295289, SourceStart: 1779078692, Range: 27408083},
	{DestinationStart: 3752992169, SourceStart: 1806486775, Range: 48835223},
	{DestinationStart: 2579693199, SourceStart: 3882082795, Range: 255038357},
	{DestinationStart: 130574867, SourceStart: 328095764, Range: 236983298},
	{DestinationStart: 4081903724, SourceStart: 3533067125, Range: 77443078},
	{DestinationStart: 367558165, SourceStart: 310812857, Range: 7454885},
	{DestinationStart: 3969271201, SourceStart: 3610510203, Range: 23582765},
	{DestinationStart: 4080320433, SourceStart: 1520717011, Range: 1583291},
	{DestinationStart: 777171298, SourceStart: 679636434, Range: 181888063},
	{DestinationStart: 2180996781, SourceStart: 2106001964, Range: 56665843},
	{DestinationStart: 3992853966, SourceStart: 4137121152, Range: 87466467},
	{DestinationStart: 0, SourceStart: 1268934022, Range: 6189473},
	{DestinationStart: 30817615, SourceStart: 318267742, Range: 9828022},
	{DestinationStart: 1676949410, SourceStart: 2162667807, Range: 403501380},
	{DestinationStart: 763616196, SourceStart: 297257755, Range: 13555102},
	{DestinationStart: 6189473, SourceStart: 565079062, Range: 24628142},
	{DestinationStart: 2237662624, SourceStart: 2637174722, Range: 66632665},
}

var TEMPERATURE_TO_HUMIDITY = []RangeMap{
	{DestinationStart: 3854764317, SourceStart: 3086190444, Range: 332386294},
	{DestinationStart: 2110554705, SourceStart: 1096342109, Range: 65650849},
	{DestinationStart: 3236082645, SourceStart: 1211923153, Range: 20175736},
	{DestinationStart: 846106827, SourceStart: 1853452836, Range: 60731419},
	{DestinationStart: 596073066, SourceStart: 1972015961, Range: 100470927},
	{DestinationStart: 1202254149, SourceStart: 2647779239, Range: 280062599},
	{DestinationStart: 2176205554, SourceStart: 2107199879, Range: 422980149},
	{DestinationStart: 2943801812, SourceStart: 432921932, Range: 58618218},
	{DestinationStart: 1824616452, SourceStart: 1161992958, Range: 3317671},
	{DestinationStart: 1201112206, SourceStart: 2106057936, Range: 1141943},
	{DestinationStart: 1840640266, SourceStart: 1631373972, Range: 178403432},
	{DestinationStart: 2767591984, SourceStart: 62542954, Range: 155027747},
	{DestinationStart: 2743880612, SourceStart: 491540150, Range: 8409669},
	{DestinationStart: 62542954, SourceStart: 1273100182, Range: 24571446},
	{DestinationStart: 4208015516, SourceStart: 3772661112, Range: 86951780},
	{DestinationStart: 3477752908, SourceStart: 217570701, Range: 215351231},
	{DestinationStart: 2019043698, SourceStart: 3681150105, Range: 91511007},
	{DestinationStart: 1482316748, SourceStart: 4027597725, Range: 172205397},
	{DestinationStart: 766244356, SourceStart: 4199803122, Range: 79862471},
	{DestinationStart: 950513678, SourceStart: 665972626, Range: 23438579},
	{DestinationStart: 89279871, SourceStart: 1232098889, Range: 41001293},
	{DestinationStart: 1654522145, SourceStart: 1914184255, Range: 36130017},
	{DestinationStart: 3002420030, SourceStart: 3660808941, Range: 20341164},
	{DestinationStart: 452765307, SourceStart: 1525178112, Range: 98766078},
	{DestinationStart: 551531385, SourceStart: 840194947, Range: 44541681},
	{DestinationStart: 1827934123, SourceStart: 2072486888, Range: 12706143},
	{DestinationStart: 87114400, SourceStart: 3859612892, Range: 2165471},
	{DestinationStart: 1091551468, SourceStart: 3873351076, Range: 109560738},
	{DestinationStart: 2714852730, SourceStart: 884736628, Range: 29027882},
	{DestinationStart: 906838246, SourceStart: 1809777404, Range: 43675432},
	{DestinationStart: 760232399, SourceStart: 1297671628, Range: 6011957},
	{DestinationStart: 696543993, SourceStart: 3861778363, Range: 11572713},
	{DestinationStart: 3693104139, SourceStart: 633914109, Range: 32058517},
	{DestinationStart: 1690652162, SourceStart: 499949819, Range: 133964290},
	{DestinationStart: 151982853, SourceStart: 3418576738, Range: 142433848},
	{DestinationStart: 2922619731, SourceStart: 689411205, Range: 21182081},
	{DestinationStart: 2599185703, SourceStart: 1165310629, Range: 46612524},
	{DestinationStart: 2752290281, SourceStart: 4279665593, Range: 15301703},
	{DestinationStart: 3022761194, SourceStart: 913764510, Range: 182577599},
	{DestinationStart: 2645798227, SourceStart: 3561010586, Range: 69054503},
	{DestinationStart: 130281164, SourceStart: 1950314272, Range: 21701689},
	{DestinationStart: 973952257, SourceStart: 2530180028, Range: 117599211},
	{DestinationStart: 715546488, SourceStart: 3982911814, Range: 44685911},
	{DestinationStart: 4187150611, SourceStart: 2085193031, Range: 20864905},
	{DestinationStart: 3256258381, SourceStart: 1303683585, Range: 221494527},
	{DestinationStart: 3725162656, SourceStart: 710593286, Range: 129601661},
	{DestinationStart: 708116706, SourceStart: 1623944190, Range: 7429782},
	{DestinationStart: 3205338793, SourceStart: 3630065089, Range: 30743852},
	{DestinationStart: 294416701, SourceStart: 2927841838, Range: 158348606},
}

var HUMIDITY_TO_LOCATION = []RangeMap{
	{DestinationStart: 3745579304, SourceStart: 2724582328, Range: 81388084},
	{DestinationStart: 2201043082, SourceStart: 981698567, Range: 150857305},
	{DestinationStart: 456470998, SourceStart: 689872919, Range: 41258774},
	{DestinationStart: 2351900387, SourceStart: 2312761976, Range: 47825019},
	{DestinationStart: 497729772, SourceStart: 314502888, Range: 115122928},
	{DestinationStart: 4036836228, SourceStart: 4279991461, Range: 14975835},
	{DestinationStart: 784121118, SourceStart: 3255986437, Range: 5687937},
	{DestinationStart: 612852700, SourceStart: 0, Range: 118278993},
	{DestinationStart: 3894168411, SourceStart: 842107405, Range: 27168107},
	{DestinationStart: 3127730376, SourceStart: 4061803810, Range: 218187651},
	{DestinationStart: 1230572305, SourceStart: 1958697503, Range: 125808488},
	{DestinationStart: 304425981, SourceStart: 163566931, Range: 106757079},
	{DestinationStart: 789809055, SourceStart: 3790646012, Range: 241507252},
	{DestinationStart: 1777374448, SourceStart: 1871672054, Range: 39148062},
	{DestinationStart: 1816522510, SourceStart: 1504159194, Range: 367512860},
	{DestinationStart: 235721965, SourceStart: 621168903, Range: 68704016},
	{DestinationStart: 4051812063, SourceStart: 1480547431, Range: 14406303},
	{DestinationStart: 3826967388, SourceStart: 1228960483, Range: 58471320},
	{DestinationStart: 411183060, SourceStart: 118278993, Range: 45287938},
	{DestinationStart: 3921336518, SourceStart: 869275512, Range: 103693352},
	{DestinationStart: 4025029870, SourceStart: 2304659906, Range: 8102070},
	{DestinationStart: 2838210137, SourceStart: 1132555872, Range: 96404611},
	{DestinationStart: 2934614748, SourceStart: 1287431803, Range: 193115628},
	{DestinationStart: 1414367080, SourceStart: 3422182046, Range: 363007368},
	{DestinationStart: 3345918027, SourceStart: 1494953734, Range: 9205460},
	{DestinationStart: 2184035370, SourceStart: 2805970412, Range: 17007712},
	{DestinationStart: 2631131158, SourceStart: 2088210279, Range: 207078979},
	{DestinationStart: 1356380793, SourceStart: 784121118, Range: 57986287},
	{DestinationStart: 2399725406, SourceStart: 2493176576, Range: 231405752},
	{DestinationStart: 1031316307, SourceStart: 3261674374, Range: 160507672},
	{DestinationStart: 44178878, SourceStart: 429625816, Range: 191543087},
	{DestinationStart: 4195166765, SourceStart: 3165556554, Range: 90429883},
	{DestinationStart: 4285596648, SourceStart: 2295289258, Range: 9370648},
	{DestinationStart: 3403000874, SourceStart: 2822978124, Range: 342578430},
	{DestinationStart: 1191823979, SourceStart: 3785189414, Range: 5456598},
	{DestinationStart: 1197280577, SourceStart: 2459884848, Range: 33291728},
	{DestinationStart: 4066218366, SourceStart: 4032153264, Range: 29650546},
	{DestinationStart: 3355123487, SourceStart: 1910820116, Range: 47877387},
	{DestinationStart: 3885438708, SourceStart: 972968864, Range: 8729703},
	{DestinationStart: 0, SourceStart: 270324010, Range: 44178878},
	{DestinationStart: 4095868912, SourceStart: 2360586995, Range: 99297853},
	{DestinationStart: 4033131940, SourceStart: 2084505991, Range: 3704288},
}

func Compute(input string) int {

	seeds := strings.Split(input, " ")
	locations := make([]int, len(seeds))

	for i, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		toSoil := computeTransition(seedInt, SEEDS_TO_SOIL)
		toFertilizer := computeTransition(toSoil, SOIL_TO_FERTILIZER)
		toWater := computeTransition(toFertilizer, FERTILIZER_TO_WATER)
		toLight := computeTransition(toWater, WATER_TO_LIGHT)
		toTemperature := computeTransition(toLight, LIGHT_TO_TEMPERATURE)
		toHumidity := computeTransition(toTemperature, TEMPERATURE_TO_HUMIDITY)
		toLocation := computeTransition(toHumidity, HUMIDITY_TO_LOCATION)
		locations[i] = toLocation
	}
	slices.Sort(locations)
	return locations[0]
}

func computeTransition(input int, transitionMap []RangeMap) int {
	for _, item := range transitionMap {
		if input >= item.SourceStart && input < (item.SourceStart+item.Range) {
			diff := input - item.SourceStart
			return item.DestinationStart + diff
		}
	}
	return input
}

func computeAllInRange(seed string) int {
	seedWithRange := strings.Split(seed, " ")
	var location []int
	start, _ := strconv.Atoi(seedWithRange[0])
	amount, _ := strconv.Atoi(seedWithRange[1])
	fmt.Printf("Start: %d, end %d\n", start, start+amount)
	for i := start; i < (start + amount); i++ {
		toSoil := computeTransition(i, SEEDS_TO_SOIL)
		toFertilizer := computeTransition(toSoil, SOIL_TO_FERTILIZER)
		toWater := computeTransition(toFertilizer, FERTILIZER_TO_WATER)
		toLight := computeTransition(toWater, WATER_TO_LIGHT)
		toTemperature := computeTransition(toLight, LIGHT_TO_TEMPERATURE)
		toHumidity := computeTransition(toTemperature, TEMPERATURE_TO_HUMIDITY)
		toLocation := computeTransition(toHumidity, HUMIDITY_TO_LOCATION)
		location = append(location, toLocation)
	}
	slices.Sort(location)
	v := location[0]
	println(v)
	return v
}

func ComputeRange() int {
	seedsWithRange := strings.Split(SEEDS_RANGE, "|")
	results := make(chan int)
	var wg sync.WaitGroup

	var output []int

	for i, seed := range seedsWithRange {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("starting goro %d\n", i)
			results <- computeAllInRange(seed)
		}(i)
	}
	go func() {
		wg.Wait()
		defer close(results)
	}()

	for result := range results {
		output = append(output, result)
	}
	slices.Sort(output)
	return output[0]
}

func assingMin(actual, in int) int {
	if actual == -1 {
		return in
	}
	if actual > in {
		return in
	}
	return actual
}
