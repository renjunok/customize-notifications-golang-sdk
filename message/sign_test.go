/*
@Time : 2021/5/7 上午2:55
@Author : RenJun
@File : sign_test
@Description :
@CopyRight:
*/
package message

import "testing"

func TestParamSign_Sign(t *testing.T) {
	type args struct {
		pushSecret string
	}
	tests := []struct {
		name    string
		p       ParamSign
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.sign(tt.args.pushSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}