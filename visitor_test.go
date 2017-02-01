package proto

type collector struct {
	proto *Proto
}

func collect(p *Proto) collector {
	return collector{p}
}

func (c collector) Comments() (list []*Comment) {
	for _, each := range c.proto.Elements {
		if c, ok := each.(*Comment); ok {
			list = append(list, c)
		}
	}
	return
}

func (c collector) Enums() (list []*Enum) {
	for _, each := range c.proto.Elements {
		if c, ok := each.(*Enum); ok {
			list = append(list, c)
		}
	}
	return
}

func (c collector) Messages() (list []*Message) {
	for _, each := range c.proto.Elements {
		if c, ok := each.(*Message); ok {
			list = append(list, c)
		}
	}
	return
}

func (c collector) Services() (list []*Service) {
	for _, each := range c.proto.Elements {
		if c, ok := each.(*Service); ok {
			list = append(list, c)
		}
	}
	return
}
