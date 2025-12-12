package nats

import (
	ansiblecfg "github.com/open-uem/openuem-ansible-config/ansible"
	"github.com/open-uem/wingetcfg/wingetcfg"
)

type CfgProfiles struct {
	AgentID string `json:"agentID,omitempty"`
}

type ProfileConfig struct {
	ProfileID     int                           `yaml:"profileID"`
	Exclusions    []string                      `yaml:"exclusions"`
	Deployments   []string                      `yaml:"deployments"`
	WinGetConfig  *wingetcfg.WinGetCfg          `yaml:"config,omitempty"`
	AnsibleConfig []*ansiblecfg.AnsiblePlaybook `yaml:"ansible,omitempty"`
	NetBirdConfig []*NetbirdTask                `yaml:"netbird,omitempty"`
}
