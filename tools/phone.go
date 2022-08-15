package tools

import (
	"github.com/dongri/phonenumber"
	"strings"
)

func ValidatePhone(phone string) bool {
	newPhone := phone

	if strings.Index(phone, "+")==0{
		newPhone = phone[1:]
	}

	if strings.Index(newPhone, "-")>=0{
		newPhone = strings.ReplaceAll(newPhone, "-", "")
	}

	if strings.Index(newPhone, " ")>=0{
		newPhone = strings.ReplaceAll(newPhone, " ", "")
	}
	country := phonenumber.GetISO3166ByNumber(newPhone, true)

	if country.CountryName!="" {

		// 验证手机号的前缀
		isMobileBegin := false
		for _,mobileBegin := range country.MobileBeginWith{
			if strings.Index(newPhone, country.CountryCode+mobileBegin)==0{
				isMobileBegin = true
				break
			}
		}

		if !isMobileBegin{
			return false
		}

		// 验证手机号的 长度
		isMobileLen := false
		countryCodeLength := len(country.CountryCode)
		for _,mobileLength := range country.PhoneNumberLengths{
			if len(newPhone) == mobileLength+countryCodeLength{
				isMobileLen = true
				break
			}
		}

		if !isMobileLen{
			return false
		}

		return true
	}

	return false
}