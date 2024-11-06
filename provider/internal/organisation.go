package api

import (
	"context"
	"fmt"
	"net/http"
)

type Organisation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateOrganisationRequest struct {
	Name string `json:"name"`
}

type CreateOrganisationResponse struct {
	Result Organisation `json:"result"`
}

type GetOrganisationResponse struct {
	Result Organisation `json:"result"`
}

type ListOrganisationResponse struct {
	Result []Organisation `json:"result"`
}

func (c *Client) CreateOrganisation(ctx context.Context, organisationName string) (*CreateOrganisationResponse, error) {
	requestPath := "/organisation"
	req := CreateOrganisationRequest{
		Name: organisationName,
	}
	var res CreateOrganisationResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListOrganisations(ctx context.Context) (*ListOrganisationResponse, error) {
	requestPath := "/organisation"
	var res ListOrganisationResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetOrganisation(ctx context.Context, id string) (*GetOrganisationResponse, *http.Response, error) {
	requestPath := fmt.Sprintf("/organisation/%s", id)

	var orgRes GetOrganisationResponse
	res, err := c.do(ctx, http.MethodGet, requestPath, nil, &orgRes)
	if err != nil {
		return nil, res, err
	}

	return &orgRes, nil, nil
}

func (c *Client) UpdateOrganisation(ctx context.Context, id int, name string) error {
	requestPath := "/organisation"

	req := Organisation{
		Id:   id,
		Name: name,
	}

	_, err := c.do(ctx, http.MethodPut, requestPath, req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteOrganisation(ctx context.Context, id int) error {
	requestPath := fmt.Sprintf("/organisation/%d", id)

	_, err := c.do(ctx, http.MethodDelete, requestPath, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
