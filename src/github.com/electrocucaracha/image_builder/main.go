package main

import (
    "fmt"
    "os"

    "io/ioutil"
    "github.com/akamensky/argparse"
    "gopkg.in/yaml.v2"

    "github.com/electrocucaracha/image_builder/utils"
    "github.com/electrocucaracha/image_builder/cmd"
)

type BuildInfo struct {
    UrlRepo      string `yaml:"url_repo"`
    RelativePath string `yaml:"relative_path"`
    Version      string `yaml:"version"`
    Profile      string `yaml:"profile"`
}

type Image struct {
    Name      string    `yaml:"name"`
    BuildInfo BuildInfo `yaml:"build_info"`
}

type ConfigurationFile struct {
    SourceFolder string  `yaml:"src_folder"`
    BaseUrlRepo  string  `yaml:"base_url_repo"`
    Images       []Image `yaml:"images"`
}

func addCloneCmd(cmds []*cmd.Cmd, repo string, path string) []*cmd.Cmd{
    if _, err := os.Stat(path); os.IsNotExist(err) {
        cloneCmd := cmd.NewCmd("git").WithArgs("clone", repo, path)
        cmds = append(cmds, cloneCmd)
    }
    return cmds
}

func addCheckoutCmd(cmds []*cmd.Cmd, path string, version string) []*cmd.Cmd{
    checkoutCmd := cmd.NewCmd("git").WithArgs("checkout", version)
    checkoutCmd.Dir = path
    return append(cmds, checkoutCmd)
}

func addMvnBuildCmd(cmds []*cmd.Cmd, path string, profile string) []*cmd.Cmd{
    if _, err := os.Stat(path + "/pom.xml"); err == nil {
        mvnCmd := cmd.NewCmd("mvn").WithArgs("package", "docker:build", "-DskipTests=true", "-Dmaven.test.skip=true", "-Dmaven.javadoc.skip=true")
        if httpProxy != "" {
            mvnCmd.WithArg("-Ddocker.buildArg.http_proxy=" + httpProxy)
        }
        if httpsProxy != "" {
            mvnCmd.WithArg("-Ddocker.buildArg.https_proxy=" + httpsProxy)
        }
        if profile != "" {
            mvnCmd.WithArgs("-P", profile)
        }
        mvnCmd.Dir = path
        cmds = append(cmds, mvnCmd)
    }
    return cmds
}

func addDockerBuildCmd(cmds []*cmd.Cmd, path string, profile string) []*cmd.Cmd{
    if _, err := os.Stat(path + "/Dockerfile"); err == nil {
        dockerCmd := cmd.NewCmd("docker").WithArgs("build", "-f", "./Dockerfile")
        if httpProxy != "" {
            dockerCmd.WithArg("--build-arg http_proxy=" + httpProxy)
        }
        if httpsProxy != "" {
            dockerCmd.WithArg("--build-arg https_proxy=" + httpsProxy)
        }
        if profile != "" {
            dockerCmd.WithArgs("-t", profile)
        }
        dockerCmd.WithArg(".")
        dockerCmd.Dir = path
        cmds = append(cmds, dockerCmd)
    }
    return cmds
}

var filename string
var bufferSize int
var numDispatchers int
var httpProxy string
var httpsProxy string

func parseConfigFile() *ConfigurationFile{
    yamlFile, err := ioutil.ReadFile(filename)
    utils.Check(err)
    config := ConfigurationFile{}
    err = yaml.Unmarshal(yamlFile, &config)
    utils.Check(err)
    return &config
}

func parseArgs() {
    parser := argparse.NewParser("image_builder", "ONAP Docker image builder")
    configFile := parser.String("c", "config-file", &argparse.Options{Required: true, Help: "Configuration file"})
    buffer := parser.Int("b", "buffer-size",  &argparse.Options{Default: 3, Help: "Commands buffer size"})
    dispatchers := parser.Int("n", "number-dispatchers",  &argparse.Options{Default: 3, Help: "Number of dispatchers"})
    proxy := parser.String("p", "http-proxy", &argparse.Options{Required: false, Help: "URL HTTP proxy server"})
    secProxy := parser.String("P", "https-proxy", &argparse.Options{Required: false, Help: "URL HTTPS proxy server"})
    err := parser.Parse(os.Args)
    if err != nil {
       fmt.Print(parser.Usage(err))
       os.Exit(0)
    }
    filename = *configFile
    bufferSize = *buffer
    numDispatchers = *dispatchers
    httpProxy = *proxy
    httpsProxy = *secProxy
}

func main() {
    parseArgs()
    config := parseConfigFile()

    cmdQueue := make(chan []*cmd.Cmd, bufferSize)
    for i := 0; i < numDispatchers; i++ {
        dispatcher := cmd.NewDispatcher(i, cmdQueue)
        defer dispatcher.Stop()
    }

    for _, image := range config.Images {
        var cmds []*cmd.Cmd
        path := config.SourceFolder + image.Name + "/"

        cmds = addCloneCmd(cmds, config.BaseUrlRepo + image.BuildInfo.UrlRepo, path)
        cmds = addCheckoutCmd(cmds, path, image.BuildInfo.Version)
        cmds = addMvnBuildCmd(cmds, path + image.BuildInfo.RelativePath, image.BuildInfo.Profile)
        cmds = addDockerBuildCmd(cmds, path + image.BuildInfo.RelativePath, image.BuildInfo.Profile)

        go func(cmds []*cmd.Cmd) {
            cmdQueue <- cmds
        } (cmds)
    }
}
