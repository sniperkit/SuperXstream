package xstream

import "net/rpc"

type HostInfo struct {
	Hostname string
	Addr     string
	Port     string
}

type Host struct {
	SGeng        ScatterGatherEngine
	Info         HostInfo
	connections  []*rpc.Client
	scatterCount int
}

func CreateHost(sgeng ScatterGatherEngine, config *Config, myPort string) Host {

	var inf HostInfo
	var conns [3]*rpc.Client

	return Host{
		SGeng:       sgeng,
		Info:        inf,
		connections: conns[:],
	}
}

func (h *Host) AppendUin(updates *UpdateList, confim *bool) error {
	//this will append updates to the Update In Buffer

}

func (h *Host) IncScatterCount() {
	//this function will increment the scatter count
	//if scatter scount equals the total number of partitions/hosts
	//Gather will be called on the local host

}