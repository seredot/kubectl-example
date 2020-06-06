package main

import (
	"errors"
	"strings"
)

var resources [][]string

func resourceList() [][]string {
	if resources == nil {
		resources = [][]string{
			{"bindings", "Binding"},
			{"componentstatuses", "cs", "ComponentStatus"},
			{"configmaps", "cm", "ConfigMap"},
			{"endpoints", "ep", "Endpoints"},
			{"events", "ev", "Event"},
			{"limitranges", "limits", "LimitRange"},
			{"namespaces", "ns", "Namespace"},
			{"nodes", "no", "Node"},
			{"persistentvolumeclaims", "pvc", "PersistentVolumeClaim"},
			{"persistentvolumes", "pv", "PersistentVolume"},
			{"pods", "po", "Pod"},
			{"podtemplates", "PodTemplate"},
			{"replicationcontrollers", "rc", "ReplicationController"},
			{"resourcequotas", "quota", "ResourceQuota"},
			{"secrets", "Secret"},
			{"serviceaccounts", "sa", "ServiceAccount"},
			{"services", "svc", "Service"},
			{"mutatingwebhookconfigurations", "MutatingWebhookConfiguration"},
			{"validatingwebhookconfigurations", "ValidatingWebhookConfiguration"},
			{"customresourcedefinitions", "crd", "crds", "CustomResourceDefinition"},
			{"apiservices", "APIService"},
			{"controllerrevisions", "ControllerRevision"},
			{"daemonsets", "ds", "DaemonSet"},
			{"deployments", "deploy", "Deployment"},
			{"replicasets", "rs", "ReplicaSet"},
			{"statefulsets", "sts", "StatefulSet"},
			{"tokenreviews", "TokenReview"},
			{"localsubjectaccessreviews", "LocalSubjectAccessReview"},
			{"selfsubjectaccessreviews", "SelfSubjectAccessReview"},
			{"selfsubjectrulesreviews", "SelfSubjectRulesReview"},
			{"subjectaccessreviews", "SubjectAccessReview"},
			{"horizontalpodautoscalers", "hpa", "HorizontalPodAutoscaler"},
			{"cronjobs", "cj", "CronJob"},
			{"jobs", "Job"},
			{"certificatesigningrequests", "csr", "CertificateSigningRequest"},
			{"leases", "Lease"},
			{"endpointslices", "EndpointSlice"},
			{"events", "ev", "Event"},
			{"ingresses", "ing", "Ingress"},
			{"ingressclasses", "IngressClass"},
			{"ingresses", "ing", "Ingress"},
			{"networkpolicies", "netpol", "NetworkPolicy"},
			{"runtimeclasses", "RuntimeClass"},
			{"poddisruptionbudgets", "pdb", "PodDisruptionBudget"},
			{"podsecuritypolicies", "psp", "PodSecurityPolicy"},
			{"clusterrolebindings", "ClusterRoleBinding"},
			{"clusterroles", "ClusterRole"},
			{"rolebindings", "RoleBinding"},
			{"roles", "Role"},
			{"priorityclasses", "pc", "PriorityClass"},
			{"csidrivers", "CSIDriver"},
			{"csinodes", "CSINode"},
			{"storageclasses", "sc", "StorageClass"},
			{"volumeattachments", "VolumeAttachment"},
		}
	}

	return resources
}

func resourceName(name string) (string, error) {
	lcaseName := strings.ToLower(name)
	list := resourceList()

	for _, resource := range list {
		for _, alias := range resource {
			if strings.ToLower(alias) == lcaseName {
				// The last item is the resource name
				return resource[len(resource)-1], nil
			}
		}
	}

	return "", errors.New("resource alias not found")
}
