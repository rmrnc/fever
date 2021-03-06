package types

// DCSO FEVER
// Copyright (c) 2017, 2019, DCSO GmbH

import (
	"encoding/json"
	"strconv"
	"time"
)

const (
	// SuricataTimestampFormat is a Go time formatting string describing the
	// timestamp format used by Suricata's EVE JSON output.
	SuricataTimestampFormat = "2006-01-02T15:04:05.999999-0700"

	// EventTypeFlow is the EventType string for the flow type.
	EventTypeFlow = "flow"
	// EventTypeAlert is the EventType string for the alert type.
	EventTypeAlert = "alert"
)

type suriTime struct{ time.Time }

func (t *suriTime) UnmarshalJSON(b []byte) error {
	data, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	t.Time, err = time.Parse(SuricataTimestampFormat, data)
	return err
}

func (t *suriTime) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.Time.Format(SuricataTimestampFormat) + "\""), nil
}

// AlertEvent is am alert sub-object of an EVE entry.
type AlertEvent struct {
	Action      string `json:"action"`
	Gid         int    `json:"gid"`
	SignatureID int    `json:"signature_id"`
	Rev         int    `json:"rev"`
	Signature   string `json:"signature"`
	Category    string `json:"category"`
	Severity    int    `json:"severity"`
}

// DNSEvent is a DNS sub-object of an EVE entry.
type DNSEvent struct {
	Type   string `json:"type"`
	ID     int    `json:"id"`
	Rcode  string `json:"rcode"`
	Rrname string `json:"rrname"`
	Rrtype string `json:"rrtype"`
	TTL    int    `json:"ttl"`
	Rdata  string `json:"rdata"`
	TxID   int    `json:"tx_id"`
}

// HTTPEvent is an HTTP sub-object of an EVE entry.
type HTTPEvent struct {
	Hostname        string `json:"hostname"`
	URL             string `json:"url"`
	HTTPUserAgent   string `json:"http_user_agent"`
	HTTPContentType string `json:"http_content_type"`
	HTTPMethod      string `json:"http_method"`
	Protocol        string `json:"protocol"`
	Status          int    `json:"status"`
	Length          int    `json:"length"`
}

type fileinfoEvent struct {
	Filename string `json:"filename"`
	Magic    string `json:"magic"`
	State    string `json:"state"`
	Md5      string `json:"md5"`
	Stored   bool   `json:"stored"`
	Size     int    `json:"size"`
	TxID     int    `json:"tx_id"`
}

type flowEvent struct {
	PktsToserver  int       `json:"pkts_toserver"`
	PktsToclient  int       `json:"pkts_toclient"`
	BytesToserver int       `json:"bytes_toserver"`
	BytesToclient int       `json:"bytes_toclient"`
	Start         *suriTime `json:"start"`
	End           *suriTime `json:"end"`
	Age           int       `json:"age"`
	State         string    `json:"state"`
	Reason        string    `json:"reason"`
}

// TLSEvent is a TLS sub-object of an EVE entry.
type TLSEvent struct {
	Subject     string `json:"subject"`
	Issuerdn    string `json:"issuerdn"`
	Fingerprint string `json:"fingerprint"`
	Sni         string `json:"sni"`
	Version     string `json:"version"`
}

