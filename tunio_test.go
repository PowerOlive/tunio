package tunio

/*

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"
)

const (
	deviceName = "tun0"
	deviceIP   = "10.0.0.2"
	deviceMask = "255.255.255.0"
)

var hostIP string

func init() {
	if os.Getenv("HOST_IP") != "" {
		hostIP = os.Getenv("HOST_IP")
	} else {
		hostIP = "10.0.0.101"
	}
}

const (
	googleHumansTxt = "Google is built by a large team of engineers, designers, researchers, robots, and others in many different sites across the globe. It is updated continuously, and built with more tools and technologies than we can shake a stick at. If you'd like to help us out, see google.com/careers.\n"
)

func TestConfigure(t *testing.T) {
	// This function dials to an external host which will take anything that
	// arrives on port 20443 and redirect it to www.google.com:443
	fn := func(proto, addr string) (net.Conn, error) {
		// The address the client wants to connect is addr, but we ignore it and
		// connect to a fixed known host instead because that's the easiest way to
		// test a connection that is not routed to tun0. We're going to manually
		// set up the external host to connect to www.google.com:443. In
		// VpnService's context this could be achieved by protecting this socket.
		_, port, _ := net.SplitHostPort(addr)
		return net.Dial("tcp", hostIP+":20"+port)
	}
	go func() {
		// Configuring the device and passing the dialer function we want to use.
		if err := Configure(deviceName, deviceIP, deviceMask, fn); err != nil {
			log.Printf("error: %v\n", err)
			//t.Fatal(err)
		}
	}()
	time.Sleep(time.Millisecond * 500)
	log.Printf("Waiting for %q...", deviceName)
}

func dialAndWaitForResponse(uri, expects string) error {
	log.Printf("Getting %s...", uri)

	cli := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:     true,
			ResponseHeaderTimeout: time.Second * 30,
			TLSHandshakeTimeout:   time.Second * 30,
		},
	}

	res, err := cli.Get(uri)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	log.Printf("OK!")

	if string(b) != googleHumansTxt {
		return errors.New(`Expecting a fixed response.`)
	}

	return nil
}

func TestSequenceDialerHTTP(t *testing.T) {
	for i := 0; i < 20; i++ {
		err := dialAndWaitForResponse("http://www.google.com/humans.txt", googleHumansTxt)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestSequenceDialerHTTPS(t *testing.T) {
	for i := 0; i < 20; i++ {
		err := dialAndWaitForResponse("https://www.google.com/humans.txt", googleHumansTxt)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestConcurrentDialerHTTP(t *testing.T) {

	for j := 0; j < 10; j++ {
		var wg sync.WaitGroup

		for i := 0; i < 200; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := dialAndWaitForResponse("http://www.google.com/humans.txt", googleHumansTxt)
				if err != nil {
					log.Printf("dialAndWaitForResponse: %q", err.Error())
				}
			}()
		}

		wg.Wait()
	}
}

func TestConcurrentDialerHTTPS(t *testing.T) {

	for j := 0; j < 10; j++ {
		var wg sync.WaitGroup

		for i := 0; i < 200; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := dialAndWaitForResponse("https://www.google.com/humans.txt", googleHumansTxt)
				if err != nil {
					log.Printf("dialAndWaitForResponse: %q", err.Error())
				}
			}()
		}

		wg.Wait()
	}
}
*/
