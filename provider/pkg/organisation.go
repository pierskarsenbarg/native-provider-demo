package pkg

import (
	"context"
	"fmt"
	"strconv"

	api "github.com/pierskarsenbarg/native-provider-demo/provider/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/property"
)

type Organisation struct{}

type OrganisationArgs struct {
	Name *string `pulumi:"orgName,optional"`
}

func (o *OrganisationArgs) Annotate(a infer.Annotator) {
	a.Describe(&o.Name, "Name of the organisation")
}

type OrganisationState struct {
	Id   int    `pulumi:"orgId"`
	Name string `pulumi:"orgName"`
}

func (o *OrganisationState) Annotate(a infer.Annotator) {
	a.Describe(&o.Id, "Id of the organisation created")
	a.Describe(&o.Name, "Name of organisation created")
}

func (o *Organisation) Create(ctx context.Context, req infer.CreateRequest[OrganisationArgs]) (infer.CreateResponse[OrganisationState], error) {
	if req.DryRun {
		return infer.CreateResponse[OrganisationState]{}, nil
	}

	config := infer.GetConfig[Config](ctx)
	org, err := o.createOrganisation(*req.Inputs.Name, config, ctx)
	if err != nil {
		return infer.CreateResponse[OrganisationState]{}, fmt.Errorf("error creating organisation: %v", err)
	}

	return infer.CreateResponse[OrganisationState]{
		ID: strconv.Itoa(org.Result.Id),
		Output: OrganisationState{
			Id:   org.Result.Id,
			Name: org.Result.Name,
		},
	}, nil
}

func (*Organisation) createOrganisation(name string, config Config, ctx context.Context) (*api.CreateOrganisationResponse, error) {
	organisation, err := config.Client.CreateOrganisation(ctx, name)
	if err != nil {
		return nil, err
	}

	return organisation, nil
}

func (o *Organisation) Delete(ctx context.Context, req infer.DeleteRequest[OrganisationState]) (infer.DeleteResponse, error) {
	config := infer.GetConfig[Config](ctx)

	orgId, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("could not delete organisation. Id not a valid number: %v", err)
	}
	err = o.deleteOrganisation(ctx, orgId, config)
	if err != nil {
		return infer.DeleteResponse{}, fmt.Errorf("error deleting organisation: %v", err)
	}
	return infer.DeleteResponse{}, nil
}

func (*Organisation) deleteOrganisation(ctx context.Context, id int, config Config) error {
	err := config.Client.DeleteOrganisation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (o *Organisation) Diff(ctx context.Context, req infer.DiffRequest[OrganisationArgs, OrganisationState]) (infer.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if req.Inputs.Name != &req.State.Name {
		diff["orgName"] = p.PropertyDiff{Kind: p.Update} // can also be UpdateReplace, DeleteReplace
	}

	return infer.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (o *Organisation) Read(ctx context.Context, req infer.ReadRequest[OrganisationArgs, OrganisationState]) (infer.ReadResponse[OrganisationArgs, OrganisationState], error) {
	config := infer.GetConfig[Config](ctx)

	organisation, err := o.getOrganisation(ctx, req.ID, config)
	if err != nil {
		return infer.ReadResponse[OrganisationArgs, OrganisationState]{}, err
	}

	if organisation == nil {
		return infer.ReadResponse[OrganisationArgs, OrganisationState]{}, nil
	}

	return infer.ReadResponse[OrganisationArgs, OrganisationState]{
		ID: req.ID,
		Inputs: OrganisationArgs{
			Name: req.Inputs.Name,
		},
		State: OrganisationState{
			Id:   req.State.Id,
			Name: organisation.Name,
		},
	}, nil
}

func (*Organisation) getOrganisation(ctx context.Context, id string, config Config) (*api.Organisation, error) {
	organisation, res, err := config.Client.GetOrganisation(ctx, id)
	if err != nil {
		if res.StatusCode == 404 {
			// database doesn't exist
			return nil, nil
		}
		return nil, err
	}
	return &organisation.Result, nil
}

func (o *Organisation) Update(ctx context.Context, req infer.UpdateRequest[OrganisationArgs, OrganisationState]) (infer.UpdateResponse[OrganisationState], error) {
	orgId, err := strconv.Atoi(req.ID)
	if err != nil {
		return infer.UpdateResponse[OrganisationState]{}, err
	}
	if !req.DryRun && req.State.Name != *req.Inputs.Name {
		config := infer.GetConfig[Config](ctx)
		err := o.updateOrganisation(ctx, orgId, *req.Inputs.Name, config)
		if err != nil {
			return infer.UpdateResponse[OrganisationState]{}, err
		}
	}

	return infer.UpdateResponse[OrganisationState]{
		Output: OrganisationState{
			Id:   orgId,
			Name: *req.Inputs.Name,
		},
	}, nil
}

func (*Organisation) updateOrganisation(ctx context.Context, id int, newName string, config Config) error {
	err := config.Client.UpdateOrganisation(ctx, id, newName)
	if err != nil {
		return err
	}
	return nil
}

func (*Organisation) Check(ctx context.Context, req infer.CheckRequest) (infer.CheckResponse[OrganisationArgs], error) {
	// Apply default arguments
	args, failures, err := infer.DefaultCheck[OrganisationArgs](ctx, req.NewInputs)

	// Apply autonaming
	//
	// If args.Name is unset, we set it to a value based off of the resource name.
	args.Name, err = autoname(args.Name, req.Name, "orgName", req.OldInputs)
	return infer.CheckResponse[OrganisationArgs]{
		Inputs:   args,
		Failures: failures,
	}, err
}

type GetOrganisation struct{}

type GetOrganisationArgs struct {
	Id int `pulumi:"orgId"`
}

func (*GetOrganisation) Invoke(ctx context.Context, req infer.FunctionRequest[GetOrganisationArgs]) (infer.FunctionResponse[OrganisationState], error) {
	config := infer.GetConfig[Config](ctx)

	orgId := strconv.Itoa(req.Input.Id)

	organisation, res, err := config.Client.GetOrganisation(ctx, orgId)
	if err != nil {
		if res.StatusCode == 404 {
			return infer.FunctionResponse[OrganisationState]{}, nil
		}
		return infer.FunctionResponse[OrganisationState]{}, err
	}
	return infer.FunctionResponse[OrganisationState]{
		Output: OrganisationState{
			Id:   organisation.Result.Id,
			Name: organisation.Result.Name,
		},
	}, nil
}

func autoname(
	field *string, name, fieldName string,
	oldInputs property.Map,
) (*string, error) {
	if field != nil {
		return field, nil
	}

	prev := oldInputs.Get(fieldName)
	if prev.IsString() && prev.AsString() != "" {
		n := prev.AsString()
		field = &n
	} else {
		n, err := resource.NewUniqueHex(name+"-", 6, 20)
		if err != nil {
			return nil, err
		}
		field = &n
	}

	return field, nil
}
