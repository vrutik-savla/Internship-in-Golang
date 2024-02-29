package dateLayout

func DateLayout() []string {
	return []string{
		"02 Jan 2006",
		"Jan 02, 2006",
		"2006-01-02",
		"2006-01-02T15:04:05Z",
		"Monday, 02 January 2006",
		"Mon Jan 02 15:04:05 2006",
		"Mon Jan 02 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"02 Jan 06 15:04 MST",
		"02 Jan 06 15:04 -0700",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"3:04PM",
	}
}