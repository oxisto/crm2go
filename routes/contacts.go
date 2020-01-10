package routes

import (
	"crm"
	"crm/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllContacts(c *gin.Context) {
	var (
		err      error
		contacts []crm.Contact
	)

	if contacts, err = db.GetAllContacts(); err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}

func GetContact(c *gin.Context) {
	var (
		contact *crm.Contact
		id      int64
		err     error
	)

	if id, err = IntParam(c, "id"); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	contact, err = db.GetContact(id)
	JSON(c, http.StatusOK, contact, err)
}

func PostContact(c *gin.Context) {
	var (
		err     error
		contact crm.Contact
	)

	if err = c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err = db.Insert(&contact); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	log.Printf("Inserted contact %d", contact.ID)

	JSON(c, http.StatusOK, contact, err)
}

func PutContact(c *gin.Context) {
	var (
		id      int64
		rows    int64
		err     error
		contact crm.Contact
	)

	if id, err = IntParam(c, "id"); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if err = c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	// make sure id is set
	contact.ID = id

	if rows, err = db.Update(&contact); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	log.Printf("Updated %d rows for contact %d", rows, contact.ID)

	JSON(c, http.StatusOK, contact, err)
}
