// Euler 13 http://projecteuler.net/problem=13
// by Daniel Rodríguez http://sadasant.com

package main

import "fmt"

func main() {
  numbers := [100]float64{
		3.7107287533902102798797998220837590246510135740250,
		4.6376937677490009712648124896970078050417018260538,
		7.4324986199524741059474233309513058123726617309629,
		9.1942213363574161572522430563301811072406154908250,
		2.3067588207539346171171980310421047513778063246676,
		8.9261670696623633820136378418383684178734361726757,
		2.8112879812849979408065481931592621691275889832738,
		4.4274228917432520321923589422876796487670272189318,
		4.7451445736001306439091167216856844588711603153276,
		7.0386486105843025439939619828917593665686757934951,
		6.2176457141856560629502157223196586755079324193331,
		6.4906352462741904929101432445813822663347944758178,
		9.2575867718337217661963751590579239728245598838407,
		5.8203565325359399008402633568948830189458628227828,
		8.0181199384826282014278194139940567587151170094390,
		3.5398664372827112653829987240784473053190104293586,
		8.6515506006295864861532075273371959191420517255829,
		7.1693888707715466499115593487603532921714970056938,
		5.4370070576826684624621495650076471787294438377604,
		5.3282654108756828443191190634694037855217779295145,
		3.6123272525000296071075082563815656710885258350721,
		4.5876576172410976447339110607218265236877223636045,
		1.7423706905851860660448207621209813287860733969412,
		8.1142660418086830619328460811191061556940512689692,
		5.1934325451728388641918047049293215058642563049483,
		6.2467221648435076201727918039944693004732956340691,
		1.5732444386908125794514089057706229429197107928209,
		5.5037687525678773091862540744969844508330393682126,
		1.8336384825330154686196124348767681297534375946515,
		8.0386287592878490201521685554828717201219257766954,
		7.8182833757993103614740356856449095527097864797581,
		1.6726320100436897842553539920931837441497806860984,
		4.8403098129077791799088218795327364475675590848030,
		8.7086987551392711854517078544161852424320693150332,
		5.9959406895756536782107074926966537676326235447210,
		6.9793950679652694742597709739166693763042633987085,
		4.1052684708299085211399427365734116182760315001271,
		6.5378607361501080857009149939512557028198746004375,
		3.5829035317434717326932123578154982629742552737307,
		9.4953759765105305946966067683156574377167401875275,
		8.8902802571733229619176668713819931811048770190271,
		2.5267680276078003013678680992525463401061632866526,
		3.6270218540497705585629946580636237993140746255962,
		2.4074486908231174977792365466257246923322810917141,
		9.1430288197103288597806669760892938638285025333403,
		3.4413065578016127815921815005561868836468420090470,
		2.3053081172816430487623791969842487255036638784583,
		1.1487696932154902810424020138335124462181441773470,
		6.3783299490636259666498587618221225225512486764533,
		6.7720186971698544312419572409913959008952310058822,
		9.5548255300263520781532296796249481641953868218774,
		7.6085327132285723110424803456124867697064507995236,
		3.7774242535411291684276865538926205024910326572967,
		2.3701913275725675285653248258265463092207058596522,
		2.9798860272258331913126375147341994889534765745501,
		1.8495701454879288984856827726077713721403798879715,
		3.8298203783031473527721580348144513491373226651381,
		3.4829543829199918180278916522431027392251122869539,
		4.0957953066405232632538044100059654939159879593635,
		2.9746152185502371307642255121183693803580388584903,
		4.1698116222072977186158236678424689157993532961922,
		6.2467957194401269043877107275048102390895523597457,
		2.3189706772547915061505504953922979530901129967519,
		8.6188088225875314529584099251203829009407770775672,
		1.1306739708304724483816533873502340845647058077308,
		8.2959174767140363198008187129011875491310547126581,
		9.7623331044818386269515456334926366572897563400500,
		4.2846280183517070527831839425882145521227251250327,
		5.5121603546981200581762165212827652751691296897789,
		3.2238195734329339946437501907836945765883352399886,
		7.5506164965184775180738168837861091527357929701337,
		6.2177842752192623401942399639168044983993173312731,
		3.2924185707147349566916674687634660915035914677504,
		9.9518671430235219628894890102423325116913619626622,
		7.3267460800591547471830798392868535206946944540724,
		7.6841822524674417161514036427982273348055556214818,
		9.7142617910342598647204516893989422179826088076852,
		8.7783646182799346313767754307809363333018982642090,
		1.0848802521674670883215120185883543223812876952786,
		7.1329612474782464538636993009049310363619763878039,
		6.2184073572399794223406235393808339651327408011116,
		6.6627891981488087797941876876144230030984490851411,
		6.0661826293682836764744779239180335110989069790714,
		8.5786944089552990653640447425576083659976645795096,
		6.6024396409905389607120198219976047599490197230297,
		6.4913982680032973156037120041377903785566085089252,
		1.6730939319872750275468906903707539413042652315011,
		9.4809377245048795150954100921645863754710598436791,
		7.8639167021187492431995700641917969777599028300699,
		1.5368713711936614952811305876380278410754449733078,
		4.0789923115535562561142322423255033685442488917353,
		4.4889911501440648020369068063960672322193204149535,
		4.1503128880339536053299340368006977710650566631954,
		8.1234880673210146739058568557934581403627822703280,
		8.2616570773948327592232845941706525094512325230608,
		2.2918802058777319719839450180888072429661980811197,
		7.7158542502016545090413245809786882778948721859617,
		7.2107838435069186155435662884062257473692284509516,
		2.0849603980134001723930671666823555245252804609722,
		5.3503534226472524250874054075591789781264330331690}
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	fmt.Printf("%.0f", sum * 10000000)
}
