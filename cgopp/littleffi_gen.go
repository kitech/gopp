// oohmyffi
package cgopp

import "log"
import "github.com/ebitengine/purego"

func litfficallgenimpl[RETY any](tycrc uint64, fnptrx uintptr, args ...any) RETY {
	var rv RETY
	switch tycrc {
	case 2913135854266351616:
		var fnv func(int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int))
	case 4906417177782910976:
		var fnv func(int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64))
	case 4906417137293123632:
		var fnv func(float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64))
	case 3303529566992870192:
		var fnv func(voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr))
	case 3303529680180094784:
		var fnv func(charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr))
	case 10432896870521055341:
		var fnv func(int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64))
	case 18209049112833628269:
		var fnv func(int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int))
	case 9207394733217734689:
		var fnv func(int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64))
	case 4037525016094906468:
		var fnv func(float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int))
	case 1587618363337354017:
		var fnv func(int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr))
	case 9121461380041965219:
		var fnv func(voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int))
	case 1587618253186279249:
		var fnv func(int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr))
	case 11149571054909292221:
		var fnv func(charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int))
	case 3370978080484024353:
		var fnv func(int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64))
	case 5966082068799308823:
		var fnv func(float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64))
	case 5118191938142612257:
		var fnv func(int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr))
	case 998566109219925464:
		var fnv func(voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64))
	case 5118191914847709009:
		var fnv func(int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr))
	case 993143478043880920:
		var fnv func(charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64))
	case 5118191991506176268:
		var fnv func(float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr))
	case 8666563580025805148:
		var fnv func(voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64))
	case 5118191859880789372:
		var fnv func(float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr))
	case 8666563660625760861:
		var fnv func(charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64))
	case 1264024342512033324:
		var fnv func(voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr))
	case 1264024389360982365:
		var fnv func(charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr))
	case 2211054885833174001:
		var fnv func(int, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(float64))
	case 7110313033498808263:
		var fnv func(int, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(int64))
	case 5737617402127050481:
		var fnv func(int64, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(float64))
	case 607717147535744692:
		var fnv func(int64, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int))
	case 1294120161901266952:
		var fnv func(float64, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int))
	case 9047505918902532104:
		var fnv func(float64, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(int64))
	case 8601972540519784689:
		var fnv func(int, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(voidptr))
	case 4448640083135211016:
		var fnv func(int, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(int64))
	case 2769566881663888881:
		var fnv func(int64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(voidptr))
	case 5688838901008398451:
		var fnv func(int64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(int))
	case 12813574247508844181:
		var fnv func(voidptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(int))
	case 15972126301297159829:
		var fnv func(voidptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(int64))
	case 8601972517224885377:
		var fnv func(int, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(charptr))
	case 4444906301819430408:
		var fnv func(int, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(int64))
	case 2769566956482142593:
		var fnv func(int64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(charptr))
	case 12311746095700473965:
		var fnv func(int64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(int))
	case 757526005966985915:
		var fnv func(charptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(int))
	case 7420195229198609083:
		var fnv func(charptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(int64))
	case 10446178606185548054:
		var fnv func(int, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(voidptr))
	case 12662681584613628230:
		var fnv func(int, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(float64))
	case 13354799327382547965:
		var fnv func(float64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(voidptr))
	case 15116452588859914367:
		var fnv func(float64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int))
	case 16518331508533350824:
		var fnv func(voidptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int))
	case 11741994578002848237:
		var fnv func(voidptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(float64))
	case 10446178527173510502:
		var fnv func(int, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(charptr))
	case 12662681639443780167:
		var fnv func(int, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(float64))
	case 13354799235770818957:
		var fnv func(float64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(charptr))
	case 3883974844702226529:
		var fnv func(float64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int))
	case 6323755006755721659:
		var fnv func(charptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int))
	case 1156697823738866174:
		var fnv func(charptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(float64))
	case 14300484722002779702:
		var fnv func(int, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(charptr))
	case 14300484742008183111:
		var fnv func(int, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(voidptr))
	case 14644746126239657629:
		var fnv func(voidptr, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(charptr))
	case 5156611318325360497:
		var fnv func(voidptr, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int))
	case 1251902826478221180:
		var fnv func(charptr, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int))
	case 8773588551704961790:
		var fnv func(charptr, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(voidptr))
	case 16591803569810614550:
		var fnv func(int64, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(voidptr))
	case 15637417818073446726:
		var fnv func(int64, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(float64))
	case 4756811858747589118:
		var fnv func(float64, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(voidptr))
	case 639354777755059975:
		var fnv func(float64, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int64))
	case 16880711023353215045:
		var fnv func(voidptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int64))
	case 10828768418933611635:
		var fnv func(voidptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(float64))
	case 16591803559517529446:
		var fnv func(int64, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(charptr))
	case 15637417872900190791:
		var fnv func(int64, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(float64))
	case 4756811865953144206:
		var fnv func(float64, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(charptr))
	case 634081649793632007:
		var fnv func(float64, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int64))
	case 16887715486180870213:
		var fnv func(charptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int64))
	case 10835345686641257587:
		var fnv func(charptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(float64))
	case 12739595372406996534:
		var fnv func(int64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(charptr))
	case 12739595323696855367:
		var fnv func(int64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(voidptr))
	case 18413219247685065475:
		var fnv func(voidptr, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(charptr))
	case 13062628665344151946:
		var fnv func(voidptr, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int64))
	case 13064205157523054986:
		var fnv func(charptr, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int64))
	case 18415649580393883507:
		var fnv func(charptr, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(voidptr))
	case 3011820139625431914:
		var fnv func(float64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(charptr))
	case 3011820152885281819:
		var fnv func(float64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(voidptr))
	case 951280892671248413:
		var fnv func(voidptr, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(charptr))
	case 3634696496493666108:
		var fnv func(voidptr, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(float64))
	case 14371278049503924076:
		var fnv func(charptr, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(float64))
	case 17884815468984000316:
		var fnv func(charptr, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(voidptr))
	case 14636640998181805801:
		var fnv func(int, int64, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(float64), args[3].(voidptr))
	case 17608338590351527609:
		var fnv func(int, int64, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(voidptr), args[3].(float64))
	case 8024764925661492737:
		var fnv func(int, float64, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(int64), args[3].(voidptr))
	case 2702537769743233272:
		var fnv func(int, float64, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(voidptr), args[3].(int64))
	case 14346906268351336378:
		var fnv func(int, voidptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(float64), args[3].(int64))
	case 13481988932804848524:
		var fnv func(int, voidptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(int64), args[3].(float64))
	case 13665731125052382233:
		var fnv func(int64, int, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(float64), args[3].(voidptr))
	case 9431883262956932169:
		var fnv func(int64, int, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(voidptr), args[3].(float64))
	case 10667867730075826418:
		var fnv func(int64, float64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int), args[3].(voidptr))
	case 18200818091515047280:
		var fnv func(int64, float64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(voidptr), args[3].(int))
	case 14439315459998434471:
		var fnv func(int64, voidptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(float64), args[3].(int))
	case 10353168398714322146:
		var fnv func(int64, voidptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(int), args[3].(float64))
	case 17291256451339633324:
		var fnv func(float64, int64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int), args[3].(voidptr))
	case 9757972802585204526:
		var fnv func(float64, int64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(voidptr), args[3].(int))
	case 13761241473902989228:
		var fnv func(float64, int, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(int64), args[3].(voidptr))
	case 17592510208780945749:
		var fnv func(float64, int, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(voidptr), args[3].(int64))
	case 1455221754820116936:
		var fnv func(float64, voidptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int), args[3].(int64))
	case 8668142463658979784:
		var fnv func(float64, voidptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int64), args[3].(int))
	case 13169916958323737472:
		var fnv func(voidptr, int64, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(float64), args[3].(int))
	case 17369229911724555205:
		var fnv func(voidptr, int64, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(int), args[3].(float64))
	case 12636635850118288700:
		var fnv func(voidptr, float64, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int64), args[3].(int))
	case 14061584415883158844:
		var fnv func(voidptr, float64, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int), args[3].(int64))
	case 15854661066165932787:
		var fnv func(voidptr, int, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(float64), args[3].(int64))
	case 11530557003775997637:
		var fnv func(voidptr, int, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(int64), args[3].(float64))
	case 14636640987889242777:
		var fnv func(int, int64, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(float64), args[3].(charptr))
	case 17608338645177746872:
		var fnv func(int, int64, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(charptr), args[3].(float64))
	case 8024764932867570289:
		var fnv func(int, float64, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(int64), args[3].(charptr))
	case 2707388944850235640:
		var fnv func(int, float64, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(charptr), args[3].(int64))
	case 14340325165508444090:
		var fnv func(int, charptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(float64), args[3].(int64))
	case 13474989430930742156:
		var fnv func(int, charptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(int64), args[3].(float64))
	case 13665731149152587881:
		var fnv func(int64, int, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(float64), args[3].(charptr))
	case 9431883186791434056:
		var fnv func(int64, int, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(charptr), args[3].(float64))
	case 10667867741576344706:
		var fnv func(int64, float64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int), args[3].(charptr))
	case 1781315069896945006:
		var fnv func(int64, float64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(charptr), args[3].(int))
	case 8835037422493582516:
		var fnv func(int64, charptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(float64), args[3].(int))
	case 4418959907599596785:
		var fnv func(int64, charptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(int), args[3].(float64))
	case 17291256581891280604:
		var fnv func(float64, int64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int), args[3].(charptr))
	case 7153268243585799984:
		var fnv func(float64, int64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(charptr), args[3].(int))
	case 13761241429300495324:
		var fnv func(float64, int, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(int64), args[3].(charptr))
	case 17597502173184477525:
		var fnv func(float64, int, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(charptr), args[3].(int64))
	case 12637184640556893670:
		var fnv func(float64, charptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int), args[3].(int64))
	case 14058792887945836006:
		var fnv func(float64, charptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int64), args[3].(int))
	case 8854838495141618046:
		var fnv func(charptr, int64, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(float64), args[3].(int))
	case 4408359466827192635:
		var fnv func(charptr, int64, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(int), args[3].(float64))
	case 7167509982422579138:
		var fnv func(charptr, float64, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int64), args[3].(int))
	case 1081582468663543746:
		var fnv func(charptr, float64, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int), args[3].(int64))
	case 1162306098125877261:
		var fnv func(charptr, int, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(float64), args[3].(int64))
	case 7791932695825388603:
		var fnv func(charptr, int, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(int64), args[3].(float64))
	case 11354142551689021897:
		var fnv func(int, int64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(voidptr), args[3].(charptr))
	case 11354142502978883256:
		var fnv func(int, int64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(charptr), args[3].(voidptr))
	case 15192545180767120636:
		var fnv func(int, voidptr, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(int64), args[3].(charptr))
	case 10960039176618731125:
		var fnv func(int, voidptr, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(charptr), args[3].(int64))
	case 10958180794470356597:
		var fnv func(int, charptr, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(voidptr), args[3].(int64))
	case 15190537884544782476:
		var fnv func(int, charptr, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(int64), args[3].(voidptr))
	case 16945815057766730553:
		var fnv func(int64, int, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(voidptr), args[3].(charptr))
	case 16945815105655540808:
		var fnv func(int64, int, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(charptr), args[3].(voidptr))
	case 16601025892961778578:
		var fnv func(int64, voidptr, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(int), args[3].(charptr))
	case 7697162731422229118:
		var fnv func(int64, voidptr, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(charptr), args[3].(int))
	case 4323878858165854835:
		var fnv func(int64, charptr, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(voidptr), args[3].(int))
	case 6096794556832446449:
		var fnv func(int64, charptr, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(int), args[3].(voidptr))
	case 11008064460727827637:
		var fnv func(voidptr, int64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(int), args[3].(charptr))
	case 1474895768858453337:
		var fnv func(voidptr, int64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(charptr), args[3].(int))
	case 14540894325976543669:
		var fnv func(voidptr, int, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(int64), args[3].(charptr))
	case 9440812074358291260:
		var fnv func(voidptr, int, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(charptr), args[3].(int64))
	case 15587571887627171727:
		var fnv func(voidptr, charptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int), args[3].(int64))
	case 12981836857465096079:
		var fnv func(voidptr, charptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int64), args[3].(int))
	case 4359325981569288121:
		var fnv func(charptr, int64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(voidptr), args[3].(int))
	case 6116668206004687419:
		var fnv func(charptr, int64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(int), args[3].(voidptr))
	case 14078053248893018463:
		var fnv func(charptr, voidptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int64), args[3].(int))
	case 12620734373632364895:
		var fnv func(charptr, voidptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int), args[3].(int64))
	case 5706343994531797442:
		var fnv func(charptr, int, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(voidptr), args[3].(int64))
	case 427251834312089403:
		var fnv func(charptr, int, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(int64), args[3].(voidptr))
	case 21567107606889343:
		var fnv func(int, float64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(voidptr), args[3].(charptr))
	case 21567087580742670:
		var fnv func(int, float64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(charptr), args[3].(voidptr))
	case 2644490050888878088:
		var fnv func(int, voidptr, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(float64), args[3].(charptr))
	case 2013588879627296553:
		var fnv func(int, voidptr, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(charptr), args[3].(float64))
	case 17217741175948889977:
		var fnv func(int, charptr, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(voidptr), args[3].(float64))
	case 15110453848188378921:
		var fnv func(int, charptr, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(float64), args[3].(voidptr))
	case 5591428462975306546:
		var fnv func(float64, int, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(voidptr), args[3].(charptr))
	case 5591428405402537027:
		var fnv func(float64, int, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(charptr), args[3].(voidptr))
	case 4670706269840767897:
		var fnv func(float64, voidptr, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int), args[3].(charptr))
	case 14727561561380758133:
		var fnv func(float64, voidptr, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(charptr), args[3].(int))
	case 11147147285879333496:
		var fnv func(float64, charptr, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(voidptr), args[3].(int))
	case 17451226266765654010:
		var fnv func(float64, charptr, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int), args[3].(voidptr))
	case 8110623998828869890:
		var fnv func(voidptr, float64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int), args[3].(charptr))
	case 18175921307335762158:
		var fnv func(voidptr, float64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(charptr), args[3].(int))
	case 6423037613678616041:
		var fnv func(voidptr, int, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(float64), args[3].(charptr))
	case 7377353000288553672:
		var fnv func(voidptr, int, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(charptr), args[3].(float64))
	case 15696755683923390833:
		var fnv func(voidptr, charptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int), args[3].(float64))
	case 11392140062930045236:
		var fnv func(voidptr, charptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(float64), args[3].(int))
	case 4747637365822169321:
		var fnv func(charptr, float64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(voidptr), args[3].(int))
	case 2990202922005427563:
		var fnv func(charptr, float64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int), args[3].(voidptr))
	case 8437058978582694206:
		var fnv func(charptr, voidptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(float64), args[3].(int))
	case 3664068996608094587:
		var fnv func(charptr, voidptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int), args[3].(float64))
	case 4580287548440592848:
		var fnv func(charptr, int, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(voidptr), args[3].(float64))
	case 59350079245209984:
		var fnv func(charptr, int, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(float64), args[3].(voidptr))
	case 4905904386726946175:
		var fnv func(int64, float64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(voidptr), args[3].(charptr))
	case 4905904341233771022:
		var fnv func(int64, float64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(charptr), args[3].(voidptr))
	case 6983610325884431880:
		var fnv func(int64, voidptr, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(float64), args[3].(charptr))
	case 6892879409292949801:
		var fnv func(int64, voidptr, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(charptr), args[3].(float64))
	case 12297198937874557305:
		var fnv func(int64, charptr, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(voidptr), args[3].(float64))
	case 10802662783174179113:
		var fnv func(int64, charptr, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(float64), args[3].(voidptr))
	case 8411151335409247856:
		var fnv func(float64, int64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(voidptr), args[3].(charptr))
	case 8411151381975802113:
		var fnv func(float64, int64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(charptr), args[3].(voidptr))
	case 4323084780779258693:
		var fnv func(float64, voidptr, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int64), args[3].(charptr))
	case 8156454756155771340:
		var fnv func(float64, voidptr, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(charptr), args[3].(int64))
	case 8159162096578558412:
		var fnv func(float64, charptr, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(voidptr), args[3].(int64))
	case 4316643429593087797:
		var fnv func(float64, charptr, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int64), args[3].(voidptr))
	case 16164209521351980986:
		var fnv func(voidptr, float64, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int64), args[3].(charptr))
	case 12294242308154522931:
		var fnv func(voidptr, float64, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(charptr), args[3].(int64))
	case 4904618685210124114:
		var fnv func(voidptr, int64, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(float64), args[3].(charptr))
	case 8886197377604814963:
		var fnv func(voidptr, int64, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(charptr), args[3].(float64))
	case 3751420033708578375:
		var fnv func(voidptr, charptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int64), args[3].(float64))
	case 5191645524155396721:
		var fnv func(voidptr, charptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(float64), args[3].(int64))
	case 12292929101081589811:
		var fnv func(charptr, float64, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(voidptr), args[3].(int64))
	case 16160506474626924234:
		var fnv func(charptr, float64, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int64), args[3].(voidptr))
	case 5188820867801450353:
		var fnv func(charptr, voidptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(float64), args[3].(int64))
	case 3747473909403624263:
		var fnv func(charptr, voidptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int64), args[3].(float64))
	case 8882633989939791474:
		var fnv func(charptr, int64, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(voidptr), args[3].(float64))
	case 4902130098487105058:
		var fnv func(charptr, int64, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(float64), args[3].(voidptr))
	case 13518211429013896447:
		var fnv func(int, int64, float64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(float64), args[3].(voidptr), args[4].(charptr))
	case 13518211383519705998:
		var fnv func(int, int64, float64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(float64), args[3].(charptr), args[4].(voidptr))
	case 11485823056774641544:
		var fnv func(int, int64, voidptr, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(voidptr), args[3].(float64), args[4].(charptr))
	case 11540332128656875689:
		var fnv func(int, int64, voidptr, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(voidptr), args[3].(charptr), args[4].(float64))
	case 6136012551547171065:
		var fnv func(int, int64, charptr, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(charptr), args[3].(voidptr), args[4].(float64))
	case 7666770546045067433:
		var fnv func(int, int64, charptr, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(int64), args[2].(charptr), args[3].(float64), args[4].(voidptr))
	case 10031058105758040048:
		var fnv func(int, float64, int64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(int64), args[3].(voidptr), args[4].(charptr))
	case 10031058152325675137:
		var fnv func(int, float64, int64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(int64), args[3].(charptr), args[4].(voidptr))
	case 14155302997736987333:
		var fnv func(int, float64, voidptr, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(voidptr), args[3].(int64), args[4].(charptr))
	case 10285834258008870988:
		var fnv func(int, float64, voidptr, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(voidptr), args[3].(charptr), args[4].(int64))
	case 10282906601341395020:
		var fnv func(int, float64, charptr, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(charptr), args[3].(voidptr), args[4].(int64))
	case 14152808068663420597:
		var fnv func(int, float64, charptr, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(float64), args[2].(charptr), args[3].(int64), args[4].(voidptr))
	case 2296234699275087418:
		var fnv func(int, voidptr, float64, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(float64), args[3].(int64), args[4].(charptr))
	case 6130243096207467699:
		var fnv func(int, voidptr, float64, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(float64), args[3].(charptr), args[4].(int64))
	case 13519708316546959058:
		var fnv func(int, voidptr, int64, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(int64), args[3].(float64), args[4].(charptr))
	case 9573824665352720883:
		var fnv func(int, voidptr, int64, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(int64), args[3].(charptr), args[4].(float64))
	case 14663908424956743623:
		var fnv func(int, voidptr, charptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(charptr), args[3].(int64), args[4].(float64))
	case 13223551164464442353:
		var fnv func(int, voidptr, charptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(voidptr), args[2].(charptr), args[3].(float64), args[4].(int64))
	case 6131178390412845491:
		var fnv func(int, charptr, float64, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(float64), args[3].(voidptr), args[4].(int64))
	case 2290841978056075082:
		var fnv func(int, charptr, float64, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(float64), args[3].(int64), args[4].(voidptr))
	case 13226358206666432241:
		var fnv func(int, charptr, voidptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(voidptr), args[3].(float64), args[4].(int64))
	case 14667837002931534535:
		var fnv func(int, charptr, voidptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(voidptr), args[3].(int64), args[4].(float64))
	case 9568573802215923698:
		var fnv func(int, charptr, int64, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(int64), args[3].(voidptr), args[4].(float64))
	case 13513278255516307362:
		var fnv func(int, charptr, int64, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int), args[1].(charptr), args[2].(int64), args[3].(float64), args[4].(voidptr))
	case 18429728682821330694:
		var fnv func(int64, int, float64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(float64), args[3].(voidptr), args[4].(charptr))
	case 18429728697474769015:
		var fnv func(int64, int, float64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(float64), args[3].(charptr), args[4].(voidptr))
	case 15797517040607046769:
		var fnv func(int64, int, voidptr, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(voidptr), args[3].(float64), args[4].(charptr))
	case 16464781829439093584:
		var fnv func(int64, int, voidptr, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(voidptr), args[3].(charptr), args[4].(float64))
	case 1260629510636022528:
		var fnv func(int64, int, charptr, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(charptr), args[3].(voidptr), args[4].(float64))
	case 3331553266867288912:
		var fnv func(int64, int, charptr, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int), args[2].(charptr), args[3].(float64), args[4].(voidptr))
	case 12832916101654359883:
		var fnv func(int64, float64, int, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int), args[3].(voidptr), args[4].(charptr))
	case 12832916147411445818:
		var fnv func(int64, float64, int, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int), args[3].(charptr), args[4].(voidptr))
	case 13789667085293789152:
		var fnv func(int64, float64, voidptr, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(voidptr), args[3].(int), args[4].(charptr))
	case 3741889349941544460:
		var fnv func(int64, float64, voidptr, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(voidptr), args[3].(charptr), args[4].(int))
	case 7295141836724979201:
		var fnv func(int64, float64, charptr, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(charptr), args[3].(voidptr), args[4].(int))
	case 981826947427406723:
		var fnv func(int64, float64, charptr, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(charptr), args[3].(int), args[4].(voidptr))
	case 10304704554279279995:
		var fnv func(int64, voidptr, float64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(float64), args[3].(int), args[4].(charptr))
	case 266499202617099415:
		var fnv func(int64, voidptr, float64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(float64), args[3].(charptr), args[4].(int))
	case 12010015064784386448:
		var fnv func(int64, voidptr, int, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(int), args[3].(float64), args[4].(charptr))
	case 11091957738333662897:
		var fnv func(int64, voidptr, int, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(int), args[3].(charptr), args[4].(float64))
	case 2763248218191009032:
		var fnv func(int64, voidptr, charptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(charptr), args[3].(int), args[4].(float64))
	case 7031826265403845965:
		var fnv func(int64, voidptr, charptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(voidptr), args[2].(charptr), args[3].(float64), args[4].(int))
	case 13721883930354435216:
		var fnv func(int64, charptr, float64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(float64), args[3].(voidptr), args[4].(int))
	case 15487956131077067026:
		var fnv func(int64, charptr, float64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(float64), args[3].(int), args[4].(voidptr))
	case 9987127777136504135:
		var fnv func(int64, charptr, voidptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(voidptr), args[3].(float64), args[4].(int))
	case 14796137779360273666:
		var fnv func(int64, charptr, voidptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(voidptr), args[3].(int), args[4].(float64))
	case 13843890424003634601:
		var fnv func(int64, charptr, int, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(int), args[3].(voidptr), args[4].(float64))
	case 18401084831427737081:
		var fnv func(int64, charptr, int, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(charptr), args[2].(int), args[3].(float64), args[4].(voidptr))
	case 14615412150555934932:
		var fnv func(float64, int64, int, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int), args[3].(voidptr), args[4].(charptr))
	case 14615412137245951909:
		var fnv func(float64, int64, int, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int), args[3].(charptr), args[4].(voidptr))
	case 14383740739873550463:
		var fnv func(float64, int64, voidptr, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(voidptr), args[3].(int), args[4].(charptr))
	case 5417318324863920531:
		var fnv func(float64, int64, voidptr, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(voidptr), args[3].(charptr), args[4].(int))
	case 2161303412950067614:
		var fnv func(float64, int64, charptr, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(charptr), args[3].(voidptr), args[4].(int))
	case 8458838383970368540:
		var fnv func(float64, int64, charptr, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(charptr), args[3].(int), args[4].(voidptr))
	case 13577783978422960676:
		var fnv func(float64, int, int64, voidptr, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(int64), args[3].(voidptr), args[4].(charptr))
	case 13577783993059689813:
		var fnv func(float64, int, int64, charptr, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(int64), args[3].(charptr), args[4].(voidptr))
	case 17521603652622755601:
		var fnv func(float64, int, voidptr, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(voidptr), args[3].(int64), args[4].(charptr))
	case 13395428667070295448:
		var fnv func(float64, int, voidptr, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(voidptr), args[3].(charptr), args[4].(int64))
	case 13397287186120809880:
		var fnv func(float64, int, charptr, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(charptr), args[3].(voidptr), args[4].(int64))
	case 17523610949432482657:
		var fnv func(float64, int, charptr, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int), args[2].(charptr), args[3].(int64), args[4].(voidptr))
	case 16732100428118194776:
		var fnv func(float64, voidptr, int, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int), args[3].(int64), args[4].(charptr))
	case 11744119432118431953:
		var fnv func(float64, voidptr, int, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int), args[3].(charptr), args[4].(int64))
	case 13347401166204368728:
		var fnv func(float64, voidptr, int64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int64), args[3].(int), args[4].(charptr))
	case 3857367209927945908:
		var fnv func(float64, voidptr, int64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(int64), args[3].(charptr), args[4].(int))
	case 10797170367013162082:
		var fnv func(float64, voidptr, charptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(charptr), args[3].(int64), args[4].(int))
	case 17991265759733988450:
		var fnv func(float64, voidptr, charptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(voidptr), args[2].(charptr), args[3].(int), args[4].(int64))
	case 7984599875355808303:
		var fnv func(float64, charptr, int, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int), args[3].(voidptr), args[4].(int64))
	case 2599168174116574422:
		var fnv func(float64, charptr, int, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int), args[3].(int64), args[4].(voidptr))
	case 10293581555139680946:
		var fnv func(float64, charptr, voidptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(voidptr), args[3].(int), args[4].(int64))
	case 16330250428052351666:
		var fnv func(float64, charptr, voidptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(voidptr), args[3].(int64), args[4].(int))
	case 2125858501823024212:
		var fnv func(float64, charptr, int64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int64), args[3].(voidptr), args[4].(int))
	case 8438966947074513366:
		var fnv func(float64, charptr, int64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(charptr), args[2].(int64), args[3].(int), args[4].(voidptr))
	case 15981629090609294681:
		var fnv func(voidptr, int64, float64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(float64), args[3].(int), args[4].(charptr))
	case 5870802020257019061:
		var fnv func(voidptr, int64, float64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(float64), args[3].(charptr), args[4].(int))
	case 17611231613426485682:
		var fnv func(voidptr, int64, int, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(int), args[3].(float64), args[4].(charptr))
	case 14638106174004633235:
		var fnv func(voidptr, int64, int, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(int), args[3].(charptr), args[4].(float64))
	case 8399929451977893162:
		var fnv func(voidptr, int64, charptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(charptr), args[3].(int), args[4].(float64))
	case 3701286485082905967:
		var fnv func(voidptr, int64, charptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int64), args[2].(charptr), args[3].(float64), args[4].(int))
	case 11977083467836069639:
		var fnv func(voidptr, float64, int64, int, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int64), args[3].(int), args[4].(charptr))
	case 3063648799861952235:
		var fnv func(voidptr, float64, int64, charptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int64), args[3].(charptr), args[4].(int))
	case 17813996985112926727:
		var fnv func(voidptr, float64, int, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int), args[3].(int64), args[4].(charptr))
	case 13688450509705449614:
		var fnv func(voidptr, float64, int, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(int), args[3].(charptr), args[4].(int64))
	case 16618717699668433981:
		var fnv func(voidptr, float64, charptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(charptr), args[3].(int), args[4].(int64))
	case 10005620058939583549:
		var fnv func(voidptr, float64, charptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(float64), args[2].(charptr), args[3].(int64), args[4].(int))
	case 2782229490749957034:
		var fnv func(voidptr, int, float64, int64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(float64), args[3].(int64), args[4].(charptr))
	case 7805675134715390243:
		var fnv func(voidptr, int, float64, charptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(float64), args[3].(charptr), args[4].(int64))
	case 9429988573750152002:
		var fnv func(voidptr, int, int64, float64, charptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(int64), args[3].(float64), args[4].(charptr))
	case 13662977334629309539:
		var fnv func(voidptr, int, int64, charptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(int64), args[3].(charptr), args[4].(float64))
	case 17492263205283840599:
		var fnv func(voidptr, int, charptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(charptr), args[3].(int64), args[4].(float64))
	case 10287019833546040929:
		var fnv func(voidptr, int, charptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(int), args[2].(charptr), args[3].(float64), args[4].(int64))
	case 5099000898586828313:
		var fnv func(voidptr, charptr, float64, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(float64), args[3].(int), args[4].(int64))
	case 3078733240262749721:
		var fnv func(voidptr, charptr, float64, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(float64), args[3].(int64), args[4].(int))
	case 6476335657127078358:
		var fnv func(voidptr, charptr, int, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int), args[3].(float64), args[4].(int64))
	case 2730103089238232544:
		var fnv func(voidptr, charptr, int, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int), args[3].(int64), args[4].(float64))
	case 8425219364170907872:
		var fnv func(voidptr, charptr, int64, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int64), args[3].(int), args[4].(float64))
	case 3684915884252444837:
		var fnv func(voidptr, charptr, int64, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(voidptr), args[1].(charptr), args[2].(int64), args[3].(float64), args[4].(int))
	case 4333666961986149722:
		var fnv func(charptr, int64, float64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(float64), args[3].(voidptr), args[4].(int))
	case 6105171995698366680:
		var fnv func(charptr, int64, float64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(float64), args[3].(int), args[4].(voidptr))
	case 635232204771128461:
		var fnv func(charptr, int64, voidptr, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(voidptr), args[3].(float64), args[4].(int))
	case 5701516949963181256:
		var fnv func(charptr, int64, voidptr, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(voidptr), args[3].(int), args[4].(float64))
	case 4780231674552532067:
		var fnv func(charptr, int64, int, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(int), args[3].(voidptr), args[4].(float64))
	case 9012937291235666995:
		var fnv func(charptr, int64, int, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int64), args[2].(int), args[3].(float64), args[4].(voidptr))
	case 5177244198117943044:
		var fnv func(charptr, float64, int64, voidptr, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int64), args[3].(voidptr), args[4].(int))
	case 3406109564532290182:
		var fnv func(charptr, float64, int64, int, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int64), args[3].(int), args[4].(voidptr))
	case 13328791355713470946:
		var fnv func(charptr, float64, voidptr, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(voidptr), args[3].(int64), args[4].(int))
	case 15313347877078528482:
		var fnv func(charptr, float64, voidptr, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(voidptr), args[3].(int), args[4].(int64))
	case 3789544237747698047:
		var fnv func(charptr, float64, int, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int), args[3].(voidptr), args[4].(int64))
	case 9100592385411940230:
		var fnv func(charptr, float64, int, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(float64), args[2].(int), args[3].(int64), args[4].(voidptr))
	case 8064915461730624790:
		var fnv func(charptr, voidptr, float64, int64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(float64), args[3].(int64), args[4].(int))
	case 258609795792393494:
		var fnv func(charptr, voidptr, float64, int, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(float64), args[3].(int), args[4].(int64))
	case 8536288359427978154:
		var fnv func(charptr, voidptr, int64, float64, int) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int64), args[3].(float64), args[4].(int))
	case 3583152468217318383:
		var fnv func(charptr, voidptr, int64, int, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int64), args[3].(int), args[4].(float64))
	case 6968977631315478255:
		var fnv func(charptr, voidptr, int, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int), args[3].(int64), args[4].(float64))
	case 2068593171453928153:
		var fnv func(charptr, voidptr, int, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(voidptr), args[2].(int), args[3].(float64), args[4].(int64))
	case 16515644808242987218:
		var fnv func(charptr, int, float64, voidptr, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(float64), args[3].(voidptr), args[4].(int64))
	case 12677559926711060011:
		var fnv func(charptr, int, float64, int64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(float64), args[3].(int64), args[4].(voidptr))
	case 551831174376210320:
		var fnv func(charptr, int, voidptr, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(voidptr), args[3].(float64), args[4].(int64))
	case 8911402361097984934:
		var fnv func(charptr, int, voidptr, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(voidptr), args[3].(int64), args[4].(float64))
	case 3814388479315590803:
		var fnv func(charptr, int, int64, voidptr, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(int64), args[3].(voidptr), args[4].(float64))
	case 841548384036564675:
		var fnv func(charptr, int, int64, float64, voidptr) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(charptr), args[1].(int), args[2].(int64), args[3].(float64), args[4].(voidptr))
	default:
		log.Println("nocare", tycrc, len(args), voidptr(fnptrx))
	} // end switch tycrc
	return rv
}
