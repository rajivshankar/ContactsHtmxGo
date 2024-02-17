package main

import (
	"contacts/contacts"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var DbContacts contacts.Contacts

func InitialiseContacts() {
	log.Println("Initializing Contacts")
	DbContacts = *new(contacts.Contacts)
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Rajiv", Email: "rajiv@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Kartik", Email: "Kamz@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "NP", Email: "NP@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Kate", Email: "kate@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Gayathri", Email: "gayathri@gmail.com"}))

	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Marcin", Email: "marcin@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Sumanth", Email: "sumanth@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Praths", Email: "praths@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Sankaran", Email: "appa@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Ganesh", Email: "ganesh@gmail.com"}))

	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Siva", Email: "siva@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Santhosh", Email: "santhosh@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Ahmad", Email: "ahmad@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Jose", Email: "jose@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Rajkumar", Email: "pooch@gmail.com"}))

	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Shabana", Email: "shabana@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Paul", Email: "feeney@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Sathish", Email: "sathish@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Preetham", Email: "preeths@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Ranjini", Email: "ranjini@gmail.com"}))

	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Anand Natraj", Email: "natz@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Mahesh S", Email: "maggu@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Ramesh", Email: "rams@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Karthik j", Email: "jka@gmail.com"}))
	DbContacts = append(DbContacts, DbContacts.GetNewContact(&contacts.Contact{Name: "Naarayan", Email: "rna@gmail.com"}))
}

func setRoutesForContacts(r *gin.Engine) {
	InitialiseContacts()
	routes := r.Group("/contacts")
	{
		routes.GET("/", ContactsHome())
		routes.GET("/all", GetContactsByPage(false))
		routes.GET("/reset", GetContactsByPage(true))
		routes.POST("/add", AddContact())
		routes.DELETE("/delete/:id", DeleteContact())
		routes.GET("/edit/:id", EditContact())
		routes.POST("/save/:id", FindContact(true))
		routes.GET("/find/:id", FindContact(false))
		routes.GET("/err/:id", ErrMsg())
		routes.GET("/validate/email", ValidateEmail())
	}
}

func ContactsHome() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Calling Index")
		ctx.HTML(http.StatusOK, "html/index.tmpl", gin.H{"contacts": DbContacts})
	}
}

func GetContacts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "html/contact_list.tmpl", gin.H{"contacts": DbContacts})
	}
}

func AddContact() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var contact *contacts.Contact
		err := ctx.ShouldBind(&contact)
		if err == nil {
			log.Printf("Added contact with name: %q and email: %q", contact.Name, contact.Email)
			contact = DbContacts.GetNewContact(contact)
			DbContacts = append(DbContacts, contact)
			var newContacts contacts.Contacts
			newContacts = append(newContacts, contact)
			ctx.HTML(http.StatusCreated, "html/contact_list.tmpl", gin.H{"contacts": newContacts})
		} else {
			log.Printf("Error [%s] whilst adding contact with name: %q and email: %q", err.Error(), contact.Name, contact.Email)
			ctx.HTML(http.StatusForbidden, "html/contact_list.tmpl", nil)
			ctx.AbortWithStatus(http.StatusNotImplemented)
		}
	}
}

func DeleteContact() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Deleting contact with ID: %q", ctx.Param("id"))
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Printf("Error [%s] occured when getting ID", err.Error())
			ctx.AbortWithStatus(http.StatusNotImplemented)
		} else {
			DbContacts = DbContacts.Delete(id)
			ctx.AbortWithStatus(http.StatusOK)
		}
	}
}

func EditContact() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Editing contact with ID: %q", ctx.Param("id"))
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Printf("Error [%s] occured when converting id", err.Error())
			ctx.AbortWithStatus(http.StatusBadRequest)
		} else {
			contact := DbContacts.ById(id)
			ctx.HTML(http.StatusOK,
				"html/contact_edit.tmpl",
				gin.H{"id": contact.Id, "name": contact.Name, "email": contact.Email})
		}
	}
}

