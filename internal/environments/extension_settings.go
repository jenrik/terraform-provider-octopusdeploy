package environments

import (
	"fmt"

	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/environments"
	"github.com/OctopusDeploy/go-octopusdeploy/v2/pkg/extensions"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ExpandJiraExtensionSettings(extensionSettings interface{}) extensions.ExtensionSettings {
	values := extensionSettings.([]interface{})
	valuesMap := values[0].(map[string]interface{})
	return environments.NewJiraExtensionSettings(
		valuesMap["environment_type"].(string),
	)
}

func ExpandJiraServiceManagementExtensionSettings(extensionSettings interface{}) extensions.ExtensionSettings {
	values := extensionSettings.([]interface{})
	valuesMap := values[0].(map[string]interface{})
	return environments.NewJiraServiceManagementExtensionSettings(
		valuesMap["is_enabled"].(bool),
	)
}

func ExpandServiceNowExtensionSettings(extensionSettings interface{}) extensions.ExtensionSettings {
	values := extensionSettings.([]interface{})
	valuesMap := values[0].(map[string]interface{})
	return environments.NewServiceNowExtensionSettings(
		valuesMap["is_enabled"].(bool),
	)
}

func GetJiraExtensionSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"environment_type": {
			Description: "The Jira environment type of this Octopus deployment environment.",
			Required:    true,
			Type:        schema.TypeString,
			ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{
				"development",
				"production",
				"staging",
				"testing",
				"unmapped",
			}, false)),
		},
	}
}

func GetJiraServiceManagementExtensionSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"is_enabled": {
			Description: "Specifies whether or not this extension is enabled for this project.",
			Required:    true,
			Type:        schema.TypeBool,
		},
	}
}

func GetServiceNowExtensionSettingsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"is_enabled": {
			Description: "Specifies whether or not this extension is enabled for this project.",
			Required:    true,
			Type:        schema.TypeBool,
		},
	}
}

func SetExtensionSettings(d *schema.ResourceData, extensionSettingsCollection []extensions.ExtensionSettings) error {
	for _, extensionSettings := range extensionSettingsCollection {
		switch extensionSettings.ExtensionID() {
		case extensions.ExtensionIDJira:
			if jiraExtensionSettings, ok := extensionSettings.(*environments.JiraExtensionSettings); ok {
				if err := d.Set("jira_extension_settings", FlattenJiraExtensionSettings(jiraExtensionSettings)); err != nil {
					return fmt.Errorf("error setting extension settings for Jira: %s", err)
				}
			}
		case extensions.ExtensionIDJiraServiceManagement:
			if jiraServiceManagementExtensionSettings, ok := extensionSettings.(*environments.JiraServiceManagementExtensionSettings); ok {
				if err := d.Set("jira_service_management_extension_settings", FlattenJiraServiceManagementExtensionSettings(jiraServiceManagementExtensionSettings)); err != nil {
					return fmt.Errorf("error setting extension settings for Jira Service Management (JSM): %s", err)
				}
			}
		case extensions.ExtensionIDServiceNow:
			if serviceNowExtensionSettings, ok := extensionSettings.(*environments.ServiceNowExtensionSettings); ok {
				if err := d.Set("servicenow_extension_settings", FlattenServiceNowExtensionSettings(serviceNowExtensionSettings)); err != nil {
					return fmt.Errorf("error setting extension settings for ServiceNow: %s", err)
				}
			}
		}
	}

	return nil
}

func FlattenJiraExtensionSettings(jiraExtensionSettings *environments.JiraExtensionSettings) []interface{} {
	if jiraExtensionSettings == nil {
		return nil
	}

	flattenedJiraExtensionSettings := make(map[string]interface{})
	flattenedJiraExtensionSettings["environment_type"] = jiraExtensionSettings.JiraEnvironmentType
	return []interface{}{flattenedJiraExtensionSettings}
}

func FlattenJiraServiceManagementExtensionSettings(jiraServiceManagementExtensionSettings *environments.JiraServiceManagementExtensionSettings) []interface{} {
	if jiraServiceManagementExtensionSettings == nil {
		return nil
	}

	flattenedJiraServiceManagementExtensionSettings := make(map[string]interface{})
	flattenedJiraServiceManagementExtensionSettings["is_enabled"] = jiraServiceManagementExtensionSettings.IsChangeControlled()
	return []interface{}{flattenedJiraServiceManagementExtensionSettings}
}

func FlattenServiceNowExtensionSettings(serviceNowExtensionSettings *environments.ServiceNowExtensionSettings) []interface{} {
	if serviceNowExtensionSettings == nil {
		return nil
	}

	flattenedServiceNowExtensionSettings := make(map[string]interface{})
	flattenedServiceNowExtensionSettings["is_enabled"] = serviceNowExtensionSettings.IsChangeControlled()
	return []interface{}{flattenedServiceNowExtensionSettings}
}