type statsEvent struct {
	Uptime  int `json:"uptime"`
	Capture struct {
		KernelPackets int `json:"kernel_packets"`
		KernelDrops   int `json:"kernel_drops"`
	} `json:"capture"`
	Decoder struct {
		Pkts       int   `json:"pkts"`
		Bytes      int64 `json:"bytes"`
		Invalid    int   `json:"invalid"`
		Ipv4       int   `json:"ipv4"`
		Ipv6       int   `json:"ipv6"`
		Ethernet   int   `json:"ethernet"`
		Raw        int   `json:"raw"`
		Null       int   `json:"null"`
		Sll        int   `json:"sll"`
		TCP        int   `json:"tcp"`
		UDP        int   `json:"udp"`
		Sctp       int   `json:"sctp"`
		Icmpv4     int   `json:"icmpv4"`
		Icmpv6     int   `json:"icmpv6"`
		Ppp        int   `json:"ppp"`
		Pppoe      int   `json:"pppoe"`
		Gre        int   `json:"gre"`
		Vlan       int   `json:"vlan"`
		VlanQinq   int   `json:"vlan_qinq"`
		Teredo     int   `json:"teredo"`
		Ipv4InIpv6 int   `json:"ipv4_in_ipv6"`
		Ipv6InIpv6 int   `json:"ipv6_in_ipv6"`
		Mpls       int   `json:"mpls"`
		AvgPktSize int   `json:"avg_pkt_size"`
		MaxPktSize int   `json:"max_pkt_size"`
		Erspan     int   `json:"erspan"`
		Ipraw      struct {
			InvalidIPVersion int `json:"invalid_ip_version"`
		} `json:"ipraw"`
		Ltnull struct {
			PktTooSmall     int `json:"pkt_too_small"`
			UnsupportedType int `json:"unsupported_type"`
		} `json:"ltnull"`
		Dce struct {
			PktTooSmall int `json:"pkt_too_small"`
		} `json:"dce"`
	} `json:"decoder"`
	Flow struct {
		Memcap           int `json:"memcap"`
		Spare            int `json:"spare"`
		EmergModeEntered int `json:"emerg_mode_entered"`
		EmergModeOver    int `json:"emerg_mode_over"`
		TCPReuse         int `json:"tcp_reuse"`
		Memuse           int `json:"memuse"`
	} `json:"flow"`
	Defrag struct {
		Ipv4 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv4"`
		Ipv6 struct {
			Fragments   int `json:"fragments"`
			Reassembled int `json:"reassembled"`
			Timeouts    int `json:"timeouts"`
		} `json:"ipv6"`
		MaxFragHits int `json:"max_frag_hits"`
	} `json:"defrag"`
	Stream struct {
		ThreeWhsAckInWrongDir           int `json:"3whs_ack_in_wrong_dir"`
		ThreeWhsAsyncWrongSeq           int `json:"3whs_async_wrong_seq"`
		ThreeWhsRightSeqWrongAckEvasion int `json:"3whs_right_seq_wrong_ack_evasion"`
	} `json:"stream"`
	TCP struct {
		Sessions           int `json:"sessions"`
		SsnMemcapDrop      int `json:"ssn_memcap_drop"`
		Pseudo             int `json:"pseudo"`
		PseudoFailed       int `json:"pseudo_failed"`
		InvalidChecksum    int `json:"invalid_checksum"`
		NoFlow             int `json:"no_flow"`
		Syn                int `json:"syn"`
		Synack             int `json:"synack"`
		Rst                int `json:"rst"`
		SegmentMemcapDrop  int `json:"segment_memcap_drop"`
		StreamDepthReached int `json:"stream_depth_reached"`
		ReassemblyGap      int `json:"reassembly_gap"`
		Memuse             int `json:"memuse"`
		ReassemblyMemuse   int `json:"reassembly_memuse"`
	} `json:"tcp"`
	Detect struct {
		Alert int `json:"alert"`
	} `json:"detect"`
	FlowMgr struct {
		ClosedPruned int `json:"closed_pruned"`
		NewPruned    int `json:"new_pruned"`
		EstPruned    int `json:"est_pruned"`
	} `json:"flow_mgr"`
	DNS struct {
		Memuse       int `json:"memuse"`
		MemcapState  int `json:"memcap_state"`
		MemcapGlobal int `json:"memcap_global"`
	} `json:"dns"`
	HTTP struct {
		Memuse int `json:"memuse"`
		Memcap int `json:"memcap"`
	} `json:"http"`
}

type sshEvent struct {
	Client struct {
		ProtoVersion    string `json:"proto_version"`
		SoftwareVersion string `json:"software_version"`
	} `json:"client"`
	Server struct {
		ProtoVersion    string `json:"proto_version"`
		SoftwareVersion string `json:"software_version"`
	} `json:"server"`
}

type smtpEvent struct {
	Helo     string   `json:"helo"`
	MailFrom string   `json:"mail_from"`
	RcptTo   []string `json:"rcpt_to"`
}

