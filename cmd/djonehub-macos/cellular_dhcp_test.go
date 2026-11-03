package main

import "testing"

func TestModemBearerReadyForDHCP(t *testing.T) {
	const ecm = "AT+QCFG=\"usbnet\"\r\n+QCFG: \"usbnet\",1\r\nOK"
	tests := []struct {
		name   string
		usbnet string
		pdp    string
		want   bool
	}{
		{
			name:   "ECM with IPv4 and IPv6 bearer (field capture)",
			usbnet: ecm,
			pdp:    "AT+CGPADDR=1\r\n+CGPADDR: 1,\"10.136.41.211,36.8.132.86.48.48.2.242.0.1.0.1.204.92.159.49\"\r\nOK",
			want:   true,
		},
		{
			name:   "ECM without bearer",
			usbnet: ecm,
			pdp:    "+CGPADDR: 1,\"0.0.0.0,0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0\"\r\nOK",
			want:   false,
		},
		{
			name:   "ECM with IPv6-only bearer",
			usbnet: ecm,
			pdp:    "+CGPADDR: 1,\"36.8.133.86.48.48.22.121.0.1.0.1.204.92.160.117\"\r\nOK",
			want:   false,
		},
		{
			name:   "usbnet 0 never serves DHCP",
			usbnet: "+QCFG: \"usbnet\",0\r\nOK",
			pdp:    "+CGPADDR: 1,\"10.136.41.211\"\r\nOK",
			want:   false,
		},
		{
			name:   "usbnet 10 is not ECM",
			usbnet: "+QCFG: \"usbnet\",10\r\nOK",
			pdp:    "+CGPADDR: 1,\"10.136.41.211\"\r\nOK",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := modemBearerReadyForDHCP(tt.usbnet, tt.pdp); got != tt.want {
				t.Fatalf("modemBearerReadyForDHCP() = %v, want %v", got, tt.want)
			}
		})
	}
}
