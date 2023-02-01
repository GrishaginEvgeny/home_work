package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
}

func NewList() List {
	return new(list)
}

func (list *list) Len() int {
	var length int
	actualItem := list.first
	if actualItem == nil {
		length = 0
	} else {
		length = 1
		for actualItem.Next != nil {
			length++
			actualItem = actualItem.Next
		}
	}
	return length
}

func (list *list) Front() *ListItem {
	return list.first
}

func (list *list) Back() *ListItem {
	return list.last
}

func (list *list) PushFront(item interface{}) *ListItem {
	newItem := &ListItem{
		Value: item,
		Prev:  nil,
		Next:  list.first,
	}
	if list.first == nil {
		list.first = newItem
		list.last = newItem
		list.first.Prev = nil
		list.first.Next = nil
	} else {
		list.first.Prev = newItem
		list.first = newItem
	}
	return newItem
}

func (list *list) PushBack(item interface{}) *ListItem {
	newItem := &ListItem{
		Value: item,
		Prev:  list.last,
		Next:  nil,
	}
	if list.last == nil {
		list.first = newItem
		list.last = newItem
		list.first.Prev = nil
		list.first.Next = nil
	} else {
		list.last.Next = newItem
		list.last = newItem
		if list.Len() == 1 {
			list.first.Next = list.last
		}
	}
	return newItem
}

func (list *list) Remove(item *ListItem) {
	switch {
	case item.Prev == nil:
		{
			list.first = list.first.Next
			list.first.Prev = nil
			break
		}
	case item.Next == nil:
		{
			list.last = list.last.Prev
			list.last.Next = nil
			break
		}
	default:
		{
			item.Prev.Next = item.Next
			item.Next.Prev = item.Prev
			break
		}
	}
}

func (list *list) MoveToFront(item *ListItem) {
	list.Remove(item)
	list.PushFront(item.Value)
}
