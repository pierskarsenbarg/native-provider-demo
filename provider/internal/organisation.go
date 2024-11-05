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

func (c *Client) GetOrganisation(ctx context.Context, id int) (*GetOrganisationResponse, error) {
	requestPath := fmt.Sprintf("/organisation/%s", string(id))

	var res GetOrganisationResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// TODO Add update name function

// TODO Add delete org function
