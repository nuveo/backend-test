package workflow

import (
	"fmt"
)

// "fmt"
// "log"
// "strings"

// "gopkg.in/mgo.v2/bson"

//Repository ...
type Repository struct {
	MockWorkflows []Workflow
}

// SERVER the DB server
// const SERVER = "mongodb://gautam:gautam@ds157233.mlab.com:57233/dummystore"

// // DBNAME the name of the DB instance
// const DBNAME = "dummyStore"

// // COLLECTION is the name of the collection in DB
// const COLLECTION = "store"

// var productId = 10

// GetProducts returns the list of Products
func (r *Repository) GetWorkflows() []Workflow {
	// session, err := mgo.Dial(SERVER)

	// if err != nil {
	// 	fmt.Println("Failed to establish connection to Mongo server:", err)
	// }

	// defer session.Close()

	// c := session.DB(DBNAME).C(COLLECTION)

	// u1, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatal("Something went wrong: %s", err)
	// }

	// u2, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatal("Something went wrong: %s", err)
	// }

	// s1 := []string{"step 01"}
	// s2 := []string{"step 01", "step 02"}
	// w1 := Workflow{UUID: u1, Status: Inserted, Steps: s1}
	// w2 := Workflow{UUID: u2, Status: Consumed, Steps: s2}

	// r.MockWorkflows = append(r.MockWorkflows, w1)
	// r.MockWorkflows = append(r.MockWorkflows, w2)

	// 	if err := c.Find(nil).All(&results); err != nil {
	// 		fmt.Println("Failed to write results:", err)
	// 	}
	return r.MockWorkflows
}

// GetProductById returns a unique Product
// func (r Repository) GetProductById(id int) Product {
// 	session, err := mgo.Dial(SERVER)

// 	if err != nil {
// 		fmt.Println("Failed to establish connection to Mongo server:", err)
// 	}

// 	defer session.Close()

// 	c := session.DB(DBNAME).C(COLLECTION)
// 	var result Product

// 	fmt.Println("ID in GetProductById", id)

// 	if err := c.FindId(id).One(&result); err != nil {
// 		fmt.Println("Failed to write result:", err)
// 	}

// 	return result
// }

// // GetProductsByString takes a search string as input and returns products
// func (r Repository) GetProductsByString(query string) Products {
// 	session, err := mgo.Dial(SERVER)

// 	if err != nil {
// 		fmt.Println("Failed to establish connection to Mongo server:", err)
// 	}

// 	defer session.Close()

// 	c := session.DB(DBNAME).C(COLLECTION)
// 	result := Products{}

// 	// Logic to create filter
// 	qs := strings.Split(query, " ")
// 	and := make([]bson.M, len(qs))
// 	for i, q := range qs {
// 		and[i] = bson.M{"title": bson.M{
// 			"$regex": bson.RegEx{Pattern: ".*" + q + ".*", Options: "i"},
// 		}}
// 	}
// 	filter := bson.M{"$and": and}

// 	if err := c.Find(&filter).Limit(5).All(&result); err != nil {
// 		fmt.Println("Failed to write result:", err)
// 	}

// 	return result
// }

// AddWorkflow adds a Workflow in the DB
func (r *Repository) AddWorkflow(workflow Workflow) bool {

	// session, err := mgo.Dial(SERVER)
	// defer session.Close()

	// productId += 1
	// product.ID = productId session.DB(DBNAME).C(COLLECTION).Insert(product)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return false
	// }

	r.MockWorkflows = append(r.MockWorkflows, workflow)
	fmt.Println("Added New Product ID- ", workflow.UUID)

	return true
}

// UpdateProduct updates a Product in the DB
func (r *Repository) UpdateWorkflow(workflowNew Workflow) bool {
	// session, err := mgo.Dial(SERVER)
	// defer session.Close()

	// err = session.DB(DBNAME).C(COLLECTION).UpdateId(product.ID, product)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return false
	// }

	for _, element := range r.MockWorkflows {

		if element.UUID == workflowNew.UUID {

			element.Status = workflowNew.Status
		}

	}
	fmt.Println("Updated Workflow UUID - ", workflowNew.UUID)

	return true
}

// // DeleteProduct deletes an Product
// func (r Repository) DeleteProduct(id int) string {
// 	session, err := mgo.Dial(SERVER)
// 	defer session.Close()

// 	// Remove product
// 	if err = session.DB(DBNAME).C(COLLECTION).RemoveId(id); err != nil {
// 		log.Fatal(err)
// 		return "INTERNAL ERR"
// 	}

// 	fmt.Println("Deleted Product ID - ", id)
// 	// Write status
// 	return "OK"
// }

// Consume by Queue and returns the list of workflows
func (r *Repository) ConsumeWorkflows() []Workflow {
	// session, err := mgo.Dial(SERVER)

	// if err != nil {
	// 	fmt.Println("Failed to establish connection to Mongo server:", err)
	// }

	// defer session.Close()

	// c := session.DB(DBNAME).C(COLLECTION)

	// u1, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatal("Something went wrong: %s", err)
	// }

	// u2, err := uuid.NewV4()
	// if err != nil {
	// 	log.Fatal("Something went wrong: %s", err)
	// }

	// s1 := []string{"step 01"}
	// s2 := []string{"step 01", "step 02"}
	// w1 := Workflow{UUID: u1, Status: Inserted, Steps: s1}
	// w2 := Workflow{UUID: u2, Status: Consumed, Steps: s2}

	// r.MockWorkflows = append(r.MockWorkflows, w1)
	// r.MockWorkflows = append(r.MockWorkflows, w2)

	// 	if err := c.Find(nil).All(&results); err != nil {
	// 		fmt.Println("Failed to write results:", err)
	// 	}
	return r.MockWorkflows
}
