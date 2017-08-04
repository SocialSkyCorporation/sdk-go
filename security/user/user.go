package user

import (
	"encoding/json"
	"errors"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
)

type SecurityUser struct {
	Kuzzle kuzzle.Kuzzle
}

/*
  Retrieves an User using its provided unique id.
*/
func (su SecurityUser) Fetch(id string, options types.QueryOptions) (types.User, error) {
	if id == "" {
		return types.User{}, errors.New("Security.User.Fetch: user id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getUser",
		Id:         id,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Create a new User in Kuzzle.
*/
func (su SecurityUser) Create(kuid string, content types.UserData, options types.QueryOptions) (types.User, error) {
	if kuid == "" {
		return types.User{}, errors.New("Security.User.Create: user kuid required")
	}

	ch := make(chan types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}
	type createBody struct {
		Content     userData              `json:"content"`
		Credentials types.UserCredentials `json:"credentials"`
	}

	body := createBody{Content: ud, Credentials: content.Credentials}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "createUser",
		Body:       body,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Create a new restricted User in Kuzzle.
*/
func (su SecurityUser) CreateRestrictedUser(kuid string, content types.UserData, options types.QueryOptions) (types.User, error) {
	if kuid == "" {
		return types.User{}, errors.New("Security.User.CreateRestrictedUser: user kuid required")
	}

	ch := make(chan types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	for key, value := range content.Content {
		ud[key] = value
	}
	type createBody struct {
		Content     userData              `json:"content"`
		Credentials types.UserCredentials `json:"credentials"`
	}

	body := createBody{Content: ud, Credentials: content.Credentials}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "createRestrictedUser",
		Body:       body,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Replace an User in Kuzzle.
*/
func (su SecurityUser) Replace(kuid string, content types.UserData, options types.QueryOptions) (types.User, error) {
	if kuid == "" {
		return types.User{}, errors.New("Security.User.Replace: user kuid required")
	}

	ch := make(chan types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "replaceUser",
		Body:       ud,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Update an User in Kuzzle.
*/
func (su SecurityUser) Update(kuid string, content types.UserData, options types.QueryOptions) (types.User, error) {
	if kuid == "" {
		return types.User{}, errors.New("Security.User.Update: user kuid required")
	}

	ch := make(chan types.KuzzleResponse)

	type userData map[string]interface{}
	ud := userData{}
	ud["profileIds"] = content.ProfileIds
	for key, value := range content.Content {
		ud[key] = value
	}

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "updateUser",
		Body:       ud,
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.User{}, errors.New(res.Error.Message)
	}

	user := types.User{}
	json.Unmarshal(res.Result, &user)

	return user, nil
}

/*
  Delete an User in Kuzzle.

  There is a small delay between user deletion and their deletion in our advanced search layer, usually a couple of seconds.
  This means that a user that has just been deleted will still be returned by this function.
*/
func (su SecurityUser) Delete(kuid string, options types.QueryOptions) (string, error) {
	if kuid == "" {
		return "", errors.New("Security.User.Delete: user kuid required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "deleteUser",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return "", errors.New(res.Error.Message)
	}

	shardResponse := types.ShardResponse{}
	json.Unmarshal(res.Result, &shardResponse)

	return shardResponse.Id, nil
}

/*
  Executes a search on Users according to filters.
*/
func (su SecurityUser) Search(filters interface{}, options types.QueryOptions) (types.KuzzleSearchUsersResult, error) {
	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "searchUsers",
		Body:       filters,
	}

	if options != nil {
		query.From = options.GetFrom()
		query.Size = options.GetSize()

		scroll := options.GetScroll()
		if scroll != "" {
			query.Scroll = scroll
		}
	}

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchUsersResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchUsersResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Executes a scroll search on Users.
*/
func (su SecurityUser) Scroll(scrollId string, options types.QueryOptions) (types.KuzzleSearchUsersResult, error) {
	if scrollId == "" {
		return types.KuzzleSearchUsersResult{}, errors.New("Security.User.Scroll: scroll id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "scrollUsers",
		ScrollId:   scrollId,
	}

	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return types.KuzzleSearchUsersResult{}, errors.New(res.Error.Message)
	}

	searchResult := types.KuzzleSearchUsersResult{}
	json.Unmarshal(res.Result, &searchResult)

	return searchResult, nil
}

/*
  Gets the rights of an User using its provided unique id.
*/
func (su SecurityUser) GetRights(kuid string, options types.QueryOptions) ([]types.UserRights, error) {
	if kuid == "" {
		return []types.UserRights{}, errors.New("Security.User.GetRights: user id required")
	}

	ch := make(chan types.KuzzleResponse)

	query := types.KuzzleRequest{
		Controller: "security",
		Action:     "getUserRights",
		Id:         kuid,
	}
	go su.Kuzzle.Query(query, options, ch)

	res := <-ch

	if res.Error.Message != "" {
		return []types.UserRights{}, errors.New(res.Error.Message)
	}

	type response struct {
		UserRights []types.UserRights `json:"hits"`
	}
	userRights := response{}
	json.Unmarshal(res.Result, &userRights)

	return userRights.UserRights, nil
}

/*
  Indicates whether an action is allowed, denied or conditional based on user rights provided as the first argument.
  An action is defined as a couple of action and controller (mandatory), plus an index and a collection(optional).
*/
func (su SecurityUser) IsActionAllowed(rights []types.UserRights, controller string, action string, index string, collection string) (string, error) {
	if rights == nil {
		return "", errors.New("Security.User.IsActionAllowed: Rights parameter is mandatory")
	}
	if controller == "" {
		return "", errors.New("Security.User.IsActionAllowed: Controller parameter is mandatory")
	}
	if action == "" {
		return "", errors.New("Security.User.IsActionAllowed: Action parameter is mandatory")
	}

	filteredUserRights := []types.UserRights{}

	for _, ur := range rights {
		if (ur.Controller == controller || ur.Controller == "*") && (ur.Action == action || ur.Action == "*") && (ur.Index == index || ur.Index == "*") && (ur.Collection == collection || ur.Collection == "*") {
			filteredUserRights = append(filteredUserRights, ur)
		}
	}

	for _, ur := range filteredUserRights {
		if ur.Value == "allowed" {
			return "allowed", nil
		}
		if ur.Value == "conditional" {
			return "conditional", nil
		}
	}

	return "denied", nil
}