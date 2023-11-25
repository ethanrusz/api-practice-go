package models

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "dummy"
	database = "albums"
)

func GetAlbums() []Album {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	results, err := db.Query("select * from album")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	var albums []Album
	for results.Next() {
		var album Album
		err = results.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil
		}

		albums = append(albums, album)
	}
	return albums
}

func GetAlbumByID(id string) *Album {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	defer db.Close()

	album := &Album{}
	results, err := db.Query("select * from album where id=$1", id)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return album
}

func AddAlbum(album Album) {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s database=%s sslmode=disable", host, port, user, password, database)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query(
		"insert into album (id, title, artist, price) values ($1, $2, $3, $4)",
		album.ID, album.Title, album.Artist, album.Price)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}
