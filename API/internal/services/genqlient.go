// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
)

// LogInWithUsernameLoginSessionIQueryableLoginSession includes the requested fields of the GraphQL interface IQueryableLoginSession.
//
// LogInWithUsernameLoginSessionIQueryableLoginSession is implemented by the following types:
// LogInWithUsernameLoginSessionQueryableLoginSession
type LogInWithUsernameLoginSessionIQueryableLoginSession interface {
	implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSession()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetLogin returns the interface-field "login" from its implementation.
	GetLogin() LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin
	// GetSessionToken returns the interface-field "sessionToken" from its implementation.
	GetSessionToken() LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken
}

func (v *LogInWithUsernameLoginSessionQueryableLoginSession) implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSession() {
}

func __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSession(b []byte, v *LogInWithUsernameLoginSessionIQueryableLoginSession) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "QueryableLoginSession":
		*v = new(LogInWithUsernameLoginSessionQueryableLoginSession)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing IQueryableLoginSession.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSession: "%v"`, tn.TypeName)
	}
}

func __marshalLogInWithUsernameLoginSessionIQueryableLoginSession(v *LogInWithUsernameLoginSessionIQueryableLoginSession) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *LogInWithUsernameLoginSessionQueryableLoginSession:
		typename = "QueryableLoginSession"

		premarshaled, err := v.__premarshalJSON()
		if err != nil {
			return nil, err
		}
		result := struct {
			TypeName string `json:"__typename"`
			*__premarshalLogInWithUsernameLoginSessionQueryableLoginSession
		}{typename, premarshaled}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSession: "%T"`, v)
	}
}

// LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin includes the requested fields of the GraphQL interface IQueryableLogin.
//
// LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin is implemented by the following types:
// LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin
type LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin interface {
	implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetUsername returns the interface-field "username" from its implementation.
	GetUsername() string
}

func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin) implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin() {
}

func __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin(b []byte, v *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "QueryableLogin":
		*v = new(LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing IQueryableLogin.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin: "%v"`, tn.TypeName)
	}
}

func __marshalLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin(v *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin:
		typename = "QueryableLogin"

		result := struct {
			TypeName string `json:"__typename"`
			*LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin: "%T"`, v)
	}
}

// LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin includes the requested fields of the GraphQL type QueryableLogin.
type LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin struct {
	Typename string `json:"__typename"`
	Username string `json:"username"`
}

// GetTypename returns LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin.Typename, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin) GetTypename() string {
	return v.Typename
}

// GetUsername returns LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin.Username, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionLoginQueryableLogin) GetUsername() string {
	return v.Username
}

// LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken includes the requested fields of the GraphQL interface IQueryableSessionToken.
//
// LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken is implemented by the following types:
// LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken
type LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken interface {
	implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetToken returns the interface-field "token" from its implementation.
	GetToken() string
}

func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken) implementsGraphQLInterfaceLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken() {
}

func __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken(b []byte, v *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "QueryableSessionToken":
		*v = new(LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing IQueryableSessionToken.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken: "%v"`, tn.TypeName)
	}
}

func __marshalLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken(v *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken:
		typename = "QueryableSessionToken"

		result := struct {
			TypeName string `json:"__typename"`
			*LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken: "%T"`, v)
	}
}

// LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken includes the requested fields of the GraphQL type QueryableSessionToken.
type LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken struct {
	Typename string `json:"__typename"`
	Token    string `json:"token"`
}

// GetTypename returns LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken.Typename, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken) GetTypename() string {
	return v.Typename
}

// GetToken returns LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken.Token, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenQueryableSessionToken) GetToken() string {
	return v.Token
}

// LogInWithUsernameLoginSessionQueryableLoginSession includes the requested fields of the GraphQL type QueryableLoginSession.
type LogInWithUsernameLoginSessionQueryableLoginSession struct {
	Typename     string                                                                                `json:"__typename"`
	Login        LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin               `json:"-"`
	SessionToken LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken `json:"-"`
}

// GetTypename returns LogInWithUsernameLoginSessionQueryableLoginSession.Typename, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionQueryableLoginSession) GetTypename() string { return v.Typename }

// GetLogin returns LogInWithUsernameLoginSessionQueryableLoginSession.Login, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionQueryableLoginSession) GetLogin() LogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin {
	return v.Login
}

// GetSessionToken returns LogInWithUsernameLoginSessionQueryableLoginSession.SessionToken, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameLoginSessionQueryableLoginSession) GetSessionToken() LogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken {
	return v.SessionToken
}

