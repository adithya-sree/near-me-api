package handler

import (
	"fmt"
	"nearme-api/src/app/db"
	"net/http"
	"time"
)

type addLocationRequest struct {
	username string
	location string
}

//AddLocation adds and saves location object do db
func (h *Handler) AddLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddLocation request received")
	header, err := checkForAddLocationHeaders(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	l, err := h.FindByUsername(header.username)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	now := time.Now()

	if l.Username == "" {
		fmt.Println("Inserting new location entry")
		err = h.Insert(db.Location{
			Username:   header.username,
			Location:   header.location,
			Created:    now.Format("2006.01.02 15:04:05"),
			LastUpdate: now.Format("2006.01.02 15:04:05"),
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		fmt.Println("Updating existing location entry")
		err = h.Update(l, db.Location{
			Username:   l.Username,
			Location:   header.location,
			Created:    l.Created,
			LastUpdate: now.Format("2006.01.02 15:04:05"),
		})

		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondJSON(w, http.StatusOK, Response{Message: "successfully added entry"})
}

//GetLocation get a location object from the db
func (h *Handler) GetLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetLocation request received")
	username, err := checkForGetLocationHeaders(r)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.FindByUsername(username)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if result.Username == "" {
		fmt.Println("No location found for user", username)
		respondError(w, http.StatusNotFound, "user was not found")
		return
	}

	respondJSON(w, http.StatusOK, result)
}

func checkForAddLocationHeaders(r *http.Request) (addLocationRequest, error) {
	username, err := checkforHeader(r, "username")
	if err != nil {
		return addLocationRequest{}, err
	}

	location, err := checkforHeader(r, "location")
	if err != nil {
		return addLocationRequest{}, err
	}

	request := addLocationRequest{
		username: username,
		location: location,
	}
	return request, nil
}

func checkForGetLocationHeaders(r *http.Request) (string, error) {
	username, err := checkforHeader(r, "username")
	if err != nil {
		return "", err
	}
	return username, nil
}
