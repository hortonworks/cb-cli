package user
import (
	"strconv"
	"strings"
	"testing"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/client/users"
	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
	"github.com/hortonworks/cb-cli/dataplane/types"
	"github.com/hortonworks/dp-cli-common/utils"
	
)

type mockUserClient struct {
}


func (*mockUserClient) GetAllUsers (params *users.GetAllUsersParams) (*users.GetAllUsersOK, error) {
	resp := []*model.UserWithRoles{
		{
			Name: &(&types.S{S: "test1"}).S,
			PreferredUsername: &(&types.S{S: "test1"}).S,
			Email: "test1",
		},
		{
			Name: &(&types.S{S: "test2"}).S,
			PreferredUsername: &(&types.S{S: "test2"}).S,
			Email: "test2",
		},
	}
	return &users.GetAllUsersOK{Payload: resp}, nil
}

func TestListUsersImpl(t *testing.T) {
	var rows []utils.Row
	listUsersImpl(new(mockUserClient), func(h []string, r []utils.Row) { rows = r })
	if len(rows) != 2 {
		t.Fatalf("row number doesn't match 2 == %d", len(rows))
	}
	for i, r := range rows {
		expected := []string{" " + "test" + strconv.Itoa(i+1), "test" + strconv.Itoa(i+1), "test" + strconv.Itoa(i+1)}
		if strings.Join(r.DataAsStringArray(), " ") != strings.Join(expected, " ") {
			t.Errorf("row data not match %s == %s", strings.Join(expected, " "), strings.Join(r.DataAsStringArray(), " "))
		}
	}
}