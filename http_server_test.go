package thousandeyes

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetHTTPServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"http-server","interval":300,"url":"https://test.com","protocol":"TCP","ipv6Policy":"USE_AGENT_POLICY","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"followRedirects":1,"sslVersionId":0,"verifyCertificate":1,"useNTLM":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := HTTPServer{
		TestID:                Int64(122621),
		Enabled:               Int(1),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:           String("2020-02-06 15:28:07"),
		SavedEvent:            Int(0),
		TestName:              String("test123"),
		Type:                  String("http-server"),
		Interval:              Int(300),
		URL:                   String("https://test.com"),
		Protocol:              String("TCP"),
		NetworkMeasurements:   Int(1),
		MTUMeasurements:       Int(1),
		BandwidthMeasurements: Int(0),
		BGPMeasurements:       Int(1),
		AlertsEnabled:         Int(1),
		LiveShare:             Int(0),
		HTTPTimeLimit:         Int(5),
		HTTPTargetTime:        Int(1000),
		HTTPVersion:           Int(2),
		FollowRedirects:       Int(1),
		NumPathTraces:         Int(3),
		SSLVersionID:          Int(0),
		VerifyCertificate:     Int(1),
		UseNTLM:               Int(0),
		AuthType:              String("NONE"),
		ContentRegex:          String(""),
		ProbeMode:             String("AUTO"),
		Agents: []Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: []string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		BGPMonitors: []Monitor{
			{
				MonitorID:   64,
				IPAddress:   "2001:240:100:ff::2497:2",
				CountryID:   "JP",
				MonitorName: "Tokyo-3",
				Network:     "IIJ Internet Initiative Japan Inc. (AS 2497)",
				MonitorType: "Public",
			},
		},
		SSLVersion: String("Auto"),
		APILinks: APILinks{
			{
				Href: "https://api.thousandeyes.com/v6/tests/1226221",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/web/http-server/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226221",
				Rel:  "data",
			},
		},
	}

	res, err := client.GetHTTPServer(122621)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_GetHTTPServerJsonError(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07",createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"http-server","interval":300,"url":"https://test.com","protocol":"TCP","ipv6Policy":"USE_AGENT_POLICY","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"followRedirects":1,"sslVersionId":0,"verifyCertificate":1,"useNTLM":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/122621.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		_, _ = w.Write([]byte(out))
	})
	_, err := client.GetHTTPServer(122621)
	assert.Error(t, err)
	assert.EqualError(t, err, "Could not decode JSON response: invalid character 'c' looking for beginning of object key string")
}

