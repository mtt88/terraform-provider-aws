// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ecrpublic

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceAuthorizationToken,
			TypeName: "aws_ecrpublic_authorization_token",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceRepository,
			TypeName: "aws_ecrpublic_repository",
		},
		{
			Factory:  ResourceRepositoryPolicy,
			TypeName: "aws_ecrpublic_repository_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ECRPublic
}

var ServicePackage = &servicePackage{}
