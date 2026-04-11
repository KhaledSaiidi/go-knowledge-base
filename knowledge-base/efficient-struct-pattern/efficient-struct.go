package main

// OptFunc is a function type that knows how to update an Opts value.
type OptFunc func(*Opts)

// Opts stores the configurable values used when creating a Server.
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

// defaultOpts returns the baseline configuration used before any custom options are applied.
func defaultOpts() Opts {
	return Opts{
		maxConn: 100,
		id:      "default",
		tls:     false,
	}
}

// withTLS enables TLS on an existing Opts value.
func withTLS(opts *Opts) {
	opts.tls = true
}

// withMaxConn builds an option function that sets the maximum connection limit.
func withMaxConn(maxConn int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = maxConn
	}
}

func withID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

// Server embeds Opts so the chosen configuration is directly available on the server.
type Server struct {
	Opts
}

// newServer starts from defaults, applies each option function, and returns the configured server.
func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}
