package concurrency

type WebsiteCheck func(string) bool
type result struct {
	index    string
	validity bool
}

func CheckWebsites(wc WebsiteCheck, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultchannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultchannel <- result{u, wc(u)}
		}(url)

	}
	for i := 0; i < len(urls); i++ {
		r := <-resultchannel
		results[r.index] = r.validity
	}
	return results
}
