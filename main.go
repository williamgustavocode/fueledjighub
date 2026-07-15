package main
import ("flag";"fmt";"os";"runtime")
var cliName = "task-runner-4a437f"
func main(){verbose:=flag.Bool("verbose",false,"enable verbose output");version:=flag.Bool("version",false,"show version");flag.Parse();if *version{fmt.Printf("%s v1.0.0\n",cliName);os.Exit(0)};fmt.Printf("[%s] GOOS=%s GOARCH=%s\n",cliName,runtime.GOOS,runtime.GOARCH);fmt.Printf("[%s] NumCPU=%d\n",cliName,runtime.NumCPU());if *verbose{fmt.Printf("[%s] Args: %v\n",cliName,flag.Args());fmt.Printf("[%s] PID: %d\n",cliName,os.Getpid())};fmt.Printf("[%s] Done.\n",cliName)}
