package tools

import "testing"

func TestTimeSpan(t *testing.T) {
	//span := GetNowTimeSpan(40 * 60)
	//t.Log(FormatUnixToString(span, BEIBING_TZ, SECONDS_FORMAT))
}

func TestLocation(t *testing.T) {
	beijingTime := GetLocationByHour(8)
	t.Log(FormatUnixToString(GetNowTimeSeconds(), beijingTime, SECONDS_FORMAT))
}
