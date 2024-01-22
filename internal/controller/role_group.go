package controller

import (
	"github.com/zncdata-labs/airbyte-operator/api/v1alpha1"
	"reflect"
)

type Role interface {
	GetRoleValue(key string) (any, error)
}

type RoleGroup interface {
	GetRole(role string) (Role, error)
}

type RoleConfig interface {
	GetRoleConfig(key string) (any, error)
}

type ClusterConfig interface {
	GetClusterConfig(key string) (any, error)
}

type ServerGroups map[string]RoleGroup

func InitConfig(airbyte v1alpha1.Airbyte) map[string]ServerGroups {
	configs := make(map[string]ServerGroups)
	v := reflect.TypeOf(airbyte.Spec)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Name == "ClusterConfig" {
			continue
		}
	}

	return configs
}
