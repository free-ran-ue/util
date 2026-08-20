package util

import (
	"net"

	"github.com/free-ran-ue/free-ran-ue/v2/logger"
	"github.com/free5gc/nas/ie"
)

func GetQosRule(ruleBytes []byte, logger *logger.UeLogger) []string {
	var rules ie.QosRules
	if err := rules.UnmarshalBinary(ruleBytes); err != nil {
		logger.PduLog.Warnf("unmarshal qos rules failed: %+v", err)
		return nil
	}

	qosRules := make([]string, 0)

	for _, r := range rules.Rules {
		for _, p := range r.PktFilterList {
			contents := p.Contents
			if contents.MatchAll {
				continue
			}
			ip, _, err := net.ParseCIDR(contents.RemoteAddr)
			if err != nil || ip.To4() == nil {
				logger.PduLog.Warnf("unsupported qos rule packet filter contents: %+v", contents)
				continue
			}
			qosRules = append(qosRules, contents.RemoteAddr)
		}
	}

	return qosRules
}