func (v *LogInWithUsernameLoginSessionQueryableLoginSession) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*LogInWithUsernameLoginSessionQueryableLoginSession
		Login        json.RawMessage `json:"login"`
		SessionToken json.RawMessage `json:"sessionToken"`
		graphql.NoUnmarshalJSON
	}
	firstPass.LogInWithUsernameLoginSessionQueryableLoginSession = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Login
		src := firstPass.Login
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal LogInWithUsernameLoginSessionQueryableLoginSession.Login: %w", err)
			}
		}
	}

	{
		dst := &v.SessionToken
		src := firstPass.SessionToken
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal LogInWithUsernameLoginSessionQueryableLoginSession.SessionToken: %w", err)
			}
		}
	}
	return nil
}

type __premarshalLogInWithUsernameLoginSessionQueryableLoginSession struct {
	Typename string `json:"__typename"`

	Login json.RawMessage `json:"login"`

	SessionToken json.RawMessage `json:"sessionToken"`
}

func (v *LogInWithUsernameLoginSessionQueryableLoginSession) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *LogInWithUsernameLoginSessionQueryableLoginSession) __premarshalJSON() (*__premarshalLogInWithUsernameLoginSessionQueryableLoginSession, error) {
	var retval __premarshalLogInWithUsernameLoginSessionQueryableLoginSession

	retval.Typename = v.Typename
	{

		dst := &retval.Login
		src := v.Login
		var err error
		*dst, err = __marshalLogInWithUsernameLoginSessionIQueryableLoginSessionLoginIQueryableLogin(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal LogInWithUsernameLoginSessionQueryableLoginSession.Login: %w", err)
		}
	}
	{

		dst := &retval.SessionToken
		src := v.SessionToken
		var err error
		*dst, err = __marshalLogInWithUsernameLoginSessionIQueryableLoginSessionSessionTokenIQueryableSessionToken(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal LogInWithUsernameLoginSessionQueryableLoginSession.SessionToken: %w", err)
		}
	}
	return &retval, nil
}

// LogInWithUsernameResponse is returned by LogInWithUsername on success.
type LogInWithUsernameResponse struct {
	LoginSession LogInWithUsernameLoginSessionIQueryableLoginSession `json:"-"`
}

// GetLoginSession returns LogInWithUsernameResponse.LoginSession, and is useful for accessing the field via an interface.
func (v *LogInWithUsernameResponse) GetLoginSession() LogInWithUsernameLoginSessionIQueryableLoginSession {
	return v.LoginSession
}

func (v *LogInWithUsernameResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*LogInWithUsernameResponse
		LoginSession json.RawMessage `json:"loginSession"`
		graphql.NoUnmarshalJSON
	}
	firstPass.LogInWithUsernameResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.LoginSession
		src := firstPass.LoginSession
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalLogInWithUsernameLoginSessionIQueryableLoginSession(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal LogInWithUsernameResponse.LoginSession: %w", err)
			}
		}
	}
	return nil
}

type __premarshalLogInWithUsernameResponse struct {
	LoginSession json.RawMessage `json:"loginSession"`
}

func (v *LogInWithUsernameResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *LogInWithUsernameResponse) __premarshalJSON() (*__premarshalLogInWithUsernameResponse, error) {
	var retval __premarshalLogInWithUsernameResponse

	{

		dst := &retval.LoginSession
		src := v.LoginSession
		var err error
		*dst, err = __marshalLogInWithUsernameLoginSessionIQueryableLoginSession(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal LogInWithUsernameResponse.LoginSession: %w", err)
		}
	}
	return &retval, nil
}

// VerifyTokenLoginIQueryableLogin includes the requested fields of the GraphQL interface IQueryableLogin.
//
// VerifyTokenLoginIQueryableLogin is implemented by the following types:
// VerifyTokenLoginQueryableLogin
type VerifyTokenLoginIQueryableLogin interface {
	implementsGraphQLInterfaceVerifyTokenLoginIQueryableLogin()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	GetId() string
}

func (v *VerifyTokenLoginQueryableLogin) implementsGraphQLInterfaceVerifyTokenLoginIQueryableLogin() {
}

func __unmarshalVerifyTokenLoginIQueryableLogin(b []byte, v *VerifyTokenLoginIQueryableLogin) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "QueryableLogin":
		*v = new(VerifyTokenLoginQueryableLogin)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"response was missing IQueryableLogin.__typename")
	default:
		return fmt.Errorf(
			`unexpected concrete type for VerifyTokenLoginIQueryableLogin: "%v"`, tn.TypeName)
	}
}

