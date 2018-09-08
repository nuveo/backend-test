package workflow

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// GET /workflow
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetWorkflows() // list of all products
	log.Println(products)
	data, _ := json.Marshal(products)
	log.Printf("jsonData: %s\n", data)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// POST /workflow
func (c *Controller) AddWorkflow(w http.ResponseWriter, r *http.Request) {
	var workflow Workflow

	// read the body of the request
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	log.Println(body)

	if err != nil {
		log.Fatalln("Error addWorkflow", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error addWorkflow", err)
	}

	if err := json.Unmarshal(body, &workflow); err != nil {
		// unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddProduct unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(workflow)
	// adds the product to the DB
	success := c.Repository.AddWorkflow(workflow)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(workflow)
	w.Write(data)
	return
}

// // SearchProduct GET /
// func (c *Controller) SearchProduct(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Println(vars)

// 	query := vars["query"] // param query
// 	log.Println("Search Query - " + query)

// 	products := c.Repository.GetProductsByString(query)
// 	data, _ := json.Marshal(products)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(data)
// 	return
// }

// // UpdateProduct PUT /
// func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
// 	var product Product
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
// 	if err != nil {
// 		log.Fatalln("Error UpdateProduct", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	if err := r.Body.Close(); err != nil {
// 		log.Fatalln("Error UpdateProduct", err)
// 	}

// 	if err := json.Unmarshal(body, &product); err != nil { // unmarshall body contents as a type Candidate
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(422) // unprocessable entity
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			log.Fatalln("Error UpdateProduct unmarshalling data", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	log.Println(product.ID)
// 	success := c.Repository.UpdateProduct(product) // updates the product in the DB

// 	if !success {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	return
// }

// // GetProduct GET - Gets a single product by ID /
// func (c *Controller) GetProduct(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Println(vars)

// 	id := vars["id"] // param id
// 	log.Println(id)

// 	productid, err := strconv.Atoi(id)

// 	if err != nil {
// 		log.Fatalln("Error GetProduct", err)
// 	}

// 	product := c.Repository.GetProductById(productid)
// 	data, _ := json.Marshal(product)

// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(data)
// 	return
// }

// // DeleteProduct DELETE /
// func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	log.Println(vars)
// 	id := vars["id"] // param id
// 	log.Println(id)

// 	productid, err := strconv.Atoi(id)

// 	if err != nil {
// 		log.Fatalln("Error GetProduct", err)
// 	}

// 	if err := c.Repository.DeleteProduct(productid); err != "" { // delete a product by id
// 		log.Println(err)
// 		if strings.Contains(err, "404") {
// 			w.WriteHeader(http.StatusNotFound)
// 		} else if strings.Contains(err, "500") {
// 			w.WriteHeader(http.StatusInternalServerError)
// 		}
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	return
// }
