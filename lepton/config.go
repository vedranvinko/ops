package lepton

// Build configs
type Config struct {
	Args         []string
	Dirs         []string
	Files        []string
	MapDirs      map[string]string
	Env          map[string]string
	Debugflags   []string
	NoTrace      []string
	Program      string
	Version      string
	Boot         string
	Kernel       string
	Mkfs         string
	NameServer   string
	NightlyBuild bool
	RunConfig    RunConfig
	CloudConfig  ProviderConfig
	Force        bool
	TargetRoot   string
	ManifestName string // save manifest to
}

type ProviderConfig struct {
	Platform   string `cloud:"platfom"`
	ProjectID  string `cloud:"projectid"`
	Zone       string `cloud:"zone"`
	BucketName string `cloud:"bucketname"`
	ImageName  string `cloud:"imagename"`
	Flavor     string `cloud:"flavor"`
}

// Runtime configs
type RunConfig struct {
	Imagename string
	Ports     []int
	Verbose   bool
	Memory    string
	Bridged   bool
	TapName   string
	Accel     bool
}

func RuntimeConfig(image string, ports []int, verbose bool) RunConfig {
	return RunConfig{Imagename: image, Ports: ports, Verbose: verbose, Memory: "2G"}
}
