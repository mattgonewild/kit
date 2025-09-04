package help

import (
	"math"
	"math/bits"
)

const (
	ZeroUint        uint = 0
	OneUint         uint = 1
	TwoUint         uint = 2
	ThreeUint       uint = 3
	FourUint        uint = 4
	FiveUint        uint = 5
	SixUint         uint = 6
	SevenUint       uint = 7
	EightUint       uint = 8
	NineUint        uint = 9
	TenUint         uint = 10
	ElevenUint      uint = 11
	TwelveUint      uint = 12
	ThirteenUint    uint = 13
	FourteenUint    uint = 14
	FifteenUint     uint = 15
	SixteenUint     uint = 16
	SeventeenUint   uint = 17
	EighteenUint    uint = 18
	NineteenUint    uint = 19
	TwentyUint      uint = 20
	TwentyOneUint   uint = 21
	TwentyTwoUint   uint = 22
	TwentyThreeUint uint = 23
	TwentyFourUint  uint = 24
	TwentyFiveUint  uint = 25
	TwentySixUint   uint = 26
	TwentySevenUint uint = 27
	TwentyEightUint uint = 28
	TwentyNineUint  uint = 29
	ThirtyUint      uint = 30
	ThirtyOneUint   uint = 31
	ThirtyTwoUint   uint = 32
	ThirtyThreeUint uint = 33
	ThirtyFourUint  uint = 34
	ThirtyFiveUint  uint = 35
	ThirtySixUint   uint = 36
	ThirtySevenUint uint = 37
	ThirtyEightUint uint = 38
	ThirtyNineUint  uint = 39
	FortyUint       uint = 40
	FortyOneUint    uint = 41
	FortyTwoUint    uint = 42
	FortyThreeUint  uint = 43
	FortyFourUint   uint = 44
	FortyFiveUint   uint = 45
	FortySixUint    uint = 46
	FortySevenUint  uint = 47
	FortyEightUint  uint = 48
	FortyNineUint   uint = 49
	FiftyUint       uint = 50
	FiftyOneUint    uint = 51
	FiftyTwoUint    uint = 52
	FiftyThreeUint  uint = 53
	FiftyFourUint   uint = 54
	FiftyFiveUint   uint = 55
	FiftySixUint    uint = 56
	FiftySevenUint  uint = 57
	FiftyEightUint  uint = 58
	FiftyNineUint   uint = 59
	SixtyUint       uint = 60
	SixtyOneUint    uint = 61
	SixtyTwoUint    uint = 62
	SixtyThreeUint  uint = 63
	SixtyFourUint   uint = 64
	MaxUint         uint = math.MaxUint
)

const (
	NegativeOneInt int = -1
	ZeroInt        int = 0
	OneInt         int = 1
	TwoInt         int = 2
	ThreeInt       int = 3
	FourInt        int = 4
	FiveInt        int = 5
	SixInt         int = 6
	SevenInt       int = 7
	EightInt       int = 8
	NineInt        int = 9
	TenInt         int = 10
	ElevenInt      int = 11
	TwelveInt      int = 12
	ThirteenInt    int = 13
	FourteenInt    int = 14
	FifteenInt     int = 15
	SixteenInt     int = 16
	SeventeenInt   int = 17
	EighteenInt    int = 18
	NineteenInt    int = 19
	TwentyInt      int = 20
	TwentyOneInt   int = 21
	TwentyTwoInt   int = 22
	TwentyThreeInt int = 23
	TwentyFourInt  int = 24
	TwentyFiveInt  int = 25
	TwentySixInt   int = 26
	TwentySevenInt int = 27
	TwentyEightInt int = 28
	TwentyNineInt  int = 29
	ThirtyInt      int = 30
	ThirtyOneInt   int = 31
	ThirtyTwoInt   int = 32
	ThirtyThreeInt int = 33
	ThirtyFourInt  int = 34
	ThirtyFiveInt  int = 35
	ThirtySixInt   int = 36
	ThirtySevenInt int = 37
	ThirtyEightInt int = 38
	ThirtyNineInt  int = 39
	FortyInt       int = 40
	FortyOneInt    int = 41
	FortyTwoInt    int = 42
	FortyThreeInt  int = 43
	FortyFourInt   int = 44
	FortyFiveInt   int = 45
	FortySixInt    int = 46
	FortySevenInt  int = 47
	FortyEightInt  int = 48
	FortyNineInt   int = 49
	FiftyInt       int = 50
	FiftyOneInt    int = 51
	FiftyTwoInt    int = 52
	FiftyThreeInt  int = 53
	FiftyFourInt   int = 54
	FiftyFiveInt   int = 55
	FiftySixInt    int = 56
	FiftySevenInt  int = 57
	FiftyEightInt  int = 58
	FiftyNineInt   int = 59
	SixtyInt       int = 60
	SixtyOneInt    int = 61
	SixtyTwoInt    int = 62
	SixtyThreeInt  int = 63
	SixtyFourInt   int = 64
	MaxInt         int = math.MaxInt
)

