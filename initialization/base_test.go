package initialization

import (
	"testing"
	"mimi/djq/service"
)

func TestTemp(t *testing.T){
	ids := "4d1aa904f2074bc09cd4605bb3d35803,bf59e95391b645bca885dc1ad809da08,7f028d43c7b143b49eb5051b1034e636,159330bbf27a4093bda78a211cc23d77,f9e3869a4a144617afc6111c06c74d25,162531075caf4428a9d9497f77b9ac7e,35bf72294c10498b9ac5bdd03046329d,f9e3869a4a144617afc6111c06c74d25,7e9472f2fb1e4ebb8bb7c300a30912f6,7e9472f2fb1e4ebb8bb7c300a30912f6,4502eac9a78e485bae1d5a278bce11f9,ad1cbb9d6032465a9ab061f703a8fd2c"
	serviceObj := &service.PresentOrder{}
	oo,err := serviceObj.Random("8fb58750aaa3484c823759c67bc01d23",ids)
	if err!=nil{
		t.Error(err)
	}else{
		t.Log(oo)
	}
}
