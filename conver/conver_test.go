package conver

import "testing"

//TODO HANK
type pb_PutResponse struct{
	Key string
}

type PutResponse pb_PutResponse

func TestConverStruct(t *testing.T){
	pbPutResponse := &pb_PutResponse{Key:"foo"}

	//注意分先后顺序
	putResponse := (*PutResponse)(pbPutResponse)

	t.Logf("%v",putResponse)
}

//interface{} gres.(*pb.CreateGroupRes)




