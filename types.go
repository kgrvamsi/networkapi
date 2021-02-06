package networkapi

import "encoding/xml"

type CommitHistory struct {
	User      string `json:"user"`
	Method    string `json:"method"`
	Log       string `json:"log"`
	Comment   string `json:"comment"`
	Timestamp string `json:"timestamp"`
}

type InterfacesInfo struct {
	XMLName              xml.Name              `xml:"rpc-reply"`
	InterfaceInformation InterfaceInformations `xml:"interface-information"`
}

type InterfaceInformations struct {
	XMLName           xml.Name            `xml:"interface-information"`
	PhysicalInterface []PhysicalInterface `xml:"physical-interface"`
	LogicalInterface  []LogicalInterface  `xml:"logical-interface"`
}

type PhysicalInterface struct {
	XMLName     xml.Name `xml:"physical-interface"`
	Name        string   `xml:"name"`
	Adminstatus string   `xml:"admin-status"`
	Operstatus  string   `xml:"oper-status"`
	Description string   `xml:"description"`
}

type LogicalInterface struct {
	XMLName     xml.Name `xml:"logical-interface"`
	Name        string   `xml:"name"`
	Adminstatus string   `xml:"admin-status"`
	Operstatus  string   `xml:"oper-status"`
	Description string   `xml:"description"`
}

type InterfacesList struct {
	Interfacename string `json:"interfacename"`
	Adminstatus   string `json:"adminstatus"`
	Operstatus    string `json:"operstatus"`
	Description   string `json:"description"`
}

type RPCReplyBgp struct {
	XMLName        xml.Name       `xml:"rpc-reply"`
	Bgpinformation BGPInformation `xml:"bgp-information"`
}

type BGPInformation struct {
	XMLName               xml.Name  `xml:"bgp-information"`
	Groupcount            string    `xml:"group-count"`
	Peercount             string    `xml:"peer-count"`
	Downpeercount         string    `xml:"down-peer-count"`
	Unconfiguredpeercount string    `xml:"unconfigured-peer-count"`
	Bgppeer               []Bgppeer `xml:"bgp-peer"`
}

type Bgppeer struct {
	XMLName         xml.Name `xml:"bgp-peer"`
	Peeraddress     string   `xml:"peer-address"`
	Peeras          string   `xml:"peer-as"`
	Inputmessages   string   `xml:"input-messages"`
	Outputmessages  string   `xml:"output-messages"`
	Routequeuecount string   `xml:"route-queue-count"`
	Flapcount       string   `xml:"flap-count"`
	Elapsedtime     string   `xml:"elapsed-time"`
	Peerstate       string   `xml:"peer-state"`
}

type Bgppeers struct {
	Peeraddress string
	Peeras      string
	Peerstate   string
}

type RouterTime struct {
	MultiRoutingEngineResults []struct {
		MultiRoutingEngineItem []struct {
			Rename []struct {
				Data string `json:"data"`
			} `json:"re-name"`
			SystemUptimeInformation []struct {
				Attributes struct {
					Xmlns string `json:"xmlns"`
				} `json:"attributes"`
				Currenttime []struct {
					Datetime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
				} `json:"current-time"`
				LastConfiguredTime []struct {
					DateTime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					TimeLength []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
					User []struct {
						Data string `json:"data"`
					} `json:"user"`
				} `json:"last-configured-time"`
				ProtocolsStartedTime []struct {
					DateTime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					TimeLength []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
				} `json:"protocols-started-time"`
				SystemBootedTime []struct {
					DateTime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					TimeLength []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"time-length"`
				} `json:"system-booted-time"`
				TimeSource []struct {
					Data string `json:"data"`
				} `json:"time-source"`
				UptimeInformation []struct {
					ActiveUserCount []struct {
						Attributes struct {
							Junosformat string `json:"junos:format"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"active-user-count"`
					DateTime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"date-time"`
					LoadAverage1 []struct {
						Data string `json:"data"`
					} `json:"load-average-1"`
					LoadAverage15 []struct {
						Data string `json:"data"`
					} `json:"load-average-15"`
					LoadAverage5 []struct {
						Data string `json:"data"`
					} `json:"load-average-5"`
					UpTime []struct {
						Attributes struct {
							Junosseconds string `json:"junos:seconds"`
						} `json:"attributes"`
						Data string `json:"data"`
					} `json:"up-time"`
					UserTable []struct{} `json:"user-table"`
				} `json:"uptime-information"`
			} `json:"system-uptime-information"`
		} `json:"multi-routing-engine-item"`
	} `json:"multi-routing-engine-results"`
}

