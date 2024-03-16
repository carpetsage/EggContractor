package pb

import (
	"time"

	"github.com/carpetsage/EggContractor/util"
)

func (c *SoloContract) GetDurationUntilProductionDeadline() time.Duration {
	return util.DoubleToDuration(c.SecondsUntilProductionDeadline)
}

func (c *SoloContract) GetDurationUntilCollectionDeadline() time.Duration {
	return util.DoubleToDuration(c.SecondsUntilCollectionDeadline)
}

func (c *SoloContract) GetServerRefreshTime() time.Time {
	return util.DoubleToTime(c.ServerRefreshTimestamp)
}
