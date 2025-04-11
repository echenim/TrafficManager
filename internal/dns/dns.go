package dns

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/sirupsen/logrus"
)

// DNSManager defines the structure for managing DNS operations
type DNSManager struct {
	service *route53.Route53
	zoneID  string
}

// NewDNSManager initializes a new DNS manager instance
func NewDNSManager(zoneID string) *DNSManager {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Specify the AWS region as needed
	}))
	service := route53.New(sess)
	return &DNSManager{
		service: service,
		zoneID:  zoneID,
	}
}

// UpdateRecord sets or updates a DNS record
func (d *DNSManager) UpdateRecord(recordName, recordType, recordValue string) error {
	input := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(d.zoneID),
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(recordName),
						Type: aws.String(recordType),
						TTL:  aws.Int64(300),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(recordValue),
							},
						},
					},
				},
			},
		},
	}

	_, err := d.service.ChangeResourceRecordSets(input)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"zoneID":      d.zoneID,
			"recordName":  recordName,
			"recordType":  recordType,
			"recordValue": recordValue,
			"error":       err,
		}).Error("Failed to update DNS records")
		return err
	}

	logrus.WithFields(logrus.Fields{
		"zoneID":      d.zoneID,
		"recordName":  recordName,
		"recordType":  recordType,
		"recordValue": recordValue,
	}).Info("DNS record updated successfully")
	return nil
}