func TestClient_CreateHTTPServer(t *testing.T) {
	out := `{"test":[{"createdDate":"2020-02-06 15:28:07","createdBy":"William Fleming (wfleming@grumpysysadm.com)","enabled":1,"savedEvent":0,"testId":122621,"testName":"test123","type":"http-server","interval":300,"url":"https://test.com","protocol":"TCP","ipv6Policy":"USE_AGENT_POLICY","networkMeasurements":1,"mtuMeasurements":1,"bandwidthMeasurements":0,"bgpMeasurements":1,"usePublicBGP":1,"alertsEnabled":1,"liveShare":0,"httpTimeLimit":5,"httpTargetTime":1000,"httpVersion":2,"followRedirects":1,"sslVersionId":0,"verifyCertificate":1,"useNTLM":0,"authType":"NONE","contentRegex":"","probeMode":"AUTO","agents":[{"agentId":48620,"agentName":"Seattle, WA (Trial) - IPv6","agentType":"Cloud","countryId":"US","ipAddresses":["135.84.184.153"],"location":"Seattle Area","network":"Astute Hosting Inc. (AS 54527)","prefix":"135.84.184.0/22"}],"sharedWithAccounts":[{"aid":176592,"name":"Cloudreach"}],"bgpMonitors":[{"monitorId":64,"ipAddress":"2001:240:100:ff::2497:2","countryId":"JP","monitorName":"Tokyo-3","network":"IIJ Internet Initiative Japan Inc. (AS 2497)","monitorType":"Public"}],"numPathTraces":3,"apiLinks":[{"rel":"self","href":"https://api.thousandeyes.com/v6/tests/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/web/http-server/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/metrics/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/path-vis/1226221"},{"rel":"data","href":"https://api.thousandeyes.com/v6/net/bgp-metrics/1226221"}],"sslVersion":"Auto"}]}`
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/http-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(out))
	})

	// Define expected values from the API (based on the JSON we print out above)
	expected := HTTPServer{
		TestID:                Int64(122621),
		Enabled:               Int(1),
		CreatedBy:             String("William Fleming (wfleming@grumpysysadm.com)"),
		CreatedDate:           String("2020-02-06 15:28:07"),
		SavedEvent:            Int(0),
		TestName:              String("test123"),
		Type:                  String("http-server"),
		Interval:              Int(300),
		URL:                   String("https://test.com"),
		Protocol:              String("TCP"),
		NetworkMeasurements:   Int(1),
		MTUMeasurements:       Int(1),
		BandwidthMeasurements: Int(0),
		BGPMeasurements:       Int(1),
		AlertsEnabled:         Int(1),
		LiveShare:             Int(0),
		HTTPTimeLimit:         Int(5),
		HTTPTargetTime:        Int(1000),
		HTTPVersion:           Int(2),
		FollowRedirects:       Int(1),
		NumPathTraces:         Int(3),
		SSLVersionID:          Int(0),
		VerifyCertificate:     Int(1),
		UseNTLM:               Int(0),
		AuthType:              String("NONE"),
		ContentRegex:          String(""),
		ProbeMode:             String("AUTO"),
		Agents: []Agent{
			{
				AgentID:     Int(48620),
				AgentType:   String("Cloud"),
				AgentName:   String("Seattle, WA (Trial) - IPv6"),
				CountryID:   String("US"),
				IPAddresses: []string{"135.84.184.153"},
				Location:    String("Seattle Area"),
				Network:     String("Astute Hosting Inc. (AS 54527)"),
				Prefix:      String("135.84.184.0/22"),
			},
		},
		SharedWithAccounts: []SharedWithAccount{
			{
				AID:              Int(176592),
				AccountGroupName: String("Cloudreach"),
			},
		},
		BGPMonitors: []Monitor{
			{
				MonitorID:   64,
				IPAddress:   "2001:240:100:ff::2497:2",
				CountryID:   "JP",
				MonitorName: "Tokyo-3",
				Network:     "IIJ Internet Initiative Japan Inc. (AS 2497)",
				MonitorType: "Public",
			},
		},
		SSLVersion: String("Auto"),
		APILinks: APILinks{
			{
				Href: "https://api.thousandeyes.com/v6/tests/1226221",
				Rel:  "self",
			},
			{
				Href: "https://api.thousandeyes.com/v6/web/http-server/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/metrics/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/path-vis/1226221",
				Rel:  "data",
			},
			{
				Href: "https://api.thousandeyes.com/v6/net/bgp-metrics/1226221",
				Rel:  "data",
			},
		},
	}
	create := HTTPServer{
		TestName: String("test1"),
		URL:      String("https://test.com"),
		Interval: Int(300),
	}
	res, err := client.CreateHTTPServer(create)
	teardown()
	assert.Nil(t, err)
	assert.Equal(t, &expected, res)
}

func TestClient_DeleteHTTPServer(t *testing.T) {
	setup()

	mux.HandleFunc("/tests/http-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
		assert.Equal(t, "POST", r.Method)
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	err := client.DeleteHTTPServer(id)

	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateHTTPServer(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"http-server","url":"https://test.com"}]}`
	mux.HandleFunc("/tests/http-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		_, _ = w.Write([]byte(out))
	})

	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	id := 1
	httpS := HTTPServer{URL: String("https://test.com")}
	res, err := client.UpdateHTTPServer(id, httpS)
	if err != nil {
		t.Fatal(err)
	}
	expected := HTTPServer{TestID: Int64(1), TestName: String("test123"), Type: String("http-server"), URL: String("https://test.com")}
	assert.Equal(t, &expected, res)

}

func TestHTTPServer_AddAgent(t *testing.T) {
	test := HTTPServer{TestName: String("test"), Agents: Agents{}}
	expected := HTTPServer{TestName: String("test"), Agents: []Agent{{AgentID: Int(1)}}}
	test.AddAgent(Int(1))
	assert.Equal(t, expected, test)
}

func TestClient_GetHTTPServerError(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/http-server/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
	})

	_, err := client.GetHTTPServer(1)
	teardown()
	assert.Error(t, err)
}

func TestClient_GetHTTPServerStatusCode(t *testing.T) {
	setup()
	out := `{"test":[{"testId":1,"testName":"test123","type":"http-server","url":"https://test.com"}]}`
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/1.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(out))
	})

	_, err := client.GetHTTPServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_CreateHTTPServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/http-server/new.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.CreateHTTPServer(HTTPServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_UpdateHTTPServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/http-server/1/update.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	_, err := client.UpdateHTTPServer(1, HTTPServer{})
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}

func TestClient_DeleteHTTPServerStatusCode(t *testing.T) {
	setup()
	var client = &Client{APIEndpoint: server.URL, AuthToken: "foo"}
	mux.HandleFunc("/tests/http-server/1/delete.json", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{}`))
	})
	err := client.DeleteHTTPServer(1)
	teardown()
	assert.EqualError(t, err, "Failed call API endpoint. HTTP response code: 400. Error: &{}")
}
