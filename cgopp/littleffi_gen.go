// oohmyffi
package cgopp

import "log"
import "github.com/ebitengine/purego"

func litfficallgenimpl[RETY any](tycrc uint64, fnptrx uintptr, args ...any) RETY {
	var rv RETY
	switch tycrc {
	case 5316925208990529324:
		var fnv func(float64, float64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 14599226747074119783:
		var fnv func(float64, float64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 14933284767871535207:
		var fnv func(float64, float64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 13016573749669964881:
		var fnv func(float64, float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 5838980223726630430:
		var fnv func(float64, float64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 6102176919137138206:
		var fnv func(float64, float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 12915894768449399889:
		var fnv func(float64, float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 1155236611261313977:
		var fnv func(float64, float64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 1562548494747677625:
		var fnv func(float64, float64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 1414575800293393593:
		var fnv func(float64, float64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 14764594021983152160:
		var fnv func(float64, float64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 14501889907781888032:
		var fnv func(float64, float64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 12887217139485101078:
		var fnv func(float64, float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 14871178468217238224:
		var fnv func(float64, float64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 14683557804053541584:
		var fnv func(float64, float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 13061220352138791958:
		var fnv func(float64, float64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 10331550043827779447:
		var fnv func(float64, float64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 9999814191588226935:
		var fnv func(float64, float64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 1692231417206539449:
		var fnv func(float64, float64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 13427024932654116794:
		var fnv func(float64, float64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 13830853563303686074:
		var fnv func(float64, float64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 14152728634776210316:
		var fnv func(float64, float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 14871178468159972304:
		var fnv func(float64, float64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 14683557803996275664:
		var fnv func(float64, float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 14119566264570858380:
		var fnv func(float64, float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 10331550043770513015:
		var fnv func(float64, float64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 9999814191530960503:
		var fnv func(float64, float64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int64))
	case 2429157946889275318:
		var fnv func(float64, int32, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(float64))
	case 3718913313307250469:
		var fnv func(float64, int32, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int32))
	case 3907800614866145061:
		var fnv func(float64, int32, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int64))
	case 5594258298860399379:
		var fnv func(float64, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 13086344553680801598:
		var fnv func(float64, int32, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int32))
	case 12698876656052539198:
		var fnv func(float64, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 5348197492168780563:
		var fnv func(float64, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 17625972978070260377:
		var fnv func(float64, int32, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int32))
	case 17382620268517853849:
		var fnv func(float64, int32, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int64))
	case 16908608284380906491:
		var fnv func(float64, int32, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(float64))
	case 2914480501864255744:
		var fnv func(float64, int32, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int32))
	case 3302581718190117120:
		var fnv func(float64, int32, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int64))
	case 6218529006754952502:
		var fnv func(float64, int32, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(float64))
	case 16067521050508140446:
		var fnv func(float64, int32, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int32))
	case 15806576154911317918:
		var fnv func(float64, int32, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int64))
	case 5897171444769977654:
		var fnv func(float64, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 11527892626118679097:
		var fnv func(float64, int32, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int32))
	case 11122832542446000697:
		var fnv func(float64, int32, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int64))
	case 17184301410477425659:
		var fnv func(float64, int32, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(float64))
	case 6828108578049218202:
		var fnv func(float64, int32, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int32))
	case 6585705846543210138:
		var fnv func(float64, int32, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int64))
	case 2376958524607918764:
		var fnv func(float64, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 16067521050484428446:
		var fnv func(float64, int32, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int32))
	case 15806576154887605918:
		var fnv func(float64, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 2839227297785410220:
		var fnv func(float64, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 11527892626094967609:
		var fnv func(float64, int32, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int32))
	case 11122832542422289209:
		var fnv func(float64, int32, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int64))
	case 2869685891967158198:
		var fnv func(float64, int64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 17710752930617248760:
		var fnv func(float64, int64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 17324868329732983800:
		var fnv func(float64, int64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 10085754502232338382:
		var fnv func(float64, int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 13086344553745521038:
		var fnv func(float64, int64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 12698876656117258638:
		var fnv func(float64, int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 10118916872437690318:
		var fnv func(float64, int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 17625972978134981673:
		var fnv func(float64, int64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 17382620268582575145:
		var fnv func(float64, int64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 3238776040427898662:
		var fnv func(float64, int64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 2914480501936244656:
		var fnv func(float64, int64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 3302581718262106032:
		var fnv func(float64, int64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 6218529006833820550:
		var fnv func(float64, int64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 16067517019369906078:
		var fnv func(float64, int64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 15806572123773083550:
		var fnv func(float64, int64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 5897171444848845702:
		var fnv func(float64, int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 11527888594980444729:
		var fnv func(float64, int64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 11122828511307766329:
		var fnv func(float64, int64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 2895787167714395942:
		var fnv func(float64, int64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 6828108578121203754:
		var fnv func(float64, int64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 6585705846615195690:
		var fnv func(float64, int64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 2376958524686787612:
		var fnv func(float64, int64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 16067517019427983006:
		var fnv func(float64, int64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 15806572123831160478:
		var fnv func(float64, int64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 2839227297864279068:
		var fnv func(float64, int64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 11527888595038522169:
		var fnv func(float64, int64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 11122828511365843769:
		var fnv func(float64, int64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int64))
	case 14756567099256403311:
		var fnv func(int32, float64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 16304054535263874354:
		var fnv func(int32, float64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 16709712753262062898:
		var fnv func(int32, float64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 11276262000390523140:
		var fnv func(int32, float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 14798823771881111620:
		var fnv func(int32, float64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 14481478327825891396:
		var fnv func(int32, float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 11233953892465330436:
		var fnv func(int32, float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 10115080159415797219:
		var fnv func(int32, float64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 9941849903436432867:
		var fnv func(int32, float64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 4285431669036016108:
		var fnv func(int32, float64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 5814607034913152634:
		var fnv func(int32, float64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 6131319160270773882:
		var fnv func(int32, float64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 3354994765065157196:
		var fnv func(int32, float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 15862503979794318435:
		var fnv func(int32, float64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 15674109259444667491:
		var fnv func(int32, float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 2960277787274967628:
		var fnv func(int32, float64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 11322875555404857796:
		var fnv func(int32, float64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 10990365646979350980:
		var fnv func(int32, float64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 4586212351943500268:
		var fnv func(int32, float64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 2738648489419102688:
		var fnv func(int32, float64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 2568971855020710368:
		var fnv func(int32, float64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 6358895716521278934:
		var fnv func(int32, float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 15862503979801253219:
		var fnv func(int32, float64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 15674109259451602275:
		var fnv func(int32, float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 6891920261969411542:
		var fnv func(int32, float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 11322875555411792068:
		var fnv func(int32, float64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 10990365646986285252:
		var fnv func(int32, float64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int64))
	case 679874354466841315:
		var fnv func(int32, int32, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(float64))
	case 12667492943623584127:
		var fnv func(int32, int32, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int32))
	case 12280341705344121215:
		var fnv func(int32, int32, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int64))
	case 15129013966179654985:
		var fnv func(int32, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12096647007055772045:
		var fnv func(int32, int32, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int32))
	case 11707279153334712717:
		var fnv func(int32, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 15451673349931916617:
		var fnv func(int32, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 16636275431445232682:
		var fnv func(int32, int32, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int32))
	case 16391022765800029226:
		var fnv func(int32, int32, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int64))
	case 8570547890262734241:
		var fnv func(int32, int32, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(float64))
	case 4193332590539777971:
		var fnv func(int32, int32, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int32))
	case 4582067125563238323:
		var fnv func(int32, int32, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int64))
	case 4904211611239940997:
		var fnv func(int32, int32, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(float64))
	case 8912196085319270418:
		var fnv func(int32, int32, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int32))
	case 9081802350973485074:
		var fnv func(int32, int32, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int64))
	case 4869819987034735493:
		var fnv func(int32, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 4228452472853956021:
		var fnv func(int32, int32, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int32))
	case 4542173926584026549:
		var fnv func(int32, int32, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int64))
	case 8227273144526009761:
		var fnv func(int32, int32, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(float64))
	case 5512887281359331369:
		var fnv func(int32, int32, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int32))
	case 5271117868550922281:
		var fnv func(int32, int32, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int64))
	case 3656714514458313759:
		var fnv func(int32, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 8912196085377618194:
		var fnv func(int32, int32, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int32))
	case 9081802351031832850:
		var fnv func(int32, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 3829488473112151071:
		var fnv func(int32, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 4228452472912303285:
		var fnv func(int32, int32, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int32))
	case 4542173926642373813:
		var fnv func(int32, int32, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int64))
	case 1083766941474415331:
		var fnv func(int32, int64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 7610794987016650146:
		var fnv func(int32, int64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 7798415651180346786:
		var fnv func(int32, int64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 1702376102104652180:
		var fnv func(int32, int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12096647006974668605:
		var fnv func(int32, int64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 11707279153253609277:
		var fnv func(int32, int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 1169351556656519572:
		var fnv func(int32, int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 16636275431364127386:
		var fnv func(int32, int64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 16391022765718923930:
		var fnv func(int32, int64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 12728777721209254268:
		var fnv func(int32, int64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 4193332590484959491:
		var fnv func(int32, int64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 4582067125508419843:
		var fnv func(int32, int64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 4904211611177194805:
		var fnv func(int32, int64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 8912201190199328786:
		var fnv func(int32, int64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 9081807455853543442:
		var fnv func(int32, int64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 4869819986971989301:
		var fnv func(int32, int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 4228457577734014389:
		var fnv func(int32, int64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 4542179031464084917:
		var fnv func(int32, int64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 13004184974282551676:
		var fnv func(int32, int64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 5512887281304516249:
		var fnv func(int32, int64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 5271117868496107161:
		var fnv func(int32, int64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 3656714514395566767:
		var fnv func(int32, int64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 8912201190175887634:
		var fnv func(int32, int64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 9081807455830102290:
		var fnv func(int32, int64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 3829488473049404079:
		var fnv func(int32, int64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 4228457577710572725:
		var fnv func(int32, int64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 4542179031440643253:
		var fnv func(int32, int64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int64))
	case 14498357444186797423:
		var fnv func(int64, float64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 14597243053048589143:
		var fnv func(int64, float64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 14930738123892583255:
		var fnv func(int64, float64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 13019102279624773473:
		var fnv func(int64, float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 14798823771987883537:
		var fnv func(int64, float64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 14481478327932663313:
		var fnv func(int64, float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 12913919698776837985:
		var fnv func(int64, float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 10115080159522567094:
		var fnv func(int64, float64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 9941849903543202742:
		var fnv func(int64, float64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 1416550745470789513:
		var fnv func(int64, float64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 5814607035007330351:
		var fnv func(int64, float64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 6131319160364951599:
		var fnv func(int64, float64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 3354994765105318937:
		var fnv func(int64, float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 15862497258467050595:
		var fnv func(int64, float64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 15674102538117399651:
		var fnv func(int64, float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 2960277787315129369:
		var fnv func(int64, float64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 11322868834077589956:
		var fnv func(int64, float64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 10990358925652083140:
		var fnv func(int64, float64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 1690274646241961865:
		var fnv func(int64, float64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 2738648489513283509:
		var fnv func(int64, float64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 2568971855114891189:
		var fnv func(int64, float64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 6358895716561439619:
		var fnv func(int64, float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 15862497258408973667:
		var fnv func(int64, float64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 15674102538059322723:
		var fnv func(int64, float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 6891920262009572227:
		var fnv func(int64, float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 11322868834019512516:
		var fnv func(int64, float64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 10990358925594005700:
		var fnv func(int64, float64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(int64))
	case 2431686648583742598:
		var fnv func(int64, int32, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(float64))
	case 12667492943651032874:
		var fnv func(int64, int32, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int32))
	case 12280341705371569962:
		var fnv func(int64, int32, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(float64), args[4].(int64))
	case 15129013966286285596:
		var fnv func(int64, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12096644656864754061:
		var fnv func(int64, int32, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int32))
	case 11707276803143694733:
		var fnv func(int64, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 15451673350038547228:
		var fnv func(int64, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 16636273081254214698:
		var fnv func(int64, int32, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int32))
	case 16391020415609011242:
		var fnv func(int64, int32, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(int64))
	case 8570547890369348596:
		var fnv func(int64, int32, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(float64))
	case 4193330241430890419:
		var fnv func(int64, int32, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int32))
	case 4582064776454350771:
		var fnv func(int64, int32, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(float64), args[4].(int64))
	case 4904213420122342277:
		var fnv func(int64, int32, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(float64))
	case 9064423663458385938:
		var fnv func(int64, int32, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int32))
	case 8677377978295189522:
		var fnv func(int64, int32, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int32), args[4].(int64))
	case 4869821795917136773:
		var fnv func(int64, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 4380680050993071541:
		var fnv func(int64, int32, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int32))
	case 4137749553905730997:
		var fnv func(int64, int32, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(int64))
	case 8227273144632624116:
		var fnv func(int64, int32, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(float64))
	case 5512884932250443817:
		var fnv func(int64, int32, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int32))
	case 5271115519442034729:
		var fnv func(int64, int32, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(float64), args[4].(int64))
	case 3656716323340715039:
		var fnv func(int64, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 9064423663516733714:
		var fnv func(int64, int32, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int32))
	case 8677377978353537298:
		var fnv func(int64, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 3829490281994552351:
		var fnv func(int64, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 4380680051051418805:
		var fnv func(int64, int32, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int32))
	case 4137749553964078261:
		var fnv func(int64, int32, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(int64))
	case 2871643216982369414:
		var fnv func(int64, int64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(float64))
	case 7610794987044097015:
		var fnv func(int64, int64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int32))
	case 7798415651207793655:
		var fnv func(int64, int64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(float64), args[4].(int64))
	case 1702376102211284929:
		var fnv func(int64, int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12096644656927960893:
		var fnv func(int64, int64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int32))
	case 11707276803206901565:
		var fnv func(int64, int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 1169351556763152321:
		var fnv func(int64, int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 16636273081317419674:
		var fnv func(int64, int64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int32))
	case 16391020415672216218:
		var fnv func(int64, int64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(int64))
	case 12728777721315870505:
		var fnv func(int64, int64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(float64))
	case 4193330241503604995:
		var fnv func(int64, int64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int32))
	case 4582064776527065347:
		var fnv func(int64, int64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(float64), args[4].(int64))
	case 4904213420200760629:
		var fnv func(int64, int64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 9064419585698365458:
		var fnv func(int64, int64, int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int32))
	case 8677373900535169042:
		var fnv func(int64, int64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 4869821795995555125:
		var fnv func(int64, int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 4380675973233051061:
		var fnv func(int64, int64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int32))
	case 4137745476145710517:
		var fnv func(int64, int64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(int64))
	case 13004184974389167913:
		var fnv func(int64, int64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(float64))
	case 5512884932323161753:
		var fnv func(int64, int64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int32))
	case 5271115519514752665:
		var fnv func(int64, int64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(float64), args[4].(int64))
	case 3656716323419132591:
		var fnv func(int64, int64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 9064419585674924306:
		var fnv func(int64, int64, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int32))
	case 8677373900511727890:
		var fnv func(int64, int64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 3829490282072969903:
		var fnv func(int64, int64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 4380675973209609397:
		var fnv func(int64, int64, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int32))
	case 4137745476122268853:
		var fnv func(int64, int64, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(int64))
	default:
		log.Println("nocare", tycrc, len(args), voidptr(fnptrx))
	} // end switch tycrc
	return rv
}