type tcpEvent struct {
	State      string `json:"state"`
	Syn        bool   `json:"syn"`
	TCPflags   string `json:"tcp_flags"`
	TCPflagsTc string `json:"tcp_flags_tc"`
	TCPflagsTs string `json:"tcp_flags_ts"`
}

type emailEvent struct {
	Status string `json:"status"`
}

type packetInfo struct {
	Linktype int `json:"linktype"`
}

// ExtraInfo contains non-EVE-standard extra information
type ExtraInfo struct {
	BloomIOC string `json:"bloom-ioc,omitempty"`
}

// EveEvent is the huge struct which can contain a parsed suricata eve.json
// log event.
type EveEvent struct {
	Timestamp        *suriTime      `json:"timestamp"`
	EventType        string         `json:"event_type"`
	FlowID           int64          `json:"flow_id,omitempty"`
	InIface          string         `json:"in_iface,omitempty"`
	SrcIP            string         `json:"src_ip,omitempty"`
	SrcPort          int            `json:"src_port,omitempty"`
	SrcHost          []string       `json:"src_host,omitempty"`
	DestIP           string         `json:"dest_ip,omitempty"`
	DestPort         int            `json:"dest_port,omitempty"`
	DestHost         []string       `json:"dest_host,omitempty"`
	Proto            string         `json:"proto,omitempty"`
	AppProto         string         `json:"app_proto,omitempty"`
	TxID             int            `json:"tx_id,omitempty"`
	TCP              *tcpEvent      `json:"tcp,omitempty"`
	PacketInfo       *packetInfo    `json:"packet_info,omitempty"`
	Alert            *AlertEvent    `json:"alert,omitempty"`
	Payload          string         `json:"payload,omitempty"`
	PayloadPrintable string         `json:"payload_printable,omitempty"`
	Stream           int            `json:"stream,omitempty"`
	Packet           string         `json:"packet,omitempty"`
	SMTP             *smtpEvent     `json:"smtp,omitempty"`
	Email            *emailEvent    `json:"email,omitempty"`
	DNS              *DNSEvent      `json:"dns,omitempty"`
	HTTP             *HTTPEvent     `json:"http,omitempty"`
	Fileinfo         *fileinfoEvent `json:"fileinfo,omitempty"`
	Flow             *flowEvent     `json:"flow,omitempty"`
	SSH              *sshEvent      `json:"ssh,omitempty"`
	TLS              *TLSEvent      `json:"tls,omitempty"`
	Stats            *statsEvent    `json:"stats,omitempty"`
	ExtraInfo        *ExtraInfo     `json:"_extra,omitempty"`
}

// EveOutEvent is the version of EveEvent that we use to marshal the output for
// downstream consumption.
type EveOutEvent EveEvent

// MarshalJSON for EveOutEvents ensures that FlowIDs are represented in JSON
// as a string. This is necessary to work around some arbitrary limitations such
// as syslog-ng's funny JSON parser implementation, which truncates large
// integers found in JSON values.
func (e EveOutEvent) MarshalJSON() ([]byte, error) {
	type Alias EveOutEvent
	v, err := json.Marshal(&struct {
		FlowID string `json:"flow_id"`
		Alias
	}{
		FlowID: strconv.FormatInt(e.FlowID, 10),
		Alias:  (Alias)(e),
	})
	return v, err
}

// UnmarshalJSON implements filling an EveOutEvent from a byte slice, converting
// the string in the FlowID field back into a number. This is necessary to
// ensure that a round-trip (write+read) works.
func (e *EveOutEvent) UnmarshalJSON(d []byte) error {
	type EveOutEvent2 EveOutEvent
	x := struct {
		EveOutEvent2
		FlowID json.Number `json:"flow_id"`
	}{EveOutEvent2: EveOutEvent2(*e)}

	if err := json.Unmarshal(d, &x); err != nil {
		return err
	}
	*e = EveOutEvent(x.EveOutEvent2)
	var err error
	e.FlowID, _ = x.FlowID.Int64() // ignore error; defaulting to zero
	return err
}
