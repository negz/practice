package url

// A Shortener gets and creates short URLs!
type Shortener interface {
	Get(path string) (string, error)
	Create(url string) (string, error)
}
