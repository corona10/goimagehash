package transforms

func forwardDCT256(input []float64) {
	var temp [256]float64
	for i := 0; i < 128; i++ {
		x, y := input[i], input[256-1-i]
		temp[i] = x + y
		temp[i+128] = (x - y) / dct256[i]
	}
	forwardDCT128(temp[:128])
	forwardDCT128(temp[128:])
	for i := 0; i < 128-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+128] + temp[i+128+1]
	}
	input[256-2], input[256-1] = temp[128-1], temp[256-1]
}

func forwardDCT128(input []float64) {
	var temp [128]float64
	for i := 0; i < 64; i++ {
		x, y := input[i], input[128-1-i]
		temp[i] = x + y
		temp[i+64] = (x - y) / dct128[i]
	}
	forwardDCT64(temp[:64])
	forwardDCT64(temp[64:])
	for i := 0; i < 64-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+64] + temp[i+64+1]
	}
	input[128-2], input[128-1] = temp[64-1], temp[128-1]
}

// forwardDCT64 function returns result of DCT-II.
// DCT type II, unscaled. Algorithm by Byeong Gi Lee, 1984.
// Static implementation by Evan Oberholster, 2022.
func forwardDCT64(input []float64) {
	var temp [64]float64
	for i := 0; i < 32; i++ {
		x, y := input[i], input[63-i]
		temp[i] = x + y
		temp[i+32] = (x - y) / dct64[i]
	}
	forwardDCT32(temp[:32])
	forwardDCT32(temp[32:])
	for i := 0; i < 32-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+32] + temp[i+32+1]
	}
	input[62], input[63] = temp[31], temp[63]
}

func forwardDCT32(input []float64) {
	var temp [32]float64
	for i := 0; i < 16; i++ {
		x, y := input[i], input[31-i]
		temp[i] = x + y
		temp[i+16] = (x - y) / dct32[i]
	}
	forwardDCT16(temp[:16])
	forwardDCT16(temp[16:])
	for i := 0; i < 16-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+16] + temp[i+16+1]
	}

	input[30], input[31] = temp[15], temp[31]
}

func forwardDCT16(input []float64) {
	var temp [16]float64
	for i := 0; i < 8; i++ {
		x, y := input[i], input[15-i]
		temp[i] = x + y
		temp[i+8] = (x - y) / dct16[i]
	}
	forwardDCT8(temp[:8])
	forwardDCT8(temp[8:])
	for i := 0; i < 8-1; i++ {
		input[i*2+0] = temp[i]
		input[i*2+1] = temp[i+8] + temp[i+8+1]
	}

	input[14], input[15] = temp[7], temp[15]
}

func forwardDCT8(input []float64) {
	var a, b = [4]float64{}, [4]float64{}

	x0, y0 := input[0], input[7]
	x1, y1 := input[1], input[6]
	x2, y2 := input[2], input[5]
	x3, y3 := input[3], input[4]

	a[0] = x0 + y0
	a[1] = x1 + y1
	a[2] = x2 + y2
	a[3] = x3 + y3
	b[0] = (x0 - y0) / 1.9615705608064609
	b[1] = (x1 - y1) / 1.6629392246050907
	b[2] = (x2 - y2) / 1.1111404660392046
	b[3] = (x3 - y3) / 0.3901806440322566

	forwardDCT4(a[:])
	forwardDCT4(b[:])

	input[0] = a[0]
	input[1] = b[0] + b[1]
	input[2] = a[1]
	input[3] = b[1] + b[2]
	input[4] = a[2]
	input[5] = b[2] + b[3]
	input[6] = a[3]
	input[7] = b[3]
}

func forwardDCT4(input []float64) {
	x0, y0 := input[0], input[3]
	x1, y1 := input[1], input[2]

	t0 := x0 + y0
	t1 := x1 + y1
	t2 := (x0 - y0) / 1.8477590650225735
	t3 := (x1 - y1) / 0.7653668647301797

	x, y := t0, t1
	t0 += t1
	t1 = (x - y) / 1.4142135623730951

	x, y = t2, t3
	t2 += t3
	t3 = (x - y) / 1.4142135623730951

	input[0] = t0
	input[1] = t2 + t3
	input[2] = t1
	input[3] = t3
}

