package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaskDirPath(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "ファイル名あり",
			filename: "file1",
			want:     "templates/tasks/file1.gtpl",
		},
		{
			name:     "ファイル名が空",
			filename: "",
			want:     "templates/tasks/.gtpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetTaskDirPath(tt.filename)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetHeaderTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "成功",
			want: "templates/_header.gtpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetHeaderTemplate()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetFooterTemplate(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "成功",
			want: "templates/_footer.gtpl",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFooterTemplate()
			assert.Equal(t, tt.want, got)
		})
	}
}