const (
	NegativeOneInt64 int64 = -1
	ZeroInt64        int64 = 0
	OneInt64         int64 = 1
	TwoInt64         int64 = 2
	ThreeInt64       int64 = 3
	FourInt64        int64 = 4
	FiveInt64        int64 = 5
	SixInt64         int64 = 6
	SevenInt64       int64 = 7
	EightInt64       int64 = 8
	NineInt64        int64 = 9
	TenInt64         int64 = 10
	ElevenInt64      int64 = 11
	TwelveInt64      int64 = 12
	ThirteenInt64    int64 = 13
	FourteenInt64    int64 = 14
	FifteenInt64     int64 = 15
	SixteenInt64     int64 = 16
	SeventeenInt64   int64 = 17
	EighteenInt64    int64 = 18
	NineteenInt64    int64 = 19
	TwentyInt64      int64 = 20
	TwentyOneInt64   int64 = 21
	TwentyTwoInt64   int64 = 22
	TwentyThreeInt64 int64 = 23
	TwentyFourInt64  int64 = 24
	TwentyFiveInt64  int64 = 25
	TwentySixInt64   int64 = 26
	TwentySevenInt64 int64 = 27
	TwentyEightInt64 int64 = 28
	TwentyNineInt64  int64 = 29
	ThirtyInt64      int64 = 30
	ThirtyOneInt64   int64 = 31
	ThirtyTwoInt64   int64 = 32
	ThirtyThreeInt64 int64 = 33
	ThirtyFourInt64  int64 = 34
	ThirtyFiveInt64  int64 = 35
	ThirtySixInt64   int64 = 36
	ThirtySevenInt64 int64 = 37
	ThirtyEightInt64 int64 = 38
	ThirtyNineInt64  int64 = 39
	FortyInt64       int64 = 40
	FortyOneInt64    int64 = 41
	FortyTwoInt64    int64 = 42
	FortyThreeInt64  int64 = 43
	FortyFourInt64   int64 = 44
	FortyFiveInt64   int64 = 45
	FortySixInt64    int64 = 46
	FortySevenInt64  int64 = 47
	FortyEightInt64  int64 = 48
	FortyNineInt64   int64 = 49
	FiftyInt64       int64 = 50
	FiftyOneInt64    int64 = 51
	FiftyTwoInt64    int64 = 52
	FiftyThreeInt64  int64 = 53
	FiftyFourInt64   int64 = 54
	FiftyFiveInt64   int64 = 55
	FiftySixInt64    int64 = 56
	FiftySevenInt64  int64 = 57
	FiftyEightInt64  int64 = 58
	FiftyNineInt64   int64 = 59
	SixtyInt64       int64 = 60
	SixtyOneInt64    int64 = 61
	SixtyTwoInt64    int64 = 62
	SixtyThreeInt64  int64 = 63
	SixtyFourInt64   int64 = 64
	MaxInt64         int64 = math.MaxInt64
)

