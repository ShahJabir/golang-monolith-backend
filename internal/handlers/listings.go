package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListingHandler struct {
	db *sql.DB
}

func NewListingHandler(db *sql.DB) *ListingHandler {
	return &ListingHandler{
		db: db,
	}
}

func (lh *ListingHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := lh.db.Query(
		`SELECT id, title, description, price, city, created_at
			FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
	if err != nil {
		log.Printf("Error querying listings: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	// Process the rows and write the response
	listings := []listing{}
	for rows.Next() {
		var l listing
		err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
		if err != nil {
			log.Printf("Error scanning listing row: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		listings = append(listings, l)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over listing rows: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(listings); err != nil {
		log.Printf("Error encoding listings to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (lh *ListingHandler) Get(w http.ResponseWriter, r *http.Request) {
	// Extract the listing ID from the request
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing listing ID", http.StatusBadRequest)
		return
	}

	// Query the database for the listing
	row := lh.db.QueryRow(
		`SELECT id, title, description, price, city, created_at
			FROM listings
			WHERE id = $1`,
		id,
	)
	var l listing
	err := row.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Listing not found", http.StatusNotFound)
			return
		}
		log.Printf("Error scanning listing row: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(l); err != nil {
		log.Printf("Error encoding listing to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (lh *ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {

	// Extract the listing ID from the request
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Missing listing ID", http.StatusBadRequest)
		return
	}

	// Delete the listing from the database
	result, err := lh.db.Exec(
		"DELETE FROM listings WHERE id = $1",
		id,
	)
	if err != nil {
		log.Printf("Error deleting listing: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Listing not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
