package webpbn

// Options puzzle parser options.
type Options struct {
	Validator Validator
}

// Option puzzle parser option.
type Option func(*Options)

// WithValidator option defines a puzzle validator.
func WithValidator(v Validator) Option {
	return func(o *Options) {
		o.Validator = v
	}
}