func FindContact(save bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginH := make(gin.H)
		log.Printf("Retrieving contact with ID: %q", ctx.Param("id"))
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			log.Printf("Error [%s] occured when converting id", err.Error())
			// ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
			ginH["status"] = "error"
			ginH["message"] = err.Error()
		} else {
			if save {
				var contact *contacts.Contact
				log.Println("Saving Contact")
				err := ctx.ShouldBind(&contact)
				if err != nil {
					log.Printf("Error [%s] whilst adding contact with name: %q and email: %q", err.Error(), contact.Name, contact.Email)
					ginH["status"] = "error"
					ginH["message"] = err.Error()
					// ctx.HTML(http.StatusBadRequest, "html/contact_edit.tmpl", gin.H{"status": "Error", "message": err.Error()})
				} else {
					err = DbContacts.IsValidEmail(contact.Id, contact.Email)
					if err != nil {
						log.Printf("error: %s", err.Error())
						// ctx.HTML(http.StatusOK, "html/error.tmpl", gin.H{"status": "error", "message": err.Error()})
						ginH["status"] = "error"
						ginH["message"] = err.Error()
					} else {
						log.Printf("Saving contact with id: %d,  name: %q and email: %q", contact.Id, contact.Name, contact.Email)
						DbContacts = DbContacts.Delete(id)
						DbContacts = append(DbContacts, contact)
						sort.Sort(contacts.SortById(DbContacts))
						log.Printf("Successfully Saved contact with id: %d,  name: %q and email: %q", contact.Id, contact.Name, contact.Email)
						var newContacts contacts.Contacts
						newContacts = append(newContacts, contact)
						ctx.HTML(http.StatusOK, "html/contact_list.tmpl", gin.H{"contacts": newContacts})
						return
					}
				}
				ctx.HTML(http.StatusBadRequest, "html/contact_edit.tmpl", ginH)
				return
			} else {
				contact := DbContacts.ById(id)
				var newContacts contacts.Contacts
				newContacts = append(newContacts, contact)
				ctx.HTML(http.StatusOK, "html/contact_list.tmpl", gin.H{"contacts": newContacts})
				return
			}
		}
	}
}

func ErrMsg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func ValidateEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			id = -1
		}
		email := ctx.Query("email")
		log.Printf("Validating email: %s for id: %d", email, id)
		err = DbContacts.IsValidEmail(id, email)
		if err == nil {
			ctx.HTML(http.StatusOK, "html/error.tmpl", nil)
		} else {
			log.Printf("error: %s", err.Error())
			ctx.HTML(http.StatusOK, "html/error.tmpl", gin.H{"status": "error", "message": err.Error()})
		}
	}
}

func GetContactsByPage(reset bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginH := make(gin.H)
		if reset {
			InitialiseContacts()
		}
		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		full := ctx.Query("full")
		if strings.EqualFold(full, "Y") {
			ginH["full"] = "Y"
		}
		log.Printf("Retrieving Page Number: %d", page)
		page_length := 10
		num_pages := int(len(DbContacts) / page_length)
		if len(DbContacts)%page_length > 1 {
			num_pages += 1
		}
		log.Printf("Number of Pages: %d", num_pages)
		if page > num_pages && num_pages > 0 {
			page = num_pages
		}
		log.Printf("number of contacts; %d", len(DbContacts))
		log.Printf("page length: %d", page_length)
		start := page_length * (page - 1)
		log.Printf("start: %d", start)
		offset := page_length
		if (start + offset) > len(DbContacts) {
			offset = len(DbContacts) - start
		}
		log.Printf("offset: %d", offset)
		result := DbContacts[start:(start + offset)]
		ginH["contacts"] = result
		if page > 1 {
			ginH["prev"] = page - 1
			log.Printf("Prev: %d", ginH["prev"])
		}
		if page < num_pages {
			ginH["next"] = page + 1
			log.Printf("Next: %d", ginH["next"])
		}
		ctx.HTML(http.StatusOK, "html/contact_list.tmpl", ginH)
	}
}
