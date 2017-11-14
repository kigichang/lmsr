package lmsr

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

const (
	w = "#"
)

func id() string {
	return uuid.New().String()
}

func prefix(p string) string {
	return p + id()
}

func combine(with string, a, b interface{}) string {
	return fmt.Sprintf("%v"+with+"%v", a, b)
}

// EventID ...
func EventID() string {
	return prefix("evt-")
}

// MarketID ...
func MarketID() string {
	return prefix("mkt-")
}

// OutcomeID ...
func OutcomeID(evt string, index int) string {
	return combine(w, evt, index)
}

// ShareID ...
func ShareID(mkt, outcome string) string {
	return combine(w, mkt, outcome)
}

func sep(str, sep string) (string, string, bool) {
	tmp := strings.Split(str, sep)
	if len(tmp) > 1 {
		return tmp[0], strings.Join(tmp[1:], sep), true
	}
	return "", "", false
}

// SepShareID ...
func SepShareID(id string) (string, string, bool) {
	return sep(id, w)
}

// SepOutcomeID ...
func SepOutcomeID(id string) (string, int, bool) {
	s1, s2, ok := sep(id, w)
	if !ok {
		return "", 0, false
	}

	idx, err := strconv.ParseInt(s2, 10, 32)
	if err != nil {
		return "", 0, false
	}

	return s1, int(idx), true
}

// Fund ...
func Fund(liquidity float64, n int) float64 {
	return liquidity * math.Log(float64(n))
}

// Liquidity ...
func Liquidity(fund float64, n int) float64 {
	return fund / math.Log(float64(n))
}

// Exp ...
func Exp(liquidity, share float64) float64 {
	return math.Exp(share / liquidity)
}

// ExpSum ...
func ExpSum(liquidity float64, shares []float64) float64 {
	sum := 0.0

	for _, x := range shares {
		sum += Exp(liquidity, x)
	}

	return sum
}

// SharePrice ...
func SharePrice(liquidity, share, sum float64) float64 {
	return Exp(liquidity, share) / sum
}

// Cost ...
func Cost(liquidity float64, shares []float64) float64 {
	return liquidity * math.Log(ExpSum(liquidity, shares))
}

// Return ...
func Return(liquidity, answer float64, shares []float64) float64 {
	return Cost(liquidity, shares) - answer
}

// Revenue ...
func Revenue(liquidity, fund, answer float64, shares []float64) float64 {
	return Return(liquidity, answer, shares) - fund
}