const (
	WordSize  uint = bits.UintSize
	WordRMask uint = WordSize - OneUint
	WordShift uint = ((WordSize >> TwoUint) / SixUint) + FourUint
)

const (
	E0 uint = iota
	E1
	E2
	E3
	E4
	E5
	E6
	E7
	E8
	E9
	E10
	E11
	E12
	E13
	E14
	E15
	E16
	E17
	E18
	E19
	E20
	E21
	E22
	E23
	E24
	E25
	E26
	E27
	E28
	E29
	E30
	E31
	E32
	E33
	E34
)

const (
	L0 uint = OneUint << iota
	L1
	L2
	L3
	L4
	L5
	L6
	L7
	L8
	L9
	L10
	L11
	L12
	L13
	L14
	L15
	L16
	L17
	L18
	L19
	L20
	L21
	L22
	L23
	L24
	L25
	L26
	L27
	L28
	L29
	L30
	L31
	L32
	L33
	L34

	WL6  uint = L6 >> WordShift
	WL7  uint = L7 >> WordShift
	WL8  uint = L8 >> WordShift
	WL9  uint = L9 >> WordShift
	WL10 uint = L10 >> WordShift
	WL11 uint = L11 >> WordShift
	WL12 uint = L12 >> WordShift
	WL13 uint = L13 >> WordShift
	WL14 uint = L14 >> WordShift
	WL15 uint = L15 >> WordShift
	WL16 uint = L16 >> WordShift
	WL17 uint = L17 >> WordShift
	WL18 uint = L18 >> WordShift
	WL19 uint = L19 >> WordShift
	WL20 uint = L20 >> WordShift
	WL21 uint = L21 >> WordShift
	WL22 uint = L22 >> WordShift
	WL23 uint = L23 >> WordShift
	WL24 uint = L24 >> WordShift
	WL25 uint = L25 >> WordShift
	WL26 uint = L26 >> WordShift
	WL27 uint = L27 >> WordShift
	WL28 uint = L28 >> WordShift
	WL29 uint = L29 >> WordShift
	WL30 uint = L30 >> WordShift
	WL31 uint = L31 >> WordShift
	WL32 uint = L32 >> WordShift
	WL33 uint = L33 >> WordShift
	WL34 uint = L34 >> WordShift

	SL6  uint = (WL6 + WordRMask) >> WordShift
	SL7  uint = (WL7 + WordRMask) >> WordShift
	SL8  uint = (WL8 + WordRMask) >> WordShift
	SL9  uint = (WL9 + WordRMask) >> WordShift
	SL10 uint = (WL10 + WordRMask) >> WordShift
	SL11 uint = (WL11 + WordRMask) >> WordShift
	SL12 uint = (WL12 + WordRMask) >> WordShift
	SL13 uint = (WL13 + WordRMask) >> WordShift
	SL14 uint = (WL14 + WordRMask) >> WordShift
	SL15 uint = (WL15 + WordRMask) >> WordShift
	SL16 uint = (WL16 + WordRMask) >> WordShift
	SL17 uint = (WL17 + WordRMask) >> WordShift
	SL18 uint = (WL18 + WordRMask) >> WordShift
	SL19 uint = (WL19 + WordRMask) >> WordShift
	SL20 uint = (WL20 + WordRMask) >> WordShift
	SL21 uint = (WL21 + WordRMask) >> WordShift
	SL22 uint = (WL22 + WordRMask) >> WordShift
	SL23 uint = (WL23 + WordRMask) >> WordShift
	SL24 uint = (WL24 + WordRMask) >> WordShift
	SL25 uint = (WL25 + WordRMask) >> WordShift
	SL26 uint = (WL26 + WordRMask) >> WordShift
	SL27 uint = (WL27 + WordRMask) >> WordShift
	SL28 uint = (WL28 + WordRMask) >> WordShift
	SL29 uint = (WL29 + WordRMask) >> WordShift
	SL30 uint = (WL30 + WordRMask) >> WordShift
	SL31 uint = (WL31 + WordRMask) >> WordShift
	SL32 uint = (WL32 + WordRMask) >> WordShift
	SL33 uint = (WL33 + WordRMask) >> WordShift
	SL34 uint = (WL34 + WordRMask) >> WordShift
)

