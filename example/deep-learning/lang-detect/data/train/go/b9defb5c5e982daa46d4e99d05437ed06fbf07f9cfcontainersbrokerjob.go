package cf_containers_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CfContainersBrokerJob struct {

	/*Broker - Descr: Determines whether dashboard verifies SSL certificates when communicating with Cloud Controller and UAA Default: true
*/
	Broker *Broker `yaml:"broker,omitempty"`

	/*Nats - Descr: Username for broker to connect to NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*Env - Descr: List of comma-separated hosts to skip connecting to the proxy Default: <nil>
*/
	Env *Env `yaml:"env,omitempty"`

	/*Cf - Descr: URL of the CloudFoundry Cloud Controller Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

}