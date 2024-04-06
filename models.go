package simpleBubble

type fancyItem struct {
	title, desc string
}

func (i fancyItem) Title() string       { return i.title }
func (i fancyItem) Description() string { return i.desc }
func (i fancyItem) FilterValue() string { return i.title }

func NewFancyItem(title, desc string) fancyItem {
	return fancyItem{
		title: title,
		desc:  desc,
	}
}

func NewFancyItems(items ...fancyItem) *[]fancyItem {
	return &items
}

type checkListItem struct {
	Text    string
	Checked bool
}

func NewCheckListItem(text string, checked bool) checkListItem {
	return checkListItem{
		Text:    text,
		Checked: checked,
	}
}

func NewCheckListItems(items ...checkListItem) *[]checkListItem {
	return &items
}
