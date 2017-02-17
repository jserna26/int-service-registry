package create

import logging "github.com/op/go-logging"

var log = logging.MustGetLogger("Create system test")

/*
func TestCreateSystem(t *testing.T) {
	r, _ := repo.NewInMemorySystemRepo("")
	h := NewCreateSystemHandler(r)
	newSys := types.NewSystemType{
		Name:        "Test",
		Description: "sdfsdf",
		Status:      types.StatusEnabled,
	}
	res, err := plumb.DoHTTP(h, http.MethodPost, plumb.StructToBody(newSys), nil)
	if err != nil {
		t.Error(err)
	}

	var sysOut types.SystemType
	json.NewDecoder(res.Body).Decode(&sysOut)
	if sysOut.ID != "1" || sysOut.Name != "Test" {
		t.Errorf("ID is %s. Name is %s", sysOut.ID, sysOut.Name)
	}
}

func TestCreateSystemNameTooLong(t *testing.T) {
	r, _ := repo.NewInMemorySystemRepo("")
	h := CreateSystemHandler(r)
	newSys := types.NewSystemType{
		Name:        "123456789012345678901234567890123456789012345678901",
		Description: "sdfsdf",
		Status:      types.StatusEnabled,
	}
	res, err := plumb.DoHTTP(h, http.MethodPost, plumb.StructToBody(newSys), nil)
	if err != nil {
		t.Error(err)
	}

	var respObj types.ErrorType
	err = json.NewDecoder(res.Body).Decode(&respObj)
	if err != nil {
		t.Error(err)
	}
	if respObj.Message != "System name can be max 50 chars" {
		t.Errorf("Did not receive 50 char error message %v", respObj)
	}

}*/
