package api

import (
	"context"
	"fmt"
	"net/http"
)

type Team struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	OrganisationId int    `json:"organisationId"`
}

type CreateTeamRequest struct {
	Name           string `json:"name"`
	OrganisationId int    `json:"organisationId"`
}

type CreateTeamResponse struct {
	Result Team `json:"result"`
}

type GetTeamResponse struct {
	Result Team `json:"result"`
}

type ListTeamResponse struct {
	Result []Team `json:"result"`
}

func (c *Client) CreateTeam(ctx context.Context, teamName string, organisationId int) (*CreateTeamResponse, error) {
	requestPath := "/team"
	req := CreateTeamRequest{
		Name:           teamName,
		OrganisationId: organisationId,
	}
	var res CreateTeamResponse
	_, err := c.do(ctx, http.MethodPost, requestPath, req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) ListTeams(ctx context.Context) (*ListTeamResponse, error) {
	requestPath := "/team"
	var res ListTeamResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeam(ctx context.Context, id int) (*GetTeamResponse, error) {
	requestPath := fmt.Sprintf("/team/%s", string(id))

	var res GetTeamResponse
	_, err := c.do(ctx, http.MethodGet, requestPath, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateTeam(ctx context.Context, id int, name string, organisationId int) error {
	requestPath := "/team"

	req := Team{
		Id:             id,
		Name:           name,
		OrganisationId: organisationId,
	}

	_, err := c.do(ctx, http.MethodPatch, requestPath, req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTeam(ctx context.Context, id int) error {
	requestPath := fmt.Sprintf("/team/%s", string(id))

	_, err := c.do(ctx, http.MethodDelete, requestPath, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
