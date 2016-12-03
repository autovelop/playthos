package ui

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Font struct {
	charmap map[string]mgl32.Vec2
}

func (f *Font) GetVec2(char string) mgl32.Vec2 {
	return f.charmap[char]
}

func (f *Font) NewFont() {
	f.charmap = make(map[string]mgl32.Vec2)
	f.charmap[" "] = mgl32.Vec2{0.0, 0.0}
	f.charmap["spc"] = mgl32.Vec2{0.0, 0.0}
	f.charmap["!"] = mgl32.Vec2{276705.0, 32776.0}
	f.charmap["\""] = mgl32.Vec2{1797408.0, 0.0}
	f.charmap["#"] = mgl32.Vec2{10738.0, 1134484.0}
	f.charmap["dol"] = mgl32.Vec2{538883.0, 19976.0}
	f.charmap["%"] = mgl32.Vec2{1664033.0, 68006.0}
	f.charmap["&"] = mgl32.Vec2{545090.0, 174362.0}
	f.charmap["'"] = mgl32.Vec2{798848.0, 0.0}
	f.charmap["("] = mgl32.Vec2{270466.0, 66568.0}
	f.charmap[")"] = mgl32.Vec2{528449.0, 33296.0}
	f.charmap["*"] = mgl32.Vec2{10471.0, 1688832.0}
	f.charmap["crs"] = mgl32.Vec2{4167.0, 1606144.0}
	f.charmap["per"] = mgl32.Vec2{0.0, 1560.0}
	f.charmap["-"] = mgl32.Vec2{7.0, 1572864.0}
	f.charmap[","] = mgl32.Vec2{0.0, 1544.0}
	f.charmap["lsl"] = mgl32.Vec2{1057.0, 67584.0}
	f.charmap["0"] = mgl32.Vec2{935221.0, 731292.0}
	f.charmap["1"] = mgl32.Vec2{274497.0, 33308.0}
	f.charmap["2"] = mgl32.Vec2{934929.0, 1116222.0}
	f.charmap["3"] = mgl32.Vec2{934931.0, 1058972.0}
	f.charmap["4"] = mgl32.Vec2{137380.0, 1302788.0}
	f.charmap["5"] = mgl32.Vec2{2048263.0, 1058972.0}
	f.charmap["6"] = mgl32.Vec2{401671.0, 1190044.0}
	f.charmap["7"] = mgl32.Vec2{2032673.0, 66576.0}
	f.charmap["8"] = mgl32.Vec2{935187.0, 1190044.0}
	f.charmap["9"] = mgl32.Vec2{935187.0, 1581336.0}
	f.charmap[":"] = mgl32.Vec2{195.0, 1560.0}
	f.charmap[";"] = mgl32.Vec2{195.0, 1544.0}
	f.charmap["<"] = mgl32.Vec2{135300.0, 66052.0}
	f.charmap["="] = mgl32.Vec2{496.0, 3968.0}
	f.charmap[">"] = mgl32.Vec2{528416.0, 541200.0}
	f.charmap["?"] = mgl32.Vec2{934929.0, 1081352.0}
	f.charmap["ats"] = mgl32.Vec2{935285.0, 714780.0}
	f.charmap["A"] = mgl32.Vec2{935188.0, 780450.0}
	f.charmap["B"] = mgl32.Vec2{1983767.0, 1190076.0}
	f.charmap["C"] = mgl32.Vec2{935172.0, 133276.0}
	f.charmap["D"] = mgl32.Vec2{1983764.0, 665788.0}
	f.charmap["E"] = mgl32.Vec2{2048263.0, 1181758.0}
	f.charmap["F"] = mgl32.Vec2{2048263.0, 1181728.0}
	f.charmap["G"] = mgl32.Vec2{935173.0, 1714334.0}
	f.charmap["H"] = mgl32.Vec2{1131799.0, 1714338.0}
	f.charmap["I"] = mgl32.Vec2{921665.0, 33308.0}
	f.charmap["J"] = mgl32.Vec2{66576.0, 665756.0}
	f.charmap["K"] = mgl32.Vec2{1132870.0, 166178.0}
	f.charmap["L"] = mgl32.Vec2{1065220.0, 133182.0}
	f.charmap["M"] = mgl32.Vec2{1142100.0, 665762.0}
	f.charmap["N"] = mgl32.Vec2{1140052.0, 1714338.0}
	f.charmap["O"] = mgl32.Vec2{935188.0, 665756.0}
	f.charmap["P"] = mgl32.Vec2{1983767.0, 1181728.0}
	f.charmap["Q"] = mgl32.Vec2{935188.0, 698650.0}
	f.charmap["R"] = mgl32.Vec2{1983767.0, 1198242.0}
	f.charmap["S"] = mgl32.Vec2{935171.0, 1058972.0}
	f.charmap["T"] = mgl32.Vec2{2035777.0, 33288.0}
	f.charmap["U"] = mgl32.Vec2{1131796.0, 665756.0}
	f.charmap["V"] = mgl32.Vec2{1131796.0, 664840.0}
	f.charmap["W"] = mgl32.Vec2{1131861.0, 699028.0}
	f.charmap["X"] = mgl32.Vec2{1131681.0, 84130.0}
	f.charmap["Y"] = mgl32.Vec2{1131794.0, 1081864.0}
	f.charmap["Z"] = mgl32.Vec2{1968194.0, 133180.0}
	f.charmap["lsb"] = mgl32.Vec2{925826.0, 66588.0}
	f.charmap["rsl"] = mgl32.Vec2{16513.0, 16512.0}
	f.charmap["rsb"] = mgl32.Vec2{919584.0, 1065244.0}
	f.charmap["pow"] = mgl32.Vec2{272656.0, 0.0}
	f.charmap["_"] = mgl32.Vec2{0.0, 62.0}
	f.charmap["a"] = mgl32.Vec2{224.0, 649374.0}
	f.charmap["b"] = mgl32.Vec2{1065444.0, 665788.0}
	f.charmap["c"] = mgl32.Vec2{228.0, 657564.0}
	f.charmap["d"] = mgl32.Vec2{66804.0, 665758.0}
	f.charmap["e"] = mgl32.Vec2{228.0, 772124.0}
	f.charmap["f"] = mgl32.Vec2{401543.0, 1115152.0}
	f.charmap["g"] = mgl32.Vec2{244.0, 665474.0}
	f.charmap["h"] = mgl32.Vec2{1065444.0, 665762.0}
	f.charmap["i"] = mgl32.Vec2{262209.0, 33292.0}
	f.charmap["j"] = mgl32.Vec2{131168.0, 1066252.0}
	f.charmap["k"] = mgl32.Vec2{1065253.0, 199204.0}
	f.charmap["l"] = mgl32.Vec2{266305.0, 33292.0}
	f.charmap["m"] = mgl32.Vec2{421.0, 698530.0}
	f.charmap["n"] = mgl32.Vec2{452.0, 1198372.0}
	f.charmap["o"] = mgl32.Vec2{228.0, 665756.0}
	f.charmap["p"] = mgl32.Vec2{484.0, 667424.0}
	f.charmap["q"] = mgl32.Vec2{244.0, 665474.0}
	f.charmap["r"] = mgl32.Vec2{354.0, 590904.0}
	f.charmap["s"] = mgl32.Vec2{228.0, 114844.0}
	f.charmap["t"] = mgl32.Vec2{8674.0, 66824.0}
	f.charmap["u"] = mgl32.Vec2{292.0, 1198868.0}
	f.charmap["v"] = mgl32.Vec2{276.0, 664840.0}
	f.charmap["w"] = mgl32.Vec2{276.0, 700308.0}
	f.charmap["x"] = mgl32.Vec2{292.0, 1149220.0}
	f.charmap["y"] = mgl32.Vec2{292.0, 1163824.0}
	f.charmap["z"] = mgl32.Vec2{480.0, 1148988.0}
	f.charmap["lpa"] = mgl32.Vec2{401542.0, 66572.0}
	f.charmap["|"] = mgl32.Vec2{266304.0, 33288.0}
	f.charmap["rpa"] = mgl32.Vec2{788512.0, 1589528.0}
	f.charmap["~"] = mgl32.Vec2{675840.0, 0.0}
	f.charmap["lar"] = mgl32.Vec2{8387.0, 1147904.0}
}
