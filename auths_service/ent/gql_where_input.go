// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"golang-boilerplate/ent/auth"
	"golang-boilerplate/ent/predicate"
	"time"
)

// AuthWhereInput represents a where input for filtering Auth queries.
type AuthWhereInput struct {
	Predicates []predicate.Auth  `json:"-"`
	Not        *AuthWhereInput   `json:"not,omitempty"`
	Or         []*AuthWhereInput `json:"or,omitempty"`
	And        []*AuthWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "username" field predicates.
	Username             *string  `json:"username,omitempty"`
	UsernameNEQ          *string  `json:"usernameNEQ,omitempty"`
	UsernameIn           []string `json:"usernameIn,omitempty"`
	UsernameNotIn        []string `json:"usernameNotIn,omitempty"`
	UsernameGT           *string  `json:"usernameGT,omitempty"`
	UsernameGTE          *string  `json:"usernameGTE,omitempty"`
	UsernameLT           *string  `json:"usernameLT,omitempty"`
	UsernameLTE          *string  `json:"usernameLTE,omitempty"`
	UsernameContains     *string  `json:"usernameContains,omitempty"`
	UsernameHasPrefix    *string  `json:"usernameHasPrefix,omitempty"`
	UsernameHasSuffix    *string  `json:"usernameHasSuffix,omitempty"`
	UsernameEqualFold    *string  `json:"usernameEqualFold,omitempty"`
	UsernameContainsFold *string  `json:"usernameContainsFold,omitempty"`

	// "password" field predicates.
	Password             *string  `json:"password,omitempty"`
	PasswordNEQ          *string  `json:"passwordNEQ,omitempty"`
	PasswordIn           []string `json:"passwordIn,omitempty"`
	PasswordNotIn        []string `json:"passwordNotIn,omitempty"`
	PasswordGT           *string  `json:"passwordGT,omitempty"`
	PasswordGTE          *string  `json:"passwordGTE,omitempty"`
	PasswordLT           *string  `json:"passwordLT,omitempty"`
	PasswordLTE          *string  `json:"passwordLTE,omitempty"`
	PasswordContains     *string  `json:"passwordContains,omitempty"`
	PasswordHasPrefix    *string  `json:"passwordHasPrefix,omitempty"`
	PasswordHasSuffix    *string  `json:"passwordHasSuffix,omitempty"`
	PasswordEqualFold    *string  `json:"passwordEqualFold,omitempty"`
	PasswordContainsFold *string  `json:"passwordContainsFold,omitempty"`

	// "access_token" field predicates.
	AccessToken             *string  `json:"accessToken,omitempty"`
	AccessTokenNEQ          *string  `json:"accessTokenNEQ,omitempty"`
	AccessTokenIn           []string `json:"accessTokenIn,omitempty"`
	AccessTokenNotIn        []string `json:"accessTokenNotIn,omitempty"`
	AccessTokenGT           *string  `json:"accessTokenGT,omitempty"`
	AccessTokenGTE          *string  `json:"accessTokenGTE,omitempty"`
	AccessTokenLT           *string  `json:"accessTokenLT,omitempty"`
	AccessTokenLTE          *string  `json:"accessTokenLTE,omitempty"`
	AccessTokenContains     *string  `json:"accessTokenContains,omitempty"`
	AccessTokenHasPrefix    *string  `json:"accessTokenHasPrefix,omitempty"`
	AccessTokenHasSuffix    *string  `json:"accessTokenHasSuffix,omitempty"`
	AccessTokenIsNil        bool     `json:"accessTokenIsNil,omitempty"`
	AccessTokenNotNil       bool     `json:"accessTokenNotNil,omitempty"`
	AccessTokenEqualFold    *string  `json:"accessTokenEqualFold,omitempty"`
	AccessTokenContainsFold *string  `json:"accessTokenContainsFold,omitempty"`

	// "refresh_token" field predicates.
	RefreshToken             *string  `json:"refreshToken,omitempty"`
	RefreshTokenNEQ          *string  `json:"refreshTokenNEQ,omitempty"`
	RefreshTokenIn           []string `json:"refreshTokenIn,omitempty"`
	RefreshTokenNotIn        []string `json:"refreshTokenNotIn,omitempty"`
	RefreshTokenGT           *string  `json:"refreshTokenGT,omitempty"`
	RefreshTokenGTE          *string  `json:"refreshTokenGTE,omitempty"`
	RefreshTokenLT           *string  `json:"refreshTokenLT,omitempty"`
	RefreshTokenLTE          *string  `json:"refreshTokenLTE,omitempty"`
	RefreshTokenContains     *string  `json:"refreshTokenContains,omitempty"`
	RefreshTokenHasPrefix    *string  `json:"refreshTokenHasPrefix,omitempty"`
	RefreshTokenHasSuffix    *string  `json:"refreshTokenHasSuffix,omitempty"`
	RefreshTokenIsNil        bool     `json:"refreshTokenIsNil,omitempty"`
	RefreshTokenNotNil       bool     `json:"refreshTokenNotNil,omitempty"`
	RefreshTokenEqualFold    *string  `json:"refreshTokenEqualFold,omitempty"`
	RefreshTokenContainsFold *string  `json:"refreshTokenContainsFold,omitempty"`

	// "user_id" field predicates.
	UserID      *int  `json:"userID,omitempty"`
	UserIDNEQ   *int  `json:"userIDNEQ,omitempty"`
	UserIDIn    []int `json:"userIDIn,omitempty"`
	UserIDNotIn []int `json:"userIDNotIn,omitempty"`
	UserIDGT    *int  `json:"userIDGT,omitempty"`
	UserIDGTE   *int  `json:"userIDGTE,omitempty"`
	UserIDLT    *int  `json:"userIDLT,omitempty"`
	UserIDLTE   *int  `json:"userIDLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *AuthWhereInput) AddPredicates(predicates ...predicate.Auth) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the AuthWhereInput filter on the AuthQuery builder.
func (i *AuthWhereInput) Filter(q *AuthQuery) (*AuthQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyAuthWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyAuthWhereInput is returned in case the AuthWhereInput is empty.
var ErrEmptyAuthWhereInput = errors.New("ent: empty predicate AuthWhereInput")

// P returns a predicate for filtering auths.
// An error is returned if the input is empty or invalid.
func (i *AuthWhereInput) P() (predicate.Auth, error) {
	var predicates []predicate.Auth
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, auth.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Auth, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, auth.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Auth, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, auth.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, auth.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, auth.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, auth.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, auth.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, auth.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, auth.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, auth.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, auth.IDLTE(*i.IDLTE))
	}
	if i.Username != nil {
		predicates = append(predicates, auth.UsernameEQ(*i.Username))
	}
	if i.UsernameNEQ != nil {
		predicates = append(predicates, auth.UsernameNEQ(*i.UsernameNEQ))
	}
	if len(i.UsernameIn) > 0 {
		predicates = append(predicates, auth.UsernameIn(i.UsernameIn...))
	}
	if len(i.UsernameNotIn) > 0 {
		predicates = append(predicates, auth.UsernameNotIn(i.UsernameNotIn...))
	}
	if i.UsernameGT != nil {
		predicates = append(predicates, auth.UsernameGT(*i.UsernameGT))
	}
	if i.UsernameGTE != nil {
		predicates = append(predicates, auth.UsernameGTE(*i.UsernameGTE))
	}
	if i.UsernameLT != nil {
		predicates = append(predicates, auth.UsernameLT(*i.UsernameLT))
	}
	if i.UsernameLTE != nil {
		predicates = append(predicates, auth.UsernameLTE(*i.UsernameLTE))
	}
	if i.UsernameContains != nil {
		predicates = append(predicates, auth.UsernameContains(*i.UsernameContains))
	}
	if i.UsernameHasPrefix != nil {
		predicates = append(predicates, auth.UsernameHasPrefix(*i.UsernameHasPrefix))
	}
	if i.UsernameHasSuffix != nil {
		predicates = append(predicates, auth.UsernameHasSuffix(*i.UsernameHasSuffix))
	}
	if i.UsernameEqualFold != nil {
		predicates = append(predicates, auth.UsernameEqualFold(*i.UsernameEqualFold))
	}
	if i.UsernameContainsFold != nil {
		predicates = append(predicates, auth.UsernameContainsFold(*i.UsernameContainsFold))
	}
	if i.Password != nil {
		predicates = append(predicates, auth.PasswordEQ(*i.Password))
	}
	if i.PasswordNEQ != nil {
		predicates = append(predicates, auth.PasswordNEQ(*i.PasswordNEQ))
	}
	if len(i.PasswordIn) > 0 {
		predicates = append(predicates, auth.PasswordIn(i.PasswordIn...))
	}
	if len(i.PasswordNotIn) > 0 {
		predicates = append(predicates, auth.PasswordNotIn(i.PasswordNotIn...))
	}
	if i.PasswordGT != nil {
		predicates = append(predicates, auth.PasswordGT(*i.PasswordGT))
	}
	if i.PasswordGTE != nil {
		predicates = append(predicates, auth.PasswordGTE(*i.PasswordGTE))
	}
	if i.PasswordLT != nil {
		predicates = append(predicates, auth.PasswordLT(*i.PasswordLT))
	}
	if i.PasswordLTE != nil {
		predicates = append(predicates, auth.PasswordLTE(*i.PasswordLTE))
	}
	if i.PasswordContains != nil {
		predicates = append(predicates, auth.PasswordContains(*i.PasswordContains))
	}
	if i.PasswordHasPrefix != nil {
		predicates = append(predicates, auth.PasswordHasPrefix(*i.PasswordHasPrefix))
	}
	if i.PasswordHasSuffix != nil {
		predicates = append(predicates, auth.PasswordHasSuffix(*i.PasswordHasSuffix))
	}
	if i.PasswordEqualFold != nil {
		predicates = append(predicates, auth.PasswordEqualFold(*i.PasswordEqualFold))
	}
	if i.PasswordContainsFold != nil {
		predicates = append(predicates, auth.PasswordContainsFold(*i.PasswordContainsFold))
	}
	if i.AccessToken != nil {
		predicates = append(predicates, auth.AccessTokenEQ(*i.AccessToken))
	}
	if i.AccessTokenNEQ != nil {
		predicates = append(predicates, auth.AccessTokenNEQ(*i.AccessTokenNEQ))
	}
	if len(i.AccessTokenIn) > 0 {
		predicates = append(predicates, auth.AccessTokenIn(i.AccessTokenIn...))
	}
	if len(i.AccessTokenNotIn) > 0 {
		predicates = append(predicates, auth.AccessTokenNotIn(i.AccessTokenNotIn...))
	}
	if i.AccessTokenGT != nil {
		predicates = append(predicates, auth.AccessTokenGT(*i.AccessTokenGT))
	}
	if i.AccessTokenGTE != nil {
		predicates = append(predicates, auth.AccessTokenGTE(*i.AccessTokenGTE))
	}
	if i.AccessTokenLT != nil {
		predicates = append(predicates, auth.AccessTokenLT(*i.AccessTokenLT))
	}
	if i.AccessTokenLTE != nil {
		predicates = append(predicates, auth.AccessTokenLTE(*i.AccessTokenLTE))
	}
	if i.AccessTokenContains != nil {
		predicates = append(predicates, auth.AccessTokenContains(*i.AccessTokenContains))
	}
	if i.AccessTokenHasPrefix != nil {
		predicates = append(predicates, auth.AccessTokenHasPrefix(*i.AccessTokenHasPrefix))
	}
	if i.AccessTokenHasSuffix != nil {
		predicates = append(predicates, auth.AccessTokenHasSuffix(*i.AccessTokenHasSuffix))
	}
	if i.AccessTokenIsNil {
		predicates = append(predicates, auth.AccessTokenIsNil())
	}
	if i.AccessTokenNotNil {
		predicates = append(predicates, auth.AccessTokenNotNil())
	}
	if i.AccessTokenEqualFold != nil {
		predicates = append(predicates, auth.AccessTokenEqualFold(*i.AccessTokenEqualFold))
	}
	if i.AccessTokenContainsFold != nil {
		predicates = append(predicates, auth.AccessTokenContainsFold(*i.AccessTokenContainsFold))
	}
	if i.RefreshToken != nil {
		predicates = append(predicates, auth.RefreshTokenEQ(*i.RefreshToken))
	}
	if i.RefreshTokenNEQ != nil {
		predicates = append(predicates, auth.RefreshTokenNEQ(*i.RefreshTokenNEQ))
	}
	if len(i.RefreshTokenIn) > 0 {
		predicates = append(predicates, auth.RefreshTokenIn(i.RefreshTokenIn...))
	}
	if len(i.RefreshTokenNotIn) > 0 {
		predicates = append(predicates, auth.RefreshTokenNotIn(i.RefreshTokenNotIn...))
	}
	if i.RefreshTokenGT != nil {
		predicates = append(predicates, auth.RefreshTokenGT(*i.RefreshTokenGT))
	}
	if i.RefreshTokenGTE != nil {
		predicates = append(predicates, auth.RefreshTokenGTE(*i.RefreshTokenGTE))
	}
	if i.RefreshTokenLT != nil {
		predicates = append(predicates, auth.RefreshTokenLT(*i.RefreshTokenLT))
	}
	if i.RefreshTokenLTE != nil {
		predicates = append(predicates, auth.RefreshTokenLTE(*i.RefreshTokenLTE))
	}
	if i.RefreshTokenContains != nil {
		predicates = append(predicates, auth.RefreshTokenContains(*i.RefreshTokenContains))
	}
	if i.RefreshTokenHasPrefix != nil {
		predicates = append(predicates, auth.RefreshTokenHasPrefix(*i.RefreshTokenHasPrefix))
	}
	if i.RefreshTokenHasSuffix != nil {
		predicates = append(predicates, auth.RefreshTokenHasSuffix(*i.RefreshTokenHasSuffix))
	}
	if i.RefreshTokenIsNil {
		predicates = append(predicates, auth.RefreshTokenIsNil())
	}
	if i.RefreshTokenNotNil {
		predicates = append(predicates, auth.RefreshTokenNotNil())
	}
	if i.RefreshTokenEqualFold != nil {
		predicates = append(predicates, auth.RefreshTokenEqualFold(*i.RefreshTokenEqualFold))
	}
	if i.RefreshTokenContainsFold != nil {
		predicates = append(predicates, auth.RefreshTokenContainsFold(*i.RefreshTokenContainsFold))
	}
	if i.UserID != nil {
		predicates = append(predicates, auth.UserIDEQ(*i.UserID))
	}
	if i.UserIDNEQ != nil {
		predicates = append(predicates, auth.UserIDNEQ(*i.UserIDNEQ))
	}
	if len(i.UserIDIn) > 0 {
		predicates = append(predicates, auth.UserIDIn(i.UserIDIn...))
	}
	if len(i.UserIDNotIn) > 0 {
		predicates = append(predicates, auth.UserIDNotIn(i.UserIDNotIn...))
	}
	if i.UserIDGT != nil {
		predicates = append(predicates, auth.UserIDGT(*i.UserIDGT))
	}
	if i.UserIDGTE != nil {
		predicates = append(predicates, auth.UserIDGTE(*i.UserIDGTE))
	}
	if i.UserIDLT != nil {
		predicates = append(predicates, auth.UserIDLT(*i.UserIDLT))
	}
	if i.UserIDLTE != nil {
		predicates = append(predicates, auth.UserIDLTE(*i.UserIDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, auth.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, auth.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, auth.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, auth.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, auth.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, auth.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, auth.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, auth.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, auth.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, auth.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, auth.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, auth.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, auth.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, auth.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, auth.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, auth.UpdatedAtLTE(*i.UpdatedAtLTE))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyAuthWhereInput
	case 1:
		return predicates[0], nil
	default:
		return auth.And(predicates...), nil
	}
}
