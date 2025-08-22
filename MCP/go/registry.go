package main

import (
	"github.com/amazon-vpc-lattice/mcp-server/config"
	"github.com/amazon-vpc-lattice/mcp-server/models"
	tools_accesslogsubscriptions_resourceidentifier "github.com/amazon-vpc-lattice/mcp-server/tools/accesslogsubscriptions_resourceidentifier"
	tools_services "github.com/amazon-vpc-lattice/mcp-server/tools/services"
	tools_servicenetworkserviceassociations "github.com/amazon-vpc-lattice/mcp-server/tools/servicenetworkserviceassociations"
	tools_targetgroups "github.com/amazon-vpc-lattice/mcp-server/tools/targetgroups"
	tools_servicenetworkvpcassociations "github.com/amazon-vpc-lattice/mcp-server/tools/servicenetworkvpcassociations"
	tools_resourcepolicy "github.com/amazon-vpc-lattice/mcp-server/tools/resourcepolicy"
	tools_tags "github.com/amazon-vpc-lattice/mcp-server/tools/tags"
	tools_accesslogsubscriptions "github.com/amazon-vpc-lattice/mcp-server/tools/accesslogsubscriptions"
	tools_authpolicy "github.com/amazon-vpc-lattice/mcp-server/tools/authpolicy"
	tools_servicenetworks "github.com/amazon-vpc-lattice/mcp-server/tools/servicenetworks"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_accesslogsubscriptions_resourceidentifier.CreateListaccesslogsubscriptionsTool(cfg),
		tools_services.CreateDeleteruleTool(cfg),
		tools_services.CreateGetruleTool(cfg),
		tools_services.CreateUpdateruleTool(cfg),
		tools_servicenetworkserviceassociations.CreateCreateservicenetworkserviceassociationTool(cfg),
		tools_servicenetworkserviceassociations.CreateListservicenetworkserviceassociationsTool(cfg),
		tools_services.CreateListrulesTool(cfg),
		tools_services.CreateBatchupdateruleTool(cfg),
		tools_services.CreateCreateruleTool(cfg),
		tools_servicenetworkserviceassociations.CreateDeleteservicenetworkserviceassociationTool(cfg),
		tools_servicenetworkserviceassociations.CreateGetservicenetworkserviceassociationTool(cfg),
		tools_targetgroups.CreateListtargetsTool(cfg),
		tools_servicenetworkvpcassociations.CreateDeleteservicenetworkvpcassociationTool(cfg),
		tools_servicenetworkvpcassociations.CreateGetservicenetworkvpcassociationTool(cfg),
		tools_servicenetworkvpcassociations.CreateUpdateservicenetworkvpcassociationTool(cfg),
		tools_resourcepolicy.CreateDeleteresourcepolicyTool(cfg),
		tools_resourcepolicy.CreateGetresourcepolicyTool(cfg),
		tools_resourcepolicy.CreatePutresourcepolicyTool(cfg),
		tools_services.CreateDeleteserviceTool(cfg),
		tools_services.CreateGetserviceTool(cfg),
		tools_services.CreateUpdateserviceTool(cfg),
		tools_targetgroups.CreateListtargetgroupsTool(cfg),
		tools_targetgroups.CreateCreatetargetgroupTool(cfg),
		tools_targetgroups.CreateDeletetargetgroupTool(cfg),
		tools_targetgroups.CreateGettargetgroupTool(cfg),
		tools_targetgroups.CreateUpdatetargetgroupTool(cfg),
		tools_services.CreateListservicesTool(cfg),
		tools_services.CreateCreateserviceTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_accesslogsubscriptions.CreateCreateaccesslogsubscriptionTool(cfg),
		tools_accesslogsubscriptions.CreateDeleteaccesslogsubscriptionTool(cfg),
		tools_accesslogsubscriptions.CreateGetaccesslogsubscriptionTool(cfg),
		tools_accesslogsubscriptions.CreateUpdateaccesslogsubscriptionTool(cfg),
		tools_targetgroups.CreateDeregistertargetsTool(cfg),
		tools_authpolicy.CreateDeleteauthpolicyTool(cfg),
		tools_authpolicy.CreateGetauthpolicyTool(cfg),
		tools_authpolicy.CreatePutauthpolicyTool(cfg),
		tools_servicenetworks.CreateListservicenetworksTool(cfg),
		tools_servicenetworks.CreateCreateservicenetworkTool(cfg),
		tools_services.CreateDeletelistenerTool(cfg),
		tools_services.CreateGetlistenerTool(cfg),
		tools_services.CreateUpdatelistenerTool(cfg),
		tools_servicenetworks.CreateDeleteservicenetworkTool(cfg),
		tools_servicenetworks.CreateGetservicenetworkTool(cfg),
		tools_servicenetworks.CreateUpdateservicenetworkTool(cfg),
		tools_services.CreateListlistenersTool(cfg),
		tools_services.CreateCreatelistenerTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_servicenetworkvpcassociations.CreateListservicenetworkvpcassociationsTool(cfg),
		tools_servicenetworkvpcassociations.CreateCreateservicenetworkvpcassociationTool(cfg),
		tools_targetgroups.CreateRegistertargetsTool(cfg),
	}
}
