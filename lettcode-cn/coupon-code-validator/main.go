package main

import "sort"

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	// 按照 businessLine 进行排序，顺序为："electronics"、"grocery"、"pharmacy"、"restaurant"。
	businessLineOrder := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	type Coupon struct {
		Code              string
		BusinessLineOrder int
	}
	var validCoupons []Coupon
	for i := 0; i < len(code); i++ {
		// businessLine[i] 必须是以下四个类别之一："electronics"、"grocery"、"pharmacy"、"restaurant"。
		// code[i] 不能为空，并且仅由字母数字字符（a-z、A-Z、0-9）和下划线（_）组成。
		if isActive[i] && (businessLine[i] == "electronics" || businessLine[i] == "grocery" || businessLine[i] == "pharmacy" || businessLine[i] == "restaurant") && isValidCode(code[i]) {
			validCoupons = append(validCoupons, Coupon{Code: code[i], BusinessLineOrder: businessLineOrder[businessLine[i]]})
		}
	}
	// 先按照其 businessLine 的顺序排序："electronics"、"grocery"、"pharmacy"、"restaurant"。
	// 在每个类别内，再按照 标识符的字典序（升序）排序。
	sort.SliceStable(validCoupons, func(i, j int) bool {
		if validCoupons[i].BusinessLineOrder == validCoupons[j].BusinessLineOrder {
			return validCoupons[i].Code < validCoupons[j].Code
		}
		return validCoupons[i].BusinessLineOrder < validCoupons[j].BusinessLineOrder
	})

	// 提取排序后的优惠券代码
	var result []string
	for _, coupon := range validCoupons {
		result = append(result, coupon.Code)
	}
	return result
}

func isValidCode(code string) bool {
	if len(code) == 0 {
		return false
	}
	for _, ch := range code {
		switch {
		case ch >= 'a' && ch <= 'z':
		case ch >= 'A' && ch <= 'Z':
		case ch >= '0' && ch <= '9':
		case ch == '_':
		default:
			return false
		}
	}
	return true
}
