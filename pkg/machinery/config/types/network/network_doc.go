// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by hack/docgen tool. DO NOT EDIT.

package network

import (
	"net/netip"

	"github.com/siderolabs/talos/pkg/machinery/config/encoder"
)

func (DefaultActionConfigV1Alpha1) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "NetworkDefaultActionConfig",
		Comments:    [3]string{"" /* encoder.HeadComment */, "NetworkDefaultActionConfig is a ingress firewall default action configuration document." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "NetworkDefaultActionConfig is a ingress firewall default action configuration document.",
		Fields: []encoder.Doc{
			{}, {
				Name:        "ingress",
				Type:        "DefaultAction",
				Note:        "",
				Description: "Default action for all not explicitly configured ingress traffic: accept or block.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Default action for all not explicitly configured ingress traffic: accept or block." /* encoder.LineComment */, "" /* encoder.FootComment */},
				Values: []string{
					"accept",
					"block",
				},
			},
		},
	}

	doc.AddExample("", exampleDefaultActionConfigV1Alpha1())

	return doc
}

func (KubespanEndpointsConfigV1Alpha1) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "KubeSpanEndpointsConfig",
		Comments:    [3]string{"" /* encoder.HeadComment */, "KubeSpanEndpointsConfig is a config document to configure KubeSpan endpoints." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "KubeSpanEndpointsConfig is a config document to configure KubeSpan endpoints.",
		Fields: []encoder.Doc{
			{}, {
				Name:        "extraAnnouncedEndpoints",
				Type:        "[]AddrPort",
				Note:        "",
				Description: "A list of extra Wireguard endpoints to announce from this machine.\n\nTalos automatically adds endpoints based on machine addresses, public IP, etc.\nThis field allows to add extra endpoints which are managed outside of Talos, e.g. NAT mapping.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "A list of extra Wireguard endpoints to announce from this machine." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
		},
	}

	doc.AddExample("", exampleKubespanEndpointsV1Alpha1())

	return doc
}

func (RuleConfigV1Alpha1) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "NetworkRuleConfig",
		Comments:    [3]string{"" /* encoder.HeadComment */, "NetworkRuleConfig is a network firewall rule config document." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "NetworkRuleConfig is a network firewall rule config document.",
		Fields: []encoder.Doc{
			{},
			{
				Name:        "name",
				Type:        "string",
				Note:        "",
				Description: "Name of the config document.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Name of the config document." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
			{
				Name:        "portSelector",
				Type:        "RulePortSelector",
				Note:        "",
				Description: "Port selector defines which ports and protocols on the host are affected by the rule.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Port selector defines which ports and protocols on the host are affected by the rule." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
			{
				Name:        "ingress",
				Type:        "[]IngressRule",
				Note:        "",
				Description: "Ingress defines which source subnets are allowed to access the host ports/protocols defined by the `portSelector`.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Ingress defines which source subnets are allowed to access the host ports/protocols defined by the `portSelector`." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
		},
	}

	doc.AddExample("", exampleRuleConfigV1Alpha1())

	return doc
}

func (RulePortSelector) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "RulePortSelector",
		Comments:    [3]string{"" /* encoder.HeadComment */, "RulePortSelector is a port selector for the network rule." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "RulePortSelector is a port selector for the network rule.",
		AppearsIn: []encoder.Appearance{
			{
				TypeName:  "RuleConfigV1Alpha1",
				FieldName: "portSelector",
			},
		},
		Fields: []encoder.Doc{
			{
				Name:        "ports",
				Type:        "PortRanges",
				Note:        "",
				Description: "Ports defines a list of port ranges or single ports.\nThe port ranges are inclusive, and should not overlap.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Ports defines a list of port ranges or single ports." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
			{
				Name:        "protocol",
				Type:        "Protocol",
				Note:        "",
				Description: "Protocol defines traffic protocol (e.g. TCP or UDP).",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Protocol defines traffic protocol (e.g. TCP or UDP)." /* encoder.LineComment */, "" /* encoder.FootComment */},
				Values: []string{
					"tcp",
					"udp",
					"icmp",
					"icmpv6",
				},
			},
		},
	}

	doc.Fields[0].AddExample("", examplePortRanges1())
	doc.Fields[0].AddExample("", examplePortRanges2())

	return doc
}

func (IngressRule) Doc() *encoder.Doc {
	doc := &encoder.Doc{
		Type:        "IngressRule",
		Comments:    [3]string{"" /* encoder.HeadComment */, "IngressRule is a ingress rule." /* encoder.LineComment */, "" /* encoder.FootComment */},
		Description: "IngressRule is a ingress rule.",
		AppearsIn: []encoder.Appearance{
			{
				TypeName:  "RuleConfigV1Alpha1",
				FieldName: "ingress",
			},
		},
		Fields: []encoder.Doc{
			{
				Name:        "subnet",
				Type:        "Prefix",
				Note:        "",
				Description: "Subnet defines a source subnet.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Subnet defines a source subnet." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
			{
				Name:        "except",
				Type:        "Prefix",
				Note:        "",
				Description: "Except defines a source subnet to exclude from the rule, it gets excluded from the `subnet`.",
				Comments:    [3]string{"" /* encoder.HeadComment */, "Except defines a source subnet to exclude from the rule, it gets excluded from the `subnet`." /* encoder.LineComment */, "" /* encoder.FootComment */},
			},
		},
	}

	doc.Fields[0].AddExample("", netip.MustParsePrefix("10.3.4.0/24"))
	doc.Fields[0].AddExample("", netip.MustParsePrefix("2001:db8::/32"))
	doc.Fields[0].AddExample("", netip.MustParsePrefix("1.3.4.5/32"))

	return doc
}

// GetFileDoc returns documentation for the file network_doc.go.
func GetFileDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "network",
		Description: "Package network provides network machine configuration documents.\n",
		Structs: []*encoder.Doc{
			DefaultActionConfigV1Alpha1{}.Doc(),
			KubespanEndpointsConfigV1Alpha1{}.Doc(),
			RuleConfigV1Alpha1{}.Doc(),
			RulePortSelector{}.Doc(),
			IngressRule{}.Doc(),
		},
	}
}
