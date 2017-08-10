package bot

func NewContainerProcess(image string, idx int) (p *Process, err error) {
	return NewProcess("docker", idx, "run", "-m 512m", "--cpus='.5'", "-i", "--rm", image)
}
