package orzbase

type Props struct {
	Service string `json:"service" yaml:"service" toml:"service"`
}

var PropsObj = Props{
	Service: "unnamed",
}
