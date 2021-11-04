package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhaoyunxing92/dingtalk/v2/request"
	"testing"
)

var client = NewDingTalk().
	SetId(1244553273).
	SetKey("dingkjy4w80esdwgjuyo").
	SetSecret("bDKa_nfJg3zYRsFrj-wTohTuoJCtxTEHaGmybYF9vgaVAZJOz-mICsLGStB288nW").
	Build()

func TestDingTalk_CreateUser(t *testing.T) {

	user, err := client.CreateUser(request.NewCreateUser("张三", "15669019211", 554656655).Build())
	assert.Nil(t, err)
	assert.NotNil(t, user.UserId)
	t.Log(user.UserId)
}

func TestDingTalk_UpdateUser(t *testing.T) {

	res, err := client.UpdateUser(
		request.NewUpdateUser("1948546245774889").
			SetName("李四").
			SetDept(1, 554656655).
			SetTelephone("").
			SetForceUpdateFields("title,userid,title,userid").
			SetDeptTitle(1, "普通员工").
			SetDeptTitle(554656655, "设计人员").
			SetJobNumber("").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, res.Code)
}

func TestDingTalk_DeleteUser(t *testing.T) {

	res, err := client.DeleteUser("1948546245774889")

	assert.Nil(t, err)
	assert.NotNil(t, res.Code)
	assert.NotNil(t, res.RequestId)
}

func TestDingTalk_GetUserDetail(t *testing.T) {

	detail, err := client.GetUserDetail(request.NewUserDetail("1948546245774889").Build())

	assert.Nil(t, err)
	assert.NotNil(t, detail.Code)
	assert.NotNil(t, detail.RequestId)
	assert.NotNil(t, detail.Name)
	assert.NotNil(t, detail.UnionId)
	assert.NotNil(t, detail.UserId)
}

func TestDingTalk_GetUserCount(t *testing.T) {

	req, err := client.GetUserCount(request.NewUserCount(true))

	assert.Nil(t, err)
	assert.NotNil(t, req.Count)
}

func TestDingTalk_GetInactiveUser(t *testing.T) {

	req, err := client.GetInactiveUser(
		request.NewInactiveUser(true, 0, 10, "20211103").
			Build())

	assert.Nil(t, err)
	assert.NotNil(t, req)
}

func TestDingTalk_GetUserIdByMobile(t *testing.T) {

	req, err := client.GetUserIdByMobile(request.NewMobileGetUserId("18357154439"))

	assert.Nil(t, err)
	assert.NotNil(t, req.UserId)
}

func TestDingTalk_GetUserIdByUnionId(t *testing.T) {

	req, err := client.GetUserIdByUnionId(request.NewUnionIdGetUserId("DmJWoRm5zkjT9kWWgVpm0giEiE"))

	assert.Nil(t, err)
	assert.NotNil(t, req.UserId)
}
