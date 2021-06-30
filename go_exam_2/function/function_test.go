package function

import "testing"

func TestValidateThailandIDLength(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Length < 13",
			args: args{
				idNo: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Length > 13",
			args: args{
				idNo: "12345678901234567890",
			},
			wantErr: true,
		},
		{
			name: "Length == 13",
			args: args{
				idNo: "1234567891234",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandIDLength(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandIDLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateThailandIDNumericOnly(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Alphabet_Only",
			args: args{
				idNo: "AsadasdjasdiasdwiqdjwqEiEi",
			},
			wantErr: true,
		},
		{
			name: "Alphabet_Numeric",
			args: args{
				idNo: "123ASD456789",
			},
			wantErr: true,
		},
		{
			name: "Numeric_Only",
			args: args{
				idNo: "123456789",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandIDNumericOnly(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandIDNumericOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateThailandIDPattern(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Logic Incorrect",
			args: args{
				idNo: "9999999999999",
			},
			wantErr: true,
		},
		{
			name: "Logic Correct",
			args: args{
				idNo: "3103109929515",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandIDPattern(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandIDPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateThailandCitizenID(t *testing.T) {
	type args struct {
		idNo string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "10 Digit",
			args:    args{idNo: "1111111111"},
			wantErr: true,
		},
		{
			name:    "14 Digit",
			args:    args{idNo: "11111111111111"},
			wantErr: true,
		},
		{
			name:    "Text and Numeric",
			args:    args{idNo: "AA11111111111"},
			wantErr: true,
		},
		{
			name:    "Text",
			args:    args{idNo: "AAAAAAAAAAAAA"},
			wantErr: true,
		},
		{
			name:    "13 Digit/ Correct format#1",
			args:    args{idNo: "4953714647109"},
			wantErr: false,
		},
		{
			name:    "13 Digit/ Correct format#2",
			args:    args{idNo: "6226429459176"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateThailandCitizenID(tt.args.idNo); (err != nil) != tt.wantErr {
				t.Errorf("ValidateThailandCitizenID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}