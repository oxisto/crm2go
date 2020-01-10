package routes

import (
	"crm"
	"crm/db"
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

	// @todo Check permissions for GetCharacter call

	if id, err = IntParam(c, "id"); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	contact, err = db.GetContact(id)
	JSON(c, http.StatusOK, contact, err)
}

func PostContact(c *gin.Context) {

}
