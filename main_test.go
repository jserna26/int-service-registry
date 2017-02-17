package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jserna26/SystemPOC/types"
)

type FakeRepo struct{}

func (f FakeRepo) CreateSystem(sys types.NewSystemType) (system types.SystemType, err error) {
	system.ID = "1"
	system.Name = sys.Name
	return
}

func (f FakeRepo) GetAllSystems() (systems []types.SystemType, err error) {
	var system types.SystemType
	system.Name = "1"
	system.Description = "Test description"
	system.Status = types.StatusEnabled
	systems = append(systems, system)
	return
}

func (f FakeRepo) GetSystem(systemName string) (system types.SystemType, err error) {
	system.Name = "1"
	system.Description = "Received test system"
	system.Status = types.StatusEnabled
	return
}

func (f FakeRepo) Connect(dburl string) (err error) {
	return
}
func (f FakeRepo) MustConnect(dburl string) {}
func (f FakeRepo) Disconnect() (err error) {
	return
}

func TestGetDbUrl(t *testing.T) {
	os.Setenv("VCAP_SERVICES", `{ "postgres": [   {    "credentials": {     "ID": 0,     "binding_id": "f888ef92-8c02-40ee-be7f-cb3ab943b299",     "database": "da650336a12d64fe88e221dfabe603fdc",     "dsn": "host=10.72.6.143 port=5432 user=ue6bfc320cb87474e82847dcc60cfb617 password=da2a616c27d04d61a982df34d29e79da dbname=da650336a12d64fe88e221dfabe603fdc connect_timeout=5 sslmode=disable",     "host": "10.72.6.143",     "instance_id": "4a4f5e6d-a796-438d-910c-95d854833a34",     "jdbc_uri": "jdbc:postgresql://10.72.6.143:5432/da650336a12d64fe88e221dfabe603fdc?user=ue6bfc320cb87474e82847dcc60cfb617\u0026password=da2a616c27d04d61a982df34d29e79da\u0026ssl=false",     "password": "da2a616c27d04d61a982df34d29e79da",     "port": "5432",     "uri": "testuri",     "username": "ue6bfc320cb87474e82847dcc60cfb617"    },    "label": "postgres",    "name": "Int-Systems",    "plan": "shared-nr",    "provider": null,    "syslog_drain_url": null,    "tags": [     "rdpg",     "postgresql"    ],    "volume_mounts": []   }  ] }`)
	os.Setenv("VCAP_APPLICATION", `{  "application_id": "e248abc1-44ba-4eca-a936-68da36b8a6cd",  "application_name": "Predix-Int-Platform-Systems",  "application_uris": [   "predix-int-platform-systems.run.aws-usw02-pr.ice.predix.io"  ],  "application_version": "d7569a01-5439-431d-a2df-7c7cfc119b45",  "limits": {   "disk": 1024,   "fds": 16384,   "mem": 64  },  "name": "Predix-Int-Platform-Systems",  "space_id": "f55809cc-f980-4237-ac9c-81208ca6f320",  "space_name": "dev",  "uris": [   "predix-int-platform-systems.run.aws-usw02-pr.ice.predix.io"  ],  "users": null,  "version": "d7569a01-5439-431d-a2df-7c7cfc119b45" }`)
	dburl, err := getDbUrl("Int-Systems", "uri")
	if err != nil {
		t.Error(err)
	}
	if dburl != "testuri" {
		t.Errorf("Expected testuri. Got %s", dburl)
	}
}

func TestFailGetDbUrlEnvVarsNotSet(t *testing.T) {
	os.Setenv("VCAP_SERVICES", `"sdfsdf"`)
	os.Setenv("VCAP_APPLICATION", `"sdfsdf"`)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic: %v", r)
		}
	}()
	dburl, err := getDbUrl("test", "uri")
	fmt.Printf("%s%s", dburl, err)
}

func TestFailGetDbUrlIncorectServiceName(t *testing.T) {

	os.Setenv("VCAP_SERVICES", `{ "postgres": [   {    "credentials": {     "ID": 0,     "binding_id": "f888ef92-8c02-40ee-be7f-cb3ab943b299",     "database": "da650336a12d64fe88e221dfabe603fdc",     "dsn": "host=10.72.6.143 port=5432 user=ue6bfc320cb87474e82847dcc60cfb617 password=da2a616c27d04d61a982df34d29e79da dbname=da650336a12d64fe88e221dfabe603fdc connect_timeout=5 sslmode=disable",     "host": "10.72.6.143",     "instance_id": "4a4f5e6d-a796-438d-910c-95d854833a34",     "jdbc_uri": "jdbc:postgresql://10.72.6.143:5432/da650336a12d64fe88e221dfabe603fdc?user=ue6bfc320cb87474e82847dcc60cfb617\u0026password=da2a616c27d04d61a982df34d29e79da\u0026ssl=false",     "password": "da2a616c27d04d61a982df34d29e79da",     "port": "5432",     "uri": "testuri",     "username": "ue6bfc320cb87474e82847dcc60cfb617"    },    "label": "postgres",    "name": "Int-Systems",    "plan": "shared-nr",    "provider": null,    "syslog_drain_url": null,    "tags": [     "rdpg",     "postgresql"    ],    "volume_mounts": []   }  ] }`)
	os.Setenv("VCAP_APPLICATION", `{  "application_id": "e248abc1-44ba-4eca-a936-68da36b8a6cd",  "application_name": "Predix-Int-Platform-Systems",  "application_uris": [   "predix-int-platform-systems.run.aws-usw02-pr.ice.predix.io"  ],  "application_version": "d7569a01-5439-431d-a2df-7c7cfc119b45",  "limits": {   "disk": 1024,   "fds": 16384,   "mem": 64  },  "name": "Predix-Int-Platform-Systems",  "space_id": "f55809cc-f980-4237-ac9c-81208ca6f320",  "space_name": "dev",  "uris": [   "predix-int-platform-systems.run.aws-usw02-pr.ice.predix.io"  ],  "users": null,  "version": "d7569a01-5439-431d-a2df-7c7cfc119b45" }`)

	dburl, err := getDbUrl("test", "uri")
	if err != nil {
		t.Log(err)
	}
	if dburl != "" {
		t.Errorf("Got Dburl %s. Expected blank dburl due to no env vars set", dburl)
	}
}

func TestGetInitRoutes(t *testing.T) {
	rt := new(mux.Router)
	ir := getInitRoutes(FakeRepo{})
	ir(rt)
	//plumb.DoHTTP()
}
