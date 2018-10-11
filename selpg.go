package main

import (
    "fmt"
	"os"
	"bufio"
	"github.com/spf13/pflag"
	"os/exec"
	"io"
)

type selpg_args struct {
	start_page  int   
	end_page    int    
	in_filename string 
	page_len    int    
	page_type   bool 
	print_dest  string
	
}

var progname string

func Parser(p *selpg_args) {
	pflag.Usage = usage
	pflag.IntVarP(&p.start_page,"start", "s", 0, "首页")
	pflag.IntVarP(&p.end_page,"end","e", 0, "尾页")
	pflag.IntVarP(&p.page_len,"linenum", "l", 15, "打印的每页行数")
	pflag.BoolVarP(&p.page_type,"printdes","f", false, "是否用换页符换页")
	pflag.StringVarP(&p.print_dest, "othertype","d", "", "打印目的地")
	pflag.Parse()
}

func processArgs(ac int, psa *selpg_args) {
	/* check the command-line arguments for validity */
	if ac < 3 {
		fmt.Fprintf(os.Stderr, "%s: 参数太少\n", progname)
		pflag.Usage()
		os.Exit(1)
	}

	/* handle 1st arg - start page */
	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "%s: 第一个参数应该为 -s=start_page\n", progname)
		pflag.Usage()
		os.Exit(2)
	}
	if psa.start_page < 1  {
		fmt.Fprintf(os.Stderr, "%s: 不合法起始页页数 %s\n", progname, psa.start_page)
		pflag.Usage()
		os.Exit(3)
	}

	/* handle 2nd arg - end page */
	if os.Args[3][0] != '-' || os.Args[3][1] != 'e' {
		fmt.Fprintf(os.Stderr, "%s: 第二个参数应该为 -e=end_page\n", progname)
		pflag.Usage()
		os.Exit(4)
	}
	if psa.end_page < 1  || psa.end_page < psa.start_page  {
		fmt.Fprintf(os.Stderr, "%s: 不合法终止页页数 %s\n", progname, psa.end_page)
		pflag.Usage()
		os.Exit(5)
	}
    
	/* now handle optional args */
	if psa.page_len != 15 {
		if psa.page_len < 1  {
			fmt.Fprintf(os.Stderr, "%s: 非法页长度 %s\n", progname, psa.page_len)
			pflag.Usage()
			os.Exit(6)
		}
	}


	/* there is one more arg */
	if pflag.NArg() > 0 {
		psa.in_filename = pflag.Arg(0)
		/* check if file exists */
		file, err := os.Open(psa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: 输入文件 \"%s\" 不存在\n", progname, psa.in_filename)
			os.Exit(7)
		}
		/* check if file is readable */
		file, err = os.OpenFile(psa.in_filename, os.O_RDONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				fmt.Fprintf(os.Stderr, "%s: 输入文件存在 但是不能被读取 \"%s\" exists but cannot be read\n", progname, psa.in_filename)
				os.Exit(8)
			}
		}
		file.Close()
	}
}

func processInput(psa *selpg_args) {
	fin := os.Stdin
	fout := os.Stdout
	var (
		 page_ctr int
		 line_ctr int
		 err error
		 err1 error
		 err2 error
		 line string
		 cmd *exec.Cmd
		 stdin io.WriteCloser
	)
	/* set the input source */
	if psa.in_filename != "" {
		fin, err1 = os.Open(psa.in_filename)
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "%s:  不能打开输入文件 \"%s\"\n", progname, psa.in_filename)
			os.Exit(11)
		}
	}

	if psa.print_dest != "" {
		cmd = exec.Command("cat", "-n")
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}

/* begin one of two main loops based on page type */
	rd := bufio.NewReader(fin)
	if psa.page_type == false {
		line_ctr = 0
		page_ctr = 1
		for true {
			line, err2 = rd.ReadString('\n')
			if err2 != nil { /* error or EOF */
				break
			}
			line_ctr++
			if line_ctr > psa.page_len {
				page_ctr++
				line_ctr = 1
			}
			if page_ctr >= psa.start_page && page_ctr <= psa.end_page {
				fmt.Fprintf(fout, "%s", line)
			}
		}
	} else {
		page_ctr = 1
		for true {
			c, err3 := rd.ReadByte()
			if err3 != nil { /* error or EOF */
				break
			}
			if c == '\f' {
				page_ctr++
			}
			if page_ctr >= psa.start_page && page_ctr <= psa.end_page {
				fmt.Fprintf(fout, "%c", c)
			}
		}
		fmt.Print("\n")
	}

	/* end main loop */
	if page_ctr < psa.start_page {
		fmt.Fprintf(os.Stderr, "%s: 起始页 (%d) 比总页数大 (%d)\n", progname, psa.start_page, page_ctr)
	} else if page_ctr < psa.end_page {
			fmt.Fprintf(os.Stderr, "%s: 终止页 (%d) 比总页数大 (%d)\n", progname, psa.end_page, page_ctr)
	}
	
	if psa.print_dest != "" {
		stdin.Close()
		cmd.Stdout = fout
		cmd.Run()
	}
	fmt.Fprintf(os.Stderr,"\n---------------\n程序结束\n")
	fin.Close()
	fout.Close()
}

func usage() {
	fmt.Fprintf(os.Stderr,"Usage error!\n")
	fmt.Fprintf(os.Stderr,"Usage:")
	fmt.Fprintf(os.Stderr,"\tselpg -s number -e number [options] [filename]\n\n")
	fmt.Fprintf(os.Stderr,"\t-s=number\t开始页数(开始<=结束)\n")
	fmt.Fprintf(os.Stderr,"\t-e=number\t结束页数(开始<=结束)\n")
	fmt.Fprintf(os.Stderr,"\t-l=number\t每页行数(可选)，默认72\n")
	fmt.Fprintf(os.Stderr,"\t-f\t\t是否用换页符来换页(可选)\n")
	fmt.Fprintf(os.Stderr,"\t[filename]\t从文件读，省略为标准输入\n\n")
}

func main() {
	sa := selpg_args{0, 0, "", 15, false, ""}
	progname = os.Args[0]
	Parser(&sa)
	processArgs(len(os.Args), &sa)
	processInput(&sa)
}
