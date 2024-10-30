package orzbase

type Props struct {
	Service string `json:"service,omitempty" yaml:"service,omitempty" toml:"service,omitempty"`
}

var PropsObj = Props{
	Service: "unnamed",
}
