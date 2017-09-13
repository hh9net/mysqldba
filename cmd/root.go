// Copyright © 2017 Hong Bin <hongbin@actionsky.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

const (
	globalStatusSQL = `show global status where Variable_name in ('Uptime','Com_commit','Com_rollback','Queries','Com_select','Com_insert','Com_update','Com_delete',
	       'Threads_running','Threads_connected','Threads_created','Innodb_buffer_pool_pages_data','Innodb_buffer_pool_pages_free',
	       'Innodb_buffer_pool_pages_dirty', 'Innodb_buffer_pool_pages_flushed','Innodb_rows_read','Innodb_rows_inserted','Innodb_rows_updated',
	       'Innodb_rows_deleted','Innodb_buffer_pool_read_requests','Innodb_buffer_pool_reads','Innodb_os_log_fsyncs','Innodb_os_log_written',
		   'Innodb_log_waits','Innodb_log_write_requests','Innodb_log_writes','Bytes_sent','Bytes_received')`

	innodbStatusSQL   = "SHOW ENGINE INNODB STATUS"
	scroll            = 40
	version           = "1.0"
	globalVariableSQL = "show global variables"
	slaveStatusSQL    = "show slave status"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "mysqlmon",
	Short: "Welcome to the MySQL Toolbox.",
	Long:  "Welcome to the MySQL Toolbox. \n\nVersion: " + version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

var (
	dbUser, dbPassWd, dbHost, dbSocket, dsn string
	dbPort                                  int
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&dbUser, "user", "u", "root", "mysql login user")
	RootCmd.PersistentFlags().StringVarP(&dbPassWd, "password", "p", "root", "mysql login password")
	RootCmd.PersistentFlags().StringVarP(&dbHost, "host", "H", "localhost", "mysql host ip")
	RootCmd.PersistentFlags().IntVarP(&dbPort, "port", "P", 3306, "mysql server port")

}

func byteHumen(i int64) aurora.Value {

	switch {

	case i >= 1024 && i < 1048576:
		return aurora.Green(strconv.FormatInt(i/1024, 10) + "K")

	case i >= 1048576 && i < 1073741824:
		return aurora.Brown(strconv.FormatInt(i/1048576, 10) + "M")

	case i >= 1073741824:
		return aurora.Red(strconv.FormatInt(i/1073741824, 10) + "G")

	default:
		return aurora.Green(strconv.FormatInt(i, 10) + "B")

	}

}