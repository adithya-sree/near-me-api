package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Client client instance for sql
type Client struct {
	*sql.DB
}

//NewClient new instance of Client
func NewClient(connection string) (*Client, error) {
	db, err := sql.Open("mysql", connection)
	if err != nil {
		return nil, err
	}

	return &Client{
		DB: db,
	}, nil
}

//Insert inserts location object
func (c *Client) Insert(l Location) error {
	insert, err := c.Query(fmt.Sprintf("INSERT INTO `location` (`username`, `created`, `lastUpdate`, `location`) VALUES ('%s', '%s', '%s', '%s');", l.Username, l.Created, l.LastUpdate, l.Location))
	if err != nil {
		defer insert.Close()
	}

	return err
}

//Update updates location object
func (c *Client) Update(old Location, new Location) error {
	update, err := c.Query(fmt.Sprintf("UPDATE `location` SET username=\"%s\", created=\"%s\", lastUpdate=\"%s\", location=\"%s\" WHERE entryId=%d", new.Username, new.Created, new.LastUpdate, new.Location, old.EntryID))
	if err != nil {
		defer update.Close()
	}

	return err
}

//FindByUsername returns location by username
func (c *Client) FindByUsername(username string) (Location, error) {
	results, err := c.Query(fmt.Sprintf("SELECT * FROM `location` where username=\"%s\";", username))
	if err != nil {
		return Location{}, err
	}

	var l Location
	for results.Next() {
		err = results.Scan(&l.Username, &l.Created, &l.LastUpdate, &l.Location, &l.EntryID)
	}

	defer results.Close()
	return l, nil
}
