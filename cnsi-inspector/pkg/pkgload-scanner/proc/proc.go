package proc

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type ProcTool struct {
	mountPoint string
}

var k8sPatterns []*regexp.Regexp
var k8sPatternsV2 []*regexp.Regexp

const k8sReg = `(?P<id>\w):cpuset:\/kubepods\/(\w+)\/(.+)\/(?P<container>\w+)`
const k8sLowReg = `(?P<id>\w):cpuset:\/kubepods.slice/(.+)/(.+)/docker-(?P<container>\w+).scope`
const k8sCrioReg = `(?P<id>\w):cpuset:\/kubepods.slice/(.+)/(.+)/crio-(?P<container>\w+).scope`
const k8sContainerdReg = `(?P<id>\w):cpuset:\/kubepods.slice/(.+)/(.+)/cri-containerd-(?P<container>\w+).scope`

const k8sLowRegV2 = `(?P<id>\w)::\/(.+)\/(.+)\/docker-(?P<container>\w+).scope`
const k8sCrioRegV2 = `(?P<id>\w)::\/(.+)\/(.+)\/crio-(?P<container>\w+).scope`
const k8sContainerdRegV2 = `(?P<id>\w)::\/kubepods.slice\/(.+)\/(.+)\/cri-containerd-(?P<container>\w+).scope`
const k8sContainerdRegV2MoreFlexible = `cri-containerd-(?P<container>\w+).scope`

// NOTE: example for kind deployment nginx-1.24.0 pid
// 0::/kubelet.slice/kubelet-kubepods.slice/kubelet-kubepods-besteffort.slice/kubelet-kubepods-besteffort-podeacee042_b28f_45dd_9afc_cbed13cf6a49.slice/cri-containerd-1ad60e6e49915933dbd00fd57046e939af9e3808bad383030bd88a1ae6e9d2f7.scope

var procTool *ProcTool

func NewProcTool(opt ...func(*ProcTool)) *ProcTool {
	for _, o := range opt {
		o(procTool)
	}
	return procTool
}

func SetMountPoint(mountPoint string) func(*ProcTool) {
	return func(k *ProcTool) {
		k.mountPoint = mountPoint
	}
}

func init() {

	procTool = &ProcTool{
		mountPoint: "/host-proc",
	}

	k8sPatterns = append(k8sPatterns, regexp.MustCompile(k8sReg))
	k8sPatterns = append(k8sPatterns, regexp.MustCompile(k8sLowReg))
	k8sPatterns = append(k8sPatterns, regexp.MustCompile(k8sCrioReg))
	k8sPatterns = append(k8sPatterns, regexp.MustCompile(k8sContainerdReg))

	k8sPatternsV2 = append(k8sPatternsV2, regexp.MustCompile(k8sLowRegV2))
	k8sPatternsV2 = append(k8sPatternsV2, regexp.MustCompile(k8sCrioRegV2))
	k8sPatternsV2 = append(k8sPatternsV2, regexp.MustCompile(k8sContainerdRegV2))
	k8sPatternsV2 = append(k8sPatternsV2, regexp.MustCompile(k8sContainerdRegV2MoreFlexible))

	if s := os.Getenv("CONTAINER_K8S_REGEXP"); s != "" {
		if e := recover(); e != nil {
			fmt.Printf("k8s regexp.MustCompile panic: %v", e)
		}
		reg := regexp.MustCompile(s)
		k8sPatterns = append(k8sPatterns, reg)
	}
}

func (k *ProcTool) GetContainerFromPID(pid int) (container string, err error) {
	for _, pattern := range k8sPatterns {
		fileName := fmt.Sprintf("%s/%d/cgroup", k.mountPoint, pid)
		s, err := k.getContainerFromFile(fileName, pattern)
		if err == nil && s != "" {
			return s, err
		} else {
			if err != nil {
				//return "", err
				//log.Printf("k8s getContainerFromFile: filename: %s, error :%v", fileName, err)
			}
		}
	}
	for _, patternV2 := range k8sPatternsV2 {
		fileName := fmt.Sprintf("%s/%d/cgroup", k.mountPoint, pid)
		s, err := k.getContainerFromFileV2(fileName, patternV2)
		if err == nil && s != "" {
			return s, err
		} else {
			if err != nil {
				//return "", err
			}
		}
	}

	return "", fmt.Errorf("%w: %d", ErrNoContainerFoundForPID, pid)
}

var ErrNoContainerFoundForPID = errors.New("no container found for pid")

func (k *ProcTool) getContainerFromFile(fileName string, pattern *regexp.Regexp) (container string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err = scanner.Err(); err != nil {
		return "", err
	}

	content, flag := "", false
	for scanner.Scan() {
		m := scanner.Text()
		if strings.Contains(m, "cpuset") {
			flag = true
			content = m
			break
		}
		flag = true
		content = m
	}

	if flag {
		template := []byte("$container")
		var result []byte

		// For each match of the regex in the content.
		for _, submatches := range pattern.FindAllSubmatchIndex([]byte(content), -1) {
			// Apply the captured submatches to the template and append the output
			// to the result.
			result = pattern.Expand(result, template, []byte(content), submatches)
		}

		s := string(result)

		return s, nil
	}

	return "", fmt.Errorf("no container id found")
}

func (k *ProcTool) getContainerFromFileV2(fileName string, pattern *regexp.Regexp) (container string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err = scanner.Err(); err != nil {
		return "", nil
	}

	content, flag := "", false
	for scanner.Scan() {
		m := scanner.Text()
		//if strings.Contains(m, "cpuset") {
		//	flag = true
		//	content = m
		//	break
		//}
		flag = true
		content = m
	}

	if flag {
		template := []byte("$container")
		var result []byte

		// For each match of the regex in the content.
		for _, submatches := range pattern.FindAllSubmatchIndex([]byte(content), -1) {
			// Apply the captured submatches to the template and append the output
			// to the result.
			result = pattern.Expand(result, template, []byte(content), submatches)
		}

		s := string(result)

		return s, nil
	}

	return "", fmt.Errorf("no container id found")
}
