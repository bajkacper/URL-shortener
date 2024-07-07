package handlers

import (
		"net/http"
	)

	func RedirectURL(w http.ResponseWriter, r *http.Request) {
			shortKey := r.URL.Path[len("/short/"):]
				if shortKey == "" {
							http.Error(w, "Shortened key missing", http.StatusBadRequest)
									return
										}

											originalURL, exists := urlMap[shortKey]
												if !exists {
															http.Error(w, "Shortened key not found", http.StatusNotFound)
																	return
																		}

																			http.Redirect(w, r, originalURL, http.StatusFound)
																		}

