// service/client.go

package service

type Client struct {
	Foo Foo
	Bar Bar
}

func NewClient(foo Foo, bar Bar) *Client {
	return &Client{
		Foo: foo,
		Bar: bar,
	}
}