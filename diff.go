package moment

import (
	"fmt"
	"math"
	"time"
)

// @todo In months/years requires the old and new to calculate correctly, right?
// @todo decide how to handle rounding (i.e. always floor?)
type Diff struct {
	duration time.Duration
}

func (d *Diff) InSeconds() int {
	return int(d.duration.Seconds())
}

func (d *Diff) InMinutes() int {
	return int(d.duration.Minutes())
}

func (d *Diff) InHours() int {
	return int(d.duration.Hours())
}

func (d *Diff) InDays() int {
	days := float64(d.InSeconds()) / 86400
	return round(days) //include today also in the diff
}

// This depends on where the weeks fall?
func (d *Diff) InWeeks() int {
	return int(math.Floor(float64(d.InDays() / 7)))
}

func (d *Diff) InMonths() int {
	// 400 years have 146097 days (taking into account leap year rules)
	// 400 years have 12 months === 4800
	days := d.InDays()
	months := float64(days) * 0.0328
	return round(months)
}

func (d *Diff) InYears() int {
	return 0
}

func (d *Diff) CustomHumanize(ignoreAgo bool) string {
	humanizedStr := d.Humanize()
	if !ignoreAgo {
		humanizedStr = humanizedStr + " " + "ago"
	}
	return humanizedStr
}

// http://momentjs.com/docs/#/durations/humanize/
func (d *Diff) Humanize() string {
	diffInSeconds := d.InSeconds()

	if diffInSeconds <= 45 {
		return fmt.Sprintf("%d seconds", diffInSeconds)
	} else if diffInSeconds <= 90 {
		return "a minute"
	}

	diffInMinutes := d.InMinutes()

	if diffInMinutes <= 45 {
		return fmt.Sprintf("%d minutes", diffInMinutes)
	} else if diffInMinutes <= 90 {
		return "an hour"
	}

	diffInHours := d.InHours()

	if diffInHours <= 22 {
		if diffInHours == 1 {
			return fmt.Sprintf("%d hour", diffInHours)
		}
		return fmt.Sprintf("%d hours", diffInHours)
	} else if diffInHours <= 36 {
		return "a day"
	}

	diffInDays := d.InDays()

	if diffInDays <= 30 {
		if diffInDays == 1 {
			return fmt.Sprintf("%d day", diffInDays)
		}
		return fmt.Sprintf("%d days", diffInDays)
	} else if diffInDays <= 31 {
		return "about 1 month"
	}

	diffInMonths := d.InMonths()

	if diffInMonths <= 12 {
		if diffInMonths == 1 {
			return fmt.Sprintf("%d month", diffInMonths)
		}
		return fmt.Sprintf("%d months", diffInMonths)
	} else if diffInMonths <= 13 {
		return "about 1 year"
	}
	return "many years"
}

// round converts 2.7 months to 3 months and 2.2 days to 2 days
func round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