type LLDpNeighborsSSH struct {
	Lldp_neighbors_information []struct {
		Attributes struct {
			Junos_style string `json:"junos:style"`
		} `json:"attributes"`
		Lldp_neighbor_information []struct {
			Lldp_local_parent_interface_name []struct {
				Data string `json:"data"`
			} `json:"lldp-local-parent-interface-name"`
			Lldp_local_port_id []struct {
				Data string `json:"data"`
			} `json:"lldp-local-port-id"`
			Lldp_remote_chassis_id []struct {
				Data string `json:"data"`
			} `json:"lldp-remote-chassis-id"`
			Lldp_remote_chassis_id_subtype []struct {
				Data string `json:"data"`
			} `json:"lldp-remote-chassis-id-subtype"`
			Lldp_remote_port_description []struct {
				Data string `json:"data"`
			} `json:"lldp-remote-port-description"`
			Lldp_remote_system_name []struct {
				Data string `json:"data"`
			} `json:"lldp-remote-system-name"`
		} `json:"lldp-neighbor-information"`
	} `json:"lldp-neighbors-information"`
}

type InterfaceDescriptions struct {
	Interface_information []struct {
		Attributes struct {
			Junos_style string `json:"junos:style"`
			Xmlns       string `json:"xmlns"`
		} `json:"attributes"`
		Logical_interface []struct {
			Admin_status []struct {
				Data string `json:"data"`
			} `json:"admin-status"`
			Description []struct {
				Data string `json:"data"`
			} `json:"description"`
			Name []struct {
				Data string `json:"data"`
			} `json:"name"`
			Oper_status []struct {
				Data string `json:"data"`
			} `json:"oper-status"`
		} `json:"logical-interface"`
		Physical_interface []struct {
			Admin_status []struct {
				Data string `json:"data"`
			} `json:"admin-status"`
			Description []struct {
				Data string `json:"data"`
			} `json:"description"`
			Name []struct {
				Data string `json:"data"`
			} `json:"name"`
			Oper_status []struct {
				Data string `json:"data"`
			} `json:"oper-status"`
		} `json:"physical-interface"`
	} `json:"interface-information"`
}

