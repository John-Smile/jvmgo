package heap
type Object struct {
	class            *Class
	fields           Slots
}

func (self *Object) IsInstanceOf(class *Class) bool  {
	return class.IsAssignableFrom(self.class)
}
