package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"unit_test/blog/pkg/blog"
	mblog "unit_test/blog/test/mocks/blog"
)

func TestListPosts(t *testing.T) {
	ctrl :=gomock.NewController(t)
	defer ctrl.Finish()

	mockBlog:= mblog.NewMockBlog(ctrl)
	mockBlog.EXPECT().ListPosts().Return([]blog.Post{})


	service:=NewService(mockBlog)
	result,err := service.ListPosts()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, []blog.Post{},result)
}