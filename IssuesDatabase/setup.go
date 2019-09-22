package issuesdatabase

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Issue : model for issues
type Issue struct {
	gorm.Model
	Sender  string
	Channel string
}

// IssueQueue : model for issues queue
type IssueQueue struct {
	gorm.Model
	Sender  string
	Channel string
}

// IssuesDatabase allows global access to database
var IssuesDatabase *gorm.DB

// Setup initiate the database
func Setup() error {
	IssuesDatabase, err := gorm.Open("sqlite3", "issues.db")
	if err != nil {
		return err
	}

	defer IssuesDatabase.Close()

	IssuesDatabase.AutoMigrate(&Issue{})
	IssuesDatabase.AutoMigrate(&IssueQueue{})

	return nil
}
