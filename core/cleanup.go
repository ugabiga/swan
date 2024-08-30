package core

type Cleanup struct {
	cleanups []func() error
}

func NewCleanup() *Cleanup {
	return &Cleanup{}
}

func (c *Cleanup) RegisterCleanup(fn func() error) {
	c.cleanups = append(c.cleanups, fn)
}

func (c *Cleanup) Run() error {
	for _, fn := range c.cleanups {
		if err := fn(); err != nil {
			return err
		}
	}
	return nil
}