type InterfacesDiagnosticsSSH struct {
	XMLName              xml.Name `xml:"rpc-reply"`
	Junos                string   `xml:"junos,attr"`
	InterfaceInformation struct {
		Xmlns             string `xml:"xmlns,attr"`
		Style             string `xml:"style,attr"`
		PhysicalInterface []struct {
			Name              string `xml:"name"`
			OpticsDiagnostics struct {
				ModuleTemperature struct {
					Celsius string `xml:"celsius,attr"`
				} `xml:"module-temperature"`
				ModuleVoltage                       string `xml:"module-voltage"`
				ModuleTemperatureHighAlarm          string `xml:"module-temperature-high-alarm"`
				ModuleTemperatureLowAlarm           string `xml:"module-temperature-low-alarm"`
				ModuleTemperatureHighWarn           string `xml:"module-temperature-high-warn"`
				ModuleTemperatureLowWarn            string `xml:"module-temperature-low-warn"`
				ModuleVoltageHighAlarm              string `xml:"module-voltage-high-alarm"`
				ModuleVoltageLowAlarm               string `xml:"module-voltage-low-alarm"`
				ModuleVoltageHighWarn               string `xml:"module-voltage-high-warn"`
				ModuleVoltageLowWarn                string `xml:"module-voltage-low-warn"`
				ModuleTemperatureHighAlarmThreshold struct {
					Celsius string `xml:"celsius,attr"`
				} `xml:"module-temperature-high-alarm-threshold"`
				ModuleTemperatureLowAlarmThreshold struct {
					Celsius string `xml:"celsius,attr"`
				} `xml:"module-temperature-low-alarm-threshold"`
				ModuleTemperatureHighWarnThreshold struct {
					Celsius string `xml:"celsius,attr"`
				} `xml:"module-temperature-high-warn-threshold"`
				ModuleTemperatureLowWarnThreshold struct {
					Celsius string `xml:"celsius,attr"`
				} `xml:"module-temperature-low-warn-threshold"`
				ModuleVoltageHighAlarmThreshold    string `xml:"module-voltage-high-alarm-threshold"`
				ModuleVoltageLowAlarmThreshold     string `xml:"module-voltage-low-alarm-threshold"`
				ModuleVoltageHighWarnThreshold     string `xml:"module-voltage-high-warn-threshold"`
				ModuleVoltageLowWarnThreshold      string `xml:"module-voltage-low-warn-threshold"`
				LaserBiasCurrentHighAlarmThreshold string `xml:"laser-bias-current-high-alarm-threshold"`
				LaserBiasCurrentLowAlarmThreshold  string `xml:"laser-bias-current-low-alarm-threshold"`
				LaserBiasCurrentHighWarnThreshold  string `xml:"laser-bias-current-high-warn-threshold"`
				LaserBiasCurrentLowWarnThreshold   string `xml:"laser-bias-current-low-warn-threshold"`
				LaserTxPowerHighAlarmThreshold     string `xml:"laser-tx-power-high-alarm-threshold"`
				LaserTxPowerHighAlarmThresholdDbm  string `xml:"laser-tx-power-high-alarm-threshold-dbm"`
				LaserTxPowerLowAlarmThreshold      string `xml:"laser-tx-power-low-alarm-threshold"`
				LaserTxPowerLowAlarmThresholdDbm   string `xml:"laser-tx-power-low-alarm-threshold-dbm"`
				LaserTxPowerHighWarnThreshold      string `xml:"laser-tx-power-high-warn-threshold"`
				LaserTxPowerHighWarnThresholdDbm   string `xml:"laser-tx-power-high-warn-threshold-dbm"`
				LaserTxPowerLowWarnThreshold       string `xml:"laser-tx-power-low-warn-threshold"`
				LaserTxPowerLowWarnThresholdDbm    string `xml:"laser-tx-power-low-warn-threshold-dbm"`
				LaserRxPowerHighAlarmThreshold     string `xml:"laser-rx-power-high-alarm-threshold"`
				LaserRxPowerHighAlarmThresholdDbm  string `xml:"laser-rx-power-high-alarm-threshold-dbm"`
				LaserRxPowerLowAlarmThreshold      string `xml:"laser-rx-power-low-alarm-threshold"`
				LaserRxPowerLowAlarmThresholdDbm   string `xml:"laser-rx-power-low-alarm-threshold-dbm"`
				LaserRxPowerHighWarnThreshold      string `xml:"laser-rx-power-high-warn-threshold"`
				LaserRxPowerHighWarnThresholdDbm   string `xml:"laser-rx-power-high-warn-threshold-dbm"`
				LaserRxPowerLowWarnThreshold       string `xml:"laser-rx-power-low-warn-threshold"`
				LaserRxPowerLowWarnThresholdDbm    string `xml:"laser-rx-power-low-warn-threshold-dbm"`
				OpticsDiagnosticsLaneValues        struct {
					LaneIndex                        string `xml:"lane-index"`
					LaserBiasCurrent                 string `xml:"laser-bias-current"`
					LaserOutputPower                 string `xml:"laser-output-power"`
					LaserOutputPowerDbm              string `xml:"laser-output-power-dbm"`
					LaserRxOpticalPower              string `xml:"laser-rx-optical-power"`
					LaserRxOpticalPowerDbm           string `xml:"laser-rx-optical-power-dbm"`
					LaserBiasCurrentHighAlarm        string `xml:"laser-bias-current-high-alarm"`
					LaserBiasCurrentLowAlarm         string `xml:"laser-bias-current-low-alarm"`
					LaserBiasCurrentHighWarn         string `xml:"laser-bias-current-high-warn"`
					LaserBiasCurrentLowWarn          string `xml:"laser-bias-current-low-warn"`
					LaserRxPowerHighAlarm            string `xml:"laser-rx-power-high-alarm"`
					LaserRxPowerLowAlarm             string `xml:"laser-rx-power-low-alarm"`
					LaserRxPowerHighWarn             string `xml:"laser-rx-power-high-warn"`
					LaserRxPowerLowWarn              string `xml:"laser-rx-power-low-warn"`
					TxLossOfSignalFunctionalityAlarm string `xml:"tx-loss-of-signal-functionality-alarm"`
					RxLossOfSignalAlarm              string `xml:"rx-loss-of-signal-alarm"`
					TxLaserDisabledAlarm             string `xml:"tx-laser-disabled-alarm"`
				} `xml:"optics-diagnostics-lane-values"`
			} `xml:"optics-diagnostics"`
		} `xml:"physical-interface"`
	} `xml:"interface-information"`
}

type RouterTimeRes struct {
	Currenttime        string `json:"current_time"`
	LastConfiguredTime string `json:"last_configured_time"`
	SystemBootedTime   string `json:"system_booted_time"`
}
