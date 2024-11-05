package api

import (
	"context"
	"fmt"
	"net/http"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	TeamId int    `json:"teamId"`
}

type CreateUserRequest struct {
	Name   string `json:"name"`
	TeamId int    `json:"teamId"`
}

type CreateUserResponse struct {
	Result User `json:"result"`
}

type GetUserResponse struct {
	Result User `json:"result"`
}

type ListUserResponse struct {
	Result []User `json:"result"`
}

func (c *Client) CreateUser(ctx context.Context, userName string, teamId int) (*CreateUserResponse, error) {
	requestPath := "/user"
	req := CreateUserRequest{
		Name:   userName,
		TeamId: teamId,
	}
	var res CreateUserResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListUsers(ctx context.Context) (*ListUserResponse, error) {
	requestPath := "/user"
	var res ListUserResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetUser(ctx context.Context, id int) (*GetUserResponse, error) {
	requestPath := fmt.Sprintf("/user/%s", string(id))

	var res GetUserResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateUser(ctx context.Context, id int, name string, teamId int) error {
	requestPath := "/user"

	req := User{
		Id:     id,
		Name:   name,
		TeamId: teamId,
	}

	_, err := c.do(ctx, http.MethodPatch, requestPath, req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUser(ctx context.Context, id int) error {
	requestPath := fmt.Sprintf("/user/%s", string(id))

	_, err := c.do(ctx, http.MethodDelete, requestPath, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
