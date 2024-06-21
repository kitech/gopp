// oohmyffi
package cgopp

import "log"
import "github.com/ebitengine/purego"

func litfficallgenimpl[RETY any](tycrc uint64, fnptrx uintptr, args ...any) RETY {
	var rv RETY
	switch tycrc {
	case 4906417137293123632:
		var fnv func(float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64))
	case 3370978100138164748:
		var fnv func(float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64))
	case 1625390155685093194:
		var fnv func(float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(float64))
	case 5938092134339604168:
		var fnv func(float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32))
	case 16726627421087814328:
		var fnv func(float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(float64))
	case 3996830784931025512:
		var fnv func(float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int64))
	case 13061220352138791958:
		var fnv func(float64, float64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 6327319250572308168:
		var fnv func(float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64))
	case 17274323849696989880:
		var fnv func(float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(float64), args[2].(int64), args[3].(float64))
	case 6299612324015391767:
		var fnv func(float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32))
	case 3194917911032941310:
		var fnv func(float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64))
	case 5703060069472926288:
		var fnv func(float64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(float64))
	case 12627726077792357462:
		var fnv func(float64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32))
	case 5594258298860399379:
		var fnv func(float64, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12698876656052539198:
		var fnv func(float64, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 12293069922669431894:
		var fnv func(float64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int64))
	case 5348197492168780563:
		var fnv func(float64, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 13348055723535124332:
		var fnv func(float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32))
	case 15060114439728310368:
		var fnv func(float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(float64))
	case 16877667243833489725:
		var fnv func(float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int64))
	case 5897171444769977654:
		var fnv func(float64, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 13590352901924865900:
		var fnv func(float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64))
	case 15463697879284885600:
		var fnv func(float64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(float64))
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
	case 12582763719507901594:
		var fnv func(float64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32))
	case 2376958524607918764:
		var fnv func(float64, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 15806576154887605918:
		var fnv func(float64, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 12338038819444030618:
		var fnv func(float64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int64))
	case 2839227297785410220:
		var fnv func(float64, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 5966082068799308823:
		var fnv func(float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64))
	case 3156163424688720638:
		var fnv func(float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64))
	case 5474336843164416592:
		var fnv func(float64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(float64))
	case 15676663025522183116:
		var fnv func(float64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32))
	case 10085754502232338382:
		var fnv func(float64, int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 12698876656117258638:
		var fnv func(float64, int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 15864424427174235084:
		var fnv func(float64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int64))
	case 10118916872437690318:
		var fnv func(float64, int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 17887684147924583115:
		var fnv func(float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32))
	case 12083235086036414458:
		var fnv func(float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(float64))
	case 16877667243908573245:
		var fnv func(float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int64))
	case 5897171444848845702:
		var fnv func(float64, int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 18274096514390180555:
		var fnv func(float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64))
	case 11541062603845185530:
		var fnv func(float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(float64), args[1].(int64), args[2].(int64), args[3].(float64))
	case 4733398028036079616:
		var fnv func(int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32))
	case 2977529939112288289:
		var fnv func(int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64))
	case 10039399914472759830:
		var fnv func(int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64))
	case 9021512218731673951:
		var fnv func(int32, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(float64))
	case 5784005156845595475:
		var fnv func(int32, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32))
	case 11276262000390523140:
		var fnv func(int32, float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 14481478327825891396:
		var fnv func(int32, float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 6188959687402007379:
		var fnv func(int32, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int64))
	case 11233953892465330436:
		var fnv func(int32, float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 2653614615967471954:
		var fnv func(int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32))
	case 3349021946929188709:
		var fnv func(int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(float64))
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
	case 10692662485627258067:
		var fnv func(int32, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32))
	case 3354994765065157196:
		var fnv func(int32, float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 15674109259444667491:
		var fnv func(int32, float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 10503669630952096979:
		var fnv func(int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int32), args[3].(int64))
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
	case 2410824856368486738:
		var fnv func(int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64))
	case 3017496101410541413:
		var fnv func(int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(float64))
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
	case 15376406098092573044:
		var fnv func(int32, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32))
	case 6358895716521278934:
		var fnv func(int32, float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 15674109259451602275:
		var fnv func(int32, float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 15043298055341556084:
		var fnv func(int32, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int64))
	case 6891920261969411542:
		var fnv func(int32, float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 1612373000591589808:
		var fnv func(int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32))
	case 6552033577527819620:
		var fnv func(int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64))
	case 9907814013011044237:
		var fnv func(int32, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(float64))
	case 688398158377360109:
		var fnv func(int32, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32))
	case 15129013966179654985:
		var fnv func(int32, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 11707279153334712717:
		var fnv func(int32, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 876757694354922221:
		var fnv func(int32, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int64))
	case 15451673349931916617:
		var fnv func(int32, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 2740278372808791648:
		var fnv func(int32, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32))
	case 8625336419849682651:
		var fnv func(int32, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(float64))
	case 1564254518328894983:
		var fnv func(int32, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int64))
	case 4869819987034735493:
		var fnv func(int32, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 2567083301201516128:
		var fnv func(int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64))
	case 8082000654356266715:
		var fnv func(int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(float64))
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
	case 5841776962842600352:
		var fnv func(int32, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32))
	case 3656714514458313759:
		var fnv func(int32, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 9081802351031832850:
		var fnv func(int32, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 6103882942718354336:
		var fnv func(int32, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int64))
	case 3829488473112151071:
		var fnv func(int32, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 1425174548892959152:
		var fnv func(int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64))
	case 6734903251947896164:
		var fnv func(int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64))
	case 10205780083590375309:
		var fnv func(int32, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(float64))
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
	case 9168676256716004727:
		var fnv func(int32, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int32))
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
	case 8834618235918589303:
		var fnv func(int32, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(float64), args[3].(int64))
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
	case 7424021985274106823:
		var fnv func(int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32))
	case 73000727473109313:
		var fnv func(int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(float64))
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
	case 1158033350268647175:
		var fnv func(int32, int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32))
	case 4904211611177194805:
		var fnv func(int32, int64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 9081807455853543442:
		var fnv func(int32, int64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 1564254518220257031:
		var fnv func(int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int32), args[3].(int64))
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
	case 7106711725590975431:
		var fnv func(int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64))
	case 475420883727497537:
		var fnv func(int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(float64))
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
	case 5841776962733961888:
		var fnv func(int32, int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32))
	case 3656714514395566767:
		var fnv func(int32, int64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 9081807455830102290:
		var fnv func(int32, int64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 6103882942609715872:
		var fnv func(int32, int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int64))
	case 3829488473049404079:
		var fnv func(int32, int64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int32), args[1].(int64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 4906417177782910976:
		var fnv func(int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64))
	case 3370978080484024353:
		var fnv func(int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64))
	case 10344359016443699734:
		var fnv func(int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64))
	case 8887836021683190111:
		var fnv func(int64, float64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(float64))
	case 10813681515688308622:
		var fnv func(int64, float64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32))
	case 13019102279624773473:
		var fnv func(int64, float64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 14481478327932663313:
		var fnv func(int64, float64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 10643864143801560974:
		var fnv func(int64, float64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int64))
	case 12913919698776837985:
		var fnv func(int64, float64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 5936738744320564936:
		var fnv func(int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32))
	case 16730623814730484664:
		var fnv func(int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(float64))
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
	case 10692662485546750563:
		var fnv func(int64, float64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32))
	case 3354994765105318937:
		var fnv func(int64, float64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(float64))
	case 15674102538117399651:
		var fnv func(int64, float64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int32), args[4].(int64))
	case 10503669630871589475:
		var fnv func(int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int32), args[3].(int64))
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
	case 6324452932553449160:
		var fnv func(int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64))
	case 17272796296921713592:
		var fnv func(int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(float64))
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
	case 15376406098012067780:
		var fnv func(int64, float64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32))
	case 6358895716561439619:
		var fnv func(int64, float64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(float64))
	case 15674102538059322723:
		var fnv func(int64, float64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int32), args[4].(int64))
	case 15043298055261050820:
		var fnv func(int64, float64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int64))
	case 6891920262009572227:
		var fnv func(int64, float64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(float64), args[2].(int64), args[3].(int64), args[4].(float64))
	case 6296116613056906263:
		var fnv func(int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32))
	case 3196851855136799486:
		var fnv func(int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64))
	case 5704548185790817104:
		var fnv func(int64, int32, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(float64))
	case 688398158320897117:
		var fnv func(int64, int32, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32))
	case 15129013966286285596:
		var fnv func(int64, int32, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(float64))
	case 11707276803143694733:
		var fnv func(int64, int32, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int32), args[4].(int64))
	case 876757694298459229:
		var fnv func(int64, int32, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int64))
	case 15451673350038547228:
		var fnv func(int64, int32, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(float64), args[3].(int64), args[4].(float64))
	case 2740278372783245152:
		var fnv func(int64, int32, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32))
	case 8625336419788564587:
		var fnv func(int64, int32, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(float64))
	case 1564259729914144263:
		var fnv func(int64, int32, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int64))
	case 4869821795917136773:
		var fnv func(int64, int32, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int32), args[3].(int64), args[4].(float64))
	case 2567083301175969632:
		var fnv func(int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64))
	case 8082000654295148651:
		var fnv func(int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(float64))
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
	case 5841782174427849632:
		var fnv func(int64, int32, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32))
	case 3656716323340715039:
		var fnv func(int64, int32, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(float64))
	case 8677377978353537298:
		var fnv func(int64, int32, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int32), args[4].(int64))
	case 6103888154303603616:
		var fnv func(int64, int32, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int64))
	case 3829490281994552351:
		var fnv func(int64, int32, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int32), args[2].(int64), args[3].(int64), args[4].(float64))
	case 5964802973282419735:
		var fnv func(int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64))
	case 3154541548188351230:
		var fnv func(int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64))
	case 5473904250107536208:
		var fnv func(int64, int64, float64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(float64))
	case 9168676256659540935:
		var fnv func(int64, int64, float64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32))
	case 1702376102211284929:
		var fnv func(int64, int64, float64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(float64))
	case 11707276803206901565:
		var fnv func(int64, int64, float64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int32), args[4].(int64))
	case 8834618235862125511:
		var fnv func(int64, int64, float64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int64))
	case 1169351556763152321:
		var fnv func(int64, int64, float64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(float64), args[3].(int64), args[4].(float64))
	case 7424021985248559815:
		var fnv func(int64, int64, int32) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32))
	case 73000727411994609:
		var fnv func(int64, int64, int32, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(float64))
	case 1564259729987958535:
		var fnv func(int64, int64, int32, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int64))
	case 4869821795995555125:
		var fnv func(int64, int64, int32, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int32), args[3].(int64), args[4].(float64))
	case 7106711725565428423:
		var fnv func(int64, int64, int64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64))
	case 475420883666382833:
		var fnv func(int64, int64, int64, float64) RETY
		purego.RegisterFunc(&fnv, fnptrx)

		rv = fnv(args[0].(int64), args[1].(int64), args[2].(int64), args[3].(float64))
	default:
		log.Println("nocare", tycrc, len(args), voidptr(fnptrx))
	} // end switch tycrc
	return rv
}
