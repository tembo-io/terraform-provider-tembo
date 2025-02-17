package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	temboclient "github.com/tembo-io/terraform-provider-tembo/temboclient"
)

func mapInstanceToDataSourceModel(instance *temboclient.Instance, state *temboInstanceDataSourceModel) {
	// Convert the existing mapping functions' results to types.List
	extraDomainsElements := make([]attr.Value, len(instance.GetExtraDomainsRw()))
	for i, v := range instance.GetExtraDomainsRw() {
		extraDomainsElements[i] = types.StringValue(v)
	}
	state.ExtraDomainsRw, _ = types.ListValue(types.StringType, extraDomainsElements)

	// Use existing mapping functions for the rest
	state.IpAllowList = getStringArrayToTypes(instance.GetIpAllowList())
	state.ConnectionPooler = mapConnectionPoolerToObject(instance.GetConnectionPooler())

	// Map extensions with proper types
	extensionElements := make([]attr.Value, len(instance.GetExtensions()))
	for i, ext := range instance.GetExtensions() {
		locationElements := make([]attr.Value, len(ext.GetLocations()))
		for j, loc := range ext.GetLocations() {
			locationElements[j] = types.ObjectValueMust(
				map[string]attr.Type{
					"database": types.StringType,
					"enabled":  types.BoolType,
					"schema":   types.StringType,
					"version":  types.StringType,
				},
				map[string]attr.Value{
					"database": types.StringValue(loc.GetDatabase()),
					"enabled":  types.BoolValue(loc.GetEnabled()),
					"schema":   types.StringValue(loc.GetSchema()),
					"version":  types.StringValue(loc.GetVersion()),
				},
			)
		}
		locations, _ := types.ListValue(
			types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"database": types.StringType,
					"enabled":  types.BoolType,
					"schema":   types.StringType,
					"version":  types.StringType,
				},
			},
			locationElements,
		)
		extensionElements[i] = types.ObjectValueMust(
			map[string]attr.Type{
				"name":        types.StringType,
				"description": types.StringType,
				"locations": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"database": types.StringType,
					"enabled":  types.BoolType,
					"schema":   types.StringType,
					"version":  types.StringType,
				}}},
			},
			map[string]attr.Value{
				"name":        types.StringValue(ext.GetName()),
				"description": types.StringValue(ext.GetDescription()),
				"locations":   locations,
			},
		)
	}
	extensions, _ := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":        types.StringType,
				"description": types.StringType,
				"locations": types.ListType{ElemType: types.ObjectType{AttrTypes: map[string]attr.Type{
					"database": types.StringType,
					"enabled":  types.BoolType,
					"schema":   types.StringType,
					"version":  types.StringType,
				}}},
			},
		},
		extensionElements,
	)
	state.Extensions = extensions

	// Map postgres_configs with proper types
	pgConfigElements := make([]attr.Value, len(instance.GetPostgresConfigs()))
	for i, config := range instance.GetPostgresConfigs() {
		pgConfigElements[i] = types.ObjectValueMust(
			map[string]attr.Type{
				"name":  types.StringType,
				"value": types.StringType,
			},
			map[string]attr.Value{
				"name":  types.StringValue(config.GetName()),
				"value": types.StringValue(config.GetValue()),
			},
		)
	}
	pgConfigs, _ := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":  types.StringType,
				"value": types.StringType,
			},
		},
		pgConfigElements,
	)
	state.PostgresConfigs = pgConfigs

	// Map trunk_installs with proper types
	trunkInstallElements := make([]attr.Value, len(instance.GetTrunkInstalls()))
	for i, install := range instance.GetTrunkInstalls() {
		trunkInstallElements[i] = types.ObjectValueMust(
			map[string]attr.Type{
				"name":    types.StringType,
				"version": types.StringType,
			},
			map[string]attr.Value{
				"name":    types.StringValue(install.GetName()),
				"version": types.StringValue(install.GetVersion()),
			},
		)
	}
	trunkInstalls, _ := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":    types.StringType,
				"version": types.StringType,
			},
		},
		trunkInstallElements,
	)
	state.TrunkInstalls = trunkInstalls
}

func getStringArrayToTypes(items []string) []types.String {
	result := make([]types.String, len(items))
	for i, item := range items {
		result[i] = types.StringValue(item)
	}
	return result
}

func mapConnectionPoolerToObject(pooler temboclient.ConnectionPooler) types.Object {
	poolerAttrTypes := map[string]attr.Type{
		"parameters": types.MapType{ElemType: types.StringType},
		"pool_mode":  types.StringType,
	}
	poolerType := types.ObjectType{AttrTypes: poolerAttrTypes}

	if !pooler.GetEnabled() {
		obj, _ := types.ObjectValue(
			map[string]attr.Type{
				"enabled": types.BoolType,
				"pooler":  poolerType,
			},
			map[string]attr.Value{
				"enabled": types.BoolValue(false),
				"pooler":  types.ObjectNull(poolerAttrTypes),
			},
		)
		return obj
	}

	poolerConfig := pooler.GetPooler()
	parameters := make(map[string]attr.Value)
	for k, v := range poolerConfig.GetParameters() {
		parameters[k] = types.StringValue(v)
	}

	poolerObj, _ := types.ObjectValue(
		map[string]attr.Type{
			"parameters": types.MapType{ElemType: types.StringType},
			"pool_mode":  types.StringType,
		},
		map[string]attr.Value{
			"parameters": types.MapValueMust(types.StringType, parameters),
			"pool_mode":  types.StringValue(string(poolerConfig.GetPoolMode())),
		},
	)

	obj, _ := types.ObjectValue(
		map[string]attr.Type{
			"enabled": types.BoolType,
			"pooler": types.ObjectType{AttrTypes: map[string]attr.Type{
				"parameters": types.MapType{ElemType: types.StringType},
				"pool_mode":  types.StringType,
			}},
		},
		map[string]attr.Value{
			"enabled": types.BoolValue(pooler.GetEnabled()),
			"pooler":  poolerObj,
		},
	)
	return obj
}
