package heap
type Object struct {
	class            *Class
	data             interface{}
}

func (self *Object) IsInstanceOf(class *Class) bool  {
	return class.isAssignableFrom(self.class)
}
func (self *Object) Fields() Slots  {
	return self.data.(Slots)
}
func (self *Object) Class() *Class  {
	return self.class
}
