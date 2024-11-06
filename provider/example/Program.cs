using System.Collections.Generic;
using Pulumi;
using MyNamespace.NativeProvider;

return await Deployment.RunAsync(() =>
{
    // Add your resources here
    // e.g. var resource = new Resource("name", new ResourceArgs { });

    var org = new Organisation("myorg", new()
    {
        OrgName = "foobar"
    });

    // Export outputs here
    return new Dictionary<string, object?>
    {
        ["outputKey"] = "outputValue"
    };
});
