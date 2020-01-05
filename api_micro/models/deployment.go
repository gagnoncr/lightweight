package models

type Deployment struct {
	ServiceName  string			    `json:"serviceName" bson:"serviceName, omitempty"`
	ReplicaCount string             `json:"replicaCount" bson:"replicaCount, omitempty"`
	ImageName    string             `json:"imageName" bson:"imageName, omitempty"`
	Repo         string			    `json:"repo" bson:"repo, omitempty"`
	Tag          string             `json:"tag" bson:"tag, omitempty"`
	PullPolicy   string             `json:"pull_policy" bson:"pull_policy, omitempty"`
	LbType       string			    `json:"lb_type" bson:"lb_type, omitempty"`
	ExternalPort string				`json:"external_port" bson:"external_port, omitempty"`
	InternalPort string			    `json:"internal_port" bson:"internal_port, omitempty"`
}
