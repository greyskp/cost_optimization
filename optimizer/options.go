package optimization

type options struct {
	observer Observer
}

type Option func(*options)

func WithObserver(o Observer) Option {
	return func(opt *options) {
		if o != nil {
			opt.observer = o
		}
	}
}

func applyOptions(opts []Option) options {
	cfg := options{observer: NoOpObserver{}}
	for _, o := range opts {
		if o != nil {
			o(&cfg)
		}
	}
	return cfg
}