const (
	BM0  uint = MaxUint >> ((WordSize - (L0 & WordRMask)) & WordRMask)
	BM1  uint = MaxUint >> ((WordSize - (L1 & WordRMask)) & WordRMask)
	BM2  uint = MaxUint >> ((WordSize - (L2 & WordRMask)) & WordRMask)
	BM3  uint = MaxUint >> ((WordSize - (L3 & WordRMask)) & WordRMask)
	BM4  uint = MaxUint >> ((WordSize - (L4 & WordRMask)) & WordRMask)
	BM5  uint = MaxUint >> ((WordSize - (L5 & WordRMask)) & WordRMask)
	BM6  uint = MaxUint >> ((WordSize - (L6 & WordRMask)) & WordRMask)
	BM7  uint = MaxUint >> ((WordSize - (L7 & WordRMask)) & WordRMask)
	BM8  uint = MaxUint >> ((WordSize - (L8 & WordRMask)) & WordRMask)
	BM9  uint = MaxUint >> ((WordSize - (L9 & WordRMask)) & WordRMask)
	BM10 uint = MaxUint >> ((WordSize - (L10 & WordRMask)) & WordRMask)
	BM11 uint = MaxUint >> ((WordSize - (L11 & WordRMask)) & WordRMask)
	BM12 uint = MaxUint >> ((WordSize - (L12 & WordRMask)) & WordRMask)
	BM13 uint = MaxUint >> ((WordSize - (L13 & WordRMask)) & WordRMask)
	BM14 uint = MaxUint >> ((WordSize - (L14 & WordRMask)) & WordRMask)
	BM15 uint = MaxUint >> ((WordSize - (L15 & WordRMask)) & WordRMask)
	BM16 uint = MaxUint >> ((WordSize - (L16 & WordRMask)) & WordRMask)
	BM17 uint = MaxUint >> ((WordSize - (L17 & WordRMask)) & WordRMask)
	BM18 uint = MaxUint >> ((WordSize - (L18 & WordRMask)) & WordRMask)
	BM19 uint = MaxUint >> ((WordSize - (L19 & WordRMask)) & WordRMask)
	BM20 uint = MaxUint >> ((WordSize - (L20 & WordRMask)) & WordRMask)
	BM21 uint = MaxUint >> ((WordSize - (L21 & WordRMask)) & WordRMask)
	BM22 uint = MaxUint >> ((WordSize - (L22 & WordRMask)) & WordRMask)
	BM23 uint = MaxUint >> ((WordSize - (L23 & WordRMask)) & WordRMask)
	BM24 uint = MaxUint >> ((WordSize - (L24 & WordRMask)) & WordRMask)
	BM25 uint = MaxUint >> ((WordSize - (L25 & WordRMask)) & WordRMask)
	BM26 uint = MaxUint >> ((WordSize - (L26 & WordRMask)) & WordRMask)
	BM27 uint = MaxUint >> ((WordSize - (L27 & WordRMask)) & WordRMask)
	BM28 uint = MaxUint >> ((WordSize - (L28 & WordRMask)) & WordRMask)
	BM29 uint = MaxUint >> ((WordSize - (L29 & WordRMask)) & WordRMask)
	BM30 uint = MaxUint >> ((WordSize - (L30 & WordRMask)) & WordRMask)
	BM31 uint = MaxUint >> ((WordSize - (L31 & WordRMask)) & WordRMask)
	BM32 uint = MaxUint >> ((WordSize - (L32 & WordRMask)) & WordRMask)
	BM33 uint = MaxUint >> ((WordSize - (L33 & WordRMask)) & WordRMask)
	BM34 uint = MaxUint >> ((WordSize - (L34 & WordRMask)) & WordRMask)

	RM0  uint = L0 - OneUint
	RM1  uint = L1 - OneUint
	RM2  uint = L2 - OneUint
	RM3  uint = L3 - OneUint
	RM4  uint = L4 - OneUint
	RM5  uint = L5 - OneUint
	RM6  uint = L6 - OneUint
	RM7  uint = L7 - OneUint
	RM8  uint = L8 - OneUint
	RM9  uint = L9 - OneUint
	RM10 uint = L10 - OneUint
	RM11 uint = L11 - OneUint
	RM12 uint = L12 - OneUint
	RM13 uint = L13 - OneUint
	RM14 uint = L14 - OneUint
	RM15 uint = L15 - OneUint
	RM16 uint = L16 - OneUint
	RM17 uint = L17 - OneUint
	RM18 uint = L18 - OneUint
	RM19 uint = L19 - OneUint
	RM20 uint = L20 - OneUint
	RM21 uint = L21 - OneUint
	RM22 uint = L22 - OneUint
	RM23 uint = L23 - OneUint
	RM24 uint = L24 - OneUint
	RM25 uint = L25 - OneUint
	RM26 uint = L26 - OneUint
	RM27 uint = L27 - OneUint
	RM28 uint = L28 - OneUint
	RM29 uint = L29 - OneUint
	RM30 uint = L30 - OneUint
	RM31 uint = L31 - OneUint
	RM32 uint = L32 - OneUint
	RM33 uint = L33 - OneUint
	RM34 uint = L34 - OneUint

	WRM6  uint = WL6 - OneUint
	WRM7  uint = WL7 - OneUint
	WRM8  uint = WL8 - OneUint
	WRM9  uint = WL9 - OneUint
	WRM10 uint = WL10 - OneUint
	WRM11 uint = WL11 - OneUint
	WRM12 uint = WL12 - OneUint
	WRM13 uint = WL13 - OneUint
	WRM14 uint = WL14 - OneUint
	WRM15 uint = WL15 - OneUint
	WRM16 uint = WL16 - OneUint
	WRM17 uint = WL17 - OneUint
	WRM18 uint = WL18 - OneUint
	WRM19 uint = WL19 - OneUint
	WRM20 uint = WL20 - OneUint
	WRM21 uint = WL21 - OneUint
	WRM22 uint = WL22 - OneUint
	WRM23 uint = WL23 - OneUint
	WRM24 uint = WL24 - OneUint
	WRM25 uint = WL25 - OneUint
	WRM26 uint = WL26 - OneUint
	WRM27 uint = WL27 - OneUint
	WRM28 uint = WL28 - OneUint
	WRM29 uint = WL29 - OneUint
	WRM30 uint = WL30 - OneUint
	WRM31 uint = WL31 - OneUint
	WRM32 uint = WL32 - OneUint
	WRM33 uint = WL33 - OneUint
	WRM34 uint = WL34 - OneUint

	SBM6  uint = MaxUint >> ((WordSize - (SL6 & WordRMask)) & WordRMask)
	SBM7  uint = MaxUint >> ((WordSize - (SL7 & WordRMask)) & WordRMask)
	SBM8  uint = MaxUint >> ((WordSize - (SL8 & WordRMask)) & WordRMask)
	SBM9  uint = MaxUint >> ((WordSize - (SL9 & WordRMask)) & WordRMask)
	SBM10 uint = MaxUint >> ((WordSize - (SL10 & WordRMask)) & WordRMask)
	SBM11 uint = MaxUint >> ((WordSize - (SL11 & WordRMask)) & WordRMask)
	SBM12 uint = MaxUint >> ((WordSize - (SL12 & WordRMask)) & WordRMask)
	SBM13 uint = MaxUint >> ((WordSize - (SL13 & WordRMask)) & WordRMask)
	SBM14 uint = MaxUint >> ((WordSize - (SL14 & WordRMask)) & WordRMask)
	SBM15 uint = MaxUint >> ((WordSize - (SL15 & WordRMask)) & WordRMask)
	SBM16 uint = MaxUint >> ((WordSize - (SL16 & WordRMask)) & WordRMask)
	SBM17 uint = MaxUint >> ((WordSize - (SL17 & WordRMask)) & WordRMask)
	SBM18 uint = MaxUint >> ((WordSize - (SL18 & WordRMask)) & WordRMask)
	SBM19 uint = MaxUint >> ((WordSize - (SL19 & WordRMask)) & WordRMask)
	SBM20 uint = MaxUint >> ((WordSize - (SL20 & WordRMask)) & WordRMask)
	SBM21 uint = MaxUint >> ((WordSize - (SL21 & WordRMask)) & WordRMask)
	SBM22 uint = MaxUint >> ((WordSize - (SL22 & WordRMask)) & WordRMask)
	SBM23 uint = MaxUint >> ((WordSize - (SL23 & WordRMask)) & WordRMask)
	SBM24 uint = MaxUint >> ((WordSize - (SL24 & WordRMask)) & WordRMask)
	SBM25 uint = MaxUint >> ((WordSize - (SL25 & WordRMask)) & WordRMask)
	SBM26 uint = MaxUint >> ((WordSize - (SL26 & WordRMask)) & WordRMask)
	SBM27 uint = MaxUint >> ((WordSize - (SL27 & WordRMask)) & WordRMask)
	SBM28 uint = MaxUint >> ((WordSize - (SL28 & WordRMask)) & WordRMask)
	SBM29 uint = MaxUint >> ((WordSize - (SL29 & WordRMask)) & WordRMask)
	SBM30 uint = MaxUint >> ((WordSize - (SL30 & WordRMask)) & WordRMask)
	SBM31 uint = MaxUint >> ((WordSize - (SL31 & WordRMask)) & WordRMask)
	SBM32 uint = MaxUint >> ((WordSize - (SL32 & WordRMask)) & WordRMask)
	SBM33 uint = MaxUint >> ((WordSize - (SL33 & WordRMask)) & WordRMask)
	SBM34 uint = MaxUint >> ((WordSize - (SL34 & WordRMask)) & WordRMask)

	SRM6  uint = SL6 - OneUint
	SRM7  uint = SL7 - OneUint
	SRM8  uint = SL8 - OneUint
	SRM9  uint = SL9 - OneUint
	SRM10 uint = SL10 - OneUint
	SRM11 uint = SL11 - OneUint
	SRM12 uint = SL12 - OneUint
	SRM13 uint = SL13 - OneUint
	SRM14 uint = SL14 - OneUint
	SRM15 uint = SL15 - OneUint
	SRM16 uint = SL16 - OneUint
	SRM17 uint = SL17 - OneUint
	SRM18 uint = SL18 - OneUint
	SRM19 uint = SL19 - OneUint
	SRM20 uint = SL20 - OneUint
	SRM21 uint = SL21 - OneUint
	SRM22 uint = SL22 - OneUint
	SRM23 uint = SL23 - OneUint
	SRM24 uint = SL24 - OneUint
	SRM25 uint = SL25 - OneUint
	SRM26 uint = SL26 - OneUint
	SRM27 uint = SL27 - OneUint
	SRM28 uint = SL28 - OneUint
	SRM29 uint = SL29 - OneUint
	SRM30 uint = SL30 - OneUint
	SRM31 uint = SL31 - OneUint
	SRM32 uint = SL32 - OneUint
	SRM33 uint = SL33 - OneUint
	SRM34 uint = SL34 - OneUint
)
