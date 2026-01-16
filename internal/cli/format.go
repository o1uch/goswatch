package cli

import (
	"fmt"
	"time"
)

func Seconds(d time.Duration) string {
	seconds := int(d.Seconds())

	if seconds%100 >= 11 && seconds%100 <= 14 {
		return fmt.Sprintf("%v секунд", seconds)
	}

	if seconds%10 == 1 {
		return fmt.Sprintf("%v секунда", seconds)
	}

	if seconds%10 > 1 && seconds%10 <= 4 {
		return fmt.Sprintf("%v секунды", seconds)
	}

	return fmt.Sprintf("%v секунд", seconds)
}

func Minutes(d time.Duration) string {
	minutes := int(d.Minutes())

	if minutes%100 >= 11 && minutes%100 <= 14 {
		return fmt.Sprintf("%v минут", minutes)
	}

	if minutes%10 == 1 {
		return fmt.Sprintf("%v минута", minutes)
	}

	if minutes%10 > 1 && minutes%10 <= 4 {
		return fmt.Sprintf("%v минуты", minutes)
	}

	return fmt.Sprintf("%v минут", minutes)

}

func Hours(d time.Duration) string {
	return fmt.Sprintf("%.1f часа", d.Hours())
}

func DefaultFormat(d time.Duration) string {
	d = d.Round(time.Second)
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02dh %02dm %02ds", hours, minutes, seconds)
}
