package lecturer_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/belajarRestApi5/config"
	"github.com/belajarRestApi5/services/lecturer"
)

var serverTest config.Server

func TestMain(m *testing.M) {

	serverTest = config.Server{}
	serverTest.Initialize()  //initialize database
	serverTest.SetupServer() //initialize all endpoint url that available

	fmt.Println("Do stuff BEFORE the tests! UP")

	code := m.Run()

	fmt.Println("Do stuff AFTER the tests! DOWN")

	os.Exit(code)

}

func TestGetAllLecturer(t *testing.T) {

	srv := httptest.NewServer(serverTest.Router) // make serverTest
	defer srv.Close()                            //	close when operation finished

	fmt.Println(srv.URL) //base url server test

	res, err := http.Get(fmt.Sprintf("%s/lecturer", srv.URL)) //hit endpoint url to server test
	if err != nil {
		fmt.Println("error")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.StatusCode)
	}
}

func TestGetLecturer(t *testing.T) {

	srv := httptest.NewServer(serverTest.Router) // make serverTest
	defer srv.Close()                            //	close when operation finished

	lecturerActual := lecturer.Lecturer{}
	lecturerActual.ID = 1
	lecturerActual.Name = "Dhuki Dwi R"
	lecturerActual.Address = "Bandung"
	lecturerActual.AddLecturer(serverTest.DB)

	res, err := http.Get(fmt.Sprintf("%s/api/lecturer/1", srv.URL)) //	hit endpoint url to server test (srv.URL == baseUrl)
	if err != nil {
		fmt.Println("error")
	}

	bodyBytes, error := ioutil.ReadAll(res.Body) //	change response to []byte so we can convert it
	defer res.Body.Close()
	if error != nil {
		log.Fatal(err)
	}

	var lecturerResult = new(lecturer.Lecturer)
	err = json.Unmarshal(bodyBytes, &lecturerResult) //	convert from []byte using json Unmarshall to Object
	if err != nil {
		fmt.Println("whoops:", err)
	}

	if res.StatusCode != http.StatusOK { // check if status is OK (200)
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, res.StatusCode)
	}

	if lecturerActual.ID != lecturerResult.ID {
		t.Errorf("Unexpected response %d. Got %d\n", lecturerActual.ID, lecturerResult.ID)
	}
}

func clearTable() {
	serverTest.DB.Exec("DELETE FROM lecturer")
	serverTest.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}
