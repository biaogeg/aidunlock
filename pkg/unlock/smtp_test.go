package unlock

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

func Test_SMTP(t *testing.T) {
	viper.SetConfigFile("/Users/mritd/tmp/aidunlock.yaml")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	var smtp SMTPConfig
	viper.UnmarshalKey("email", &smtp)
	smtp.Send(fmt.Sprintf("Apple ID [%s] Password Reset Success!\n\nNew Password: %s\n", "aaa", "bbb"))
}
