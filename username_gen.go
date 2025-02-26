package usernaemgen

import (
	"math/rand/v2"
	"strconv"
	"strings"
)

type PostfixType byte

const (
	// Do not use any postfix
	NoPostfix PostfixType = iota

	// Use the provided number as postfix for the result in
	// case the number overflows the total possible values
	UseProvidedNumberAfterOverflow PostfixType = iota

	// Use the provided number as postfix for the result
	UseProvidedNumber PostfixType = iota
)

type usernameGen struct {
	dicts               []DictType
	delimiter           string
	postfixType         PostfixType
	totalPossibleValues int
}

func NewUsernameGen() *usernameGen {
	return NewUsernameGenWithOptions("-", UseProvidedNumberAfterOverflow, Adjectives, Colors, Animals)
}

func NewUsernameGenWithOptions(delimiter string, postfixType PostfixType, dicts ...DictType) *usernameGen {
	if len(dicts) == 0 {
		panic("the dicts can not be empty")
	}

	return &usernameGen{
		dicts:       dicts,
		delimiter:   delimiter,
		postfixType: postfixType,
	}
}

func (u *usernameGen) SetDelimiter(delimiter string) {
	u.delimiter = delimiter
}

func (u *usernameGen) SetPostfixType(postfixType PostfixType) {
	u.postfixType = postfixType
}

func (u *usernameGen) SetDicts(dicts ...DictType) {
	if len(dicts) == 0 {
		panic("the dicts can not be empty")
	}

	u.totalPossibleValues = 0
	u.dicts = dicts
}

func (u usernameGen) CalcTotalPossibleValues() int {
	if u.totalPossibleValues != 0 {
		return u.totalPossibleValues
	}

	var t int = 1
	for _, dictType := range u.dicts {
		t *= DictsLen[dictType]
	}

	u.totalPossibleValues = t
	return t
}

func (u usernameGen) GenerateRand() string {
	return u.Generate(rand.Int64())
}

func (u usernameGen) Generate(randNum int64) string {
	dictsCount := len(u.dicts)
	indices := u.generateIndices(randNum)

	if len(indices) != dictsCount {
		panic("the lenguth of te array of indices should equal the dictsCount")
	}

	sb := strings.Builder{}

	for i, dictType := range u.dicts {
		word := dicts[dictType][indices[i]]
		sb.WriteString(word)

		if i != dictsCount-1 {
			sb.WriteString(u.delimiter)
		}
	}

	u.addPostfix(&sb, randNum)

	return sb.String()
}

func (u usernameGen) generateIndices(randNum int64) []int64 {
	dictsCount := len(u.dicts)
	indexArr := make([]int64, 0, dictsCount)

	carry := randNum
	for _, dictType := range u.dicts {
		dictLen := int64(DictsLen[dictType])
		index := carry % dictLen
		indexArr = append(indexArr, index)
		carry = carry / dictLen
	}

	return indexArr
}

func (u usernameGen) addPostfix(sb *strings.Builder, randNum int64) {
	switch u.postfixType {
	case NoPostfix:
		return

	case UseProvidedNumber:
		sb.WriteString(u.delimiter)
		sb.WriteString(strconv.FormatInt(randNum, 10))

	case UseProvidedNumberAfterOverflow:
		if randNum >= int64(u.CalcTotalPossibleValues()) {
			sb.WriteString(u.delimiter)
			sb.WriteString(strconv.FormatInt(randNum, 10))
		}
	}
}
