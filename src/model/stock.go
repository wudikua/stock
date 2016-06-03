package model

import (
	"fmt"
	"strings"
)

type StockType int

const (
	CHUANGYE StockType = iota
	HU_A
	HU_B
	SHEN_A
	SHEN_B
	ZHONG_XIAO
	UNKNOWN
)

type Stock struct {
	CnName string
	PyName string
	Code   string
	Type   StockType
}

func (this *Stock) String() string {
	return fmt.Sprintf("[%s][%s]%s(%s)", this.Type, this.Code, this.CnName, this.PyName)
}

func Code2Type(code string) StockType {
	if strings.HasPrefix(code, "300") {
		return CHUANGYE
	} else if strings.HasPrefix(code, "600") ||
		strings.HasPrefix(code, "601") ||
		strings.HasPrefix(code, "603") {
		return HU_A
	} else if strings.HasPrefix(code, "900") {
		return HU_B
	} else if strings.HasPrefix(code, "000") {
		return SHEN_A
	} else if strings.HasPrefix(code, "002") {
		return ZHONG_XIAO
	} else if strings.HasPrefix(code, "200") {
		return SHEN_B
	}
	return UNKNOWN
}

func (this StockType) String() string {
	switch this {
	case CHUANGYE:
		return "创"
	case HU_A:
		return "沪A"
	case HU_B:
		return "沪B"
	case SHEN_A:
		return "深A"
	case SHEN_B:
		return "深B"
	case ZHONG_XIAO:
		return "中小"
	}
	return "??"
}
