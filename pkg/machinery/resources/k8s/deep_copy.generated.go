// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type AdmissionControlConfigSpec -type APIServerConfigSpec -type AuditPolicyConfigSpec -type BootstrapManifestsConfigSpec -type ConfigStatusSpec -type ControllerManagerConfigSpec -type EndpointSpec -type ExtraManifestsConfigSpec -type KubeletLifecycleSpec -type KubePrismConfigSpec -type KubePrismEndpointsSpec -type KubePrismStatusesSpec -type KubeletSpecSpec -type ManifestSpec -type ManifestStatusSpec -type NodeCordonedSpecSpec -type NodeLabelSpecSpec -type NodeTaintSpecSpec -type KubeletConfigSpec -type NodeIPSpec -type NodeIPConfigSpec -type NodeStatusSpec -type NodenameSpec -type SchedulerConfigSpec -type SecretsStatusSpec -type StaticPodSpec -type StaticPodStatusSpec -type StaticPodServerStatusSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package k8s

import (
	"net/netip"

	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// DeepCopy generates a deep copy of AdmissionControlConfigSpec.
func (o AdmissionControlConfigSpec) DeepCopy() AdmissionControlConfigSpec {
	var cp AdmissionControlConfigSpec = o
	if o.Config != nil {
		cp.Config = make([]AdmissionPluginSpec, len(o.Config))
		copy(cp.Config, o.Config)
		for i2 := range o.Config {
			if o.Config[i2].Configuration != nil {
				cp.Config[i2].Configuration = make(map[string]interface{}, len(o.Config[i2].Configuration))
				for k4, v4 := range o.Config[i2].Configuration {
					cp.Config[i2].Configuration[k4] = v4
				}
			}
		}
	}
	return cp
}

// DeepCopy generates a deep copy of APIServerConfigSpec.
func (o APIServerConfigSpec) DeepCopy() APIServerConfigSpec {
	var cp APIServerConfigSpec = o
	if o.EtcdServers != nil {
		cp.EtcdServers = make([]string, len(o.EtcdServers))
		copy(cp.EtcdServers, o.EtcdServers)
	}
	if o.ServiceCIDRs != nil {
		cp.ServiceCIDRs = make([]string, len(o.ServiceCIDRs))
		copy(cp.ServiceCIDRs, o.ServiceCIDRs)
	}
	if o.ExtraArgs != nil {
		cp.ExtraArgs = make(map[string]string, len(o.ExtraArgs))
		for k2, v2 := range o.ExtraArgs {
			cp.ExtraArgs[k2] = v2
		}
	}
	if o.ExtraVolumes != nil {
		cp.ExtraVolumes = make([]ExtraVolume, len(o.ExtraVolumes))
		copy(cp.ExtraVolumes, o.ExtraVolumes)
	}
	if o.EnvironmentVariables != nil {
		cp.EnvironmentVariables = make(map[string]string, len(o.EnvironmentVariables))
		for k2, v2 := range o.EnvironmentVariables {
			cp.EnvironmentVariables[k2] = v2
		}
	}
	if o.Resources.Requests != nil {
		cp.Resources.Requests = make(map[string]string, len(o.Resources.Requests))
		for k3, v3 := range o.Resources.Requests {
			cp.Resources.Requests[k3] = v3
		}
	}
	if o.Resources.Limits != nil {
		cp.Resources.Limits = make(map[string]string, len(o.Resources.Limits))
		for k3, v3 := range o.Resources.Limits {
			cp.Resources.Limits[k3] = v3
		}
	}
	return cp
}

// DeepCopy generates a deep copy of AuditPolicyConfigSpec.
func (o AuditPolicyConfigSpec) DeepCopy() AuditPolicyConfigSpec {
	var cp AuditPolicyConfigSpec = o
	if o.Config != nil {
		cp.Config = make(map[string]interface{}, len(o.Config))
		for k2, v2 := range o.Config {
			cp.Config[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of BootstrapManifestsConfigSpec.
func (o BootstrapManifestsConfigSpec) DeepCopy() BootstrapManifestsConfigSpec {
	var cp BootstrapManifestsConfigSpec = o
	if o.PodCIDRs != nil {
		cp.PodCIDRs = make([]string, len(o.PodCIDRs))
		copy(cp.PodCIDRs, o.PodCIDRs)
	}
	if o.ProxyArgs != nil {
		cp.ProxyArgs = make([]string, len(o.ProxyArgs))
		copy(cp.ProxyArgs, o.ProxyArgs)
	}
	return cp
}

// DeepCopy generates a deep copy of ConfigStatusSpec.
func (o ConfigStatusSpec) DeepCopy() ConfigStatusSpec {
	var cp ConfigStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of ControllerManagerConfigSpec.
func (o ControllerManagerConfigSpec) DeepCopy() ControllerManagerConfigSpec {
	var cp ControllerManagerConfigSpec = o
	if o.PodCIDRs != nil {
		cp.PodCIDRs = make([]string, len(o.PodCIDRs))
		copy(cp.PodCIDRs, o.PodCIDRs)
	}
	if o.ServiceCIDRs != nil {
		cp.ServiceCIDRs = make([]string, len(o.ServiceCIDRs))
		copy(cp.ServiceCIDRs, o.ServiceCIDRs)
	}
	if o.ExtraArgs != nil {
		cp.ExtraArgs = make(map[string]string, len(o.ExtraArgs))
		for k2, v2 := range o.ExtraArgs {
			cp.ExtraArgs[k2] = v2
		}
	}
	if o.ExtraVolumes != nil {
		cp.ExtraVolumes = make([]ExtraVolume, len(o.ExtraVolumes))
		copy(cp.ExtraVolumes, o.ExtraVolumes)
	}
	if o.EnvironmentVariables != nil {
		cp.EnvironmentVariables = make(map[string]string, len(o.EnvironmentVariables))
		for k2, v2 := range o.EnvironmentVariables {
			cp.EnvironmentVariables[k2] = v2
		}
	}
	if o.Resources.Requests != nil {
		cp.Resources.Requests = make(map[string]string, len(o.Resources.Requests))
		for k3, v3 := range o.Resources.Requests {
			cp.Resources.Requests[k3] = v3
		}
	}
	if o.Resources.Limits != nil {
		cp.Resources.Limits = make(map[string]string, len(o.Resources.Limits))
		for k3, v3 := range o.Resources.Limits {
			cp.Resources.Limits[k3] = v3
		}
	}
	return cp
}

// DeepCopy generates a deep copy of EndpointSpec.
func (o EndpointSpec) DeepCopy() EndpointSpec {
	var cp EndpointSpec = o
	if o.Addresses != nil {
		cp.Addresses = make([]netip.Addr, len(o.Addresses))
		copy(cp.Addresses, o.Addresses)
	}
	return cp
}

// DeepCopy generates a deep copy of ExtraManifestsConfigSpec.
func (o ExtraManifestsConfigSpec) DeepCopy() ExtraManifestsConfigSpec {
	var cp ExtraManifestsConfigSpec = o
	if o.ExtraManifests != nil {
		cp.ExtraManifests = make([]ExtraManifest, len(o.ExtraManifests))
		copy(cp.ExtraManifests, o.ExtraManifests)
		for i2 := range o.ExtraManifests {
			if o.ExtraManifests[i2].ExtraHeaders != nil {
				cp.ExtraManifests[i2].ExtraHeaders = make(map[string]string, len(o.ExtraManifests[i2].ExtraHeaders))
				for k4, v4 := range o.ExtraManifests[i2].ExtraHeaders {
					cp.ExtraManifests[i2].ExtraHeaders[k4] = v4
				}
			}
		}
	}
	return cp
}

// DeepCopy generates a deep copy of KubeletLifecycleSpec.
func (o KubeletLifecycleSpec) DeepCopy() KubeletLifecycleSpec {
	var cp KubeletLifecycleSpec = o
	return cp
}

// DeepCopy generates a deep copy of KubePrismConfigSpec.
func (o KubePrismConfigSpec) DeepCopy() KubePrismConfigSpec {
	var cp KubePrismConfigSpec = o
	if o.Endpoints != nil {
		cp.Endpoints = make([]KubePrismEndpoint, len(o.Endpoints))
		copy(cp.Endpoints, o.Endpoints)
	}
	return cp
}

// DeepCopy generates a deep copy of KubePrismEndpointsSpec.
func (o KubePrismEndpointsSpec) DeepCopy() KubePrismEndpointsSpec {
	var cp KubePrismEndpointsSpec = o
	if o.Endpoints != nil {
		cp.Endpoints = make([]KubePrismEndpoint, len(o.Endpoints))
		copy(cp.Endpoints, o.Endpoints)
	}
	return cp
}

// DeepCopy generates a deep copy of KubePrismStatusesSpec.
func (o KubePrismStatusesSpec) DeepCopy() KubePrismStatusesSpec {
	var cp KubePrismStatusesSpec = o
	return cp
}

// DeepCopy generates a deep copy of KubeletSpecSpec.
func (o KubeletSpecSpec) DeepCopy() KubeletSpecSpec {
	var cp KubeletSpecSpec = o
	if o.Args != nil {
		cp.Args = make([]string, len(o.Args))
		copy(cp.Args, o.Args)
	}
	if o.ExtraMounts != nil {
		cp.ExtraMounts = make([]specs.Mount, len(o.ExtraMounts))
		copy(cp.ExtraMounts, o.ExtraMounts)
		for i2 := range o.ExtraMounts {
			if o.ExtraMounts[i2].Options != nil {
				cp.ExtraMounts[i2].Options = make([]string, len(o.ExtraMounts[i2].Options))
				copy(cp.ExtraMounts[i2].Options, o.ExtraMounts[i2].Options)
			}
			if o.ExtraMounts[i2].UIDMappings != nil {
				cp.ExtraMounts[i2].UIDMappings = make([]specs.LinuxIDMapping, len(o.ExtraMounts[i2].UIDMappings))
				copy(cp.ExtraMounts[i2].UIDMappings, o.ExtraMounts[i2].UIDMappings)
			}
			if o.ExtraMounts[i2].GIDMappings != nil {
				cp.ExtraMounts[i2].GIDMappings = make([]specs.LinuxIDMapping, len(o.ExtraMounts[i2].GIDMappings))
				copy(cp.ExtraMounts[i2].GIDMappings, o.ExtraMounts[i2].GIDMappings)
			}
		}
	}
	if o.Config != nil {
		cp.Config = make(map[string]any, len(o.Config))
		for k2, v2 := range o.Config {
			cp.Config[k2] = v2
		}
	}
	if o.CredentialProviderConfig != nil {
		cp.CredentialProviderConfig = make(map[string]any, len(o.CredentialProviderConfig))
		for k2, v2 := range o.CredentialProviderConfig {
			cp.CredentialProviderConfig[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of ManifestSpec.
func (o ManifestSpec) DeepCopy() ManifestSpec {
	var cp ManifestSpec = o
	if o.Items != nil {
		cp.Items = make([]SingleManifest, len(o.Items))
		copy(cp.Items, o.Items)
		for i2 := range o.Items {
			if o.Items[i2].Object != nil {
				cp.Items[i2].Object = make(map[string]interface{}, len(o.Items[i2].Object))
				for k4, v4 := range o.Items[i2].Object {
					cp.Items[i2].Object[k4] = v4
				}
			}
		}
	}
	return cp
}

// DeepCopy generates a deep copy of ManifestStatusSpec.
func (o ManifestStatusSpec) DeepCopy() ManifestStatusSpec {
	var cp ManifestStatusSpec = o
	if o.ManifestsApplied != nil {
		cp.ManifestsApplied = make([]string, len(o.ManifestsApplied))
		copy(cp.ManifestsApplied, o.ManifestsApplied)
	}
	return cp
}

// DeepCopy generates a deep copy of NodeCordonedSpecSpec.
func (o NodeCordonedSpecSpec) DeepCopy() NodeCordonedSpecSpec {
	var cp NodeCordonedSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of NodeLabelSpecSpec.
func (o NodeLabelSpecSpec) DeepCopy() NodeLabelSpecSpec {
	var cp NodeLabelSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of NodeTaintSpecSpec.
func (o NodeTaintSpecSpec) DeepCopy() NodeTaintSpecSpec {
	var cp NodeTaintSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of KubeletConfigSpec.
func (o KubeletConfigSpec) DeepCopy() KubeletConfigSpec {
	var cp KubeletConfigSpec = o
	if o.ClusterDNS != nil {
		cp.ClusterDNS = make([]string, len(o.ClusterDNS))
		copy(cp.ClusterDNS, o.ClusterDNS)
	}
	if o.ExtraArgs != nil {
		cp.ExtraArgs = make(map[string]string, len(o.ExtraArgs))
		for k2, v2 := range o.ExtraArgs {
			cp.ExtraArgs[k2] = v2
		}
	}
	if o.ExtraMounts != nil {
		cp.ExtraMounts = make([]specs.Mount, len(o.ExtraMounts))
		copy(cp.ExtraMounts, o.ExtraMounts)
		for i2 := range o.ExtraMounts {
			if o.ExtraMounts[i2].Options != nil {
				cp.ExtraMounts[i2].Options = make([]string, len(o.ExtraMounts[i2].Options))
				copy(cp.ExtraMounts[i2].Options, o.ExtraMounts[i2].Options)
			}
			if o.ExtraMounts[i2].UIDMappings != nil {
				cp.ExtraMounts[i2].UIDMappings = make([]specs.LinuxIDMapping, len(o.ExtraMounts[i2].UIDMappings))
				copy(cp.ExtraMounts[i2].UIDMappings, o.ExtraMounts[i2].UIDMappings)
			}
			if o.ExtraMounts[i2].GIDMappings != nil {
				cp.ExtraMounts[i2].GIDMappings = make([]specs.LinuxIDMapping, len(o.ExtraMounts[i2].GIDMappings))
				copy(cp.ExtraMounts[i2].GIDMappings, o.ExtraMounts[i2].GIDMappings)
			}
		}
	}
	if o.ExtraConfig != nil {
		cp.ExtraConfig = make(map[string]any, len(o.ExtraConfig))
		for k2, v2 := range o.ExtraConfig {
			cp.ExtraConfig[k2] = v2
		}
	}
	if o.CredentialProviderConfig != nil {
		cp.CredentialProviderConfig = make(map[string]any, len(o.CredentialProviderConfig))
		for k2, v2 := range o.CredentialProviderConfig {
			cp.CredentialProviderConfig[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of NodeIPSpec.
func (o NodeIPSpec) DeepCopy() NodeIPSpec {
	var cp NodeIPSpec = o
	if o.Addresses != nil {
		cp.Addresses = make([]netip.Addr, len(o.Addresses))
		copy(cp.Addresses, o.Addresses)
	}
	return cp
}

// DeepCopy generates a deep copy of NodeIPConfigSpec.
func (o NodeIPConfigSpec) DeepCopy() NodeIPConfigSpec {
	var cp NodeIPConfigSpec = o
	if o.ValidSubnets != nil {
		cp.ValidSubnets = make([]string, len(o.ValidSubnets))
		copy(cp.ValidSubnets, o.ValidSubnets)
	}
	if o.ExcludeSubnets != nil {
		cp.ExcludeSubnets = make([]string, len(o.ExcludeSubnets))
		copy(cp.ExcludeSubnets, o.ExcludeSubnets)
	}
	return cp
}

// DeepCopy generates a deep copy of NodeStatusSpec.
func (o NodeStatusSpec) DeepCopy() NodeStatusSpec {
	var cp NodeStatusSpec = o
	if o.Labels != nil {
		cp.Labels = make(map[string]string, len(o.Labels))
		for k2, v2 := range o.Labels {
			cp.Labels[k2] = v2
		}
	}
	if o.Annotations != nil {
		cp.Annotations = make(map[string]string, len(o.Annotations))
		for k2, v2 := range o.Annotations {
			cp.Annotations[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of NodenameSpec.
func (o NodenameSpec) DeepCopy() NodenameSpec {
	var cp NodenameSpec = o
	return cp
}

// DeepCopy generates a deep copy of SchedulerConfigSpec.
func (o SchedulerConfigSpec) DeepCopy() SchedulerConfigSpec {
	var cp SchedulerConfigSpec = o
	if o.ExtraArgs != nil {
		cp.ExtraArgs = make(map[string]string, len(o.ExtraArgs))
		for k2, v2 := range o.ExtraArgs {
			cp.ExtraArgs[k2] = v2
		}
	}
	if o.ExtraVolumes != nil {
		cp.ExtraVolumes = make([]ExtraVolume, len(o.ExtraVolumes))
		copy(cp.ExtraVolumes, o.ExtraVolumes)
	}
	if o.EnvironmentVariables != nil {
		cp.EnvironmentVariables = make(map[string]string, len(o.EnvironmentVariables))
		for k2, v2 := range o.EnvironmentVariables {
			cp.EnvironmentVariables[k2] = v2
		}
	}
	if o.Resources.Requests != nil {
		cp.Resources.Requests = make(map[string]string, len(o.Resources.Requests))
		for k3, v3 := range o.Resources.Requests {
			cp.Resources.Requests[k3] = v3
		}
	}
	if o.Resources.Limits != nil {
		cp.Resources.Limits = make(map[string]string, len(o.Resources.Limits))
		for k3, v3 := range o.Resources.Limits {
			cp.Resources.Limits[k3] = v3
		}
	}
	return cp
}

// DeepCopy generates a deep copy of SecretsStatusSpec.
func (o SecretsStatusSpec) DeepCopy() SecretsStatusSpec {
	var cp SecretsStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of StaticPodSpec.
func (o StaticPodSpec) DeepCopy() StaticPodSpec {
	var cp StaticPodSpec = o
	if o.Pod != nil {
		cp.Pod = make(map[string]any, len(o.Pod))
		for k2, v2 := range o.Pod {
			cp.Pod[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of StaticPodStatusSpec.
func (o StaticPodStatusSpec) DeepCopy() StaticPodStatusSpec {
	var cp StaticPodStatusSpec = o
	if o.PodStatus != nil {
		cp.PodStatus = make(map[string]any, len(o.PodStatus))
		for k2, v2 := range o.PodStatus {
			cp.PodStatus[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of StaticPodServerStatusSpec.
func (o StaticPodServerStatusSpec) DeepCopy() StaticPodServerStatusSpec {
	var cp StaticPodServerStatusSpec = o
	return cp
}
