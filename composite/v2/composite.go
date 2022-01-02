package composite

type IOrganization interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e *Employee) Count() int {
	return 1
}

type Department struct {
	Name             string
	SubOrganizations []IOrganization
}

func (d *Department) Count() int {
	c := 0
	for _, o := range d.SubOrganizations {
		c += o.Count()
	}

	return c
}

func (d *Department) AddSub(o IOrganization) {
	d.SubOrganizations = append(d.SubOrganizations, o)
}

func NewOrganization() IOrganization {
	root := &Department{ Name: "root"}
	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Department{ Name: "sub", SubOrganizations: []IOrganization{&Employee{}}})
	}

	return  root
}