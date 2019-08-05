package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//port
const (
	ListeningPort = ":8081"
)

//response struct
type Response struct {
	Title  string
	Detail string
}

func helloWord(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("helloword"))
	resp := Response{}
	resp.Title = "selamat datang"
	resp.Detail = "jelas online golang"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

//penyimpanan
type Categories struct {
	gorm.Model
	Name string `form:"name" json:"name"`
}
type Products struct {
	gorm.Model
	Item        string `form:"item" json:"item"`
	Category_id string `form:"category_id" json:"category_id"`
}

// handler mrupakan kumpuan dari aktivitas siswa
type Handler struct {
	DB *SQLORM
}

//get semua product
func (h *Handler) ReadAll(w http.ResponseWriter, r *http.Request) {
	SemuaCategories := []Categories{}
	h.DB.db.Find(&SemuaCategories)
	encodedResp, err := json.Marshal(SemuaCategories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	SemuaProducts := []Products{}
	h.DB.db.Find(&SemuaProducts)
	encodedResp, err := json.Marshal(SemuaProducts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

func (h *Handler) createCategories(w http.ResponseWriter, r *http.Request) {
	// Get the Body
	categoriesBaru := Categories{}
	// mengubah inputan json menjadi struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&categoriesBaru)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	h.DB.db.Create(&categoriesBaru)
	resp := Response{}
	resp.Title = "sukses"
	resp.Detail = "sudah masuk"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	// Get the Body
	ProductsBaru := Products{}
	// mengubah inputan json menjadi struct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ProductsBaru)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	h.DB.db.Create(&ProductsBaru)
	resp := Response{}
	resp.Title = "sukses"
	resp.Detail = "sudah masuk"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

//function detail categories
func (h *Handler) GetDetailCategories(w http.ResponseWriter, r *http.Request) {
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var CategoriesTertentu Categories
	h.DB.db.First(&CategoriesTertentu, id)
	// searching siswa tertentu

	//kalau siswa id tidak di temukan
	if CategoriesTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	encodedResp, err := json.Marshal(CategoriesTertentu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

func (h *Handler) GetDetailProducts(w http.ResponseWriter, r *http.Request) {
	// mengambil id
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var ProductsTertentu Products
	h.DB.db.First(&ProductsTertentu, id)

	// jika tidak ditemukan
	if ProductsTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id tidak ditemukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}
		w.Write(encodedResp)
		return
	}
	// jika ditemukan
	encodedResp, err := json.Marshal(ProductsTertentu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func (h *Handler) updateCategories(w http.ResponseWriter, r *http.Request) {
	// decode json body
	categoriesUpdate := Categories{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&categoriesUpdate)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	//temukan id
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var CategoriesTertentu Categories
	h.DB.db.First(&CategoriesTertentu, id)
	// searching siswa tertentu
	//kalau siswa id tidak di temukan
	if CategoriesTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	CategoriesTertentu.Name = categoriesUpdate.Name
	h.DB.db.Save(&CategoriesTertentu)
	encodedResp, err := json.Marshal(categoriesUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}
func (h *Handler) updateProducts(w http.ResponseWriter, r *http.Request) {
	ProductsUpdate := Products{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&ProductsUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	//temukan idnya
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var ProductsTertentu Products
	h.DB.db.First(&ProductsTertentu, id)

	// jika tidak ditemukan
	if ProductsTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}
		w.Write(encodedResp)
		return
	}
	//jika ditemukan
	ProductsTertentu.Item = ProductsUpdate.Item
	h.DB.db.Save(&ProductsTertentu)
	encodedResp, err := json.Marshal(ProductsUpdate)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}
	w.Write(encodedResp)
}

func (h *Handler) deleteSiswa(w http.ResponseWriter, r *http.Request) {
	//temukan id
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var CategoriesTertentu Categories
	h.DB.db.First(&CategoriesTertentu, id)
	//kalau siswa id tidak di temukan
	if CategoriesTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id tidak di temukan"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}

		w.Write(encodedResp)
		return
	}
	//kalau di temukan
	h.DB.db.Delete(&CategoriesTertentu)
	resp := Response{}
	resp.Detail = "categoriesCategories berhasil di hapus"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

	w.Write(encodedResp)
	return
}
func (h *Handler) DeleteProducts(w http.ResponseWriter, r *http.Request) {
	id, err := getVarsID(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
		return
	}
	var ProductsTertentu Products
	h.DB.db.First(&ProductsTertentu, id)
	if ProductsTertentu.ID != uint(id) {
		resp := Response{}
		resp.Detail = "id not found"
		encodedResp, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(""))
		}
		w.Write(encodedResp)
		return
	}

	// jika ditemukan
	h.DB.db.Delete(ProductsTertentu)
	resp := Response{}
	resp.Detail = "product berhasil di hapus"
	encodedResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(""))
	}

	w.Write(encodedResp)
	return
}

func getVarsID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	if val, ok := vars["id"]; ok {
		convertedVal, err := strconv.Atoi(val)
		if err != nil {
			return id, err
		}
		id = convertedVal
	}
	return
}

type SQLORM struct {
	db *gorm.DB
}

func connect() *SQLORM {
	ormDB, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/logistik?parseTime=true")
	if err != nil {
		log.Printf("Unable to open mysql DB: %v", err)
		return nil
	}
	// Migrate Schema
	ormDB.AutoMigrate(&Categories{})
	db := new(SQLORM)
	db.db = ormDB
	if ok := db.db.HasTable(&Categories{}); !ok {
		log.Println("No Siswa Table found in DB")
		return nil
	}
	// Migrate Schema product
	ormDB.AutoMigrate(&Products{})
	db.db = ormDB
	if ok := db.db.HasTable(&Products{}); !ok {
		log.Println("No Siswa Table found in DB")
		return nil
	}
	return db
}

func main() {
	db := connect()
	if db == nil {
		log.Println("Unable to Connect to DB")
		return
	}
	log.Println("Successfully Connected to DB")
	//handler
	handler := new(Handler)
	handler.DB = db
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", helloWord).Methods(http.MethodGet)
	r.HandleFunc("/api/categories", handler.ReadAll).Methods(http.MethodGet)
	r.HandleFunc("/api/categories", handler.createCategories).Methods(http.MethodPost)
	r.HandleFunc("/api/categories/{id:[0-9]+}", handler.GetDetailCategories).Methods(http.MethodGet)
	r.HandleFunc("/api/categories/{id:[0-9]+}", handler.updateCategories).Methods(http.MethodPatch)
	r.HandleFunc("/api/categories/{id:[0-9]+}", handler.deleteSiswa).Methods(http.MethodDelete)
	//product
	r.HandleFunc("/api/products", handler.GetProducts).Methods(http.MethodGet)
	r.HandleFunc("/api/products/{id:[0-9]+}", handler.GetDetailProducts).Methods(http.MethodGet)
	r.HandleFunc("/api/products", handler.createItem).Methods(http.MethodPost)
	r.HandleFunc("/api/products/{id:[0-9]+}", handler.updateProducts).Methods(http.MethodPatch)
	r.HandleFunc("/api/products/{id:[0-9]+}", handler.DeleteProducts).Methods(http.MethodDelete)

	// running port
	log.Printf("Starting http server at %v", ListeningPort)
	err := http.ListenAndServe(ListeningPort, r)
	if err != nil {
		log.Fatalf("Unable to run http server: %v", err)
	}
	log.Println("Stopping API Service...")
}
