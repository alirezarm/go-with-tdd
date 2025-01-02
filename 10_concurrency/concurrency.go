package concurrency

type WebsiteChecker func(string) bool

// we don't need to name struct's values (anonymous within the struct) when 
// it's hard to know what to name a value or it's not important
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// Each goroutine adds its result to the map
			// results[url] = wc(url)
			resultChannel <- result{url, wc(url)}
		}()
	}
	
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}