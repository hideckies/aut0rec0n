package recon

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/hideckies/aut0rec0n/pkg/output"

	"github.com/fatih/color"
)

type PortConfig struct {
	Host string
}

type port struct {
	id      int
	proto   string
	service string
	status  string
}

type PortResult struct {
	Ports []port
}

type Port struct {
	Config PortConfig
	Result PortResult
}

// Intialize a new Port
func NewPort(host string) Port {
	var p Port
	p.Config = PortConfig{Host: host}
	p.Result = PortResult{}
	return p
}

// Execute scanning port
func (p *Port) Execute() error {
	p.portScan()
	p.Print()
	return nil
}

// Port scan with nmap
func (p *Port) portScan() {
	cmd := exec.Command("sudo", "nmap", "-sS", "-p-", p.Config.Host)
	result, err := cmd.CombinedOutput()
	if err == nil {
		reader := bytes.NewReader(result)
		scanner := bufio.NewScanner(reader)

		re := regexp.MustCompile(`\d+\/[a-z]+`)

		for scanner.Scan() {
			line := scanner.Text()
			if re.MatchString(line) {
				sep := strings.Split(line, " ")
				idProto := strings.Split(sep[0], "/")

				id, err := strconv.Atoi(idProto[0])
				if err != nil {
					continue
				}
				newPort := port{
					id:      id,
					proto:   idProto[1],
					service: sep[3],
					status:  sep[1],
				}
				p.Result.Ports = append(p.Result.Ports, newPort)
			}
		}
	} else {
		color.Yellow("nmap could not be executed.\naut0rec0n tries a custom scanner.")
		p.customScan()
	}
}

// Custom port scan
func (p *Port) customScan() {
	maxPort := 65535
	bar := output.NewProgressBar(maxPort, "scanning...")

	check := func(id int, proto string) bool {
		addr := fmt.Sprintf("%s:%d", p.Config.Host, id)
		conn, err := net.Dial(proto, addr)
		if err != nil {
			return false
		}
		conn.Close()
		return true
	}

	for i := 1; i <= maxPort; i++ {
		bar.Add(1)
		if check(i, "tcp") {
			newPort := port{
				id:      i,
				proto:   "tcp",
				service: "unknown",
				status:  "open",
			}
			p.Result.Ports = append(p.Result.Ports, newPort)
		}
		time.Sleep(100 * time.Microsecond)
	}
}

// Print result
func (p *Port) Print() {
	output.Headline("PORT SCAN")
	if len(p.Result.Ports) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w,
			"%s/%s\t%s\t%s\n",
			color.CyanString("PORT"),
			color.CyanString("PROTO"),
			color.CyanString("STATUS"),
			color.CyanString("SERVICE"))
		for _, port := range p.Result.Ports {
			fmt.Fprintf(w,
				"%s/%s\t%s\t%s\n",
				color.GreenString("%d", port.id),
				color.GreenString(port.proto),
				color.GreenString(port.status),
				color.GreenString(port.service))
		}
		w.Flush()
	}
}
