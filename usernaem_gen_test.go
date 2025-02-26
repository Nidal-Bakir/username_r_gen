package usernaemgen

import (
	"strconv"
	"strings"
	"testing"
)

func TestCalcTotalPossibleValues(t *testing.T) {
	type testCase struct {
		Gen      *usernameGen
		List     []DictType
		Expected int
	}

	testCaseArr := []testCase{
		testCase{
			Gen:      NewUsernameGenWithOptions("-", UseProvidedNumberAfterOverflow, Adjectives, Colors, Animals),
			List:     []DictType{Adjectives, Colors, Animals},
			Expected: DictsLen[Adjectives] * DictsLen[Colors] * DictsLen[Animals],
		},
		testCase{
			Gen:      NewUsernameGenWithOptions("-", UseProvidedNumberAfterOverflow, Adjectives, Animals),
			List:     []DictType{Adjectives, Animals},
			Expected: DictsLen[Adjectives] * DictsLen[Animals],
		},
		testCase{
			Gen:      NewUsernameGenWithOptions("-", UseProvidedNumberAfterOverflow, Adjectives),
			List:     []DictType{Adjectives},
			Expected: DictsLen[Adjectives],
		},
	}

	for _, tc := range testCaseArr {
		acual := tc.Gen.CalcTotalPossibleValues()
		if tc.Expected != acual {
			t.Fatalf("The values must be equal, Expected: %d AND Acual: %d", tc.Expected, acual)
		}
	}
}

func TestAllPossibleValuesWithNoConflicts(t *testing.T) {
	nameGen := NewUsernameGenWithOptions("-", NoPostfix, Adjectives)
	testCount := nameGen.CalcTotalPossibleValues()
	set := make(map[string]bool, testCount)

	var conflictCount int = 0

	for i := range testCount {
		actual := nameGen.Generate(int64(i))
		if set[actual] {
			conflictCount++
		} else {
			set[actual] = true
		}
	}

	if conflictCount != 0 {
		t.Fatalf("There should not be any conflicats, but got %d conflicts", conflictCount)
	}
}

func TestAllPossibleValuesWithConflicts(t *testing.T) {
	nameGen := NewUsernameGenWithOptions("-", NoPostfix, Adjectives)
	testCount := nameGen.CalcTotalPossibleValues()
	set := make(map[string]bool, testCount)

	// added new 5000 test case to overflow to the beginning
	overloadedValue := 5000
	testCount += overloadedValue

	var conflictCount int = 0

	for i := range testCount {
		actual := nameGen.Generate(int64(i))
		if set[actual] {
			conflictCount++
		} else {
			set[actual] = true
		}
	}

	if conflictCount != overloadedValue {
		t.Fatalf("There should be exactly %d conflicats, but got %d", overloadedValue, conflictCount)
	}
}

func TestGenerateFn(t *testing.T) {
	nameGen := NewUsernameGenWithOptions("-", NoPostfix, Adjectives, Colors, Animals)
	acual := nameGen.Generate(0)

	sb := strings.Builder{}
	sb.WriteString(dicts[Adjectives][0])
	sb.WriteString("-")
	sb.WriteString(dicts[Colors][0])
	sb.WriteString("-")
	sb.WriteString(dicts[Animals][0])
	expected := sb.String()

	if expected != acual {
		t.Fatalf("The values must be equal, Expected: %s AND Acual: %s", expected, acual)
	}

	// -------------------------------------------------------------------------

	nameGen.SetDelimiter("_")
	acual = nameGen.Generate(35790)

	carry := 35790
	col1 := carry % DictsLen[Adjectives]
	carry = carry / DictsLen[Adjectives]

	col2 := carry % DictsLen[Colors]
	carry = carry / DictsLen[Colors]

	col3 := carry % DictsLen[Animals]

	sb = strings.Builder{}
	sb.WriteString(dicts[Adjectives][col1])
	sb.WriteString("_")
	sb.WriteString(dicts[Colors][col2])
	sb.WriteString("_")
	sb.WriteString(dicts[Animals][col3])
	expected = sb.String()

	if expected != acual {
		t.Fatalf("The values must be equal, Expected: %s AND Acual: %s", expected, acual)
	}
}

func TestGenerateWithUseProvidedNumberAfterOverflow(t *testing.T) {
	nameGen := NewUsernameGenWithOptions("-", UseProvidedNumberAfterOverflow, Adjectives, Colors, Animals)

	randNumber := int64(nameGen.CalcTotalPossibleValues()) + 500
	expected := strconv.FormatInt(randNumber, 10)

	ResultStr := nameGen.Generate(randNumber)
	arr := strings.Split(ResultStr, "-")
	acual := arr[len(arr)-1]

	if expected != acual {
		t.Fatalf("The values must be equal, Expected: %s AND Acual: %s", expected, acual)
	}
}
