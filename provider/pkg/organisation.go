package pkg

import (
	"context"
	"fmt"
	"strconv"

	api "github.com/pierskarsenbarg/native-provider-demo/provider/internal"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
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

func (o *Organisation) Create(ctx context.Context, name string, input OrganisationArgs, preview bool) (id string, output OrganisationState, err error) {
	if preview {
		return "", OrganisationState{}, nil
	}

	config := infer.GetConfig[Config](ctx)
	org, err := o.createOrganisation(*input.Name, config, ctx)
	if err != nil {
		return "", OrganisationState{}, fmt.Errorf("error creating organisation: %v", err)
	}

	return strconv.Itoa(org.Result.Id), OrganisationState{
		Id:   org.Result.Id,
		Name: org.Result.Name,
	}, nil
}

func (*Organisation) createOrganisation(name string, config Config, ctx context.Context) (*api.CreateOrganisationResponse, error) {
	organisation, err := config.Client.CreateOrganisation(ctx, name)
	if err != nil {
		return nil, err
	}

	return organisation, nil
}

func (o *Organisation) Delete(ctx context.Context, id string, props OrganisationState) error {
	config := infer.GetConfig[Config](ctx)

	orgId, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("could not delete organisation. Id not a valid number: %v", err)
	}
	err = o.deleteOrganisation(ctx, orgId, config)
	if err != nil {
		return fmt.Errorf("error deleting organisation: %v", err)
	}
	return nil
}

func (*Organisation) deleteOrganisation(ctx context.Context, id int, config Config) error {
	err := config.Client.DeleteOrganisation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (o *Organisation) Diff(ctx context.Context, id string, olds OrganisationState, news OrganisationArgs) (p.DiffResponse, error) {
	diff := map[string]p.PropertyDiff{}

	if news.Name != &olds.Name {
		diff["orgName"] = p.PropertyDiff{Kind: p.Update} // can also be UpdateReplace, DeleteReplace
	}

	return p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (o *Organisation) Read(ctx context.Context, id string, inputs OrganisationArgs, state OrganisationState) (
	string, OrganisationArgs, OrganisationState, error,
) {
	config := infer.GetConfig[Config](ctx)

	organisation, err := o.getOrganisation(ctx, id, config)
	if err != nil {
		return "", OrganisationArgs{}, OrganisationState{}, err
	}

	if organisation != nil {
		return "", OrganisationArgs{}, OrganisationState{}, nil
	}

	return id, OrganisationArgs{
			Name: &organisation.Name,
		}, OrganisationState{
			Id:   organisation.Id,
			Name: organisation.Name,
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

func (o *Organisation) Update(ctx context.Context, id string, olds OrganisationState, news OrganisationArgs, preview bool) (OrganisationState, error) {
	orgId, err := strconv.Atoi(id)
	if err != nil {
		return OrganisationState{}, err
	}
	if !preview && olds.Name != *news.Name {
		config := infer.GetConfig[Config](ctx)
		err := o.updateOrganisation(ctx, orgId, *news.Name, config)
		if err != nil {
			return OrganisationState{}, err
		}
	}

	return OrganisationState{
		Id:   orgId,
		Name: *news.Name,
	}, nil
}

func (*Organisation) updateOrganisation(ctx context.Context, id int, newName string, config Config) error {
	err := config.Client.UpdateOrganisation(ctx, id, newName)
	if err != nil {
		return err
	}
	return nil
}

var _ infer.CustomCheck[OrganisationArgs] = ((*Organisation)(nil))

func (*Organisation) Check(
	ctx context.Context, name string, oldInputs, newInputs resource.PropertyMap,
) (OrganisationArgs, []p.CheckFailure, error) {
	// Apply default arguments
	args, failures, err := infer.DefaultCheck[OrganisationArgs](ctx, newInputs)
	if err != nil {
		return args, failures, err
	}

	// Apply autonaming
	//
	// If args.Name is unset, we set it to a value based off of the resource name.
	args.Name, err = autoname(args.Name, name, "orgName", oldInputs)
	return args, failures, err
}

type GetOrganisation struct{}

type GetOrganisationArgs struct {
	Id int `pulumi:"orgId"`
}

func (GetOrganisation) Call(ctx context.Context, args GetOrganisationArgs) (OrganisationState, error) {
	config := infer.GetConfig[Config](ctx)

	orgId := strconv.Itoa(args.Id)

	organisation, res, err := config.Client.GetOrganisation(ctx, orgId)
	if err != nil {
		if res.StatusCode == 404 {
			return OrganisationState{}, nil
		}
		return OrganisationState{}, err
	}
	return OrganisationState{
		Id:   organisation.Result.Id,
		Name: organisation.Result.Name,
	}, nil
}

func autoname(
	field *string, name string, fieldName resource.PropertyKey,
	oldInputs resource.PropertyMap,
) (*string, error) {
	if field != nil {
		return field, nil
	}

	prev := oldInputs[fieldName]
	if prev.IsSecret() {
		prev = prev.SecretValue().Element
	}

	if prev.IsString() && prev.StringValue() != "" {
		n := prev.StringValue()
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
