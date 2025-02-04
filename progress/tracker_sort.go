package progress

import "sort"

// SortBy helps sort a list of Trackers by various means.
type SortBy int

const (
	// SortByNone doesn't do any sorting == sort by insertion order.
	SortByNone SortBy = iota

	// SortByMessage sorts by the Message alphabetically in ascending order.
	SortByMessage

	// SortByMessageDsc sorts by the Message alphabetically in descending order.
	SortByMessageDsc

	// SortByPercent sorts by the Percentage complete in ascending order.
	SortByPercent

	// SortByPercentDsc sorts by the Percentage complete in descending order.
	SortByPercentDsc

	// SortByValue sorts by the Value in ascending order.
	SortByValue

	// SortByValueDsc sorts by the Value in descending order.
	SortByValueDsc
)

// Sort applies the sorting method defined by SortBy.
func (sb SortBy) Sort(trackers []*Tracker) {
	switch sb {
	case SortByMessage:
		sort.Sort(sortByMessage(trackers))
	case SortByMessageDsc:
		sort.Sort(sortDsc{sortByMessage(trackers)})
	case SortByPercent:
		sort.Sort(sortByPercent(trackers))
	case SortByPercentDsc:
		sort.Sort(sortDsc{sortByPercent(trackers)})
	case SortByValue:
		sort.Sort(sortByValue(trackers))
	case SortByValueDsc:
		sort.Sort(sortDsc{sortByValue(trackers)})
	default:
		// no sort
	}
}

type sortByMessage []*Tracker

func (sb sortByMessage) Len() int           { return len(sb) }
func (sb sortByMessage) Swap(i, j int)      { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByMessage) Less(i, j int) bool { return sb[i].message() < sb[j].message() }

type sortByPercent []*Tracker

func (sb sortByPercent) Len() int           { return len(sb) }
func (sb sortByPercent) Swap(i, j int)      { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByPercent) Less(i, j int) bool {
	if sb[i].PercentDone() == sb[j].PercentDone() {
		return sb[i].timeStart.Before(sb[j].timeStart)
	}
	return sb[i].PercentDone() < sb[j].PercentDone()
}

type sortByValue []*Tracker

func (sb sortByValue) Len() int           { return len(sb) }
func (sb sortByValue) Swap(i, j int)      { sb[i], sb[j] = sb[j], sb[i] }
func (sb sortByValue) Less(i, j int) bool {
	if sb[i].value == sb[j].value {
		return sb[i].timeStart.Before(sb[j].timeStart)
	}
	return sb[i].value < sb[j].value
}

type sortDsc struct{ sort.Interface }

func (sd sortDsc) Less(i, j int) bool { return !sd.Interface.Less(i, j) && sd.Interface.Less(j, i) }
