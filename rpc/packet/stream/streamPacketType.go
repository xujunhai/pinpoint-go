package stream

const (
	ApplicationStreamCreate        int16 = 10
	ApplicationStreamCreateSuccess int16 = 12
	ApplicationStreamCreateFail    int16 = 14

	ApplicationStreamClose int16 = 15

	ApplicationStreamPing int16 = 17
	ApplicationStreamPong int16 = 18

	ApplicationStreamResponse int16 = 20
)
