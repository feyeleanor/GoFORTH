package container

import(
	"golightly";
)

const(
	SPARE_CAPACITY = 32;
)

type Container []Value

func (c *Container) New(length, capacity uint) {
	if length + SPARE_CAPACITY > capacity { capacity = length + SPARE_CAPACITY }
	make(Container, length, capacity);
}

func (c *Container) CopyTo(destination *Container) {
	for i, x := range c { destination[i] = x }
}


func (c *Container) CopyFrom(source *Container) {
	for i, x := range source { c[i] = x }
}

func (c *Container) Duplicate() Container {
	new_container := New(c.length, c.length);
	new_container.CopyFrom(c);
	return new_container;
}

func (c *Container) Resize(length, capacity uint) {
	new_container := New(length, capacity);
	new_container.CopyFrom(c);
	c = new_container;
}

func (c *Container) Cap() int {
	return cap(c)
}

func (c *Container) Len() int {
	return len(c)
}

func (c *Container) Last() Value {
	c[c.Len() - 1];
}

func (c *Container) First() Value {
	c[0];
}

func (c *Container) Expand(position, elements uint) Container {
	starting_capacity := c.cap();
	starting_length := c.len();
	desired_length := starting_length + elements;
	switch {
	case desired_length < (starting_capacity - SPARE_CAPACITY):
		c.Resize(desired_length, desired_length + SPARE_CAPACITY);

	case desired_length < starting_capacity:
		c = c[0 : desired_length];

	case desired_length > starting_capacity:
		desired_capacity := starting_capacity + SPARE_CAPACITY;
		if desired_capacity < desired_length { desired_capacity = desired_length + SPARE_CAPACITY }
		c.Resize(desired_length, desired_capacity);

	default:
		return c;
	}

	for j := starting_length - 1; j >= position; j-- {
		c[j + n] = c[j];
		c[j] = nil;
	}
	return c;
}