package contacts

import (
	"errors"
	"log"
	"sort"
	"strings"
)

type SortById Contacts

func (a SortById) Len() int           { return len(a) }
func (a SortById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortById) Less(i, j int) bool { return a[i].Id < a[j].Id }

func (contacts Contacts) GetNextId() int {
	maxId := 0
	for _, contact := range contacts {
		if maxId < contact.Id {
			maxId = contact.Id
		}
	}
	return maxId + 1
}

func (contacts Contacts) GetNewContact(contact *Contact) *Contact {
	contact.Id = contacts.GetNextId()
	return contact
}

func (contacts Contacts) Delete(id int) Contacts {
	log.Printf("Deleting Id: %d", id)
	sort.Sort(SortById(contacts))
	for i, c := range contacts {
		if id == c.Id {
			contacts = append(contacts[:i], contacts[i+1:]...)
		}
	}
	return contacts
}

func (contacts Contacts) Log() {
	log.Printf("Printing %d items of contacts", len(contacts))
	for i := range contacts {
		log.Println(*contacts[i])
	}
}

func (contacts Contacts) ById(id int) *Contact {
	for i := range contacts {
		if contacts[i].Id == id {
			return contacts[i]
		}
	}
	return nil
}

func (contacts Contacts) IsValidEmail(id int, email string) (err error) {
	log.Printf("Validating email: %s for id: %d", email, id)
	for i := range contacts {
		if contacts[i].Id != id && strings.EqualFold(contacts[i].Email, email) {
			err = errors.New("email must be unique")
		}
	}
	return err
}
