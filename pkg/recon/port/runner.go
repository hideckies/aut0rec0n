package port

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

type Config struct {
	Host string
}

type port struct {
	id      int
	proto   string
	service string
	status  string
}

type Result struct {
	Ports []port
}

type Runner struct {
	Config Config
	Result Result
}

// Intialize a new Port
func NewRunner(host string) Runner {
	var r Runner
	r.Config = Config{Host: host}
	r.Result = Result{}
	return r
}

// Execute scanning port
func (r *Runner) Run() error {
	r.portScan()
	r.Print()
	return nil
}

// Port scan with nmap
func (r *Runner) portScan() {
	cmd := exec.Command("sudo", "nmap", "-sS", "-p-", r.Config.Host)
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
				r.Result.Ports = append(r.Result.Ports, newPort)
			}
		}
	} else {
		color.Yellow("nmap could not be executed.\naut0rec0n tries a custom scanner.")
		r.customScan()
	}
}

// Custom port scan
func (r *Runner) customScan() {
	maxPort := 65535
	bar := output.NewProgressBar(maxPort, "scanning...")

	check := func(id int, proto string) bool {
		addr := fmt.Sprintf("%s:%d", r.Config.Host, id)
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
			r.Result.Ports = append(r.Result.Ports, newPort)
		}
		time.Sleep(100 * time.Microsecond)
	}
}

// Print result
func (r *Runner) Print() {
	output.Headline("PORT SCAN")
	if len(r.Result.Ports) > 0 {
		w := tabwriter.NewWriter(os.Stdout, 0, 1, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w,
			"%s/%s\t%s\t%s\n",
			color.CyanString("PORT"),
			color.CyanString("PROTO"),
			color.CyanString("STATUS"),
			color.CyanString("SERVICE"))
		for _, port := range r.Result.Ports {
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
