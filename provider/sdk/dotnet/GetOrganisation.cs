// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace MyNamespace.NativeProvider
{
    public static class GetOrganisation
    {
        public static Task<GetOrganisationResult> InvokeAsync(GetOrganisationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrganisationResult>("nativeProvider:index:getOrganisation", args ?? new GetOrganisationArgs(), options.WithDefaults());

        public static Output<GetOrganisationResult> Invoke(GetOrganisationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrganisationResult>("nativeProvider:index:getOrganisation", args ?? new GetOrganisationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrganisationArgs : global::Pulumi.InvokeArgs
    {
        [Input("orgId", required: true)]
        public int OrgId { get; set; }

        public GetOrganisationArgs()
        {
        }
        public static new GetOrganisationArgs Empty => new GetOrganisationArgs();
    }

    public sealed class GetOrganisationInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("orgId", required: true)]
        public Input<int> OrgId { get; set; } = null!;

        public GetOrganisationInvokeArgs()
        {
        }
        public static new GetOrganisationInvokeArgs Empty => new GetOrganisationInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrganisationResult
    {
        /// <summary>
        /// Id of the organisation created
        /// </summary>
        public readonly int OrgId;
        /// <summary>
        /// Name of organisation created
        /// </summary>
        public readonly string OrgName;

        [OutputConstructor]
        private GetOrganisationResult(
            int orgId,

            string orgName)
        {
            OrgId = orgId;
            OrgName = orgName;
        }
    }
}