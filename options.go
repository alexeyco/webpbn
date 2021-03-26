package webpbn

type Options struct {
	Validator Validator
}

type Option func(*Options)

func WithValidator(v Validator) Option {
	return func(o *Options) {
		o.Validator = v
	}
}
