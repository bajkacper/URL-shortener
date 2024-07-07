package main

import (
		"fmt"
			"net/http"
				"my-url-shortener/handlers"
			)

			func main() {
					http.HandleFunc("/shorten", handlers.ShortenURL)
						http.HandleFunc("/short/", handlers.RedirectURL)

							fmt.Println("URL Shortener service is running on port 8080")
								http.ListenAndServe(":8080", nil)
							}

