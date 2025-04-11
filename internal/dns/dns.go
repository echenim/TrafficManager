package dns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
)

type DNSManager struct {
	Service *route53.Route53
}

func NewDNSManager(zoneID string) *DNSManager {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // or dynamically from config
	}))
	service := route53.New(sess)
	return &DNSManager{
		Service: service,
	}
}

func (m *DNSManager) UpdateRecord(name, recordType, value string) {
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(name),
						Type: aws.String(recordType),
						TTL:  aws.Int64(300),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(value),
							},
						},
					},
				},
			},
		},
		HostedZoneId: aws.String("ZONE_ID"),
	}
	_, err := m.Service.ChangeResourceRecordSets(input)
	if err != nil {
		// logger.Log.WithFields(logrus.Fields{
		// 	"name": name, "type": recordType, "value": value, "error": err,
		// }).Error("Failed to update DNS record")
	}
}