func __marshalVerifyTokenLoginIQueryableLogin(v *VerifyTokenLoginIQueryableLogin) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *VerifyTokenLoginQueryableLogin:
		typename = "QueryableLogin"

		result := struct {
			TypeName string `json:"__typename"`
			*VerifyTokenLoginQueryableLogin
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`unexpected concrete type for VerifyTokenLoginIQueryableLogin: "%T"`, v)
	}
}

// VerifyTokenLoginQueryableLogin includes the requested fields of the GraphQL type QueryableLogin.
type VerifyTokenLoginQueryableLogin struct {
	Typename string `json:"__typename"`
	Id       string `json:"id"`
}

// GetTypename returns VerifyTokenLoginQueryableLogin.Typename, and is useful for accessing the field via an interface.
func (v *VerifyTokenLoginQueryableLogin) GetTypename() string { return v.Typename }

// GetId returns VerifyTokenLoginQueryableLogin.Id, and is useful for accessing the field via an interface.
func (v *VerifyTokenLoginQueryableLogin) GetId() string { return v.Id }

// VerifyTokenResponse is returned by VerifyToken on success.
type VerifyTokenResponse struct {
	Login VerifyTokenLoginIQueryableLogin `json:"-"`
}

// GetLogin returns VerifyTokenResponse.Login, and is useful for accessing the field via an interface.
func (v *VerifyTokenResponse) GetLogin() VerifyTokenLoginIQueryableLogin { return v.Login }

func (v *VerifyTokenResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*VerifyTokenResponse
		Login json.RawMessage `json:"login"`
		graphql.NoUnmarshalJSON
	}
	firstPass.VerifyTokenResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.Login
		src := firstPass.Login
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalVerifyTokenLoginIQueryableLogin(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"unable to unmarshal VerifyTokenResponse.Login: %w", err)
			}
		}
	}
	return nil
}

type __premarshalVerifyTokenResponse struct {
	Login json.RawMessage `json:"login"`
}

func (v *VerifyTokenResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *VerifyTokenResponse) __premarshalJSON() (*__premarshalVerifyTokenResponse, error) {
	var retval __premarshalVerifyTokenResponse

	{

		dst := &retval.Login
		src := v.Login
		var err error
		*dst, err = __marshalVerifyTokenLoginIQueryableLogin(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"unable to marshal VerifyTokenResponse.Login: %w", err)
		}
	}
	return &retval, nil
}

// __LogInWithUsernameInput is used internally by genqlient
type __LogInWithUsernameInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUsername returns __LogInWithUsernameInput.Username, and is useful for accessing the field via an interface.
func (v *__LogInWithUsernameInput) GetUsername() string { return v.Username }

// GetPassword returns __LogInWithUsernameInput.Password, and is useful for accessing the field via an interface.
func (v *__LogInWithUsernameInput) GetPassword() string { return v.Password }

// __VerifyTokenInput is used internally by genqlient
type __VerifyTokenInput struct {
	SessionToken string `json:"sessionToken"`
}

// GetSessionToken returns __VerifyTokenInput.SessionToken, and is useful for accessing the field via an interface.
func (v *__VerifyTokenInput) GetSessionToken() string { return v.SessionToken }

// The mutation executed by LogInWithUsername.
const LogInWithUsername_Operation = `
mutation LogInWithUsername ($username: String!, $password: String!) {
	loginSession: logInWithPassword(username: $username, password: $password) {
		__typename
		login {
			__typename
			username
		}
		sessionToken {
			__typename
			token
		}
	}
}
`

func LogInWithUsername(
	ctx_ context.Context,
	client_ graphql.Client,
	username string,
	password string,
) (data_ *LogInWithUsernameResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "LogInWithUsername",
		Query:  LogInWithUsername_Operation,
		Variables: &__LogInWithUsernameInput{
			Username: username,
			Password: password,
		},
	}

	data_ = &LogInWithUsernameResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}

// The query executed by VerifyToken.
const VerifyToken_Operation = `
query VerifyToken ($sessionToken: String!) {
	login(sessionToken: $sessionToken) {
		__typename
		id
	}
}
`

func VerifyToken(
	ctx_ context.Context,
	client_ graphql.Client,
	sessionToken string,
) (data_ *VerifyTokenResponse, err_ error) {
	req_ := &graphql.Request{
		OpName: "VerifyToken",
		Query:  VerifyToken_Operation,
		Variables: &__VerifyTokenInput{
			SessionToken: sessionToken,
		},
	}

	data_ = &VerifyTokenResponse{}
	resp_ := &graphql.Response{Data: data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return data_, err_
}