// Static DCT Tables
var (
	//for i := 0; i < len(dct256); i++ {
	//	dct256[i] = (math.Cos((float64(i)+0.5)*math.Pi/float64(256)) * 2)
	//}
	dct256 = [128]float64{
		1.9999623505652022, 1.9996611635916468, 1.9990588350021863, 1.9981554555052907, 1.9969511611465895, 1.9954461332883833, 1.9936405985823316, 1.9915348289353196,
		1.9891291414685108, 1.986423898469589, 1.983419507338199, 1.9801164205245942, 1.976515135461499, 1.9726161944891973, 1.968420184773858, 1.9639277382191105,
		1.9591395313708813, 1.9540562853155086, 1.9486787655711517, 1.9430077819725036, 1.9370441885488345, 1.9307888833953788, 1.9242428085380832, 1.917406949791743,
		1.9102823366115416, 1.9028700419380167, 1.8951711820354824, 1.8871869163239208, 1.8789184472043798, 1.8703670198778952, 1.8615339221579674, 1.8524204842766228,
		1.843028078684084, 1.8333581198420854, 1.8234120640108598, 1.8131914090298307, 1.802697694092044, 1.7919324995123704, 1.7808974464895158, 1.7695941968618756,
		1.758024452857267, 1.7461899568365802, 1.7340924910313853, 1.7217338772755346, 1.709115976730801, 1.6962406896065945, 1.6831099548737969, 1.6697257499727602,
		1.6560900905155114, 1.6422050299822093, 1.6280726594118968, 1.6136951070875987, 1.59907453821581, 1.5842131546004248, 1.5691131943111505, 1.5537769313464649,
		1.5382066752911592, 1.5224047709685236, 1.506373598087225, 1.490115570882932, 1.4736331377547398, 1.4569287808964504, 1.4400050159227635, 1.4228643914904329,
		1.4055094889144508, 1.387942921779308, 1.370167335545401, 1.352185407150632, 1.3339998446072752, 1.3156133865941575, 1.297028802044225, 1.2782488897275517,
		1.2592764778298542, 1.2401144235265784, 1.220765612552619, 1.201232958767738, 1.1815194037177486, 1.1616279161915293, 1.1415614917739347, 1.121323152394672,
		1.1009159458732098, 1.080342945459786, 1.0596072493725897, 1.0387119803311793, 1.0176602850862142, 0.9964553339455638, 0.9751003202968722, 0.9535984601266445,
		0.9319529915359323, 0.9101671742526877, 0.8882442891408585, 0.866187637706304, 0.8440005415995996, 0.8216863421158078, 0.7992483996912936, 0.7766900933976526,
		0.7540148204328366, 0.7312259956095479, 0.708327050840981, 0.6853214346239888, 0.6622126115197529, 0.6390040616320315, 0.61569928008307, 0.5923017764872479,
		0.5688150744225436, 0.545242710899898, 0.5215882358305511, 0.4978552114914405, 0.4740472119887347, 0.45016782271958555, 0.4262206398321827, 0.40220926968418386,
		0.37813732829961255, 0.3540084408242977, 0.3298262409799401, 0.3055943705168868, 0.2813164786656985, 0.25699622158758645, 0.23263726182380973, 0.20824326774410945,
		0.1838179129942654, 0.15936487594286025, 0.1348878391273282, 0.11039048869938006, 0.08587651386988192, 0.06134960635327317, 0.03681345981160964, 0.012271769298309032}
	dct128 = [64]float64{
		1.9998494036782892, 1.9986447691766989, 1.9962362258002984, 1.992625224365556, 1.9878139400047121, 1.98180527085556, 1.9746028363157169, 1.9662109748624328,
		1.9566347414392553, 1.9458799044111204, 1.9339529420897041, 1.9208610388311316, 1.9066120807083875, 1.8912146507610426, 1.87467802382515, 1.8570121609464312,
		1.8382277033801153, 1.818335966181045, 1.7973489313879076, 1.7752792408057079, 1.7521401883908132, 1.7279457122431734, 1.7027103862105304, 1.6764494111096762,
		1.6491786055700506, 1.6209143965051895, 1.5916738092177671, 1.561474457144189, 1.530334531244918, 1.4982727890469187, 1.4653085433448259, 1.4314616505676372,
		1.3967524988179458, 1.3612019955909063, 1.3248315551803436, 1.287663085779583, 1.249718976284773, 1.211022082808651, 1.1715957149128777, 1.1314636215672265,
		1.0906499768440931, 1.049179365356938, 1.0070767674514352, 0.9643675441582458, 0.92107742191648, 0.8772324770770554, 0.8328591201952746, 0.7879840801220962,
		0.7426343879036752, 0.696837360498869, 0.650620584324526, 0.6040118986384564, 0.5570393787701061, 0.5097313192090293, 0.4621162165613425, 0.4142227523844371,
		0.36607977591028207, 0.3177162866677228, 0.26916141701425245, 0.22044441458776637, 0.17159462468887976, 0.12264147260441731, 0.07361444588271798, 0.02454307657143989}
	dct64 = [32]float64{
		1.9993976373924083, 1.9945809133573804, 1.9849590691974202, 1.9705552847778824, 1.9514042600770571, 1.9275521315908797, 1.8990563611860733, 1.8659855976694777,
		1.8284195114070614, 1.7864486023910306, 1.7401739822174227, 1.6897071304994142, 1.6351696263031674, 1.5766928552532127, 1.5144176930129691, 1.448494165902934,
		1.3790810894741339, 1.3063456859075537, 1.2304631811612539, 1.151616382835691, 1.0699952397741948, 0.9857963844595683, 0.8992226593092132, 0.8104826280099796,
		0.7197900730699766, 0.627363480797783, 0.5334255149497968, 0.43820248031373954, 0.3419237775206027, 0.24482135039843256, 0.1471291271993349, 0.049082457045824535,
	}
	dct32 = [16]float64{
		1.9975909124103448, 1.978353019929562, 1.9400625063890882, 1.8830881303660416, 1.8079785862468867, 1.7154572200005442, 1.6064150629612899, 1.4819022507099182,
		1.3431179096940369, 1.191398608984867, 1.0282054883864435, 0.8551101868605644, 0.6737797067844401, 0.48596035980652796, 0.2934609489107235, 0.09813534865483627,
	}
	dct16 = [8]float64{
		1.9903694533443936, 1.9138806714644176, 1.76384252869671, 1.546020906725474, 1.2687865683272912, 0.9427934736519956, 0.5805693545089246, 0.19603428065912154,
	}
)
