package handlers

import (
		"fmt"
			"my-url-shortener/utils"
				"net/http"
			)

			var urlMap = make(map[string]string)

			func ShortenURL(w http.ResponseWriter, r *http.Request) {
					if r.Method != http.MethodPost {
								http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
										return
											}

												originalURL := r.FormValue("url")
													if originalURL == "" {
																http.Error(w, "URL parameter missing", http.StatusBadRequest)
																		return
																			}

																				shortKey := utils.GenerateKey()
																					urlMap[shortKey] = originalURL

																						shortenedURL := fmt.Sprintf("http://localhost:8080/short/%s", shortKey)

																							w.Header().Set("Content-Type", "text/html")
																								responseHTML := fmt.Sprintf(`
																										<h2>URL Shortener Service</h2>
																												<p>Original URL: %s</p>
																														<p>Shortened URL: <a href="%s">%s</a></p>
																																<form method="post" action="/shorten">
																																			<input type="text" name="url" placeholder="Enter a URL">
																																						<input type="submit" value="Shorten">
																																								</form>
																																									`, originalURL, shortenedURL, shortenedURL)
																																										fmt.Fprintf(w, responseHTML)
																																									}

