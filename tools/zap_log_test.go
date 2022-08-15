package tools

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestZapLog(t *testing.T){
	zap.L().Info("[STATISTIC]", zap.Uint64("time",uint64(time.Now().UnixNano() / 1e6)), zap.String("xx","xx"))
}
