// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_objectplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/airbytehq/terraform-provider-airbyte/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/airbytehq/terraform-provider-airbyte/internal/provider/types"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/models/operations"
	"github.com/airbytehq/terraform-provider-airbyte/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &DestinationMssqlResource{}
var _ resource.ResourceWithImportState = &DestinationMssqlResource{}

func NewDestinationMssqlResource() resource.Resource {
	return &DestinationMssqlResource{}
}

// DestinationMssqlResource defines the resource implementation.
type DestinationMssqlResource struct {
	client *sdk.SDK
}

// DestinationMssqlResourceModel describes the resource data model.
type DestinationMssqlResourceModel struct {
	Configuration   tfTypes.DestinationMssql `tfsdk:"configuration"`
	DefinitionID    types.String             `tfsdk:"definition_id"`
	DestinationID   types.String             `tfsdk:"destination_id"`
	DestinationType types.String             `tfsdk:"destination_type"`
	Name            types.String             `tfsdk:"name"`
	WorkspaceID     types.String             `tfsdk:"workspace_id"`
}

func (r *DestinationMssqlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_destination_mssql"
}

func (r *DestinationMssqlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DestinationMssql Resource",
		Attributes: map[string]schema.Attribute{
			"configuration": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Required: true,
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Required:    true,
						Description: `The name of the MSSQL database.`,
					},
					"host": schema.StringAttribute{
						Required:    true,
						Description: `The host name of the MSSQL database.`,
					},
					"jdbc_url_params": schema.StringAttribute{
						Optional:    true,
						Description: `Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).`,
					},
					"password": schema.StringAttribute{
						Optional:    true,
						Sensitive:   true,
						Description: `The password associated with this username.`,
					},
					"port": schema.Int64Attribute{
						Computed:    true,
						Optional:    true,
						Default:     int64default.StaticInt64(1433),
						Description: `The port of the MSSQL database. Default: 1433`,
					},
					"raw_data_schema": schema.StringAttribute{
						Optional:    true,
						Description: `The schema to write raw tables into (default: airbyte_internal)`,
					},
					"schema": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Default:     stringdefault.StaticString("public"),
						Description: `The default schema tables are written to if the source does not specify a namespace. The usual value for this field is "public". Default: "public"`,
					},
					"ssl_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"encrypted_trust_server_certificate": schema.SingleNestedAttribute{
								Optional:    true,
								Attributes:  map[string]schema.Attribute{},
								Description: `Use the certificate provided by the server without verification. (For testing purposes only!)`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("encrypted_verify_certificate"),
									}...),
								},
							},
							"encrypted_verify_certificate": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"host_name_in_certificate": schema.StringAttribute{
										Optional:    true,
										Description: `Specifies the host name of the server. The value of this property must match the subject property of the certificate.`,
									},
								},
								Description: `Verify and use the certificate provided by the server.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("encrypted_trust_server_certificate"),
									}...),
								},
							},
						},
						Description: `The encryption method which is used to communicate with the database.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"tunnel_method": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"no_tunnel": schema.SingleNestedAttribute{
								Optional:   true,
								Attributes: map[string]schema.Attribute{},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("password_authentication"),
										path.MatchRelative().AtParent().AtName("ssh_key_authentication"),
									}...),
								},
							},
							"password_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(22),
										Description: `Port on the proxy/jump server that accepts inbound ssh connections. Default: 22`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host`,
									},
									"tunnel_user_password": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level password for logging into the jump server host`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("no_tunnel"),
										path.MatchRelative().AtParent().AtName("ssh_key_authentication"),
									}...),
								},
							},
							"ssh_key_authentication": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"ssh_key": schema.StringAttribute{
										Required:    true,
										Sensitive:   true,
										Description: `OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )`,
									},
									"tunnel_host": schema.StringAttribute{
										Required:    true,
										Description: `Hostname of the jump server host that allows inbound ssh tunnel.`,
									},
									"tunnel_port": schema.Int64Attribute{
										Computed:    true,
										Optional:    true,
										Default:     int64default.StaticInt64(22),
										Description: `Port on the proxy/jump server that accepts inbound ssh connections. Default: 22`,
									},
									"tunnel_user": schema.StringAttribute{
										Required:    true,
										Description: `OS-level username for logging into the jump server host.`,
									},
								},
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("no_tunnel"),
										path.MatchRelative().AtParent().AtName("password_authentication"),
									}...),
								},
							},
						},
						Description: `Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.`,
						Validators: []validator.Object{
							validators.ExactlyOneChild(),
						},
					},
					"username": schema.StringAttribute{
						Required:    true,
						Description: `The username which is used to access the database.`,
					},
				},
			},
			"definition_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The UUID of the connector definition. One of configuration.destinationType or definitionId must be provided. Requires replacement if changed. `,
			},
			"destination_id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"destination_type": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `Name of the destination e.g. dev-mysql-instance.`,
			},
			"workspace_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
		},
	}
}

func (r *DestinationMssqlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DestinationMssqlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *DestinationMssqlResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedDestinationMssqlCreateRequest()
	res, err := r.client.Destinations.CreateDestinationMssql(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	destinationID := data.DestinationID.ValueString()
	request1 := operations.GetDestinationMssqlRequest{
		DestinationID: destinationID,
	}
	res1, err := r.client.Destinations.GetDestinationMssql(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationMssqlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *DestinationMssqlResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.GetDestinationMssqlRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.GetDestinationMssql(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res.DestinationResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationMssqlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *DestinationMssqlResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	destinationMssqlPutRequest := data.ToSharedDestinationMssqlPutRequest()
	destinationID := data.DestinationID.ValueString()
	request := operations.PutDestinationMssqlRequest{
		DestinationMssqlPutRequest: destinationMssqlPutRequest,
		DestinationID:              destinationID,
	}
	res, err := r.client.Destinations.PutDestinationMssql(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	destinationId1 := data.DestinationID.ValueString()
	request1 := operations.GetDestinationMssqlRequest{
		DestinationID: destinationId1,
	}
	res1, err := r.client.Destinations.GetDestinationMssql(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.DestinationResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedDestinationResponse(res1.DestinationResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DestinationMssqlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *DestinationMssqlResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	destinationID := data.DestinationID.ValueString()
	request := operations.DeleteDestinationMssqlRequest{
		DestinationID: destinationID,
	}
	res, err := r.client.Destinations.DeleteDestinationMssql(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if fmt.Sprintf("%v", res.StatusCode)[0] != '2' {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *DestinationMssqlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("destination_id"), req.ID)...)
}
